package main

import (
	"encoding/hex"
	"fmt"
	"io"
	"net"

	extractor "github.com/docbull/ethereum-block-extractor/proto"
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

	block := &extractor.BlockInfo{}
	err := proto.Unmarshal(data[:length], block)
	if err != nil {
		fmt.Println(err)
		conn.Write([]byte("error"))
		return
	}

	fmt.Println("difficulty:", block.Difficulty)
	fmt.Println("extraData:", hex.EncodeToString(block.Extra))
	fmt.Println("gasLimit:", block.GasLimit)
	fmt.Println("gasUsed:", block.GasUsed)
	fmt.Println("hash:", hex.EncodeToString(block.TxHash))
	fmt.Println("logsBloom:", hex.EncodeToString(block.Bloom))
	fmt.Println("miner:", hex.EncodeToString(block.Coinbase))
	fmt.Println("mixHash:", hex.EncodeToString(block.MixDigest))
	fmt.Println("nonce:", block.Nonce)
	fmt.Println("number:", block.Number)
	fmt.Println("parentHash:", hex.EncodeToString(block.ParentHash))
	fmt.Println("receiptsRoot:", hex.EncodeToString(block.ReceiptHash))
	fmt.Println("sha3Uncles:", hex.EncodeToString(block.UncleHash))
	fmt.Println("stateRoot:", hex.EncodeToString(block.Root))
	fmt.Println("timestamp:", block.Time)
	fmt.Println("transactionRoot:", hex.EncodeToString(block.TxHash))
}

func Start2Listen() {
	lis, err := net.Listen("tcp", ":4242")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer lis.Close()

	fmt.Println("🚀 Ethereum Block Extractor ...")
	for {
		conn, err := lis.Accept()
		fmt.Println("🎁 New Block Arrived!")
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
