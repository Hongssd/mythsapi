package codesdef

import (
	"fmt"
	"strings"
)

type Code string

type Codes []Code

func (cs *Codes) AddCodes(newCodes ...Code) {
	for _, code := range newCodes {
		*cs = append(*cs, code)
	}
}

func (c *Codes) UnmarshalJSON(data []byte) error {
	s := string(data)
	s = strings.Trim(s, "\"")
	if s == "" {
		return nil
	}
	codes := strings.Split(s, ",")
	var result Codes
	for _, code := range codes {
		result = append(result, Code(code))
	}
	*c = result
	return nil
}


func (c *Codes) MarshalJSON() ([]byte, error) {
	if c == nil || len(*c) == 0 {
		return []byte("\"\""), nil
	}
	var codes []string
	for _, code := range *c {
		codes = append(codes, string(code))
	}
	result := strings.Join(codes, ",")
	result = fmt.Sprintf("\"%s\"", result)
	return []byte(result), nil
}
