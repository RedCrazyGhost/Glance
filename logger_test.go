/**
 @author: RedCrazyGhost
 @date: 2023/2/14

**/

package main

import (
	"fmt"
	"testing"
)

func TestLog(t *testing.T) {
	log := NewLog("test")

	log.Print(0)
	log.Println(1)
	log.Printf("-%d-", 2)
}

func TestNewLogPool(t *testing.T) {
	pool := NewLogPool()
	pool.append(NewLog("1826test"))
	l := pool.indexOf(0)
	l.Println("ok")
	fmt.Printf("%s", l.name)
	pool.show()
}
