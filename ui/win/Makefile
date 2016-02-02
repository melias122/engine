all: 386 amd64

386:
	GOOS=windows GOARCH=386 go build -a -ldflags="-H windowsgui" -o Generator_386_v1.exe

amd64:
	GOOS=windows GOARCH=amd64 go build -a -ldflags="-H windowsgui" -o Generator_amd64_v1.exe

clean:
	rm ./*.exe