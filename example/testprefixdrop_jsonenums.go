// generated by jsonenums -type=TestPrefixDrop -prefix_to_drop=prefixDrop; DO NOT EDIT

package main

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

var (
	_TestPrefixDropNameToValue = map[string]TestPrefixDrop{
		"ThisS": prefixDropThisS,
		"hould": prefixDrophould,
		"Work":  prefixDropWork,
	}

	_TestPrefixDropValueToName = map[TestPrefixDrop]string{
		prefixDropThisS: "ThisS",
		prefixDrophould: "hould",
		prefixDropWork:  "Work",
	}
)

type _TestPrefixDropInvalidValueError struct {
	invalidValue string
}

func (e _TestPrefixDropInvalidValueError) Error() string {
	return fmt.Sprintf("invalid TestPrefixDrop: %s", e.invalidValue)
}

func (e _TestPrefixDropInvalidValueError) InvalidValueError() string {
	return e.Error()
}

func init() {
	var v TestPrefixDrop
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_TestPrefixDropNameToValue = map[string]TestPrefixDrop{
			interface{}(prefixDropThisS).(fmt.Stringer).String(): prefixDropThisS,
			interface{}(prefixDrophould).(fmt.Stringer).String(): prefixDrophould,
			interface{}(prefixDropWork).(fmt.Stringer).String():  prefixDropWork,
		}
	}
}

func ListTestPrefixDropValues() map[string]string {
	TestPrefixDropList := make(map[string]string)
	for k := range _TestPrefixDropNameToValue {
		TestPrefixDropList[k] = k
	}
	return TestPrefixDropList
}

func (r TestPrefixDrop) toString() (string, error) {
	s, ok := _TestPrefixDropValueToName[r]
	if !ok {
		return "", fmt.Errorf("invalid TestPrefixDrop: %d", r)
	}
	return s, nil
}

func (r TestPrefixDrop) getString() (string, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return s.String(), nil
	}
	return r.toString()
}

func (r *TestPrefixDrop) setValue(str string) error {
	v, ok := _TestPrefixDropNameToValue[str]
	if !ok {
		return fmt.Errorf("invalid TestPrefixDrop %q", str)
	}
	*r = v
	return nil
}

// MarshalJSON is generated so TestPrefixDrop satisfies json.Marshaler.
func (r TestPrefixDrop) MarshalJSON() ([]byte, error) {
	s, err := r.getString()
	if err != nil {
		return nil, err
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so TestPrefixDrop satisfies json.Unmarshaler.
func (r *TestPrefixDrop) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("TestPrefixDrop should be a string, got %s", data)
	}
	return r.setValue(s)
}

//Scan an input string into this structure for use with GORP
func (r *TestPrefixDrop) Scan(i interface{}) error {
	switch t := i.(type) {
	case []byte:
		return r.setValue(string(t))
	case string:
		return r.setValue(t)
	default:
		return fmt.Errorf("can't scan %T into type %T", i, r)
	}
}

func (r TestPrefixDrop) Value() (driver.Value, error) {
	return r.getString()
}
