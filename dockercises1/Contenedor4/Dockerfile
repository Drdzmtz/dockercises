FROM golang:latest
LABEL "Correo:"=drodriguez@tredicom.com

WORKDIR /home/Tredicom/Dockercises/1/Contenedor4

ENV GO111MODULE=on


COPY ./main.go .

# COPY go.mod .
# COPY go.sum .
RUN go mod download

RUN go build main.go

CMD ["./main"]