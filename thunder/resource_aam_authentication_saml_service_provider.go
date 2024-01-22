package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAamAuthenticationSamlServiceProvider() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_aam_authentication_saml_service_provider`: Authentication service provider\n\n__PLACEHOLDER__",
		CreateContext: resourceAamAuthenticationSamlServiceProviderCreate,
		UpdateContext: resourceAamAuthenticationSamlServiceProviderUpdate,
		ReadContext:   resourceAamAuthenticationSamlServiceProviderRead,
		DeleteContext: resourceAamAuthenticationSamlServiceProviderDelete,

		Schema: map[string]*schema.Schema{
			"acs_uri_bypass": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "After user authenticated, bypass requests with assertion-consuming-service location URI",
			},
			"adfs_ws_federation": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ws_federation_enable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable ADFS WS-Federation",
						},
					},
				},
			},
			"artifact_resolution_service": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"artifact_index": {
							Type: schema.TypeInt, Optional: true, Description: "The index of artifact resolution service",
						},
						"artifact_location": {
							Type: schema.TypeString, Optional: true, Description: "The location of artifact resolution service. (ex. /SAML/POST)",
						},
						"artifact_binding": {
							Type: schema.TypeString, Optional: true, Description: "'soap': SOAP binding of artifact resolution service;",
						},
					},
				},
			},
			"assertion_consuming_service": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"assertion_index": {
							Type: schema.TypeInt, Optional: true, Description: "The index of assertion consuming service",
						},
						"assertion_location": {
							Type: schema.TypeString, Optional: true, Description: "The location of assertion consuming service endpoint. (ex. /SAML/POST)",
						},
						"assertion_binding": {
							Type: schema.TypeString, Optional: true, Description: "'artifact': Artifact binding of assertion consuming service; 'paos': PAOS binding of assertion consuming service; 'post': POST binding of assertion consuming service;",
						},
					},
				},
			},
			"bad_request_redirect_url": {
				Type: schema.TypeString, Optional: true, Description: "Specify URL to redirect",
			},
			"certificate": {
				Type: schema.TypeString, Optional: true, Description: "SAML service provider certificate file (PFX format is required.)",
			},
			"entity_id": {
				Type: schema.TypeString, Optional: true, Description: "SAML service provider entity ID",
			},
			"metadata_export_service": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"md_export_location": {
							Type: schema.TypeString, Optional: true, Description: "Specify the URI to export SP metadata (Export URI. Default is /A10SP_Metadata)",
						},
						"sign_xml": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Sign exported SP metadata XML with SP's certificate",
						},
					},
				},
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Specify SAML authentication service provider name",
			},
			"packet_capture_template": {
				Type: schema.TypeString, Optional: true, Description: "Name of the packet capture template to be bind with this object",
			},
			"require_assertion_signed": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"require_assertion_signed_enable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable required signing of SAML assertion",
						},
					},
				},
			},
			"saml_request_signed": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"saml_request_signed_disable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable signing signature for SAML (Authn/Artifact Resolve) requests",
						},
					},
				},
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'sp-metadata-export-req': Metadata Export Request; 'sp-metadata-export-success': Metadata Export Success; 'login-auth-req': Login Authentication Request; 'login-auth-resp': Login Authentication Response; 'acs-req': SAML Single-Sign-On Request; 'acs-success': SAML Single-Sign-On Success; 'acs-authz-fail': SAML Single-Sign-On Authorization Fail; 'acs-error': SAML Single-Sign-On Error; 'slo-req': Single Logout Request; 'slo-success': Single Logout Success; 'slo-error': Single Logout Error; 'sp-slo-req': SP-initiated Single Logout Request; 'glo-slo-success': Total Global Logout Success; 'loc-slo-success': Total Local Logout Success; 'par-slo-success': Total Partial Logout Success; 'other-error': Other Error;",
						},
					},
				},
			},
			"service_url": {
				Type: schema.TypeString, Optional: true, Description: "SAML service provider service URL (ex. https://www.a10networks.com/saml.sso)",
			},
			"signature_algorithm": {
				Type: schema.TypeString, Optional: true, Default: "SHA1", Description: "'SHA1': use SHA1 as signature algorithm (default); 'SHA256': use SHA256 as signature algorithm;",
			},
			"single_logout_service": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"slo_location": {
							Type: schema.TypeString, Optional: true, Description: "The location of name-id management service. (ex. /SAML/POST)",
						},
						"slo_binding": {
							Type: schema.TypeString, Optional: true, Description: "'post': POST binding of single logout service; 'redirect': Redirect binding of single logout service; 'soap': SOAP binding of single logout service;",
						},
					},
				},
			},
			"soap_tls_certificate_validate": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"soap_tls_certificate_validate_disable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable verification for server certificate in TLS session when resolving artificate",
						},
					},
				},
			},
			"sp_initiated_single_logout_service": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sp_slo_location": {
							Type: schema.TypeString, Optional: true, Description: "The location of SP-initiated single logout service endpoint. (ex. /Logout)",
						},
						"asynchronous": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "the IDP will not send a logout response to AX",
						},
					},
				},
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceAamAuthenticationSamlServiceProviderCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationSamlServiceProviderCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationSamlServiceProvider(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAamAuthenticationSamlServiceProviderRead(ctx, d, meta)
	}
	return diags
}

func resourceAamAuthenticationSamlServiceProviderUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationSamlServiceProviderUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationSamlServiceProvider(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAamAuthenticationSamlServiceProviderRead(ctx, d, meta)
	}
	return diags
}
func resourceAamAuthenticationSamlServiceProviderDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationSamlServiceProviderDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationSamlServiceProvider(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceAamAuthenticationSamlServiceProviderRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationSamlServiceProviderRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationSamlServiceProvider(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectAamAuthenticationSamlServiceProviderAdfsWsFederation(d []interface{}) edpt.AamAuthenticationSamlServiceProviderAdfsWsFederation {

	count1 := len(d)
	var ret edpt.AamAuthenticationSamlServiceProviderAdfsWsFederation
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.WsFederationEnable = in["ws_federation_enable"].(int)
	}
	return ret
}

func getSliceAamAuthenticationSamlServiceProviderArtifactResolutionService(d []interface{}) []edpt.AamAuthenticationSamlServiceProviderArtifactResolutionService {

	count1 := len(d)
	ret := make([]edpt.AamAuthenticationSamlServiceProviderArtifactResolutionService, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AamAuthenticationSamlServiceProviderArtifactResolutionService
		oi.ArtifactIndex = in["artifact_index"].(int)
		oi.ArtifactLocation = in["artifact_location"].(string)
		oi.ArtifactBinding = in["artifact_binding"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceAamAuthenticationSamlServiceProviderAssertionConsumingService(d []interface{}) []edpt.AamAuthenticationSamlServiceProviderAssertionConsumingService {

	count1 := len(d)
	ret := make([]edpt.AamAuthenticationSamlServiceProviderAssertionConsumingService, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AamAuthenticationSamlServiceProviderAssertionConsumingService
		oi.AssertionIndex = in["assertion_index"].(int)
		oi.AssertionLocation = in["assertion_location"].(string)
		oi.AssertionBinding = in["assertion_binding"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectAamAuthenticationSamlServiceProviderMetadataExportService(d []interface{}) edpt.AamAuthenticationSamlServiceProviderMetadataExportService {

	count1 := len(d)
	var ret edpt.AamAuthenticationSamlServiceProviderMetadataExportService
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.MdExportLocation = in["md_export_location"].(string)
		ret.SignXml = in["sign_xml"].(int)
	}
	return ret
}

func getObjectAamAuthenticationSamlServiceProviderRequireAssertionSigned(d []interface{}) edpt.AamAuthenticationSamlServiceProviderRequireAssertionSigned {

	count1 := len(d)
	var ret edpt.AamAuthenticationSamlServiceProviderRequireAssertionSigned
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RequireAssertionSignedEnable = in["require_assertion_signed_enable"].(int)
	}
	return ret
}

func getObjectAamAuthenticationSamlServiceProviderSamlRequestSigned(d []interface{}) edpt.AamAuthenticationSamlServiceProviderSamlRequestSigned {

	count1 := len(d)
	var ret edpt.AamAuthenticationSamlServiceProviderSamlRequestSigned
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SamlRequestSignedDisable = in["saml_request_signed_disable"].(int)
	}
	return ret
}

func getSliceAamAuthenticationSamlServiceProviderSamplingEnable(d []interface{}) []edpt.AamAuthenticationSamlServiceProviderSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.AamAuthenticationSamlServiceProviderSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AamAuthenticationSamlServiceProviderSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceAamAuthenticationSamlServiceProviderSingleLogoutService(d []interface{}) []edpt.AamAuthenticationSamlServiceProviderSingleLogoutService {

	count1 := len(d)
	ret := make([]edpt.AamAuthenticationSamlServiceProviderSingleLogoutService, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AamAuthenticationSamlServiceProviderSingleLogoutService
		oi.SloLocation = in["slo_location"].(string)
		oi.SloBinding = in["slo_binding"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectAamAuthenticationSamlServiceProviderSoapTlsCertificateValidate(d []interface{}) edpt.AamAuthenticationSamlServiceProviderSoapTlsCertificateValidate {

	count1 := len(d)
	var ret edpt.AamAuthenticationSamlServiceProviderSoapTlsCertificateValidate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SoapTlsCertificateValidateDisable = in["soap_tls_certificate_validate_disable"].(int)
	}
	return ret
}

func getSliceAamAuthenticationSamlServiceProviderSpInitiatedSingleLogoutService(d []interface{}) []edpt.AamAuthenticationSamlServiceProviderSpInitiatedSingleLogoutService {

	count1 := len(d)
	ret := make([]edpt.AamAuthenticationSamlServiceProviderSpInitiatedSingleLogoutService, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AamAuthenticationSamlServiceProviderSpInitiatedSingleLogoutService
		oi.SpSloLocation = in["sp_slo_location"].(string)
		oi.Asynchronous = in["asynchronous"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointAamAuthenticationSamlServiceProvider(d *schema.ResourceData) edpt.AamAuthenticationSamlServiceProvider {
	var ret edpt.AamAuthenticationSamlServiceProvider
	ret.Inst.AcsUriBypass = d.Get("acs_uri_bypass").(int)
	ret.Inst.AdfsWsFederation = getObjectAamAuthenticationSamlServiceProviderAdfsWsFederation(d.Get("adfs_ws_federation").([]interface{}))
	ret.Inst.ArtifactResolutionService = getSliceAamAuthenticationSamlServiceProviderArtifactResolutionService(d.Get("artifact_resolution_service").([]interface{}))
	ret.Inst.AssertionConsumingService = getSliceAamAuthenticationSamlServiceProviderAssertionConsumingService(d.Get("assertion_consuming_service").([]interface{}))
	ret.Inst.BadRequestRedirectUrl = d.Get("bad_request_redirect_url").(string)
	ret.Inst.Certificate = d.Get("certificate").(string)
	ret.Inst.EntityId = d.Get("entity_id").(string)
	ret.Inst.MetadataExportService = getObjectAamAuthenticationSamlServiceProviderMetadataExportService(d.Get("metadata_export_service").([]interface{}))
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.PacketCaptureTemplate = d.Get("packet_capture_template").(string)
	ret.Inst.RequireAssertionSigned = getObjectAamAuthenticationSamlServiceProviderRequireAssertionSigned(d.Get("require_assertion_signed").([]interface{}))
	ret.Inst.SamlRequestSigned = getObjectAamAuthenticationSamlServiceProviderSamlRequestSigned(d.Get("saml_request_signed").([]interface{}))
	ret.Inst.SamplingEnable = getSliceAamAuthenticationSamlServiceProviderSamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.ServiceUrl = d.Get("service_url").(string)
	ret.Inst.SignatureAlgorithm = d.Get("signature_algorithm").(string)
	ret.Inst.SingleLogoutService = getSliceAamAuthenticationSamlServiceProviderSingleLogoutService(d.Get("single_logout_service").([]interface{}))
	ret.Inst.SoapTlsCertificateValidate = getObjectAamAuthenticationSamlServiceProviderSoapTlsCertificateValidate(d.Get("soap_tls_certificate_validate").([]interface{}))
	ret.Inst.SpInitiatedSingleLogoutService = getSliceAamAuthenticationSamlServiceProviderSpInitiatedSingleLogoutService(d.Get("sp_initiated_single_logout_service").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
