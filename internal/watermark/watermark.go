package watermark

import (
	"fmt"
	"net"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/gin-gonic/gin"
)

// 零宽字符
const (
	ZeroWidthSpace     = "\u200B" // 零宽空格
	ZeroWidthNonJoiner = "\u200C" // 零宽非连接符
	ZeroWidthJoiner    = "\u200D" // 零宽连接符
)

// Encode 将水印文本编码到原始文本中
func Encode(originalText, watermark, ipAddress string) string {
	// 将水印转换为二进制
	encodedWatermark := strToEncoded(watermark)

	encodedIpAddress := ""

	if ipAddress != "" {
		encodedIpAddress = ZeroWidthJoiner + strToEncoded(ipAddress)
	}

	// 在原始文本中插入水印
	return originalText + ZeroWidthJoiner + encodedWatermark + encodedIpAddress
}

// Decode 从文本中提取水印
func Decode(text string, c *gin.Context) string {
	// 查找零宽连接符的位置
	joinerIndexList := make([]int, 0)
	for i := 0; i < len(text); {
		r, size := utf8.DecodeRuneInString(text[i:])
		if r == '\u200D' {
			joinerIndexList = append(joinerIndexList, i)
		}
		i += size
	}
	if len(joinerIndexList) == 0 {
		return ""
	}

	if len(joinerIndexList) == 2 {
		ipAddress := text[joinerIndexList[1]+3:]
		ip := strByDecode(ipAddress)
		clientIp := getClientIP(c)
		if ip != clientIp {
			return ""
		}
	}
	watermarkPart := getWatermark(text, joinerIndexList)
	return strByDecode(watermarkPart)
}

func getWatermark(text string, joinerIndexList []int) string {
	// 提取水印部分
	if len(joinerIndexList) == 1 {
		// 跳过零宽连接符
		return text[joinerIndexList[0]+3:]
	} else {
		// 跳过零宽连接符
		return text[joinerIndexList[0]+3 : joinerIndexList[1]]
	}
}

func getClientIP(c *gin.Context) string {
	// 1. 检查X-Forwarded-For头部
	xForwardedFor := c.GetHeader("X-Forwarded-For")
	if xForwardedFor != "" {
		// 可能有多个IP，第一个是原始客户端IP
		ips := strings.Split(xForwardedFor, ",")
		return strings.TrimSpace(ips[0])
	}

	// 2. 检查X-Real-Ip头部
	xRealIP := c.GetHeader("X-Real-Ip")
	if xRealIP != "" {
		return xRealIP
	}

	// 3. 最后回退到RemoteAddr
	ip, _, err := net.SplitHostPort(c.Request.RemoteAddr)
	if err != nil {
		return c.Request.RemoteAddr
	}

	// 最后使用 Gin 的 ClientIP()
	clientIp := c.ClientIP()
	if clientIp == "::1" {
		return "127.0.0.1"
	}

	return ip
}

func strByDecode(str string) string {
	// 将零宽字符转换回二进制
	var binaryStr string
	for i := 0; i < len(str); {
		r, size := utf8.DecodeRuneInString(str[i:])
		if r == '\u200B' {
			binaryStr += "1"
		} else if r == '\u200C' {
			binaryStr += "0"
		}
		i += size
	}

	// 将二进制字符串转换回字节
	var bytes []byte
	for i := 0; i < len(binaryStr); i += 8 {
		if i+8 <= len(binaryStr) {
			byteStr := binaryStr[i : i+8]
			b, _ := strconv.ParseUint(byteStr, 2, 8)
			bytes = append(bytes, byte(b))
		}
	}

	return string(bytes)
}

func strToEncoded(str string) string {
	strBytes := []byte(str)
	var binaryStr string
	for _, b := range strBytes {
		binaryStr += fmt.Sprintf("%08b", b)
	}

	// 将二进制字符串转换为零宽字符
	var encodedStr string
	for _, bit := range binaryStr {
		if bit == '1' {
			encodedStr += ZeroWidthSpace
		} else {
			encodedStr += ZeroWidthNonJoiner
		}
	}
	return encodedStr
}
