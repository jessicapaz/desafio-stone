FROM golang:latest
ENV GO111MODULE=on
LABEL maintainer="Jessica Paz <jessicamorim.42@gmail.com>"
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
EXPOSE 8966
ENTRYPOINT ["/app/desafio-stone"]
