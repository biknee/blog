package handlers

import (
	"bytes"
	"html/template"
	"net/http"
	"strconv"
	"strings"

	"blog/models"
	"blog/services"

	"github.com/gin-gonic/gin"
	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting/v2"
)

type PostHandler struct {
	svc *services.PostService
	md  goldmark.Markdown
}

func NewPostHandler(svc *services.PostService) *PostHandler {
	md := goldmark.New(
		goldmark.WithExtensions(
			highlighting.NewHighlighting(
				highlighting.WithStyle("monokai"),
			),
		),
	)
	return &PostHandler{svc: svc, md: md}
}

func (h *PostHandler) Home(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	if page < 1 {
		page = 1
	}

	posts, total, err := h.svc.ListPosts(page, 12)
	if err != nil {
		c.String(http.StatusInternalServerError, "加载文章列表失败")
		return
	}

	totalPages := (total + 11) / 12
	if totalPages < 1 {
		totalPages = 1
	}

	hasMore := page < totalPages

	c.HTML(http.StatusOK, "home.html", gin.H{
		"Title":      "我的博客",
		"Posts":      posts,
		"Page":       page,
		"TotalPages": totalPages,
		"HasMore":    hasMore,
		"NextPage":   page + 1,
	})
}

func (h *PostHandler) View(c *gin.Context) {
	slug := c.Param("slug")
	post, err := h.svc.GetPost(slug)
	if err != nil {
		c.HTML(http.StatusNotFound, "404.html", gin.H{
			"Title": "页面不存在 — 我的博客",
		})
		return
	}

	var buf bytes.Buffer
	if err := h.md.Convert([]byte(post.Content), &buf); err != nil {
		c.String(http.StatusInternalServerError, "渲染文章失败")
		return
	}

	c.HTML(http.StatusOK, "post.html", gin.H{
		"Title":   post.Title + " — 我的博客",
		"Post":    post,
		"Content": template.HTML(buf.String()),
	})
}

func (h *PostHandler) NewForm(c *gin.Context) {
	c.HTML(http.StatusOK, "new.html", gin.H{
		"Title": "新建文章 — 我的博客",
	})
}

func (h *PostHandler) Create(c *gin.Context) {
	title := strings.TrimSpace(c.PostForm("title"))
	tags := strings.TrimSpace(c.PostForm("tags"))
	content := strings.TrimSpace(c.PostForm("content"))

	if title == "" || content == "" {
		c.HTML(http.StatusBadRequest, "new.html", gin.H{
			"Title":       "新建文章 — 我的博客",
			"Error":       "标题和内容不能为空。",
			"FormTitle":   title,
			"FormTags":    tags,
			"FormContent": content,
		})
		return
	}

	post, err := h.svc.CreatePost(title, tags, content)
	if err != nil {
		c.String(http.StatusInternalServerError, "创建文章失败")
		return
	}

	c.Redirect(http.StatusFound, "/post/"+post.Slug)
}

func (h *PostHandler) EditForm(c *gin.Context) {
	slug := c.Param("slug")
	post, err := h.svc.GetPost(slug)
	if err != nil {
		c.HTML(http.StatusNotFound, "404.html", gin.H{
			"Title": "页面不存在 — 我的博客",
		})
		return
	}

	c.HTML(http.StatusOK, "edit.html", gin.H{
		"Title":    "编辑：" + post.Title + " — 我的博客",
		"Post":     post,
		"FormTags": strings.Join(post.Tags, ", "),
	})
}

func (h *PostHandler) Update(c *gin.Context) {
	slug := c.Param("slug")
	title := strings.TrimSpace(c.PostForm("title"))
	tags := strings.TrimSpace(c.PostForm("tags"))
	content := strings.TrimSpace(c.PostForm("content"))

	if title == "" || content == "" {
		c.HTML(http.StatusBadRequest, "edit.html", gin.H{
			"Title":   "编辑文章 — 我的博客",
			"Error":   "标题和内容不能为空。",
			"Post":    &models.Post{Slug: slug, Title: title, Content: content},
			"FormTags": tags,
		})
		return
	}

	post, err := h.svc.UpdatePost(slug, title, tags, content)
	if err != nil {
		c.String(http.StatusInternalServerError, "更新文章失败")
		return
	}

	c.Redirect(http.StatusFound, "/post/"+post.Slug)
}

func (h *PostHandler) Delete(c *gin.Context) {
	slug := c.Param("slug")
	if err := h.svc.DeletePost(slug); err != nil {
		c.String(http.StatusInternalServerError, "删除文章失败")
		return
	}
	c.Redirect(http.StatusFound, "/")
}

func (h *PostHandler) Tag(c *gin.Context) {
	tag := c.Param("tag")
	posts, err := h.svc.ListPostsByTag(tag)
	if err != nil {
		c.String(http.StatusInternalServerError, "加载文章列表失败")
		return
	}

	c.HTML(http.StatusOK, "home.html", gin.H{
		"Title":      "标签：" + tag + " — 我的博客",
		"Posts":      posts,
		"Page":       1,
		"TotalPages": 1,
		"Tag":        tag,
	})
}
