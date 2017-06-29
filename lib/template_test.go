package dba

import "testing"

func TestEvalFail(t *testing.T) {

	Ctx := newTemplateContext()
	templates := []string{
		"${first_name()}",
		"${first_name()}",
		"${username()}",
		"${paragraph()}",
		"${gender_abbrev()}",
		"${row.age > 18 ? 1 : 0}",
	}

	for _, v := range templates {
		if _, err := EvalTemplate(v, Ctx); err == nil {
			t.Errorf("expect %v to fail", v)
		}
	}
}

func TestEvalPass(t *testing.T) {

	Ctx := buildTemplateContext()

	templates := []string{
		"${digits_n(2)}",
		"${brand()}",
		"${first_name()}",
		"${company()}",
		"${1 + 3}",
		"${2 > 7}",
	}

	for _, v := range templates {
		if _, err := EvalTemplate(v, Ctx); err != nil {
			t.Error(err)
		}
	}
}
