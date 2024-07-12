package main

import "fmt"

// Constants: create a custom memory cache

const GlobalLimit = 100 // Constants require an initial value, types are optional

const MaxCacheSize = 10 * GlobalLimit

const (
	CacheKeyBook = "book_"
	CacheKeyCD = "cd_"
)

var cache map[string]string

func cacheGet(key string) string { // Get items from cache
	return cache[key]
}

func cacheSet(key, val string) { // Set items in the cache
	if len(cache)+1 >= MaxCacheSize {
		return
	}
	cache[key] = val
}

func GetBook(isbn string) string { // Add a book to the cache
	return cacheGet(CacheKeyBook + isbn)
}

func SetBook(isbn string, name string) { // Add a book to the cache
	cacheSet(CacheKeyBook + isbn, name)
}

func GetCD(sku string) string {
	return cacheGet(CacheKeyCD + sku)
}

func SetCD(sku string, title string) {
	cacheSet(CacheKeyCD+sku, title)
}

func main() {
	cache = make(map[string]string)
	SetBook("1234-5678", "Get Ready To Go")
	SetCD("1234-5678", "Get Ready To Go Audio Book")
	fmt.Println("Book :", GetBook("1234-5678"))
	fmt.Println("CD :", GetCD("1234-5678"))
}