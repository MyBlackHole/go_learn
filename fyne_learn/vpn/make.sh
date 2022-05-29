# 用于交叉编译
# GOOS=linux GOARCH=amd64 go build -ldflags="-r ./lib/" -o release/linux/amd64/bin/URun-vpn
# GOOS=windows GOARCH=amd64 go build -ldflags="-r ./lib/" -o release/windows/amd64/URun-vpn.exe
CC=x86_64-w64-mingw32-gcc CGO_ENABLED=1 GOOS=windows go build -o URun-vpn.exe -ldflags="-H windowsgui"
# GOOS=windows GOARCH=386 go build -o release/windows/386/URun-vpn.exe
