package jobs

import (
	"os"

	"github.com/colinmarc/hdfs"
)

func main() {
	client, err = hdfs.New(os.Getenv("HADOOP_HOST"))
}
