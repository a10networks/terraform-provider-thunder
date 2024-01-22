

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDnsCacheOper struct {
    
    Name string `json:"name"`
    Oper DdosDnsCacheOperOper `json:"oper"`
    ZoneTransfer DdosDnsCacheOperZoneTransfer `json:"zone-transfer"`

}
type DataDdosDnsCacheOper struct {
    DtDdosDnsCacheOper DdosDnsCacheOper `json:"dns-cache"`
}


type DdosDnsCacheOperOper struct {
    DomainEntries []DdosDnsCacheOperOperDomainEntries `json:"domain-entries"`
    ResponseStatus string `json:"response-status"`
    ResponseFlag string `json:"response-flag"`
    AnswerSectionRecordCount int `json:"answer-section-record-count"`
    AnswerSectionSize int `json:"answer-section-size"`
    AuthoritySectionRecordCount int `json:"authority-section-record-count"`
    AuthoritySectionSize int `json:"authority-section-size"`
    AdditionalSectionRecordCount int `json:"additional-section-record-count"`
    AdditionalSectionSize int `json:"additional-section-size"`
    AnswerSection []DdosDnsCacheOperOperAnswerSection `json:"answer-section"`
    AuthoritativeSection []DdosDnsCacheOperOperAuthoritativeSection `json:"authoritative-section"`
    AdditionalSection []DdosDnsCacheOperOperAdditionalSection `json:"additional-section"`
    AllCachedFqdn int `json:"all-cached-fqdn"`
    CachedFqdnName string `json:"cached-fqdn-name"`
    RecordType string `json:"record-type"`
    DebugMode int `json:"debug-mode"`
}


type DdosDnsCacheOperOperDomainEntries struct {
    FqdnName string `json:"fqdn-name"`
    FqdnManualOverrideAction string `json:"fqdn-manual-override-action"`
    WildCardNode string `json:"wild-card-node"`
    DelegationNode string `json:"delegation-node"`
    EmptyNonTerminalNode string `json:"empty-non-terminal-node"`
    RecordTypes string `json:"record-types"`
}


type DdosDnsCacheOperOperAnswerSection struct {
    RecordDomainName string `json:"record-domain-name"`
    RecordType string `json:"record-type"`
    RecordClass string `json:"record-class"`
    RecordTtl int `json:"record-ttl"`
    RecordData string `json:"record-data"`
}


type DdosDnsCacheOperOperAuthoritativeSection struct {
    RecordDomainName string `json:"record-domain-name"`
    RecordType string `json:"record-type"`
    RecordClass string `json:"record-class"`
    RecordTtl int `json:"record-ttl"`
    RecordData string `json:"record-data"`
}


type DdosDnsCacheOperOperAdditionalSection struct {
    RecordDomainName string `json:"record-domain-name"`
    RecordType string `json:"record-type"`
    RecordClass string `json:"record-class"`
    RecordTtl int `json:"record-ttl"`
    RecordData string `json:"record-data"`
}


type DdosDnsCacheOperZoneTransfer struct {
    Oper DdosDnsCacheOperZoneTransferOper `json:"oper"`
}


type DdosDnsCacheOperZoneTransferOper struct {
    ZoneTransferStatusList []DdosDnsCacheOperZoneTransferOperZoneTransferStatusList `json:"zone-transfer-status-list"`
    ZoneName string `json:"zone-name"`
    SflowSourceId string `json:"sflow-source-id"`
    LocalIp string `json:"local-ip"`
    RemoteIp string `json:"remote-ip"`
    EstimatedNextUpdate string `json:"estimated-next-update"`
    ZoneTransferHistoryList []DdosDnsCacheOperZoneTransferOperZoneTransferHistoryList `json:"zone-transfer-history-list"`
    ZoneTransferStatistics []DdosDnsCacheOperZoneTransferOperZoneTransferStatistics `json:"zone-transfer-statistics"`
    ZtsSflowSourceId string `json:"zts-sflow-source-id"`
    Status string `json:"status"`
    Zone string `json:"zone"`
    Statistics int `json:"statistics"`
    ZtStatistics int `json:"zt-statistics"`
    DebugMode int `json:"debug-mode"`
}


type DdosDnsCacheOperZoneTransferOperZoneTransferStatusList struct {
    ZoneName string `json:"zone-name"`
    SflowSourceId string `json:"sflow-source-id"`
    LastUpdate string `json:"last-update"`
    LastCompleteUpdate string `json:"last-complete-update"`
    LastCompleteSerial string `json:"last-complete-serial"`
    EstimatedNextUpdate string `json:"estimated-next-update"`
}


type DdosDnsCacheOperZoneTransferOperZoneTransferHistoryList struct {
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


type DdosDnsCacheOperZoneTransferOperZoneTransferStatistics struct {
    StatsName string `json:"stats-name"`
    StatsCount int `json:"stats-count"`
}

func (p *DdosDnsCacheOper) GetId() string{
    return "1"
}

func (p *DdosDnsCacheOper) getPath() string{
    
    return "ddos/dns-cache/"+p.Name+"/oper"
}

func (p *DdosDnsCacheOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosDnsCacheOper,error) {
logger.Println("DdosDnsCacheOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosDnsCacheOper
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
