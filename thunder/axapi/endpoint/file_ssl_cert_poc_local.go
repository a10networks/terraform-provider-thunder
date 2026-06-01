package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
	"io/ioutil"
	"net/http"
	"os"
)

// based on ACOS 6_0_8-219
type FileSslCertPocLocal struct {
	Inst struct {
		Action string `json:"action"`

		CertificateType string `json:"certificate-type"`

		DstFile string `json:"dst-file"`

		File string `json:"file"`

		FileHandle string `json:"file-handle"`

		PfxPassword string `json:"pfx-password"`

		FileContent []byte `json:"-"`
	} `json:"ssl-cert-poc"`
}
type DeleteFileSslCertPocLocal struct {
	Inst struct {
		FileName string `json:"filename"`
	} `json:"ssl-cert-poc"`
}

func (p *FileSslCertPocLocal) GetId() string {
	return "1"
}

func (p *FileSslCertPocLocal) getPath() string {
	return "file/ssl-cert-poc"
}

func (p *FileSslCertPocLocal) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("FileSslCertPocLocal::Post")
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
	s := &FileSslCertPocLocal{}
	s.Inst.Action = p.Inst.Action
	s.Inst.CertificateType = p.Inst.CertificateType
	s.Inst.DstFile = p.Inst.DstFile
	s.Inst.File = p.Inst.File
	if p.Inst.FileHandle != "" {
		s.Inst.FileHandle = p.Inst.File
	}
	s.Inst.PfxPassword = p.Inst.PfxPassword
	s.Inst.FileContent = data
	_, err = axapi.NormalizeMultipartObject(http.MethodPost, p.getPath(), s.Inst.File, s.Inst.FileContent, s, headers, host, logger)
	return err
}

func (p *FileSslCertPocLocal) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("FileSslCertPocLocal::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, "file/ssl-cert-poc/oper", "", nil, headers, logger)
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
func (p *FileSslCertPocLocal) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("FileSslCertPocLocal::Put")
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

func (p *FileSslCertPocLocal) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("FileSslCertPocLocal::Delete")
	headers := axapi.GenRequestHeader(authToken)
	s := &DeleteFileSslCertPocLocal{}
	s.Inst.FileName = p.Inst.File
	payloadBytes, err := json.Marshal(s)
	logger.Println(s)
	if err != nil {
		logger.Println("json.Marshal() failed with error", err)
		return err
	}
	logger.Println("payload:", string(payloadBytes))
	_, _, err = axapi.SendPost(host, "delete/ssl-cert-poc", payloadBytes, headers, logger)
	return err
}
