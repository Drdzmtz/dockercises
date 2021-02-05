FROM golang:latest
LABEL "Correo:"=drodriguez@tredicom.com

WORKDIR /home/Tredicom/Dockercises/1/Contenedor4/

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY ./main.go .

RUN go build main.go

CMD ["./main"]