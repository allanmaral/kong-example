FROM golang:1.21.5-alpine3.19 AS builder
WORKDIR /app
COPY . .
RUN go mod tidy
RUN GOOS=linux CGO_ENABLE=0 go build -ldflags="-w -s" -o server .

FROM scratch
COPY --from=builder /app/server .
CMD ["./server"]
