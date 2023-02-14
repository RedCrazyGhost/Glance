/**
 @author: RedCrazyGhost
 @date: 2023/2/14

**/

package main

import (
	"Glance/core"
	"fmt"
	"io"
	"log"
	"os"
)

type LogPool struct {
	Logs []*Log
	Len  int
}

type Log struct {
	name string
	l    *log.Logger
	f    *os.File
}

var logPool *LogPool

var gloabalLog *Log
var failLog *Log

func initLog() {
	err := os.MkdirAll("./log/"+core.NowDateString(), os.ModePerm)
	if err != nil {
		gloabalLog.Fatalln("日志文件初始化失败")
	}
	logPool = NewLogPool()
	gloabalLog = NewLog("gloabal")
	logPool.append(gloabalLog)
	failLog = NewLog("fail")
	logPool.append(failLog)

	gloabalLog.Printf("Glance Version:%s\n", VERSION)
}

func NewLogPool() *LogPool {
	p := &LogPool{}
	p.Logs = make([]*Log, 0)
	p.Len = 2
	return p
}

func (p *LogPool) indexOf(index int) *Log {
	return p.Logs[index]
}
func (p *LogPool) show() {
	fmt.Println(logPool)
}
func (p *LogPool) append(l *Log) {
	p.Logs = append(p.Logs, l)
	p.Len = len(p.Logs)
}

func NewLog(name string) *Log {
	l := &Log{name: name}
	filepath := fmt.Sprintf("./log/%s/%s.log", core.NowDateString(), name)
	f, err := os.OpenFile(filepath, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	if err != nil {
		l.l.Fatalf("%s日志初始化失败", name)
	}
	multiWriter := io.MultiWriter(os.Stdout, f)
	l.l = log.New(multiWriter, "", log.LstdFlags)
	return l
}

func (l *Log) Printf(format string, v ...any) {
	l.l.Printf(fmt.Sprintf(format, v...))
}

func (l *Log) Println(v ...any) {
	l.l.Println(fmt.Sprintln(v...))
}
func (l *Log) Print(v ...any) {
	l.l.Print(fmt.Sprintln(v...))
}
func (l *Log) Fatalf(format string, v ...any) {
	l.l.Fatalf(fmt.Sprintf(format, v...))
}
func (l *Log) Fatalln(v ...any) {
	l.l.Fatalln(fmt.Sprintln(v...))
}
func (l *Log) Fatal(v ...any) {
	l.l.Fatal(fmt.Sprintln(v...))
}

func (l *Log) Panicf(format string, v ...any) {
	l.l.Panicf(fmt.Sprintf(format, v...))
}
func (l *Log) Panicln(v ...any) {
	l.l.Panicln(fmt.Sprintln(v...))
}
func (l *Log) Panic(v ...any) {
	l.l.Panic(fmt.Sprintln(v...))
}
