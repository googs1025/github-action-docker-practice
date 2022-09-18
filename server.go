package main
import (
	"log"
	"net"
	"net/rpc"
)
type Math struct {}
type MathRequest struct{
	A int
	B int
}
func (p *Math) Sum(request MathRequest, response *int) error {
	*response = request.A + request.B
	return nil
}
func main(){
	rpc.RegisterName("Sum", new(Math))

	listener, err := net.Listen("tcp", "0.0.0.0:8888")
	if(err != nil){
		log.Fatal("ListenTCP error:", err)
	}

	for {
		socket, err := listener.Accept()
		if(err != nil){
			log.Fatal("Accept error:", err)
		}
		go rpc.ServeConn(socket)
	}
}
