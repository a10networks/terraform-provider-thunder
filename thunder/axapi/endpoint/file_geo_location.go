package endpoint

import (
	"io/ioutil"
	"net/http"
	"os"

	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 5_2_1-P6_74
type FileGeoLocation struct {
	Inst struct {
		Action      string `json:"action"`
		File        string `json:"file"`
		FileHandle  string `json:"file-handle"`
		FileContent []byte `json:"-"`
	} `json:"geo-location"`
}

type DeleteFileGeoLocation struct {
	Inst struct {
		FileName string `json:"filename"`
	} `json:"geo-location"`
}

func (p *FileGeoLocation) GetId() string {
	return "1"
}
func (p *FileGeoLocation) getPath() string {
	return "file/geo-location"
}
func (p *FileGeoLocation) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	headers := axapi.GenRequestHeader(authToken)
	f, error := os.Open(p.Inst.FileHandle)
	if error != nil {
		logger.Println("Failed to open a file: ", error)
		return error
	}

	data, error := ioutil.ReadAll(f)
	defer f.Close()
	if error != nil {
		logger.Println("Failed to read file: ", error)
		return error
	}

	s := &FileGeoLocation{}
	s.Inst.Action = p.Inst.Action
	s.Inst.FileContent = data
	s.Inst.FileHandle = p.Inst.File
	s.Inst.File = p.Inst.File

	_, err := axapi.NormalizeMultipartObject(http.MethodPost, p.getPath(), s.Inst.File, s.Inst.FileContent, s, headers, host, logger)
	return err
}

func (p *FileGeoLocation) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("FileGeoLocation::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, "file/geo-location/oper", "", nil, headers, logger)
	if err == nil {
		err = json.Unmarshal(axResp, &p)
		if err != nil {
			logger.Println("json.Unmarshal() failed with error", err)
		}
	}
	return err
}
func (p *FileGeoLocation) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("FileGeoLocation::Put")
	headers := axapi.GenRequestHeader(authToken)
	payloadBytes, err := axapi.SerializeToJson(p)
	if err != nil {
		logger.Println("Failed to serialize struct as json", err)
		return err
	}
	logger.Println("payload: " + string(payloadBytes))
	_, _, err = axapi.SendPut(host, p.getPath(), "", payloadBytes, headers, logger)
	return err
}
func (p *FileGeoLocation) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("FileGeoLocation::Delete")
	headers := axapi.GenRequestHeader(authToken)

	s := &DeleteFileGeoLocation{}
	s.Inst.FileName = p.Inst.File
	payloadBytes, err := json.Marshal(s)
	logger.Println(s)
	if err != nil {
		logger.Println("json.Marshal() failed with error", err)
		return err
	}
	logger.Println("payload:", string(payloadBytes))
	_, _, err = axapi.SendPost(host, "delete/geo-location", payloadBytes, headers, logger)
	return err
}
