package controllers

import (
	"mvc/services"
	"mvc/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {

	//fmt.Println("sahi h ")
	//userIdParam := req.URL.Query().Get("user_id")
	//c.Query("caller")

	userId, err := strconv.ParseInt(c.Param("user_Id"), 10, 64)

	if err != nil {
		userErr := &utils.ApplicationError{
			Message:    "user id not found",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		utils.Respond(c, userErr.StatusCode, userErr)
		//c.JSON(userErr.StatusCode, userErr)

		return
		// jsonValue, _ := json.Marshal(userErr)
		// resp.WriteHeader(userErr.StatusCode)
		// resp.Write([]byte(jsonValue))

	}

	user, userErr := services.UsersService.GetUser(userId)

	if userErr != nil {
		utils.Respond(c, userErr.StatusCode, userErr)
		//c.JSON(userErr.StatusCode, userErr)

		return

		// jsonValue, _ := json.Marshal(userErr)
		// resp.WriteHeader(userErr.StatusCode)
		// resp.Write([]byte(jsonValue))

	}
	utils.Respond(c, http.StatusOK, user)
	c.JSON(http.StatusOK, user)
}
