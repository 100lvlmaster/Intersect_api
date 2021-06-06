# Build Stage
FROM golang:alpine as build

WORKDIR /src/Intersect_api

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o /app/Intersect_api

# Final Stage
FROM alpine:latest
WORKDIR /app
COPY --from=build /app/Intersect_api /app/

EXPOSE 8080

CMD ./Intersect_api

