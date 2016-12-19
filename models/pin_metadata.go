package models

import (
	"github.com/BrandonRomano/iso8601"
)

type PinMetadata struct {
	Article Article `json:"article"`
	Link    Link    `json:"link"`
	Place   Place   `json:"place"`
	Movie   Movie   `json:"movie"`
	Product Product `json:"product"`
}

type MetadataPerson struct {
	Name string `json:"name"`
}

// ========== Meta: Article ==========

type Article struct {
	PublishedAt iso8601.Time     `json:"published_at"`
	Description string           `json:"description"`
	Name        string           `json:"name"`
	Authors     []MetadataPerson `json:"authors"`
}

// ========== Meta: Link ==========

type Link struct {
	Locale      string `json:"locale"`
	Title       string `json:"title"`
	SiteName    string `json:"site_name"`
	Description string `json:"description"`
	Favicon     string `json:"favicon"`
}

// ========== Meta: Place ==========

type Place struct {
	Category   string  `json:"category"`
	Name       string  `json:"name"`
	Locality   string  `json:"locality"`
	Country    string  `json:"country"`
	Region     string  `json:"region"`
	Longitude  float32 `json:"longitude"`
	SourceUrl  string  `json:"source_url"`
	Street     string  `json:"street"`
	PostalCode string  `json:"postal_code"`
	Latitude   float32 `json:"latitude"`
}

// ========== Meta: Movie ==========

type Movie struct {
	Rating      string           `json:"rating"`
	Directors   []MetadataPerson `json:"directors"`
	Actors      []MetadataPerson `json:"actors"`
	Name        string           `json:"name"`
	PublishedAt iso8601.Time     `json:"published_at"`
}

// ========== Meta: Product ==========

type Product struct {
	Name  string       `json:"name"`
	Offer ProductOffer `json:"offer"`
}

type ProductOffer struct {
	Price   string `json:"price"`
	InStock bool   `json:"in_stock"`
}

// ========== Meta: Recipe ==========

type Recipe struct {
	Servings    RecipeServings   `json:"servings"`
	Name        string           `json:"name"`
	Ingredients []RecipeCategory `json:"ingredients"`
}

type RecipeServings struct {
	Serves  string `json:"serves"`
	Summary string `json:"summary"`
}

type RecipeCategory struct {
	Category    string             `json:"category"`
	Ingredients []RecipeIngredient `json:"ingredients"`
}

type RecipeIngredient struct {
	Amount string `json:"amount"`
	Name   string `json:"name"`
}
