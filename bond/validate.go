package bond

import (
	"fmt"
)

type ValidatorError struct {
	field  string
	reason string
}

func (error ValidatorError) Error() string {
	return fmt.Sprintf("数据校验失败：字段%s,reason:%s", error.field, error.reason)
}

func NewValidatorError(field, reason string) error {
	return &ValidatorError{field, reason}
}

type Validator interface {
	Validate(param map[string]string) ValidatorError
}
