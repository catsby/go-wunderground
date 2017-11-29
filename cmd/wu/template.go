package main

import (
	"path"
	"strconv"
	"strings"
	"text/template"
	"time"
)

type UnitSelect string

const (
	// Use customary units
	UCustomary = UnitSelect("us")
	// Use SI units
	UMetric = UnitSelect("si")
)

// Var wraps an arbitrary value, allowing the user to effectively modify variables in templates.
type Var struct{ v interface{} }

func NewVar(val interface{}) *Var { return &Var{val} }
func (v *Var) Set(val interface{}) *Var { v.v = val; return v }
func (v *Var) Get() interface{}        { return v.v }

func TemplateEngine(units UnitSelect) (*template.Template, error) {
	funcs := template.FuncMap{
		"atoi": func (s string) int { i, _ := strconv.Atoi(s); return i },
		"atof": func (s string) float64 { f, _ := strconv.ParseFloat(s, 64); return f },
		"var": NewVar,
		"join": strings.Join,
		"timef": time.Time.Format,
	}
	te, err := template.New("_").Funcs(funcs).ParseGlob(
		path.Join("templates", "*.tmpl"))
	if err != nil {
		return nil, err
	}
	return te.ParseGlob(
		path.Join("templates", string(units), "*.tmpl"))
}
