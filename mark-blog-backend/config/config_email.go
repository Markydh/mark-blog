package config

type Email struct {
	// 邮箱账号
	Account string 		`yaml:"account"`
	// QQ邮箱填写授权码
	Password string 	`yaml:"password"`
	// QQ：POP/SMTP 587
	Port int 			`yaml:"port"`
	// QQ：smtp.qq.com
	Host string			`yaml:"host"`
}
