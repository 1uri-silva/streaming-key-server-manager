FROM golang:1.22 as build

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o stream-key-manager ./cmd/server/main.go

FROM scratch

COPY --from=build /app/stream-key-manager /stream-key-manager

CMD ["/stream-key-manager"]

EXPOSE 8000