package env

import (
	"flag"
	"fmt"
	"strings"
)

var active string

func init() {
	env := flag.String("e", "", "运行环境:\n dev:开发环境\n fat:测试环境\n uat:预上线环境\n pro:正式环境\n")
	flag.Parse()

	switch strings.ToLower(strings.TrimSpace(*env)) {
	case "prod":
		active = "prod"
	case "dev":
		active = "dev"
	default:
		active = "dev"
		fmt.Println("Warning: '-e' is invalid. The default 'dev' will be used.")
	}
}

func Active() string {
	return active
}
