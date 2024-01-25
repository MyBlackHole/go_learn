go run main.go server

curl 127.0.0.1:9000/foo -v
*   Trying 127.0.0.1:9000...
* Connected to 127.0.0.1 (127.0.0.1) port 9000 (#0)
> GET /foo HTTP/1.1
> Host: 127.0.0.1:9000
> User-Agent: curl/7.88.1
> Accept: */*
>
< HTTP/1.1 200 OK
< Access-Control-Allow-Origin: *
< Date: Thu, 18 Jan 2024 04:08:20 GMT
< Content-Length: 3
< Content-Type: text/plain; charset=utf-8
<
* Connection #0 to host 127.0.0.1 left intact
foo%


❯ curl 127.0.0.1:9000/ -v
*   Trying 127.0.0.1:9000...
* Connected to 127.0.0.1 (127.0.0.1) port 9000 (#0)
> GET / HTTP/1.1
> Host: 127.0.0.1:9000
> User-Agent: curl/7.88.1
> Accept: */*
>
< HTTP/1.1 200 OK
< Access-Control-Allow-Origin: *
< Date: Thu, 18 Jan 2024 06:40:52 GMT
< Content-Length: 3
< Content-Type: text/plain; charset=utf-8
<
* Connection #0 to host 127.0.0.1 left intact
foo%


# 静态编译
CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w"
# 提供更全的 debug 信息
go build -gcflags="all=-N -l"

# debug
sudo /home/black/go/bin/dlv exec ./emulator -- server --port 80
./ossutil -e 127.0.0.1 -i 1234 -k 1234 --loglevel=debug mb oss://wdg1
./ossutil -e 192.168.63.183 -i 1234 -k 1234 mb oss://ob-1-1921686062-1705309643

config max-string-len 99999
break ListObjectsHandler
break HeadObjectHandler
break AppendObjectHandler
break PutObjectHandler
break errorResponseHandler

byebug /opt/aio/airflow/tools/oss-emulator/bin/emulator -r /volmountpoint/aiopool/aio_s3 -p 80
break /opt/aio/airflow/tools/oss-emulator/lib/emulator/object.rb:20

cd /opt/aio/airflow/tools/oss-emulator/bin/ &&  ruby ./emulator -r /volmountpoint/aiopool/aio_s3 -p 80 -L debug &
