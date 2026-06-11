package main

import (
	"fmt"
	"study/feature1"
	"study/feature2"
	simpleconnection "study/feature_postgres/simple_connection/simple_connection"
)

func main() {
	fmt.Println("hello Git")
	feature1.Feature1()
	feature2.Feature2()
	simpleconnection.CheckConnection()
}