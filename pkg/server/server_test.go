package server

import "testing"

func TestCreateServer(t *testing.T) {
	_ = CreateServer("test-service")
}
