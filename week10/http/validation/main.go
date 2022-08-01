package main

import (
	"fmt"
	"github.com/go-playground/validator"
	"regexp"
)

var val = validator.New()

type RegistRequest struct {
	UserName   string `validate:"gt=0"`             // >0 长度大于0
	Password   string `validate:"min=6,max=12"`     //密码长度【6,12】
	PassRepeat string `validate:"eqfield=Password"` //跨字段相等校验
	Email      string `validate:"my_email"`            //需要满足email的格式
}

//自定义校验函数标签
func validateEmail(f1 validator.FieldLevel) bool {
	input := f1.Field().String()
	if pass, _ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,})\.([a-z]{2,4})$`, input); pass {
		return true
	}
	return false
}

type InnerRequest struct {
	Pass  string `validate:"min=6,max=12"`     //跨结构体相等校验
	Email string `validate:"eqfield=PassWord"` //跨字段相等校验
}

type OutterRequest struct {
	PassWord   string `validate:"eqcsfield=Nest.Pass"` //跨结构体相等校验
	PassRepeat string `validate:"eqfield=PassWord"`    //跨字段相等校验
	Nest       InnerRequest
}

//自定义错误输出
func processErr(err error) {
	if err == nil {
		return
	}

	//	给Validate.Struct()函数传了一个非法的参数
	invalid,ok := err.(*validator.InvalidValidationError)
	if ok{
		fmt.Println("param error:",invalid)
		return
	}

//	ValidationErrors 是一个错误切片，它保存了每个字段违反的每个约束信息
	validationErrs := err.(validator.ValidationErrors)
	for _,validationErr:=range validationErrs{
		fmt.Printf("field %s 不满足条件 %s\n",validationErr.Field(),validationErr.Tag())
	}
}

func main(){
	val.RegisterValidation("my_email",validateEmail) //注册一个自定义的validator

	req := RegistRequest{
		UserName: "jasmine",
		Password: "123456",
		PassRepeat: "123456",
		Email: "123qq",
	}
	fmt.Println(val.Struct(req))
	processErr(val.Struct(req)) //Struct()返回的error分为两种类型：InvalidValidationError和ValidationErrors
	processErr(val.Struct(3))
	fmt.Println("=========================")

	inreq := InnerRequest{
		Pass:"1234567",
		Email: "123qq.com",
	}

	outreq:=OutterRequest{
		PassWord: "123456",
		PassRepeat: "123456",
		Nest: inreq,
	}
	processErr(val.Struct(outreq))

}