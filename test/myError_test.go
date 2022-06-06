package main

import (
	"context"
	"fmt"
	"time"
	"github.com/spf13/viper"
	"github.com/fsnotify/fsnotify"
)



var config *viper.Viper
func main() {
	v := viper.New()
	v.SetConfigName("config") // 设置文件名称（无后缀）
	v.SetConfigType("yaml")   // 设置后缀名 {"1.6以后的版本可以不设置该后缀"}
	v.AddConfigPath("./")  // 设置文件所在路径
	v.Set("verbose", true) // 设置默认参数

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic(" Config file not found; ignore error if desired")
		} else {
			panic("Config file was found but another error was produced")
		}
	}
	// 监控配置和重新获取配置
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})

	// 父context(利用根context得到)
	ctx, cancel := context.WithCancel(context.Background())

	// 父context的子协程
	go watch1(ctx)

	// 子context，注意：这里虽然也返回了cancel的函数对象，但是未使用
	valueCtx, cancelC := context.WithCancel(ctx)
	// 子context的子协程
	go watch2(valueCtx)

	fmt.Println("现在开始等待3秒,time=", time.Now().Unix())
	time.Sleep(3 * time.Second)
	cancelC()
	fmt.Println("等待2秒结束,调用cancel()函数")
	time.Sleep(2*time.Second)
	// 调用cancel()
	fmt.Println("等待3秒结束,调用cancel()函数")
	cancel()

	// 再等待5秒看输出，可以发现父context的子协程和子context的子协程都会被结束掉
	time.Sleep(5 * time.Second)
	fmt.Println("最终结束,time=", time.Now().Unix())
}

// 父context的协程
func watch1(ctx context.Context) {
	for {
		select {
		case <-ctx.Done(): //取出值即说明是结束信号
			fmt.Println("收到信号，父context的协程退出,time=", time.Now().Unix())
			return
		default:
			fmt.Println("父context的协程监控中,time=", time.Now().Unix())
			time.Sleep(1 * time.Second)
		}
	}
}

// 子context的协程
func watch2(ctx context.Context) {
	for {
		select {
		case <-ctx.Done(): //取出值即说明是结束信号
			fmt.Println("收到信号，子context的协程退出,time=", time.Now().Unix())
			return
		default:
			fmt.Println("子context的协程监控中,time=", time.Now().Unix())
			time.Sleep(1 * time.Second)
		}
	}
}