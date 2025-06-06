package watermark

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

// 零宽字符
const (
	ZeroWidthSpace     = "\u200B" // 零宽空格
	ZeroWidthNonJoiner = "\u200C" // 零宽非连接符
	ZeroWidthJoiner    = "\u200D" // 零宽连接符
)

// Encode 将水印文本编码到原始文本中
func Encode(originalText, watermark string) string {
	// 将水印转换为二进制
	watermarkBytes := []byte(watermark)
	var binaryStr string
	for _, b := range watermarkBytes {
		binaryStr += fmt.Sprintf("%08b", b)
	}

	// 将二进制字符串转换为零宽字符
	var encodedWatermark string
	for _, bit := range binaryStr {
		if bit == '1' {
			encodedWatermark += ZeroWidthSpace
		} else {
			encodedWatermark += ZeroWidthNonJoiner
		}
	}

	// 在原始文本中插入水印
	return originalText + ZeroWidthJoiner + encodedWatermark
}

// Decode 从文本中提取水印
func Decode(text string) string {
	// 查找零宽连接符的位置
	joinerIndex := -1
	for i := 0; i < len(text); {
		r, size := utf8.DecodeRuneInString(text[i:])
		if r == '\u200D' {
			joinerIndex = i
			break
		}
		i += size
	}

	if joinerIndex == -1 {
		return ""
	}

	// 提取水印部分
	watermarkPart := text[joinerIndex+3:] // 跳过零宽连接符

	// 将零宽字符转换回二进制
	var binaryStr string
	for i := 0; i < len(watermarkPart); {
		r, size := utf8.DecodeRuneInString(watermarkPart[i:])
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
