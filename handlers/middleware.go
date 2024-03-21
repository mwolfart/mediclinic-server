package handlers

import (
	"medclin/models"
	"medclin/utils"

	"github.com/gin-gonic/gin"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func HasPermission(c *gin.Context, roles ...string) (bool, error) {
	userId, err := utils.UserEmailFromToken(c)
	if err != nil {
		return false, err
	}

	user, err := models.Genericusers(qm.Where("id = ?", userId), qm.Load(models.GenericuserRels.RoleidRoles)).One(c, db)
	if err != nil {
		return false, err
	}

	userRoles := user.R.RoleidRoles
	for _, roleString := range roles {
		for _, userRole := range userRoles {
			if userRole.Description == roleString {
				return true, nil
			}
		}
	}

	return false, nil
}
