FROM golang:1.16
LABEL maintainer="sam <sinloes@gmail.com>"
WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

ENV PORT 8080

RUN go build

CMD ["./go-crud-app.exe"]