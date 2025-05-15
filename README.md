# Distributed Automated Report Generation System for User Actions in Database


## 📌 Overview

This project implements a **distributed, scalable data processing pipeline** using modern big data tools and container orchestration. 
The pipeline integrates the following key technologies:

- **Apache Airflow** for orchestration
- **Apache Spark** for distributed processing
- **Hadoop (HDFS)** for storage
- **Go** for synthetic data generation
### Since this is a test bench, all the data we process is generated using a Go script in the generation directory. On a real stand there will be no need to use this script.
- **Kubernetes** for container orchestration
- **Docker** for containerization

---

## 🎯 Project Objectives

- Develop and orchestrate a modular data pipeline
- Use Apache Spark in a **distributed cluster** configuration
- Containerize and deploy services via Docker and Kubernetes

---

## ⚙️ Architecture

### **Data Generation**

A Go-based script (`main.go`) generates synthetic CSV data simulating user actions (e.g., `INSERT`, `DELETE`, `SELECT`, etc.). Each record includes:
- User ID
- Timestamp
- Action Type

Generated CSV files are stored in a shared volume accessible by other services.

---

### **Spark Processing**

Apache Spark is deployed in **distributed mode** with:
- **1 Master node**
- **1 Worker node**

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

- **Start → Generate CSV → Spark Processing → End**

Airflow is containerized and deployed with a web interface on:

- `8081` (Airflow Web UI)

DAGs are written in Python and handle retries, logging, and scheduling.

---



## 🔌 Ports Summary

| Service      | Description             | Port   |
|--------------|-------------------------|--------|
| **Spark Master** | Cluster communication     | `7077` |
| **Spark Master UI** | Web interface             | `8080` |
| **Airflow UI**     | DAG management interface  | `8081` |

---
