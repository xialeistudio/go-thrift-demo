package main

import (
	"thrift-demo/example"
	"strings"
	"git.apache.org/thrift.git/lib/go/thrift"
	"log"
)

type FormatDataImpl struct {
}

func (s *FormatDataImpl) DoFormat(data *example.Data) (r *example.Data, err error) {
	var rData example.Data
	rData.Text = strings.ToUpper(data.Text)
	return &rData, nil
}

const (
	host = "localhost"
	port = "9090"
)

func main() {
	handler := &FormatDataImpl{}
	processor := example.NewFormatDataProcessor(handler)
	serverTransport, err := thrift.NewTServerSocket(host + ":" + port)
	if err != nil {
		log.Fatalln(err)
	}
	transport := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocol := thrift.NewTBinaryProtocolFactoryDefault()
	server := thrift.NewTSimpleServer4(processor, serverTransport, transport, protocol)
	log.Println("Running at: " + host + ":" + port)
	log.Fatalln(server.Serve())
}
