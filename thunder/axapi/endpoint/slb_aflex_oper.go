package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 7_0_2-102
type SlbAflexOper struct {
	Oper SlbAflexOperOper `json:"oper"`
}
type DataSlbAflexOper struct {
	DtSlbAflexOper SlbAflexOper `json:"aflex"`
}

type SlbAflexOperOper struct {
	Name             string                       `json:"name"`
	Exact_match      int                          `json:"exact_match"`
	Event            string                       `json:"event"`
	Filter_debug     int                          `json:"filter_debug"`
	Substring        int                          `json:"substring"`
	AflexFileSizeMax int                          `json:"aflex-file-size-max"`
	FileList         []SlbAflexOperOperFileList   `json:"file-list"`
	ThreadList       []SlbAflexOperOperThreadList `json:"thread-list"`
	AflexCount       int                          `json:"aflex-count"`
}

type SlbAflexOperOperFileList struct {
	File        string                              `json:"file"`
	Syntax      string                              `json:"syntax"`
	Vport_bound string                              `json:"vport_bound"`
	VportList   []SlbAflexOperOperFileListVportList `json:"vport-list"`
	Events      []SlbAflexOperOperFileListEvents    `json:"events"`
}

type SlbAflexOperOperFileListVportList struct {
	Vserver string `json:"vserver"`
	Port    int    `json:"port"`
}

type SlbAflexOperOperFileListEvents struct {
	EventType         string `json:"event-type"`
	TotalExecutions   int    `json:"total-executions"`
	Failures          int    `json:"failures"`
	Aborts            int    `json:"aborts"`
	Exceed_time_limit int    `json:"exceed_time_limit"`
}

type SlbAflexOperOperThreadList struct {
	Thread    int                                   `json:"thread"`
	ErrLine   int                                   `json:"err-line"`
	LastMsg   string                                `json:"last-msg"`
	ErrorList []SlbAflexOperOperThreadListErrorList `json:"error-list"`
}

type SlbAflexOperOperThreadListErrorList struct {
	EventName  string `json:"event-name"`
	CmdName    string `json:"cmd-name"`
	ErrMsg     string `json:"err-msg"`
	LineNumber int    `json:"line-number"`
	FileName   string `json:"file-name"`
}

func (p *SlbAflexOper) GetId() string {
	return "1"
}

func (p *SlbAflexOper) getPath() string {
	return "slb/aflex/oper"
}

func (p *SlbAflexOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbAflexOper, error) {
	logger.Println("SlbAflexOper::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
	var payload DataSlbAflexOper
	if err == nil {
		if len(axResp) > 0 {
			err = json.Unmarshal(axResp, &p)
		}
		if err != nil {
			logger.Println("json.Unmarshal() failed with error", err)
		}
	}
	return payload, err
}
