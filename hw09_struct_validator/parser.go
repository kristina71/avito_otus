package hw09structvalidator

import "reflect"

func parseFieldRules(typ reflect.StructField, validateTag string) (fieldRules, error) {
	kind := typ.Type.Kind()
	field := typ.Name
	var rules fieldRules
	var err error

	switch kind {
	case reflect.String:
		rules, err = fillStringRules(field, validateTag, validateRegular)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		rules, err = fillIntRules(field, validateTag, validateRegular)
	case reflect.Slice:
		sliceKind := typ.Type.Elem().Kind()
		switch sliceKind {
		case reflect.String:
			rules, err = fillStringRules(field, validateTag, validateSlice)
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			rules, err = fillIntRules(field, validateTag, validateSlice)
		default:
			return nil, &ErrIncorrectUse{reason: IncorrectFieldType, field: field, kind: sliceKind}
		}
	default:
		return nil, &ErrIncorrectUse{reason: IncorrectFieldType, field: field, kind: kind}
	}

	return rules, err
}

func parseStructRules(value reflect.Value) (structRules, error) {
	var sr structRules
	for i := 0; i < value.NumField(); i++ {
		fieldType := value.Type().Field(i)
		if fieldType.PkgPath != "" {
			sr = append(sr, nil)
			continue
		}

		validateTag := fieldType.Tag.Get(validatorTag)
		if validateTag == "" {
			sr = append(sr, nil)
			continue
		}

		rules, err := parseFieldRules(fieldType, validateTag)
		if err != nil {
			return nil, err
		}
		sr = append(sr, rules)
	}
	return sr, nil
}
