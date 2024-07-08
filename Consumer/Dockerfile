FROM golang:1.22 AS builder

WORKDIR /build

ADD go.mod .

COPY . .

RUN go mod tidy
RUN go build -o src

FROM golang:1.22

WORKDIR /src

COPY --from=builder /build/src /src/app

CMD ["./app"]