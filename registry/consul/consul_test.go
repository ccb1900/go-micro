package consul

import "testing"

func TestRegistry_Open(t *testing.T) {
	r := New(&Option{
		Address:          "",
		Debug:            false,
		HealthTimeout:    "",
		HealthInterval:   "",
		HealthDeregister: "",
	})
	err := r.HttpRegister("1", "1", "1", 12345, nil)
	if err != nil {
		return
	}
}
