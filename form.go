package form

import "github.com/creamsensation/gox"

type Form struct {
	Method      string
	Action      string
	CsrfToken   string
	CsrfName    string
	IsValid     bool
	IsSubmitted bool
}

func (f Form) Csrf() gox.Node {
	return Csrf(f.CsrfName, f.CsrfToken)
}
