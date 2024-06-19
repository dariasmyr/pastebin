// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

type CreateDocumentInput struct {
	Value        string  `json:"value"`
	Title        *string `json:"title,omitempty"`
	AccessKey    string  `json:"accessKey"`
	MaxViewCount *int    `json:"maxViewCount,omitempty"`
	TTLMs        *int    `json:"ttlMs,omitempty"`
}

type Document struct {
	ID           int       `json:"id"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	Title        *string   `json:"title,omitempty"`
	Value        string    `json:"value"`
	AccessKey    string    `json:"accessKey"`
	ViewCount    int       `json:"viewCount"`
	MaxViewCount int       `json:"maxViewCount"`
	TTLMs        int       `json:"ttlMs"`
}

type Mutation struct {
}

type Query struct {
}
