package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 6_0_8-219
type DdosDstZoneSrcIpFilteringOper struct {
	Oper DdosDstZoneSrcIpFilteringOperOper `json:"oper"`

	ZoneName string
}
type DataDdosDstZoneSrcIpFilteringOper struct {
	DtDdosDstZoneSrcIpFilteringOper DdosDstZoneSrcIpFilteringOper `json:"src-ip-filtering"`
}

type DdosDstZoneSrcIpFilteringOperOper struct {
	ClassList []DdosDstZoneSrcIpFilteringOperOperClassList `json:"class-list"`
}

type DdosDstZoneSrcIpFilteringOperOperClassList struct {
	Name string `json:"name"`
	Hit  int    `json:"hit"`
}

func (p *DdosDstZoneSrcIpFilteringOper) GetId() string {
	return "1"
}

func (p *DdosDstZoneSrcIpFilteringOper) getPath() string {

	return "ddos/dst/zone/" + p.ZoneName + "/src-ip-filtering/oper"
}

func (p *DdosDstZoneSrcIpFilteringOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosDstZoneSrcIpFilteringOper, error) {
	logger.Println("DdosDstZoneSrcIpFilteringOper::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
	var payload DataDdosDstZoneSrcIpFilteringOper
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
