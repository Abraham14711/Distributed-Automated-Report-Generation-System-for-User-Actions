package jobs

import (
	"os"
	"path/filepath"
	"time"

	"github.com/colinmarc/hdfs"
)

func main() {
	client, err := hdfs.New(os.Getenv("HADOOP_HOST"))
	if err != nil {
		panic(err)
	}
	defer client.Close()

	hdfsDir := "/data/reports"

	files, err := client.ReadDir("/opt/spark/output_data/date=" + time.Now().String()[0:10])
	for _, file := range files {
		if filepath.Ext(file.Name()) == ".parquet" {
			localFile := "/opt/spark/output_data/date=" + time.Now().String()[0:10] + "/" + file.Name()
			remoteFile := hdfsDir + "/" + file.Name()
			err = client.CopyToRemote(localFile, remoteFile)
			if err != nil {
				panic(err)
			} else {
				os.Remove(localFile)
			}
		}
	}
}
