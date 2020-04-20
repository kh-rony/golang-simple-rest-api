package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type Product struct {
	ID       int       `json:"id,omitempty"`
	Name     string    `json:"name,omitempty"`
	Price    int       `json:"price,omitempty"`
	Category *Category `json:"category,omitempty"`
}

type Category struct {
	Name   string `json:"category-name,omitempty"`
	Vendor string `json:"vendor,omitempty"`
}

var products []Product

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/products", GetAllProducts).Methods("GET")
	router.HandleFunc("/product/{id}", GetProduct).Methods("GET")
	router.HandleFunc("/product", CreateNewProduct).Methods("POST")
	router.HandleFunc("/product/{id}", UpdateProduct).Methods("PUT")
	router.HandleFunc("/product/{id}", DeleteProduct).Methods("DELETE")

	products = append(products, Product{ID: 1, Name: "Apple", Price: 150, Category: &Category{Name: "Fruit", Vendor: "Vendor1"}})
	products = append(products, Product{ID: 2, Name: "Orange", Price: 120, Category: &Category{Name: "Fruit", Vendor: "Vendor1"}})
	products = append(products, Product{ID: 3, Name: "Mango", Price: 80, Category: &Category{Name: "Fruit", Vendor: "Vendor1"}})
	products = append(products, Product{ID: 4, Name: "Milk", Price: 80, Category: &Category{Name: "Dairy", Vendor: "Vendor2"}})
	products = append(products, Product{ID: 5, Name: "Cheese", Price: 90, Category: &Category{Name: "Dairy", Vendor: "Vendor2"}})

	log.Fatal(http.ListenAndServe(":8100", router))
}

func GetAllProducts(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func GetProduct(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	for _, product := range products {

		//if strconv.Itoa(product.ID) == params["id"] {
		//
		//	w.Header().Set("Content-Type", "application/json")
		//	json.NewEncoder(w).Encode(product)
		//	return
		//}

		if id, _ := strconv.Atoi(params["id"]); id == product.ID {

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(product)
			return
		}
	}
}

func CreateNewProduct(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)

	var newProduct Product
	decoder.Decode(&newProduct)

	products = append(products, Product{ID: newProduct.ID, Name: newProduct.Name, Price: newProduct.Price, Category: &Category{Name: newProduct.Category.Name, Vendor: newProduct.Category.Vendor}})

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)

	var updatedProduct Product
	decoder.Decode(&updatedProduct)


	params := mux.Vars(r)

	for i := 0; i < len(products); i++ {

		if id, _ := strconv.Atoi(params["id"]); id == products[i].ID {

			products[i].Name = updatedProduct.Name
			products[i].Price = updatedProduct.Price
			products[i].Category.Name = updatedProduct.Category.Name
			products[i].Category.Vendor = updatedProduct.Category.Vendor

			break
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	for i := 0; i < len(products); i++ {

		if id, _ := strconv.Atoi(params["id"]); id == products[i].ID {

			products = append( products[:i], products[i+1:]...)
			break
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}