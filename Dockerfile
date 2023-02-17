FROM golang:1.19.6-alpine3.17 AS build

WORKDIR /go/src/numbers_reservation
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/numbers_reservation cmd/numbers_reservation.go

FROM scratch

ENV GOPROXY=https://proxy.golang.org
ENV GIN_MODE=release

COPY --from=build /go/bin/numbers_reservation /app

ENTRYPOINT ["/app"]
