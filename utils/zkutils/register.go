package zkutils

import (
	"encoding/json"
)


//func NewClient(zkList []string, zkRoot string, timeout time.Duration) (client *SdClient) {
//	client = NewConn().SetZKList(zkList).SetZkroot(zkRoot).SetTimeout(timeout).Set()
//	return
//}

func Register(node *ServiceNode) (err error) {
	path := client.zkRoot + "/" + node.Name
	if data, e := json.Marshal(node); e == nil {
		_, err = client.createNode(path, data)
	}else {
		err = e
	}
	return
}
