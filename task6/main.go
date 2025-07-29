package main

import (
	"task_manager/data"
	"task_manager/router"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func main() {

    data.InitDB()
	r:= router.SetUpRouter()
		// ROUTING REGISTRATION 
		router.POST("/register", func( ctx *gin.Context){
			var user User
			if err := ctx.ShouldBindBodyWithJSON(&user) ; err != nil {
				ctx.JSON(400, gin.H{
					"message":"Invalid payload",
				})
				return 
			}
			// The registration logic
			
			// handle password encryption 
			hashedPassword,err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
			if err != nil {
				ctx.JSON(500, gin.H{"messsage":"Interal server errorrrrr",})
				return 
			}
	
			user.Password = string(hashedPassword)
			users[user.Email] = &user
	
	
			ctx.JSON(200, gin.H{"message": "user registered successfullly",})
		})
	
		// LOGIN LOGIC 
		router.POST("/login", func(ctx *gin.Context) {
			var user User
			if err := ctx.ShouldBindJSON(&user); err != nil {
				ctx.JSON(400, gin.H{"message": "Invalid payload"})
				return
			}
		
			existingUser, ok := users[user.Email]
			if !ok {
				ctx.JSON(401, gin.H{"error": "Wrong email or password- E"})
				return
			}
		
			err := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password))
			if err != nil {
				ctx.JSON(401, gin.H{"error": "Wrong email or password- P"})
				return     
			}
		
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"email": existingUser.Email, 
				"role":  existingUser.Role,  
			})
		
			tokenString, err := token.SignedString(jwtSecret)
			if err != nil {
				ctx.JSON(500, gin.H{"error": "Token generation failed"})
				return
			}
		
		
			ctx.JSON(200, gin.H{"message": "Login success", "token": tokenString})
		})
		




	r.Run(":8080")
}