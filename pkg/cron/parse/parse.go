package parse

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/perithompson/cron-parser/pkg/cron"
)

const (
	MinutesMin         = 0
	MinutesMax         = 59
	HourMin            = 0
	HourMax            = 23
	DayOfMonthMin      = 1
	DayOfMonthMax      = 31
	MonthMin           = 1
	MonthMax           = 12
	DayOfWeekMin       = 0
	DayOfWeekMax       = 6
	MinutePosition     = 0
	HourPosition       = 1
	DayOfMonthPosition = 2
	MonthPosition      = 3
	DayOfWeekPosition  = 4
	CommandPosition    = 5
)

//Parse returns a Cron Object from an expression string
func Parse(cronStr string) (*cron.Cron, error) {
	c := new(cron.Cron)
	var err error

	args := strings.Split(cronStr, " ")
	if len(args) < 6 {
		return c, fmt.Errorf("invalid input, expected inputs: %d, got: %d", 6, len(args))
	}
	c.Minute, err = ParseArg(args[MinutePosition], MinutesMin, MinutesMax)
	if err != nil {
		return c, err
	}
	c.Hour, err = ParseArg(args[HourPosition], HourMin, HourMax)
	if err != nil {
		return c, err
	}
	c.DayOfMonth, err = ParseArg(args[DayOfMonthPosition], DayOfMonthMin, DayOfMonthMax)
	if err != nil {
		return c, err
	}
	args[MonthPosition] = ParseShortMonthExpr(args[MonthPosition])
	c.Month, err = ParseArg(args[MonthPosition], MonthMin, MonthMax)
	if err != nil {
		return c, err
	}
	args[DayOfWeekPosition] = ParseDayExpr(args[DayOfWeekPosition])
	c.DayOfWeek, err = ParseArg(args[DayOfWeekPosition], DayOfWeekMin, DayOfWeekMax)
	if err != nil {
		return c, err
	}
	if args[CommandPosition] == "" {
		return c, fmt.Errorf("missing command in cron spec")
	}
	c.Command = args[CommandPosition]
	return c, nil
}

//ParseArg returns an array of int from a given expression string
func ParseArg(expr string, min, max int) ([]int, error) {
	var err error
	var values []int
	switch {
	case strings.Split(expr, "")[0] == "*":
		values, err = EveryOrRange(expr, min, max)
	case strings.Contains(expr, "-"):
		values, err = EveryOrRange(expr, min, max)
	case strings.Contains(expr, ","):
		values, err = ListInt(expr)
	default:
		values, err = IntVal(expr, min, max)
	}
	if duplicateInArray(values) != -1 {
		return values, fmt.Errorf("duplicate values found in argument: %s --- duplicate found %d", expr, duplicateInArray(values))
	}
	return values, err
}

//EveryOrRange returns a slice of integer and validates if all in correct range
func EveryOrRange(expr string, min, max int) ([]int, error) {
	var values []int
	var err error
	offset := 1
	if strings.Contains(expr, "/") {
		skip, err := strconv.Atoi(strings.SplitAfter(expr, "/")[1])
		if err != nil {
			return []int{}, err
		}
		offset = skip
		expr = strings.Trim(strings.SplitAfter(expr, "/")[0], "/")
	}
	i := min
	if strings.Contains(expr, "-") {
		limits := strings.SplitAfter(expr, "-")
		i, err = strconv.Atoi(strings.TrimRight(limits[0], "-"))
		if err != nil {
			return values, err
		}
		upper, err := strconv.Atoi(limits[1])
		if err != nil {
			return values, err
		}
		if i < min {
			return values, fmt.Errorf("lower bound in expression is lower that expected minimum: %s, please set lower bound greated than %d", expr, min)
		}
		if upper > max {
			return values, fmt.Errorf("upper bound in expression is greater that expected maximum: %s, please set upped bound less than %d", expr, max)
		}
		max = upper

	}
	for i <= max {
		values = append(values, i)
		i += offset
	}
	return values, nil
}

//ListInt returns a slice of integer and validates if all in correct range
func ListInt(expr string) ([]int, error) {
	var values []int

	for _, v := range strings.Split(expr, ",") {
		i, err := strconv.Atoi(v)
		if err != nil {
			return []int{}, err
		}
		values = append(values, i)
	}
	return values, nil
}

//IntVal returns a single integer and validates if in correct range
func IntVal(expr string, min, max int) ([]int, error) {
	var values []int

	val, err := strconv.Atoi(expr)
	if err != nil {
		return values, err
	}
	if val < min {
		return values, fmt.Errorf("value given is lower that expected minimum: %s", expr)
	}
	if val > max {
		return values, fmt.Errorf("value given is greater that expected maximum: %s", expr)
	}
	values = append(values, val)
	return values, nil
}

//ParseDayExpr returns a slice of integer converting days to corresponding integers and validates if in correct range
func ParseDayExpr(expr string) string {
	days := map[string]string{
		"sunday":    "0",
		"monday":    "1",
		"tuesday":   "2",
		"wednesday": "3",
		"thursday":  "4",
		"friday":    "5",
		"saturday":  "6",
	}
	for str, index := range days {
		expr = strings.Replace(strings.ToLower(expr), str, index, -1)
	}

	//in case short days are used run through another check
	return ParseShortDayExpr(expr)
}

//ParseShortDayExpr returns a slice of integer converting first three chars of
//days to corresponding integers and validates if in correct range
func ParseShortDayExpr(expr string) string {
	days := map[string]string{
		"sun": "0",
		"mon": "1",
		"tue": "2",
		"wed": "3",
		"thu": "4",
		"fri": "5",
		"sat": "6",
	}
	for str, index := range days {
		expr = strings.Replace(strings.ToLower(expr), str, index, -1)
	}
	return expr
}

//ParseShortMonthExpr returns a slice of integer converting first three chars of
//month to corresponding integers and validates if in correct range
func ParseShortMonthExpr(expr string) string {
	days := map[string]string{
		"jan": "1",
		"feb": "2",
		"mar": "3",
		"apr": "4",
		"may": "5",
		"jun": "6",
		"jul": "7",
		"aug": "8",
		"sep": "9",
		"oct": "10",
		"nov": "11",
		"dec": "12",
	}
	for str, index := range days {
		expr = strings.Replace(strings.ToLower(expr), str, index, -1)
	}
	return expr
}

//duplicateInArray checks for duplicates
func duplicateInArray(arr []int) int {
	seen := make(map[int]bool, 0)
	for i := 0; i < len(arr); i++ {
		if seen[arr[i]] == true {
			return arr[i]
		} else {
			seen[arr[i]] = true
		}
	}
	return -1
}
