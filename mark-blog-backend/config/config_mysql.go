package config

import "strconv"

type Mysql struct{
	Host 		string  `yaml:"host"`
	Port 		int		`yaml:"port"`
	Config      string  `yaml:"config"`   //高级配置，例如charset
	Db 	 		string  `yaml:"db"`
	User		string 	`yaml:"user"`
	Password	string 	`yaml:"password"`
	LogLevel	string 	`yaml:"log_level"`
}	


// Mysqldb = "root:root@tcp(127.0.0.1:3306)/markBlog?charset=utf8"

func (m Mysql)Dsn() string{
	return m.User + ":" + m.Password + "@tcp(" +m.Host+":"+strconv.Itoa(m.Port) + ")/" + m.Db + m.Config;
}