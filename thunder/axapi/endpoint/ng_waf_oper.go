

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type NgWafOper struct {
    
    Cpu NgWafOperCpu `json:"cpu"`
    CustomPage NgWafOperCustomPage `json:"custom-page"`
    CustomSignals NgWafOperCustomSignals `json:"custom-signals"`
    Oper NgWafOperOper `json:"oper"`
    Status NgWafOperStatus `json:"status"`

}
type DataNgWafOper struct {
    DtNgWafOper NgWafOper `json:"ng-waf"`
}


type NgWafOperCpu struct {
    Oper NgWafOperCpuOper `json:"oper"`
}


type NgWafOperCpuOper struct {
    NumberOfCpus int `json:"number-of-cpus"`
    CpuInfo []NgWafOperCpuOperCpuInfo `json:"cpu-info"`
}


type NgWafOperCpuOperCpuInfo struct {
    CpuId int `json:"cpu-id"`
    Sec1 int `json:"sec1"`
    Sec5 int `json:"sec5"`
    Sec10 int `json:"sec10"`
    Sec30 int `json:"sec30"`
    Sec60 int `json:"sec60"`
}


type NgWafOperCustomPage struct {
    Oper NgWafOperCustomPageOper `json:"oper"`
}


type NgWafOperCustomPageOper struct {
    FileList []NgWafOperCustomPageOperFileList `json:"file-list"`
}


type NgWafOperCustomPageOperFileList struct {
    File string `json:"file"`
    Size int `json:"size"`
}


type NgWafOperCustomSignals struct {
    Oper NgWafOperCustomSignalsOper `json:"oper"`
}


type NgWafOperCustomSignalsOper struct {
    SignalList []NgWafOperCustomSignalsOperSignalList `json:"signal-list"`
}


type NgWafOperCustomSignalsOperSignalList struct {
    Signal string `json:"signal"`
}


type NgWafOperOper struct {
    NgwafStatsList []NgWafOperOperNgwafStatsList `json:"ngwaf-stats-list"`
    Vserver string `json:"vserver"`
    Vport string `json:"vport"`
    CacheVserver string `json:"cache-vserver"`
    CacheVport string `json:"cache-vport"`
    ClearAll int `json:"clear-all"`
}


type NgWafOperOperNgwafStatsList struct {
    Type string `json:"type"`
    Name string `json:"name"`
    Count1 int `json:"count1"`
}


type NgWafOperStatus struct {
    Oper NgWafOperStatusOper `json:"oper"`
}


type NgWafOperStatusOper struct {
    NgwafVersion string `json:"ngwaf-version"`
    PartitionList []NgWafOperStatusOperPartitionList `json:"partition-list"`
}


type NgWafOperStatusOperPartitionList struct {
    PartitionName string `json:"partition-name"`
    Status string `json:"status"`
    AgentName string `json:"agent-name"`
    AccessKeyId string `json:"access-key-id"`
    SecretAccessKey string `json:"secret-access-key"`
    CacheEntries int `json:"cache-entries"`
    TrackedCustomSignal int `json:"tracked-custom-signal"`
}

func (p *NgWafOper) GetId() string{
    return "1"
}

func (p *NgWafOper) getPath() string{
    return "ng-waf/oper"
}

func (p *NgWafOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataNgWafOper,error) {
logger.Println("NgWafOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataNgWafOper
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
