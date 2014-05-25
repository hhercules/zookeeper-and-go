package main

import (
	"fmt"
	"os"
	"strings"
	"time"
	"github.com/samuel/go-zookeeper/zk"
)

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func connect() *zk.Conn {
	zksStr := os.Getenv("ZOOKEEPER_SERVERS")
	zks := strings.Split(zksStr, ",")
	conn, _, err := zk.Connect(zks, time.Second)
	must(err)
	return conn
}

func main() {
	conn1 := connect()
	defer conn1.Close()

	acl := zk.WorldACL(zk.PermAll)
	_, err = conn1.Create("/ephemeral", []byte("here"), zk.FlagEphemeral, acl)
	must(err)

	conn2 := connect()
	defer conn2.Close()

	exists, _, err := conn2.Exists("/ephemeral")
	must(err)
	fmt.Printf("before disconnect: %+v\n", exists)

	conn1.Close()
	time.Sleep(time.Second * 2)

	exists, _, err = conn2.Exists("/ephemeral")
	must(err)
	fmt.Printf("after disconnect: %+v\n", exists)
}