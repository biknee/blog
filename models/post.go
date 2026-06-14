package models

import (
	"fmt"
	"os"
	"strings"
	"time"

	"gopkg.in/yaml.v3"
)

type Post struct {
	Title   string   `yaml:"title"`
	Date    string   `yaml:"date"`
	Tags    []string `yaml:"tags"`
	Slug    string   `yaml:"slug"`
	Summary string   `yaml:"summary"`
	Content string   `yaml:"-"`
}

func ParseFile(path string) (*Post, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return ParseContent(string(data))
}

func ParseContent(raw string) (*Post, error) {
	raw = strings.TrimSpace(raw)
	if !strings.HasPrefix(raw, "---") {
		return nil, fmt.Errorf("missing frontmatter")
	}

	parts := strings.SplitN(raw, "---", 3)
	if len(parts) < 3 {
		return nil, fmt.Errorf("invalid frontmatter format")
	}

	var post Post
	if err := yaml.Unmarshal([]byte(strings.TrimSpace(parts[1])), &post); err != nil {
		return nil, fmt.Errorf("bad frontmatter: %w", err)
	}

	post.Content = strings.TrimSpace(parts[2])

	if post.Slug == "" {
		post.Slug = Slugify(post.Title)
	}
	if post.Date == "" {
		post.Date = time.Now().Format("2006-01-02")
	}
	if post.Title == "" {
		post.Title = post.Slug
	}
	if post.Summary == "" && len(post.Content) > 0 {
		post.Summary = truncateContent(post.Content, 150)
	}

	return &post, nil
}

func Slugify(title string) string {
	s := strings.ToLower(strings.TrimSpace(title))
	s = strings.ReplaceAll(s, " ", "-")
	var b strings.Builder
	for _, r := range s {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || r == '-' {
			b.WriteRune(r)
		}
	}
	slug := b.String()
	for strings.Contains(slug, "--") {
		slug = strings.ReplaceAll(slug, "--", "-")
	}
	slug = strings.Trim(slug, "-")
	if slug == "" {
		slug = "post-" + time.Now().Format("20060102-150405")
	}
	return slug
}

func truncateContent(s string, maxLen int) string {
	s = strings.ReplaceAll(s, "\n", " ")
	s = strings.Join(strings.Fields(s), " ")
	runes := []rune(s)
	if len(runes) <= maxLen {
		return s
	}
	return string(runes[:maxLen]) + "..."
}
