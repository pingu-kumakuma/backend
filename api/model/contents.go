package model

import (
	"log"
	"time"
)

type Contents struct {
	ID         uint      `json:"id"`
	Title      string      `json:"title"`
	Category   string    `json:"category"`
	Curriculum string    `json:"curriculum"`
	Content    string    `json:"content"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func GetContent(id uint) (contents Contents, err error) {
	err = Db.QueryRow("SELECT title, category, curriculum, content, created_at, updated_at FROM contents WHERE id = ?", id).Scan(
		&contents.ID,
		&contents.Title,
		&contents.Category,
		&contents.Curriculum,
		&contents.Content,
		&contents.CreatedAt,
		&contents.UpdatedAt,
	)
	if err != nil {
		log.Fatalln(err)
	}
	return contents, err
}

func GetContents() (contents []Contents, err error) {
	rows, err := Db.Query("SELECT title, category, curriculum, content, created_at, updated_at FROM contents")
	if err != nil {
		log.Fatalln(err)
	}
	defer rows.Close()

	for rows.Next() {
		var item Contents
		err = rows.Scan(
			&item.ID,
			&item.Title,
			&item.Category,
			&item.Curriculum,
			&item.Content,
			&item.CreatedAt,
			&item.UpdatedAt,
		)
		contents = append(contents, item)
	}

	return contents, err
}

func CreateContent(title string, category string, curriculum string, content string) (err error) {
	now := time.Now()
	_, err = Db.Exec("INSERT INTO contents (title,category, curriculum, content, created_at, updated_at) VALUES(?,?,?,?, ?, ?)", title, category, curriculum, content, now, now)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func (contents *Contents) UpdateContent() error {
	_, err = Db.Exec("update contents set content = ? where title = ?", contents.Content, contents.Title)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func (contents *Contents) DeleteContent() error {
	_, err = Db.Exec("delete from contents where title = ?", contents.Title)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}
