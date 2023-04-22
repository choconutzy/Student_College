package controllers

import (
	"Jobhun_Mahasiswa/src/config"
	"Jobhun_Mahasiswa/src/models"
	"Jobhun_Mahasiswa/src/utils"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var db = config.DB()

func GetMajors(ctx *gin.Context) {
	var results = []models.Majors{}

	sqlStatement := `SELECT * FROM majors;`

	rows, err := db.Query(sqlStatement)

	if err != nil {
		utils.InternalServerErrorResponse(err, ctx, err.Error())
		return
	}

	for rows.Next() {

		var major models.Majors
		err = rows.Scan(&major.ID, &major.Major)

		if err != nil {
			utils.InternalServerErrorResponse(err, ctx, err.Error())
			return
		}

		results = append(results, major)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":    results,
		"message": "success get all majors data",
	})
}

func GetHobbies(ctx *gin.Context) {
	var results = []models.Hobbies{}

	sqlStatement := `SELECT * FROM hobbies;`

	rows, err := db.Query(sqlStatement)

	if err != nil {
		utils.InternalServerErrorResponse(err, ctx, err.Error())
		return
	}

	for rows.Next() {

		var hobby models.Hobbies
		err = rows.Scan(&hobby.ID, &hobby.Hobby)

		if err != nil {
			utils.InternalServerErrorResponse(err, ctx, err.Error())
			return
		}

		results = append(results, hobby)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":    results,
		"message": "success get all hobbies data",
	})
}

func GetStudents(ctx *gin.Context) {
	var results = []models.Student{}

	sqlStatement := `SELECT st.id, st.fullname, st.age, st.gender, st.regist_date, sm.major_id, m.major, GROUP_CONCAT(sh.hobby_id) hobby_id, GROUP_CONCAT(h.hobby) hobbies FROM students st JOIN student_hobby sh ON st.id = sh.student_id JOIN hobbies h ON sh.hobby_id = h.id JOIN student_major sm ON sm.student_id = st.id JOIN majors m ON sm.major_id = m.id GROUP BY st.id;`

	rows, err := db.Query(sqlStatement)

	if err != nil {
		// panic(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
		fmt.Print(err)
		return
	}

	for rows.Next() {

		var student_ models.Student
		var hobby_id string
		var hobbies string

		err = rows.Scan(&student_.ID, &student_.Fullname, &student_.Age, &student_.Gender, &student_.Regist_date, &student_.MajorID, &student_.Major, &hobby_id, &hobbies)

		if err != nil {
			// panic(err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "internal server error",
			})
			fmt.Print(err)
			return
		}

		student_.HobbyID = strings.Split(hobby_id, ",")
		student_.Hobbies = strings.Split(hobbies, ",")

		results = append(results, student_)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":    results,
		"message": "success get all student data",
	})
}

func GetStudentById(ctx *gin.Context) {
	id := ctx.Param("id")

	var student_data models.Student

	sqlStatement := fmt.Sprintf(`SELECT st.id, st.fullname, st.age, st.gender, st.regist_date, sm.major_id, m.major, GROUP_CONCAT(sh.hobby_id) hobby_id, GROUP_CONCAT(h.hobby) hobbies FROM students st JOIN student_hobby sh ON st.id = sh.student_id JOIN hobbies h ON sh.hobby_id = h.id JOIN student_major sm ON sm.student_id = st.id JOIN majors m ON sm.major_id = m.id WHERE st.id = '%s' GROUP BY st.id;`, id)

	result, err := utils.ExecuteSQLQueryRow(db, sqlStatement, id, &student_data, ctx)
	if err == nil {
		fmt.Println(result)
		ctx.JSON(http.StatusOK, gin.H{
			"data":    result,
			"message": fmt.Sprint("success get student data by id", id),
		})
	}

}

func CreateStudent(ctx *gin.Context) {
	student_id := uuid.New().String()
	var student models.Student

	if err := ctx.ShouldBindJSON(&student); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	student.ID = student_id

	queries := []string{
		fmt.Sprintf("INSERT INTO students (id,fullname, gender, age, regist_date) VALUES ('%s','%s','%s',%d,'%s');", student_id, student.Fullname, student.Gender, student.Age, student.Regist_date),
		fmt.Sprintf("INSERT INTO student_major (student_id, major_id) VALUES ('%s','%s');", student_id, student.MajorID),
		"INSERT INTO student_hobby (student_id, hobby_id) VALUES ",
	}
	for i, query := range queries {
		column := "major"
		data := student.HobbyID

		if i == 2 {
			for h := 0; h < len(data); h++ {
				query = utils.AddValues(student_id, query, column, data[h], h, data)
			}
		}
		_, err := db.Exec(query)
		if err != nil {
			// panic(err.Error())
			utils.InternalServerErrorResponse(err, ctx, fmt.Sprintf("Failed to add %s data", strings.Split(query, " ")[2]))
			return
		}
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "student has been successfully created",
	})
}

func UpdateStudent(ctx *gin.Context) {
	student_id := ctx.Param("id")

	var student models.Student
	if err := ctx.ShouldBindJSON(&student); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	queries := []string{
		fmt.Sprintf("UPDATE students SET fullname = '%s', age = %d, gender = '%s', regist_date = '%s' WHERE id = '%s';", student.Fullname, student.Age, student.Gender, student.Regist_date, student_id),
		fmt.Sprintf("UPDATE student_major SET major_id = '%s' WHERE student_id = '%s';", student.MajorID, student_id),
		fmt.Sprintf("DELETE FROM student_hobby WHERE student_id = '%s';", student_id),
		"INSERT INTO student_hobby (student_id, hobby_id) VALUES ",
	}

	for query_index, query := range queries {
		column := "major"
		data := student.HobbyID

		if query_index == 3 {
			for h := 0; h < len(data); h++ {
				query = utils.AddValues(student_id, query, column, data[h], h, data)
			}
		}
		_, err := db.Exec(query)
		if err != nil {
			utils.InternalServerErrorResponse(err, ctx, fmt.Sprintf("Failed to update %s data", strings.Split(query, " ")[2]))
			return
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprint("success update student data by id", student_id),
	})
}

func DeleteStudent(ctx *gin.Context) {
	student_id := ctx.Param("id")

	queries := []string{
		fmt.Sprintf("DELETE FROM students WHERE id = '%s';", student_id),
		fmt.Sprintf("DELETE FROM student_major WHERE student_id = '%s';", student_id),
		fmt.Sprintf("DELETE FROM student_hobby WHERE student_id = '%s';", student_id),
	}

	for _, query := range queries {
		_, err := db.Exec(query)
		if err != nil {
			utils.InternalServerErrorResponse(err, ctx, fmt.Sprintf("Failed to delete %s data", strings.Split(query, " ")[2]))
			return
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprint("success delete student data by id", student_id),
	})
}
