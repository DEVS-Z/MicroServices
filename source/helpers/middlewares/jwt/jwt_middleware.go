package jwt_middleware

import (
	"errors"
	"main/core/security/token"

	"github.com/miqueaz/FrameGo/pkg/client"

	"github.com/gin-gonic/gin"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Obtener el token desde el encabezado Authorization
		t := c.GetHeader("Authorization")

		// Validar el token usando la función de tu paquete token
		tokenData, err := token.ValidToken(t)
		if err != nil {
			// Si el token no es válido, respondemos con un error Unauthorized
			client.Forbidden(c, errors.New("Unauthorized"))
			c.Abort() // Interrumpir el procesamiento de la solicitud
			return
		}

		c.Set("tokenData", tokenData.Claims)

		// Si el token es válido, continuamos con la siguiente función del middleware
		c.Next()
	}
}
