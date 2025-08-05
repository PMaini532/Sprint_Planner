package routes
import (
	"github.com/gin-gonic/gin"
	"github.com/PMaini532/Sprint_Planner/controllers"
)

func UserRoutes(r *gin.Engine) {
	users := r.Group("/users")
	{
		users.GET("", controllers.GetAllUsers)
		users.GET("/admin", controllers.GetAdmins)
		users.GET("/developer", controllers.GetDevelopers)
	}
}