# bash
export PATH=$PATH:/usr/local/go/bin
go build -buildmode=c-shared -o lib/ela.so lib/ela.go