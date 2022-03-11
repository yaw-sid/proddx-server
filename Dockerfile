FROM golang:1.16-alpine

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . ./

RUN go build -o bin/proddx-server cmd/server/*.go

EXPOSE 5000

CMD [ "bin/proddx-server" ]
