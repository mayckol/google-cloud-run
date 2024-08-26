# Stage 1: Build the application
FROM golang:1.22 AS build
WORKDIR /app
COPY . .
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o cloudrun

# Stage 2: Test the application
FROM golang:1.22 AS test
WORKDIR /app
COPY . .
CMD ["go", "test", "./...", "-v"]

# Stage 3: Final application image
FROM scratch
WORKDIR /app
COPY --from=build /app/cloudrun .
CMD ["./cloudrun"]
