

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type GslbZoneService struct {
	Inst struct {

    Action string `json:"action"`
    Disable int `json:"disable"`
    DnsARecord GslbZoneServiceDnsARecord399 `json:"dns-a-record"`
    DnsCaaRecordList []GslbZoneServiceDnsCaaRecordList `json:"dns-caa-record-list"`
    DnsCnameRecordList []GslbZoneServiceDnsCnameRecordList `json:"dns-cname-record-list"`
    DnsMxRecordList []GslbZoneServiceDnsMxRecordList `json:"dns-mx-record-list"`
    DnsNaptrRecordList []GslbZoneServiceDnsNaptrRecordList `json:"dns-naptr-record-list"`
    DnsNsRecordList []GslbZoneServiceDnsNsRecordList `json:"dns-ns-record-list"`
    DnsPtrRecordList []GslbZoneServiceDnsPtrRecordList `json:"dns-ptr-record-list"`
    DnsRecordList []GslbZoneServiceDnsRecordList `json:"dns-record-list"`
    DnsSrvRecordList []GslbZoneServiceDnsSrvRecordList `json:"dns-srv-record-list"`
    DnsTxtRecordList []GslbZoneServiceDnsTxtRecordList `json:"dns-txt-record-list"`
    ForwardType string `json:"forward-type"`
    GeoLocationList []GslbZoneServiceGeoLocationList `json:"geo-location-list"`
    HealthCheckGateway string `json:"health-check-gateway" dval:"enable"`
    HealthCheckPort []GslbZoneServiceHealthCheckPort `json:"health-check-port"`
    Policy string `json:"policy"`
    SamplingEnable []GslbZoneServiceSamplingEnable `json:"sampling-enable"`
    ServiceName string `json:"service-name"`
    ServicePort int `json:"service-port"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"service"`
}


type GslbZoneServiceDnsARecord399 struct {
    DnsARecordSrvList []GslbZoneServiceDnsARecordDnsARecordSrvList `json:"dns-a-record-srv-list"`
    DnsARecordIpv4List []GslbZoneServiceDnsARecordDnsARecordIpv4List `json:"dns-a-record-ipv4-list"`
    DnsARecordIpv6List []GslbZoneServiceDnsARecordDnsARecordIpv6List `json:"dns-a-record-ipv6-list"`
}


type GslbZoneServiceDnsARecordDnsARecordSrvList struct {
    Svrname string `json:"svrname"`
    NoResp int `json:"no-resp"`
    AsBackup int `json:"as-backup"`
    Weight int `json:"weight"`
    Ttl int `json:"ttl"`
    AsReplace int `json:"as-replace"`
    Disable int `json:"disable"`
    Static int `json:"static"`
    AdminIp int `json:"admin-ip"`
    Uuid string `json:"uuid"`
}


type GslbZoneServiceDnsARecordDnsARecordIpv4List struct {
    DnsARecordIp string `json:"dns-a-record-ip"`
    NoResp int `json:"no-resp"`
    AsBackup int `json:"as-backup"`
    Weight int `json:"weight"`
    Ttl int `json:"ttl"`
    AsReplace int `json:"as-replace"`
    Disable int `json:"disable"`
    Static int `json:"static"`
    AdminIp int `json:"admin-ip"`
    Uuid string `json:"uuid"`
}


type GslbZoneServiceDnsARecordDnsARecordIpv6List struct {
    DnsARecordIpv6 string `json:"dns-a-record-ipv6"`
    NoResp int `json:"no-resp"`
    AsBackup int `json:"as-backup"`
    Weight int `json:"weight"`
    Ttl int `json:"ttl"`
    AsReplace int `json:"as-replace"`
    Disable int `json:"disable"`
    Static int `json:"static"`
    AdminIp int `json:"admin-ip"`
    Uuid string `json:"uuid"`
}


type GslbZoneServiceDnsCaaRecordList struct {
    CriticalFlag int `json:"critical-flag"`
    PropertyTag string `json:"property-tag"`
    Rdata string `json:"rdata"`
    Ttl int `json:"ttl"`
    Uuid string `json:"uuid"`
    SamplingEnable []GslbZoneServiceDnsCaaRecordListSamplingEnable `json:"sampling-enable"`
}


type GslbZoneServiceDnsCaaRecordListSamplingEnable struct {
    Counters1 string `json:"counters1"`
}


type GslbZoneServiceDnsCnameRecordList struct {
    AliasName string `json:"alias-name"`
    AdminPreference int `json:"admin-preference" dval:"100"`
    Weight int `json:"weight" dval:"1"`
    AsBackup int `json:"as-backup"`
    Uuid string `json:"uuid"`
    SamplingEnable []GslbZoneServiceDnsCnameRecordListSamplingEnable `json:"sampling-enable"`
}


type GslbZoneServiceDnsCnameRecordListSamplingEnable struct {
    Counters1 string `json:"counters1"`
}


type GslbZoneServiceDnsMxRecordList struct {
    MxName string `json:"mx-name"`
    Priority int `json:"priority"`
    Ttl int `json:"ttl"`
    Uuid string `json:"uuid"`
    SamplingEnable []GslbZoneServiceDnsMxRecordListSamplingEnable `json:"sampling-enable"`
}


type GslbZoneServiceDnsMxRecordListSamplingEnable struct {
    Counters1 string `json:"counters1"`
}


type GslbZoneServiceDnsNaptrRecordList struct {
    NaptrTarget string `json:"naptr-target"`
    ServiceProto string `json:"service-proto"`
    Flag string `json:"flag"`
    Order int `json:"order"`
    Preference int `json:"preference"`
    Regexp int `json:"regexp"`
    Ttl int `json:"ttl"`
    Uuid string `json:"uuid"`
    SamplingEnable []GslbZoneServiceDnsNaptrRecordListSamplingEnable `json:"sampling-enable"`
}


type GslbZoneServiceDnsNaptrRecordListSamplingEnable struct {
    Counters1 string `json:"counters1"`
}


type GslbZoneServiceDnsNsRecordList struct {
    NsName string `json:"ns-name"`
    Ttl int `json:"ttl"`
    Uuid string `json:"uuid"`
    SamplingEnable []GslbZoneServiceDnsNsRecordListSamplingEnable `json:"sampling-enable"`
}


type GslbZoneServiceDnsNsRecordListSamplingEnable struct {
    Counters1 string `json:"counters1"`
}


type GslbZoneServiceDnsPtrRecordList struct {
    PtrName string `json:"ptr-name"`
    Ttl int `json:"ttl"`
    Uuid string `json:"uuid"`
    SamplingEnable []GslbZoneServiceDnsPtrRecordListSamplingEnable `json:"sampling-enable"`
}


type GslbZoneServiceDnsPtrRecordListSamplingEnable struct {
    Counters1 string `json:"counters1"`
}


type GslbZoneServiceDnsRecordList struct {
    Type int `json:"type"`
    Data string `json:"data"`
    Uuid string `json:"uuid"`
}


type GslbZoneServiceDnsSrvRecordList struct {
    SrvName string `json:"srv-name"`
    Port int `json:"port"`
    Priority int `json:"priority"`
    Weight int `json:"weight" dval:"10"`
    Ttl int `json:"ttl"`
    Uuid string `json:"uuid"`
    SamplingEnable []GslbZoneServiceDnsSrvRecordListSamplingEnable `json:"sampling-enable"`
}


type GslbZoneServiceDnsSrvRecordListSamplingEnable struct {
    Counters1 string `json:"counters1"`
}


type GslbZoneServiceDnsTxtRecordList struct {
    RecordName string `json:"record-name"`
    TxtData string `json:"txt-data"`
    Ttl int `json:"ttl"`
    Uuid string `json:"uuid"`
    SamplingEnable []GslbZoneServiceDnsTxtRecordListSamplingEnable `json:"sampling-enable"`
}


type GslbZoneServiceDnsTxtRecordListSamplingEnable struct {
    Counters1 string `json:"counters1"`
}


type GslbZoneServiceGeoLocationList struct {
    GeoName string `json:"geo-name"`
    Alias []GslbZoneServiceGeoLocationListAlias `json:"alias"`
    Action int `json:"action"`
    ActionType string `json:"action-type"`
    ForwardType string `json:"forward-type"`
    Policy string `json:"policy"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type GslbZoneServiceGeoLocationListAlias struct {
    Alias string `json:"alias"`
}


type GslbZoneServiceHealthCheckPort struct {
    HealthCheckPort int `json:"health-check-port"`
}


type GslbZoneServiceSamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *GslbZoneService) GetId() string{
    return strconv.Itoa(p.Inst.ServicePort)+"+"+url.QueryEscape(p.Inst.ServiceName)
}

func (p *GslbZoneService) getPath() string{
    return "gslb/zone/" +p.Inst.Name + "/service"
}

func (p *GslbZoneService) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("GslbZoneService::Post")
    headers := axapi.GenRequestHeader(authToken)
        payloadBytes, err := axapi.SerializeToJson(p)
        if err != nil {
            logger.Println("Failed to serialize struct as json", err)
            return err
        }
        logger.Println("payload:", string(payloadBytes))
        _, _, err = axapi.SendPost(host, p.getPath(), payloadBytes, headers, logger)
    return err
}

func (p *GslbZoneService) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("GslbZoneService::Get")
    headers := axapi.GenRequestHeader(authToken)
        _, axResp, err := axapi.SendGet(host, p.getPath(), instId, nil, headers, logger)
    if err == nil {
        if len(axResp) > 0{
        err = json.Unmarshal(axResp, &p)
        }
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        }
    }
    return err
}
func (p *GslbZoneService) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("GslbZoneService::Put")
    headers := axapi.GenRequestHeader(authToken)
    payloadBytes, err := axapi.SerializeToJson(p)
    if err != nil {
        logger.Println("Failed to serialize struct as json", err)
        return err
    }
    logger.Println("payload: " + string(payloadBytes))
    _, _, err = axapi.SendPut(host, p.getPath(), p.GetId(), payloadBytes, headers, logger)
    return err
}

func (p *GslbZoneService) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("GslbZoneService::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
