package avatar

import (
	"math/rand"
	"reflect"
	"strings"
	"time"
	"unicode/utf8"
)

// isNilOrZero checks if the given value is either nil or the zero value of its type.
func isNilOrZero(input any) bool {
	val := reflect.ValueOf(input)
	switch val.Kind() {
	case reflect.Ptr, reflect.Interface, reflect.Map, reflect.Slice, reflect.Chan, reflect.Func:
		return val.IsNil() || val.IsZero()
	default:
		return val.IsZero()
	}
}

// extractLetter get first letter from utf string.
func extractLetter(text string) string {
	text = strings.ToLower(strings.TrimSpace(text))
	if len(text) > 0 {
		first, _ := utf8.DecodeRuneInString(text)
		return string(first)
	}
	return ""
}

// randomize generates a random integer in the range 0 to length-1.
func randomize(length int) int {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	return r.Intn(length)
}

// filter filters elements from a slice based on a given condition.
// It takes a slice of type T (where T is any type that supports comparison),
// and a condition function that checks whether an element should be included in the result.
// It returns a new slice containing only the elements that satisfy the condition.
func filter[T comparable](slice []T, condition func(T) bool) []T {
	var result []T
	for _, item := range slice {
		if condition(item) {
			result = append(result, item)
		}
	}
	return result
}

// mapKeys to get the keys of a map.
func mapKeys[T any](m map[string]T) []string {
	result := make([]string, 0)
	for key := range m {
		result = append(result, key)
	}
	return result
}

// toStringSlice convert type aliases slice string slice.
func toStringSlice[T ~string](vals []T) []string {
	res := make([]string, 0)
	for _, v := range vals {
		res = append(res, string(v))
	}
	return res
}
