# Toy reverse proxy

## How it works

Run followings:

```shell
$ bash run-apis.sh
```

Three API servers is working, by vanila [http](https://golang.org/pkg/net/http/) at port 1385, [Echo](https://github.com/labstack/echo) at port 1387 and [Gin](https://github.com/gin-gonic/gin) at port 1389.

We put them behind a toy reverse proxy (load balancer?). 

this toy reverse proxy picks server one by one, by round-robin algorithm.

Run toy reverse proxy server as follows:

```shell
$ cd toy-reverse-proxy && go run main.go
```

When you go to reverse proxy server (at `localhost:1335/api`), you can see that pick server one by one as follows:

[![Image from Gyazo](https://i.gyazo.com/94e2a24b1be2596cae4c02b93d2ad57f.gif)](https://gyazo.com/94e2a24b1be2596cae4c02b93d2ad57f)
