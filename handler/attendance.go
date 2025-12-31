package handler

// import (
// 	"errors"
// 	"net/http"
// 	"time"

// 	"github.com/ilhamosaurus/HRIS/model"
// 	"github.com/ilhamosaurus/HRIS/pkg/types"
// 	"github.com/ilhamosaurus/HRIS/pkg/util"
// 	"github.com/labstack/echo/v4"
// )

// func (h *Handler) CheckIn(c echo.Context) error {
// 	auth := util.GetUserAuth(c)
// 	now := time.Now()
// 	attendance := h.model.GetAttendace(auth.Username, now)

// 	overtime := h.model.GetOvertime(auth.Username, types.Approved.String(), now)
// 	if now.Weekday() == time.Saturday && overtime.Username != auth.Username || now.Weekday() == time.Sunday && overtime.Username != auth.Username {
// 		return c.JSON(types.GenerateReponse(http.StatusBadRequest, "no overtime approved", nil))
// 	}

// 	if attendance.Username == auth.Username && !attendance.CheckIn.IsZero() {
// 		return c.JSON(types.GenerateReponse(http.StatusBadRequest, "already checked in", nil))
// 	}

// 	attendance.Username = auth.Username
// 	attendance.Date = now
// 	attendance.CheckIn = now
// 	if err := h.model.AddAttendance(attendance); err != nil {
// 		return c.JSON(types.GenerateReponse(http.StatusInternalServerError, err.Error(), nil))
// 	}
// 	return c.JSON(types.GenerateReponse(http.StatusOK, "OK", nil))
// }

// func (h *Handler) CheckOut(c echo.Context) error {
// 	auth := util.GetUserAuth(c)
// 	now := time.Now()
// 	attendance := h.model.GetAttendace(auth.Username, now)

// 	if attendance.Username == auth.Username && attendance.CheckOut != nil {
// 		return c.JSON(types.GenerateReponse(http.StatusBadRequest, "already checked out", nil))
// 	}

// 	if attendance.CheckIn.IsZero() {
// 		return c.JSON(types.GenerateReponse(http.StatusBadRequest, "not checked in", nil))
// 	}

// 	attendance.CheckOut = &now
// 	if err := h.model.UpdateAttendance(attendance); err != nil {
// 		return c.JSON(types.GenerateReponse(http.StatusInternalServerError, err.Error(), nil))
// 	}
// 	return c.JSON(types.GenerateReponse(http.StatusOK, "OK", nil))
// }

// func (h *Handler) SetAttendance(c echo.Context) error {
// 	auth := util.GetUserAuth(c)
// 	if auth.Role != types.Admin {
// 		return c.JSON(types.GenerateReponse(http.StatusUnauthorized, "unauthorized", nil))
// 	}

// 	var req types.Attendance
// 	if err := c.Bind(&req); err != nil {
// 		return c.JSON(types.GenerateReponse(http.StatusBadRequest, err.Error(), nil))
// 	}

// 	if err := c.Validate(req); err != nil {
// 		return c.JSON(types.GenerateReponse(http.StatusBadRequest, err.Error(), nil))
// 	}

// 	attendance, err := ToModelAttendace(req)
// 	if err != nil {
// 		return c.JSON(types.GenerateReponse(http.StatusBadRequest, err.Error(), nil))
// 	}

// 	if req.ID != nil {
// 		if err := h.model.UpdateAttendance(attendance); err != nil {
// 			return c.JSON(types.GenerateReponse(http.StatusInternalServerError, err.Error(), nil))
// 		}
// 		return c.JSON(types.GenerateReponse(http.StatusOK, "OK", nil))
// 	}

// 	if err := h.model.AddAttendance(attendance); err != nil {
// 		return c.JSON(types.GenerateReponse(http.StatusInternalServerError, err.Error(), nil))
// 	}
// 	return c.JSON(types.GenerateReponse(http.StatusOK, "OK", nil))
// }

// func (h *Handler) GetAttendances(c echo.Context) error {
// 	auth := util.GetUserAuth(c)

// 	var existingAttendances []model.Attendance
// 	if auth.Role != types.Admin {
// 		existingAttendances = h.model.GetAttendaces(&model.Attendance{Username: auth.Username})
// 	} else {
// 		cond := model.Attendance{}
// 		if c.QueryParam("username") != "" {
// 			cond.Username = c.QueryParam("username")
// 		}
// 		if c.QueryParam("date") != "" {
// 			date, err := StringToDate(c.QueryParam("date"))
// 			if err != nil {
// 				return c.JSON(types.GenerateReponse(http.StatusBadRequest, err.Error(), nil))
// 			}
// 			cond.Date = date
// 		}

// 		existingAttendances = h.model.GetAttendaces(&cond)
// 	}

// 	attendances := make([]*types.Attendance, len(existingAttendances))
// 	for i := range existingAttendances {
// 		attendances[i] = ToTypeAttendance(existingAttendances[i])
// 	}
// 	return c.JSON(types.GenerateReponse(http.StatusOK, "OK", attendances))
// }

// func ToModelAttendace(req types.Attendance) (model.Attendance, error) {
// 	attendance := model.Attendance{Username: req.Username}

// 	if req.ID != nil {
// 		attendance.ID = req.GetID()
// 	}

// 	if req.Date != nil {
// 		date, err := StringToDate(req.GetDate())
// 		if err != nil {
// 			return model.Attendance{}, err
// 		}
// 		attendance.Date = date
// 	}

// 	if req.CheckIn != nil {
// 		checkIn, err := ParseTimeFromString(req.GetCheckIn())
// 		if err != nil {
// 			return model.Attendance{}, err
// 		}
// 		attendance.CheckIn = checkIn
// 	}

// 	if req.CheckOut != nil && req.CheckIn != nil {
// 		checkOut, err := ParseTimeFromString(req.GetCheckOut())
// 		if attendance.CheckIn.After(checkOut) {
// 			return model.Attendance{}, errors.New("check out must be after check in")
// 		}
// 		if err != nil {
// 			return model.Attendance{}, err
// 		}
// 		attendance.CheckOut = &checkOut
// 	}
// 	return attendance, nil
// }

// func ToTypeAttendance(attendance model.Attendance) *types.Attendance {
// 	typeAttendance := types.Attendance{
// 		ID:       &attendance.ID,
// 		Username: attendance.Username,
// 	}

// 	date := DateToString(attendance.Date)
// 	typeAttendance.Date = &date

// 	checkIn := TimeToString(attendance.CheckIn)
// 	typeAttendance.CheckIn = &checkIn

// 	if attendance.CheckOut != nil && !attendance.CheckOut.IsZero() {
// 		checkOut := TimeToString(*attendance.CheckOut)
// 		typeAttendance.CheckOut = &checkOut
// 	}

// 	return &typeAttendance
// }

// func ParseTimeFromString(s string) (time.Time, error) {
// 	return time.Parse("2006-01-02T15:04:05Z07:00", s)
// }

// func StringToDate(s string) (time.Time, error) {
// 	t, err := time.Parse("20060102", s)
// 	return t, err
// }

// func DateToString(t time.Time) string {
// 	return t.Format("20060102")
// }

// func TimeToString(t time.Time) string {
// 	return t.Format("2006-01-02T15:04:05Z07:00")
// }
