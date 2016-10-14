package fileSvcProvider

import (
	"fmt"

	"git-wip-us.apache.org/repos/asf/thrift.git/lib/go/thrift"
	"github.com/ibrohimislam/tugas-2/services/tugas"
)

type FSServer struct {
	host             string
	handler          *FSHandler
	processor        *tugas.FileSvcProcessor
	transport        *thrift.TServerSocket
	transportFactory thrift.TTransportFactory
	protocolFactory  *thrift.TBinaryProtocolFactory
	server           *thrift.TSimpleServer
}

func NewFSServer(host string) *FSServer {
	handler := NewFSHandler()
	processor := tugas.NewFileSvcProcessor(handler)
	transport, err := thrift.NewTServerSocket(host)
	if err != nil {
		panic(err)
	}

	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	server := thrift.NewTSimpleServer4(processor, transport, transportFactory, protocolFactory)
	return &FSServer{
		host:             host,
		handler:          handler,
		processor:        processor,
		transport:        transport,
		transportFactory: transportFactory,
		protocolFactory:  protocolFactory,
		server:           server,
	}
}

func (ps *FSServer) Run() {
	fmt.Printf("[INFO] server listening on %s\n", ps.host)
	ps.server.Serve()
}

func (ps *FSServer) Stop() {
	fmt.Println("[INFO] stopping server...")
	ps.server.Stop()
}
