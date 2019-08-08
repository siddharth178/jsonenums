// generated by jsonenums -type=TestCasing -snake_case_json=true; DO NOT EDIT

package main

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

var (
	_TestCasingNameToValue = map[string]TestCasing{
		"case_madness_a":      caseMadnessA,
		"case_ma_dn_e_es_b":   caseMaDnEEsB,
		"normal_case_example": normalCaseExample,
	}

	_TestCasingValueToName = map[TestCasing]string{
		caseMadnessA:      "case_madness_a",
		caseMaDnEEsB:      "case_ma_dn_e_es_b",
		normalCaseExample: "normal_case_example",
	}
)

type _TestCasingInvalidValueError struct {
	invalidValue string
}

func (e _TestCasingInvalidValueError) Error() string {
	return fmt.Sprintf("invalid TestCasing: %s", e.invalidValue)
}

func (e _TestCasingInvalidValueError) InvalidValueError() string {
	return e.Error()
}

func init() {
	var v TestCasing
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_TestCasingNameToValue = map[string]TestCasing{
			interface{}(caseMadnessA).(fmt.Stringer).String():      caseMadnessA,
			interface{}(caseMaDnEEsB).(fmt.Stringer).String():      caseMaDnEEsB,
			interface{}(normalCaseExample).(fmt.Stringer).String(): normalCaseExample,
		}
	}
}

func ListTestCasingValues() map[string]string {
	TestCasingList := make(map[string]string)
	for k := range _TestCasingNameToValue {
		TestCasingList[k] = k
	}
	return TestCasingList
}

func (r TestCasing) toString() (string, error) {
	s, ok := _TestCasingValueToName[r]
	if !ok {
		return "", fmt.Errorf("invalid TestCasing: %d", r)
	}
	return s, nil
}

func (r TestCasing) getString() (string, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return s.String(), nil
	}
	return r.toString()
}

func (r *TestCasing) setValue(str string) error {
	v, ok := _TestCasingNameToValue[str]
	if !ok {
		return fmt.Errorf("invalid TestCasing %q", str)
	}
	*r = v
	return nil
}

// MarshalJSON is generated so TestCasing satisfies json.Marshaler.
func (r TestCasing) MarshalJSON() ([]byte, error) {
	s, err := r.getString()
	if err != nil {
		return nil, err
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so TestCasing satisfies json.Unmarshaler.
func (r *TestCasing) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("TestCasing should be a string, got %s", data)
	}
	return r.setValue(s)
}

//Scan an input string into this structure for use with GORP
func (r *TestCasing) Scan(i interface{}) error {
	switch t := i.(type) {
	case []byte:
		return r.setValue(string(t))
	case string:
		return r.setValue(t)
	default:
		return fmt.Errorf("can't scan %T into type %T", i, r)
	}
}

func (r TestCasing) Value() (driver.Value, error) {
	return r.getString()
}
