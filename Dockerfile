FROM golang:1.23-alpine AS build
RUN mkdir /var/app
WORKDIR /var/app
COPY . /var/app/

RUN go mod download

RUN go build -o /app ./cmd/api/...

FROM alpine:latest

# We need to copy the binary from the build image to the production image.
COPY --from=build /app .
EXPOSE 8080
CMD [ "/app" ]