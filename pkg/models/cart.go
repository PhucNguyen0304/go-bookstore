package models

import (
    "errors"
    "github.com/jinzhu/gorm"
)

type Cart struct {
    gorm.Model
    UserId    uint `json:"userId"`    
    ProductID uint `json:"productId"`
    Quantity  int  `json:"quantity"` 
}

// Add product to cart
func AddItemToCart(userID uint, productID uint, quantity int) error {
    var cart Cart
    err := db.Where("user_id = ? AND product_id = ?", userID, productID).First(&cart).Error
    if gorm.IsRecordNotFoundError(err) {
        // If product not create yet => Create product
        cart = Cart{
            UserId:    userID,
            ProductID: productID,
            Quantity:  quantity,
        }
        return db.Create(&cart).Error
    } else if err != nil {
        return errors.New("failed to query cart")
    }

    // Update quantity product
    cart.Quantity = quantity
    return db.Save(&cart).Error
}


// Get cart
func GetCartByUserID(userID uint) ([]Cart, error) {
    var carts []Cart
    err := db.Where("user_id= ?", userID).Find(&carts).Error
    return carts, err
}