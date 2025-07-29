package Middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthRoleMiddleWare() gin.HandlerFunc  {
	// check for exist, string, admin 

	 return func (ctx *gin.Context) {
		role, exists := ctx.Get("role")
			if !exists{
				ctx.AbortWithStatusJSON(403, gin.H{"error":"role not found", "message":"you have to entera role"})
				return 
			}
			roleString, ok := role.(string)
			if !ok {
				ctx.AbortWithStatusJSON(403, gin.H{"error":"invalid role format", "message":"you have to enter a valid role format"})	
				return 
			}

			if roleString != "admin" {

				ctx.AbortWithStatusJSON(403, gin.H{"error":"unauthorized role", "message":"only admin can access this route"})
				return 
			}

			ctx.Next()

		}
	 }




	var jwtSecret = [] byte("My_Jwt_Secret")

func AuthMiddleWare() gin.HandlerFunc {       
	// THE STEP  

	return func(ctx *gin.Context) {
		// IMPLEMENTATION LOGIC IS KEPT HERE 
		// STEP1: check for presenece of the authorization header 
		authHeader := ctx.GetHeader("authorization")
		if authHeader == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message":"Unauthorized token", "error":"Missing token"})
			return

		}
		authParts := strings.Split(authHeader, " ")
		// STPE2: check if the bearer is there 
		if len(authParts) != 2 || strings.ToLower(authParts[0]) != "bearer" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message":"Unauthorized auth header format", "error":"invalid authoirzation header"})
			return 
		}

		// STEP3: parse the header
		token, err := jwt.Parse(authParts[1], func (token *jwt.Token) (interface{}, error){
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("invalid signing method %v", token.Header["alg"])
			}
			return jwtSecret, nil
		})

				// Check if the JWT is valid and has the type MapClaims 
				if claims, ok := token.Claims.(jwt.MapClaims); err == nil && ok && token.Valid {
					// Get role and store it for the next handlers to authorize role
					ctx.Set("role", claims["role"].(string))
				} else {
					ctx.JSON(401, gin.H{"error" : "Invalid JWT"})
					ctx.Abort()
					return 
				}
				ctx.Next()
		//   The last step is protecting the routes  





		
	}
	// stepe 1: EXTRACT AUTH HEADER 
	

}