# 用于交叉编译
# GOOS=linux GOARCH=amd64 go build -o test
# GOOS=windows GOARCH=amd64 go build -ldflags="-r ./lib/" -o release/windows/amd64/URun-vpn.exe
GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc CXX=x86_64-w64-mingw32-g++ go build -x ./
# GOOS=windows GOARCH=386 go build -o release/windows/386/URun-vpn.exe
