FROM golang:1.22.1

WORKDIR /api

COPY go.mod ./
RUN go mod download

COPY ./ .

RUN go mod tidy
RUN go build -o /backend

EXPOSE 3000 
CMD ["/backend"]
