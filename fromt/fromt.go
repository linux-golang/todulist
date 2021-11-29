package fromt

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

func WriterTable(slice [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Sing", "Rating"})
	for _, v := range slice {
		table.Append(v)
	}
	table.Render()

}
