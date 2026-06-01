package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
	"io/ioutil"
	"net/http"
	"os"
)

// based on ACOS 6_0_8-219
type FileAuthPortalImageOper struct {
	Inst struct {
		Oper FileAuthPortalImageOperOper `json:"oper"`

		FileContent []byte `json:"-"`

		File string `json:"file"`

		FileHandle string `json:"file-handle"`
	} `json:"auth-portal-image"`
}
type DeleteFileAuthPortalImageOper struct {
	Inst struct {
		FileName string `json:"filename"`
	} `json:"auth-portal-image"`
}

type FileAuthPortalImageOperOper struct {
	FileList []FileAuthPortalImageOperOperFileList `json:"file-list"`
}

type FileAuthPortalImageOperOperFileList struct {
	File string `json:"file"`
}

func (p *FileAuthPortalImageOper) GetId() string {
	return "1"
}

func (p *FileAuthPortalImageOper) getPath() string {
	return "file/auth-portal-image/oper"
}

func (p *FileAuthPortalImageOper) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("FileAuthPortalImageOper::Post")
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
	s := &FileAuthPortalImageOper{}
	s.Inst.Oper = p.Inst.Oper
	s.Inst.FileContent = data
	s.Inst.File = p.Inst.File
	s.Inst.FileHandle = p.Inst.File
	_, err := axapi.NormalizeMultipartObject(http.MethodPost, p.getPath(), s.Inst.File, s.Inst.FileContent, s, headers, host, logger)
	return err
}

func (p *FileAuthPortalImageOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("FileAuthPortalImageOper::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, "file/auth-portal-image/oper", "", nil, headers, logger)
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
func (p *FileAuthPortalImageOper) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("FileAuthPortalImageOper::Put")
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

func (p *FileAuthPortalImageOper) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("FileAuthPortalImageOper::Delete")
	headers := axapi.GenRequestHeader(authToken)
	s := &DeleteFileAuthPortalImageOper{}
	s.Inst.FileName = p.Inst.File
	payloadBytes, err := json.Marshal(s)
	logger.Println(s)
	if err != nil {
		logger.Println("json.Marshal() failed with error", err)
		return err
	}
	logger.Println("payload:", string(payloadBytes))
	_, _, err = axapi.SendPost(host, "delete/auth-portal-image", payloadBytes, headers, logger)
	return err
}
