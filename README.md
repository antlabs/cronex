# cronex
高性能cron库，相比目前使用得最多的cron，只是优化了性能。

# 特性
* 继承robfig/cron全部的解析器代码
* 优化调度相关性能

# demo
```go
import(
    "github.com/antlabs/cronex"
)

func main() {
    cron := cronex.New()
    cron.AddFunc("* * * * * *", func() {

    })
    cron.Run() //阻塞，如果要异步就用cron.Start()
}
```
