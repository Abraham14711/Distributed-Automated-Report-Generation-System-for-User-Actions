from pyspark.sql import SparkSession
from pyspark.sql.types import StructType, StructField, StringType, TimestampType
from datetime import datetime 

date = datetime.now().strftime("%Y-%m-%d")

spark = SparkSession.builder.appName("DAG").getOrCreate()
spark.sparkContext.setLogLevel("ERROR")

schema = StructType([
StructField("email", StringType(), False),
StructField("action", StringType(), False),
StructField("date", TimestampType(), False)
])