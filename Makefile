VERSION?="1.0"

bin:
	go get github.com/rakyll/statik
	statik -include=laugh.mp3 -src=./
	bash build.sh

.PHONY: bin
