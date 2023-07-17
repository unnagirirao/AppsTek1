package daos

import (
	"database/sql"
	"errors"
	log "github.com/sirupsen/logrus"
	"github.com/unnagirirao/AppsTek1/appstek/pkg/rest/server/daos/clients/sqls"
	"github.com/unnagirirao/AppsTek1/appstek/pkg/rest/server/models"
)

type AppsTekDao struct {
	sqlClient *sqls.SQLiteClient
}

func migrateAppsTeks(r *sqls.SQLiteClient) error {
	query := `
	CREATE TABLE IF NOT EXISTS appsTeks(
		Id INTEGER PRIMARY KEY AUTOINCREMENT,
        
		Employees TEXT NOT NULL,
		Trainees TEXT NOT NULL,
        CONSTRAINT id_unique_key UNIQUE (Id)
	)
	`
	_, err1 := r.DB.Exec(query)
	return err1
}

func NewAppsTekDao() (*AppsTekDao, error) {
	sqlClient, err := sqls.InitSqliteDB()
	if err != nil {
		return nil, err
	}
	err = migrateAppsTeks(sqlClient)
	if err != nil {
		return nil, err
	}
	return &AppsTekDao{
		sqlClient,
	}, nil
}

func (appsTekDao *AppsTekDao) CreateAppsTek(m *models.AppsTek) (*models.AppsTek, error) {
	insertQuery := "INSERT INTO appsTeks(Employees, Trainees)values(?, ?)"
	res, err := appsTekDao.sqlClient.DB.Exec(insertQuery, m.Employees, m.Trainees)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	m.Id = id

	log.Debugf("appsTek created")
	return m, nil
}

func (appsTekDao *AppsTekDao) UpdateAppsTek(id int64, m *models.AppsTek) (*models.AppsTek, error) {
	if id == 0 {
		return nil, errors.New("invalid updated ID")
	}
	if id != m.Id {
		return nil, errors.New("id and payload don't match")
	}

	appsTek, err := appsTekDao.GetAppsTek(id)
	if err != nil {
		return nil, err
	}
	if appsTek == nil {
		return nil, sql.ErrNoRows
	}

	updateQuery := "UPDATE appsTeks SET Employees = ?, Trainees = ? WHERE Id = ?"
	res, err := appsTekDao.sqlClient.DB.Exec(updateQuery, m.Employees, m.Trainees, id)
	if err != nil {
		return nil, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}
	if rowsAffected == 0 {
		return nil, sqls.ErrUpdateFailed
	}

	log.Debugf("appsTek updated")
	return m, nil
}

func (appsTekDao *AppsTekDao) DeleteAppsTek(id int64) error {
	deleteQuery := "DELETE FROM appsTeks WHERE Id = ?"
	res, err := appsTekDao.sqlClient.DB.Exec(deleteQuery, id)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return sqls.ErrDeleteFailed
	}

	log.Debugf("appsTek deleted")
	return nil
}

func (appsTekDao *AppsTekDao) ListAppsTeks() ([]*models.AppsTek, error) {
	selectQuery := "SELECT * FROM appsTeks"
	rows, err := appsTekDao.sqlClient.DB.Query(selectQuery)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)
	var appsTeks []*models.AppsTek
	for rows.Next() {
		m := models.AppsTek{}
		if err = rows.Scan(&m.Id, &m.Employees, &m.Trainees); err != nil {
			return nil, err
		}
		appsTeks = append(appsTeks, &m)
	}
	if appsTeks == nil {
		appsTeks = []*models.AppsTek{}
	}

	log.Debugf("appsTek listed")
	return appsTeks, nil
}

func (appsTekDao *AppsTekDao) GetAppsTek(id int64) (*models.AppsTek, error) {
	selectQuery := "SELECT * FROM appsTeks WHERE Id = ?"
	row := appsTekDao.sqlClient.DB.QueryRow(selectQuery, id)
	m := models.AppsTek{}
	if err := row.Scan(&m.Id, &m.Employees, &m.Trainees); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sqls.ErrNotExists
		}
		return nil, err
	}

	log.Debugf("appsTek retrieved")
	return &m, nil
}
