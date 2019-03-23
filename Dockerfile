FROM golang:1.11 as builder

WORKDIR /box
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY main.go .
COPY controllers ./controllers
COPY logic ./logic
COPY routers ./routers

RUN CGO_ENABLED="0" go build

FROM scratch

COPY --from=builder /box/vin .
COPY conf conf

EXPOSE 8095

ENTRYPOINT [ "./vin" ]