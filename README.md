# Encoding Detection Tool

This is a simple console application written in Go that detects the encoding of a text file.

## Dependencies
This application uses the golang.org/x/net/html/charset library to perform encoding detection. You can download and install this library by running:

```bash
 go get golang.org/x/net/html/charset
```

## Getting Started

To run this application, you'll need to ensure you have Go installed on your system. You can download and install it from the official website: [https://golang.org/dl/](https://golang.org/dl/)

Once you have Go installed, you can download and run the application as follows:

1. Clone this repository or download the `EncodingDetect.go` file.

2. Open a terminal and navigate to the directory where the `EncodingDetect.go` file is located.

3. Run the program by executing the following command:

   ```bash
   go run EncodingDetect.go
   ```

You will be prompted to enter the name of the file you want to analyze. Provide the file's name or path.

The program will read the file and detect its encoding, then display the result.

## Disclaimer
This is a basic encoding detection tool and may not be 100% accurate. The tool detects common encodings and can be further customized based on your needs.
