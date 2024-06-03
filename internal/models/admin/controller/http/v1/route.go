package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/jmoiron/sqlx"
	_accRepo "github.com/nutikuli/internProject_backend/internal/models/account/repository"
	_accUse "github.com/nutikuli/internProject_backend/internal/models/account/usecase"
	_adperRepo "github.com/nutikuli/internProject_backend/internal/models/adminpermission/repository"
	"github.com/nutikuli/internProject_backend/internal/models/admin/repository"
	_AdminUse "github.com/nutikuli/internProject_backend/internal/models/admin/usecase"
	_fileRepo "github.com/nutikuli/internProject_backend/internal/services/file/repository"
)

func UseAdminRoute(db *sqlx.DB, app fiber.Router) {
	authR := app.Group("/admin", func(c *fiber.Ctx) error {
		log.Infof("admin : %v", c.Request().URI().String())
		return c.Next()
	})

	fileRepo := _fileRepo.NewFileRepository(db)
	adminRepo := repository.NewFileRepository(db)
	adperRepo := _adperRepo.NewAdminPermissionRepository(db)

	accRepo := _accRepo.NewAccountRepository(db)
	accUse := _accUse.NewAccountUsecase(accRepo, fileRepo, adminRepo, nil, nil)

	AdminUseCase := _AdminUse.NewAdminUsecase(adminRepo , fileRepo ,accUse , adperRepo)

	AdminConn := NewAdminHandler(AdminUseCase)

	authR.Post("/admin-register", AdminConn.RegisterAdminAccount)
	authR.Get("/:admin_id", AdminConn.GetAdmineById)
	authR.Put("/admin_id", AdminConn.UpdateAdminById)
	authR.Delete("/admin_id", AdminConn.DeletedAdminByID)
}