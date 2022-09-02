package conf


import(
	"github.com/caarlos0/env/v6"
	"github.com/BurntSushi/toml"
)

var (
//	为了保护该变量不被外部修改，未暴露
	config *Config
)

//单独提供一个方式Getter
func C() *Config{
	if config == nil{
		panic( "load config first")
	}
	return config
}


//从环境变量中加载配置,加载成一个全局变量
func  LoadConfigFromEnv() error{
	config=NewDefaultConfig()
	//if err:= env.Parse(config);err!=nil{
	//	return err
	//}
	//return nil
	return env.Parse(config)
}

func LoadConfigFromToml(filepath string)error{
	config=NewDefaultConfig()
	_,err:=toml.DecodeFile(filepath,config)
	if err != nil {
		return err
	}
	return nil
}