package config

import (
	"testing"
)

type NotValidated struct {
	A int
	B bool
	C string
}

type Validated struct {
	A         int
	B         bool
	C         string
	validated bool
}

func (t *Validated) Validate() error {
	t.validated = true
	return nil
}

type X struct {
	V1  Validated
	V2  *Validated
	V3  *Validated
	NV1 NotValidated
	NV2 *NotValidated
	NV3 *NotValidated
	Y   struct {
		V1  Validated
		V2  *Validated
		V3  *Validated
		NV1 NotValidated
		NV2 *NotValidated
		NV3 *NotValidated
	}
	validated bool
}

func (t *X) Validate() error {
	t.validated = true
	return nil
}

func Test_validate(t *testing.T) {
	x := X{
		V2:  &Validated{},
		NV2: &NotValidated{},
	}
	x.Y.V2 = &Validated{}
	x.Y.NV2 = &NotValidated{}

	err := validate(&x)
	failIfError(t, err)

	ok := x.validated && x.V1.validated && x.V2.validated && x.Y.V1.validated && x.Y.V2.validated
	if !ok {
		t.Errorf("%+v", x)
	}
}
