package main

import (
	"net/http"
	"time"

	"net/http/pprof"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// Booking contains binded and validated data.
type Booking struct {
	CheckIn  time.Time `form:"check_in" binding:"required,bookabledate" time_format:"2006-01-02"`
	CheckOut time.Time `form:"check_out" binding:"required,gtfield=CheckIn" time_format:"2006-01-02"`
}

// func bookableDate(
// 	v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
// 	field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string,
// ) bool {
// 	if date, ok := field.Interface().(time.Time); ok {
// 		today := time.Now()
// 		if today.Year() > date.Year() || today.YearDay() > date.YearDay() {
// 			return false
// 		}
// 	}
// 	return true
// }

func main() {
	route := gin.Default()
	// if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
	// 	v.RegisterValidation("bookabledate", bookableDate)
	// }
	route.GET("/index", pprofHandler(pprof.Index))
	route.GET("/block", pprofHandler(pprof.Handler("block").ServeHTTP))
	route.GET("/goroutine", pprofHandler(pprof.Handler("goroutine").ServeHTTP))
	route.GET("/heap", pprofHandler(pprof.Handler("heap").ServeHTTP))
	route.GET("/mutex", pprofHandler(pprof.Handler("mutex").ServeHTTP))
	route.GET("/threadcreate", pprofHandler(pprof.Handler("threadcreate").ServeHTTP))
	route.GET("/bookable", getBookable)
	route.Run(":8080")

}

func getBookable(c *gin.Context) {
	var b Booking
	if err := c.ShouldBindWith(&b, binding.Query); err == nil {
		c.JSON(http.StatusOK, gin.H{"message": "Booking dates are valid!"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func pprofHandler(h http.HandlerFunc) gin.HandlerFunc {
	handler := http.HandlerFunc(h)
	return func(c *gin.Context) {
		handler.ServeHTTP(c.Writer, c.Request)
	}

}
