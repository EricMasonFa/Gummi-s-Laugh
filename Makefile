VERSION?="1.0"

bin:
	go get github.com/rakyll/statik
	statik -include=laugh.mp3 -src=./
	go mod download
	bash build.sh

.PHONY: bin
