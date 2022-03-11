FROM golang:1.17-alpine

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . ./

RUN go build -o bin/proddx-server cmd/server/*.go

EXPOSE 5000

RUN go build -o bin/migration cmd/migration/*.go

CMD [ "bin/migration; bin/proddx-server" ]
