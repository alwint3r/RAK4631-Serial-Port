
linux:
	GOOS=linux GOARCH=amd64 go build -o ./bin/linux/linux-amd64/rak4631autoserial-linux-amd64
	GOOS=linux GOARCH=arm64 go build -o ./bin/linux/linux-arm64/rak4631autoserial-linux-arm64
	GOOS=linux GOARCH=arm go build -o ./bin/linux/linux-arm/rak4631autoserial-linux-arm
	GOOS=linux GOARCH=386 go build -o ./bin/linux/linux-386/rak4631autoserial-linux-i386

windows:
	GOOS=windows GOARCH=amd64 go build -o ./bin/windows/windows-amd64/rak4631autoserial-win-amd64.exe
	GOOS=windows GOARCH=386 go build -o ./bin/windows/windows-386/rak4631autoserial-win-i386.exe

macos:
	GOOS=darwin GOARCH=amd64 go build -o ./bin/macos/macos-amd64/rak4631autoserial-macos-amd64

clean:
	rm -rf bin/*