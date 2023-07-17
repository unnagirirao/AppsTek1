package services

import (
	"github.com/unnagirirao/AppsTek1/appstek/pkg/rest/server/daos"
	"github.com/unnagirirao/AppsTek1/appstek/pkg/rest/server/models"
)

type AppService struct {
	appDao *daos.AppDao
}

func NewAppService() (*AppService, error) {
	appDao, err := daos.NewAppDao()
	if err != nil {
		return nil, err
	}
	return &AppService{
		appDao: appDao,
	}, nil
}

func (appService *AppService) CreateApp(app *models.App) (*models.App, error) {
	return appService.appDao.CreateApp(app)
}

func (appService *AppService) UpdateApp(id int64, app *models.App) (*models.App, error) {
	return appService.appDao.UpdateApp(id, app)
}

func (appService *AppService) DeleteApp(id int64) error {
	return appService.appDao.DeleteApp(id)
}

func (appService *AppService) ListApps() ([]*models.App, error) {
	return appService.appDao.ListApps()
}

func (appService *AppService) GetApp(id int64) (*models.App, error) {
	return appService.appDao.GetApp(id)
}
