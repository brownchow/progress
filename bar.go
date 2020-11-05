package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// 以下载文件为例，cur 是当前已下载的字节数byte，total是总字节数byte
type Bar struct {
	percent      uint64 // 已下载大小的占总大小的比重
	cur          uint64 // 已下载字节数
	total        uint64 // 总共需要下载的字节数
	rate         string // 进度条，就是打印出来的条
	symbol       string // symbol, you can choose #，=， @
	totalSymbols uint64
}

// 构造函数
func (bar *Bar) NewBar(cur, total uint64) {
	bar.cur = cur
	bar.total = total
	bar.totalSymbols = 50
	if bar.symbol == "" {
		bar.symbol = "█"
	}
	bar.percent = bar.getPercent()
	//
	//       Bytes Downloaded
	//+---------------------------+ * Total Symbols =  current progress(the symbols to display)
	//    total Bytes to download
	//
	bar.rate = strings.Repeat(bar.symbol, int(bar.percent*bar.totalSymbols/100))
}

// 指定符号
func (bar *Bar) NewBarWithSymbol(cur, total uint64, symbol string) {
	bar.symbol = symbol
	bar.NewBar(cur, total)
}

// 计算已下载字节数占总字节数的百分比
func (bar *Bar) getPercent() uint64 {
	return uint64(float32(bar.cur) / float32(bar.total) * 100)
}

// 该方法会循环执行，cur是当前已下载的字节数Bytes
func (bar *Bar) Plot(cur uint64) {
	bar.cur = cur
	last := bar.percent
	bar.percent = bar.getPercent()
	if bar.percent != last {
		// 这样写bar.percent/100*bar.totalSymbols 永远是0，因为 bar.percent永远小于100，整数相除，结果就是0
		bar.rate = strings.Repeat(bar.symbol, int(bar.percent*bar.totalSymbols/100))
	}
	fmt.Printf("\r[%-50s]%3d%% %8dBytes/%dBytes", bar.rate, bar.percent, bar.cur, bar.total)
}

func (bar *Bar) Finish() {
	fmt.Printf("\nDone!")
}

func main() {
	rand.Seed(time.Now().Unix())
	var cur uint64 = 788
	var total uint64 = 45677
	var bar Bar
	bar.NewBarWithSymbol(cur, total, "#")
	// 每隔100ms执行一次
	ch := time.Tick(100 * time.Millisecond)
	for {
		// 定时到了会解除阻塞
		<-ch
		if bar.percent == 100 {
			break
		}
		cur = cur + uint64(rand.Intn(500))
		bar.Plot(cur)
	}
	bar.Finish()
}
