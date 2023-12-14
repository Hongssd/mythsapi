package indicatorsdef

import (
	"fmt"
	"strings"
)

type Indicator string

type Indicators []Indicator

func (is Indicator) String() string {
	return string(is)
}

func (is *Indicators) AddIndicators(newIndicators ...Indicator) {
	for _, indicator := range newIndicators {
		*is = append(*is, indicator)
	}
}

func (is *Indicators) UnmarshalJSON(data []byte) error {
	s := string(data)
	s = strings.Trim(s, "\"")
	if s == "" {
		return nil
	}
	indicators := strings.Split(s, ",")
	var result Indicators
	for _, indicator := range indicators {
		result = append(result, Indicator(indicator))
	}
	*is = result
	return nil
}

func (is *Indicators) MarshalJSON() ([]byte, error) {
	if is == nil || len(*is) == 0 {
		return []byte("\"\""), nil
	}
	var indicators []string
	for _, indicator := range *is {
		indicators = append(indicators, string(indicator))
	}
	result := strings.Join(indicators, ",")
	result = fmt.Sprintf("\"%s\"", result)
	return []byte(result), nil
}
