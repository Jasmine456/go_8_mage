package main

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"go_8_mage/old_project/cloud-station/cli"
	"os"
)

//防止密码被别人窥见到, 我们可以使用一个第三方库来加密我们的输入: https://github.com/AlecAivazis/survey
var AccessSecret string
func getSecretKeyFromInputV2() {
	prompt := &survey.Password{
		Message: "请输入access key: ",
	}
	survey.AskOne(prompt, &AccessSecret)
}

func main(){
	//getSecretKeyFromInputV2()
	if err:=cli.RootCmd.Execute();err!=nil{
		fmt.Println(err)
		os.Exit(-1)
	}
}