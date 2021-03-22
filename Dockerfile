FROM golang:1.15 as build
WORKDIR /go/src/github.com/st-vasyl/echo/
COPY . .
RUN go get -d -v github.com/sirupsen/logrus && \
    echo $GOPATH $GOROOT && \
    CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -a -installsuffix cgo -o echo .

FROM alpine:3.9
LABEL MAINTAINER "Vasyl Stetsuryn <vasyl@vasyl.org>"

ARG APK_FLAGS_COMMON="-q"
ARG APK_FLAGS_PERSISTANT="${APK_FLAGS_COMMON} --clean-protected --no-cache"

ENV LANG C.UTF-8
ENV TERM=xterm
USER root

RUN apk update && \
    apk add ${APK_FLAGS_PERSISTANT} \
            less \
            bash && \
    addgroup echo && \
    adduser -u 1000 \
            -S \
            -D -G echo \
            -h /home/echo \
            -s /bin/bash \
            echo && \
    mkdir -p /opt/echo && \
    chown -R echo:echo /opt/echo

COPY --from=build /go/src/github.com/st-vasyl/echo/echo /opt/echo/echo
USER echo
WORKDIR /opt/echo


CMD ["/opt/echo/echo"]
