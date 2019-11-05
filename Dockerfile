FROM golang:alpine as builder
WORKDIR /src
COPY ["main.go", "./"]
RUN go build -tags netgo -ldflags '-s -w' -o app

FROM scratch
LABEL maintainer=cornelius.weig@gmail.com
ENTRYPOINT ["/app"]
EXPOSE 80
COPY --from=builder ["/src/app", "/app"]

