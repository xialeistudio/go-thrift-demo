package main

import (
	"git.apache.org/thrift.git/lib/go/thrift"
	"net"
	"log"
	"example"
)

const (
	host = "localhost"
	port = "9090"
)

func main() {
	socket, err := thrift.NewTSocket(net.JoinHostPort(host, port))
	if err != nil {
		log.Fatalln(err)
	}
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	transport := transportFactory.GetTransport(socket)
	if err != nil {
		log.Fatalln(err)
	}
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	client := example.NewFormatDataClientFactory(transport, protocolFactory)
	if err := transport.Open(); err != nil {
		log.Fatalln(err)
	}
	defer transport.Close()

	data := &example.Data{Text: "Hello World"}
	d, err := client.DoFormat(data)
	log.Println(d.Text)
}
