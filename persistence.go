package main

import (
	"encoding/json"
	//"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"encoding/csv"
)

var confFile = "sgjp/ElasticServer/conf"

var taskDurationFile = "sgjp/ElasticServer/TaskDuration.csv"


func getConfiguration() Configuration {
	confValue := readFromFile(confFile)
	var configuration Configuration

	json.Unmarshal([]byte(confValue), &configuration)
	return configuration

}


func readFromFile(fileName string) string {

	stream, err := ioutil.ReadFile(fileName)

	if err != nil {
		log.Fatal(err)
	}

	readString := string(stream)

	return readString
}



func saveTaskDuration(elapsed int64, qty int){
	record := []string{
		strconv.Itoa(qty), strconv.FormatInt(elapsed,10)}

	file, er := os.OpenFile(taskDurationFile, os.O_RDWR|os.O_APPEND, 0666)

	if er != nil {
		log.Fatal(er)
	}
	defer file.Close()
	writer := csv.NewWriter(file)

	err := writer.Write(record)


	if err != nil {
		log.Fatal(er)
	}

	defer writer.Flush()
}

type Configuration struct {
	Fwd       bool
	Proxy 	  Proxy
}
type Proxy struct {
	Id   int
	Host string
}

