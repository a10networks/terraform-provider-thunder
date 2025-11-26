package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
	"io/ioutil"
	"net/http"
	"os"
)

// based on ACOS 7_0_2-102
type FileSystemVarlog struct {
	Inst struct {
		Filename string `json:"filename"`

		Rm int `json:"rm"`

		Uuid string `json:"uuid"`

		FileContent []byte `json:"-"`

		File string `json:"file"`

		FileHandle string `json:"file-handle"`
	} `json:"varlog"`
}
type DeleteFileSystemVarlog struct {
	Inst struct {
		FileName string `json:"filename"`
	} `json:"varlog"`
}

func (p *FileSystemVarlog) GetId() string {
	return "1"
}

func (p *FileSystemVarlog) getPath() string {
	return "file-system/varlog"
}

func (p *FileSystemVarlog) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("FileSystemVarlog::Post")
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
	s := &FileSystemVarlog{}
	s.Inst.Filename = p.Inst.Filename
	s.Inst.Rm = p.Inst.Rm
	s.Inst.Uuid = p.Inst.Uuid
	s.Inst.FileContent = data
	s.Inst.File = p.Inst.File
	s.Inst.FileHandle = p.Inst.File
	_, err := axapi.NormalizeMultipartObject(http.MethodPost, p.getPath(), s.Inst.File, s.Inst.FileContent, s, headers, host, logger)
	return err
}

func (p *FileSystemVarlog) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("FileSystemVarlog::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, "file/varlog/oper", "", nil, headers, logger)
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
func (p *FileSystemVarlog) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("FileSystemVarlog::Put")
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

func (p *FileSystemVarlog) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("FileSystemVarlog::Delete")
	headers := axapi.GenRequestHeader(authToken)
	s := &DeleteFileSystemVarlog{}
	s.Inst.FileName = p.Inst.File
	payloadBytes, err := json.Marshal(s)
	logger.Println(s)
	if err != nil {
		logger.Println("json.Marshal() failed with error", err)
		return err
	}
	logger.Println("payload:", string(payloadBytes))
	_, _, err = axapi.SendPost(host, "delete/varlog", payloadBytes, headers, logger)
	return err
}
