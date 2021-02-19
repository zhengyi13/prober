package prober

import (
	"crypto/tls"
	"log"
)


// Basic type is a HostPort string like "localhost:8443"
type HostPort string

func Probe(hp HostPort) (timestamp int64, err error) {
	// Given a HostPort, attempt to connect via SSL.
	// On failure, return an error
	// On success, return the certificate's expiration date as a Unix timestamp
	config := tls.Config{
		InsecureSkipVerify: true,
	}
	conn, err := tls.Dial("tcp", string(hp), &config)
	if err != nil {
		log.Printf("ERROR dial: %s\n", err)
		return timestamp, err
	}
	defer conn.Close()
	log.Printf("client: connected to %s\n ", conn.RemoteAddr())
	s := conn.ConnectionState()
	for _, cert := range s.PeerCertificates {
		log.Printf("CN: %s\n", cert.Subject)
		timestamp = cert.NotAfter.Unix() // thus we naturally always get the last timestamp
	}
	return timestamp, err
}
