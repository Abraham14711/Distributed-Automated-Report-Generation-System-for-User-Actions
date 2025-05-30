package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

var events_counter = rand.Intn(90) + 10
var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var actions = []string{
	"CREATE",
	"READ",
	"UPDATE",
	"DELETE",
	"INSERT",
	"SELECT",
	"DROP",
	"GRANT",
	"REVOKE",
	"ALTER",
	"RENAME",
	"TRUNCATE",
	"REPLACE",
}

func randString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}

func generateEmail(email_name string) string {
	email_providers := []string{"gmail.com", "yandex.ru", "yahoo.com", "mail.ru", "ya.ru"}
	return fmt.Sprintf("%s@%s", email_name, email_providers[rand.Intn(len(email_providers))])
}

func main() {
	file, err := os.Create("/opt/airflow/input_data/pure_data_" + time.Now().String()[0:10] + ".csv")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	email_names := []string{}
	for i := 0; i < events_counter/2; i++ {
		email_names = append(email_names, randString(10))
	}
	for i := 0; i < events_counter; i++ {
		file.WriteString(generateEmail(email_names[rand.Int()%len(email_names)]) + "," + actions[rand.Intn(len(actions))] + "," + strings.Replace(time.Now().String()[0:19], " ", ",", 1) + "\n")
	}
	fmt.Printf("Done")
}
