package main

import (
	"fmt"

	"github.com/panjf2000/gnet/v2"
	"github.com/sirupsen/logrus"
)

type Client struct {
	gnet.BuiltinEventEngine
}

func (*Client) OnTraffic(c gnet.Conn) gnet.Action {
	buf, _ := c.Next(-1)
	c.Write(buf)
	return gnet.None
}

var client, _ = gnet.NewClient(&Client{}, gnet.WithLogger(logrus.StandardLogger()))

type Server struct {
	gnet.BuiltinEventEngine

	addr      string
	multicore bool
}

func (s *Server) OnBoot(eng gnet.Engine) gnet.Action {
	logrus.Info(fmt.Sprintf("echo server with multi-core=%t is listening on %s", s.multicore, s.addr))
	return gnet.None
}

func (*Server) OnOpen(conn gnet.Conn) (out []byte, action gnet.Action) {
	logrus.Info("OnOpen", conn)

	return
}

func (*Server) OnTraffic(c gnet.Conn) gnet.Action {
	buf, _ := c.Next(-1)

	cli, err := client.Dial("tcp", "1.1.1.1")
	if err != nil {
		logrus.Error(err)
	}
	cli.Write(buf)

	return gnet.None
}

func ServerPool(cycle Cycle) {
	var multicore bool = true

	server := &Server{addr: fmt.Sprintf("tcp://:%d", cycle.Localport), multicore: multicore}

	logrus.Info("start server")
	logrus.Fatal(gnet.Run(server, server.addr, gnet.WithMulticore(multicore), gnet.WithLogger(logrus.StandardLogger())))
}
