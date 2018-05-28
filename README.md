# coding_challenge_1

# golang version
1.9.0

# docker
### 使用したdocker image
```docker
golang:latest
```

### docker file
```docker:docker_file
$ docker build -t golang_app .
$ docker run -p 80:8080 --rm golang_app
```
