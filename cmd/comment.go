package cmd

import (
	"fmt"
	"strings"
)

// comment represents a single line comment
type comment struct {
	start string
	end   string
}

var comments = map[string]comment{
	"go":    {"//", ""},
	"c":     {"//", ""},
	"cpp":   {"//", ""},
	"cs":    {"//", ""},
	"java":  {"//", ""},
	"js":    {"//", ""},
	"ts":    {"//", ""},
	"swift": {"//", ""},
	"kt":    {"//", ""},
	"rs":    {"//", ""},
	"php":   {"//", ""},
	"py":    {"#", ""},
	"rb":    {"#", ""},
	"sh":    {"#", ""},
	"lua":   {"--", ""},
	"html":  {"<!--", "-->"},
	"xml":   {"<!--", "-->"},
	"css":   {"/*", "*/"},
	"scss":  {"//", ""},
	"sql":   {"--", ""},
}

func generateComment(fname, nit string) string {
	spl := strings.Split(fname, ".")
  c, ok := comments[spl[len(spl) - 1]]
  if !ok {
    return ""
  }
	return strings.TrimSpace(fmt.Sprintf("%s %s %s", c.start, nit, c.end))
}
