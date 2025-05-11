FROM apache/airflow:2.7.1-python3.11

USER root

RUN apt-get update && \
    apt-get upgrade -y && \
    apt-get install -y gcc python3-dev openjdk-11-jdk procps

USER airflow

RUN pip install apache-airflow==2.7.1 apache-airflow-providers-apache-spark pyspark