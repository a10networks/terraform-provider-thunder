

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type Admin struct {
	Inst struct {

    Access AdminAccess60 `json:"access"`
    AccessList int `json:"access-list"`
    Action string `json:"action" dval:"enable"`
    AwsAccesskey AdminAwsAccesskey61 `json:"aws-accesskey"`
    AzureCred AdminAzureCred62 `json:"azure-cred"`
    CloudCred AdminCloudCred63 `json:"cloud-cred"`
    Encrypted string `json:"encrypted"`
    GcpCred AdminGcpCred64 `json:"gcp-cred"`
    PasswdString string `json:"passwd-string"`
    Password AdminPassword65 `json:"password"`
    PasswordKey int `json:"password-key"`
    PrivilegeGlobal string `json:"privilege-global"`
    PrivilegeList []AdminPrivilegeList `json:"privilege-list"`
    PrivilegeShellRoot int `json:"privilege-shell-root"`
    SshPubkey AdminSshPubkey66 `json:"ssh-pubkey"`
    TrustedHost int `json:"trusted-host"`
    TrustedHostAclId int `json:"trusted-host-acl-id"`
    TrustedHostCidr string `json:"trusted-host-cidr"`
    User string `json:"user"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"admin"`
}


type AdminAccess60 struct {
    AccessType string `json:"access-type" dval:"axapi,cli,web"`
    Uuid string `json:"uuid"`
}


type AdminAwsAccesskey61 struct {
    Import int `json:"import"`
    UseMgmtPort int `json:"use-mgmt-port"`
    FileUrl string `json:"file-url"`
    Delete int `json:"delete"`
    Show int `json:"show"`
}


type AdminAzureCred62 struct {
    Import int `json:"import"`
    UseMgmtPort int `json:"use-mgmt-port"`
    FileUrl string `json:"file-url"`
    Delete int `json:"delete"`
    Show int `json:"show"`
}


type AdminCloudCred63 struct {
    Type string `json:"type"`
    Import int `json:"import"`
    UseMgmtPort int `json:"use-mgmt-port"`
    FileUrl string `json:"file-url"`
    Delete int `json:"delete"`
    Show int `json:"show"`
}


type AdminGcpCred64 struct {
    Import int `json:"import"`
    UseMgmtPort int `json:"use-mgmt-port"`
    FileUrl string `json:"file-url"`
    Delete int `json:"delete"`
    Show int `json:"show"`
}


type AdminPassword65 struct {
    PasswordInModule string `json:"password-in-module"`
    EncryptedInModule string `json:"encrypted-in-module"`
    Uuid string `json:"uuid"`
}


type AdminPrivilegeList struct {
    PrivilegePartition string `json:"privilege-partition"`
    PartitionName string `json:"partition-name"`
}


type AdminSshPubkey66 struct {
    Import int `json:"import"`
    UseMgmtPort int `json:"use-mgmt-port"`
    FileUrl string `json:"file-url"`
    Delete int `json:"delete"`
    List int `json:"list"`
}

func (p *Admin) GetId() string{
    return url.QueryEscape(p.Inst.User)
}

func (p *Admin) getPath() string{
    return "admin"
}

func (p *Admin) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Admin::Post")
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

func (p *Admin) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Admin::Get")
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
func (p *Admin) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Admin::Put")
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

func (p *Admin) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Admin::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
