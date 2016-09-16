default:
	go build -v

bins: osx_386 osx_amd64 linux_386 linux_amd64

osx_386:
	GOARCH=386 GOOS=darwin go build -o generate-docker-certs.386.osx

osx_amd64:
	GOARCH=amd64 GOOS=darwin go build -o generate-docker-certs.amd64.osx

linux_amd64:
	GOARCH=amd64 GOOS=linux go build -o generate-docker-certs.amd64.linux

linux_386:
	GOARCH=386 GOOS=linux go build -o generate-docker-certs.386.linux
