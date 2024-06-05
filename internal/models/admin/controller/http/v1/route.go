package v1

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/jmoiron/sqlx"
	_accRepo "github.com/nutikuli/internProject_backend/internal/models/account/repository"
	_accUse "github.com/nutikuli/internProject_backend/internal/models/account/usecase"
	"github.com/nutikuli/internProject_backend/internal/models/admin/repository"
	_AdminUse "github.com/nutikuli/internProject_backend/internal/models/admin/usecase"
	_adperRepo "github.com/nutikuli/internProject_backend/internal/models/adminpermission/repository"
	_fileRepo "github.com/nutikuli/internProject_backend/internal/services/file/repository"
	_fileUse "github.com/nutikuli/internProject_backend/internal/services/file/usecase"
	"github.com/nutikuli/internProject_backend/pkg/middlewares"

	_logRepo "github.com/nutikuli/internProject_backend/internal/models/logdata/repository"
)

func UseAdminRoute(db *sqlx.DB, app fiber.Router) {
	authR := app.Group("/admin", func(c *fiber.Ctx) error {
		log.Infof("admin : %v", c.Request().URI().String())
		return c.Next()
	})

	logRepo := _logRepo.NewLoggerRepository(db)
	logger := middlewares.NewLogger(logRepo)

	fileRepo := _fileRepo.NewFileRepository(db)
	adminRepo := repository.NewFileRepository(db)
	adperRepo := _adperRepo.NewAdminPermissionRepository(db)
	fileUse := _fileUse.NewFileUsecase(fileRepo)

	accRepo := _accRepo.NewAccountRepository(db)
	accUse := _accUse.NewAccountUsecase(accRepo, fileRepo, adminRepo, nil, nil)

	AdminUseCase := _AdminUse.NewAdminUsecase(adminRepo, fileRepo, accUse, adperRepo, fileUse)

	AdminConn := NewAdminHandler(AdminUseCase)

	authR.Post("/admin-register",
		func(c *fiber.Ctx) error {
			logAction := &middlewares.LoggerAction{
				Menu:   "ผู้ดูแลระบบ",
				Action: "สร้างข้อมูล",
			}

			return logger.LogRequest(c, logAction)
		},
		AdminConn.RegisterAdminAccount)
	authR.Get("/getalladmin", AdminConn.GetAllAdmin)
	authR.Get("/:admin_id",
		func(c *fiber.Ctx) error {
			logAction := &middlewares.LoggerAction{
				Menu:   "แอดมิน",
				Action: fmt.Sprintf("ดูข้อมูล Admin Account หมายเลข  %s", c.Params("admin_id")),
			}

			return logger.LogRequest(c, logAction)
		},

		AdminConn.GetAdmineById)
	authR.Patch("/:admin_id",
		func(c *fiber.Ctx) error {
			logAction := &middlewares.LoggerAction{
				Menu:   "แอดมิน",
				Action: fmt.Sprintf("อัพเดทข้อมูล Admin Account หมายเลข  %s", c.Params("admin_id")),
			}

			return logger.LogRequest(c, logAction)
		},
		AdminConn.UpdateAdminById)
	authR.Delete("/:admin_id",
		func(c *fiber.Ctx) error {
			logAction := &middlewares.LoggerAction{
				Menu:   "แอดมิน",
				Action: fmt.Sprintf("ลบข้อมูล Admin Account หมายเลข  %s", c.Params("admin_id")),
			}

			return logger.LogRequest(c, logAction)
		},
		AdminConn.DeletedAdminByID)
}
