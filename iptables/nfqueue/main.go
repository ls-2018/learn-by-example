package main

import (
	"context"
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	fnfqueue "github.com/florianl/go-nfqueue"
	snfqueue "github.com/subgraph/go-nfnetlink/nfqueue"
)

type packet []byte

func (p packet) srcIP() net.IP {
	return net.IP(p[12:16])
}

func (p packet) dstIP() net.IP {
	return net.IP(p[16:20])
}

func (p packet) srcPort() uint16 {
	tcphdr := p[20:]
	return binary.BigEndian.Uint16(tcphdr[:2])
}

func (p packet) dstPort() uint16 {
	tcphdr := p[20:]
	return binary.BigEndian.Uint16(tcphdr[2:4])
}

func handlePacket(q *fnfqueue.Nfqueue, a fnfqueue.Attribute) int {
	if a.Payload != nil && len(*a.Payload) != 0 {
		pkt := packet(*a.Payload)
		fmt.Printf("tcp connect: %s:%d -> %s:%d\n", pkt.srcIP(), pkt.srcPort(), pkt.dstIP(), pkt.dstPort())
	}
	_ = q.SetVerdict(*a.PacketID, fnfqueue.NfAccept)
	return 0
}

func Method1() func() {
	cfg := fnfqueue.Config{
		NfQueue:     1,
		MaxQueueLen: 2,
		Copymode:    fnfqueue.NfQnlCopyPacket,
	}

	nfq, err := fnfqueue.Open(&cfg)
	if err != nil {
		fmt.Println("failed to open nfqueue, err:", err)
		os.Exit(1)
	}

	if err := nfq.RegisterWithErrorFunc(context.Background(), func(a fnfqueue.Attribute) int {
		return handlePacket(nfq, a)
	}, func(e error) int {
		return 0
	}); err != nil {
		fmt.Println("failed to register handlers, err:", err)
		os.Exit(1)
	}

	return func() {
		nfq.Close()
	}
}

func Method2() func() {
	q := snfqueue.NewNFQueue(1)

	ps, err := q.Open()
	if err != nil {
		fmt.Printf("Error opening NFQueue: %v\n", err)
		os.Exit(1)
	}

	for p := range ps {
		networkLayer := p.Packet.NetworkLayer()
		ipsrc, ipdst := networkLayer.NetworkFlow().Endpoints()

		transportLayer := p.Packet.TransportLayer()
		tcpsrc, tcpdst := transportLayer.TransportFlow().Endpoints()

		fmt.Printf("A new tcp connection will be established: %s:%s -> %s:%s\n", ipsrc, tcpsrc, ipdst, tcpdst)
		p.Accept()
	}
	return func() {
		q.Close()
	}
}
func main() {
	ex := Method2()

	// 创建一个信号通道
	sigs := make(chan os.Signal, 1)
	// 注册要接收的信号
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	sig := <-sigs
	fmt.Println("接收到信号:", sig)
	if ex != nil {
		ex() // 一定要关闭，不然无法再次 ssh 登录
	}
}
