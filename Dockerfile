FROM golang:1.17-alpine

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . ./

RUN go build -o bin/proddx-server cmd/server/*.go

EXPOSE 5000

RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

RUN migrate -database ${DATABASE_URL} -path db/migrations up

CMD [ "bin/proddx-server" ]
