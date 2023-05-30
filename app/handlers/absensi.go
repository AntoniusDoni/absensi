package handlers

import (
	"fmt"
	"time"

	"github.com/absensi/app/models"
	"github.com/absensi/app/services"
	"github.com/absensi/pkg/authorize"
	"github.com/absensi/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/datatypes"
)

type AbsensiHandler struct {
	Ser *services.Services
}

func (absen *Handler) AddAbsen(ctx *fiber.Ctx) error {
	absenRequest := new(models.RequestAbsen)
	res := models.ResponseRequest{}
	err := ctx.BodyParser(absenRequest)
	if err != nil {
		res.Code = fiber.StatusExpectationFailed
		res.Message = "Faild to Absen"
	}
	timenow := time.Now()
	date := utils.ConvertDate(timenow.Year(), int(timenow.Month()), timenow.Day())
	absens := new(models.Absensi)
	absens.Nip = authorize.GetNip(ctx)
	absens.DateAt = timenow
	user_id, _ := authorize.GetRole(ctx)
	absens.IdPegawai = user_id
	errs := absen.Ser.Validator.ValidateRequest(absenRequest)
	if errs != nil {
		res.Code = fiber.StatusBadRequest
		res.Message = fmt.Sprintf("%s %v", "Request body invalid", errs)

	}
	check, err := absen.Ser.Repository.GetDetailAbsen(user_id, date, "")
	if err != nil {
		absens.TimeIn = datatypes.NewTime(timenow.Hour(), timenow.Minute(), timenow.Second(), 0)
		absens.IsIn = true
		absens.IsOut = false
		err = absen.Ser.Repository.CreateAbsen(absens)
		if err != nil {
			res.Code = fiber.StatusUnprocessableEntity
			res.Message = fmt.Sprintf("%s %v", "Failed to Save", err)
		}
		res.Code = fiber.StatusOK
		res.Message = "Sucsess"
		res.Data = fiber.Map{
			"nip":    absens.Nip,
			"add_at": absens.TimeIn,
			"absen":  "masuk",
		}
		return ctx.Status(res.Code).JSON(res)
	}
	absens.TimeOut = datatypes.NewTime(timenow.Hour(), timenow.Minute(), timenow.Second(), 0)
	updateresult := absen.Ser.Repository.UpdateAbsen(check.Id, absens.TimeOut)

	if updateresult == 0 {
		res.Code = fiber.StatusAccepted
		res.Message = "you are has a logout"
		return ctx.Status(res.Code).JSON(res)
	}

	if err == nil {
		res.Code = fiber.StatusCreated
		res.Message = "Sucsess"
		res.Data = fiber.Map{
			"nip":    absens.Nip,
			"add_at": absens.TimeOut,
			"absen":  "Keluar",
		}
	}

	return ctx.Status(res.Code).JSON(res)
}
