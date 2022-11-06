// guonaihong apache 2.0

package cronex

import "github.com/antlabs/timer"

type TimerNoder = timer.TimeNoder

// cronex
type Cronex struct {
	tm timer.Timer
}

// 初始化一个cronex
func New() *Cronex {
	return &Cronex{
		tm: timer.NewTimer(timer.WithMinHeap()),
	}
}

// 添加函数
func (c *Cronex) AddFunc(spec string, cmd func()) (node TimerNoder, err error) {
	var schedule timer.Next
	schedule, err = standardParser.Parse(spec)
	if err != nil {
		return
	}

	return c.tm.CustomFunc(schedule, cmd), nil
}

// 运行消费者循环
func (c *Cronex) Run() {
	c.tm.Run()
}

// 异步运行消费者循环
func (c *Cronex) Start() {
	go c.Run()
}

// 关闭cronex的任务循环
func (c *Cronex) Stop() {
	c.tm.Stop()
}
