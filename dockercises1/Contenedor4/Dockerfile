FROM golang:latest
LABEL "Correo:"=drodriguez@tredicom.com

WORKDIR /home/Tredicom/Dockercises/1/
ENV GO111MODULE=on

COPY ./main.go .Contenedor4/

RUN go mod download
RUN go build ./Contenedor4/main.go

CMD ["./main"]