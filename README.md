# 文本暗水印工具

一个基于零宽字符的文本暗水印工具，支持中文水印的添加和提取。该工具提供了Web界面和命令行两种使用方式。   
参考：https://github.com/guofei9987/text_blind_watermark   

## 功能特点

- 支持中文水印
- 使用零宽字符实现，不影响文本显示
- 提供美观的Web界面
- 支持命令行操作
- 加解密稳定可靠

## 安装

1. 确保已安装Go环境（1.16或更高版本）
2. 克隆项目
3. 安装依赖：
```bash
go mod tidy
```

## 使用方法

### Web界面

1. 启动Web服务器：
```bash
go run cmd/server/main.go
```
2. 在浏览器中访问：http://localhost:8080
   
<img src="https://github.com/Ed1s0nZ/text_blind_watermark_gin/blob/main/image/web.png">  

### 命令行工具

1. 编译命令行工具：
```bash
go build -o watermark cmd/cli/main.go
```

2. 添加水印：
```bash
./watermark encode input.txt "水印内容" output.txt
```

3. 提取水印：
```bash
./watermark decode input.txt output.txt
```

## 示例

1. 创建一个测试文件：
```bash
echo "这是一段测试文本" > test.txt
```

2. 添加水印：
```bash
./watermark encode test.txt "测试水印" test_with_watermark.txt
```

3. 提取水印：
```bash
./watermark decode test_with_watermark.txt extracted_watermark.txt
```

## 注意事项

- 水印内容支持中文和英文
- 添加水印后的文本看起来与原文完全相同
- 建议在复制粘贴文本时使用纯文本编辑器，以避免格式丢失 
