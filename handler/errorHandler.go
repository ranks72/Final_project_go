package handler

import "github.com/go-playground/validator/v10"

type ErrorMsg struct {
	Error string `json:"message"`
}

func getErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return fe.Field() + " is required"
	case "min":
		if fe.Field() == "age" {
			return fe.Field() + " Umur kurang dari 8 tahun"
		}
		return fe.Field() + " harus memiliki karakter min. " + fe.Param() + " karakter"
	}
	return fe.Field() + " is invalid"
}
