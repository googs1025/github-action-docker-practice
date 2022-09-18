package main
import (
	"fmt"
	"log"
	"net/rpc"
)

// MathRequest 消息
type MathRequest struct{
	A int
	B int
}

func main(){
	client, err := rpc.Dial("tcp", "127.0.0.1:8888")
	if(err != nil){
		log.Fatal("Dial failed", err)
	}
	var sum int
	err = client.Call("Sum.Sum", MathRequest{1, 5} ,&sum)
	if(err != nil){
		log.Fatal("Call failed", err)
	}
	fmt.Println("sum:", sum)//sum: 6
}
