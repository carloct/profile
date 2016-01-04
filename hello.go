package main

import (
	"fmt"
	"os"
	"log"
	"html"
	"net/http"
	"io/ioutil"

	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
)

type Response struct {
	Status 	int 	`json:"Status"`
	Message string 	`json:"Message"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := httprouter.New()
	r.GET("/", Index)
	r.GET("/api/categories", Categories)
	r.GET("/api/category/:id", ProductsByCategory)
	r.GET("/api/products/:sku", ProductBySku)
	r.NotFound(NotFound)
	log.Fatal(http.ListenAndServe(":8080", r))
}

func Index(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func Categories(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	resp, err := http.Get(os.Getenv("API_PRODUCTS") + "all_categories")
	if err != nil {
		fmt.Printf("%s", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
        log.Fatalf("ERROR: %s", err)
    }
    defer resp.Body.Close()

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(200)
    w.Write(body);
}

func ProductsByCategory(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	url := os.Getenv("API_PRODUCTS") + "category/" + p.ByName("id")
	
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("%s", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
        log.Fatalf("ERROR: %s", err)
    }
    defer resp.Body.Close()

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(200)
    w.Write(body);
}

func ProductBySku(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	resp, err := http.Get(os.Getenv("API_PRODUCTS") + "product/" + p.ByName("sku"))
	if err != nil {
		fmt.Printf("%s", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
        log.Fatalf("ERROR: %s", err)
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(200)
    w.Write(body)
}

func NotFound(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	response = &Response{Status: 404, Message: "Resource not found"}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(404)
	w.Write(response)
}