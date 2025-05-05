package client

import (
	"encoding/json"
	"errors"
	"github.com/go-resty/resty/v2"
	"nurcenter/productivity-service/config"
	"strconv"
)

func VerifyUser(userID uint, cfg *config.Config) error {
	client := resty.New()
	resp, err := client.R().Get(cfg.UserServiceURL + "/users/" + strconv.Itoa(int(userID)))
	if err != nil {
		return err
	}
	if resp.StatusCode() != 200 {
		return errors.New("user not found")
	}

	var user map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &user); err != nil {
		return err
	}
	if userID != uint(user["id"].(float64)) {
		return errors.New("user ID mismatch")
	}
	return nil
}
