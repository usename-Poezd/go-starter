FROM golang:1.20-alpine AS build
WORKDIR /src

RUN apk add --no-cache make
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
RUN go build -o /src/app ./cmd/app

EXPOSE 8000

CMD ["/src/app"]