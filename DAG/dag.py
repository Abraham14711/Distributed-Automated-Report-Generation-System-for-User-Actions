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
    start_date = days_ago(0),
    catchup = False,
    description = "Weekly report generation"
) as dag: 
    
    start = EmptyOperator(task_id = "start")
    
    get_newest_data = BashOperator(
        task_id  = "newest_data_download",
        bash_command = "go run /opt/airflow/generation/main.go"
    )

    create_daily_report = SparkSubmitOperator(
        task_id = "daily_report_creation",
        conn_id = "spark_connection",
        conf={"spark.master": "local"},
        application = "/opt/airflow/jobs/daily_report_generation.py",
    )

    end = EmptyOperator(task_id = 'end')

    (
        start
        # >> get_newest_data 
        >> create_daily_report
        >> end
    )