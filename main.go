package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/docker/machine/libmachine/auth"
	"github.com/docker/machine/libmachine/cert"
)

var ip = flag.String("ip", "", "IP of the server")

func main() {
	pwd, _ := os.Getwd()
	output := pwd + "/output"

	flag.Parse()

	if (*ip) == "" {
		fmt.Println("USAGE: \n\ngenerate-docker-certs -ip <IP>\n")
		os.Exit(1)
	}

	authOptions := &auth.Options{
		CertDir:          output,
		CaCertPath:       output + "/ca.pem",
		CaPrivateKeyPath: output + "/ca-key.pem",
		ClientCertPath:   output + "/cert.pem",
		ClientKeyPath:    output + "/key.pem",
		ServerCertPath:   output + "/server.pem",
		ServerKeyPath:    output + "/server-key.pem",
	}

	if err := cert.BootstrapCertificates(authOptions); err != nil {
		fmt.Println(err)
	}

	err := cert.GenerateCert(&cert.Options{
		Hosts:     []string{*ip, "localhost"},
		CertFile:  authOptions.ServerCertPath,
		KeyFile:   authOptions.ServerKeyPath,
		CAFile:    authOptions.CaCertPath,
		CAKeyFile: authOptions.CaPrivateKeyPath,
		Org:       "example.com",
		Bits:      2048,
	})

	if err != nil {
		fmt.Println("error generating server cert: %s", err)
	} else {
		fmt.Println("Success!")
	}
}
