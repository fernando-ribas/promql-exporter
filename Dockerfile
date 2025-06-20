FROM golang:1.22 as builder

WORKDIR /app
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o promql_exporter .

FROM alpine:3.19
WORKDIR /
COPY --from=builder /app/promql_exporter .

EXPOSE 9517
ENTRYPOINT ["/promql_exporter"]
