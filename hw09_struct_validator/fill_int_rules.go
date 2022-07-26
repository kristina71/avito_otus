package hw09structvalidator

import "strings"

func fillIntRules(field, value string, vKind validateKind) (fieldRules, error) {
	rules := &intRules{
		field: field,
		vKind: vKind,
	}
	strs := strings.Split(value, "|")
	for _, str := range strs {
		pair := strings.Split(str, ":")
		if len(pair) != 2 {
			return nil, &ErrIncorrectUse{reason: IncorrectCondition, field: field, rule: str}
		}

		ruleName := pair[0]
		ruleValue := pair[1]
		var rule intRule
		var err error
		switch ruleName {
		case "min":
			rule, err = newIntMin(ruleValue)
		case "max":
			rule, err = newIntMax(ruleValue)
		case "in":
			rule, err = newIntIn(ruleValue)
		default:
			return nil, &ErrIncorrectUse{reason: UnknownRule, field: field, rule: ruleName}
		}
		if err != nil {
			return nil, &ErrIncorrectUse{reason: IncorrectCondition, field: field, rule: ruleName, err: err}
		}
		rules.rules = append(rules.rules, rule)
	}
	if len(rules.rules) == 0 {
		return nil, nil
	}
	return rules, nil
}
