package app

import (
	"fmt"
	"regexp"
)

// Validate 正则表达式定义
var Validate = map[string]*regexp.Regexp{
	"email":    regexp.MustCompile(`^\w+@\w+\.\w+$`),
	"username": regexp.MustCompile(`([a-z]|[A-z])+[0-9]*`),
	"num":      regexp.MustCompile(`^[1-9]+[0-9]*$`),
	"zoronum":  regexp.MustCompile(`^[0-9]+$`),
	"mobile":   regexp.MustCompile(`1[0-9]{10}`),
	"password": regexp.MustCompile(`^\w{6,20}$`),
	"notempty": regexp.MustCompile(`\S`),
}

//ValidateForm 批量定义表单验证
type ValidateForm struct {
	Key   string
	Field string
	Msg   string
}

// Check 验证字段是否符合要求
func Check(reg string, field string) error {
	if ok := Validate[reg].MatchString(field); !ok {
		return fmt.Errorf("validate: %s is not match", field)
	}
	return nil
}

//AutoCheck 批量验证
func AutoCheck(vfs []ValidateForm) error {
	for _, vf := range vfs {
		if err := Check(vf.Key, vf.Field); err != nil {
			return fmt.Errorf("validate: %s", vf.Msg)
		}
	}
	return nil
}
