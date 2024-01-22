

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type ExportGeoLocationArchive struct {
	Inst struct {

    GeoLocationArchiveName string `json:"geo-location-archive-name"`
    Password string `json:"password"`
    RemoteFile string `json:"remote-file"`
    UseMgmtPort int `json:"use-mgmt-port"`

	} `json:"geo-location-archive"`
}

func (p *ExportGeoLocationArchive) GetId() string{
    return "1"
}

func (p *ExportGeoLocationArchive) getPath() string{
    return "export/geo-location-archive"
}

func (p *ExportGeoLocationArchive) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("ExportGeoLocationArchive::Post")
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

func (p *ExportGeoLocationArchive) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("ExportGeoLocationArchive::Get")
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
func (p *ExportGeoLocationArchive) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("ExportGeoLocationArchive::Put")
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

func (p *ExportGeoLocationArchive) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("ExportGeoLocationArchive::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
