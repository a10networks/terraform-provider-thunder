package go_thunder

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"github.com/clarketm/json" // "encoding/json"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"
	"util"
)

type FileSslKey struct {
	FileSslKeyInstanceFile FileSslKeyInstance `json:"ssl-key,omitempty"`
}

type FileSslKeyInstance struct {
	FileSslKeyInstanceAction     string `json:"action,omitempty"`
	FileSslKeyInstanceDstFile    string `json:"dst-file,omitempty"`
	FileSslKeyInstanceFile       string `json:"file,omitempty"`
	FileSslKeyInstanceFileHandle string `json:"file-handle,omitempty"`
	FileSslKeyInstanceSecured    int    `json:"secured,omitempty"`
	FileSslKeyInstanceSize       int    `json:"size,omitempty"`
	FileSslKeyInstanceUUID       string `json:"uuid,omitempty"`
}

func PostFileSslKeyImport(id string, inst FileSslKey, host string, FileHandleLocal string) error {

	logger := util.GetLoggerInstance()
	method := "POST"
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	url := "https://" + host + "/axapi/v3/file/ssl-key"
	mimeHeader1 := make(map[string][]string)
	mimeHeader1["Content-Disposition"] = append(mimeHeader1["Content-Disposition"], "form-data; name=\"json\"")
	mimeHeader1["Content-Type"] = append(mimeHeader1["Content-Type"], "application/json")
	fieldWriter1, _ := writer.CreatePart(mimeHeader1)

	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	fieldWriter1.Write(payloadBytes)
	file, errFile2 := os.Open(string(FileHandleLocal))
	defer file.Close()

	part2, errFile2 := writer.CreateFormFile("file", filepath.Base(FileHandleLocal))

	_, errFile2 = io.Copy(part2, file)
	if errFile2 != nil {
		logger.Println(errFile2)
		return errFile2
	}

	err = writer.Close()
	if err != nil {
		logger.Println(err)
		return err
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	//client := &http.Client{Transport: tr}
	client := &http.Client{Transport: tr, Timeout: 30 * time.Second}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		logger.Println(err)
		return err
	}

	req.Header.Add("Authorization", id)
	req.Header.Add("Accept", "application/json")

	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logger.Println(err)
		return err
	}
	//logger.Println(string(data))

	err = check_api_status("PostFileSslCert", data)
	if err != nil {
		return err
	}

	return err
}

func PostFileSslKey(id string, inst FileSslKey, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostFileSslKey")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
		return err
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/file/ssl-key", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m FileSslKey
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Post REQ RES..........................", m)
			err := check_api_status("PostFileSslKey", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

// func GetFileSslKey(id string, host string) (*FileSslKey, error) {

// 	logger := util.GetLoggerInstance()

// 	var headers = make(map[string]string)
// 	headers["Accept"] = "application/json"
// 	headers["Content-Type"] = "application/json"
// 	headers["Authorization"] = id
// 	logger.Println("[INFO] Inside GetFileSslKey")

// 	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/file/ssl-key", nil, headers)

// 	if err != nil {
// 		logger.Println("The HTTP request failed with error ", err)
// 		return nil, err
// 	} else {
// 		data, _ := ioutil.ReadAll(resp.Body)
// 		var m FileSslKey
// 		err := json.Unmarshal(data, &m)
// 		if err != nil {
// 			logger.Println("Unmarshal error ", err)
// 			return nil, err
// 		} else {
// 			logger.Println("[INFO] Get REQ RES..........................", m)
// 			err := check_api_status("GetFileSslKey", data)
// 			if err != nil {
// 				return nil, err
// 			}
// 			return &m, nil
// 		}
// 	}

// }
