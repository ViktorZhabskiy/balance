FROM golang:1.16

ENV GO111MODULE="on"

RUN apt-get update && \
    apt-get install -y \
    lsb-release \
    apt-transport-https \
    ca-certificates \
    curl \
    gnupg-agent

RUN git config --global --add url."git@github.com:".insteadOf "https://github.com/"

#install migrate cli #details https://github.com/golang-migrate/migrate
RUN curl -sSL https://packagecloud.io/golang-migrate/migrate/gpgkey | apt-key add -
RUN echo "deb https://packagecloud.io/golang-migrate/migrate/debian/ $(lsb_release -sc) main" > /etc/apt/sources.list.d/migrate.list
RUN apt-get update && \
    apt-get install -y migrate

EXPOSE "8001"

WORKDIR /code
