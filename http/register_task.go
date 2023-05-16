package http

import "github.com/gin-gonic/gin"

//github.com/robertobouses/

type RegisterTaskRequest struct {
	ID          int    `json:"id" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Description string `description:"lastName" binding:"required"`
	DueDate     string `duedate:"password" binding:"required"`
	Completed   bool   `completed:"password" binding:"required"`
}

func (h Handler) RegisterTask(c *gin.Context) {

	//h.app.Create_task
}

/*var req RegisterTaskRequest
	if err := c.BindJSON(&req); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return

	}

	params := user.CreateTaskParams(req)
	usr, err := h.app.CreateUser(params)
	if errors.Is(err, user.ValidationError) {
		c.AbortWithStatus(http.StatusBadRequest)
		fmt.Println(usr)
		fmt.Print(err)
		return
	}
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, RegisterResponse{
		ID:        usr.ID.String(),
		Email:     usr.Email,
		FirstName: usr.FirstName,
		LastName:  usr.LastName,
		Password:  usr.Password,
	},
	)
	fmt.Println("c.JSON password_____________", usr.Password)
}	*/
