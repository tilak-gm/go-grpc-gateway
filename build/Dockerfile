FROM golang:1.22.3-alpine3.19 as dep
WORKDIR /app
ADD go.mod go.sum ./
RUN go mod download

FROM dep as build
WORKDIR /app
COPY . .
RUN find ./ -type f \
    \! -name "*.go" \! -name "*.mod" \! -name "*.sum" -delete
RUN go build -o user-service ./cmd/server
RUN go build -o gateway-service ./cmd/client

FROM alpine:latest as user-service
WORKDIR /app
COPY --from=build /app/user-service /usr/local/bin/user-service
EXPOSE 8000
ENTRYPOINT ["user-service"]

FROM alpine:latest as gateway-service
WORKDIR /app
COPY --from=build /app/gateway-service /usr/local/bin/gateway-service
EXPOSE 8080
ENTRYPOINT ["gateway-service"]
