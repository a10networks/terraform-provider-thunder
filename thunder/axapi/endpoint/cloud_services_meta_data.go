

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type CloudServicesMetaData struct {
	Inst struct {

    Action string `json:"action" dval:"disable"`
    PreventAdminPasswd int `json:"prevent-admin-passwd"`
    PreventAdminSshKey int `json:"prevent-admin-ssh-key"`
    PreventAutofill int `json:"prevent-autofill"`
    PreventBlob int `json:"prevent-blob"`
    PreventCloudService int `json:"prevent-cloud-service"`
    PreventLicense int `json:"prevent-license"`
    PreventUserOps int `json:"prevent-user-ops"`
    PreventWebservice int `json:"prevent-webservice"`
    Provider1 string `json:"provider1"`
    Uuid string `json:"uuid"`

	} `json:"meta-data"`
}

func (p *CloudServicesMetaData) GetId() string{
    return "1"
}

func (p *CloudServicesMetaData) getPath() string{
    return "cloud-services/meta-data"
}

func (p *CloudServicesMetaData) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("CloudServicesMetaData::Post")
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

func (p *CloudServicesMetaData) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("CloudServicesMetaData::Get")
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
func (p *CloudServicesMetaData) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("CloudServicesMetaData::Put")
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

func (p *CloudServicesMetaData) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("CloudServicesMetaData::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
