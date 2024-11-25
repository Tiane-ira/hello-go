package configs

const (
	// MinGoVersion 最小 Go 版本
	MinGoVersion = 1.18
	// ProjectVersion 项目版本
	SdkVersion = "v0.0.1"
	// ProjectName 项目名称
	ProjectName = "hello-go"
	// ProjectDomain 项目域名
	ProjectDomain = "http://127.0.0.1"
	// ProjectPort 项目端口
	ProjectPort = ":9999"
	// ProjectAccessLogFile 项目访问日志存放文件
	ProjectLogFile = "./logs/" + ProjectName + ".log"
	// ProjectCronLogFile 项目计划任务日志存放文件
	ProjectCronLogFile = "./logs/" + ProjectName + "-cron.log"
	// ConfigPath 配置文件目录
	ConfigPath = "./configs/"

	// ZhCN 简体中文 - 中国
	ZhCN = "zh-cn"
	// EnUS 英文 - 美国
	EnUS = "en-us"
)
