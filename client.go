package main


import (
	"strconv"
	"github.com/sgjp/go-coap"
	"log"
	"time"
)

var PrimeNumsQty int


func startClientWorker(c chan *coap.Message) {
	conf := getConfiguration()

	for{
		m:= <- c
		if (m!= nil) {

			primeNumber,_ := getPrimeNumber(m)

			//Call the proxy
			if cipherAES{
				sendCoapMsg("/fwd", encrypt(strconv.Itoa(primeNumber)), coap.PUT, coap.Confirmable, "calcPrimeNumberResult", conf.Proxy.Host, m.MessageID, m.Token)
			}else {
				sendCoapMsg("/fwd", strconv.Itoa(primeNumber), coap.PUT, coap.Confirmable, "calcPrimeNumberResult", conf.Proxy.Host, m.MessageID, m.Token)
			}


		}

	}
}

func startClientManager(){
	conf := getConfiguration()
	for i:=1;i<=PrimeNumsQty;i++{
		//log.Printf("Sending %v",i)
		var rv *coap.Message
		if cipherAES{
			rv = sendCoapMsg("/fwd",encrypt(strconv.Itoa(i)),coap.GET,coap.Confirmable,"calcPrimeNumber",conf.Proxy.Host,GenerateMessageID(),make([]byte, 8))
		}else{
			rv = sendCoapMsg("/fwd",strconv.Itoa(i),coap.GET,coap.Confirmable,"calcPrimeNumber",conf.Proxy.Host,GenerateMessageID(),make([]byte, 8))
		}

		if rv.Code ==coap.Content{
			//This means the response is cached and it has a result
			PrimeNumsCalculated++
			log.Printf("Calculated: %v prime numbers",PrimeNumsCalculated)
			if cipherAES{
				//Decrypt it just for timing
				decrypt(string(rv.Payload))
			}
			if PrimeNumsCalculated==PrimeNumsQty{
				elapsed := time.Since(StartTime)
				saveTaskDuration(int64(elapsed/time.Millisecond), PrimeNumsQty)
				log.Printf("%v prime numbers calculated in %v ms STARTDATE: %v",PrimeNumsQty,(int64(elapsed/time.Millisecond)),StartTime)
			}

		}
		//log.Printf(string(rv.Payload))
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

func sendCoapMsg(path string, payload string, coapCode coap.COAPCode, coapType coap.COAPType, proxyUri, host string, messageId uint16, token []byte) *coap.Message{
	req := coap.Message{
		Type:      coapType,
		Code:      coapCode,
		MessageID: messageId,
		Token:	   token,
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