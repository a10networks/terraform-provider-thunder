

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Import struct {
	Inst struct {

    Aflex string `json:"aflex"`
    AuthJwks string `json:"auth-jwks"`
    AuthPortal string `json:"auth-portal"`
    AuthPortalImage string `json:"auth-portal-image"`
    AuthSamlIdp ImportAuthSamlIdp439 `json:"auth-saml-idp"`
    Background int `json:"background"`
    BiosFile string `json:"bios-file"`
    BwList string `json:"bw-list"`
    CaCert string `json:"ca-cert"`
    CertificateType string `json:"certificate-type"`
    ClassList string `json:"class-list"`
    ClassListConvert string `json:"class-list-convert"`
    ClassListType string `json:"class-list-type"`
    CloudConfig string `json:"cloud-config"`
    CloudCreds string `json:"cloud-creds"`
    CsrGenerate int `json:"csr-generate"`
    DdosScript string `json:"ddos-script"`
    Digest string `json:"digest"`
    DnssecDnskey string `json:"dnssec-dnskey"`
    DnssecDs string `json:"dnssec-ds"`
    DomainList string `json:"domain-list"`
    GeoLocation string `json:"geo-location"`
    GeoLocationArchive ImportGeoLocationArchive440 `json:"geo-location-archive"`
    GlmCert string `json:"glm-cert"`
    GlmLicense string `json:"glm-license"`
    HealthExternal ImportHealthExternal441 `json:"health-external"`
    HealthPostfile ImportHealthPostfile442 `json:"health-postfile"`
    IpMapList string `json:"ip-map-list"`
    LocalUriFile string `json:"local-uri-file"`
    Lw4o6 string `json:"lw-4o6"`
    NgWafCustomPage ImportNgWafCustomPage443 `json:"ng-waf-custom-page"`
    NgWafModule ImportNgWafModule444 `json:"ng-waf-module"`
    Overwrite int `json:"overwrite"`
    Password string `json:"password"`
    PfxPassword string `json:"pfx-password"`
    RemoteFile string `json:"remote-file"`
    RemoteFileZoneTransfer string `json:"remote-file-zone-transfer"`
    Rpz string `json:"rpz"`
    Secured int `json:"secured"`
    SslCert string `json:"ssl-cert"`
    SslCertKey string `json:"ssl-cert-key"`
    SslCrl string `json:"ssl-crl"`
    SslKey string `json:"ssl-key"`
    Store ImportStore445 `json:"store"`
    StoreName string `json:"store-name"`
    Terminal int `json:"terminal"`
    ThalesKmdata string `json:"thales-kmdata"`
    ThalesSecworld string `json:"thales-secworld"`
    ToDevice ImportToDevice446 `json:"to-device"`
    Tsig string `json:"tsig"`
    UsbLicense string `json:"usb-license"`
    UseMgmtPort int `json:"use-mgmt-port"`
    UserTag string `json:"user-tag"`
    WebCategoryLicense string `json:"web-category-license"`
    XmlSchema string `json:"xml-schema"`
    ZoneTransfer string `json:"zone-transfer"`

	} `json:"import"`
}


type ImportAuthSamlIdp439 struct {
    SamlIdpName string `json:"saml-idp-name"`
    VerifyXmlSignature int `json:"verify-xml-signature"`
    Overwrite int `json:"overwrite"`
    UseMgmtPort int `json:"use-mgmt-port"`
    RemoteFile string `json:"remote-file"`
    Password string `json:"password"`
}


type ImportGeoLocationArchive440 struct {
    GeoLocationArchiveFormat string `json:"geo-location-archive-format"`
    UseMgmtPort int `json:"use-mgmt-port"`
    RemoteFile string `json:"remote-file"`
    Password string `json:"password"`
}


type ImportHealthExternal441 struct {
    Externalfilename string `json:"externalfilename"`
    Description string `json:"description"`
    Overwrite int `json:"overwrite"`
    UseMgmtPort int `json:"use-mgmt-port"`
    RemoteFile string `json:"remote-file"`
    Password string `json:"password"`
}


type ImportHealthPostfile442 struct {
    Postfilename string `json:"postfilename"`
    Overwrite int `json:"overwrite"`
    UseMgmtPort int `json:"use-mgmt-port"`
    RemoteFile string `json:"remote-file"`
    Password string `json:"password"`
}


type ImportNgWafCustomPage443 struct {
    CustomPage string `json:"custom-page"`
    Overwrite int `json:"overwrite"`
    UseMgmtPort int `json:"use-mgmt-port"`
    RemoteFile string `json:"remote-file"`
    Password string `json:"password"`
}


type ImportNgWafModule444 struct {
    Overwrite int `json:"overwrite"`
    UseMgmtPort int `json:"use-mgmt-port"`
    RemoteFile string `json:"remote-file"`
    Password string `json:"password"`
}


type ImportStore445 struct {
    Delete int `json:"delete"`
    Create int `json:"create"`
    Name string `json:"name"`
    RemoteFile string `json:"remote-file"`
}


type ImportToDevice446 struct {
    Device int `json:"device"`
    GlmLicense string `json:"glm-license"`
    GlmCert string `json:"glm-cert"`
    WebCategoryLicense string `json:"web-category-license"`
    Overwrite int `json:"overwrite"`
    UseMgmtPort int `json:"use-mgmt-port"`
    RemoteFile string `json:"remote-file"`
}

func (p *Import) GetId() string{
    return "1"
}

func (p *Import) getPath() string{
    return "import"
}

func (p *Import) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Import::Post")
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

func (p *Import) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Import::Get")
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
func (p *Import) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Import::Put")
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

func (p *Import) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Import::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
