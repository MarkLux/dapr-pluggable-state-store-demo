package main

import "dapr-pluggable-state-store-demo/server"

func main() {
	s := &server.DemoStateStoreServer{
		BindPort: "8099",
	}
	s.Serve()
}
