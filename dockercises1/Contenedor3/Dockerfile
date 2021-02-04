FROM golang:latest
LABEL "Correo:"=drodriguez@tredicom.com

WORKDIR /home/Tredicom/Dockercises/1/

ENV GO111MODULE=on

RUN go mod init github.com/Drdzmtz/dockercises
RUN go mod download

COPY ./main.go ./Contenedor3/
COPY ./people.xml ./Contenedor3/

RUN go build ./Contenedor3/main.go

CMD ["./main"]