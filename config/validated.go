package config

import "reflect"

type IValidated interface {
	Validate() error
}

func validate(ptrToStruct interface{}) error {
	if validated, ok := ptrToStruct.(IValidated); validated != nil && ok {
		if err := validated.Validate(); err != nil {
			return err
		}
	}

	v := reflect.ValueOf(ptrToStruct)

	if v.Kind() != reflect.Ptr {
		panic(v)
	}

	v = v.Elem()
	if !v.IsValid() { // nil ptr
		panic(v)
	}

	// теперь v - сама структура
	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)

		var maybeValidated reflect.Value
		if f.Kind() == reflect.Struct {
			maybeValidated = f.Addr()
		} else if f.Kind() == reflect.Ptr && !f.IsNil() {
			maybeValidated = f
		}
		if maybeValidated.IsValid() && maybeValidated.CanInterface() {
			if validated, ok := maybeValidated.Interface().(IValidated); validated != nil && ok {
				if err := validated.Validate(); err != nil {
					return err
				}
			}
			if err := validate(maybeValidated.Interface()); err != nil {
				return err
			}
		}
	}

	return nil
}

/*
func validate(target interface{}) error {
	if validated, ok := target.(IValidated); ok && validated != nil {
		if err := validated.Validate(); err != nil {
			return err
		}
	}

	pv := reflect.ValueOf(target)
	if pv.Kind() != reflect.Ptr {
		return errors.New("")
	}
	v := pv.Elem()
	if v.Kind() != reflect.Struct {
		return errors.New("")
	}
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		ft := t.Field(i)
		fv := v.Field(i)

		for fv.Kind() == reflect.Ptr && !fv.IsNil() {
			fvv := fv.Elem()
		}

		ft.Type.
		if validated, ok := fv.Interface().(IValidated); ok && validated != nil {
			if err := validated.Validate(); err != nil {
				return err
			}
		}
	}

	return nil
}
*/
