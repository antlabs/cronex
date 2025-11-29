# cronex
![Go CI](https://github.com/antlabs/cronex/actions/workflows/ci.yml/badge.svg?branch=master)

高性能cron库，相比目前使用得最多的cron，只是优化了性能。

# 特性
* 继承robfig/cron全部的解析器代码
* 优化调度相关性能

# cpu占用对比(越低越好)
![cronex.png](https://github.com/guonaihong/images/blob/master/cronex/cronex.png)    
测试代码位置 https://github.com/guonaihong/crontest
# 快速开始
```go
import(
    "github.com/antlabs/cronex"
)

func main() {
    cron := cronex.New()
    cron.AddFunc("* * * * * *", func() {
        //TODO
    })
    cron.Run() //开启阻塞消费者循环，如果要异步就用cron.Start()
}
```

# 关闭任务
```go
import(
    "github.com/antlabs/cronex"
)

func main() {
    cron := cronex.New()
    tm, err := cron.AddFunc("* * * * * *", func() {
        //TODO
    })
    if err != nil {
        return
    }
    tm.Stop()  //删除这个任务
    cron.Run() //开启阻塞消费者循环，如果要异步就用cron.Start()
}
```
