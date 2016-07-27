package heiyo

import (
	"testing"
)

func TestNewHYConn(t *testing.T) {
	hy_conn1 := NewHYConn()
	hy_conn2 := NewHYConn()
	hy_conn3 := NewHYConn()
	t.Log(hy_conn1.GetTag())
	t.Log(hy_conn3.GetTag())

	t.Log(hy_conn1)
	t.Log(hy_conn2)
	t.Log(hy_conn3)
}
