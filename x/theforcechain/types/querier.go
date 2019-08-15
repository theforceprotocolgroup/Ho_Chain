package types

import "strings"

// Query Result Payload for an id query
type QueryResIds []string

// implement fmt.Stringer
func (n QueryResIds) String() string {
	return strings.Join(n[:], "\n")
}
