// Issue 89
// Should avoid Passing hard coded credential into tls.X509KeyPair

package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"os"
)

func main() {
	certPEM := []byte(os.Getenv("CertPem"))
	keyPEM := []byte("-----BEGIN PRIVATE KEY-----" +
		"MIGHAgEAMBMGByqGSM49AgEGCCqGSM49AwEHBG0wawIBAQQgUwukxU6jIyoZZZ1U" +
		"uoC8W9S0QQvfehNc7NFnLTr8WFKhRANCAATMqlKaWUafyYeUviY7iwSMoMULZ7er" +
		"P/0PZQ/uiw5dyZmIpPI2k4661Kkvb01w3/F+WMqAUVWyNb0G9ntUl+HA" +
		"-----END PRIVATE KEY-----")

	cert, err := tls.X509KeyPair(certPEM, keyPEM)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(cert)
}
