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
	filename1 := "/Users/wenxingzhan/Desktop/202211_2088132603923197/test.csv"
	filename2 := "/Users/wenxingzhan/Desktop/202211_2088132603923197/20881326039231970156_202211_账务明细_1.csv"
	readCSV(filename1)
	readCSV(filename2)
}

func TestNowDateString(t *testing.T) {
	fmt.Printf(NowDateString())
}
