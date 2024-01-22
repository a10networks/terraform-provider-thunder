

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type ImportGeoLocationArchive struct {
	Inst struct {

    GeoLocationArchiveFormat string `json:"geo-location-archive-format"`
    Password string `json:"password"`
    RemoteFile string `json:"remote-file"`
    UseMgmtPort int `json:"use-mgmt-port"`

	} `json:"geo-location-archive"`
}

func (p *ImportGeoLocationArchive) GetId() string{
    return "1"
}

func (p *ImportGeoLocationArchive) getPath() string{
    return "import/geo-location-archive"
}

func (p *ImportGeoLocationArchive) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("ImportGeoLocationArchive::Post")
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

func (p *ImportGeoLocationArchive) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("ImportGeoLocationArchive::Get")
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
func (p *ImportGeoLocationArchive) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("ImportGeoLocationArchive::Put")
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

func (p *ImportGeoLocationArchive) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("ImportGeoLocationArchive::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
