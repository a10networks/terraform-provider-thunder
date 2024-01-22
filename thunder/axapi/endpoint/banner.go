

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Banner struct {
	Inst struct {

    ExecBannerCfg BannerExecBannerCfg `json:"exec-banner-cfg"`
    LoginBannerCfg BannerLoginBannerCfg `json:"login-banner-cfg"`
    Uuid string `json:"uuid"`

	} `json:"banner"`
}


type BannerExecBannerCfg struct {
    Exec int `json:"exec"`
    ExecBanner string `json:"exec-banner"`
}


type BannerLoginBannerCfg struct {
    Login int `json:"login"`
    LoginBanner string `json:"login-banner"`
}

func (p *Banner) GetId() string{
    return "1"
}

func (p *Banner) getPath() string{
    return "banner"
}

func (p *Banner) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Banner::Post")
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

func (p *Banner) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Banner::Get")
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
func (p *Banner) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Banner::Put")
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

func (p *Banner) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Banner::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
