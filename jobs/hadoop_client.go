package jobs

import (
	"os"
	"path/filepath"

	"github.com/colinmarc/hdfs"
)

func main() {
	client, err := hdfs.New(os.Getenv("HADOOP_HOST"))
	if err != nil {
		panic(err)
	}
	defer client.Close()

	hdfsDir := "/data/reports"

	files, err := client.ReadDir("/opt/spark/output_data")
	for _, file := range files {
		if filepath.Ext(file.Name()) == ".parquet" {
			localFile := "/opt/spark/output_data/" + file.Name()
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
