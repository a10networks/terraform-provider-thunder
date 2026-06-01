package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 6_0_8-219
type SlbDynamicServiceDnsStats struct {
	Stats SlbDynamicServiceDnsStatsStats `json:"stats"`
}
type DataSlbDynamicServiceDnsStats struct {
	DtSlbDynamicServiceDnsStats SlbDynamicServiceDnsStats `json:"dynamic-service-dns"`
}

type SlbDynamicServiceDnsStatsStats struct {
	N_query            int `json:"n_query"`
	N_query_err        int `json:"n_query_err"`
	N_query_drop       int `json:"n_query_drop"`
	N_resp             int `json:"n_resp"`
	N_resp_err         int `json:"n_resp_err"`
	N_resp_f_formerr   int `json:"n_resp_f_formerr"`
	N_resp_f_servfail  int `json:"n_resp_f_servfail"`
	N_resp_f_nxdomain  int `json:"n_resp_f_nxdomain"`
	N_resp_f_notimp    int `json:"n_resp_f_notimp"`
	N_resp_f_refused   int `json:"n_resp_f_refused"`
	N_resp_f_yxdomain  int `json:"n_resp_f_yxdomain"`
	N_resp_f_yxrrset   int `json:"n_resp_f_yxrrset"`
	N_resp_f_nxrrset   int `json:"n_resp_f_nxrrset"`
	N_resp_f_notauth   int `json:"n_resp_f_notauth"`
	N_resp_f_notzone   int `json:"n_resp_f_notzone"`
	N_resp_f_dsotypeni int `json:"n_resp_f_dsotypeni"`
	N_resp_f_badvers   int `json:"n_resp_f_badvers"`
	N_resp_f_badkey    int `json:"n_resp_f_badkey"`
	N_resp_f_badtime   int `json:"n_resp_f_badtime"`
	N_resp_f_badmode   int `json:"n_resp_f_badmode"`
	N_resp_f_badname   int `json:"n_resp_f_badname"`
	N_resp_f_badalg    int `json:"n_resp_f_badalg"`
	N_resp_f_badtrunc  int `json:"n_resp_f_badtrunc"`
	N_resp_f_badcookie int `json:"n_resp_f_badcookie"`
	N_resp_f_invalid   int `json:"n_resp_f_invalid"`
}

func (p *SlbDynamicServiceDnsStats) GetId() string {
	return "1"
}

func (p *SlbDynamicServiceDnsStats) getPath() string {
	return "slb/dynamic-service-dns/stats"
}

func (p *SlbDynamicServiceDnsStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbDynamicServiceDnsStats, error) {
	logger.Println("SlbDynamicServiceDnsStats::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
	var payload DataSlbDynamicServiceDnsStats
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
