package middlewares

import (
	"fmt"
	"log"

	"github.com/PranitRout07/Practice-Golang/chi_and_htmx/initializers"
)

func CheckDuplicateDatas(articleName string) (bool) {
	var count int
	sqlQuery := fmt.Sprintf("SELECT COUNT(*) FROM posts WHERE title = '%s';",articleName)
	err := initializers.DBConnection.QueryRow(sqlQuery).Scan(&count)

	// log.Println("Print duplicate result", res)
	if err != nil {
		log.Fatal(err)
		return false
	}
	log.Println("Count : ",count)
	return count>=1
}
