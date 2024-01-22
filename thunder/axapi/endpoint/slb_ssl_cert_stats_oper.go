

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbSslCertStatsOper struct {
    
    Oper SlbSslCertStatsOperOper `json:"oper"`

}
type DataSlbSslCertStatsOper struct {
    DtSlbSslCertStatsOper SlbSslCertStatsOper `json:"ssl-cert-stats"`
}


type SlbSslCertStatsOperOper struct {
    CertCount int `json:"cert-count"`
    KeyCount int `json:"key-count"`
    CaCertCount int `json:"ca-cert-count"`
    Partition string `json:"partition"`
}

func (p *SlbSslCertStatsOper) GetId() string{
    return "1"
}

func (p *SlbSslCertStatsOper) getPath() string{
    return "slb/ssl-cert-stats/oper"
}

func (p *SlbSslCertStatsOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbSslCertStatsOper,error) {
logger.Println("SlbSslCertStatsOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbSslCertStatsOper
    if err == nil {
        if len(axResp) > 0{
        err = json.Unmarshal(axResp, &p)
        }
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        }
    }
    return payload,err
}
