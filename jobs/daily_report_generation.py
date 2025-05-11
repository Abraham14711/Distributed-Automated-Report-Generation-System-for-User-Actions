from pyspark.sql import SparkSession
from pyspark.sql.types import StructType, StructField, StringType, TimestampType
import pyspark.sql.functions as F
from datetime import datetime 

date = datetime.now().strftime("%Y-%m-%d")

spark = SparkSession.builder.appName("daily report generation").getOrCreate()
spark.sparkContext.setLogLevel("ERROR")

schema = StructType([
StructField("email", StringType(), False),
StructField("action", StringType(), False),
StructField("date", TimestampType(), False),
StructField("time", StringType(), False)
])

file = spark.read.csv("input_data/pure_data_" + date + ".csv", header=False, sep=",",schema=schema)

counted_actions = file.groupBy("email").agg(
    F.count(F.when(F.col("action") == "CREATE", 1)).alias("create_count"),
    F.count(F.when(F.col("action") == "READ", 1)).alias("read_count"),
    F.count(F.when(F.col("action") == "UPDATE", 1)).alias("update_count"),
    F.count(F.when(F.col("action") == "DELETE", 1)).alias("delete_count"),
    F.count(F.when(F.col("action") == "INSERT", 1)).alias("insert_count"),
    F.count(F.when(F.col("action") == "SELECT", 1)).alias("select_count"),
    F.count(F.when(F.col("action") == "DROP", 1)).alias("drop_count"),
    F.count(F.when(F.col("action") == "GRANT", 1)).alias("grant_count"),
    F.count(F.when(F.col("action") == "REVOKE", 1)).alias("revoke_count"),
    F.count(F.when(F.col("action") == "ALTER", 1)).alias("alter_count"),
    F.count(F.when(F.col("action") == "RENAME", 1)).alias("rename_count"),
    F.count(F.when(F.col("action") == "TRUNCATE", 1)).alias("truncate_count"),
    F.count(F.when(F.col("action") == "REPLACE", 1)).alias("replace_count")
)

counted_actions.show()

counted_actions.write.parquet("output_data")