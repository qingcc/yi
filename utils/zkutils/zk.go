package zkutils

import (
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"log"
	"sync"
	"time"
)

var(
	zkOnce sync.Once
	zkConn *zk.Conn
	//zkList = []string{"47.112.210.86:6031", "47.112.210.86:6032", "47.112.210.86:6033"}
	//zkList = []string{"localhost:2181"}
	zkList = []string{"47.112.210.86:6035"}
	client = &SdClient{zkServers:zkList, timeout:time.Second * 10, conn:zkConn}
)

func newConnect() {
	conn, _, err := zk.Connect(client.zkServers, client.timeout)
	if err != nil {
		fmt.Println("connect to zk failed:", err)
	}
	client.conn = conn
	client.createNode(client.zkRoot, []byte(""))
	return
}

func GetConn() *zk.Conn {
	zkOnce.Do(func() {
		newConnect()
	})
	return zkConn
}

func (c *SdClient)createNode(path string, data []byte) (value string, err error) {
	var exists bool
	if exists, err = c.exists(path); err != nil {
		return
	}
	if !exists {
		log.Println("create node")
		value, err = c.conn.Create(path, data, 0, zk.WorldACL(zk.PermAll))
	}
	return
}

func (c *SdClient)exists(path string) (exists bool, err error) {
	exists, _, err = c.conn.Exists(path)
	return
}

func (c *SdClient)Create(path string, data []byte) (value string, err error) {
	return c.createNode(path, data)
}

func (c *SdClient)Children(path string) ([]string, *zk.Stat, error) {
	return c.conn.Children(path)
}




type ZkSession struct {
	SdClient
}

func NewConn() *ZkSession {
	return &ZkSession{*client}
}
func (client *ZkSession)SetZKList(zklist []string) *ZkSession {
	client.zkServers = zklist
	return client
}
func (client *ZkSession)SetTimeout(timeout time.Duration) *ZkSession {
	client.timeout = timeout
	return client
}
func (client *ZkSession)SetZkroot(zkroot string) *ZkSession {
	client.zkRoot = zkroot
	return client
}
func (c *ZkSession)Set() *zk.Conn {
	client = &c.SdClient
	return GetConn()
}