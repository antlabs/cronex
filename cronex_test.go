package cronex

import (
	"fmt"
	"sync/atomic"
	"testing"
	"time"

	"github.com/antlabs/timer"
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
		if err != nil {
			t.Logf("err(%v)", err)
			return
		}
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

	if count != 1 {
		t.Logf("count(%d), count != 1, callback 没有调用", count)
		return
	}
}

// 测试下Next函数的时间可正确
func Test_Cronex_ParseNext(t *testing.T) {

	var schedule timer.Next
	schedule, err := standardParser.Parse("* * * * * *")
	if err != nil {
		t.Logf("err(%v)", err)
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

// 多次运行的例子
func Test_Multiple(t *testing.T) {
	cron := New()
	count := int32(0)

	cron.Start()
	max := int32(10)
	for i := int32(0); i < max; i++ {
		cron.AddFunc("* * * * * *", func() {
			fmt.Printf("Every Second")
			atomic.AddInt32(&count, 1)
		})
	}

	time.Sleep(time.Duration(1.1 * float64(time.Second)))
	cron.Stop()
	if count != max {
		t.Errorf("expected %d, got %d", max, count)
	}
}
