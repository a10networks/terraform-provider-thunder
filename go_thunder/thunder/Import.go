package go_thunder

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"util"
)

type Import struct {
	GeoLocation ImportInstance `json:"import,omitempty"`
}
type HealthExternal struct {
	Description      string `json:"description,omitempty"`
	RemoteFile       string `json:"remote-file,omitempty"`
	Externalfilename string `json:"externalfilename,omitempty"`
	Password         string `json:"password,omitempty"`
	UseMgmtPort      int    `json:"use-mgmt-port,omitempty"`
	Overwrite        int    `json:"overwrite,omitempty"`
}
type HealthPostfile struct {
	Postfilename string `json:"postfilename,omitempty"`
	Password     string `json:"password,omitempty"`
	UseMgmtPort  int    `json:"use-mgmt-port,omitempty"`
	RemoteFile   string `json:"remote-file,omitempty"`
	Overwrite    int    `json:"overwrite,omitempty"`
}
type ToDevice struct {
	WebCategoryLicense string `json:"web-category-license,omitempty"`
	RemoteFile         string `json:"remote-file,omitempty"`
	GlmLicense         string `json:"glm-license,omitempty"`
	GlmCert            string `json:"glm-cert,omitempty"`
	Device             int    `json:"device,omitempty"`
	UseMgmtPort        int    `json:"use-mgmt-port,omitempty"`
	Overwrite          int    `json:"overwrite,omitempty"`
}
type Store struct {
	Create     int    `json:"create,omitempty"`
	Name       string `json:"name,omitempty"`
	RemoteFile string `json:"remote-file,omitempty"`
	Delete     int    `json:"delete,omitempty"`
}
type AuthSamlIdp struct {
	RemoteFile         string `json:"remote-file,omitempty"`
	SamlIdpName        string `json:"saml-idp-name,omitempty"`
	VerifyXMLSignature int    `json:"verify-xml-signature,omitempty"`
	Password           string `json:"password,omitempty"`
	UseMgmtPort        int    `json:"use-mgmt-port,omitempty"`
	Overwrite          int    `json:"overwrite,omitempty"`
}
type ImportInstance struct {
	GeoLocation      string `json:"geo-location,omitempty"`
	SslCertKey       string `json:"ssl-cert-key,omitempty"`
	ClassListConvert string `json:"class-list-convert,omitempty"`
	BwList           string `json:"bw-list,omitempty"`
	UsbLicense       string `json:"usb-license,omitempty"`
	IPMapList        string `json:"ip-map-list,omitempty"`
	//Externalfilename       HealthExternal `json:"health-external,omitempty"`
	AuthPortal           string `json:"auth-portal,omitempty"`
	LocalURIFile         string `json:"local-uri-file,omitempty"`
	Aflex                string `json:"aflex,omitempty"`
	Overwrite            int    `json:"overwrite,omitempty"`
	ClassListType        string `json:"class-list-type,omitempty"`
	PfxPassword          string `json:"pfx-password,omitempty,omitempty"`
	WebCategoryLicense   string `json:"web-category-license,omitempty"`
	ThalesKmdata         string `json:"thales-kmdata,omitempty"`
	Secured              int    `json:"secured,omitempty"`
	SslCrl               string `json:"ssl-crl,omitempty"`
	Terminal             int    `json:"terminal,omitempty"`
	Policy               string `json:"policy,omitempty"`
	FileInspectionBwList string `json:"file-inspection-bw-list,omitempty"`
	ThalesSecworld       string `json:"thales-secworld,omitempty"`
	Lw4O6                string `json:"lw-4o6,omitempty"`
	AuthPortalImage      string `json:"auth-portal-image,omitempty"`
	//Postfilename       HealthPostfile `json:"health-postfile,omitempty"`
	ClassList   string `json:"class-list,omitempty"`
	GlmLicense  string `json:"glm-license,omitempty"`
	DnssecDs    string `json:"dnssec-ds,omitempty"`
	CloudCreds  string `json:"cloud-creds,omitempty"`
	AuthJwks    string `json:"auth-jwks,omitempty"`
	Wsdl        string `json:"wsdl,omitempty"`
	Password    string `json:"password,omitempty"`
	SslKey      string `json:"ssl-key,omitempty"`
	UseMgmtPort int    `json:"use-mgmt-port,omitempty"`
	RemoteFile  string `json:"remote-file,omitempty"`
	CloudConfig string `json:"cloud-config,omitempty"`
	//Device             ToDevice       `json:"to-device,omitempty"`
	UserTag   string `json:"user-tag,omitempty"`
	StoreName string `json:"store-name,omitempty"`
	CaCert    string `json:"ca-cert,omitempty"`
	GlmCert   string `json:"glm-cert,omitempty"`
	//Store                Store          `json:"store,omitempty"`
	XMLSchema       string `json:"xml-schema,omitempty"`
	CertificateType string `json:"certificate-type,omitempty"`
	//SamlIdpName          AuthSamlIdp    `json:"auth-saml-idp,omitempty"`
	SslCert      string `json:"ssl-cert,omitempty"`
	DnssecDnskey string `json:"dnssec-dnskey,omitempty"`
}

func GetImport(id string, host string) (*Import, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/import", nil, headers)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		logger.Println("The HTTP request failed with error \n", err)
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m Import
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			fmt.Printf("Unmarshal error %s\n", err)
			logger.Println("[INFO] Unmarshal error- %s\n", err)
			return nil, err
		} else {
			fmt.Print(m)
			logger.Println("[INFO] GetImport REQ RES..........................", m)
			err := check_api_status("GetImport", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}
}

func PostImport(id string, vc Import, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id

	payloadBytes, err := json.Marshal(vc)

	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))

	if err != nil {
		logger.Println("[INFO] Marshalling failed with error \n", err)
	}
	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/import", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		logger.Println("The HTTP request failed with error \n", err)
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var v Import
		erro := json.Unmarshal(data, &v)
		if erro != nil {
			fmt.Printf("Unmarshal error %s\n", err)
		} else {
			fmt.Println("response Body:", string(data))
			logger.Println("response Body:", string(data))
			err := check_api_status("PostImport", data)
			if err != nil {
				return err
			}
		}
	}
	return err
}
