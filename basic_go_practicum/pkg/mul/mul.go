package mul

import "strings"

func Mul(a interface{}, b int) interface{} {
	switch v2 := a.(type) {
	case int:
		return v2 * b
	case string:
		builder := strings.Builder{}
		for i := 0; i < b; i++ {
			builder.WriteString(v2)
		}
		return builder.String()
	}

	return nil
}
