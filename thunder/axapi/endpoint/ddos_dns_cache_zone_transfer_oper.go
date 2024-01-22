

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDnsCacheZoneTransferOper struct {
    
    Oper DdosDnsCacheZoneTransferOperOper `json:"oper"`
    Name string 

}
type DataDdosDnsCacheZoneTransferOper struct {
    DtDdosDnsCacheZoneTransferOper DdosDnsCacheZoneTransferOper `json:"zone-transfer"`
}


type DdosDnsCacheZoneTransferOperOper struct {
    ZoneTransferStatusList []DdosDnsCacheZoneTransferOperOperZoneTransferStatusList `json:"zone-transfer-status-list"`
    ZoneName string `json:"zone-name"`
    SflowSourceId string `json:"sflow-source-id"`
    LocalIp string `json:"local-ip"`
    RemoteIp string `json:"remote-ip"`
    EstimatedNextUpdate string `json:"estimated-next-update"`
    ZoneTransferHistoryList []DdosDnsCacheZoneTransferOperOperZoneTransferHistoryList `json:"zone-transfer-history-list"`
    ZoneTransferStatistics []DdosDnsCacheZoneTransferOperOperZoneTransferStatistics `json:"zone-transfer-statistics"`
    ZtsSflowSourceId string `json:"zts-sflow-source-id"`
    Status string `json:"status"`
    Zone string `json:"zone"`
    Statistics int `json:"statistics"`
    ZtStatistics int `json:"zt-statistics"`
    DebugMode int `json:"debug-mode"`
}


type DdosDnsCacheZoneTransferOperOperZoneTransferStatusList struct {
    ZoneName string `json:"zone-name"`
    SflowSourceId string `json:"sflow-source-id"`
    LastUpdate string `json:"last-update"`
    LastCompleteUpdate string `json:"last-complete-update"`
    LastCompleteSerial string `json:"last-complete-serial"`
    EstimatedNextUpdate string `json:"estimated-next-update"`
}


type DdosDnsCacheZoneTransferOperOperZoneTransferHistoryList struct {
    UpdateStatus string `json:"update-status"`
    ZoneTransferResult string `json:"zone-transfer-result"`
    ZoneTransferBeginTime string `json:"zone-transfer-begin-time"`
    ZoneTransferEndTime string `json:"zone-transfer-end-time"`
    TcpConnectionBeginTime string `json:"tcp-connection-begin-time"`
    TcpConnectionEndTime string `json:"tcp-connection-end-time"`
    SerialNumber string `json:"serial-number"`
    DnsMessageProcessed int `json:"dns-message-processed"`
    RecordsProcessed int `json:"records-processed"`
    DnsMessagePendingProcessed int `json:"dns-message-pending-processed"`
    TotalFailure string `json:"total-failure"`
}


type DdosDnsCacheZoneTransferOperOperZoneTransferStatistics struct {
    StatsName string `json:"stats-name"`
    StatsCount int `json:"stats-count"`
}

func (p *DdosDnsCacheZoneTransferOper) GetId() string{
    return "1"
}

func (p *DdosDnsCacheZoneTransferOper) getPath() string{
    
    return "ddos/dns-cache/" +p.Name + "/zone-transfer/oper"
}

func (p *DdosDnsCacheZoneTransferOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosDnsCacheZoneTransferOper,error) {
logger.Println("DdosDnsCacheZoneTransferOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosDnsCacheZoneTransferOper
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
