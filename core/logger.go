/**
 @author: RedCrazyGhost
 @date: 2023/2/14

**/

package core

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type LogPool struct {
	Logs []*Log
}

type Log struct {
	name string
	l    *log.Logger
	f    *os.File
}

var SystemLogPool *LogPool

var GloabalLog *Log
var FailLog *Log

func InitLog() {
	err := os.MkdirAll("./log/"+NowDateString(), os.ModePerm)
	if err != nil {
		GloabalLog.Fatalln("日志文件初始化失败")
	}
	SystemLogPool = NewLogPool()
	GloabalLog = NewLog("Gloabal")
	SystemLogPool.append(GloabalLog)
	FailLog = NewLog("Fail")
	SystemLogPool.append(FailLog)

	GloabalLog.Println(NewLogMessage("logPool", "初始化完成！"))
}

// NewLogPool 创建LogPool
func NewLogPool() *LogPool {
	p := &LogPool{}
	p.Logs = make([]*Log, 0)
	return p
}

// indexOf 根据index获取Log
func (p *LogPool) indexOf(index int) *Log {
	return p.Logs[index]
}

// getLog 通过logName获取Log
func (p *LogPool) getLog(logName string) (*Log, error) {
	for _, l := range p.Logs {
		if strings.EqualFold(logName, l.name) {
			return l, nil
		}
	}
	return nil, errors.New("对象不存在！")
}

// append 添加Log
func (p *LogPool) append(l *Log) {
	p.Logs = append(p.Logs, l)
}

// closeLog 根据logName关闭Log
func (p *LogPool) closeLog(logName string) error {
	closeLog, err := p.getLog(logName)
	if err != nil {
		GloabalLog.Println(NewLogMessage(logName, err.Error()))
	}
	err = closeLog.close()
	if err != nil {
		return err
	}
	err = p.deleteLog(closeLog)
	if err != nil {
		FailLog.Println(NewLogMessage("LogPool", err.Error()))
		return err
	}
	return nil
}

// logOfIndex 根据Log查找在LogPool中的index
func (p *LogPool) logOfIndex(target *Log) (int, error) {
	for i, l := range p.Logs {
		if target == l {
			return i, nil
		}
	}
	return -1, errors.New("日志池中不存在" + target.name + "日志")
}

// deleteLog 通过index删除Log
func (p *LogPool) deleteLog(l *Log) error {
	index, err := p.logOfIndex(l)
	if err != nil {
		return err
	}
	p.Logs = append(p.Logs[:index], p.Logs[index+1:]...)
	return nil
}

// closeAll 关闭所有日志文件流
// 最后一个关闭GlobalLog日志
func (p *LogPool) CloseAll() error {
	Len := len(p.Logs)

	if Len > 2 {
		for i := 2; i < Len; i++ {
			err := p.Logs[i].close()
			if err != nil {
				continue
			}
			err = p.deleteLog(p.Logs[i])
			if err != nil {
				FailLog.Println(NewLogMessage("LogPool", err.Error()))
				return err
			}
		}
	}
	for i := Len - 1; i > -1; i-- {
		err := p.Logs[i].close()
		if err != nil {
			continue
		}
		err = p.deleteLog(p.Logs[i])
		if err != nil {
			FailLog.Println(NewLogMessage("LogPool", err.Error()))
			return err
		}
	}
	return nil

}

// NewLog 实例化Log
func NewLog(name string) *Log {
	l := &Log{name: name}
	filepath := fmt.Sprintf("./log/%s/%s.log", NowDateString(), name)
	file, err := os.OpenFile(filepath, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	l.f = file
	if err != nil {
		l.l.Fatalln(NewLogMessage(name, "日志初始化失败!"))
	}
	multiWriter := io.MultiWriter(os.Stdout, file)
	l.l = log.New(multiWriter, "", log.LstdFlags|log.Lshortfile)
	return l
}

// close 关闭文件流
func (l *Log) close() error {
	err := l.f.Close()
	if err != nil {
		GloabalLog.l.Print(NewLogMessage(l.name, "文件流关闭失败！原因->"+err.Error()))
		return err
	}
	GloabalLog.l.Print(NewLogMessage(l.name, "文件流关闭成功！"))
	return nil
}

// Println 日志池内所有Log打印日志
func (p *LogPool) Println(v ...any) {
	for _, l := range p.Logs {
		l.Println(fmt.Sprint(v...))
	}
}

// Printf 封装log.Logger的Printf
func (l *Log) Printf(format string, v ...any) {
	l.l.Printf(fmt.Sprintf(format, v...))
}

// Println 封装log.Logger的Println
func (l *Log) Println(v ...any) {
	l.l.Println(fmt.Sprint(v...))
}

// Print 封装log.Logger的Print
func (l *Log) Print(v ...any) {
	l.l.Print(fmt.Sprint(v...))
}

// Fatalf 封装log.Logger的Fatalf
func (l *Log) Fatalf(format string, v ...any) {
	l.l.Fatalf(fmt.Sprintf(format, v...))
}

// Fatalln 封装log.Logger的Fatalln
func (l *Log) Fatalln(v ...any) {
	l.l.Fatalln(fmt.Sprint(v...))
}

// Fatal 封装log.Logger的Fatal
func (l *Log) Fatal(v ...any) {
	l.l.Fatal(fmt.Sprint(v...))
}

// Panicf 封装log.Logger的Panicf
func (l *Log) Panicf(format string, v ...any) {
	l.l.Panicf(fmt.Sprintf(format, v...))
}

// Panicln 封装log.Logger的Panicln
func (l *Log) Panicln(v ...any) {
	l.l.Panicln(fmt.Sprint(v...))
}

// Panic 封装log.Logger的Panic
func (l *Log) Panic(v ...any) {
	l.l.Panic(fmt.Sprint(v...))
}
