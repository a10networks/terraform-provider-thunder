package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAamAuthenticationServerOcsp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_aam_authentication_server_ocsp`: OCSP Authentication Server\n\n__PLACEHOLDER__",
		CreateContext: resourceAamAuthenticationServerOcspCreate,
		UpdateContext: resourceAamAuthenticationServerOcspUpdate,
		ReadContext:   resourceAamAuthenticationServerOcspRead,
		DeleteContext: resourceAamAuthenticationServerOcspDelete,

		Schema: map[string]*schema.Schema{
			"instance_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type: schema.TypeString, Required: true, Description: "Specify OCSP authentication server name",
						},
						"url": {
							Type: schema.TypeString, Optional: true, Description: "Specify the OCSP server's address (Format: http://host[:port]/) (The OCSP server's address(Format: http://host[:port]/))",
						},
						"responder_ca": {
							Type: schema.TypeString, Optional: true, Description: "Specify the trusted OCSP responder's CA cert filename",
						},
						"responder_cert": {
							Type: schema.TypeString, Optional: true, Description: "Specify the trusted OCSP responder's cert filename",
						},
						"health_check": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Check server's health status",
						},
						"health_check_string": {
							Type: schema.TypeString, Optional: true, Description: "Health monitor name",
						},
						"health_check_disable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable configured health check configuration",
						},
						"port_health_check": {
							Type: schema.TypeString, Optional: true, Description: "Check port's health status",
						},
						"port_health_check_disable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable configured port health check configuration",
						},
						"http_version": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Set HTTP version (default 1.0)",
						},
						"version_type": {
							Type: schema.TypeString, Optional: true, Description: "'1.1': HTTP version 1.1;",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"sampling_enable": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"counters1": {
										Type: schema.TypeString, Optional: true, Description: "'all': all; 'request': Request; 'certificate-good': Good Certificate Response; 'certificate-revoked': Revoked Certificate Response; 'certificate-unknown': Unknown Certificate Response; 'timeout': Timeout; 'fail': Handle OCSP response failed; 'stapling-request': OCSP Stapling Request Send; 'stapling-certificate-good': OCSP Stapling Good Certificate Response; 'stapling-certificate-revoked': OCSP Stapling Revoked Certificate Response; 'stapling-certificate-unknown': OCSP Stapling Unknown Certificate Response; 'stapling-timeout': OCSP Stapling Timeout; 'stapling-fail': Handle OCSP response failed;",
									},
								},
							},
						},
						"packet_capture_template": {
							Type: schema.TypeString, Optional: true, Description: "Name of the packet capture template to be bind with this object",
						},
					},
				},
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'stapling-certificate-good': Total OCSP Stapling Good Certificate Response; 'stapling-certificate-revoked': Total OCSP Stapling Revoked Certificate Response; 'stapling-certificate-unknown': Total OCSP Stapling Unknown Certificate Response; 'stapling-request-normal': Total OSCP Stapling Normal Request; 'stapling-request-dropped': Total OCSP Stapling Dropped Request; 'stapling-response-success': Total OCSP Stapling Success Response; 'stapling-response-failure': Total OCSP Stapling Failure Response; 'stapling-response-error': Total OCSP Stapling Error Response; 'stapling-response-timeout': Total OCSP Stapling Timeout Response; 'stapling-response-other': Total OCSP Stapling Other Response; 'request-normal': Total OSCP Normal Request; 'request-dropped': Total OCSP Dropped Request; 'response-success': Total OCSP Success Response; 'response-failure': Total OCSP Failure Response; 'response-error': Total OCSP Error Response; 'response-timeout': Total OCSP Timeout Response; 'response-other': Total OCSP Other Response; 'job-start-error': Total OCSP Job Start Error; 'polling-control-error': Total OCSP Polling Control Error;",
						},
					},
				},
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceAamAuthenticationServerOcspCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationServerOcspCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationServerOcsp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAamAuthenticationServerOcspRead(ctx, d, meta)
	}
	return diags
}

func resourceAamAuthenticationServerOcspUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationServerOcspUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationServerOcsp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAamAuthenticationServerOcspRead(ctx, d, meta)
	}
	return diags
}
func resourceAamAuthenticationServerOcspDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationServerOcspDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationServerOcsp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceAamAuthenticationServerOcspRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationServerOcspRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationServerOcsp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceAamAuthenticationServerOcspInstanceList(d []interface{}) []edpt.AamAuthenticationServerOcspInstanceList {

	count1 := len(d)
	ret := make([]edpt.AamAuthenticationServerOcspInstanceList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AamAuthenticationServerOcspInstanceList
		oi.Name = in["name"].(string)
		oi.Url = in["url"].(string)
		oi.ResponderCa = in["responder_ca"].(string)
		oi.ResponderCert = in["responder_cert"].(string)
		oi.HealthCheck = in["health_check"].(int)
		oi.HealthCheckString = in["health_check_string"].(string)
		oi.HealthCheckDisable = in["health_check_disable"].(int)
		oi.PortHealthCheck = in["port_health_check"].(string)
		oi.PortHealthCheckDisable = in["port_health_check_disable"].(int)
		oi.HttpVersion = in["http_version"].(int)
		oi.VersionType = in["version_type"].(string)
		//omit uuid
		oi.SamplingEnable = getSliceAamAuthenticationServerOcspInstanceListSamplingEnable(in["sampling_enable"].([]interface{}))
		oi.PacketCaptureTemplate = in["packet_capture_template"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceAamAuthenticationServerOcspInstanceListSamplingEnable(d []interface{}) []edpt.AamAuthenticationServerOcspInstanceListSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.AamAuthenticationServerOcspInstanceListSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AamAuthenticationServerOcspInstanceListSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceAamAuthenticationServerOcspSamplingEnable(d []interface{}) []edpt.AamAuthenticationServerOcspSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.AamAuthenticationServerOcspSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AamAuthenticationServerOcspSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointAamAuthenticationServerOcsp(d *schema.ResourceData) edpt.AamAuthenticationServerOcsp {
	var ret edpt.AamAuthenticationServerOcsp
	ret.Inst.InstanceList = getSliceAamAuthenticationServerOcspInstanceList(d.Get("instance_list").([]interface{}))
	ret.Inst.SamplingEnable = getSliceAamAuthenticationServerOcspSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
