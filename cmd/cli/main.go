package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"text_blind_watermark/internal/watermark"
)

func main() {
	// 定义命令行参数
	encodeCmd := flag.NewFlagSet("encode", flag.ExitOnError)
	decodeCmd := flag.NewFlagSet("decode", flag.ExitOnError)

	// 检查是否提供了子命令
	if len(os.Args) < 2 {
		fmt.Println("请使用 'encode' 或 'decode' 子命令")
		os.Exit(1)
	}

	// 根据子命令执行相应的操作
	switch os.Args[1] {
	case "encode":
		encodeCmd.Parse(os.Args[2:])
		if encodeCmd.NArg() < 3 {
			fmt.Println("用法: encode <输入文件> <水印内容> <输出文件>")
			os.Exit(1)
		}
		inputFile := encodeCmd.Arg(0)
		watermarkText := encodeCmd.Arg(1)
		outputFile := encodeCmd.Arg(2)

		// 读取输入文件
		content, err := ioutil.ReadFile(inputFile)
		if err != nil {
			fmt.Printf("读取输入文件失败: %v\n", err)
			os.Exit(1)
		}

		// 添加水印
		encoded := watermark.Encode(string(content), watermarkText, "")

		// 写入输出文件
		err = ioutil.WriteFile(outputFile, []byte(encoded), 0644)
		if err != nil {
			fmt.Printf("写入输出文件失败: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("水印已添加到文件: %s\n", outputFile)

	case "decode":
		decodeCmd.Parse(os.Args[2:])
		if decodeCmd.NArg() < 2 {
			fmt.Println("用法: decode <输入文件> <输出文件>")
			os.Exit(1)
		}
		inputFile := decodeCmd.Arg(0)
		outputFile := decodeCmd.Arg(1)

		// 读取输入文件
		content, err := ioutil.ReadFile(inputFile)
		if err != nil {
			fmt.Printf("读取输入文件失败: %v\n", err)
			os.Exit(1)
		}

		// 提取水印
		decoded := watermark.Decode(string(content), nil)

		// 写入输出文件
		err = ioutil.WriteFile(outputFile, []byte(decoded), 0644)
		if err != nil {
			fmt.Printf("写入输出文件失败: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("水印已提取到文件: %s\n", outputFile)

	default:
		fmt.Println("未知命令。请使用 'encode' 或 'decode'")
		os.Exit(1)
	}
}
