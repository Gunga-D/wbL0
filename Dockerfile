FROM golang:1.18.3-buster

ENV GOPATH=/

RUN apt-get update && apt-get install -y \
    make \
    build-essential \
    postgresql-client

COPY go.mod ./
RUN go mod download

COPY ./ ./

RUN chmod +x wait-for-postgres.sh

RUN make ord-serv

CMD ["./order-service"]