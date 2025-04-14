FROM golang:1.23-alpine

WORKDIR /app

RUN go install github.com/air-verse/air@latest

COPY go.* ./

RUN go mod download

COPY ./src/*.go ./
COPY . .

# RUN go build -o main main.go

EXPOSE 8888

#CMD ["./main"]
CMD ["air", "-c", ".air.toml"]