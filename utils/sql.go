package utils

import (
	"fmt"
	"strings"
)

// Like sql:Add '%' before and after the query string
func Like(q string) string {
	q = strings.TrimSpace(q)
	if q == "" {
		return ""
	}
	q = strings.Replace(q, "/", "//", -1)
	q = strings.Replace(q, "%", "/%", -1)
	q = strings.Replace(q, "_", "/_", -1)
	return fmt.Sprintf("%%%s%%", q)
}