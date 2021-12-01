package beater

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

var db *sql.DB

var server = "DESKTOP-UBNV5QG"
var port = 1433
var user = "usuario_dev"
var password = "david"
var database = "SeusInventariosFarmacias"

type Inventario struct {
	uiCodSucursal         int
	uiCodProducto         int
	bExistenciaINV        bool
	bExistenciaPro        bool
	nExistencia           string
	nExistenciaApartada   string
	nExistenciaDisponible string
	bAgotado              bool
	bDescontinuado        bool
	bAlertaSanitaria      bool
	dFechaVenta           string
	cControlCOFEPRIS      string
}

func main() {
	startConection()
	ReadEmployees()
}

func startConection() {
	// INICIAMOS CONEXION
	// Build connection string
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		server, user, password, port, database)

	var err error

	// Create connection pool
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}
	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("Connected!\n")
}

// ReadEmployees reads all employee records
func ReadEmployees() ([]Inventario, error) {
	ctx := context.Background()

	// Check if database is alive.
	err := db.PingContext(ctx)
	if err != nil {
		return nil, err
	}
	tsql := fmt.Sprintf("EXEC spECommerceConsExistenciaSucursalesTres;")

	// Execute query
	rows, err := db.QueryContext(ctx, tsql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var producto []Inventario
	// Iterate through the result set.

	for rows.Next() {
		var uiCodSucursal, uiCodProducto int
		var nExistencia, nExistenciaApartada, nExistenciaDisponible, dFechaVenta, cControlCOFEPRIS string
		var bExistenciaINV, bExistenciaPro, bAgotado, bDescontinuado, bAlertaSanitaria bool

		// Get values from row.
		err := rows.Scan(&uiCodSucursal, &uiCodProducto, &bExistenciaINV, &bExistenciaPro, &nExistencia, &nExistenciaApartada, &nExistenciaDisponible, &bAgotado, &bDescontinuado, &bAlertaSanitaria, &dFechaVenta, &cControlCOFEPRIS)
		p := Inventario{uiCodSucursal: uiCodSucursal, uiCodProducto: uiCodProducto, bExistenciaINV: bExistenciaINV, bExistenciaPro: bExistenciaPro, nExistencia: nExistencia, nExistenciaApartada: nExistenciaApartada, nExistenciaDisponible: nExistenciaDisponible, bAgotado: bAgotado, bDescontinuado: bDescontinuado, bAlertaSanitaria: bAlertaSanitaria, dFechaVenta: dFechaVenta, cControlCOFEPRIS: cControlCOFEPRIS}
		producto = append(producto, p)

		if err != nil {
			fmt.Printf(err.Error())
			return nil, err
		}

		fmt.Printf("uiCodSucursal: %d, uiCodProducto: %d, bExistenciaINV: %t,\n bExistenciaPro: %t,\n nExistencia: %s,\n %s, %s, %t, %t, %t, %s, %s", uiCodSucursal, uiCodProducto, bExistenciaINV, bExistenciaPro, nExistencia, nExistenciaApartada, nExistenciaDisponible, bAgotado, bDescontinuado, bAlertaSanitaria, dFechaVenta, cControlCOFEPRIS)
	}
	return producto, nil
}
