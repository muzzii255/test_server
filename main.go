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

type Order struct {
	ID        int     `json:"id"`
	UserID    int     `json:"userId"`
	ProductID int     `json:"productId"`
	Quantity  int     `json:"quantity"`
	Total     float64 `json:"total"`
}

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type Review struct {
	ID        int    `json:"id"`
	ProductID int    `json:"productId"`
	UserID    int    `json:"userId"`
	Rating    int    `json:"rating"`
	Comment   string `json:"comment"`
}

var (
	users      = make(map[int]*models.User)
	products   = make(map[int]*Product)
	orders     = make(map[int]*Order)
	categories = make(map[int]*Category)
	reviews    = make(map[int]*Review)

	userID     = 1
	productID  = 1
	orderID    = 1
	categoryID = 1
	reviewID   = 1

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

	r.HandleFunc("/api/v1/orders", createOrder).Methods("POST")
	r.HandleFunc("/api/v1/orders", getOrders).Methods("GET")
	r.HandleFunc("/api/v1/orders/{id}", getOrder).Methods("GET")
	r.HandleFunc("/api/v1/orders/{id}", deleteOrder).Methods("DELETE")

	r.HandleFunc("/api/v1/categories", createCategory).Methods("POST")
	r.HandleFunc("/api/v1/categories", getCategories).Methods("GET")
	r.HandleFunc("/api/v1/categories/{id}", getCategory).Methods("GET")
	r.HandleFunc("/api/v1/categories/{id}", updateCategory).Methods("PUT")
	r.HandleFunc("/api/v1/categories/{id}", deleteCategory).Methods("DELETE")

	r.HandleFunc("/api/v1/reviews", createReview).Methods("POST")
	r.HandleFunc("/api/v1/reviews", getReviews).Methods("GET")
	r.HandleFunc("/api/v1/reviews/{id}", getReview).Methods("GET")
	r.HandleFunc("/api/v1/reviews/{id}", updateReview).Methods("PUT")
	r.HandleFunc("/api/v1/reviews/{id}", deleteReview).Methods("DELETE")

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

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

func createOrder(w http.ResponseWriter, r *http.Request) {
	var order Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	mu.Lock()
	order.ID = orderID
	orders[orderID] = &order
	orderID++
	mu.Unlock()
	respondJSON(w, http.StatusCreated, order)
}

func getOrders(w http.ResponseWriter, r *http.Request) {
	mu.RLock()
	list := make([]*Order, 0, len(orders))
	for _, o := range orders {
		list = append(list, o)
	}
	mu.RUnlock()
	respondJSON(w, http.StatusOK, list)
}

func getOrder(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	mu.RLock()
	order, ok := orders[id]
	mu.RUnlock()
	if !ok {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	respondJSON(w, http.StatusOK, order)
}

func deleteOrder(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	mu.Lock()
	delete(orders, id)
	mu.Unlock()
	respondJSON(w, http.StatusOK, map[string]string{"status": "deleted"})
}

func createCategory(w http.ResponseWriter, r *http.Request) {
	var category Category
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	mu.Lock()
	category.ID = categoryID
	categories[categoryID] = &category
	categoryID++
	mu.Unlock()
	respondJSON(w, http.StatusCreated, category)
}

func getCategories(w http.ResponseWriter, r *http.Request) {
	mu.RLock()
	list := make([]*Category, 0, len(categories))
	for _, c := range categories {
		list = append(list, c)
	}
	mu.RUnlock()
	respondJSON(w, http.StatusOK, list)
}

func getCategory(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	mu.RLock()
	category, ok := categories[id]
	mu.RUnlock()
	if !ok {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	respondJSON(w, http.StatusOK, category)
}

func updateCategory(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	mu.Lock()
	category, ok := categories[id]
	if !ok {
		mu.Unlock()
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	json.NewDecoder(r.Body).Decode(category)
	category.ID = id
	mu.Unlock()
	respondJSON(w, http.StatusOK, category)
}

func deleteCategory(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	mu.Lock()
	delete(categories, id)
	mu.Unlock()
	respondJSON(w, http.StatusOK, map[string]string{"status": "deleted"})
}

// Reviews handlers
func createReview(w http.ResponseWriter, r *http.Request) {
	var review Review
	if err := json.NewDecoder(r.Body).Decode(&review); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	mu.Lock()
	review.ID = reviewID
	reviews[reviewID] = &review
	reviewID++
	mu.Unlock()
	respondJSON(w, http.StatusCreated, review)
}

func getReviews(w http.ResponseWriter, r *http.Request) {
	mu.RLock()
	list := make([]*Review, 0, len(reviews))
	for _, rev := range reviews {
		list = append(list, rev)
	}
	mu.RUnlock()
	respondJSON(w, http.StatusOK, list)
}

func getReview(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	mu.RLock()
	review, ok := reviews[id]
	mu.RUnlock()
	if !ok {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	respondJSON(w, http.StatusOK, review)
}

func updateReview(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	mu.Lock()
	review, ok := reviews[id]
	if !ok {
		mu.Unlock()
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	json.NewDecoder(r.Body).Decode(review)
	review.ID = id
	mu.Unlock()
	respondJSON(w, http.StatusOK, review)
}

func deleteReview(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	mu.Lock()
	delete(reviews, id)
	mu.Unlock()
	respondJSON(w, http.StatusOK, map[string]string{"status": "deleted"})
}

func respondJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
