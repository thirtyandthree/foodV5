.PHONU: all build run clean help

FILE="foodV5-api"

build:
	@go build -o ${FILE} main.go wire_gen.go

windows-amd64:
	@CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc CXX=x86_64-w64-mingw32-g++ GOOS=windows GOARCH=amd64 go build -o ${FILE}.exe main.go wire_gen.go

windows-386:
	@CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc CXX=x86_64-w64-mingw32-g++ GOOS=windows GOARCH=386 go build -o ${FILE}.exe main.go wire_gen.go

linux-amd64:
	@GOOS=linux GOARCH=amd64 go build -o ${FILE} main.go wire_gen.go

linux-386:
	@GOOS=linux GOARCH=386 go build -o ${FILE} main.go wire_gen.go

run:
	@go run main.go wire_gen.go -c config.yaml

run-hot:
	@realize s

clean:
	@if [ -f ${FILE} ]; then
		rm -rf ${FILE}
	fi

help:
	@echo "make build - 编译成二进制文件"
	@echo "make run - 运行服务"
	@echo "make clean - 清除编译后的文件"
	@echo "make linux-amd64 - 生成amd架构可执行文件"