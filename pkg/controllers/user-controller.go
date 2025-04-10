package controllers

import(
	"encoding/json"
	"net/http"
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"strconv"
	"github.com/gorilla/mux"
	"github.com/PhucNguyen0304/go-bookstore/pkg/models"
	"github.com/PhucNguyen0304/go-bookstore/pkg/utils"
	"golang.org/x/crypto/bcrypt"
	"time"
    "github.com/dgrijalva/jwt-go"
)

// Global variable
var User *models.User
var jwtSecret []byte
var adminEmail string
var adminPassword string

func init() {
    // Load .env file
    err := godotenv.Load("c:/Users/Nguyen Phuc/OneDrive - Thuyloi University/Máy tính/go-bookstore/.env")
    if err != nil {
        fmt.Println("Error loading .env file:", err)
    }

    // Load JWT secret
    jwtSecret = []byte(os.Getenv("JWT_SECRET"))
    if len(jwtSecret) == 0 {
        fmt.Println("JWT_SECRET is not set in the .env file")
    }

    // Load admin credentials
    adminEmail = os.Getenv("ADMIN_EMAIL")
    adminPassword = os.Getenv("ADMIN_PASSWORD")
    if adminEmail == "" || adminPassword == "" {
        fmt.Println("Admin credentials are not set in the .env file")
    }
}

func Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	utils.ParseBody(r, &user)

	// Hash Password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password),bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Err while Hash Pass")
	}
	user.Password = string(hashedPassword)
	// Save User
	createUser := user.RegisterUser()
	res, _ := json.Marshal(createUser)
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var LoginRequest struct {
		Email string `json:"email"`
		Password string `json:"password"`
	}
	utils.ParseBody(r, &LoginRequest)
	// Find User
	user, err := models.GetUserByEmail(LoginRequest.Email)
	if err != nil {
		http.Error(w, "Invalid email", http.StatusUnauthorized)
		return
	}
	// Compare Hash And Password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(LoginRequest.Password))
	if err != nil {
		http.Error(w,"Invalid Password", http.StatusUnauthorized)
		return
	}
	// Generate Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
		"userId": user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})
	
	tokenString,err := token.SignedString(jwtSecret)

	if err != nil {
		http.Error(w,"Err while Generate token",http.StatusUnauthorized)
		return
	}
	res := map[string]string{"token" : tokenString}
	json.NewEncoder(w).Encode(res)
}

func GetAllUser(w http.ResponseWriter, r *http.Request) {
	Users := models.GetAllUser()
	res,err := json.Marshal(Users)
	if err != nil {
		fmt.Println("Err while parse")
	}
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	UserId := vars["userId"]
	id,err := strconv.ParseInt(UserId, 0, 0)
	if err != nil {
		fmt.Println("Err While Parse Id To Int")
	}
	User,_ := models.GetUserById(id)
	res,_ := json.Marshal(User)
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func DeleteUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	UserId := vars["userId"]
	id,err := strconv.ParseInt(UserId, 0, 0)
	if err != nil {
		fmt.Println("Err while Parse")
	}
	user := models.DeleteUser(id)
	res, _ := json.Marshal(user)
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var updateUser = &models.User{}
	utils.ParseBody(r, updateUser)
	vars := mux.Vars(r)
	userId := vars["userId"]
	id, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		fmt.Println("Err While Parse")
	}
	User,db := models.GetUserById(id)
	if User.Name != "" {
		User.Name = updateUser.Name
	}
	if User.Email != "" {
		User.Email = updateUser.Email
	}
	if User.Password != "" {
		User.Password = updateUser.Password
	}
	// Save user to db
	db.Save(&User)
	res, _ := json.Marshal(User)
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

 /*-------------Admin Login --------------------*/
func AdminLogin(w http.ResponseWriter, r *http.Request) {
    var loginRequest struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }
    utils.ParseBody(r, &loginRequest)

    // Check email and password admin
    if loginRequest.Email != adminEmail || loginRequest.Password != adminPassword {
        http.Error(w, "Invalid admin credentials", http.StatusUnauthorized)
        return
    }

    // Create jwt token
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "admin": true,
        "exp":   time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
    })

    tokenString, err := token.SignedString(jwtSecret)
    if err != nil {
        http.Error(w, "Error generating token", http.StatusInternalServerError)
        return
    }

    // Return Token
    res := map[string]string{"token": tokenString}
    json.NewEncoder(w).Encode(res)
}