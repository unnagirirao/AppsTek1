package daos

import (
	"database/sql"
	"errors"
	log "github.com/sirupsen/logrus"
	"github.com/unnagirirao/AppsTek1/appstek/pkg/rest/server/daos/clients/sqls"
	"github.com/unnagirirao/AppsTek1/appstek/pkg/rest/server/models"
)

type AppDao struct {
	sqlClient *sqls.SQLiteClient
}

func migrateApps(r *sqls.SQLiteClient) error {
	query := `
	CREATE TABLE IF NOT EXISTS apps(
		Id INTEGER PRIMARY KEY AUTOINCREMENT,
        
		Trainees TEXT NOT NULL,
		Employees TEXT NOT NULL,
        CONSTRAINT id_unique_key UNIQUE (Id)
	)
	`
	_, err1 := r.DB.Exec(query)
	return err1
}

func NewAppDao() (*AppDao, error) {
	sqlClient, err := sqls.InitSqliteDB()
	if err != nil {
		return nil, err
	}
	err = migrateApps(sqlClient)
	if err != nil {
		return nil, err
	}
	return &AppDao{
		sqlClient,
	}, nil
}

func (appDao *AppDao) CreateApp(m *models.App) (*models.App, error) {
	insertQuery := "INSERT INTO apps(Trainees, Employees)values(?, ?)"
	res, err := appDao.sqlClient.DB.Exec(insertQuery, m.Trainees, m.Employees)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	m.Id = id

	log.Debugf("app created")
	return m, nil
}

func (appDao *AppDao) UpdateApp(id int64, m *models.App) (*models.App, error) {
	if id == 0 {
		return nil, errors.New("invalid updated ID")
	}
	if id != m.Id {
		return nil, errors.New("id and payload don't match")
	}

	app, err := appDao.GetApp(id)
	if err != nil {
		return nil, err
	}
	if app == nil {
		return nil, sql.ErrNoRows
	}

	updateQuery := "UPDATE apps SET Trainees = ?, Employees = ? WHERE Id = ?"
	res, err := appDao.sqlClient.DB.Exec(updateQuery, m.Trainees, m.Employees, id)
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

	log.Debugf("app updated")
	return m, nil
}

func (appDao *AppDao) DeleteApp(id int64) error {
	deleteQuery := "DELETE FROM apps WHERE Id = ?"
	res, err := appDao.sqlClient.DB.Exec(deleteQuery, id)
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

	log.Debugf("app deleted")
	return nil
}

func (appDao *AppDao) ListApps() ([]*models.App, error) {
	selectQuery := "SELECT * FROM apps"
	rows, err := appDao.sqlClient.DB.Query(selectQuery)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)
	var apps []*models.App
	for rows.Next() {
		m := models.App{}
		if err = rows.Scan(&m.Id, &m.Trainees, &m.Employees); err != nil {
			return nil, err
		}
		apps = append(apps, &m)
	}
	if apps == nil {
		apps = []*models.App{}
	}

	log.Debugf("app listed")
	return apps, nil
}

func (appDao *AppDao) GetApp(id int64) (*models.App, error) {
	selectQuery := "SELECT * FROM apps WHERE Id = ?"
	row := appDao.sqlClient.DB.QueryRow(selectQuery, id)
	m := models.App{}
	if err := row.Scan(&m.Id, &m.Trainees, &m.Employees); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sqls.ErrNotExists
		}
		return nil, err
	}

	log.Debugf("app retrieved")
	return &m, nil
}
