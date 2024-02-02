package arangodb

import (
	"time"

	"github.com/arangodb/go-driver"
)

type Edge struct {
	Document

	From string `json:"_from"`
	To   string `json:"_to"`
}

type Document struct {
	ArangoInterface `json:"-"`
	driver.DocumentMeta

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ArangoInterface interface {
	GetID() string
	GetKey() string
	Set(ID, Key, Rev string)
	InitializeTimestamp()
	UpdateTimestamp()
}

func (d *Document) Set(ID, Key, Rev string) {
	d.ID = driver.DocumentID(ID)
	d.Key = Key
	d.Rev = Rev
}

func (d *Document) GetID() string {
	return string(d.ID)
}

func (d *Document) GetKey() string {
	return d.Key
}

func (d *Document) InitializeTimestamp() {
	d.CreatedAt = time.Now()
	d.UpdatedAt = time.Now()
}

func (d *Document) UpdateTimestamp() {
	d.UpdatedAt = time.Now()
}
