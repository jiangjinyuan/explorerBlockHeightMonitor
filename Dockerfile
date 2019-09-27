FROM golang:alpine AS build-env
RUN apk --no-cache add git
ADD . /go/src/github.com/jiangjinyuan/explorerBlockHeightMonitor
RUN cd /go/src/github.com/jiangjinyuan/explorerBlockHeightMonitor && \
   go mod download && \
   go build -v -o /src/bin/explorerBlockHeightMonitor main.go

FROM alpine
RUN apk --no-cache add openssl ca-certificates tzdata
WORKDIR /app
COPY --from=build-env /src/bin /app/
COPY --from=build-env /go/src/github.com/jiangjinyuan/explorerBlockHeightMonitor/configs /app/configs
COPY --from=build-env /go/src/github.com/jiangjinyuan/explorerBlockHeightMonitor/logs /app/logs
ENTRYPOINT ./explorerBlockHeightMonitor