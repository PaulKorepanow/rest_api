package server

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

const (
	amountOfUser = 20
	defaultTo    = "4000-01-01T12:37:18.485Z"
	defaultFrom  = "0001-01-01T12:37:18.485Z"
)

func HandleUsers() echo.HandlerFunc {
	return func(c echo.Context) error {
		var fromParam string
		fromParam = c.QueryParam("from")
		if fromParam == "" {
			fromParam = defaultFrom
		}
		from, err := time.Parse(time.RFC3339, fromParam)
		if err != nil {
			return err
		}

		var toParam string
		toParam = c.QueryParam("to")
		if toParam == "" {
			toParam = defaultTo
		}
		to, err := time.Parse(time.RFC3339, toParam)
		if err != nil {
			return err
		}

		var resp *http.Response
		var counter int
		for {
			if counter == 5 {
				return echo.NewHTTPError(http.StatusGatewayTimeout, "can't get response")
			}
			counter++
			resp, err = http.Get(fmt.Sprintf("https://randomuser.me/api/?results=%d", amountOfUser))
			if err != nil {
				return err
			}

			if resp.StatusCode == http.StatusOK {
				break
			}
			resp.Body.Close()
			time.Sleep(500 * time.Millisecond)
		}

		defer resp.Body.Close()

		res := Users{}
		if err = json.NewDecoder(resp.Body).Decode(&res); err != nil {
			return err
		}

		var users []UserLocal
		for _, u := range res.Users {
			date, err := time.Parse(time.RFC3339, u.Registration.Date)
			if err != nil {
				return err
			}
			if date.After(from) && date.Before(to) {
				postcode, err := u.Location.getString()
				if err != nil {
					return err
				}
				users = append(users, UserLocal{
					Gender:    u.Gender,
					FirstName: u.Name.First,
					LastName:  u.Name.Last,
					Postcode:  postcode,
					CreatedAt: u.Registration.Date,
				})
			}

		}

		return c.JSON(http.StatusOK, &users)
	}
}

func HandlePost() echo.HandlerFunc {
	type test struct {
		Status string
		From   string `json:"from"`
		To     string `json:"to"`
	}
	return func(c echo.Context) error {
		newTest := new(test)
		if err := c.Bind(newTest); err != nil {
			return err
		}
		newTest.Status = "Success"
		return c.JSON(http.StatusOK, newTest)
	}
}
