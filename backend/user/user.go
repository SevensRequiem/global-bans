package user

import (
	"globalbans/backend/auth"
	"globalbans/backend/bans"
	"globalbans/backend/models"

	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo) {
	e.GET("/api/user", GetUserData)
}

func GetUserData(c echo.Context) error {
	// Get user data
	user, _ := auth.GetCurrentUser(c)
	if user == nil {
		user = &models.LoggedInUser{
			ID:       "0",
			Username: "Guest",
			Banned:   bans.BannedCheck(c),
		}
	} else {
		user = &models.LoggedInUser{
			ID:       user.ID,
			Username: user.Username,
			Banned:   bans.BannedCheck(c),
		}
	}

	// Return user data
	return c.JSON(200, user)
}
