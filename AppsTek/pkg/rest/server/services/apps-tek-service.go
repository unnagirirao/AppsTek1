package services

import (
	"github.com/unnagirirao/AppsTek1/appstek/pkg/rest/server/daos"
	"github.com/unnagirirao/AppsTek1/appstek/pkg/rest/server/models"
)

type AppsTekService struct {
	appsTekDao *daos.AppsTekDao
}

func NewAppsTekService() (*AppsTekService, error) {
	appsTekDao, err := daos.NewAppsTekDao()
	if err != nil {
		return nil, err
	}
	return &AppsTekService{
		appsTekDao: appsTekDao,
	}, nil
}

func (appsTekService *AppsTekService) CreateAppsTek(appsTek *models.AppsTek) (*models.AppsTek, error) {
	return appsTekService.appsTekDao.CreateAppsTek(appsTek)
}

func (appsTekService *AppsTekService) UpdateAppsTek(id int64, appsTek *models.AppsTek) (*models.AppsTek, error) {
	return appsTekService.appsTekDao.UpdateAppsTek(id, appsTek)
}

func (appsTekService *AppsTekService) DeleteAppsTek(id int64) error {
	return appsTekService.appsTekDao.DeleteAppsTek(id)
}

func (appsTekService *AppsTekService) ListAppsTeks() ([]*models.AppsTek, error) {
	return appsTekService.appsTekDao.ListAppsTeks()
}

func (appsTekService *AppsTekService) GetAppsTek(id int64) (*models.AppsTek, error) {
	return appsTekService.appsTekDao.GetAppsTek(id)
}
