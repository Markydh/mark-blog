package config


//记录配置结构体
type Config struct{
	Mysql Mysql `yaml:"mysql"`
	Logger Logger `yaml:"logger"`
	System System `yaml:"system"`
}





