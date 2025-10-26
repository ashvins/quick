package stringset

// Implement Set as a collection of unique string values.
//
// For Set.String, use '{' and '}', output elements as double-quoted strings
// safely escaped with Go syntax, and use a comma and a single space between
// elements. For example, a set with 2 elements, "a" and "b", should be formatted as {"a", "b"}.
// Format the empty set as {}.

// Define the Set type here.
import (
    "strings"
)

type Set struct {
	data map[string]struct{}
}

func New() Set {
	s := Set{}
	s.data = make(map[string]struct{}, 0)
	return s
}

func NewFromSlice(l []string) Set {
	s := New()
	for _, str := range l {
		s.data[str] = struct{}{} // insert into map
	}
	return s
}

func (s Set) String() string {
	sb := strings.Builder{}
	sb.WriteRune('{')
	doubleQuote := "\""
	prefix := doubleQuote
	for str := range s.data {
		sb.WriteString(prefix)
		sb.WriteString(str)
		sb.WriteString(doubleQuote)
		prefix = ", \"" // after the first string, this will be the prefix
	}
	sb.WriteRune('}')
	return sb.String()
}

func (s Set) IsEmpty() bool {
	return len(s.data) == 0
}

func (s Set) Has(elem string) bool {
	_, ok := s.data[elem]
	return ok
}

func (s Set) Add(elem string) {
	s.data[elem] = struct{}{}
}

func Subset(s1, s2 Set) bool {
	for k := range s1.data {
		if _, ok := s2.data[k]; !ok {
			return false
		}
	}
	return true
}

func Disjoint(s1, s2 Set) bool {
	return Intersection(s1, s2).IsEmpty()
}

func Equal(s1, s2 Set) bool {
	return Difference(s1, s2).IsEmpty() && Difference(s2, s1).IsEmpty()
}

func Intersection(s1, s2 Set) Set {
	sl := make([]string, 0, max(len(s1.data), len(s2.data))) // estimate elements in intersection
	for k := range s1.data {
		if _, ok := s2.data[k]; ok {
			sl = append(sl, k)
		}
	}

	return NewFromSlice(sl)
}

func Difference(s1, s2 Set) Set {
	sl := make([]string, 0, 10) // estimate 10 elements in difference
	for k := range s1.data {
		if _, ok := s2.data[k]; !ok {
			sl = append(sl, k)
		}
	}

	return NewFromSlice(sl)
}

func Union(s1, s2 Set) Set {
	sl := make([]string, 0, len(s1.data)+len(s2.data))
	for k := range s1.data {
		sl = append(sl, k)
	}
	for k := range s2.data {
		sl = append(sl, k)
	}
	return NewFromSlice(sl)
}
