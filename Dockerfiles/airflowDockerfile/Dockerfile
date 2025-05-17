FROM apache/airflow:2.7.1-python3.11

USER root

RUN apt-get update && \
    apt-get install -y --no-install-recommends \
        gcc \
        python3-dev \
        openjdk-11-jdk \
        procps && \
    rm -rf /var/lib/apt/lists/*

RUN apt-get update && \
    apt-get install -y wget tar gcc && \
    rm -rf /var/lib/apt/lists/* && \
    wget https://golang.org/dl/go1.21.3.linux-amd64.tar.gz -O /tmp/go.tar.gz && \
    tar -C /usr/local -xzf /tmp/go.tar.gz && \
    rm /tmp/go.tar.gz

RUN mkdir -p /go && \
    mkdir -p /opt/airflow/input_data && \
    mkdir -p /opt/airflow/output_data && \
    chown -R airflow /go && \
    chown -R airflow /opt/airflow/input_data && \
    chown -R airflow /opt/airflow/output_data && \

USER airflow

ENV PATH="/usr/local/go/bin:${PATH}"
ENV GOPATH="/go"
ENV PATH="${GOPATH}/bin:${PATH}"


RUN pip install --no-cache-dir \
    apache-airflow==2.7.1 \
    apache-airflow-providers-apache-spark \
    pyspark \
    pandas