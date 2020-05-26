FROM golang:latest

ARG app_env
ENV APP_ENV $app_env

WORKDIR /

COPY . .

CMD ["go", "run", "main.go"]