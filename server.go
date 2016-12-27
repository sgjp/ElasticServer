package main

import (
	//"fmt"
	"github.com/sgjp/go-coap"
	"log"
	"net"
	"strconv"
	"time"
	"strings"
)

var PrimeNumsCalculated int

func startServer(c chan *coap.Message) {

	resourcesToPublish := []string{"calcPrimeNumber", "calcPrimeNumberResult"}

	StartTime = time.Now()

	log.Fatal(coap.ListenAndServe("udp", port,
		coap.FuncHandler(func(l *net.UDPConn, a *net.UDPAddr, m *coap.Message) *coap.Message {
			log.Printf("Got message path=%q: PayLoad: %#v from %v", m.Path(), string(m.Payload), a)
			if len(m.Path()) > 0 {

				switch m.Path()[0] {

				case "calcPrimeNumber":



					res := &coap.Message{
						Type:      coap.Acknowledgement,
						Code:      coap.Content,
						MessageID: m.MessageID,
						Token:     m.Token,
						Payload:   []byte("2.05"),
					}
					res.SetOption(coap.ContentFormat, coap.TextPlain)
					res.SetOption(coap.MaxAge, 60)

					c <-m
					return res

				case "calcPrimeNumberResult":

					PrimeNumsCalculated++
					log.Printf("Calculated: %v",PrimeNumsCalculated)
					if PrimeNumsCalculated==PrimeNumsQty{
						elapsed := time.Since(StartTime)
						saveTaskDuration(int64(elapsed/time.Millisecond), PrimeNumsQty)
						log.Printf("%v prime numbers calculated in %v ms STARTDATE: %v",PrimeNumsQty,(int64(elapsed/time.Millisecond)),StartTime)
					}

					res := &coap.Message{
						Type:      coap.Acknowledgement,
						Code:      coap.Content,
						MessageID: m.MessageID,
						Token:     m.Token,
						Payload:   []byte("2.05"),
					}
					res.SetOption(coap.ContentFormat, coap.TextPlain)

					return res

				case ".well-known":
					if(len(m.Path())>1) && m.Path()[1] == "core"{
						res := coreHandler(m, resourcesToPublish)
						return res
					}else{
						res := notFoundHandler(m)
						return res
					}

				default:
					res := notFoundHandler(m)
					return res

				}
			} else {
				res := notFoundHandler(m)
				return res
			}
			return nil
		})))
}


func notFoundHandler(m *coap.Message) *coap.Message {

	res := &coap.Message{
		Type:      coap.Acknowledgement,
		Code:      coap.NotFound,
		MessageID: m.MessageID,
		Token:     m.Token,
		Payload:   []byte("4.05"),
	}
	res.SetOption(coap.ContentFormat, coap.TextPlain)
	return res

}

func badRequestHandler(m *coap.Message) *coap.Message {

	res := &coap.Message{
		Type:      coap.Acknowledgement,
		Code:      coap.BadRequest,
		MessageID: m.MessageID,
		Token:     m.Token,
		Payload:   []byte("4.00"),
	}
	res.SetOption(coap.ContentFormat, coap.TextPlain)
	return res

}

func coreHandler (m *coap.Message, resources []string) *coap.Message{
	var payloadValue string


	for _, resource := range resources{
		payloadValue += "</" + resource + ">,"
	}
	payloadValue += "</.well-known/core>,"
	res := &coap.Message{
		Type:      coap.Acknowledgement,
		Code:      coap.Content,
		MessageID: m.MessageID,
		Token:     m.Token,
		Payload:   []byte(payloadValue),
	}

	res.SetOption(coap.ContentFormat, coap.AppLinkFormat)

	return res
}


func parsePayload(payload []byte) (int, string){
	splited := strings.Split(string(payload),",")
	qty,_:=strconv.Atoi(splited[0])
	return qty,splited[1]
}

