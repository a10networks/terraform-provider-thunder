package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceWebService() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_web_service`: Configure Web Services\n\n__PLACEHOLDER__",
		CreateContext: resourceWebServiceCreate,
		UpdateContext: resourceWebServiceUpdate,
		ReadContext:   resourceWebServiceRead,
		DeleteContext: resourceWebServiceDelete,

		Schema: map[string]*schema.Schema{
			"auto_redirt_disable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Diable",
			},
			"axapi_idle": {
				Type: schema.TypeInt, Optional: true, Default: 10, Description: "Idle timeout of a xml service connection in minutes (Connection idle timeout value in minutes, default 10, 0 means never timeout)",
			},
			"axapi_session_limit": {
				Type: schema.TypeInt, Optional: true, Default: 30, Description: "Set the max allowed aXAPI sessions (Session limit (default 30))",
			},
			"gui_idle": {
				Type: schema.TypeInt, Optional: true, Default: 10, Description: "Idle timeout of a connection in minutes (Connection idle timeout value in minutes, default 10, 0 means never timeout)",
			},
			"gui_session_limit": {
				Type: schema.TypeInt, Optional: true, Default: 30, Description: "Set the max allowed GUI sessions (Session limit (default 30))",
			},
			"keep_alive_timeout": {
				Type: schema.TypeInt, Optional: true, Default: 30, Description: "Set KeepAliveTimeout in seconds (KeepAliveTimeout in seconds (default 30))",
			},
			"login_message": {
				Type: schema.TypeString, Optional: true, Description: "Set GUI login message",
			},
			"max_keep_alive_requests": {
				Type: schema.TypeInt, Optional: true, Default: 100, Description: "Set MaxKeepAliveRequests (MaxKeepAliveRequests (default 100))",
			},
			"mpm_max_conn": {
				Type: schema.TypeInt, Optional: true, Description: "Set Max Connections of MPM",
			},
			"mpm_max_conn_per_child": {
				Type: schema.TypeInt, Optional: true, Description: "Set Max Connections Per Child of MPM",
			},
			"mpm_min_spare_conn": {
				Type: schema.TypeInt, Optional: true, Description: "Set Min Spare Connections of MPM",
			},
			"port": {
				Type: schema.TypeInt, Optional: true, Default: 80, Description: "Set Web Server Port (Port number(default 80))",
			},
			"pre_login_message": {
				Type: schema.TypeString, Optional: true, Description: "Set Pre GUI login message",
			},
			"public_apis": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"api_uri": {
							Type: schema.TypeString, Optional: true, Description: "API URI",
						},
					},
				},
			},
			"secure": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"restart": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Restart WEB service",
						},
						"wipe": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Wipe WEB private-key and certificate",
						},
						"generate": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"domain_name": {
										Type: schema.TypeString, Optional: true, Description: "The domain name",
									},
									"country": {
										Type: schema.TypeString, Optional: true, Description: "The country name",
									},
									"state": {
										Type: schema.TypeString, Optional: true, Description: "The location",
									},
								},
							},
						},
						"regenerate": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"domain_name": {
										Type: schema.TypeString, Optional: true, Description: "The domain name",
									},
									"country": {
										Type: schema.TypeString, Optional: true, Description: "The country name",
									},
									"state": {
										Type: schema.TypeString, Optional: true, Description: "The location",
									},
								},
							},
						},
						"certificate": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"load": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Load WEB certificate",
									},
									"use_mgmt_port": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use management port as source port",
									},
									"file_url": {
										Type: schema.TypeString, Optional: true, Description: "File URL",
									},
								},
							},
						},
						"private_key": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"load": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Load WEB private-key",
									},
									"use_mgmt_port": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use management port as source port",
									},
									"file_url": {
										Type: schema.TypeString, Optional: true, Description: "File URL",
									},
								},
							},
						},
					},
				},
			},
			"secure_port": {
				Type: schema.TypeInt, Optional: true, Default: 443, Description: "Set web secure server port number for listening (Secure Port Number(default 443))",
			},
			"secure_server_disable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable",
			},
			"server_disable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceWebServiceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceWebServiceCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointWebService(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceWebServiceRead(ctx, d, meta)
	}
	return diags
}

func resourceWebServiceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceWebServiceUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointWebService(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceWebServiceRead(ctx, d, meta)
	}
	return diags
}
func resourceWebServiceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceWebServiceDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointWebService(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceWebServiceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceWebServiceRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointWebService(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceWebServicePublicApis(d []interface{}) []edpt.WebServicePublicApis {

	count1 := len(d)
	ret := make([]edpt.WebServicePublicApis, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.WebServicePublicApis
		oi.ApiUri = in["api_uri"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectWebServiceSecure3676(d []interface{}) edpt.WebServiceSecure3676 {

	count1 := len(d)
	var ret edpt.WebServiceSecure3676
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Restart = in["restart"].(int)
		ret.Wipe = in["wipe"].(int)
		ret.Generate = getObjectWebServiceSecureGenerate3677(in["generate"].([]interface{}))
		ret.Regenerate = getObjectWebServiceSecureRegenerate3678(in["regenerate"].([]interface{}))
		ret.Certificate = getObjectWebServiceSecureCertificate3679(in["certificate"].([]interface{}))
		ret.PrivateKey = getObjectWebServiceSecurePrivateKey3680(in["private_key"].([]interface{}))
	}
	return ret
}

func getObjectWebServiceSecureGenerate3677(d []interface{}) edpt.WebServiceSecureGenerate3677 {

	count1 := len(d)
	var ret edpt.WebServiceSecureGenerate3677
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DomainName = in["domain_name"].(string)
		ret.Country = in["country"].(string)
		ret.State = in["state"].(string)
	}
	return ret
}

func getObjectWebServiceSecureRegenerate3678(d []interface{}) edpt.WebServiceSecureRegenerate3678 {

	count1 := len(d)
	var ret edpt.WebServiceSecureRegenerate3678
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DomainName = in["domain_name"].(string)
		ret.Country = in["country"].(string)
		ret.State = in["state"].(string)
	}
	return ret
}

func getObjectWebServiceSecureCertificate3679(d []interface{}) edpt.WebServiceSecureCertificate3679 {

	count1 := len(d)
	var ret edpt.WebServiceSecureCertificate3679
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Load = in["load"].(int)
		ret.UseMgmtPort = in["use_mgmt_port"].(int)
		ret.FileUrl = in["file_url"].(string)
	}
	return ret
}

func getObjectWebServiceSecurePrivateKey3680(d []interface{}) edpt.WebServiceSecurePrivateKey3680 {

	count1 := len(d)
	var ret edpt.WebServiceSecurePrivateKey3680
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Load = in["load"].(int)
		ret.UseMgmtPort = in["use_mgmt_port"].(int)
		ret.FileUrl = in["file_url"].(string)
	}
	return ret
}

func dataToEndpointWebService(d *schema.ResourceData) edpt.WebService {
	var ret edpt.WebService
	ret.Inst.AutoRedirtDisable = d.Get("auto_redirt_disable").(int)
	ret.Inst.AxapiIdle = d.Get("axapi_idle").(int)
	ret.Inst.AxapiSessionLimit = d.Get("axapi_session_limit").(int)
	ret.Inst.GuiIdle = d.Get("gui_idle").(int)
	ret.Inst.GuiSessionLimit = d.Get("gui_session_limit").(int)
	ret.Inst.KeepAliveTimeout = d.Get("keep_alive_timeout").(int)
	ret.Inst.LoginMessage = d.Get("login_message").(string)
	ret.Inst.MaxKeepAliveRequests = d.Get("max_keep_alive_requests").(int)
	ret.Inst.MpmMaxConn = d.Get("mpm_max_conn").(int)
	ret.Inst.MpmMaxConnPerChild = d.Get("mpm_max_conn_per_child").(int)
	ret.Inst.MpmMinSpareConn = d.Get("mpm_min_spare_conn").(int)
	ret.Inst.Port = d.Get("port").(int)
	ret.Inst.PreLoginMessage = d.Get("pre_login_message").(string)
	ret.Inst.PublicApis = getSliceWebServicePublicApis(d.Get("public_apis").([]interface{}))
	ret.Inst.Secure = getObjectWebServiceSecure3676(d.Get("secure").([]interface{}))
	ret.Inst.SecurePort = d.Get("secure_port").(int)
	ret.Inst.SecureServerDisable = d.Get("secure_server_disable").(int)
	ret.Inst.ServerDisable = d.Get("server_disable").(int)
	//omit uuid
	return ret
}
