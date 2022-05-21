package main

import (
	"fmt"
	"io"
	"net"

	auditchain "github.com/docbull/ethaudit/auditchain/proto"
	proto "github.com/golang/protobuf/proto"
)

func ReceiveBlockInfo(conn net.Conn) {
	data := make([]byte, 0)
	buf := make([]byte, 1024)
	length := 0

	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
				return
			}
			break
		}

		data = append(data, buf[:n]...)
		length += n
	}

	auditBlock := &auditchain.BlockInfo{}
	err := proto.Unmarshal(data[:length], auditBlock)
	if err != nil {
		fmt.Println(err)
		conn.Write([]byte("error"))
		return
	}
	fmt.Println("difficulty:", auditBlock.Difficulty)
	fmt.Println("extraData:", auditBlock.Extra)
	fmt.Println("gasLimit:", auditBlock.GasLimit)
	fmt.Println("gasUsed:", auditBlock.GasUsed)
	fmt.Println("hash:", auditBlock.TxHash)
	fmt.Println("logsBloom:", auditBlock.Bloom)
	fmt.Println("miner:", auditBlock.Coinbase)
	fmt.Println("mixHash:", auditBlock.MixDigest)
	fmt.Println("nonce:", auditBlock.Nonce)
	fmt.Println("number:", auditBlock.Number)
	fmt.Println("parentHash:", auditBlock.ParentHash)
	fmt.Println("receiptsRoot:", auditBlock.ReceiptHash)
	fmt.Println("sha3Uncles:", auditBlock.UncleHash)
	fmt.Println("stateRoot:", auditBlock.Root)
	fmt.Println("timestamp:", auditBlock.Time)
	fmt.Println("transactionRoot:", auditBlock.TxHash)
}

func Start2Listen() {
	lis, err := net.Listen("tcp", ":4242")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer lis.Close()

	fmt.Println("Auditchain ...")
	for {
		conn, err := lis.Accept()
		fmt.Println("SOME PEER CONNECTED!")
		if err != nil {
			fmt.Println(err)
			return
		}

		go ReceiveBlockInfo(conn)
	}
}

func main() {
	Start2Listen()
}
