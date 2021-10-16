package command

import (
	"fmt"
	"strings"
	"time"

	"github.com/robfig/cron/v3"
	"github.com/wzhanjun/stock/config"
	"github.com/wzhanjun/stock/utils/xueqiu"
)

func StartJobs() {

	cron := cron.New(cron.WithSeconds(), cron.WithLocation(time.Local))

	// 行情任务
	cron.AddFunc("*/10 * * * * *", QuoteJob)

	cron.Start()

	select {}
}

func QuoteJob() {
	// 获取配置项
	cfg, err := config.NewConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	codeStr := cfg.StockConfig.Codes
	if codeStr == "" {
		fmt.Println("获取股票codes为空")
		return
	}

	codes := strings.Split(codeStr, ",")

	fmt.Printf("\n===============Start: %s==================\n\n", time.Now().Format("2006-01-02 15:03:04"))

	for _, stockCode := range codes {
		if stockCode != "" {
			result, err := xueqiu.NewXQApiClient().GetQuote(stockCode)
			if err != nil {
				fmt.Println(err)
				continue
			}
			if result.Quote.Name == "" {
				fmt.Printf("%c[1;0;31m[查询失败]%c[0m，请确认股票代码:%s是否有效\n", 0x1B, 0x1B, stockCode)
				continue
			}
			fmt.Printf(
				"%c[1;0;32m[查询成功]%c[0m: %s(%s), 当前价格: %.2f, 最低: %.2f, 最高: %.2f 跌涨幅：%.2f%%\n",
				0x1B,
				0x1B,
				result.Quote.Name,
				result.Quote.Symbol,
				result.Quote.Current,
				result.Quote.Low,
				result.Quote.High,
				result.Quote.Percent,
			)
		}
	}
}
