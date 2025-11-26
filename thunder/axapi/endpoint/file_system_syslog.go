package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
	"io/ioutil"
	"net/http"
	"os"
)

// based on ACOS 7_0_2-102
type FileSystemSyslog struct {
	Inst struct {
		Filename string `json:"filename"`

		Rm int `json:"rm"`

		Uuid string `json:"uuid"`

		FileContent []byte `json:"-"`

		File string `json:"file"`

		FileHandle string `json:"file-handle"`
	} `json:"syslog"`
}
type DeleteFileSystemSyslog struct {
	Inst struct {
		FileName string `json:"filename"`
	} `json:"syslog"`
}

func (p *FileSystemSyslog) GetId() string {
	return "1"
}

func (p *FileSystemSyslog) getPath() string {
	return "file-system/syslog"
}

func (p *FileSystemSyslog) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("FileSystemSyslog::Post")
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
	s := &FileSystemSyslog{}
	s.Inst.Filename = p.Inst.Filename
	s.Inst.Rm = p.Inst.Rm
	s.Inst.Uuid = p.Inst.Uuid
	s.Inst.FileContent = data
	s.Inst.File = p.Inst.File
	s.Inst.FileHandle = p.Inst.File
	_, err := axapi.NormalizeMultipartObject(http.MethodPost, p.getPath(), s.Inst.File, s.Inst.FileContent, s, headers, host, logger)
	return err
}

func (p *FileSystemSyslog) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("FileSystemSyslog::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, "file/syslog/oper", "", nil, headers, logger)
	if err == nil {
		if len(axResp) > 0 {
			err = json.Unmarshal(axResp, &p)
		}
		if err != nil {
			logger.Println("json.Unmarshal() failed with error", err)
		}
	}
	return err
}
func (p *FileSystemSyslog) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("FileSystemSyslog::Put")
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

func (p *FileSystemSyslog) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("FileSystemSyslog::Delete")
	headers := axapi.GenRequestHeader(authToken)
	s := &DeleteFileSystemSyslog{}
	s.Inst.FileName = p.Inst.File
	payloadBytes, err := json.Marshal(s)
	logger.Println(s)
	if err != nil {
		logger.Println("json.Marshal() failed with error", err)
		return err
	}
	logger.Println("payload:", string(payloadBytes))
	_, _, err = axapi.SendPost(host, "delete/syslog", payloadBytes, headers, logger)
	return err
}
