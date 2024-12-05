package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/nabinyd/backend-15_days_challenge.git/util"
)

var validCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		// check currency is supported
		return util.IsSupportedCurrency(currency)

	}
	return false

}
