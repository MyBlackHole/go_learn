curl 127.0.0.1:8080/foo -v
*   Trying 127.0.0.1:8080...
* Connected to 127.0.0.1 (127.0.0.1) port 8080 (#0)
> GET /foo HTTP/1.1
> Host: 127.0.0.1:8080
> User-Agent: curl/7.88.1
> Accept: */*
>
< HTTP/1.1 200 OK
< Access-Control-Allow-Methods: GET,PUT,PATCH,OPTIONS
< Access-Control-Allow-Origin: *
< Date: Thu, 18 Jan 2024 02:34:29 GMT
< Content-Length: 3
< Content-Type: text/plain; charset=utf-8
<
* Connection #0 to host 127.0.0.1 left intact
foo%
