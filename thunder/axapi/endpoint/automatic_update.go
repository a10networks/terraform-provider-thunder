package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 6_0_8-219
type AutomaticUpdate struct {
	Inst struct {
		CheckNow AutomaticUpdateCheckNow70 `json:"check-now"`

		Checknow AutomaticUpdateChecknow71 `json:"checknow"`

		ConfigList []AutomaticUpdateConfigList `json:"config-list"`

		GlmSourceUrl string `json:"glm-source-url"`

		Info AutomaticUpdateInfo72 `json:"info"`

		ProxyServer AutomaticUpdateProxyServer73 `json:"proxy-server"`

		Reset AutomaticUpdateReset74 `json:"reset"`

		Revert AutomaticUpdateRevert75 `json:"revert"`

		UseMgmtPort int `json:"use-mgmt-port"`

		Uuid string `json:"uuid"`
	} `json:"automatic-update"`
}

type AutomaticUpdateCheckNow70 struct {
	FeatureName       string `json:"feature-name"`
	ProdVer           string `json:"prod-ver"`
	FromStagingServer int    `json:"from-staging-server"`
	StageVer          string `json:"stage-ver"`
}

type AutomaticUpdateChecknow71 struct {
	Uuid string `json:"uuid"`
}

type AutomaticUpdateConfigList struct {
	FeatureName      string `json:"feature-name"`
	Debug            int    `json:"debug"`
	DisableSslVerify int    `json:"disable-ssl-verify"`
	Schedule         int    `json:"schedule"`
	Weekly           int    `json:"weekly"`
	WeekDay          string `json:"week-day"`
	WeekTime         string `json:"week-time"`
	Daily            int    `json:"daily"`
	DayTime          string `json:"day-time"`
	Uuid             string `json:"uuid"`
}

type AutomaticUpdateInfo72 struct {
	Uuid string `json:"uuid"`
}

type AutomaticUpdateProxyServer73 struct {
	ProxyHost    string `json:"proxy-host"`
	HttpsPort    int    `json:"https-port"`
	AuthType     string `json:"auth-type" dval:"ntlm"`
	Domain       string `json:"domain"`
	Username     string `json:"username"`
	Password     int    `json:"password"`
	SecretString string `json:"secret-string"`
	Encrypted    string `json:"encrypted"`
	Uuid         string `json:"uuid"`
}

type AutomaticUpdateReset74 struct {
	FeatureName string `json:"feature-name"`
}

type AutomaticUpdateRevert75 struct {
	FeatureName string `json:"feature-name"`
}

func (p *AutomaticUpdate) GetId() string {
	return "1"
}

func (p *AutomaticUpdate) getPath() string {
	return "automatic-update"
}

func (p *AutomaticUpdate) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("AutomaticUpdate::Post")
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

func (p *AutomaticUpdate) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("AutomaticUpdate::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
	if err == nil {
		if len(axResp) > 0 {
			err = json.Unmarshal(axResp, &p)
		}
		if err != nil {
			logger.Println("json.Unmarshal() failed with error", err)
		}
	}
	return err
}
func (p *AutomaticUpdate) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("AutomaticUpdate::Put")
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

func (p *AutomaticUpdate) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("AutomaticUpdate::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
	return err
}
