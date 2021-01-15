package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// HelloServer prints out tls details
func HelloServer(w http.ResponseWriter, req *http.Request) {
	// for i, cert := range req.TLS.PeerCertificates {
	// 	subject := cert.Subject
	// 	issuer := cert.Issuer
	// 	log.Printf(" %d s:/C=%v/ST=%v/L=%v/O=%v/OU=%v/CN=%s", i, subject.Country, subject.Province, subject.Locality, subject.Organization, subject.OrganizationalUnit, subject.CommonName)
	// 	log.Printf("   i:/C=%v/ST=%v/L=%v/O=%v/OU=%v/CN=%s", issuer.Country, issuer.Province, issuer.Locality, issuer.Organization, issuer.OrganizationalUnit, issuer.CommonName)
	// }
	if req.Header != nil {
		fmt.Printf("%v\n", req.Header)
	}

	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("This is an example server.\n"))
	// fmt.Fprintf(w, "This is an example server.\n")
	// io.WriteString(w, "This is an example server.\n")
}

// func main() {
// 	http.HandleFunc("/hello", HelloServer)
// 	err := http.ListenAndServeTLS(":443", "server.crt", "server.key", nil)
// 	if err != nil {
// 		log.Fatal("ListenAndServe: ", err)
// 	}
// }

func main() {
	http.HandleFunc("/hello", HelloServer)
	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func createServerConfig(ca, crt, key string) (*tls.Config, error) {
	caCertPEM, err := ioutil.ReadFile(ca)
	if err != nil {
		return nil, err
	}

	roots := x509.NewCertPool()
	ok := roots.AppendCertsFromPEM(caCertPEM)
	if !ok {
		panic("failed to parse root certificate")
	}

	cert, err := tls.LoadX509KeyPair(crt, key)
	if err != nil {
		return nil, err
	}
	return &tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    roots,
	}, nil
}

// func main() {
// 	tlsConfig, err := createServerConfig("./ca.crt", "server.crt", "server.key")
// 	if err != nil {
// 		log.Fatal("ListenAndServe: ", err)
// 	}
// 	server := &http.Server{
// 		TLSConfig: tlsConfig,
// 	}
// 	http.HandleFunc("/hello", HelloServer)
// 	err = server.ListenAndServeTLS("", "")
// 	if err != nil {
// 		log.Fatal("ListenAndServe: ", err)
// 	}
// }
