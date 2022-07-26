package hw09structvalidator

import "strings"

func fillStringRules(field, value string, vKind validateKind) (fieldRules, error) {
	rules := &stringRules{
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
		var rule stringRule
		var err error
		switch ruleName {
		case "len":
			rule, err = newStrLen(ruleValue)
		case "regexp":
			rule, err = newStrRegexp(ruleValue)
		case "in":
			rule = newStrIn(ruleValue)
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
