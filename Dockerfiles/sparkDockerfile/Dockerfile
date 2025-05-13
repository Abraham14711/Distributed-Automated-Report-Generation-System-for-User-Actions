FROM apache/spark:3.5.0

USER root 

RUN python3 -m pip install --no-cache-dir pyspark

RUN mkdir -p /opt/spark/generation && \
    mkdir -p /opt/spark/input_data && \
    mkdir -p /opt/spark/jobs

RUN mkdir -p /home/spark && \
    chown spark:spark /home/spark && \
    chmod 755 /home/spark

USER spark
WORKDIR /opt/spark