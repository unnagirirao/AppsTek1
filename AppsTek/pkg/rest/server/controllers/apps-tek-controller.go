package controllers

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/unnagirirao/AppsTek1/appstek/pkg/rest/server/models"
	"github.com/unnagirirao/AppsTek1/appstek/pkg/rest/server/services"
	"net/http"
	"strconv"
)

type AppsTekController struct {
	appsTekService *services.AppsTekService
}

func NewAppsTekController() (*AppsTekController, error) {
	appsTekService, err := services.NewAppsTekService()
	if err != nil {
		return nil, err
	}
	return &AppsTekController{
		appsTekService: appsTekService,
	}, nil
}

func (appsTekController *AppsTekController) CreateAppsTek(context *gin.Context) {
	// validate input
	var input models.AppsTek
	if err := context.ShouldBindJSON(&input); err != nil {
		log.Error(err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// trigger appsTek creation
	if _, err := appsTekController.appsTekService.CreateAppsTek(&input); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "AppsTek created successfully"})
}

func (appsTekController *AppsTekController) UpdateAppsTek(context *gin.Context) {
	// validate input
	var input models.AppsTek
	if err := context.ShouldBindJSON(&input); err != nil {
		log.Error(err)
		context.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// trigger appsTek update
	if _, err := appsTekController.appsTekService.UpdateAppsTek(id, &input); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "AppsTek updated successfully"})
}

func (appsTekController *AppsTekController) FetchAppsTek(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// trigger appsTek fetching
	appsTek, err := appsTekController.appsTekService.GetAppsTek(id)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, appsTek)
}

func (appsTekController *AppsTekController) DeleteAppsTek(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// trigger appsTek deletion
	if err := appsTekController.appsTekService.DeleteAppsTek(id); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "AppsTek deleted successfully",
	})
}

func (appsTekController *AppsTekController) ListAppsTeks(context *gin.Context) {
	// trigger all appsTeks fetching
	appsTeks, err := appsTekController.appsTekService.ListAppsTeks()
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, appsTeks)
}

func (*AppsTekController) PatchAppsTek(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "PATCH",
	})
}

func (*AppsTekController) OptionsAppsTek(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "OPTIONS",
	})
}

func (*AppsTekController) HeadAppsTek(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "HEAD",
	})
}
