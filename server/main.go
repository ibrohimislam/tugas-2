package main

import "github.com/ibrohimislam/tugas-2/server/fileSvcProvider"

func main() {
	host := "localhost:9090"
	server := fileSvcProvider.NewFSServer(host)
	server.Run()
}
