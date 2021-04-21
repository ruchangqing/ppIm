FROM golang:1.16.3
WORKDIR /go/src/ppim
COPY . .
RUN go env -w GOPROXY=https://goproxy.cn,direct \
 && go get -d -v \
 && go install -v \