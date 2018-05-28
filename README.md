# Version
### golang
```
$ go verison
# go version go1.10.2 linux/amd64
```

### docker
```
$ docker -v
# Docker version 18.05.0-ce, build f150324
```

# Image

### 使用したdocker image
```docker
golang:latest
```


# Usage
### build and run
```docker:docker_file
$ docker build -t golang_app .
$ docker run -d -p 8080:8080 --rm golang_app
```
### confirmation

##### ステータスコードの確認
```
$ curl -LI http://localhost:8080/ -o /dev/null -w '%{http_code}' -s
```

##### レスポンスの確認
```
$ curl -XGET -H 'Content-Type:application/json' http://localhost:8080/
```
