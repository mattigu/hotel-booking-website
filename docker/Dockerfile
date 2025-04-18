FROM golang:1.24-alpine

WORKDIR /app


COPY go.* ./


COPY ./src/*.go ./
COPY . .

RUN go install github.com/air-verse/air@latest
RUN go get github.com/jackc/pgx/v5@latest
RUN go get github.com/joho/godotenv@latest
#RUN go get github.com/jackc/pgx/v5/pgxpool
#RUN go mod tidy
RUN go mod download

EXPOSE 8888

CMD ["air", "-c", ".air.toml"]