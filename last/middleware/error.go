package middleware

import (
	"fmt"
	// "log"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

const InternalServerError = "Something went wrong!"

func CustomErrors(ctx *gin.Context) {
	ctx.Next()
	if len(ctx.Errors) > 0 {
		for _, err := range ctx.Errors {
			// Check error type
			switch err.Type {
			case gin.ErrorTypePublic:
				// Show public errors only if nothing has been written yet
				if !ctx.Writer.Written() {
					ctx.AbortWithStatusJSON(ctx.Writer.Status(), gin.H{"error": err.Error()})
				}
			case gin.ErrorTypeBind:
				errMap := make(map[string]string)
				if errs, ok := err.Err.(validator.ValidationErrors); ok {
					for _, fieldErr := range []validator.FieldError(errs) {
						errMap[fieldErr.Field()] = customValidationError(fieldErr)
					}
				}

				status := http.StatusBadRequest
				// Preserve current status
				if ctx.Writer.Status() != http.StatusOK {
					status = ctx.Writer.Status()
				}
				ctx.AbortWithStatusJSON(status, gin.H{"error": errMap})
			default:
				// Log other errors

				// log.Fatal().Err(err.Err).Msg("Other error")
			}
		}

		// If there was no public or bind error, display default 500 message
		if !ctx.Writer.Written() {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": InternalServerError})
		}
	}
}

func customValidationError(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return fmt.Sprintf("%s is required.", err.Field())
	case "min":
		return fmt.Sprintf("%s must be longer than or equal %s characters.", err.Field(), err.Param())
	case "max":
		return fmt.Sprintf("%s cannot be longer than %s characters.", err.Field(), err.Param())
	case "email":
		return fmt.Sprintf("%s is invalid %s characters.", err.Field(), err.Param())
	default:
		return err.Error()
	}
}
