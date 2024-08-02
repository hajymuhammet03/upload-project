FROM golang:1.22

RUN apt-get update && apt-get install -y \
    libwebp-dev \
    libde265-dev \
    && apt-get clean

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

WORKDIR /app/cmd/main

RUN go build -o /app/app .

CMD ["/app/app"]
