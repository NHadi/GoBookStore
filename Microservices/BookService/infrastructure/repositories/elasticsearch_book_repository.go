package repositories

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"BookStore/domain/entities"

	"github.com/elastic/go-elasticsearch/v7"
)

type ElasticsearchBookRepository struct {
	client *elasticsearch.Client
	index  string
}

// NewElasticsearchBookRepository creates a new instance of ElasticsearchBookRepository
func NewElasticsearchBookRepository() (*ElasticsearchBookRepository, error) {
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://elasticsearch:9200", // Match docker service name
		},
	}
	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return nil, err
	}

	return &ElasticsearchBookRepository{
		client: client,
		index:  "books",
	}, nil
}

// IndexBook adds a new book document to Elasticsearch
func (r *ElasticsearchBookRepository) IndexBook(book *entities.Book) error {
	data, err := json.Marshal(book)
	if err != nil {
		return err
	}

	req := bytes.NewReader(data)
	res, err := r.client.Index(r.index, req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("Error indexing book: %s", res.Status())
	}

	return nil
}

// SearchBooks searches for books in Elasticsearch by title
func (r *ElasticsearchBookRepository) SearchBooks(title string) ([]entities.Book, error) {
	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"title": title,
			},
		},
	}

	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		return nil, err
	}

	res, err := r.client.Search(
		r.client.Search.WithContext(context.Background()),
		r.client.Search.WithIndex(r.index),
		r.client.Search.WithBody(&buf),
	)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.IsError() {
		return nil, fmt.Errorf("Error searching books: %s", res.Status())
	}

	var results struct {
		Hits struct {
			Hits []struct {
				Source entities.Book `json:"_source"`
			} `json:"hits"`
		} `json:"hits"`
	}

	if err := json.NewDecoder(res.Body).Decode(&results); err != nil {
		return nil, err
	}

	var books []entities.Book
	for _, hit := range results.Hits.Hits {
		books = append(books, hit.Source)
	}

	return books, nil
}

// Add adds a book to Elasticsearch (similar to IndexBook)
func (r *ElasticsearchBookRepository) Add(book *entities.Book) {
	if err := r.IndexBook(book); err != nil {
		// Handle the error appropriately
	}
}

// List retrieves all books from Elasticsearch
func (r *ElasticsearchBookRepository) List() []*entities.Book {
	// Implement the logic to retrieve all books from Elasticsearch
	// You might want to define a match_all query to fetch all documents
	var books []*entities.Book

	res, err := r.client.Search(
		r.client.Search.WithIndex(r.index),
		r.client.Search.WithBody(bytes.NewReader([]byte(`{"query": {"match_all": {}}}`))),
	)
	if err != nil {
		// Handle the error appropriately
		return nil
	}
	defer res.Body.Close()

	var results struct {
		Hits struct {
			Hits []struct {
				Source entities.Book `json:"_source"`
			} `json:"hits"`
		} `json:"hits"`
	}

	if err := json.NewDecoder(res.Body).Decode(&results); err != nil {
		// Handle the error appropriately
		return nil
	}

	for _, hit := range results.Hits.Hits {
		books = append(books, &hit.Source)
	}

	return books
}

// Update modifies an existing book in Elasticsearch
func (r *ElasticsearchBookRepository) Update(book *entities.Book) error {
	// Implement the logic to update a book in Elasticsearch
	data, err := json.Marshal(book)
	if err != nil {
		return err
	}

	req := bytes.NewReader(data)
	res, err := r.client.Update(r.index, book.ID, req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("Error updating book: %s", res.Status())
	}

	return nil
}

// Delete removes a book from Elasticsearch
func (r *ElasticsearchBookRepository) Delete(id string) error {
	res, err := r.client.Delete(r.index, id)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("Error deleting book: %s", res.Status())
	}

	return nil
}

// FindById retrieves a book from Elasticsearch by ID
func (r *ElasticsearchBookRepository) FindById(id string) (*entities.Book, error) {
	res, err := r.client.Get(r.index, id)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.IsError() {
		return nil, fmt.Errorf("Error retrieving book: %s", res.Status())
	}

	var book entities.Book
	if err := json.NewDecoder(res.Body).Decode(&book); err != nil {
		return nil, err
	}

	return &book, nil
}
