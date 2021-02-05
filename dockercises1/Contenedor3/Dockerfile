FROM golang:latest
LABEL "Correo:"=drodriguez@tredicom.com

WORKDIR /home/Tredicom/Dockercises/1/Contenedor3/

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY ./main.go .
COPY ./people.xml .

RUN go build main.go

CMD ["./main"]