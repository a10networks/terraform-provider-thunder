package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosNotificationTemplateApi() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_notification_template_api`: Violation action notification api configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosNotificationTemplateApiCreate,
		UpdateContext: resourceDdosNotificationTemplateApiUpdate,
		ReadContext:   resourceDdosNotificationTemplateApiRead,
		DeleteContext: resourceDdosNotificationTemplateApiDelete,

		Schema: map[string]*schema.Schema{
			"authentication": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"relative_login_uri": {
							Type: schema.TypeString, Optional: true, Description: "Configure the authentication login uri",
						},
						"relative_logoff_uri": {
							Type: schema.TypeString, Optional: true, Description: "Configure the authentication logoff uri",
						},
						"auth_username": {
							Type: schema.TypeString, Optional: true, Description: "Configure the authentication user name",
						},
						"auth_password": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure the authentication user password (Authentication password)",
						},
						"auth_password_val": {
							Type: schema.TypeString, Optional: true, Description: "Configure the authentication user password (Authentication password)",
						},
						"api_key": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure api-key as a mode of authentication",
						},
						"api_key_string": {
							Type: schema.TypeString, Optional: true, Description: "Configure api-key as a mode of authentication",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"disable_authentication": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable authentication to communicate to the host",
			},
			"host_ipv4_address": {
				Type: schema.TypeString, Optional: true, Description: "Configure the host IPv4 address to send notification (IPv4 address of the host)",
			},
			"host_ipv6_address": {
				Type: schema.TypeString, Optional: true, Description: "Configure the host IPv6 address to send notification (IPv6 address of the host)",
			},
			"hostname": {
				Type: schema.TypeString, Optional: true, Description: "host name(e.g www.a10networks.com)",
			},
			"http_port": {
				Type: schema.TypeInt, Optional: true, Default: 80, Description: "Configure the http port to use(default 80) (http port(default 80))",
			},
			"http_protocol": {
				Type: schema.TypeString, Optional: true, Default: "https", Description: "'http': Use http protocol; 'https': Use https protocol(default);  (http protocol)",
			},
			"https_port": {
				Type: schema.TypeInt, Optional: true, Default: 443, Description: "Configure the https port to use(default 443) (https port(default 443))",
			},
			"relative_uri": {
				Type: schema.TypeString, Optional: true, Default: "/", Description: "Configure the relative uri for the api(e.g /example , default /) (api relative uri)",
			},
			"timeout": {
				Type: schema.TypeInt, Optional: true, Default: 10, Description: "Configure the api execution timeout(default 10secs) (api timeout)",
			},
			"use_mgmt_port": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use management port to send out notification",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}
func resourceDdosNotificationTemplateApiCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNotificationTemplateApiCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNotificationTemplateApi(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosNotificationTemplateApiRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosNotificationTemplateApiUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNotificationTemplateApiUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNotificationTemplateApi(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosNotificationTemplateApiRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosNotificationTemplateApiDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNotificationTemplateApiDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNotificationTemplateApi(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosNotificationTemplateApiRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNotificationTemplateApiRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNotificationTemplateApi(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectDdosNotificationTemplateApiAuthentication284(d []interface{}) edpt.DdosNotificationTemplateApiAuthentication284 {

	count1 := len(d)
	var ret edpt.DdosNotificationTemplateApiAuthentication284
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RelativeLoginUri = in["relative_login_uri"].(string)
		ret.RelativeLogoffUri = in["relative_logoff_uri"].(string)
		ret.AuthUsername = in["auth_username"].(string)
		ret.AuthPassword = in["auth_password"].(int)
		ret.AuthPasswordVal = in["auth_password_val"].(string)
		//omit encrypted
		ret.ApiKey = in["api_key"].(int)
		ret.ApiKeyString = in["api_key_string"].(string)
		//omit api_key_encrypted
		//omit uuid
	}
	return ret
}

func dataToEndpointDdosNotificationTemplateApi(d *schema.ResourceData) edpt.DdosNotificationTemplateApi {
	var ret edpt.DdosNotificationTemplateApi
	ret.Inst.Authentication = getObjectDdosNotificationTemplateApiAuthentication284(d.Get("authentication").([]interface{}))
	ret.Inst.DisableAuthentication = d.Get("disable_authentication").(int)
	ret.Inst.HostIpv4Address = d.Get("host_ipv4_address").(string)
	ret.Inst.HostIpv6Address = d.Get("host_ipv6_address").(string)
	ret.Inst.Hostname = d.Get("hostname").(string)
	ret.Inst.HttpPort = d.Get("http_port").(int)
	ret.Inst.HttpProtocol = d.Get("http_protocol").(string)
	ret.Inst.HttpsPort = d.Get("https_port").(int)
	ret.Inst.RelativeUri = d.Get("relative_uri").(string)
	ret.Inst.Timeout = d.Get("timeout").(int)
	ret.Inst.UseMgmtPort = d.Get("use_mgmt_port").(int)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
