package main


import (
	"strconv"
	"github.com/sgjp/go-coap"
	"log"
)

var PrimeNumsQty int


func startClientWorker(c chan *coap.Message) {
	conf := getConfiguration()

	for{
		m:= <- c
		if (m!= nil) {

			primeNumber,_ := getPrimeNumber(m)

			//Call the proxy

			sendCoapMsg("/fwd",strconv.Itoa(primeNumber),coap.GET,coap.Confirmable,"calcPrimeNumberResult",conf.Proxy.Host)



		}

	}
}

func startClientManager(){
	conf := getConfiguration()
	for i:=1;i<=PrimeNumsQty;i++{
		log.Printf("Sending %v",i)

		rv := sendCoapMsg("/fwd",strconv.Itoa(i),coap.GET,coap.Confirmable,"calcPrimeNumber",conf.Proxy.Host)


		log.Printf(string(rv.Payload))
		//calcPrimeNumberRemotely(i)
	}
}


func calcPrimeNumberRemotely(number int){
	//Call the proxy

	conf := getConfiguration()

	//resourcePayload := ResourcePayload {Rn: "/calcPrimeNumber", Rd: string(m.Payload)}

	req := coap.Message{
		Type:      coap.Confirmable,
		Code:      coap.GET,
		MessageID: GenerateMessageID(),
		Payload:   []byte(strconv.Itoa(number)),
	}

	req.SetPathString("/fwd")
	req.SetOption(coap.ProxyURI,"calcPrimeNumber")

	c, err := coap.Dial("udp", conf.Proxy.Host)
	if err != nil {
		log.Fatalf("Error dialing: %v", err)
	}

	rv, err := c.Send(req)
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}

	if rv != nil {

	}


	if rv.Option(coap.MaxAge) == nil{
		//return num, 60
	}
	log.Printf("SENT: %v",number)
	//return num, rv.Option(coap.MaxAge)

}

func sendCoapMsg(path string, payload string, coapCode coap.COAPCode, coapType coap.COAPType, proxyUri, host string) *coap.Message{
	req := coap.Message{
		Type:      coapType,
		Code:      coapCode,
		MessageID: GenerateMessageID(),
		Payload:   []byte(payload),
	}


	req.SetOption(coap.MaxAge, 3)
	req.SetOption(coap.ProxyURI, proxyUri)

	req.SetPathString(path)

	c, err := coap.Dial("udp", host)
	if err != nil {
		log.Fatalf("Error dialing: %v", err)
	}

	rv, err := c.Send(req)
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}

	if rv != nil {
		//log.Printf("Response payload: %s", rv.Payload)
	}
	return rv
}