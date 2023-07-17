package controllers

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/unnagirirao/AppsTek1/appstek/pkg/rest/server/models"
	"github.com/unnagirirao/AppsTek1/appstek/pkg/rest/server/services"
	"net/http"
	"strconv"
)

type AppController struct {
	appService *services.AppService
}

func NewAppController() (*AppController, error) {
	appService, err := services.NewAppService()
	if err != nil {
		return nil, err
	}
	return &AppController{
		appService: appService,
	}, nil
}

func (appController *AppController) CreateApp(context *gin.Context) {
	// validate input
	var input models.App
	if err := context.ShouldBindJSON(&input); err != nil {
		log.Error(err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// trigger app creation
	if _, err := appController.appService.CreateApp(&input); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "App created successfully"})
}

func (appController *AppController) UpdateApp(context *gin.Context) {
	// validate input
	var input models.App
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

	// trigger app update
	if _, err := appController.appService.UpdateApp(id, &input); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "App updated successfully"})
}

func (appController *AppController) FetchApp(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// trigger app fetching
	app, err := appController.appService.GetApp(id)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, app)
}

func (appController *AppController) DeleteApp(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// trigger app deletion
	if err := appController.appService.DeleteApp(id); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "App deleted successfully",
	})
}

func (appController *AppController) ListApps(context *gin.Context) {
	// trigger all apps fetching
	apps, err := appController.appService.ListApps()
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, apps)
}

func (*AppController) PatchApp(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "PATCH",
	})
}

func (*AppController) OptionsApp(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "OPTIONS",
	})
}

func (*AppController) HeadApp(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "HEAD",
	})
}
