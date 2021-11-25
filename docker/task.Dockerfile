FROM golang:1.17.1-alpine3.14 AS build

ENV GO111MODULE=on
WORKDIR /app
COPY ../go.mod ../go.sum ./
RUN go mod download

COPY .. .
WORKDIR /app/services/task/cmd
RUN go build

FROM alpine:3.14

ENV ENV_FILE=task
COPY --from=build /app/env/task.env /env/
COPY --from=build /app/services/task/cmd /usr/local/bin/

EXPOSE 9090
CMD ["cmd"]
