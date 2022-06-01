package smtp

import (
	"fmt"
	"testing"
	"time"
)

func TestEmail_Send(t *testing.T) {
	e := New(&Option{})

	err := e.Send("sdfsdf", nil, nil)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log("pass")

	//2022-05-17T09:15:24.480Z
	var zone = time.FixedZone("CST", 0)
	time.Local = zone
	fmt.Println(time.Now().Format(time.RFC3339Nano))
}
