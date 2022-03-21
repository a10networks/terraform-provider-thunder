package endpoint

//Operational

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

//based on ACOS 5_2_1-P3_70
type Import struct {
	Inst ImportInst `json:"import"`
}

type ImportInst struct {
	Aflex              string               `json:"aflex,omitempty"`
	AuthJwks           string               `json:"auth-jwks,omitempty"`
	AuthPortal         string               `json:"auth-portal,omitempty"`
	AuthPortalImage    string               `json:"auth-portal-image,omitempty"`
	AuthSamlIdp        ImportAuthSamlIdp    `json:"auth-saml-idp,omitempty"`
	BiosFile           string               `json:"bios-file,omitempty"`
	BwList             string               `json:"bw-list,omitempty"`
	CaCert             string               `json:"ca-cert,omitempty"`
	CertificateType    string               `json:"certificate-type,omitempty"`
	ClassList          string               `json:"class-list,omitempty"`
	ClassListConvert   string               `json:"class-list-convert,omitempty"`
	ClassListType      string               `json:"class-list-type,omitempty"`
	CloudConfig        string               `json:"cloud-config,omitempty"`
	CloudCreds         string               `json:"cloud-creds,omitempty"`
	DnssecDnskey       string               `json:"dnssec-dnskey,omitempty"`
	DnssecDs           string               `json:"dnssec-ds,omitempty"`
	GeoLocation        string               `json:"geo-location,omitempty"`
	GlmCert            string               `json:"glm-cert,omitempty"`
	GlmLicense         string               `json:"glm-license,omitempty"`
	HealthExternal     ImportHealthExternal `json:"health-external,omitempty"`
	HealthPostfile     ImportHealthPostfile `json:"health-postfile,omitempty"`
	IpMapList          string               `json:"ip-map-list,omitempty"`
	LocalUriFile       string               `json:"local-uri-file,omitempty"`
	Lw4o6              string               `json:"lw-4o6,omitempty"`
	Overwrite          int                  `json:"overwrite,omitempty"`
	Password           string               `json:"password,omitempty"`
	PfxPassword        string               `json:"pfx-password,omitempty"`
	Policy             string               `json:"policy,omitempty"`
	RemoteFile         string               `json:"remote-file,omitempty"`
	Rpz                string               `json:"rpz,omitempty"`
	Secured            int                  `json:"secured,omitempty"`
	SslCert            string               `json:"ssl-cert,omitempty"`
	SslCertKey         string               `json:"ssl-cert-key,omitempty"`
	SslCrl             string               `json:"ssl-crl,omitempty"`
	SslKey             string               `json:"ssl-key,omitempty"`
	Store              ImportStore          `json:"store,omitempty"`
	StoreName          string               `json:"store-name,omitempty"`
	Terminal           int                  `json:"terminal,omitempty"`
	ThalesKmdata       string               `json:"thales-kmdata,omitempty"`
	ThalesSecworld     string               `json:"thales-secworld,omitempty"`
	ToDevice           ImportToDevice       `json:"to-device,omitempty"`
	UsbLicense         string               `json:"usb-license,omitempty"`
	UseMgmtPort        int                  `json:"use-mgmt-port,omitempty"`
	UserTag            string               `json:"user-tag,omitempty"`
	WebCategoryLicense string               `json:"web-category-license,omitempty"`
	Wsdl               string               `json:"wsdl,omitempty"`
	XmlSchema          string               `json:"xml-schema,omitempty"`
}

type ImportAuthSamlIdp struct {
	SamlIdpName        string `json:"saml-idp-name,omitempty"`
	VerifyXmlSignature int    `json:"verify-xml-signature,omitempty"`
	Overwrite          int    `json:"overwrite,omitempty"`
	UseMgmtPort        int    `json:"use-mgmt-port,omitempty"`
	RemoteFile         string `json:"remote-file,omitempty"`
	Password           string `json:"password,omitempty"`
}

type ImportHealthExternal struct {
	Externalfilename string `json:"externalfilename,omitempty"`
	Description      string `json:"description,omitempty"`
	Overwrite        int    `json:"overwrite,omitempty"`
	UseMgmtPort      int    `json:"use-mgmt-port,omitempty"`
	RemoteFile       string `json:"remote-file,omitempty"`
	Password         string `json:"password,omitempty"`
}

type ImportHealthPostfile struct {
	Postfilename string `json:"postfilename,omitempty"`
	Overwrite    int    `json:"overwrite,omitempty"`
	UseMgmtPort  int    `json:"use-mgmt-port,omitempty"`
	RemoteFile   string `json:"remote-file,omitempty"`
	Password     string `json:"password,omitempty"`
}

type ImportStore struct {
	Delete     int    `json:"delete,omitempty"`
	Create     int    `json:"create,omitempty"`
	Name       string `json:"name,omitempty"`
	RemoteFile string `json:"remote-file,omitempty"`
}

type ImportToDevice struct {
	Device             int    `json:"device,omitempty"`
	GlmLicense         string `json:"glm-license,omitempty"`
	GlmCert            string `json:"glm-cert,omitempty"`
	WebCategoryLicense string `json:"web-category-license,omitempty"`
	Overwrite          int    `json:"overwrite,omitempty"`
	UseMgmtPort        int    `json:"use-mgmt-port,omitempty"`
	RemoteFile         string `json:"remote-file,omitempty"`
}

func (p *Import) GetId() string {
	return "1"
}

func (p *Import) Post(auth_token string, host string, logger *axapi.ThunderLog) error {
	logger.Println("Import::Post")
	headers := axapi.GenRequestHeader(auth_token)
	payloadBytes, err := json.Marshal(p)
	if err != nil {
		logger.Println("json.Marshal() failed with error", err)
		return err
	}
	logger.Println("payload:", string(payloadBytes))
	_, _, err = axapi.SendPost(host, "import", payloadBytes, headers, logger)
	return err
}
