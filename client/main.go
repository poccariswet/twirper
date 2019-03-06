package main

import (
	"context"
	"fmt"
	"net/http"
)

func main() {
	c := pb.NewHelloWorldProtobufClient("http://localhost:8080", &http.Client{})

	var subject string
	if len(os.Args) > 1 {
		subject = os.Args[1]
	}

	resp, err := c.Hello(context.Background, &pb.HelloReq{Subject: subject})
	if err != nil {
		fmt.Fprint(err, os.Stdderr)
		os.Exit(1)
	}

	fmt.Println(resp.GetText())
}
