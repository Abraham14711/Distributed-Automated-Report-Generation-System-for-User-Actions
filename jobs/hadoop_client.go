package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/colinmarc/hdfs"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}
func main() {
	client, err := hdfs.NewClient(hdfs.ClientOptions{
		Addresses: []string{os.Getenv("HADOOP_HOST")},
		User:      "root",
	})
	if err != nil {
		log.Fatalf("Failed to create HDFS client: %v", err)
	}
	defer client.Close()

	if _, err := client.Stat("/"); err != nil {
		log.Fatalf("HDFS connection failed: %v", err)
	}

	hdfsDir := "/data/reports"
	today := time.Now().Format("2006-01-02")
	localDir := filepath.Join("/opt/airflow/output_data", "date="+today)

	if err := client.MkdirAll(hdfsDir, 0755); err != nil {
		log.Fatalf("Failed to create HDFS directory: %v", err)
	}

	if _, err := os.Stat(localDir); os.IsNotExist(err) {
		log.Fatalf("Local directory does not exist: %s", localDir)
	}

	files, err := os.ReadDir(localDir)
	if err != nil {
		log.Fatalf("Failed to read local directory: %v", err)
	}

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".parquet" {
			localPath := filepath.Join(localDir, file.Name())
			remotePath := filepath.Join(hdfsDir, file.Name())

			if _, err := client.Stat(remotePath); err == nil {
				log.Printf("File already exists in HDFS, skipping: %s", remotePath)
				continue
			}

			if err := client.CopyToRemote(localPath, remotePath); err != nil {
				log.Printf("Failed to copy %s: %v", localPath, err)
			} else {
				log.Printf("Successfully copied %s to HDFS", localPath)
				if err := os.Remove(localPath); err != nil {
					log.Printf("Failed to remove local file: %v", err)
				}
			}
		}
	}

	log.Println("\nChecking HDFS directory after copy:")
	listHDFSFiles(client, hdfsDir)
}

func listHDFSFiles(client *hdfs.Client, path string) {
	files, err := client.ReadDir(path)
	if err != nil {
		log.Printf("Failed to list HDFS directory: %v", err)
		return
	}

	for _, file := range files {
		fmt.Printf("%s\t%d\t%s\n",
			file.Mode(),
			file.Size(),
			file.Name(),
		)
	}
}
