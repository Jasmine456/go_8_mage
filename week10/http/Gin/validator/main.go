package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"//注意必须 要用新版本v10
	"net/http"
	"time"
)

type Student struct {
	Name       string    `form:"name" binding:"required"`
	Score      int       `form:"score" binding:"gt=0"`
	Enrollment time.Time `form:"enrollment" binding:"required,before_today" time_format:"2006-01-02" time_utc:"8"`
	Graduation time.Time `form:"graduation" binding:"required,gtfield=Enrollment" time_format:"2006-01-02" time_utc:"8"`
}

//自定义验证器
var beforeToday validator.Func = func(f1 validator.FieldLevel) bool {
	if date, ok := f1.Field().Interface().(time.Time); ok {
		today := time.Now()
		if date.Before(today) {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

func processErr(err error) {
	if err == nil {
		return
	}

	//	给validate.Struct()函数传了一个非法参数
	invalid, ok := err.(*validator.InvalidValidationError)
	if ok {
		fmt.Println("param error:", invalid)
		return
	}

	//	ValidationErrors是一个错误的切片，他保存了每个字段违反的每个约束信息
	validationErrs := err.(validator.ValidationErrors)
	for _, validationErr := range validationErrs {
		fmt.Printf("field %s 不满足条件%s\n", validationErr.Field(), validationErr.Tag())
	}
}

func main() {
	engine := gin.Default()

	//	注册验证器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("before_today", beforeToday)
	}

	engine.GET("/", func(ctx *gin.Context) {
		var stu Student
		if err := ctx.ShouldBind(&stu); err != nil {
			processErr(err) //校验不符合时，打印出那时不符合
			ctx.String(http.StatusBadRequest, "parse parameter failed")
		} else {
			ctx.JSON(http.StatusOK, stu)
		}
	})

	engine.Run(":5656")
}

//http://localhost:5656?name=zcy&score=1&enrollment=2021-08-23&graduation=2021-09-23
