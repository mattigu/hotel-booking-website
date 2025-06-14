FROM golang:1.24-alpine

WORKDIR /app

COPY ./src/go.* ./

RUN go install github.com/air-verse/air@latest
RUN go get github.com/joho/godotenv@latest
RUN go get github.com/jackc/pgx/v5/pgxpool@latest
#RUN go mod tidy
RUN go mod download

COPY ./src/*.go ./
COPY ./src/.air.toml ./
COPY ./src/.env ./

EXPOSE 8888

CMD ["air", "-c", ".air.toml"]