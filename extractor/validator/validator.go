package main

import (
	"fmt"
	"net"

	verification "github.com/docbull/auditchain/verification/proto"
)

type Verifier struct {
	Auditors   []verification.AuditChainInfo
	PublicKeys verification.PublicKey
}

func (verifier *Verifier) Init() {

}

func (verifier *Verifier) OpenSocket(port string) {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println(err)
	}
	defer lis.Close()

	for {
		conn, err := lis.Accept()
		if err != nil {
			fmt.Println(err)
		}
		go verifier.StoreAuditNodeInfo(conn)
	}
}

func (verifier *Verifier) StoreAuditNodeInfo(conn net.Conn) {
	
}

func main() {
	verifier := Verifier{}
	verifier.Init()
	verifier.OpenSocket("31418")
}
