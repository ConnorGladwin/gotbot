FROM golang:1.22

WORKDIR /api

COPY go.mod ./
RUN go mod download

COPY ./.env ./
# COPY *.go ./
# COPY ./web_server/*.go ./web_server/
# COPY ./web_server/auth/*.go ./web_server/auth/
# COPY ./web_server/queries/*.go ./web_server/queries/
# COPY ./web_server/database/*.go ./web_server/database/

COPY ./ .

RUN go mod tidy
RUN go build -o /backend

EXPOSE 8080

CMD ["/backend"]
