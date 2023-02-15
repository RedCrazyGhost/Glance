/**
 @author: RedCrazyGhost
 @date: 2023/2/14

**/

package core

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
	pool.append(NewLog("1830test"))
	l := pool.indexOf(0)
	l.Println("ok")
	fmt.Printf("%s\n", l.name)
	err := pool.closeLog("1826test")
	if err != nil {
		return
	}
	err = pool.closeAll()
	if err != nil {
		return
	}
}
