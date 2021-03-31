/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/1/7 下午8:52
 * Description:
 **/

package main

import (
	"DistributedCache/cacheClient/client"
	"flag"
	"fmt"
)

func main() {
	server := flag.String("h", "localhost", "cache server address")
	op := flag.String("c", "get", "command, could be get/set/del")
	key := flag.String("k", "", "key")
	value := flag.String("v", "", "value")
	typ := flag.String("t", "http", "cache type")

	flag.Parse()

	c := client.New(*typ, *server)
	cmd := &client.Cmd{*op, *key, *value, nil}
	c.Run(cmd)
	if cmd.Error != nil {
		fmt.Println("error", cmd.Error)
	} else {
		fmt.Println(cmd.Value)
	}
}
