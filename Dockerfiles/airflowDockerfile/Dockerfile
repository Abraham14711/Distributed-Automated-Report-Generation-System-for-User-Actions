FROM apache/airflow:2.7.1-python3.11

USER root

RUN apt-get update && \
    apt-get install -y --no-install-recommends \
        gcc \
        python3-dev \
        openjdk-11-jdk \
        procps && \
    rm -rf /var/lib/apt/lists/*

USER airflow

RUN pip install --no-cache-dir \
    apache-airflow==2.7.1 \
    apache-airflow-providers-apache-spark \
    pyspark \
    pandas