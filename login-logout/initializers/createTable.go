package initializers

import (
	"log"
)

func Create_table() {
	sqlQuery := `CREATE TABLE person(  
					id int NOT NULL PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
					email VARCHAR(255),
					password VARCHAR(255)
					);`
	_, err := DBConnection.Exec(sqlQuery)
	if err != nil {
		log.Fatal(err)
	}

}
