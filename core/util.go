/**
 @author: RedCrazyGhost
 @date: 2023/2/13

**/

package core

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

// todo 补充开发工具函数

// readCSV 读取CSV文件内容
// 注意！.csv 文件编码集有存在问题
// 统一使用UTF-8编码集
// # 前缀忽略
// todo 思考 输入数据后的处理
func readCSV(filepath string) {
	file, err := os.Open(filepath)
	if err != nil {
		log.Panic("文件名称不对！/解析文件不存在！")
	}
	defer file.Close()

	//r := csv.NewReader(transform.NewReader(file, simplifiedchinese.GB18030.NewDecoder()))
	r := csv.NewReader(file)
	r.Comment = '#'
	r.LazyQuotes = true
	r.ReuseRecord = true
	r.TrimLeadingSpace = true
	for {
		row, err := r.Read()
		if err != nil && err != io.EOF {
			log.Panic("读取文件错误！")
		} else if err == io.EOF {
			log.Println("读取完毕！")
			return
		}
		fmt.Printf("%v\n", row)
	}
}

func NowDateString() string {
	year, month, day := time.Now().Date()
	return fmt.Sprintf("%d-%d-%d", year, month, day)
}
