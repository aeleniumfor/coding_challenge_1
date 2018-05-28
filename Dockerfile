FROM golang:latest

WORKDIR /go/src/coding_challenge_1
COPY ./ ./
RUN pwd
RUN go install -v ./
CMD ["coding_challenge_1"]