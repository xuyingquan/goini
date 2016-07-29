package goini

import (
	"fmt"
	"testing"
)

func Test_Config(t *testing.T) {
	conf := SetConfig("/test.conf")
	host := conf.GetValue("DEFAULT", "host")
	fmt.Println(host)
}
