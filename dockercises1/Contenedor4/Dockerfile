FROM golang:latest
LABEL "Correo:"=drodriguez@tredicom.com

WORKDIR /home/Tredicom/Dockercises/1/Contenedor4

COPY ./main.go .
COPY ./mongoConnection.go .
COPY ./people.go .

RUN go get -v -u go.mongodb.org/mongo-driver/mongo
RUN go get -u github.com/go-chi/chi
RUN go get -u github.com/go-chi/render

CMD ["go", "run", "main.go", "mongoConnection.go", "people.go"]