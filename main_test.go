package main

import (
	"fmt"
	"testing"
)

func TestHandleCmd(t *testing.T){
	fmt.Println(HandleCmd("keys * -sort=key only=string,zset expect=hash"))
}