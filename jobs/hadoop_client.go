package main

import (
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/colinmarc/hdfs"
)

func main() {
	namenodeAddress := os.Getenv("HADOOP_HOST")
	client, err := hdfs.NewClient(hdfs.ClientOptions{
		Addresses: []string{namenodeAddress},
		User:      "root",
	})
	if _, err := client.Stat("/"); err != nil {
		log.Fatalf("HDFS connection failed: %v", err)
	}

	if err != nil {
		panic(err)
	}
	defer client.Close()

	hdfsDir := "/data/reports"

	files, err := os.ReadDir("output_data/date=" + time.Now().String()[0:10] + "/")
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		if filepath.Ext(file.Name()) == ".parquet" {
			localFile := "output_data/date=" + time.Now().String()[0:10] + "/" + file.Name()
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
