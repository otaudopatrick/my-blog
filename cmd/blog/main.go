package main

import (
	"bytes"
	"context"
	"html/template"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/otaudopatrick/my-blog/database"
	"github.com/otaudopatrick/my-blog/internal/models"
	"github.com/otaudopatrick/my-blog/internal/utils"
	"github.com/yuin/goldmark"
)

func ConvertMarkdownToHTML(markdown string) (string, error) {
	var buf bytes.Buffer
	if err := goldmark.Convert([]byte(markdown), &buf); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func main() {
	database.ConnectDb()
	engine := html.New("./web/templates", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/static", "./web/static")

	app.Get("/", func(c *fiber.Ctx) error {
		var posts []models.Post
		database.Connection.Find(&posts)

		formattedPosts := make([]map[string]interface{}, len(posts))

		for i, post := range posts {
			formattedPosts[i] = map[string]interface{}{
				"Title": post.Title,
				"Date":  post.CreatedAt.Format("02-01-2006"),
				"Slug":  post.Slug,
			}
		}

		meta := utils.DefaultMetaTags()

		return c.Render("home", fiber.Map{
			"Meta":  meta,
			"Posts": formattedPosts,
		})
	})

	app.Get("/about", func(c *fiber.Ctx) error {

		meta := utils.DefaultMetaTags()

		return c.Render("about", fiber.Map{
			"Meta": meta,
		})
	})

	app.Get("/:slug", func(c *fiber.Ctx) error {
		slug := c.Params("slug")

		//TODO: Remover essa validação e adicionar o tratameno no static
		if slug == "favicon.ico" {
			return c.SendStatus(fiber.StatusNotFound)
		}

		var post models.Post
		ctx := context.Background()

		err := database.Connection.WithContext(ctx).
			Where("slug = ?", slug).
			First(&post).Error

		if err != nil {
			log.Fatal(err)
		}

		meta := utils.DefaultMetaTags()

		html, err := ConvertMarkdownToHTML(post.Body)

		if err != nil {
			log.Fatal(err)
		}

		return c.Render("posts/show", fiber.Map{
			"Meta":    meta,
			"Content": template.HTML(html),
			"Post":    post,
		})
	})

	log.Fatal(app.Listen(":3000"))
}
