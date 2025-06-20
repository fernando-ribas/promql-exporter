FROM golang:1.22 as builder

WORKDIR /app
COPY . .

RUN go mod tidy
RUN go build -o promql_exporter .

FROM alpine
WORKDIR /
COPY --from=builder /app/promql_exporter .

ENTRYPOINT ["/promql_exporter"]
