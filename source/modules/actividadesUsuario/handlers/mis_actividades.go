package handlers

import (
	"errors"
	"main/pkg/client"
	util "main/source/modules/actividadesUsuario/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func MyActividadesByEstado(c *gin.Context) {
	var filters = make(map[string]interface{})

	claimsAny, exists := c.Get("tokenData")
	if !exists {
		client.Forbidden(c, errors.New("User role not found"))
		c.Abort()
		return
	}

	tokenData, ok := claimsAny.(jwt.MapClaims)
	if !ok {
		client.Forbidden(c, errors.New("Invalid token data type"))
		c.Abort()
		return
	}
	//imprimir la data del token
	filters["usuario"] = tokenData["id"]
	data, err := util.ActividadesByEstado(filters)
	if err != nil {
		client.Error(c, "Error fetching activities by state", err)
		c.Abort()
		return
	}

	if data == nil {
		client.NotFound(c, errors.New("No activities found for the user"))
		c.Abort()
		return
	}
	client.Success(c, "Actividades por estado", data)
}
