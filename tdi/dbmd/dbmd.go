package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {

	context := flag.String("context", "none", "Context name")
	if *context == "none" {
		fmt.Println("No se ha pasado ningún contexto como parámetro")
	}

	fmt.Println("Esto es dbmd")
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=myuser1 dbname=mydb1 password=mypassword1")
	fmt.Println(err)

	var table0 Table0
	table0.ID = 1
	table0.Texto = "ESTO ES UN TEXTO QUE INSERTO"
	table0.Fechahora = time.Now()
	table0.Fecha = time.Now()
	table0.Hora = time.Now()
	db.NewRecord(table0)
	db.Create(&table0)

	db.Commit()
	db.Close()
}
