FROM apache/airflow:2.7.1-python3.11

USER root

RUN apt-get update && \
    apt-get install -y --no-install-recommends \
        gcc \
        python3-dev \
        openjdk-11-jdk \
        procps \
        wget \
        tar \
        gcc-aarch64-linux-gnu \
        build-essential && \
    rm -rf /var/lib/apt/lists/*


RUN wget https://golang.org/dl/go1.21.3.linux-arm64.tar.gz -O /tmp/go.tar.gz && \
    tar -C /usr/local -xzf /tmp/go.tar.gz && \
    rm /tmp/go.tar.gz

    RUN mkdir -p /go && \
    chown -R airflow /go && \
    mkdir -p /opt/airflow/jobs && \
    chown -R airflow /opt/airflow


ENV PATH="/usr/local/go/bin:${PATH}"
ENV GOPATH="/go"
ENV PATH="${GOPATH}/bin:${PATH}"
ENV GOARCH=arm64
ENV GOOS=linux
ENV CGO_ENABLED=1
ENV CC=aarch64-linux-gnu-gcc
ENV CXX=aarch64-linux-gnu-g++

USER airflow
RUN pip install --no-cache-dir \
    apache-airflow==2.7.1 \
    apache-airflow-providers-apache-spark \
    pyspark \
    pandas

COPY .env /opt/airflow/.env