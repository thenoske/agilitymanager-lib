ARG ARCH=amd64
FROM --platform=linux/${ARCH} python:3.10.2-slim-bullseye

RUN apt-get update
RUN apt-get update \
    && apt-get -y install libpq-dev gcc \
    && apt-get -y install gettext \
    && apt-get -y install cron \
    && apt-get -y install git \
    && apt-get -y install nano \
    && apt-get -y install mc \
    && apt-get -y install wget

# Set environment variables
ENV PIP_DISABLE_PIP_VERSION_CHECK 1
ENV PYTHONDONTWRITEBYTECODE 1
ENV PYTHONUNBUFFERED 1

# Set work directory
WORKDIR /code

# Copy project
COPY . /code/

# Install dependencies
RUN  rm -rf /usr/local/go && tar -C /usr/local -xzf go/go1.23.3.linux-amd64.tar.gz
RUN export PATH=$PATH:/usr/local/go/bin



