FROM golang:latest
LABEL "Correo:"=drodriguez@tredicom.com

WORKDIR /home/Tredicom/Dockercises/1/

RUN go mod init github.com/Drdzmtz/dockercises/dockercises1

COPY ./main.go ./Contenedor3/
COPY ./people.xml ./Contenedor3/

RUN go build ./Contenedor3/main.go

CMD ["Contenedor3./main"]