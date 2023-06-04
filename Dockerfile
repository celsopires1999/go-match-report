FROM golang:1.20.4

RUN curl -s https://packagecloud.io/install/repositories/golang-migrate/migrate/script.deb.sh | bash 
RUN apt-get update 
RUN apt-get install migrate

RUN useradd -u 1000 dev
USER dev

WORKDIR /home/dev/src

ENV PATH="/go/bin:${PATH}"

RUN go install github.com/kyleconroy/sqlc/cmd/sqlc@latest

CMD ["tail", "-f", "/dev/null"]
