FROM golang:latest as builder
WORKDIR /root
COPY go.mod ./
COPY main.go ./
RUN CGO_ENABLED=0 go build -o /app .

FROM alpine:latest
COPY --from=builder /app .
EXPOSE 8080
CMD /app
