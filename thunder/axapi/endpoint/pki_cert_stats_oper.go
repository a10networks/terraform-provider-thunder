

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type PkiCertStatsOper struct {
    
    Oper PkiCertStatsOperOper `json:"oper"`

}
type DataPkiCertStatsOper struct {
    DtPkiCertStatsOper PkiCertStatsOper `json:"cert-stats"`
}


type PkiCertStatsOperOper struct {
    CertCount int `json:"cert-count"`
    KeyCount int `json:"key-count"`
    CaCertCount int `json:"ca-cert-count"`
    Partition string `json:"partition"`
}

func (p *PkiCertStatsOper) GetId() string{
    return "1"
}

func (p *PkiCertStatsOper) getPath() string{
    return "pki/cert-stats/oper"
}

func (p *PkiCertStatsOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataPkiCertStatsOper,error) {
logger.Println("PkiCertStatsOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataPkiCertStatsOper
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
