package main

import (
	"log"
	"os"

	"github.com/Sofiviola/Examen.git/cmd/server/handler"
	"github.com/Sofiviola/Examen.git/docs"
	"github.com/Sofiviola/Examen.git/internal/dentista"
	"github.com/Sofiviola/Examen.git/internal/paciente"
	"github.com/Sofiviola/Examen.git/internal/turno"
	"github.com/Sofiviola/Examen.git/pkg/middleware"
	"github.com/Sofiviola/Examen.git/pkg/store/dentistaStore"
	"github.com/Sofiviola/Examen.git/pkg/store/pacienteStore"
	"github.com/Sofiviola/Examen.git/pkg/store/turnoStore"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)
var(
	storageDB *sql.DB
)
func main() {
	dataSource := "root:2021@tcp(localhost:3306)/viola_examen_go"
	storageDB, err := sql.Open("mysql", dataSource)
	if err != nil {
		log.Fatal(err)
	}

	if err := godotenv.Load("./.env"); err != nil {
		panic("Error loading .env file: " + err.Error())
	}

	//Dentista
	storage := dentistaStore.NewSqlStore(storageDB)
	repo := dentista.NewRepository(storage)
	service := dentista.NewService(repo)
	dentistaHandler := handler.NewDentistaHandler(service)

	r := gin.Default()
	docs.SwaggerInfo.Host = os.Getenv("HOST")
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })
	dentistas := r.Group("/dentistas")
	{
		dentistas.GET(":id", dentistaHandler.GetByID())
		dentistas.GET("matricula/:matricula", dentistaHandler.GetByMatricula())
		dentistas.POST("", middleware.Authentication(), dentistaHandler.Post())
		dentistas.PUT(":id", middleware.Authentication(), dentistaHandler.Put())
		dentistas.DELETE(":id", middleware.Authentication(), dentistaHandler.Delete())
		dentistas.PATCH(":id", middleware.Authentication(), dentistaHandler.Patch())
	}

   //paciente
    storagePaciente := pacienteStore.NewSqlStore(storageDB)
	repoPaciente := paciente.NewRepository(storagePaciente)
	servicePaciente := paciente.NewService(repoPaciente)
	pacienteHandler := handler.NewPacienteHandler(servicePaciente)
	
	pacientes := r.Group("/pacientes")
	{
		pacientes.GET(":id", pacienteHandler.GetByID())
		pacientes.GET("/dni/:dni", pacienteHandler.GetByDni())
		pacientes.POST("", middleware.Authentication(), pacienteHandler.Post())
		pacientes.PUT(":id", middleware.Authentication(), pacienteHandler.Put())
		pacientes.DELETE(":id", middleware.Authentication(), pacienteHandler.Delete())
		pacientes.PATCH(":id", middleware.Authentication(), pacienteHandler.Patch())
	}

	//Turno
    storageTurno := turnoStore.NewSqlStore(storageDB)
	repoTurno := turno.NewRepository(storageTurno)
	serviceTurno := turno.NewService(repoTurno,repoPaciente,repo)
	turnoHandler := handler.NewTurnoHandler(serviceTurno)
	
	turnos := r.Group("/turnos")
	{
		turnos.GET(":id", turnoHandler.GetByID())
		turnos.GET("dni/:dni", turnoHandler.GetByDniPaciente())
		turnos.POST("", middleware.Authentication(), turnoHandler.Post())
		turnos.PUT(":id", middleware.Authentication(), turnoHandler.Put())
		turnos.DELETE(":id", middleware.Authentication(), turnoHandler.Delete())
		turnos.PATCH(":id", middleware.Authentication(), turnoHandler.Patch())
	}
	r.Run(":8080")
}
