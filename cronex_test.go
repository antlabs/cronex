package cronex

import (
	"fmt"
	"testing"
	"time"

	"github.com/antlabs/timer"
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
		close(durationChan)
	}()

	count = 0
	first := time.Duration(0)
	for tv := range durationChan {
		if first == 0 {
			first = tv
			continue
		}

		left := first + time.Duration(count)*time.Second
		right := first + time.Duration(1.2*float64(time.Duration(count)*time.Second))

		if tv < left || tv > right {
			t.Logf("count(%d), tv(%v), tv < left(%v) || tv > right(%v)", count, tv, left, right)
			return
		}
		count++
	}

	assert.NotEqual(t, count, 1, fmt.Sprintf("callback没有调用"))
}

func Test_Cronex_ParseNext(t *testing.T) {

	var schedule timer.Next
	schedule, err := standardParser.Parse("* * * * * *")
	assert.NoError(t, err)
	if err != nil {
		return
	}

	first := time.Duration(0)
	for count := 1; count < 4; count++ {
		now := schedule.Next(time.Now())
		left := first + time.Duration(0.8*float64(time.Second))
		right := first + time.Duration(1.2*float64(time.Second))

		tv := now.Sub(time.Now())
		if tv < left || tv > right {
			t.Logf("tv(%v), tv < left(%v) || tv > right(%v)", tv, left, right)
			return
		}
	}
}
