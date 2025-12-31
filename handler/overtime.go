package handler

// import (
// 	"net/http"
// 	"strings"
// 	"time"

// 	"github.com/ilhamosaurus/HRIS/model"
// 	"github.com/ilhamosaurus/HRIS/pkg/types"
// 	"github.com/ilhamosaurus/HRIS/pkg/util"
// 	"github.com/labstack/echo/v4"
// )

// func (h *Handler) SetOvertime(c echo.Context) error {
// 	auth := util.GetUserAuth(c)

// 	overtime := new(types.Overtime)
// 	if err := c.Bind(overtime); err != nil {
// 		return c.JSON(types.GenerateReponse(http.StatusBadRequest, err.Error(), nil))
// 	}

// 	if err := c.Validate(overtime); err != nil {
// 		return c.JSON(types.GenerateReponse(http.StatusBadRequest, err.Error(), nil))
// 	}

// 	if !auth.Role.IsAdmin() && auth.Username != overtime.Username {
// 		return c.JSON(types.GenerateReponse(http.StatusUnauthorized, "unauthorized", nil))
// 	}

// 	modelOvertime, err := toModelOvertime(overtime)
// 	if overtime.ID != nil {
// 		if err := h.model.UpdateOvertime(modelOvertime); err != nil {
// 			return c.JSON(types.GenerateReponse(http.StatusInternalServerError, err.Error(), nil))
// 		}
// 		return c.JSON(types.GenerateReponse(http.StatusOK, "OK", nil))
// 	}

// 	if err != nil {
// 		return c.JSON(types.GenerateReponse(http.StatusBadRequest, err.Error(), nil))
// 	}

// 	if err := h.model.AddOvertime(modelOvertime); err != nil {
// 		return c.JSON(types.GenerateReponse(http.StatusInternalServerError, err.Error(), nil))
// 	}
// 	return c.JSON(types.GenerateReponse(http.StatusOK, "OK", nil))
// }

// func (h *Handler) GetOvertime(c echo.Context) error {
// 	auth := util.GetUserAuth(c)
// 	username := c.QueryParam("username")
// 	status := c.QueryParam("status")

// 	var overtimes []model.Overtime
// 	var err error
// 	cond := new(model.Overtime)
// 	if !auth.Role.IsAdmin() {
// 		cond.Username = auth.Username
// 		if status != "" {
// 			cond.Status = strings.ToUpper(status)
// 		}

// 		overtimes, err = h.model.GetOvertimes(cond)
// 		if err != nil {
// 			return c.JSON(types.GenerateReponse(http.StatusInternalServerError, err.Error(), nil))
// 		}
// 	} else {
// 		if username != "" {
// 			cond.Username = username
// 		}
// 		if status != "" {
// 			cond.Status = strings.ToUpper(status)
// 		}

// 		overtimes, err = h.model.GetOvertimes(cond)
// 		if err != nil {
// 			return c.JSON(types.GenerateReponse(http.StatusInternalServerError, err.Error(), nil))
// 		}
// 	}

// 	res := make([]*types.Overtime, len(overtimes))
// 	for i, o := range overtimes {
// 		res[i] = toTypesOvertime(o)
// 	}

// 	return c.JSON(types.GenerateReponse(http.StatusOK, "OK", res))
// }

// func toModelOvertime(o *types.Overtime) (*model.Overtime, error) {
// 	overtime := model.Overtime{
// 		Username: o.Username,
// 	}

// 	date, err := time.Parse("20060102", o.Date)
// 	if err != nil {
// 		return nil, err
// 	}
// 	overtime.Date = date

// 	if o.ID != nil {
// 		overtime.ID = o.GetID()
// 	}

// 	if o.StartTime != nil {
// 		startTime, err := ParseTimeFromString(o.GetStartTime())
// 		if err != nil {
// 			return nil, err
// 		}
// 		if sameDate(date, startTime) {
// 			overtime.StartTime = startTime
// 		}
// 	}

// 	if o.EndTime != nil {
// 		endTime, err := ParseTimeFromString(o.GetEndTime())
// 		if err != nil {
// 			return nil, err
// 		}
// 		if sameDate(date, endTime) {
// 			overtime.EndTime = endTime
// 		}
// 	}

// 	overtime.Hours = overtime.CalculateHours()
// 	overtime.Description = o.GetDescription()
// 	if o.Status != nil {
// 		overtime.Status = o.GetStatus().String()
// 	} else {
// 		overtime.Status = types.Pending.String()
// 	}

// 	return &overtime, nil
// }

// func toTypesOvertime(o model.Overtime) *types.Overtime {
// 	startTime := o.StartTime.Format(time.RFC3339)
// 	endTime := o.EndTime.Format(time.RFC3339)
// 	return &types.Overtime{
// 		ID:          &o.ID,
// 		Username:    o.Username,
// 		Date:        o.Date.Format("20060102"),
// 		StartTime:   &startTime,
// 		EndTime:     &endTime,
// 		Hours:       &o.Hours,
// 		Status:      &o.Status,
// 		Approval:    &o.Approval,
// 		Description: &o.Description,
// 	}
// }

// func sameDate(date1, date2 time.Time) bool {
// 	return date1.Year() == date2.Year() && date1.Month() == date2.Month() && date1.Day() == date2.Day()
// }
