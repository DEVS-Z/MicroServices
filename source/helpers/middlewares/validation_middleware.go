package middleware

import (
	"errors"
	"net/http"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// DtoKey es la clave para almacenar el DTO en el contexto de Gin
const DtoKey = "dto"

// instancia global de validator
var validate = validator.New()

// ValidatorMiddleware crea un middleware Gin que:
// 1) Deserializa el JSON en un T genérico.
// 2) Ejecuta la validación estándar con go-playground/validator.
// 3) Ejecuta validaciones personalizadas.
// 4) Deposita el DTO validado en el contexto para uso en handlers.
func ValidatorMiddleware[T any]() gin.HandlerFunc {
	return func(c *gin.Context) {
		var dto T

		// 1. Bind JSON (decodifica y valida que sea JSON válido)
		if err := c.ShouldBindJSON(&dto); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON: " + err.Error()})
			c.Abort()
			return
		}

		// 2. Validación con tags `validate:"..."` en el struct
		if err := validate.Struct(dto); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Validation error: " + err.Error()})
			c.Abort()
			return
		}

		// 3. Validaciones personalizadas
		if err := runCustomValidator(&dto); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Custom validation error: " + err.Error()})
			c.Abort()
			return
		}

		// 4. Guardar el DTO en el contexto de Gin para que los handlers lo usen
		c.Set(DtoKey, dto)

		// Continuar con el siguiente handler
		c.Next()
	}
}

func runCustomValidator(dto interface{}) error {
	v := reflect.ValueOf(dto)
	t := reflect.TypeOf(dto)

	//Execute every methods to dto type
	for i := 0; i < t.NumMethod(); i++ {
		method := t.Method(i)

		if strings.HasPrefix(method.Name, "Valid") {
			result := v.MethodByName(method.Name).Call(nil)
			if len(result) != 1 {
				return errors.New("method" + method.Name + " should return a error")
			}
			if err, ok := result[0].Interface().(error); ok && err != nil {
				return err
			}
		}
	}
	return nil
}
