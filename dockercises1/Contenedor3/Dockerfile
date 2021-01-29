FROM golang:latest
LABEL "Correo:"=drodriguez@tredicom.com

WORKDIR /home/Tredicom/Dockercises/1/Contenedor3

COPY ./main.go .
COPY ./people.xml .
COPY ./mongoConnection.go .
COPY ./people.go .

RUN go get -v -u go.mongodb.org/mongo-driver/mongo

CMD ["go", "run", "main.go", "mongoConnection.go", "people.go"]