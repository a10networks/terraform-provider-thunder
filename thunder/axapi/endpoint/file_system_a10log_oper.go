package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
	"io/ioutil"
	"net/http"
	"os"
)

// based on ACOS 6_0_8-219
type FileSystemA10logOper struct {
	Inst struct {
		Oper FileSystemA10logOperOper `json:"oper"`

		FileContent []byte `json:"-"`

		File string `json:"file"`

		FileHandle string `json:"file-handle"`
	} `json:"a10log"`
}
type DeleteFileSystemA10logOper struct {
	Inst struct {
		FileName string `json:"filename"`
	} `json:"a10log"`
}

type FileSystemA10logOperOper struct {
	Resp []FileSystemA10logOperOperResp `json:"resp"`
}

type FileSystemA10logOperOperResp struct {
	Message string `json:"message"`
}

func (p *FileSystemA10logOper) GetId() string {
	return "1"
}

func (p *FileSystemA10logOper) getPath() string {
	return "file-system/a10log/oper"
}

func (p *FileSystemA10logOper) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("FileSystemA10logOper::Post")
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
	s := &FileSystemA10logOper{}
	s.Inst.Oper = p.Inst.Oper
	s.Inst.FileContent = data
	s.Inst.File = p.Inst.File
	s.Inst.FileHandle = p.Inst.File
	_, err := axapi.NormalizeMultipartObject(http.MethodPost, p.getPath(), s.Inst.File, s.Inst.FileContent, s, headers, host, logger)
	return err
}

func (p *FileSystemA10logOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("FileSystemA10logOper::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, "file/a10log/oper", "", nil, headers, logger)
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
func (p *FileSystemA10logOper) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("FileSystemA10logOper::Put")
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

func (p *FileSystemA10logOper) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("FileSystemA10logOper::Delete")
	headers := axapi.GenRequestHeader(authToken)
	s := &DeleteFileSystemA10logOper{}
	s.Inst.FileName = p.Inst.File
	payloadBytes, err := json.Marshal(s)
	logger.Println(s)
	if err != nil {
		logger.Println("json.Marshal() failed with error", err)
		return err
	}
	logger.Println("payload:", string(payloadBytes))
	_, _, err = axapi.SendPost(host, "delete/a10log", payloadBytes, headers, logger)
	return err
}
