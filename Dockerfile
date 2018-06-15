FROM golang:alpine AS build
WORKDIR /go/src/github.com/cuigh/swirl-agent/
ADD . .
#RUN dep ensure
RUN CGO_ENABLED=0 go build -ldflags "-s -w"

FROM alpine:3.7
LABEL maintainer="cuigh <noname@live.com>"
WORKDIR /app
COPY --from=build /go/src/github.com/cuigh/swirl-agent/swirl-agent .
COPY --from=build /go/src/github.com/cuigh/swirl-agent/config ./config/
EXPOSE 8002
ENTRYPOINT ["/app/swirl-agent"]
