package main

import "dapr-pluggable-state-store-demo/server"

func main() {
	s := &server.DemoStateStoreServer{
		SockPath: "/var/sock/dapr-demo-state-store.sock",
	}
	s.Serve()
}
