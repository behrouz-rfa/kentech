FROM golang:1.22.1-alpine3.19 as dev

ENV GO111MODULE=on

ENV APP_HOME /app
WORKDIR $APP_HOME

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY cmd ./cmd
COPY pkg ./pkg
COPY firebase.json ./firebase.json

RUN go build -o ./server ./cmd/server/main.go

FROM alpine as release

ENV APP_HOME /app

WORKDIR $APP_HOME

COPY --from=dev ./app/server ./server
COPY --from=dev ./app/firebase.json ./firebase.json
COPY ./scripts/docker-entrypoint.sh ./docker-entrypoint.sh

ENTRYPOINT ["sh", "./docker-entrypoint.sh"]
EXPOSE 80
