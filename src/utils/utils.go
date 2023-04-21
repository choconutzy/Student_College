package utils

import (
	"Jobhun_Mahasiswa/src/models"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func ExecuteSQLQueryRow(db *sql.DB, sqlStatement string, id string, result *models.StudentDB, ctx *gin.Context) (*models.StudentDB, error) {
	err := db.QueryRow(sqlStatement, id).Scan(&result.ID, &result.Fullname, &result.Age, &result.Gender, &result.Regist_date, &result.Major, &result.Hobby)

	if err != nil {
		BadRequest(err, ctx, "Failed to find student", "data not found")
		return nil, err
	}
	return result, nil
}
