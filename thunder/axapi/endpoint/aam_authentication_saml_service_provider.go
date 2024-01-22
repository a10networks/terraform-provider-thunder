

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type AamAuthenticationSamlServiceProvider struct {
	Inst struct {

    AcsUriBypass int `json:"acs-uri-bypass"`
    AdfsWsFederation AamAuthenticationSamlServiceProviderAdfsWsFederation `json:"adfs-ws-federation"`
    ArtifactResolutionService []AamAuthenticationSamlServiceProviderArtifactResolutionService `json:"artifact-resolution-service"`
    AssertionConsumingService []AamAuthenticationSamlServiceProviderAssertionConsumingService `json:"assertion-consuming-service"`
    BadRequestRedirectUrl string `json:"bad-request-redirect-url"`
    Certificate string `json:"certificate"`
    EntityId string `json:"entity-id"`
    MetadataExportService AamAuthenticationSamlServiceProviderMetadataExportService `json:"metadata-export-service"`
    Name string `json:"name"`
    PacketCaptureTemplate string `json:"packet-capture-template"`
    RequireAssertionSigned AamAuthenticationSamlServiceProviderRequireAssertionSigned `json:"require-assertion-signed"`
    SamlRequestSigned AamAuthenticationSamlServiceProviderSamlRequestSigned `json:"saml-request-signed"`
    SamplingEnable []AamAuthenticationSamlServiceProviderSamplingEnable `json:"sampling-enable"`
    ServiceUrl string `json:"service-url"`
    SignatureAlgorithm string `json:"signature-algorithm" dval:"SHA1"`
    SingleLogoutService []AamAuthenticationSamlServiceProviderSingleLogoutService `json:"single-logout-service"`
    SoapTlsCertificateValidate AamAuthenticationSamlServiceProviderSoapTlsCertificateValidate `json:"soap-tls-certificate-validate"`
    SpInitiatedSingleLogoutService []AamAuthenticationSamlServiceProviderSpInitiatedSingleLogoutService `json:"SP-initiated-single-logout-service"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"service-provider"`
}


type AamAuthenticationSamlServiceProviderAdfsWsFederation struct {
    WsFederationEnable int `json:"ws-federation-enable"`
}


type AamAuthenticationSamlServiceProviderArtifactResolutionService struct {
    ArtifactIndex int `json:"artifact-index"`
    ArtifactLocation string `json:"artifact-location"`
    ArtifactBinding string `json:"artifact-binding"`
}


type AamAuthenticationSamlServiceProviderAssertionConsumingService struct {
    AssertionIndex int `json:"assertion-index"`
    AssertionLocation string `json:"assertion-location"`
    AssertionBinding string `json:"assertion-binding"`
}


type AamAuthenticationSamlServiceProviderMetadataExportService struct {
    MdExportLocation string `json:"md-export-location"`
    SignXml int `json:"sign-xml"`
}


type AamAuthenticationSamlServiceProviderRequireAssertionSigned struct {
    RequireAssertionSignedEnable int `json:"require-assertion-signed-enable"`
}


type AamAuthenticationSamlServiceProviderSamlRequestSigned struct {
    SamlRequestSignedDisable int `json:"saml-request-signed-disable"`
}


type AamAuthenticationSamlServiceProviderSamplingEnable struct {
    Counters1 string `json:"counters1"`
}


type AamAuthenticationSamlServiceProviderSingleLogoutService struct {
    SloLocation string `json:"SLO-location"`
    SloBinding string `json:"SLO-binding"`
}


type AamAuthenticationSamlServiceProviderSoapTlsCertificateValidate struct {
    SoapTlsCertificateValidateDisable int `json:"soap-tls-certificate-validate-disable"`
}


type AamAuthenticationSamlServiceProviderSpInitiatedSingleLogoutService struct {
    SpSloLocation string `json:"SP-SLO-location"`
    Asynchronous int `json:"asynchronous"`
}

func (p *AamAuthenticationSamlServiceProvider) GetId() string{
    return p.Inst.Name
}

func (p *AamAuthenticationSamlServiceProvider) getPath() string{
    return "aam/authentication/saml/service-provider"
}

func (p *AamAuthenticationSamlServiceProvider) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationSamlServiceProvider::Post")
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

func (p *AamAuthenticationSamlServiceProvider) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationSamlServiceProvider::Get")
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
func (p *AamAuthenticationSamlServiceProvider) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationSamlServiceProvider::Put")
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

func (p *AamAuthenticationSamlServiceProvider) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationSamlServiceProvider::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
