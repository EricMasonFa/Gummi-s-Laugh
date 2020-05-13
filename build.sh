for OS in darwin linux windows; do
    for ARCH in amd64; do
        CGO_ENABLED=1 GOOS=$OS GOARCH=$ARCH go build -v -o bin/gummi-$OS-$ARCH
    done
done
