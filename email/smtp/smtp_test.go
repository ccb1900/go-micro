package smtp

import "testing"

func TestEmail_Send(t *testing.T) {
	e := New(&Option{})

	err := e.Send("sdfsdf", nil, nil)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log("pass")
}
