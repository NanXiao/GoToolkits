package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"strconv"
)

type tcpConnType int

const (
	client tcpConnType = iota
	mnp
)

const (
	countNum int = 3000
)

const (
	clientMSGRspLen uint32 = 64
	mnpMSGRspLen    uint32 = 64
)

const (
	clientMSGBaseLen          uint32 = 10
	clientMSGTotalLenParamLen uint32 = 4
	clientMSGReqTypeParamLen  uint32 = 2
	clientMSGMaxParamLen      uint32 = 1024
)

const (
	clientMSGReqTypeGetVersion uint16 = 3
)

type clientMSG struct {
	totalLen uint32
	reqType  uint16
	seqID    uint32
	param    [clientMSGMaxParamLen]byte
}

const (
	netProtocol string = "tcp"
	colon       string = ":"
	ipAddr      string = "192.168.23.192"
)

func main() {
	clientRspChan := tcpConnRoutine(client)
	mnpRspChan := tcpConnRoutine(mnp)

	i := <-clientRspChan
	j := <-mnpRspChan

	fmt.Println(i)
	fmt.Println(j)

	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func tcpConnRoutine(connType tcpConnType) chan int {
	var (
		port    int64
		rspChan chan int
	)

	switch connType {
	case client:
		port = 5678
	case mnp:
		port = 1213
	default:
		return nil
	}

	rspChan = make(chan int)

	tcpAddr, err := net.ResolveTCPAddr(netProtocol, string(strconv.AppendInt(append([]byte(ipAddr), colon...), port, 10)))
	checkError(err)

	switch connType {
	case client:
		go clientRoutine(tcpAddr, rspChan)
	case mnp:
		go mnpRoutine(tcpAddr, rspChan)
	default:
		return nil
	}

	return rspChan
}

func mnpRoutine(tcpAddr *net.TCPAddr, rspChan chan int) {
	tcpAck := make([]byte, mnpMSGRspLen)

	for i := 0; i < countNum; i++ {
		tcpConn, err := net.DialTCP("tcp", nil, tcpAddr)
		checkError(err)

		_, err = tcpConn.Write([]byte("11400123456\n"))
		checkError(err)

		_, err = tcpConn.Read(tcpAck)
		checkError(err)

		tcpConn.Close()
	}

	rspChan <- 1
}

func clientRoutine(tcpAddr *net.TCPAddr, rspChan chan int) {
	var msg clientMSG

	req := make([]byte, (clientMSGBaseLen + clientMSGMaxParamLen))
	rsp := make([]byte, clientMSGRspLen)

	msg.totalLen = clientMSGBaseLen
	msg.reqType = clientMSGReqTypeGetVersion

	binary.BigEndian.PutUint32(req, msg.totalLen)
	binary.BigEndian.PutUint16(req[clientMSGTotalLenParamLen:], msg.reqType)

	for i := 0; i < countNum; i++ {
		tcpConn, err := net.DialTCP("tcp", nil, tcpAddr)
		checkError(err)

		_, err = tcpConn.Write(req[0:clientMSGBaseLen])
		checkError(err)

		_, err = tcpConn.Read(rsp)
		checkError(err)

		tcpConn.Close()
	}

	rspChan <- 1
}
