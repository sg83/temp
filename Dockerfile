FROM golang:1.20-alpine

RUN mkdir /app
ADD . /app
WORKDIR /app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
#COPY go.mod go.sum ./
#RUN go mod download && go mod verify

#COPY . .
RUN go mod tidy
RUN go build -v -o /app/api

EXPOSE 8080

CMD ["./app/api"]