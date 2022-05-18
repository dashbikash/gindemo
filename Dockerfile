FROM golang:alpine3.15

WORKDIR /app

COPY . ./

RUN go mod download

RUN go build -o /bin/go-app

ENV APP_PORT=3030
EXPOSE 3030

CMD [ "/bin/go-app" ]