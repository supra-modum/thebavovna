package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func signUp(c *gin.Context) {
	//user := new(store.User)
	//if err := ctx.Bind(user); err != nil {
	//	ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//	return
	//}
	//if err := store.AddUser(user); err != nil {
	//	ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//	return
	//}
	c.JSON(http.StatusOK, gin.H{
		"msg": "Signed up successfully.",
		"jwt": "123456789",
	})
}

func signIn(c *gin.Context) {
	//user := new(store.User)
	//if err := ctx.Bind(user); err != nil {
	//	ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//	return
	//}
	//user, err := store.Authenticate(user.Username, user.Password)
	//if err != nil {
	//	ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Sign in failed."})
	//	return
	//}

	c.JSON(http.StatusOK, gin.H{
		"msg": "Signed in successfully.",
		"jwt": "123456789",
	})
}
