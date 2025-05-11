FROM apache/spark:3.5.0

RUN mkdir -p /opt/spark/generation && \
    mkdir -p /opt/spark/input_data && \
    mkdir -p /opt/spark/jobs

WORKDIR /opt/spark