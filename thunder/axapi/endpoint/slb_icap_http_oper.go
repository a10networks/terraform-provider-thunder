

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbIcapHttpOper struct {
    
    Oper SlbIcapHttpOperOper `json:"oper"`

}
type DataSlbIcapHttpOper struct {
    DtSlbIcapHttpOper SlbIcapHttpOper `json:"icap_http"`
}


type SlbIcapHttpOperOper struct {
    Icap_httpCpuList []SlbIcapHttpOperOperIcap_httpCpuList `json:"icap_http-cpu-list"`
    CpuCount int `json:"cpu-count"`
}


type SlbIcapHttpOperOperIcap_httpCpuList struct {
    Status_2xx int `json:"status_2xx"`
    Status_200 int `json:"status_200"`
    Status_201 int `json:"status_201"`
    Status_202 int `json:"status_202"`
    Status_203 int `json:"status_203"`
    Status_204 int `json:"status_204"`
    Status_205 int `json:"status_205"`
    Status_206 int `json:"status_206"`
    Status_207 int `json:"status_207"`
    Status_1xx int `json:"status_1xx"`
    Status_100 int `json:"status_100"`
    Status_101 int `json:"status_101"`
    Status_102 int `json:"status_102"`
    Status_3xx int `json:"status_3xx"`
    Status_300 int `json:"status_300"`
    Status_301 int `json:"status_301"`
    Status_302 int `json:"status_302"`
    Status_303 int `json:"status_303"`
    Status_304 int `json:"status_304"`
    Status_305 int `json:"status_305"`
    Status_306 int `json:"status_306"`
    Status_307 int `json:"status_307"`
    Status_4xx int `json:"status_4xx"`
    Status_400 int `json:"status_400"`
    Status_401 int `json:"status_401"`
    Status_402 int `json:"status_402"`
    Status_403 int `json:"status_403"`
    Status_404 int `json:"status_404"`
    Status_405 int `json:"status_405"`
    Status_406 int `json:"status_406"`
    Status_407 int `json:"status_407"`
    Status_408 int `json:"status_408"`
    Status_409 int `json:"status_409"`
    Status_410 int `json:"status_410"`
    Status_411 int `json:"status_411"`
    Status_412 int `json:"status_412"`
    Status_413 int `json:"status_413"`
    Status_414 int `json:"status_414"`
    Status_415 int `json:"status_415"`
    Status_416 int `json:"status_416"`
    Status_417 int `json:"status_417"`
    Status_418 int `json:"status_418"`
    Status_422 int `json:"status_422"`
    Status_423 int `json:"status_423"`
    Status_424 int `json:"status_424"`
    Status_425 int `json:"status_425"`
    Status_426 int `json:"status_426"`
    Status_449 int `json:"status_449"`
    Status_450 int `json:"status_450"`
    Status_5xx int `json:"status_5xx"`
    Status_500 int `json:"status_500"`
    Status_501 int `json:"status_501"`
    Status_502 int `json:"status_502"`
    Status_503 int `json:"status_503"`
    Status_504 int `json:"status_504"`
    Status_505 int `json:"status_505"`
    Status_506 int `json:"status_506"`
    Status_507 int `json:"status_507"`
    Status_508 int `json:"status_508"`
    Status_509 int `json:"status_509"`
    Status_510 int `json:"status_510"`
    Status_6xx int `json:"status_6xx"`
}

func (p *SlbIcapHttpOper) GetId() string{
    return "1"
}

func (p *SlbIcapHttpOper) getPath() string{
    return "slb/icap_http/oper"
}

func (p *SlbIcapHttpOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbIcapHttpOper,error) {
logger.Println("SlbIcapHttpOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbIcapHttpOper
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
