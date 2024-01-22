package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAamAuthenticationServerOcspInstance() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_aam_authentication_server_ocsp_instance`: Specify OCSP authentication server name\n\n__PLACEHOLDER__",
		CreateContext: resourceAamAuthenticationServerOcspInstanceCreate,
		UpdateContext: resourceAamAuthenticationServerOcspInstanceUpdate,
		ReadContext:   resourceAamAuthenticationServerOcspInstanceRead,
		DeleteContext: resourceAamAuthenticationServerOcspInstanceDelete,

		Schema: map[string]*schema.Schema{
			"health_check": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Check server's health status",
			},
			"health_check_disable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable configured health check configuration",
			},
			"health_check_string": {
				Type: schema.TypeString, Optional: true, Description: "Health monitor name",
			},
			"http_version": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Set HTTP version (default 1.0)",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Specify OCSP authentication server name",
			},
			"packet_capture_template": {
				Type: schema.TypeString, Optional: true, Description: "Name of the packet capture template to be bind with this object",
			},
			"port_health_check": {
				Type: schema.TypeString, Optional: true, Description: "Check port's health status",
			},
			"port_health_check_disable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable configured port health check configuration",
			},
			"responder_ca": {
				Type: schema.TypeString, Optional: true, Description: "Specify the trusted OCSP responder's CA cert filename",
			},
			"responder_cert": {
				Type: schema.TypeString, Optional: true, Description: "Specify the trusted OCSP responder's cert filename",
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
			"url": {
				Type: schema.TypeString, Optional: true, Description: "Specify the OCSP server's address (Format: http://host[:port]/) (The OCSP server's address(Format: http://host[:port]/))",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"version_type": {
				Type: schema.TypeString, Optional: true, Description: "'1.1': HTTP version 1.1;",
			},
		},
	}
}
func resourceAamAuthenticationServerOcspInstanceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationServerOcspInstanceCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationServerOcspInstance(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAamAuthenticationServerOcspInstanceRead(ctx, d, meta)
	}
	return diags
}

func resourceAamAuthenticationServerOcspInstanceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationServerOcspInstanceUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationServerOcspInstance(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAamAuthenticationServerOcspInstanceRead(ctx, d, meta)
	}
	return diags
}
func resourceAamAuthenticationServerOcspInstanceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationServerOcspInstanceDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationServerOcspInstance(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceAamAuthenticationServerOcspInstanceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationServerOcspInstanceRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationServerOcspInstance(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceAamAuthenticationServerOcspInstanceSamplingEnable(d []interface{}) []edpt.AamAuthenticationServerOcspInstanceSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.AamAuthenticationServerOcspInstanceSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AamAuthenticationServerOcspInstanceSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointAamAuthenticationServerOcspInstance(d *schema.ResourceData) edpt.AamAuthenticationServerOcspInstance {
	var ret edpt.AamAuthenticationServerOcspInstance
	ret.Inst.HealthCheck = d.Get("health_check").(int)
	ret.Inst.HealthCheckDisable = d.Get("health_check_disable").(int)
	ret.Inst.HealthCheckString = d.Get("health_check_string").(string)
	ret.Inst.HttpVersion = d.Get("http_version").(int)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.PacketCaptureTemplate = d.Get("packet_capture_template").(string)
	ret.Inst.PortHealthCheck = d.Get("port_health_check").(string)
	ret.Inst.PortHealthCheckDisable = d.Get("port_health_check_disable").(int)
	ret.Inst.ResponderCa = d.Get("responder_ca").(string)
	ret.Inst.ResponderCert = d.Get("responder_cert").(string)
	ret.Inst.SamplingEnable = getSliceAamAuthenticationServerOcspInstanceSamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.Url = d.Get("url").(string)
	//omit uuid
	ret.Inst.VersionType = d.Get("version_type").(string)
	return ret
}
