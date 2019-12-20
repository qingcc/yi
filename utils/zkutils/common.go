package zkutils

import (
	"github.com/samuel/go-zookeeper/zk"
	"time"
)

type ServiceNode struct {
	Name string `json:"name"`
	Host string `json:"host"`
	Port int `json:"port"`
}

type SdClient struct{
	zkServers []string	// 多个节点地址
	zkRoot string		// 服务根节点
	conn *zk.Conn		// zk的客户端连接
	timeout time.Duration
}
