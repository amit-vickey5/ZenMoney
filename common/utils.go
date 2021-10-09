package common

import (
	"fmt"
	"io"
	"net/http"
	"reflect"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

var logger io.Writer

func InitLogger() {
	logger = gin.DefaultErrorWriter
}

func LogInfo(ctx *gin.Context, info string) {
	logger.Write([]byte(info + "\n"))
}

func GetSetuAACustomerHandle(customerNumber string) string {
	return fmt.Sprintf("%s@%s", customerNumber, SetuHandle)
}

func StructToMap(item interface{}) map[string]interface{} {
	res := map[string]interface{}{}
	if item == nil {
		return res
	}
	v := reflect.TypeOf(item)
	reflectValue := reflect.ValueOf(item)
	reflectValue = reflect.Indirect(reflectValue)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	for i := 0; i < v.NumField(); i++ {
		tag := v.Field(i).Tag.Get("json")
		field := reflectValue.Field(i).Interface()
		if tag != "" && tag != "-" {
			if v.Field(i).Type.Kind() == reflect.Struct {
				res[tag] = StructToMap(field)
			} else {
				res[tag] = field
			}
		}
	}
	return res
}

func GetTimestampFromISO(isoTimeStr string) (int64, error) {
	layout := "2006-01-02T15:04:05.000Z"
	t, err := time.Parse(layout, isoTimeStr)
	if err != nil {
		return 0, errors.Wrap(err, "error-parsing-iso-time")
	}
	return t.Unix(), nil
}

func ErrorResponse(ctx *gin.Context, err error) {
	LogInfo(ctx, fmt.Sprintf("error-occurred: %v", err.Error()))
	ctx.JSON(http.StatusInternalServerError, gin.H{
		"errror": errors.WithStack(err).Error(),
	})
}
