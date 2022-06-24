# 用于交叉编译
# CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o release/mac/amd64/bin
# CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o release/linux/amd64/bin
# CGO_ENABLED=0 GOOS=freebsd GOARCH=amd64 go build -o release/freebsd/amd64/bin
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o release/windows/amd64/URun-time_judgment.exe
CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -o release/windows/386/URun-time_judgment.exe