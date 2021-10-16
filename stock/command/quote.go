package command

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/wzhanjun/stock/utils/xueqiu"
)

var QuoteCommand = &cobra.Command{
	Use:   "quote [code]",
	Short: "股票行情",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return fmt.Errorf("请输入股票代码")
		}
		stockCode := strings.ToUpper(args[0])
		result, err := xueqiu.NewXQApiClient().GetQuote(stockCode)
		if err != nil {
			fmt.Println(err)
		}
		if result.Quote.Name == "" {
			return fmt.Errorf("查询股票信息失败，请确认code:%s是否有效", stockCode)
		}
		fmt.Printf(
			"查询成功: %s(%s), 当前价格: %.2f, 最低: %.2f, 最高: %.2f 跌涨幅：%.2f%%\n",
			result.Quote.Name,
			result.Quote.Symbol,
			result.Quote.Current,
			result.Quote.Low,
			result.Quote.High,
			result.Quote.Percent,
		)

		return nil
	},
}
