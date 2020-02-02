# https://app.getpocket.com/read/2871142125

FROM golang:1.13-alpine as builder
RUN apk add make
RUN mkdir /build 
ADD . /build/
WORKDIR /build
# Set environment variable here to avoid error 
# that happens when running go test in builder container.
# https://github.com/golang/go/issues/27303
ENV CGO_ENABLED 0
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o server ./cmd/server.go

FROM scratch
COPY --from=builder /build/server /app/
WORKDIR /app
CMD ["./server"]
