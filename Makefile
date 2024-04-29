build:
	go build -o bin/gnp cmd/gnp/main.go

compile:
	echo "Compiling for every OS and Platform"
	GOOS=linux GOARCH=arm go build -o bin/gnp-linux-arm cmd/gnp/main.go
	GOOS=linux GOARCH=arm64 go build -o bin/gnp-linux-arm64 cmd/gnp/main.go
	GOOS=freebsd GOARCH=386 go build -o bin/gnp-freebsd-386 cmd/gnp/main.go
	GOOS=linux GOARCH=amd64 go build -o bin/gnp-linux-amd64 cmd/gnp/main.go

clean:
	rm -rf bin/*

