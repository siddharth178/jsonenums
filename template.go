// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Added as a .go file to avoid embedding issues of the template.

package main

import "text/template"

var generatedTmpl = template.Must(template.New("generated").Parse(`
// generated by jsonenums {{.Command}}; DO NOT EDIT

package {{.PackageName}}

import (
    "database/sql/driver"
    "encoding/json"
    "fmt"
)

{{range $typename, $values := .TypesAndValues}}

var (
    _{{$typename}}NameToValue = map[string]{{$typename}} {
        {{range $values}}"{{.SnakeRep}}": {{.CammelRep}},
        {{end}}
    }

    _{{$typename}}ValueToName = map[{{$typename}}]string {
        {{range $values}}{{.CammelRep}}: "{{.SnakeRep}}",
        {{end}}
    }
)

func init() {
    var v {{$typename}}
    if _, ok := interface{}(v).(fmt.Stringer); ok {
        _{{$typename}}NameToValue = map[string]{{$typename}} {
            {{range $values}}interface{}({{.CammelRep}}).(fmt.Stringer).String(): {{.CammelRep}},
            {{end}}
        }
    }
}

// MarshalJSON is generated so {{$typename}} satisfies json.Marshaler.
func (r {{$typename}}) MarshalJSON() ([]byte, error) {
    if s, ok := interface{}(r).(fmt.Stringer); ok {
        return json.Marshal(s.String())
    }
    s, ok := _{{$typename}}ValueToName[r]
    if !ok {
        return nil, fmt.Errorf("invalid {{$typename}}: %d", r)
    }
    return json.Marshal(s)
}

// UnmarshalJSON is generated so {{$typename}} satisfies json.Unmarshaler.
func (r *{{$typename}}) UnmarshalJSON(data []byte) error {
    var s string
    if err := json.Unmarshal(data, &s); err != nil {
        return fmt.Errorf("{{$typename}} should be a string, got %s", data)
    }
    v, ok := _{{$typename}}NameToValue[s]
    if !ok {
        return fmt.Errorf("invalid {{$typename}} %q", s)
    }
    *r = v
    return nil
}

//Scan an input string into this structure for use with GORP
func (r *{{$typename}}) Scan(i interface{}) error {
	switch t := i.(type) {
	case string:
		r.UnmarshalJSON([]byte(t))
	default:
		return fmt.Errorf("Can't scan %T into type %T", i, r)
	}
	return nil
}

func (r {{$typename}}) Value() (driver.Value, error) {
	bytes, err := r.MarshalJSON()
	return string(bytes), err 
}

{{end}}
`))
