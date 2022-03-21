package main

import (
	"fmt"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

type Test struct {
	Id   string `dynamo:"Id,hash"`
	Name string `dynamo:"Name"`
}

func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Hello Usagisan")
	ScanTest()
	fmt.Fprintf(writer, "Hello Usagisan")
}

func helth_handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Hello helth")
	fmt.Fprintf(writer, "helth")
}

func ScanTest() {

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("ap-northeast-1"),
		Endpoint:    aws.String("http://dynamodb-localstack.default.svc.cluster.local:4566"),
		Credentials: credentials.NewStaticCredentials("test-key", "test-secret", ""),
	})
	if err != nil {
		panic(err)
	}
	db := dynamo.New(sess)
	table := db.Table("Test")
	var result Test

	err = table.Get("Id", "0001").One(&result)
	if err != nil {
		panic(err)
	}
	fmt.Printf("GetDB%+v\n", result)
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/helthz", helth_handler)
	http.ListenAndServe(":8080", nil)
}
