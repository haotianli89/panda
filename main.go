package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/micro/go-micro/cmd"
	"github.com/haotianli89/driversvc/pb"
	"github.com/micro/go-micro/client"
	"golang.org/x/net/context"
)

func main() {
	fmt.Println("listening 0.0.0.0:8080")
	http.HandleFunc("/", handler) // each request calls handler
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}

// handler echoes the Path component of the request URL r.
func handler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("http request: ", *r)

	cmd.Init()

	client := driversvc.NewDriversvcClient("driversvc", client.DefaultClient)

	rsp, err := client.GetDrivers(context.Background(), &driversvc.GetDriversRequest{/*Id: "add45514-7dfb-4ce7-a74b-26d380939833"*/})
	if err != nil {
		fmt.Println(err)
	}else{
		fmt.Println(rsp)
	}

	fmt.Fprintf(w, rsp.String())
}
