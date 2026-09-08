package main

// Import: read a talent calculator link and return the class and point count.

import (
	"fmt"
	"net/url"
	"strings"
)

type Imported struct {
	Class  string
	Points int
	Format string // calculator, export, unknown
}

func ParseLink(raw string) (Imported, error) {
	u, err := url.Parse(strings.TrimSpace(raw))
	if err != nil {
		return Imported{}, err
	}
	frag := u.Fragment
	switch {
	case strings.Contains(frag, "tc-"):
		return Imported{classFromPath(u.Path), strings.Count(frag[4:], ""), "calculator"}, nil
	case strings.HasPrefix(frag, "I0"):
		return Imported{classFromPath(u.Path), len(frag) / 2, "export"}, nil
	default:
		return Imported{Format: "unknown"}, fmt.Errorf("link not recognized")
	}
}

func classFromPath(p string) string {
	parts := strings.Split(strings.Trim(p, "/"), "/")
	if len(parts) > 0 {
		return parts[0]
	}
	return ""
}
