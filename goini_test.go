package goini

import (
	"fmt"
	"testing"
)

func Test_Config(t *testing.T) {
	conf := NewConfig("/home/xyq/test.conf")
	host, err := conf.GetString("DEFAULT", "host")
	fmt.Println(err)
	fmt.Println(host)
}
