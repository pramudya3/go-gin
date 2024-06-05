package main

import "errors"

type Article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var articleList = []Article{
	{
		ID:      1,
		Title:   "Title 1",
		Content: "Content 1",
	},
	{
		ID:      2,
		Title:   "Title 2",
		Content: "Content 2",
	},
	{
		ID:      3,
		Title:   "Title 3",
		Content: "Content 3",
	},
}

func getAllArticles() []Article {
	return articleList
}

func getArticleByID(id int) (*Article, error) {
	for _, a := range articleList {
		if a.ID == id {
			return &a, nil
		}
	}

	return nil, errors.New("Article not found")
}
