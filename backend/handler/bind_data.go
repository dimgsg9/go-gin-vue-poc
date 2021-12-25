package handler

import (
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/dimgsg9/booker_proto/backend/model/apperrors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/validator/v10"

	ut "github.com/go-playground/universal-translator"

	enTranslations "github.com/go-playground/validator/v10/translations/en"
)

// bindData is helper function, returns false if data is not bound
func bindData(c *gin.Context, req interface{}) bool {

	var trans ut.Translator

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})

		en := en.New()
		uni := ut.New(en, en)
		trans, _ = uni.GetTranslator("en")
		enTranslations.RegisterDefaultTranslations(v, trans)

	}

	if c.ContentType() != "application/json" {
		msg := fmt.Sprintf("%s only accepts Content-Type application/json", c.FullPath())

		err := apperrors.NewUnsupportedMediaType(msg)

		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return false
	}

	// Bind incoming json to struct and check for validation errors
	if err := c.ShouldBind(req); err != nil {
		log.Printf("Error binding data: %+v\n", err)

		if verrs, ok := err.(validator.ValidationErrors); ok {
			invalidArgs := make(map[string]string)

			for _, verr := range verrs {
				// if verr.Param() != "" {
				// 	invalidArgs[verr.Field()] = verr.Translate(trans)
				// }
				invalidArgs[verr.Field()] = verr.Translate(trans)
			}

			err := apperrors.NewBadRequest("Invalid request parameters. See invalidArgs")

			c.JSON(err.Status(), gin.H{
				"error":       err,
				"invalidArgs": invalidArgs,
			})
			return false
		}

		// if we aren't able to properly extract validation errors,
		// we'll fallback and return an internal server error
		fallBack := apperrors.NewInternal()

		c.JSON(fallBack.Status(), gin.H{"error": fallBack})
		return false
	}

	return true
}
