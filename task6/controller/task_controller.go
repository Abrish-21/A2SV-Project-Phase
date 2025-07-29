package controller

import (
	"net/http"
	"task_manager/data"
	"task_manager/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = "for testing only"

// ROUTING REGISTRATION
	func UserRegisterController( ctx *gin.Context){
		var user models.User
		// bind the payload to the user struct 
		if err := ctx.ShouldBindBodyWithJSON(&user) ; err != nil {
			ctx.JSON(400, gin.H{
				"message":"Invalid payload",
			})
			return 
		}
		// check if the user is in the DB using the FInderUserService 
		if data.FindUserService(user.Username) {
			ctx.JSON(http.StatusConflict, gin.H{"error":"username already exisits"})
			return 
		}
		
		// handle password encryption 
		hashedPassword,err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error" : "error hashing using bcrypt"})
			return 
		}
		err = data.UserRegisterService(user, string(hashedPassword))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error":"can't insert into the database"})
			return 
		}

		ctx.JSON(200, gin.H{"message": "user registered successfullly",})
	}

	func GetUserProfileController(c *gin.Context) {
		c.JSON(200, gin.H{"message" : "Only logged in users can see this"})
	}
	
	func GetAdminPageController(c *gin.Context) {
		c.JSON(200, gin.H{"message" : "Hello, welcome to the admin page!"})
	}

	// LOGIN LOGIC 
	func UserLoginController(ctx *gin.Context) {
		var user models.User
		if err := ctx.ShouldBindJSON(&user); err != nil {
			ctx.JSON(400, gin.H{"message": "Invalid payload"})
			return
		}
	
		existingUser, err := data.UserLoginService(user)
		if err !=  nil {
			ctx.JSON(401, gin.H{"error": "Wrong email or password- E"})
			return
		}
	//   check if if the password is correct or wrong  
		err = bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password))
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
	}


// CONTROLLER FOR GETTING ALL TASKS
func GetAllTasks(ctx *gin.Context) {
	tasks,err := data.GetAllTasks()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tasks"})
		return 
	}
	ctx.JSONP(http.StatusOK, tasks)
}
// GETTING  A TASK BY ID 
func GetTaskByID(ctx *gin.Context){
	id:= ctx.Param("id")
	task, err := data.GetTaskByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound,gin.H{ "Error": " Task not Found"})
		return 
	}
	ctx.JSON(http.StatusOK,task )
}

// CREATING A NEW TASK 
func AddNewTask(ctx *gin.Context) {
	var newTask  models.Task
	if err:= ctx.ShouldBindJSON(&newTask); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "error": err.Error()})
		return 
		
	}
	err:= data.AddNewTask(newTask)
	if  err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H {"error":"can't create task"})
		return 
	}
	ctx.JSON(http.StatusCreated, gin.H{ "Message:": "Task created"})

}

// UPDATING TASK 
func UpdateTask(ctx *gin.Context) {
	id := ctx.Param("id")
	var updatedTask models.Task

	if err := ctx.ShouldBindJSON(&updatedTask); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := data.UpdateTask(id, updatedTask)
	if err != nil {
		if err == mongo.ErrNoDocuments{
			ctx.JSON(http.StatusNotFound, gin.H{"message":"task can't be found"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message":"Failed to update task"})
		}
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "task successfully updated"})
}

// DELETING TASK 
func DeleteTask(ctx *gin.Context) {
	id := ctx.Param("id")

	err:= data.DeleteTask(id) 
	if err != nil{
		if err == mongo.ErrNoDocuments{
			ctx.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})

		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Task can't be deleted"})
		}
		return 
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
}

