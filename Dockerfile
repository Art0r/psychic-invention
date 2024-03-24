FROM golang:latest

WORKDIR /app

ADD . /app/

EXPOSE 80

CMD [ "go", "run", "main.go" ]