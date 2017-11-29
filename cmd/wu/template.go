//go:generate go-bindata -nomemcopy templates/...
package main

import (
	"fmt"
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

	te := template.New("_").Funcs(funcs)

	loadAssets := func(dir string, assets []string) (*template.Template, error) {
		for _, name := range assets {
			if !strings.HasSuffix(name, ".tmpl") {
				continue
			}
			name = path.Join(dir, name)
			asset, err := Asset(name)
			if err != nil {
				return nil, fmt.Errorf("unable to load asset %s: %s", name, err)
			}

			te, err = te.Parse(string(asset))
			if err != nil {
				return nil, fmt.Errorf("unable to parse template %s: %s", name, err)
			}
		}

		return te, nil
	}

	dir := "templates"
	assets, err := AssetDir(dir)
	if err != nil {
		return nil, fmt.Errorf("failed to find top-level template assets: %s", err)
	}

	te, err = loadAssets(dir, assets)
	if err != nil {
		return nil, fmt.Errorf("failed to load top-level template assets: %s", err)
	}

	dir = path.Join(dir, string(units))
	assets, err = AssetDir(dir)
	if err != nil {
		return nil, fmt.Errorf("failed to find template assets: %s", err)
	}

	te, err = loadAssets(dir, assets)
	if err != nil {
		return nil, fmt.Errorf("failed to load template assets: %s", err)
	}

	return te, nil
}
