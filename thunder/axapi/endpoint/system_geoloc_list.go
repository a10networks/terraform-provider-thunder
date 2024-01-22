

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemGeolocList struct {
	Inst struct {

    ExcludeGeolocNameList []SystemGeolocListExcludeGeolocNameList `json:"exclude-geoloc-name-list"`
    IncludeGeolocNameList []SystemGeolocListIncludeGeolocNameList `json:"include-geoloc-name-list"`
    Name string `json:"name"`
    SamplingEnable []SystemGeolocListSamplingEnable `json:"sampling-enable"`
    Shared int `json:"shared"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"geoloc-list"`
}


type SystemGeolocListExcludeGeolocNameList struct {
    ExcludeGeolocNameVal string `json:"exclude-geoloc-name-val"`
}


type SystemGeolocListIncludeGeolocNameList struct {
    IncludeGeolocNameVal string `json:"include-geoloc-name-val"`
}


type SystemGeolocListSamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *SystemGeolocList) GetId() string{
    return p.Inst.Name
}

func (p *SystemGeolocList) getPath() string{
    return "system/geoloc-list"
}

func (p *SystemGeolocList) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SystemGeolocList::Post")
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

func (p *SystemGeolocList) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SystemGeolocList::Get")
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
func (p *SystemGeolocList) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SystemGeolocList::Put")
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

func (p *SystemGeolocList) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SystemGeolocList::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
