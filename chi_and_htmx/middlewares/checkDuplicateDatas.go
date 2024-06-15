package middlewares

import (
	"fmt"
	"log"

	"github.com/PranitRout07/Practice-Golang/chi_and_htmx/initializers"
)

func CheckDuplicateDatas(articleName string) (bool) {
	sqlQuery := fmt.Sprintf("SELECT COUNT(*) FROM posts WHERE title = '%s';",articleName)
	res, err := initializers.DBConnection.Exec(sqlQuery)

	log.Println("Print duplicate result", res)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}
