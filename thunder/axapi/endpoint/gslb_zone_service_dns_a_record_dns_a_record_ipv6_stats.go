package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 6_0_8-219
type GslbZoneServiceDnsARecordDnsARecordIpv6Stats struct {
	DnsARecordIpv6 string `json:"dns-a-record-ipv6"`

	Stats GslbZoneServiceDnsARecordDnsARecordIpv6StatsStats `json:"stats"`

	ServiceName string

	Zone_name string

	ServicePort string
}
type DataGslbZoneServiceDnsARecordDnsARecordIpv6Stats struct {
    DtGslbZoneServiceDnsARecordDnsARecordIpv6Stats GslbZoneServiceDnsARecordDnsARecordIpv6StatsObj `json:"dns-a-record-ipv6"`
}


type GslbZoneServiceDnsARecordDnsARecordIpv6StatsObj struct {
    Stats GslbZoneServiceDnsARecordDnsARecordIpv6StatsStats `json:"stats"`
    DnsARecordIpv6 string `json:"dns-a-record-ipv6"`
    A10URL string `json:"a10-url"`
}


type GslbZoneServiceDnsARecordDnsARecordIpv6StatsStats struct {
	Hits int `json:"hits"`
}

func (p *GslbZoneServiceDnsARecordDnsARecordIpv6Stats) GetId() string {
	return "1"
}

func (p *GslbZoneServiceDnsARecordDnsARecordIpv6Stats) getPath() string {

	return "gslb/zone/" + p.Zone_name + "/service/" + p.ServicePort + "+" + p.ServiceName + "/dns-a-record/dns-a-record-ipv6/" + p.DnsARecordIpv6 + "/stats"
}

func (p *GslbZoneServiceDnsARecordDnsARecordIpv6Stats) Get(
    authToken string,
    host string,
    instId string,
    logger *axapi.ThunderLog,
) (DataGslbZoneServiceDnsARecordDnsARecordIpv6Stats, error) {

    logger.Println("GslbZoneServiceDnsARecordDnsARecordIpv6Stats::Get")
    headers := axapi.GenRequestHeader(authToken)

    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)

    var payload DataGslbZoneServiceDnsARecordDnsARecordIpv6Stats

    if err == nil && len(axResp) > 0 {
        err = json.Unmarshal(axResp, &payload) 
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        }
    }

    return payload, err
}
