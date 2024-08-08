package validation

import (
	"forth/utils/exception"
	"net/http"
	"strconv"
	"strings"
)

func FormValidate(r *http.Request, validation map[string]string) map[string][]string {
	var errors = make(map[string][]string)

	for k, v := range validation {
		itemValidations := strings.Split(v, "|")
		for _, val := range itemValidations {
			if val == "required" {
				if r.Form.Get(k) == "" {
					_, hasError := errors[k]
					if hasError {
						errors[k] = append(errors[k], k+" cannot be null")
					} else {
						errors[k] = []string{k + " cannot be null"}
					}

				}
			}
			if strings.Contains(val, "min") {

				minLength, err := strconv.Atoi(strings.Split(val, ":")[1])
				if err != nil {
					exception.Exception("bad validation", err.Error())
				}

				if len(r.Form.Get(k)) < minLength {
					_, hasError := errors[k]
					if hasError {
						errors[k] = append(errors[k], k+" must be at least "+strconv.Itoa(minLength)+" characters")
					} else {
						errors[k] = []string{k + " must be at least " + strconv.Itoa(minLength) + " characters"}
					}
				}
			}
			if strings.Contains(val, "max") {

				maxLength, err := strconv.Atoi(strings.Split(val, ":")[1])
				if err != nil {
					exception.Exception("bad validation", err.Error())
				}

				if len(r.Form.Get(k)) > maxLength {
					_, hasError := errors[k]
					if hasError {
						errors[k] = append(errors[k], k+" must not be greater than "+strconv.Itoa(maxLength)+" characters")
					} else {
						errors[k] = []string{k + " must not be greater than " + strconv.Itoa(maxLength) + " characters"}
					}
				}
			}
		}
	}
	return errors

}
