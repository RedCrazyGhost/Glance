/**
 @author: RedCrazyGhost
 @date: 2023/2/14

**/

package core

import (
	"fmt"
	"testing"
)

func TestReadCSV(t *testing.T) {
	filename := "test.csv"
	ReadCSV(filename, nil)
}

func TestNowDateString(t *testing.T) {
	fmt.Println(NowDateString())
}
