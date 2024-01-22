

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Export struct {
	Inst struct {

    Aflex string `json:"aflex"`
    AuthJwks string `json:"auth-jwks"`
    AuthPortal string `json:"auth-portal"`
    AuthPortalImage string `json:"auth-portal-image"`
    Axdebug string `json:"axdebug"`
    BwList string `json:"bw-list"`
    CaCert string `json:"ca-cert"`
    CaptureConfig string `json:"capture-config"`
    CaptureConfigRealtime string `json:"capture-config-realtime"`
    ClassList string `json:"class-list"`
    Csr string `json:"csr"`
    DebugMonitor string `json:"debug-monitor"`
    DebugTrafficCapture string `json:"debug-traffic-capture"`
    DebugTrafficCaptureChassis string `json:"debug-traffic-capture-chassis"`
    DebugTrafficCaptureChassisSlot int `json:"debug-traffic-capture-chassis-slot"`
    DnssecDnskey string `json:"dnssec-dnskey"`
    DnssecDs string `json:"dnssec-ds"`
    DomainList string `json:"domain-list"`
    Externalfilename string `json:"externalfilename"`
    FixedNat string `json:"fixed-nat"`
    FixedNatArchive string `json:"fixed-nat-archive"`
    GeoLocation string `json:"geo-location"`
    GeoLocationArchive ExportGeoLocationArchive347 `json:"geo-location-archive"`
    IpMapList string `json:"ip-map-list"`
    IpsecErrorDump string `json:"ipsec-error-dump"`
    LocalUriFile string `json:"local-uri-file"`
    Lw4o6 string `json:"lw-4o6"`
    Lw4o6BindingTableValidationLog string `json:"lw-4o6-binding-table-validation-log"`
    MergedPcap int `json:"merged-pcap"`
    MonEntityDebugFile string `json:"mon-entity-debug-file"`
    Password string `json:"password"`
    PerCpu int `json:"per-cpu"`
    PktCount int `json:"pkt-count"`
    PktcaptureFile string `json:"pktcapture-file"`
    Profile string `json:"profile"`
    RemoteFile string `json:"remote-file"`
    Rpz string `json:"rpz"`
    RunningConfig int `json:"running-config"`
    SamlIdpName string `json:"saml-idp-name"`
    SslCert string `json:"ssl-cert"`
    SslCertKey string `json:"ssl-cert-key"`
    SslCrl string `json:"ssl-crl"`
    SslKey string `json:"ssl-key"`
    StartupConfig int `json:"startup-config"`
    StatusCheck int `json:"status-check"`
    Store ExportStore348 `json:"store"`
    StoreName string `json:"store-name"`
    Syslog string `json:"syslog"`
    Tgz int `json:"tgz"`
    ThalesKmdata string `json:"thales-kmdata"`
    ThalesSecworld string `json:"thales-secworld"`
    Tsig string `json:"tsig"`
    UseMgmtPort int `json:"use-mgmt-port"`
    Visibility int `json:"visibility"`
    XmlSchema string `json:"xml-schema"`

	} `json:"export"`
}


type ExportGeoLocationArchive347 struct {
    GeoLocationArchiveName string `json:"geo-location-archive-name"`
    UseMgmtPort int `json:"use-mgmt-port"`
    RemoteFile string `json:"remote-file"`
    Password string `json:"password"`
}


type ExportStore348 struct {
    Delete int `json:"delete"`
    Create int `json:"create"`
    Name string `json:"name"`
    RemoteFile string `json:"remote-file"`
}

func (p *Export) GetId() string{
    return "1"
}

func (p *Export) getPath() string{
    return "export"
}

func (p *Export) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Export::Post")
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

func (p *Export) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Export::Get")
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
func (p *Export) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Export::Put")
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

func (p *Export) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Export::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
