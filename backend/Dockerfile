FROM golang:1.22

WORKDIR /api

COPY go.mod ./
RUN go mod download

COPY ./ .

RUN go mod tidy
RUN go build -o /backend

EXPOSE 8080
CMD ["/backend"]
