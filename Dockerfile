FROM golang:latest
WORKDIR /data/ppim
COPY . .
RUN go env -w GOPROXY=https://goproxy.cn,direct \
 && go run main.go
