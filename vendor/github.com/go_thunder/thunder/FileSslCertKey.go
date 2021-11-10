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

type FileSslCertKey struct {
	FileSslCertKeyInstanceFile FileSslCertKeyInstance `json:"ssl-cert-key,omitempty"`
}

type FileSslCertKeyInstance struct {
	FileSslCertKeyInstanceAction     string `json:"action,omitempty"`
	FileSslCertKeyInstanceDstFile    string `json:"dst-file,omitempty"`
	FileSslCertKeyInstanceFile       string `json:"file,omitempty"`
	FileSslCertKeyInstanceFileHandle string `json:"file-handle,omitempty"`
	FileSslCertKeyInstanceSecured    int    `json:"secured,omitempty"`
	FileSslCertKeyInstanceSize       int    `json:"size,omitempty"`
}

func PostFileSslCertKeyImport(id string, inst FileSslCertKey, host string, FileHandleLocal string) error {

	logger := util.GetLoggerInstance()
	method := "POST"
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	url := "https://" + host + "/axapi/v3/file/ssl-cert-key"
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

	err = check_api_status("PostFileSslCertKeyImport", data)
	if err != nil {
		return err
	}
	// b := string(body)
	// data, _ := ioutil.ReadAll(b)
	// var m FileSslCert
	// err = json.Unmarshal(data, &m)
	// if err != nil {
	// 	logger.Println("Unmarshal error ", err)
	// 	return err
	// }

	return err
}

func PostFileSslCertKey(id string, inst FileSslCertKey, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostFileSslCertKey")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
		return err
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/file/ssl-cert-key", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m FileSslCertKey
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Post REQ RES..........................", m)
			err := check_api_status("PostFileSslCertKey", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetFileSslCertKey(id string, host string) (*FileSslCertKey, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetFileSslCertKey")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/file/ssl-cert-key", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m FileSslCertKey
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return nil, nil
		} else {
			logger.Println("[INFO] Get REQ RES..........................", m)
			err := check_api_status("GetFileSslCertKey", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}
