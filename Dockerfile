ARG TARGET

FROM golang:1.21.7-alpine3.18 as build

RUN apk update && apk add --no-cache make

ENV GOFLAGS="-mod=readonly"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

# Copy the rest of the application files to the container
COPY . .

RUN make bins

EXPOSE 8080

FROM debian:stable-slim AS alpine
RUN apt update -y
SHELL ["/bin/bash", "-c"]

FROM alpine as e-commerce-products-service
COPY --from=build /app/bins/products-service /usr/local/bin
ENTRYPOINT ["products-service"]

FROM alpine as e-commerce-users-service
COPY --from=build /app/bins/users-service /usr/local/bin
ENTRYPOINT ["users-service"]

FROM e-commerce-${TARGET}-service