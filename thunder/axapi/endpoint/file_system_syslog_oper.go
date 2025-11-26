package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
	"io/ioutil"
	"net/http"
	"os"
)

// based on ACOS 7_0_2-102
type FileSystemSyslogOper struct {
	Inst struct {
		Oper FileSystemSyslogOperOper `json:"oper"`

		FileContent []byte `json:"-"`

		File string `json:"file"`

		FileHandle string `json:"file-handle"`
	} `json:"syslog"`
}
type DeleteFileSystemSyslogOper struct {
	Inst struct {
		FileName string `json:"filename"`
	} `json:"syslog"`
}

type FileSystemSyslogOperOper struct {
	Resp []FileSystemSyslogOperOperResp `json:"resp"`
}

type FileSystemSyslogOperOperResp struct {
	Message string `json:"message"`
}

func (p *FileSystemSyslogOper) GetId() string {
	return "1"
}

func (p *FileSystemSyslogOper) getPath() string {
	return "file-system/syslog/oper"
}

func (p *FileSystemSyslogOper) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("FileSystemSyslogOper::Post")
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
	s := &FileSystemSyslogOper{}
	s.Inst.Oper = p.Inst.Oper
	s.Inst.FileContent = data
	s.Inst.File = p.Inst.File
	s.Inst.FileHandle = p.Inst.File
	_, err := axapi.NormalizeMultipartObject(http.MethodPost, p.getPath(), s.Inst.File, s.Inst.FileContent, s, headers, host, logger)
	return err
}

func (p *FileSystemSyslogOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("FileSystemSyslogOper::Get")
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
func (p *FileSystemSyslogOper) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("FileSystemSyslogOper::Put")
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

func (p *FileSystemSyslogOper) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("FileSystemSyslogOper::Delete")
	headers := axapi.GenRequestHeader(authToken)
	s := &DeleteFileSystemSyslogOper{}
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
