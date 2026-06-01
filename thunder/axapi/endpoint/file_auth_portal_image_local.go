package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
	"io/ioutil"
	"net/http"
	"os"
)

// based on ACOS 6_0_8-219
type FileAuthPortalImageLocal struct {
	Inst struct {
		Action string `json:"action"`

		DstFile string `json:"dst-file"`

		File string `json:"file"`

		FileHandle string `json:"file-handle"`

		Uuid string `json:"uuid"`

		FileContent []byte `json:"-"`
	} `json:"auth-portal-image"`
}
type DeleteFileAuthPortalImageLocal struct {
	Inst struct {
		FileName string `json:"filename"`
	} `json:"auth-portal-image"`
}

func (p *FileAuthPortalImageLocal) GetId() string {
	return "1"
}

func (p *FileAuthPortalImageLocal) getPath() string {
	return "file/auth-portal-image"
}

func (p *FileAuthPortalImageLocal) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("FileAuthPortalImageLocal::Post")
	headers := axapi.GenRequestHeader(authToken)
	if p.Inst.Action == "delete" {
		payloadBytes, err := axapi.SerializeToJson(p)
		if err != nil {
			logger.Println("Failed to serialize struct as json", err)
			return err
		}
		logger.Println("payload:", string(payloadBytes))
		_, _, err = axapi.SendPost(host, p.getPath(), payloadBytes, headers, logger)
		return err
	}
	filePath := p.Inst.FileHandle
	if filePath == "" {
		filePath = p.Inst.File
	}
	f, err := os.Open(filePath)
	if err != nil {
		logger.Println("Failed to open a file:", err)
		return err
	}
	defer f.Close()
	data, err := ioutil.ReadAll(f)
	if err != nil {
		logger.Println("Failed to read file:", err)
		return err
	}
	s := &FileAuthPortalImageLocal{}
	s.Inst.Action = p.Inst.Action
	s.Inst.DstFile = p.Inst.DstFile
	s.Inst.File = p.Inst.File
	if p.Inst.FileHandle != "" {
		s.Inst.FileHandle = p.Inst.FileHandle
	}
	s.Inst.Uuid = p.Inst.Uuid
	s.Inst.FileContent = data
	_, err = axapi.NormalizeMultipartObject(http.MethodPost,p.getPath(),s.Inst.File,s.Inst.FileContent,s,headers,host,logger,)
	return err
}

func (p *FileAuthPortalImageLocal) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("FileAuthPortalImageLocal::Get")
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
func (p *FileAuthPortalImageLocal) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("FileAuthPortalImageLocal::Put")
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

func (p *FileAuthPortalImageLocal) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("FileAuthPortalImageLocal::Delete")
	headers := axapi.GenRequestHeader(authToken)
	s := &DeleteFileAuthPortalImageLocal{}
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
