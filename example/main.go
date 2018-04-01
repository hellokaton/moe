package main

import (
	"time"

	moe "github.com/biezhi/moe"
)

func main() {

	for name := range moe.SpinnerMap {
		func() {
			moe := moe.New("正在为您加载").Spinner(name).Start()
			time.Sleep(2 * time.Second)
			defer moe.Stop()
		}()
	}

	// moe := moe.New("正在为您加载").Color("blue")
	// moe.Start()
	// time.Sleep(2 * time.Second)

	// moe.Text("稍后看到精彩内容..").Spinner("clock")
	// time.Sleep(2 * time.Second)

	// moe.Text("这或许需要点儿时间..").Color("yellow").Spinner("bouncingBall")
	// time.Sleep(2 * time.Second)

	// moe.Stop()
}
