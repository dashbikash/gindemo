FROM golang:alpine3.15

WORKDIR /app
COPY . /app
RUN go install

ENV GIN_MODE=release
ENV APP_PORT=3030

EXPOSE 3030
CMD go build -o bin/app
ENTRYPOINT ["./bin/app" ]