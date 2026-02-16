package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"sync"

	"test_server/models"

	"github.com/gorilla/mux"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
}

var (
	users    = make(map[int]*models.User)
	products = make(map[int]*Product)

	userID    = 1
	productID = 1

	mu sync.RWMutex
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/v1/users", createUser).Methods("POST")
	r.HandleFunc("/api/v1/users", getUsers).Methods("GET")
	r.HandleFunc("/api/v1/users/{id}", getUser).Methods("GET")
	r.HandleFunc("/api/v1/users/{id}", updateUser).Methods("PUT")
	r.HandleFunc("/api/v1/users/{id}", deleteUser).Methods("DELETE")

	r.HandleFunc("/api/v1/products", createProduct).Methods("POST")
	r.HandleFunc("/api/v1/products", getProducts).Methods("GET")
	r.HandleFunc("/api/v1/products/{id}", getProduct).Methods("GET")
	r.HandleFunc("/api/v1/products/{id}", updateProduct).Methods("PUT")
	r.HandleFunc("/api/v1/products/{id}", deleteProduct).Methods("DELETE")

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

// @testgen router=/api/v1/users struct=models.User
func createUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	mu.Lock()
	user.ID = userID
	users[userID] = &user
	userID++
	mu.Unlock()
	respondJSON(w, http.StatusCreated, user)
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	mu.RLock()
	list := make([]*models.User, 0, len(users))
	for _, u := range users {
		list = append(list, u)
	}
	mu.RUnlock()
	respondJSON(w, http.StatusOK, list)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	mu.RLock()
	user, ok := users[id]
	mu.RUnlock()
	if !ok {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	respondJSON(w, http.StatusOK, user)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	mu.Lock()
	user, ok := users[id]
	if !ok {
		mu.Unlock()
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	json.NewDecoder(r.Body).Decode(user)
	user.ID = id
	mu.Unlock()
	respondJSON(w, http.StatusOK, user)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	mu.Lock()
	delete(users, id)
	mu.Unlock()
	respondJSON(w, http.StatusOK, map[string]string{"status": "deleted"})
}

// @testgen router=/api/v1/products struct=Product
func createProduct(w http.ResponseWriter, r *http.Request) {
	var product Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	mu.Lock()
	product.ID = productID
	products[productID] = &product
	productID++
	mu.Unlock()
	respondJSON(w, http.StatusCreated, product)
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	mu.RLock()
	list := make([]*Product, 0, len(products))
	for _, p := range products {
		list = append(list, p)
	}
	mu.RUnlock()
	respondJSON(w, http.StatusOK, list)
}

func getProduct(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	mu.RLock()
	product, ok := products[id]
	mu.RUnlock()
	if !ok {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	respondJSON(w, http.StatusOK, product)
}

func updateProduct(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	mu.Lock()
	product, ok := products[id]
	if !ok {
		mu.Unlock()
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	json.NewDecoder(r.Body).Decode(product)
	product.ID = id
	mu.Unlock()
	respondJSON(w, http.StatusOK, product)
}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	mu.Lock()
	delete(products, id)
	mu.Unlock()
	respondJSON(w, http.StatusOK, map[string]string{"status": "deleted"})
}

func respondJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
