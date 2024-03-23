package main

import "math/rand"

type Idea struct {
	ID         int         `json:"id"`
	Name       string      `json:"name"`
	Category   string      `json:"category"`
	Attributes []Attribute `json:"attributes"`
}

type Attribute struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type CreateIdeaRequest struct {
	Name       string      `json:"name"`
	Category   string      `json:"category"`
	Attributes []Attribute `json:"attributes"`
}

func NewIdea(name string, category string, attributes []Attribute) *Idea {
	return &Idea{
		Name:       name,
		Category:   category,
		Attributes: attributes,
	}
}