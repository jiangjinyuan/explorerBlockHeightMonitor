FROM golang:alpine AS build-env
RUN apk --no-cache add git
ADD . /go/src/github.com/jiangjinyuan/blockHeightMonitor
RUN cd /go/src/github.com/jiangjinyuan/blockHeightMonitor && \
   go mod download && \
   go build -v -o /src/bin/blockHeightMonitor main.go

FROM alpine
RUN apk --no-cache add openssl ca-certificates tzdata
WORKDIR /app
COPY --from=build-env /src/bin /app/
COPY --from=build-env /go/src/github.com/jiangjinyuan/blockHeightMonitor/configs /app/configs
COPY --from=build-env /go/src/github.com/jiangjinyuan/blockHeightMonitor/logs /app/logs
ENTRYPOINT ./stratum-server-monitor