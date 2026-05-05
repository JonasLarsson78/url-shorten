package store

import (
	"math/rand"

	"go-rest-api/db"
)

const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func Save(url string) (string, error) {
	code := make([]byte, 6)
	for i := range code {
		code[i] = chars[rand.Intn(len(chars))]
	}
	key := string(code)

	_, err := db.DB.Exec("INSERT INTO urls (code, url) VALUES (?, ?)", key, url)
	return key, err
}

func Get(code string) (string, bool) {
	var url string
	err := db.DB.QueryRow("SELECT url FROM urls WHERE code = ? AND deleted_at IS NULL", code).Scan(&url)
	if err != nil {
		return "", false
	}
	return url, true
}

func Delete(code string) (bool, error) {
	res, err := db.DB.Exec("UPDATE urls SET deleted_at = NOW() WHERE code = ? AND deleted_at IS NULL", code)
	if err != nil {
		return false, err
	}
	rows, _ := res.RowsAffected()
	return rows > 0, nil
}
