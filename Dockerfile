FROM golang:1.23 as build
WORKDIR /go/src/github.com/st-vasyl/echo-server/
COPY . .

ARG GIT_BRANCH
ARG GIT_HASH
ARG APP_VERSION

ENV BUILD_FLAGS="-X 'main.Branch=${GIT_BRANCH}' -X 'main.CommitHash=${GIT_HASH}' -X 'main.Version=${APP_VERSION}'"

RUN go mod tidy && \
    CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -a -installsuffix cgo -o echo -ldflags "${BUILD_FLAGS}" .

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
            -h /home/echo-server \
            -s /bin/bash \
            echo && \
    mkdir -p /opt/echo-server && \
    chown -R echo:echo /opt/echo-server

COPY --from=build /go/src/github.com/st-vasyl/echo-server/echo /opt/echo-server/echo
USER echo
WORKDIR /opt/echo-server

LABEL Name="${GIT_REPO}" \
    Version="${APP_VERSION}"

CMD ["/opt/echo-server/echo"]
