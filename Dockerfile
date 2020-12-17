FROM golang:alpine AS builder
WORKDIR /build/
COPY . .
RUN go build -o server .

FROM alpine
WORKDIR /app/
COPY --from=builder /build/server .
ENTRYPOINT ["/app/server"]
