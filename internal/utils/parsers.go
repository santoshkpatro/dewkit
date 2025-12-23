package utils

import "strings"

func SetNestedSettingsValue(m map[string]any, key string, value any) {
	parts := strings.Split(key, ".")
	current := m

	for i, part := range parts {
		if i == len(parts)-1 {
			current[part] = value
			return
		}

		if _, ok := current[part]; !ok {
			current[part] = make(map[string]any)
		}

		current = current[part].(map[string]any)
	}
}
