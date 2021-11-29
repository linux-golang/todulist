package fromt

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

func WriterTable(slice [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "任务名", "任务详情", "创建时间", "完成时间", "状态", "执行人"})
	table.SetRowLine(true)
	for _, v := range slice {
		table.Append(v)
	}
	table.Render()

}
