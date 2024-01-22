

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type GslbZoneServiceDnsARecordDnsARecordSrv struct {
	Inst struct {

    AdminIp int `json:"admin-ip"`
    AsBackup int `json:"as-backup"`
    AsReplace int `json:"as-replace"`
    Disable int `json:"disable"`
    NoResp int `json:"no-resp"`
    Static int `json:"static"`
    Svrname string `json:"svrname"`
    Ttl int `json:"ttl"`
    Uuid string `json:"uuid"`
    Weight int `json:"weight"`
    ServiceName string 
    ServicePort string 
    Name string 

	} `json:"dns-a-record-srv"`
}

func (p *GslbZoneServiceDnsARecordDnsARecordSrv) GetId() string{
    return p.Inst.Svrname
}

func (p *GslbZoneServiceDnsARecordDnsARecordSrv) getPath() string{
    return "gslb/zone/" +p.Inst.Name + "/service/" +p.Inst.ServicePort + "+" +p.Inst.ServiceName + "/dns-a-record/dns-a-record-srv"
}

func (p *GslbZoneServiceDnsARecordDnsARecordSrv) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("GslbZoneServiceDnsARecordDnsARecordSrv::Post")
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

func (p *GslbZoneServiceDnsARecordDnsARecordSrv) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("GslbZoneServiceDnsARecordDnsARecordSrv::Get")
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
func (p *GslbZoneServiceDnsARecordDnsARecordSrv) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("GslbZoneServiceDnsARecordDnsARecordSrv::Put")
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

func (p *GslbZoneServiceDnsARecordDnsARecordSrv) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("GslbZoneServiceDnsARecordDnsARecordSrv::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
