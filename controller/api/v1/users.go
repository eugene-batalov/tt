package v1

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"

	"tt/db"
	"tt/model"
)

// CreateUser
// @Summary about CreateUser
// @Tags Users
// @ID CreateUser
// @Param request body model.User true "Create user"
// @Success 201 {object} int "User id"
// @Failure 400 {object} string "Bad Request"
// @Router /users [post]
func CreateUser(c echo.Context) (err error) {
	u := new(model.User)
	if err = c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err = c.Validate(u); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err = db.CreateUser(u); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusCreated, u.ID)
}

// UpdateUser
// @Summary about UpdateUser
// @Tags Users
// @ID UpdateUser
// @Param request body model.User true "Update user"
// @Success 200 {object} string "ok"
// @Failure 400 {object} string "Bad Request"
// @Router /users [put]
func UpdateUser(c echo.Context) (err error) {
	u := new(model.User)
	if err = c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err = db.UpdateUser(u); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, "ok")
}

// DeleteUser
// @Summary about DeleteUser
// @Tags Users
// @ID DeleteUser
// @Param id query int true "Delete user"
// @Success 204
// @Failure 400 {object} string "Bad Request"
// @Router /users [delete]
func DeleteUser(c echo.Context) (err error) {
	var id int64
	if id, err = strconv.ParseInt(c.QueryParam("id"), 10, 64); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	u := &model.User{ID: id}
	if err = db.DeleteUser(u); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}

// SearchUsersOrdered
// @Summary about SearchUsersOrdered
// @Tags Users
// @ID SearchUsersOrdered
// @Param email query string false "user email substring"
// @Param phone query string false "user phone substring"
// @Param orderBy query string false "order by field name"
// @Success 200 {array} model.User "list of users"
// @Failure 400 {object} string "Bad Request"
// @Router /users [get]
func SearchUsersOrdered(c echo.Context) (err error) {
	emailSubstr := "%" + c.QueryParam("email") + "%"
	phoneSubstr := "%" + c.QueryParam("phone") + "%"
	orderBy := c.QueryParam("orderBy")

	var users []model.User
	if users, err = db.SearchUsersOrdered(emailSubstr, phoneSubstr, orderBy); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, users)
}
