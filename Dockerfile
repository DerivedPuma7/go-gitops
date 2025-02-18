FROM golang:1.23 AS build
WORKDIR /app
COPY . .
RUN [ ! -f go.mod ] && go mod init go-kubernetes || true
RUN go mod tidy || true
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main

FROM scratch
WORKDIR /app
COPY --from=build /app/main .

CMD ["./main"]