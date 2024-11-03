package utils

type MetaTags struct {
	Title       string
	Description string
	Keywords    string
	Author      string
}

func GenerateMetaTags(title, description, keywords, author string) MetaTags {
	return MetaTags{
		Title:       title,
		Description: description,
		Keywords:    keywords,
		Author:      author,
	}
}

func DefaultMetaTags() MetaTags {
	return GenerateMetaTags(
		"Patrick Luz | Blog",
		"Eu sou Patrick Luz, desenvolvedor apaixonado por tecnologia.",
		"blog, programação",
		"Patrick Luz",
	)
}
