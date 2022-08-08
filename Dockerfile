FROM golang:1.18.3-alpine3.16

RUN apk update && apk upgrade && \
    apk --update add git make

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY ./ ./
RUN make builded-order-service

CMD [ "./bin/order-service"]