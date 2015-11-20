all: 386 amd64

386:
	GOOS=windows GOARCH=386 go build -ldflags="-H windowsgui" -o Generator_386.exe

amd64:
	GOOS=windows GOARCH=amd64 go build -ldflags="-H windowsgui" -o Generator_amd64.exe

clean:
	rm ./*.exe