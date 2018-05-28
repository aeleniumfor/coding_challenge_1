# golang version
1.10.2

# docker
### 使用したdocker image
```docker
golang:latest
```

### ビルドと実行
```docker:docker_file
$ docker build -t golang_app .
$ docker run -p 8080:8080 --rm golang_app
```
