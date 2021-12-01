FROM golang:1.17

WORKDIR /currencies-go-producer

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o /currencies-go-producer

CMD ["/currencies-go-producer"]

