package utils

import "regexp"

type ZipCode string

func (z *ZipCode) IsValid() bool {
	re := regexp.MustCompile(`\D`)

	cleaned := re.ReplaceAllString(string(*z), "")

	return len(cleaned) == 8
}

func (z *ZipCode) Masked() string {
	// Remove any non-digit characters before masking
	re := regexp.MustCompile(`\D`)
	cleaned := re.ReplaceAllString(string(*z), "")

	if len(cleaned) != 8 {
		return string(*z)
	}

	return cleaned[:5] + "-" + cleaned[5:]
}

func (z *ZipCode) Raw() string {
	re := regexp.MustCompile(`\D`)
	return re.ReplaceAllString(string(*z), "")
}
