import (
    "database/sql"
    "fmt"
    "github.com/lib/pq"
)

const (
    host     = "localhost"
    port     = 5432
    user     = "tu_usuario"
    password = "tu_contraseña"
    dbname   = "tu_base_de_datos"
)

connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
db, err := sql.Open("postgres", connStr)
if err != nil {
    panic(err)
}
defer db.Close()

