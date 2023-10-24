FROM golang:1.19-bullseye

WORKDIR /app

COPY . .
RUN go mod download

COPY . .

RUN go build -o /go-docker

EXPOSE 3000

CMD [ "/go-docker" ]