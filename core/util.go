/**
 @author: RedCrazyGhost
 @date: 2023/2/13

**/

package core

import (
	"crypto/md5"
	"encoding/csv"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"os"
	"time"
)

// ReadCSV 读取CSV文件内容
// 注意！.csv 文件编码集有存在问题
// 统一使用UTF-8编码集
// # 前缀忽略
// todo 思考 输入数据后的处理
func ReadCSV(filepath string, parserName string, flow *TradeFlow) error {
	if len(filepath) <= 0 {
		return errors.New("文件名称不对！")
	}
	if flow == nil {
		return errors.New("交易流不能为空！")
	}
	file, err := os.Open(filepath)
	if err != nil {
		return errors.New("文件不存在！")
	}
	defer file.Close()

	//r := csv.NewReader(transform.NewReader(file, simplifiedchinese.GB18030.NewDecoder()))
	r := csv.NewReader(file)
	r.Comment = StringFirstWordToRune(SystemConfig.Base.CSV.Comment)
	r.LazyQuotes = SystemConfig.Base.CSV.LazyQuotes
	r.ReuseRecord = SystemConfig.Base.CSV.ReuseRecord
	r.TrimLeadingSpace = SystemConfig.Base.CSV.TrimLeadingSpace

	parser := NewParser(parserName)

	title := make([]string, 0)
	isFrist := false
	for {
		row, err := r.Read()
		if err != nil && err != io.EOF {
			SystemLogPool.Println("读取文件错误！")
			return err
		} else if err == io.EOF {
			GloabalLog.Println("读取完毕！")
			return nil
		}

		if !isFrist {
			title = append(title, row...)
			isFrist = !isFrist
		} else {
			metadata := make([]string, 0, len(title))
			metadata = append(metadata, row...)
			node, err := NewTradeNode(parser, title, metadata...)
			if err != nil {
				FailLog.Println(err.Error())
			} else {
				err = flow.insert(node, 0)
				if err != nil {
					return err
				}
			}
		}
	}
}

// NowDateString 返回当前日期字符串
func NowDateString() string {
	return time.Now().Format("2006-01-02")
}

// NowDateTimeString 返回当前日期时间字符串
func NowDateTimeString() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// MD5 hash算法加密
func MD5(sum []byte, value ...string) string {
	hash := md5.New()
	hash.Write([]byte(NowDateTimeString() + fmt.Sprintf("%v", value)))
	return hex.EncodeToString(hash.Sum(sum))
}

// RemoveRepeatedElement 去除重复的channel元素
func RemoveRepeatedElement(channels []TradeChannel) []TradeChannel {
	m := make(map[TradeChannel]int)
	for _, channel := range channels {
		m[channel] = 1
	}
	if len(m) > 0 {
		tradeChannels := make([]TradeChannel, 0)
		for key := range m {
			tradeChannels = append(tradeChannels, key)
		}
		return tradeChannels
	}
	return nil
}

// readParserFile
func readParserFile() {

}

func StringFirstWordToRune(str string) rune {
	return []rune(str)[0]
}
