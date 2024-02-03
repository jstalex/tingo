FROM golang:alpine AS build

WORKDIR /app

COPY . .
RUN go build -o main cmd/backtest/main.go


FROM scratch

COPY --from=build app/main main
COPY --from=build app/config.yaml config.yaml
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENTRYPOINT ["/main"]