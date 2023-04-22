package utils

import (
	"Jobhun_Mahasiswa/src/models"
	"database/sql"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

func ExecuteSQLQueryRow(db *sql.DB, sqlStatement string, id string, result *models.Student, ctx *gin.Context) (*models.Student, error) {
	var hobby_id string
	var hobbies string
	err := db.QueryRow(sqlStatement).Scan(&result.ID, &result.Fullname, &result.Age, &result.Gender, &result.Regist_date, &result.Major, &result.MajorID, &hobby_id, &hobbies)
	result.HobbyID = strings.Split(hobby_id, ",")
	result.Hobbies = strings.Split(hobbies, ",")

	if err != nil {
		// panic(err.Error())
		BadRequest(err, ctx, "Failed to find student", err.Error())
		return nil, err
	}
	return result, nil
}

func AddValues(student_id string, sqlStatement string, column string, h_id string, i int, data []string) string {
	sqlStatement += fmt.Sprintf("('%s', '%v')", student_id, h_id)
	if i == (len(data) - 1) {
		sqlStatement += "; "
	} else {
		sqlStatement += ", "
	}
	return sqlStatement
}
