package services

import (
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"blog/models"

	"gopkg.in/yaml.v3"
)

const dataDir = "data"

type PostService struct{}

func NewPostService() *PostService {
	os.MkdirAll(dataDir, 0755)
	return &PostService{}
}

func (s *PostService) ListPosts(page, perPage int) ([]models.Post, int, error) {
	entries, err := os.ReadDir(dataDir)
	if err != nil {
		return nil, 0, err
	}

	var posts []models.Post
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".md") {
			continue
		}
		post, err := models.ParseFile(filepath.Join(dataDir, entry.Name()))
		if err != nil {
			continue
		}
		posts = append(posts, *post)
	}

	sort.Slice(posts, func(i, j int) bool {
		return posts[i].Date > posts[j].Date
	})

	total := len(posts)
	start := (page - 1) * perPage
	if start >= total {
		return []models.Post{}, total, nil
	}
	end := start + perPage
	if end > total {
		end = total
	}

	return posts[start:end], total, nil
}

func (s *PostService) GetPost(slug string) (*models.Post, error) {
	path := filepath.Join(dataDir, slug+".md")
	return models.ParseFile(path)
}

func (s *PostService) CreatePost(title, tagsStr, content string) (*models.Post, error) {
	slug := models.Slugify(title)
	path := filepath.Join(dataDir, slug+".md")

	if _, err := os.Stat(path); err == nil {
		slug = slug + "-" + time.Now().Format("0102-150405")
		path = filepath.Join(dataDir, slug+".md")
	}

	var tags []string
	for _, t := range strings.Split(tagsStr, ",") {
		if t = strings.TrimSpace(t); t != "" {
			tags = append(tags, t)
		}
	}

	post := models.Post{
		Title:   title,
		Date:    time.Now().Format("2006-01-02"),
		Tags:    tags,
		Slug:    slug,
		Content: content,
	}

	fm, _ := yaml.Marshal(map[string]interface{}{
		"title": post.Title,
		"date":  post.Date,
		"tags":  post.Tags,
		"slug":  post.Slug,
	})

	fileContent := "---\n" + string(fm) + "---\n\n" + post.Content

	if err := os.WriteFile(path, []byte(fileContent), 0644); err != nil {
		return nil, err
	}

	return &post, nil
}

func (s *PostService) UpdatePost(slug, title, tagsStr, content string) (*models.Post, error) {
	path := filepath.Join(dataDir, slug+".md")

	existing, err := models.ParseFile(path)
	if err != nil {
		return nil, err
	}

	var tags []string
	for _, t := range strings.Split(tagsStr, ",") {
		if t = strings.TrimSpace(t); t != "" {
			tags = append(tags, t)
		}
	}

	existing.Title = title
	existing.Tags = tags
	existing.Content = content

	fm, _ := yaml.Marshal(map[string]interface{}{
		"title": existing.Title,
		"date":  existing.Date,
		"tags":  existing.Tags,
		"slug":  existing.Slug,
	})

	fileContent := "---\n" + string(fm) + "---\n\n" + existing.Content

	if err := os.WriteFile(path, []byte(fileContent), 0644); err != nil {
		return nil, err
	}

	return existing, nil
}

func (s *PostService) SaveDraft(slug, title, tagsStr, content string) (*models.Post, error) {
	path := filepath.Join(dataDir, slug+".md")

	existing, err := models.ParseFile(path)
	if err != nil {
		return nil, err
	}

	if title != "" {
		existing.Title = title
	}

	var tags []string
	for _, t := range strings.Split(tagsStr, ",") {
		if t = strings.TrimSpace(t); t != "" {
			tags = append(tags, t)
		}
	}
	if len(tags) > 0 {
		existing.Tags = tags
	}

	existing.Content = content

	fm, _ := yaml.Marshal(map[string]interface{}{
		"title": existing.Title,
		"date":  existing.Date,
		"tags":  existing.Tags,
		"slug":  existing.Slug,
	})

	fileContent := "---\n" + string(fm) + "---\n\n" + existing.Content

	if err := os.WriteFile(path, []byte(fileContent), 0644); err != nil {
		return nil, err
	}

	return existing, nil
}

func (s *PostService) DeletePost(slug string) error {
	path := filepath.Join(dataDir, slug+".md")
	return os.Remove(path)
}

func (s *PostService) ListPostsByTag(tag string) ([]models.Post, error) {
	entries, err := os.ReadDir(dataDir)
	if err != nil {
		return nil, err
	}

	var posts []models.Post
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".md") {
			continue
		}
		post, err := models.ParseFile(filepath.Join(dataDir, entry.Name()))
		if err != nil {
			continue
		}
		for _, t := range post.Tags {
			if strings.EqualFold(t, tag) {
				posts = append(posts, *post)
				break
			}
		}
	}

	sort.Slice(posts, func(i, j int) bool {
		return posts[i].Date > posts[j].Date
	})

	return posts, nil
}
