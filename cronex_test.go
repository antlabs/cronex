package cronex

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_Cronex(t *testing.T) {
	table := []string{
		"*/1 * * * * *", //每秒执行一次
	}

	cron := New()
	cron.Start() //运行事件循环

	count := 2
	durationChan := make(chan time.Duration, count)
	now := time.Now()
	var tm TimerNoder
	var err error

	for _, tc := range table {
		tm, err = cron.AddFunc(tc, func() {
			durationChan <- time.Since(now)
		})
		assert.NoError(t, err)
	}

	// 3s之后关闭
	go func() {
		time.Sleep(time.Second * 3)
		tm.Stop()
		cron.Stop()
	}()

	count = 1
	for tv := range durationChan {
		fmt.Printf("%v\n", tv)
		left := time.Duration(count) * time.Second
		right := 2 * time.Duration(count) * time.Second

		if tv < left || tv > right {
			t.Logf("tv(%v), tv < left(%v) || tv > right(%v)", tv, left, right)
			return
		}
		count++
	}

	assert.NotEqual(t, count, 1, fmt.Sprintf("callback没有调用"))
}

func Test_ParseNext(t *testing.T) {

}
