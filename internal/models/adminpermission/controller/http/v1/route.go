package v1 

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/jmoiron/sqlx"
	"github.com/nutikuli/internProject_backend/internal/models/adminpermission/repository"
	_AdminPerUse "github.com/nutikuli/internProject_backend/internal/models/adminpermission/usecase"
	_fileRepo "github.com/nutikuli/internProject_backend/internal/services/file/repository" 
)

func UseAdminPermissionRoute(db *sqlx.DB, app fiber.Router) {
	authR := app.Group("/adminpermission", func(c *fiber.Ctx) error {
		log.Infof("adminpermission : %v", c.Request().URI().String())
		return c.Next()
	})

	fileRepo := _fileRepo.NewFileRepository(db)
	adminperRepo :=repository.NewAdminPermissionRepository(db)

	

	AdminPermissionUseCase := _AdminPerUse.NewAdminpermissionUsecase(adminperRepo ,fileRepo)

	AdminPermissionConn := NewAdminPermissionHandler(AdminPermissionUseCase)

	authR.Post("/adminper-register", AdminPermissionConn.CreateAdminPermissionAccount)
	authR.Get("/:admin_id", AdminPermissionConn.GetAdminePermissionById)

}