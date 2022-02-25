FROM golang:1.17.7-alpine
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
RUN go build -o /go-backend-app
ENV DB_HOST localhost
ENV HTTP_PORT 3000
EXPOSE 3000
CMD [ "/go-backend-app" ]