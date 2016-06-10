package main

import (
	"flag"
	"github.com/paradoxxl/gonnect/server"
)

var genCert = flag.Bool("c", false, "genCert")
var keyFile = flag.String("kf", "priv.key", "keyFile")
var pemFile = flag.String("pf", "pub.pem", "pemFile")
var port = flag.Int("p", 8758, "port")


func main() {

	flag.Parse()

	if *genCert {
		server.GenCert(server.DefaultCert, *keyFile, *pemFile)
	}
	srv := server.NewGonectServer(*pemFile, *keyFile, *port)
	defer srv.Close()

}
