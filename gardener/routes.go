package gardener

import (
	"os"

	"github.com/ihorbryk/kalyna/router"
	"github.com/olekukonko/tablewriter"
)

func PrintRoutes(routes router.Routes) {
	table := tablewriter.NewWriter(os.Stdout)

	table.Header([]string{"Path", "Name"})

	for _, route := range routes {
		table.Append([]string{route.Path, route.Name})
	}

	table.Render()

}
