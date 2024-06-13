FROM golang AS builder

WORKDIR /build

ADD go.mod .

COPY . .

RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o src

FROM scratch

WORKDIR /src

COPY --from=builder /build/src /src/app

EXPOSE 5555

CMD ["./app"]