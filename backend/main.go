package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq" // The database driver
)

// User model for demonstration
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	// For local development, load .env file
	err := godotenv.Load("../.env")
	if err != nil {
		log.Println("Could not load .env file, using environment variables")
	}

	// TODO: Setup database connection using os.Getenv("DATABASE_URL")
	// db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer db.Close()
	// log.Println("Database connection successful!")

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"}, // Your frontend URL
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	// --- API Routes ---
	api := e.Group("/api")

	// Auth routes
	api.POST("/login", login)

	// User CRUD routes
	api.GET("/users", getUsers)
	api.POST("/users", createUser)
	api.GET("/users/:id", getUser)
	api.PUT("/users/:id", updateUser)
	api.DELETE("/users/:id", deleteUser)

	// Start server
	port := os.Getenv("BACKEND_PORT")
	if port == "" {
		port = "8080" // Default port if not specified
	}
	e.Logger.Fatal(e.Start(":" + port))
}

// --- Handler Functions (Placeholders) ---

func login(c echo.Context) error {
	// TODO: Implement actual login logic (check user/pass, create JWT)
	type LoginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	req := new(LoginRequest)
	if err := c.Bind(req); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	// For demonstration, we'll just check if email is "test@test.com"
	if req.Email == "test@test.com" && req.Password == "password" {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "Login successful!",
			"token":   "fake-jwt-token-for-demonstration",
		})
	}
	return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Invalid credentials"})
}

func getUsers(c echo.Context) error {
	// TODO: Fetch users from the database
	users := []User{
		{ID: 1, Name: "Alice", Email: "alice@example.com"},
		{ID: 2, Name: "Bob", Email: "bob@example.com"},
	}
	return c.JSON(http.StatusOK, users)
}

func createUser(c echo.Context) error {
	// TODO: Create a new user in the database
	u := new(User)
	if err := c.Bind(u); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	u.ID = 3 // Simulate new ID
	return c.JSON(http.StatusCreated, u)
}

func getUser(c echo.Context) error {
	// TODO: Fetch a single user from the database by ID
	id := c.Param("id")
	user := User{ID: 1, Name: "Alice", Email: "alice@example.com"}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Get user with ID: " + id,
		"user":    user,
	})
}

func updateUser(c echo.Context) error {
	// TODO: Update a user in the database
	id := c.Param("id")
	u := new(User)
	if err := c.Bind(u); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Updated user with ID: " + id,
		"user":    u,
	})
}

func deleteUser(c echo.Context) error {
	// TODO: Delete a user from the database
	id := c.Param("id")
	return c.JSON(http.StatusOK, map[string]string{"message": "Deleted user with ID: " + id})
}
