package main


import (
	"time"
	"github.com/sgjp/go-coap"
	"log"
)

var mode int

var StartTime time.Time

var port string

var cipherAES bool

var hostManager string
func main() {

	//mode sets the device to Client(Manager): 0 or Server(Worker): 1
	mode = 1;
	PrimeNumsQty = 1000
	port = ":5684"
	cipherAES = false
	c := make(chan *coap.Message, 100000)

	

	if mode==1{
		log.Printf("Starting as a Worker...")
		go startClientWorker(c)
		startServer(c)
	}else if mode ==0{
		log.Printf("Starting as a Manager...")
		go startClientManager()
		startServer(c)


	}




}
