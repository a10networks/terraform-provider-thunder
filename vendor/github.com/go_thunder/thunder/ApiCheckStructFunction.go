package go_thunder

import (
	"errors"
	"github.com/clarketm/json" // "encoding/json"
	"util"
)

type response_struct1 struct {
	Response struct {
		Status string `json:"status"`
		Err    struct {
			Code int    `json:"code"`
			From string `json:"from"`
			Msg  string `json:"msg"`
		} `json:"err"`
	} `json:"response"`
}

func check_api_status(resource_name_with_method string, byte_body []uint8) error {

	logger := util.GetLoggerInstance()

	var response_struct_var response_struct1
	string_data := string(byte_body)
	json.Unmarshal([]byte(string_data), &response_struct_var)
	if response_struct_var.Response.Status == "fail" {
		panic_string := string(resource_name_with_method) + "  --->  " + string(response_struct_var.Response.Err.Msg)
		logger.Println(panic_string)
		err := errors.New(panic_string)
		return err
	}
	return nil
}
