package graceful

import (
	"fmt"
	"net"
	"testing"
	"time"
)

func TestShutdown(t *testing.T) {
	fmt.Println("start...")
	t.Log("start...")
	l, err := net.Listen("tcp", ":18801")
	if err != nil {
		t.Errorf("err : %v", err)
		return
	}
	t.Log("enter...")
	fmt.Println("enter...")
	//var lock sync.Mutex
	//closed := false
	Shutdown(func() {
		for {
			t.Log("wait...")
			fmt.Println("wait...")
			accept, err := l.Accept()
			if err != nil {
				t.Errorf("接收 err : %v", err)
				return
			}
			fmt.Println("收到连接...")
			go func(accept net.Conn) {

			}(accept)
		}
	}, func() {
		time.Sleep(3 * time.Second)
		err := l.Close()
		if err != nil {
			t.Errorf("关闭 err : %v", err)
			return
		}
		fmt.Println("关闭成功")
	})
}
