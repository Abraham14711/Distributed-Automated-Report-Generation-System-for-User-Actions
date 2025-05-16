# Distributed Automated Report Generation System for User Actions in Database


## 📌 Overview

This project implements a **distributed, scalable data processing pipeline** using modern big data tools and container orchestration. 
The pipeline integrates the following key technologies:

- **Apache Airflow** for orchestration
- **Apache Spark** for distributed processing
- **Hadoop (HDFS)** for storage
- **Python** for managing aiflow pipeline using DAG files 
- **Go** for synthetic data generation and implementing an HDFS client
- **Kubernetes** for the orchestration of distributed database components 
- **Docker/Docker-compose** for containerization and orchestration 

---

## 🎯 Project Objectives

- Develop and orchestrate a modular data pipeline
- Use Apache Spark in a **distributed cluster** configuration
- Containerize and deploy services via Docker and Kubernetes

---

## ⚙️ Architecture

The entire system runs on a **Kubernetes cluster consisting of 3 nodes**.

### **Data Generation**

A Go-based script (`main.go`) generates synthetic CSV data simulating user actions (e.g., `INSERT`, `DELETE`, `SELECT`, etc.). Each record includes:
- User Email
- Timestamp
- Action Type

Generated CSV files are stored in a shared volume accessible by other services.

---

### **Spark Processing**

Apache Spark is deployed in **distributed mode** with:
- **1 Master node**
- **2 Worker node**

Spark listens on:
- `7077` for cluster communication
- `8080` for Web UI (Master)

#### How It Works:
- Spark reads the raw CSV data
- Groups records by email and event type
- Aggregates frequencies
- Stores results in **Parquet format**

Airflow communicates with the Spark Master via the `SparkSubmitOperator`. The Master distributes tasks to the workers for execution.

---

### **Airflow Orchestration**

Apache Airflow manages the pipeline execution flow:

- **Start → Generate CSV → Spark Processing → Send to HDFS → End**

Airflow is containerized and deployed with a web interface on:

- `8081` (Airflow Web UI)

DAGs are written in Python and handle retries, logging, and scheduling.

---

## 🧱 Hadoop Deployment in Kubernetes

Hadoop is deployed in Kubernetes using custom Helm charts:

```bash
helm repo add pfisterer-hadoop https://pfisterer.github.io/apache-hadoop-helm/
helm install hadoop   --set persistence.dataNode.size=10Gi   --set persistence.nameNode.size=10Gi pfisterer-hadoop/hadoop
```




## 🔌 Ports Summary

| Service            | Description                  | Port   |
|--------------------|------------------------------|--------|
| **Spark Master**   | Cluster communication        | `7077` |
| **Spark Master UI**| Web interface (Spark Master) | `8080` |
| **Airflow UI**     | DAG management interface     | `8081` |
| **HDFS** |  NameNode client interface   | `8020` |
| **Hadoop UI**      | HDFS Web UI       | `9870` |

---

---
