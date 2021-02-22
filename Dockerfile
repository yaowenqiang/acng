FROM golang:1.8.1-alpine

RUn apk update && apk upgrade && apk add --no-cache bash git
RUN go get github.com/gin-gonic/gin
ENV SOURCES /go/src/github.com/yaowenqiang/acng
COPY . ${SOURCES}

RUN cd ${SOURCES} && CGO_ENABLED=0 go build
workdir ${SOURCES}
CMD ${SOURCES}Gin-web
EXPOSE 8080
