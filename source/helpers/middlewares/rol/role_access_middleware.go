package role_middleware

import (
	"errors"
	"log"
	modulos_rol_service "main/source/modules/modulosRol/service"

	base_models "github.com/miqueaz/FrameGo/pkg/base/models"
	"github.com/miqueaz/FrameGo/pkg/client"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RoleAccessMiddleware[T any](modulo base_models.Model[T]) gin.HandlerFunc {
	return func(c *gin.Context) {

		// Obtener el tokenData (claims del JWT) desde el contexto
		claimsAny, exists := c.Get("tokenData") // Asegúrate que estás usando "tokenClaims" si lo guardaste así
		if !exists {
			client.Forbidden(c, errors.New("User role not found"))
			c.Abort()
			return
		}

		// Realizar type assertion a jwt.MapClaims
		tokenData, ok := claimsAny.(jwt.MapClaims)
		if !ok {
			client.Forbidden(c, errors.New("Invalid token data type"))
			c.Abort()
			return
		}

		// Obtener el rol del usuario
		role, ok := tokenData["role"].(float64)
		if !ok {
			client.Forbidden(c, errors.New("User role is missing or invalid"))
			c.Abort()
			return
		}

		// Debug opcional
		log.Printf("Modulo ID: %v", modulo.ID)
		log.Printf("User Role: %v", int(role))

		// Verificar permisos en base al módulo y rol
		haveAccess, err := modulos_rol_service.Service.Service.Read(map[string]any{
			"Modulo": modulo.ID,
			"Rol":    int(role),
		})
		if err != nil {
			client.Forbidden(c, errors.New("Not authorized to access this module"))
			c.Abort()
			return
		}

		if len(haveAccess) > 0 {
			c.Next() // Rol permitido, continuar
			return
		}

		// Rol no tiene permiso
		client.Forbidden(c, errors.New("Not authorized to access this module"))
		c.Abort()
	}
}
