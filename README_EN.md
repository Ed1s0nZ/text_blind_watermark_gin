# Text Blind Watermark Tool

A text watermarking tool based on zero-width characters that supports adding and extracting watermarks from text. The tool provides both a web interface and command-line interface.

## Features

- Support for Chinese and English watermarks
- Uses zero-width characters for invisible watermarking
- Beautiful and user-friendly web interface
- Command-line support for automation
- Robust watermarking that can survive partial text deletion
- Watermark information is distributed throughout the text

## Installation

1. Ensure you have Go environment installed (version 1.16 or higher)
2. Clone the repository
3. Install dependencies:
```bash
go mod tidy
```

## Usage

### Web Interface

1. Start the web server:
```bash
go run cmd/server/main.go
```

2. Open your browser and visit: http://localhost:8080

### Command Line Tool

1. Build the command-line tool:
```bash
go build -o watermark cmd/cli/main.go
```

2. Add watermark:
```bash
./watermark encode input.txt "watermark content" output.txt
```

3. Extract watermark:
```bash
./watermark decode input.txt output.txt
```

## Example

1. Create a test file:
```bash
echo "This is a test text" > test.txt
```

2. Add watermark:
```bash
./watermark encode test.txt "test watermark" test_with_watermark.txt
```

3. Extract watermark:
```bash
./watermark decode test_with_watermark.txt extracted_watermark.txt
```

## How It Works

The tool uses zero-width characters to embed watermark information:
- Zero-width space (\u200B)
- Zero-width non-joiner (\u200C)
- Zero-width joiner (\u200D)

The watermark information is distributed throughout the text, making it more robust against partial text deletion. Even if some parts of the text are removed, the remaining watermark information can still be extracted.

## Notes

- The watermark is invisible to the naked eye
- When copying and pasting text, it's recommended to use a plain text editor to avoid format loss
- The watermark can survive partial text deletion, but complete deletion of the text will result in watermark loss
- The tool is designed for watermarking purposes only and does not provide data integrity verification

## License

This project is open source and available under the MIT License. 