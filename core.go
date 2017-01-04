package main

import (
	"github.com/sgjp/go-coap"
	"log"
	"strconv"
)

func getPrimeNumber(m *coap.Message) (int, interface{}) {

	//conf := getConfiguration()
	//if conf.Fwd {
	   //return calcPrimeNumberRemotely(m, conf.Proxy)
	//}
	//qty := parseRequestPayload(string(m.Payload))
	if cipherAES{
		qty,_ :=  strconv.Atoi(decrypt(string(m.Payload)))
		return calcPrimeNumberLocally(qty), 60
	}else{
		qty,_ :=  strconv.Atoi(string(m.Payload))
		return calcPrimeNumberLocally(qty), 60
	}

}

func calcPrimeNumberLocally(qty int) int {
	var num int

	for i := 2; i < qty; i++ {
		for j := 2; j < i; j++ {
			if i%j == 0 {
				break
			} else if i == j+1 {
				num = i
			}
		}
	}

	log.Println("Calculated Locally: " + strconv.Itoa(num))
	return num
}








