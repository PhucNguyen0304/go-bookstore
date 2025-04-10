package controllers

import (
    "encoding/json"
    "net/http"
    "strconv"

    "github.com/PhucNguyen0304/go-bookstore/pkg/models"
    "github.com/PhucNguyen0304/go-bookstore/pkg/utils"
    "github.com/gorilla/mux"
)

// Add Product To Cart
func AddItemToCart(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    userID, err := strconv.ParseUint(params["userId"], 10, 32)
    if err != nil {
        http.Error(w, "Invalid user ID", http.StatusBadRequest)
        return
    }

    var requestBody struct {
        ProductID uint `json:"productId"`
        Quantity  int  `json:"quantity"`
    }
    utils.ParseBody(r, &requestBody)
    if err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }   

    err = models.AddItemToCart(uint(userID), requestBody.ProductID, requestBody.Quantity)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"message": "Item added to cart successfully"})
}

// Get Cart By User ID
func GetCart(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    userID, err := strconv.ParseUint(params["userId"], 10, 32)
    if err != nil {
        http.Error(w, "Invalid user ID", http.StatusBadRequest)
        return
    }

    carts, err := models.GetCartByUserID(uint(userID))
    if err != nil {
        http.Error(w, "Failed to retrieve cart", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(carts)
}