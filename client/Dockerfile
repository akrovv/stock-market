FROM golang:alpine AS builder

RUN mkdir /client
WORKDIR /client

COPY . .
COPY .env .

RUN go mod download && go build cmd/client/main.go

FROM alpine:latest

RUN mkdir /client && mkdir /client/front
WORKDIR /client

COPY --from=builder /client/front ./front
COPY --from=builder /client/.env .
COPY --from=builder /client/main .

EXPOSE 8000

CMD [ "./main" ]