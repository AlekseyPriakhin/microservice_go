FROM golang:alpine AS builder

WORKDIR /build

ADD go.mod .

COPY . .

RUN CGO_ENABLED=0 go build -o src

FROM scratch

WORKDIR /src

COPY --from=builder /build/src /src/app

EXPOSE 5555

CMD ["./app"]