package printer

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/perithompson/cron-parser/pkg/cron"
)

func intarrayToString(a []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
}

func Print(c *cron.Cron) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 0, ' ', tabwriter.DiscardEmptyColumns)
	fmt.Fprintln(w, "minute", "\t", intarrayToString(c.Minute, " "))
	fmt.Fprintln(w, "hour", "\t", intarrayToString(c.Hour, " "))
	fmt.Fprintln(w, "day of month", "\t", intarrayToString(c.DayOfMonth, " "))
	fmt.Fprintln(w, "month", "\t", intarrayToString(c.Month, " "))
	fmt.Fprintln(w, "day of week", "\t", intarrayToString(c.DayOfWeek, " "))
	fmt.Fprintln(w, "command", "\t", c.Command)
	w.Flush()
}
