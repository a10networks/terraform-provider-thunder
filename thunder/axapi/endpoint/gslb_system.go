

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type GslbSystem struct {
	Inst struct {

    AgeInterval int `json:"age-interval" dval:"10"`
    GeoLocationIana int `json:"geo-location-iana" dval:"1"`
    GslbGroup int `json:"gslb-group"`
    GslbLoadFileList []GslbSystemGslbLoadFileList `json:"gslb-load-file-list"`
    GslbServiceIp int `json:"gslb-service-ip"`
    GslbSite int `json:"gslb-site"`
    Hostname int `json:"hostname"`
    IpTtl int `json:"ip-ttl"`
    Module int `json:"module"`
    SlbDevice int `json:"slb-device"`
    SlbServer int `json:"slb-server"`
    SlbVirtualServer int `json:"slb-virtual-server"`
    Ttl int `json:"ttl" dval:"300"`
    Uuid string `json:"uuid"`
    Wait int `json:"wait"`

	} `json:"system"`
}


type GslbSystemGslbLoadFileList struct {
    GeoLocationLoadFilename string `json:"geo-location-load-filename"`
    TemplateName string `json:"template-name"`
}

func (p *GslbSystem) GetId() string{
    return "1"
}

func (p *GslbSystem) getPath() string{
    return "gslb/system"
}

func (p *GslbSystem) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("GslbSystem::Post")
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

func (p *GslbSystem) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("GslbSystem::Get")
    headers := axapi.GenRequestHeader(authToken)
        _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
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
func (p *GslbSystem) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("GslbSystem::Put")
    headers := axapi.GenRequestHeader(authToken)
    payloadBytes, err := axapi.SerializeToJson(p)
    if err != nil {
        logger.Println("Failed to serialize struct as json", err)
        return err
    }
    logger.Println("payload: " + string(payloadBytes))
    _, _, err = axapi.SendPut(host, p.getPath(), "", payloadBytes, headers, logger)
    return err
}

func (p *GslbSystem) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("GslbSystem::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
