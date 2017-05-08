# redirect-server

Performs a DNS lookup and redirect the user to your specified domain.

## [Docker](https://www.docker.com/)

`redirect-server` is ready to run on docker. Official images can be found on [peleteiro/redirect-server](https://hub.docker.com/r/peleteiro/redirect-server/).

Run `docker run -p 8080:8080 peleteiro/redirect-server` on terminal to start `redirect-server` locally.

Test it locally using:

```shell
$ curl -i --header 'Host: foo.com' localhost:8080/path/file.jpg

HTTP/1.1 301 Moved Permanently
Location: https://www.foo.com/path/file.jpg
Date: Wed, 08 Jun 2016 20:38:48 GMT
Content-Length: 0
Content-Type: text/plain; charset=utf-8
```
