package dba

import (
	"fmt"
	"reflect"

	"github.com/hashicorp/hil"
	"github.com/hashicorp/hil/ast"
)

// TemplateError denotes failing to table update template.
type TemplateError struct {
	err     error
	tblName string
	field   string
	input   interface{}
}

// Error returns the formatted template error.
func (te TemplateError) Error() string {
	return fmt.Sprintf("While parsing template for %s, %s in %s table : %s",
		te.field, te.input, te.tblName, te.err.Error())
}

func newTemplateContext() EvalConfig {
	return EvalConfig{
		&hil.EvalConfig{
			GlobalScope: &ast.BasicScope{
				VarMap: map[string]ast.Variable{
					"hello": {
						Type:  ast.TypeString,
						Value: "Hello World!",
					},
				},
				FuncMap: map[string]ast.Function{
					"sentences_n":     sentencesN,
					"words_n":         wordsN,
					"year":            year,
					"characters_n":    charactersN,
					"digits_n":        digitsN,
					"paragraphs_n":    paragraphsN,
					"credit_card_num": creditCardNum,
					"password":        password,
				},
			},
		},
	}
}

func buildTemplateContext() EvalConfig {
	EvalConfig := newTemplateContext()
	scope := EvalConfig.GlobalScope

	for k := range funcMap {
		// create a closure over the actual function call
		fn := func(k string) func([]interface{}) (interface{}, error) {
			return func(inputs []interface{}) (interface{}, error) {
				val, err := funcCall(funcMap, k)
				result := val[0].Interface().(string)
				return result, err
			}
		}(k)

		// build func AST
		funcAst := ast.Function{
			ArgTypes:   []ast.Type{},
			ReturnType: ast.TypeString,
			Variadic:   false,
			Callback:   fn,
		}

		scope.FuncMap[k] = funcAst
	}

	return EvalConfig
}

func funcCall(m map[string]interface{}, name string, params ...interface{}) ([]reflect.Value, error) {
	f := reflect.ValueOf(m[name])
	// TODO: get args number from len(ArgTypes) from scope.FuncMap
	funcArgNum := f.Type().NumIn()
	if len(params) != funcArgNum {
		err := fmt.Errorf("Function %s expected %d arguments (%d given)",
			name, funcArgNum, len(params))
		return nil, TemplateError{err: err}
	}

	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	result := f.Call(in)
	return result, nil
}

// EvalTemplate evaluates template if string, otherwise returns value as given
func EvalTemplate(v interface{}, Ctx EvalConfig) (interface{}, error) {
	switch v := v.(type) {
	case string:
		return eval(v, Ctx)
	default:
		return v, nil
	}
}

func eval(tmpl string, Ctx EvalConfig) (interface{}, error) {
	tree, err := hil.Parse(tmpl)
	if err != nil {
		return nil, err
	}

	result, err := hil.Eval(tree, Ctx.EvalConfig)
	if err != nil {
		return nil, err
	}
	return result.Value, nil
}
