

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type GslbZone struct {
	Inst struct {

    Disable int `json:"disable"`
    DnsCaaRecordList []GslbZoneDnsCaaRecordList `json:"dns-caa-record-list"`
    DnsMxRecordList []GslbZoneDnsMxRecordList `json:"dns-mx-record-list"`
    DnsNsRecordList []GslbZoneDnsNsRecordList `json:"dns-ns-record-list"`
    DnsSoaRecord GslbZoneDnsSoaRecord `json:"dns-soa-record"`
    Name string `json:"name"`
    Policy string `json:"policy" dval:"default"`
    SamplingEnable []GslbZoneSamplingEnable `json:"sampling-enable"`
    ServiceList []GslbZoneServiceList `json:"service-list"`
    Template GslbZoneTemplate `json:"template"`
    Ttl int `json:"ttl" dval:"10"`
    UseServerTtl int `json:"use-server-ttl"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"zone"`
}


type GslbZoneDnsCaaRecordList struct {
    CriticalFlag int `json:"critical-flag"`
    PropertyTag string `json:"property-tag"`
    Rdata string `json:"rdata"`
    Ttl int `json:"ttl"`
    Uuid string `json:"uuid"`
    SamplingEnable []GslbZoneDnsCaaRecordListSamplingEnable `json:"sampling-enable"`
}


type GslbZoneDnsCaaRecordListSamplingEnable struct {
    Counters1 string `json:"counters1"`
}


type GslbZoneDnsMxRecordList struct {
    MxName string `json:"mx-name"`
    Priority int `json:"priority"`
    Ttl int `json:"ttl"`
    Uuid string `json:"uuid"`
    SamplingEnable []GslbZoneDnsMxRecordListSamplingEnable `json:"sampling-enable"`
}


type GslbZoneDnsMxRecordListSamplingEnable struct {
    Counters1 string `json:"counters1"`
}


type GslbZoneDnsNsRecordList struct {
    NsName string `json:"ns-name"`
    Ttl int `json:"ttl"`
    Uuid string `json:"uuid"`
    SamplingEnable []GslbZoneDnsNsRecordListSamplingEnable `json:"sampling-enable"`
}


type GslbZoneDnsNsRecordListSamplingEnable struct {
    Counters1 string `json:"counters1"`
}


type GslbZoneDnsSoaRecord struct {
    SoaName string `json:"soa-name"`
    Mail string `json:"mail"`
    Expire int `json:"expire" dval:"1209600"`
    Refresh int `json:"refresh" dval:"3600"`
    Retry int `json:"retry" dval:"900"`
    Serial int `json:"serial"`
    SoaTtl int `json:"soa-ttl"`
    External string `json:"external"`
    ExMail string `json:"ex-mail"`
    ExExpire int `json:"ex-expire" dval:"1209600"`
    ExRefresh int `json:"ex-refresh" dval:"3600"`
    ExRetry int `json:"ex-retry" dval:"900"`
    ExSerial int `json:"ex-serial"`
    ExSoaTtl int `json:"ex-soa-ttl"`
}


type GslbZoneSamplingEnable struct {
    Counters1 string `json:"counters1"`
}


type GslbZoneServiceList struct {
    ServicePort int `json:"service-port"`
    ServiceName string `json:"service-name"`
    Action string `json:"action"`
    ForwardType string `json:"forward-type"`
    Disable int `json:"disable"`
    HealthCheckGateway string `json:"health-check-gateway" dval:"enable"`
    HealthCheckPort []GslbZoneServiceListHealthCheckPort `json:"health-check-port"`
    Policy string `json:"policy"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    SamplingEnable []GslbZoneServiceListSamplingEnable `json:"sampling-enable"`
    DnsARecord GslbZoneServiceListDnsARecord `json:"dns-a-record"`
    DnsCnameRecordList []GslbZoneServiceListDnsCnameRecordList `json:"dns-cname-record-list"`
    DnsMxRecordList []GslbZoneServiceListDnsMxRecordList `json:"dns-mx-record-list"`
    DnsNsRecordList []GslbZoneServiceListDnsNsRecordList `json:"dns-ns-record-list"`
    DnsPtrRecordList []GslbZoneServiceListDnsPtrRecordList `json:"dns-ptr-record-list"`
    DnsSrvRecordList []GslbZoneServiceListDnsSrvRecordList `json:"dns-srv-record-list"`
    DnsNaptrRecordList []GslbZoneServiceListDnsNaptrRecordList `json:"dns-naptr-record-list"`
    DnsTxtRecordList []GslbZoneServiceListDnsTxtRecordList `json:"dns-txt-record-list"`
    DnsCaaRecordList []GslbZoneServiceListDnsCaaRecordList `json:"dns-caa-record-list"`
    DnsRecordList []GslbZoneServiceListDnsRecordList `json:"dns-record-list"`
    GeoLocationList []GslbZoneServiceListGeoLocationList `json:"geo-location-list"`
}


type GslbZoneServiceListHealthCheckPort struct {
    HealthCheckPort int `json:"health-check-port"`
}


type GslbZoneServiceListSamplingEnable struct {
    Counters1 string `json:"counters1"`
}


type GslbZoneServiceListDnsARecord struct {
    DnsARecordSrvList []GslbZoneServiceListDnsARecordDnsARecordSrvList `json:"dns-a-record-srv-list"`
    DnsARecordIpv4List []GslbZoneServiceListDnsARecordDnsARecordIpv4List `json:"dns-a-record-ipv4-list"`
    DnsARecordIpv6List []GslbZoneServiceListDnsARecordDnsARecordIpv6List `json:"dns-a-record-ipv6-list"`
}


type GslbZoneServiceListDnsARecordDnsARecordSrvList struct {
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


type GslbZoneServiceListDnsARecordDnsARecordIpv4List struct {
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


type GslbZoneServiceListDnsARecordDnsARecordIpv6List struct {
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


type GslbZoneServiceListDnsCnameRecordList struct {
    AliasName string `json:"alias-name"`
    AdminPreference int `json:"admin-preference" dval:"100"`
    Weight int `json:"weight" dval:"1"`
    AsBackup int `json:"as-backup"`
    Uuid string `json:"uuid"`
    SamplingEnable []GslbZoneServiceListDnsCnameRecordListSamplingEnable `json:"sampling-enable"`
}


type GslbZoneServiceListDnsCnameRecordListSamplingEnable struct {
    Counters1 string `json:"counters1"`
}


type GslbZoneServiceListDnsMxRecordList struct {
    MxName string `json:"mx-name"`
    Priority int `json:"priority"`
    Ttl int `json:"ttl"`
    Uuid string `json:"uuid"`
    SamplingEnable []GslbZoneServiceListDnsMxRecordListSamplingEnable `json:"sampling-enable"`
}


type GslbZoneServiceListDnsMxRecordListSamplingEnable struct {
    Counters1 string `json:"counters1"`
}


type GslbZoneServiceListDnsNsRecordList struct {
    NsName string `json:"ns-name"`
    Ttl int `json:"ttl"`
    Uuid string `json:"uuid"`
    SamplingEnable []GslbZoneServiceListDnsNsRecordListSamplingEnable `json:"sampling-enable"`
}


type GslbZoneServiceListDnsNsRecordListSamplingEnable struct {
    Counters1 string `json:"counters1"`
}


type GslbZoneServiceListDnsPtrRecordList struct {
    PtrName string `json:"ptr-name"`
    Ttl int `json:"ttl"`
    Uuid string `json:"uuid"`
    SamplingEnable []GslbZoneServiceListDnsPtrRecordListSamplingEnable `json:"sampling-enable"`
}


type GslbZoneServiceListDnsPtrRecordListSamplingEnable struct {
    Counters1 string `json:"counters1"`
}


type GslbZoneServiceListDnsSrvRecordList struct {
    SrvName string `json:"srv-name"`
    Port int `json:"port"`
    Priority int `json:"priority"`
    Weight int `json:"weight" dval:"10"`
    Ttl int `json:"ttl"`
    Uuid string `json:"uuid"`
    SamplingEnable []GslbZoneServiceListDnsSrvRecordListSamplingEnable `json:"sampling-enable"`
}


type GslbZoneServiceListDnsSrvRecordListSamplingEnable struct {
    Counters1 string `json:"counters1"`
}


type GslbZoneServiceListDnsNaptrRecordList struct {
    NaptrTarget string `json:"naptr-target"`
    ServiceProto string `json:"service-proto"`
    Flag string `json:"flag"`
    Order int `json:"order"`
    Preference int `json:"preference"`
    Regexp int `json:"regexp"`
    Ttl int `json:"ttl"`
    Uuid string `json:"uuid"`
    SamplingEnable []GslbZoneServiceListDnsNaptrRecordListSamplingEnable `json:"sampling-enable"`
}


type GslbZoneServiceListDnsNaptrRecordListSamplingEnable struct {
    Counters1 string `json:"counters1"`
}


type GslbZoneServiceListDnsTxtRecordList struct {
    RecordName string `json:"record-name"`
    TxtData string `json:"txt-data"`
    Ttl int `json:"ttl"`
    Uuid string `json:"uuid"`
    SamplingEnable []GslbZoneServiceListDnsTxtRecordListSamplingEnable `json:"sampling-enable"`
}


type GslbZoneServiceListDnsTxtRecordListSamplingEnable struct {
    Counters1 string `json:"counters1"`
}


type GslbZoneServiceListDnsCaaRecordList struct {
    CriticalFlag int `json:"critical-flag"`
    PropertyTag string `json:"property-tag"`
    Rdata string `json:"rdata"`
    Ttl int `json:"ttl"`
    Uuid string `json:"uuid"`
    SamplingEnable []GslbZoneServiceListDnsCaaRecordListSamplingEnable `json:"sampling-enable"`
}


type GslbZoneServiceListDnsCaaRecordListSamplingEnable struct {
    Counters1 string `json:"counters1"`
}


type GslbZoneServiceListDnsRecordList struct {
    Type int `json:"type"`
    Data string `json:"data"`
    Uuid string `json:"uuid"`
}


type GslbZoneServiceListGeoLocationList struct {
    GeoName string `json:"geo-name"`
    Alias []GslbZoneServiceListGeoLocationListAlias `json:"alias"`
    Action int `json:"action"`
    ActionType string `json:"action-type"`
    ForwardType string `json:"forward-type"`
    Policy string `json:"policy"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type GslbZoneServiceListGeoLocationListAlias struct {
    Alias string `json:"alias"`
}


type GslbZoneTemplate struct {
    Dnssec string `json:"dnssec"`
}

func (p *GslbZone) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *GslbZone) getPath() string{
    return "gslb/zone"
}

func (p *GslbZone) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("GslbZone::Post")
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

func (p *GslbZone) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("GslbZone::Get")
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
func (p *GslbZone) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("GslbZone::Put")
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

func (p *GslbZone) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("GslbZone::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
