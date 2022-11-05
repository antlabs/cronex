// guonaihong apache 2.0

package cronex

import "github.com/antlabs/timer"

type Cronex struct {
	tm timer.Timer
}

func New() *Cronex {
	return &Cronex{
		tm: timer.NewTimer(timer.WithMinHeap()),
	}
}

func (c *Cronex) AddFunc(spec string, cmd func()) (node timer.TimeNoder, err error) {
	return
}

func (c *Cronex) Run() {

}

func (c *Cronex) Start() {
	go c.Run()
}
