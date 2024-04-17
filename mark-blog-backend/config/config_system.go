package config

type System struct{
	Host		string 			`yaml:"host"`
	Port 		uint			`yaml:"port"`
	Env		string			`yaml:"env"`
  }