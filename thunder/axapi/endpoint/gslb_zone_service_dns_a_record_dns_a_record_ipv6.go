

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type GslbZoneServiceDnsARecordDnsARecordIpv6 struct {
	Inst struct {

    AdminIp int `json:"admin-ip"`
    AsBackup int `json:"as-backup"`
    AsReplace int `json:"as-replace"`
    Disable int `json:"disable"`
    DnsARecordIpv6 string `json:"dns-a-record-ipv6"`
    NoResp int `json:"no-resp"`
    Static int `json:"static"`
    Ttl int `json:"ttl"`
    Uuid string `json:"uuid"`
    Weight int `json:"weight"`
    ServiceName string 
    ServicePort string 
    Name string 

	} `json:"dns-a-record-ipv6"`
}

func (p *GslbZoneServiceDnsARecordDnsARecordIpv6) GetId() string{
    return p.Inst.DnsARecordIpv6
}

func (p *GslbZoneServiceDnsARecordDnsARecordIpv6) getPath() string{
    return "gslb/zone/" +p.Inst.Name + "/service/" +p.Inst.ServicePort + "+" +p.Inst.ServiceName + "/dns-a-record/dns-a-record-ipv6"
}

func (p *GslbZoneServiceDnsARecordDnsARecordIpv6) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("GslbZoneServiceDnsARecordDnsARecordIpv6::Post")
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

func (p *GslbZoneServiceDnsARecordDnsARecordIpv6) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("GslbZoneServiceDnsARecordDnsARecordIpv6::Get")
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
func (p *GslbZoneServiceDnsARecordDnsARecordIpv6) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("GslbZoneServiceDnsARecordDnsARecordIpv6::Put")
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

func (p *GslbZoneServiceDnsARecordDnsARecordIpv6) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("GslbZoneServiceDnsARecordDnsARecordIpv6::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
