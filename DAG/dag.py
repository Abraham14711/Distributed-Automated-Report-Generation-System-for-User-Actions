from datetime import datetime, timedelta
from airflow import DAG
from airflow.operators.python import BranchPythonOperator
from airflow.providers.apache.spark.operators.spark_submit import SparkSubmitOperator
from airflow.operators.empty import EmptyOperator
from airflow.operators.bash import BashOperator
from airflow.utils.dates import days_ago

default_args = {
    "owner":"airflow"    
}

with DAG(
    dag_id = "spark_dag",
    default_args = default_args,
    schedule_interval = "@daily",
    start_date = days_ago(6),
    catchup = True,
    description = "Weekly report generation"
) as dag: 
    
    start = EmptyOperator(task_id = "start")
    
    get_newest_data = BashOperator(
        task_id  = "newest data download",
        bash_command = "go run /opt/airflow/generation/main.go"
    )

    create_daily_report = SparkSubmitOperator(
        task_id = "daily report creation",
        conn_id = "spark connection",
        application = "/opt/airflow/"# TODO write dayly report generation and change path
    )