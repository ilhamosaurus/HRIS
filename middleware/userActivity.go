package middleware

import (
	"bytes"
	"encoding/json"
	"io"
	"mime"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/ilhamosaurus/HRIS/model"
	"github.com/ilhamosaurus/HRIS/pkg/types"
	"github.com/ilhamosaurus/HRIS/pkg/util"
	"github.com/labstack/echo/v4"
)

var (
	skippedMethod          = []string{http.MethodGet, http.MethodOptions}
	skkipedBodyCheck       = []string{"login", "logout"}
	maxBodySize      int64 = 10 << 20
)

func ActivityMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := next(c); err != nil {
			return err
		}

		for i := range skippedMethod {
			if c.Request().Method == skippedMethod[i] {
				return nil
			}
		}

		var details map[string]any
		for i := range skkipedBodyCheck {
			if !strings.Contains(c.Path(), skkipedBodyCheck[i]) {
				details = make(map[string]any)
				if c.ParamValues() != nil && len(c.ParamValues()) > 0 {
					details["params"] = c.ParamValues()
				}

				if c.QueryParams() != nil && len(c.QueryParams()) > 0 {
					details["query"] = c.QueryParams()
				}

				req := c.Request()
				if req.Body != nil {
					// limit reader to avoid OOM
					lr := io.LimitReader(req.Body, maxBodySize+1)
					raw, err := io.ReadAll(lr)
					// restore body regardless so downstream can read it
					req.Body = io.NopCloser(bytes.NewReader(raw))

					if err != nil {
						details["body_error"] = err.Error()
					} else if len(raw) > 0 {
						if int64(len(raw)) > maxBodySize {
							details["body_error"] = "body too large"
						} else {
							// try to decide how to store body based on content-type
							ct := req.Header.Get("Content-Type")
							mediaType, _, _ := mime.ParseMediaType(ct)

							switch mediaType {
							case "application/json", "application/ld+json":
								var v interface{}
								if err := json.Unmarshal(raw, &v); err == nil {
									details["body"] = v // parsed JSON (map[string]any or []any)
								} else {
									details["body_raw"] = string(raw)
								}
							case "application/x-www-form-urlencoded":
								if vals, err := url.ParseQuery(string(raw)); err == nil {
									details["body"] = vals // url.Values
								} else {
									details["body_raw"] = string(raw)
								}
							default:
								// fallback — store raw string
								details["body_raw"] = string(raw)
							}
						}
					}
				}
			}
		}

		var username string
		if !strings.Contains(c.Path(), "login") {
			auth := util.GetUserAuth(c)
			username = auth.Username
		} else {
			var req types.LoginRequest
			if err := c.Bind(&req); err != nil {
				return err
			}
			username = req.Username
		}
		log := model.UserActivity{
			Time:       time.Now(),
			Username:   username,
			Address:    c.RealIP(),
			Feature:    c.Path(),
			AccessType: c.Request().Method,
		}

		if details != nil {
			d, err := json.Marshal(details)
			if err != nil {
				return err
			}
			stringD := string(d)
			log.AccessDetails = &stringD
		}

		return model.AddUserActivity(log)
	}
}
