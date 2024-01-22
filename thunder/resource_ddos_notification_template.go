package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosNotificationTemplate() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_notification_template`: Notification template configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosNotificationTemplateCreate,
		UpdateContext: resourceDdosNotificationTemplateUpdate,
		ReadContext:   resourceDdosNotificationTemplateRead,
		DeleteContext: resourceDdosNotificationTemplateDelete,

		Schema: map[string]*schema.Schema{
			"api": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"host_ipv4_address": {
							Type: schema.TypeString, Optional: true, Description: "Configure the host IPv4 address to send notification (IPv4 address of the host)",
						},
						"host_ipv6_address": {
							Type: schema.TypeString, Optional: true, Description: "Configure the host IPv6 address to send notification (IPv6 address of the host)",
						},
						"hostname": {
							Type: schema.TypeString, Optional: true, Description: "host name(e.g www.a10networks.com)",
						},
						"http_protocol": {
							Type: schema.TypeString, Optional: true, Default: "https", Description: "'http': Use http protocol; 'https': Use https protocol(default);  (http protocol)",
						},
						"http_port": {
							Type: schema.TypeInt, Optional: true, Default: 80, Description: "Configure the http port to use(default 80) (http port(default 80))",
						},
						"https_port": {
							Type: schema.TypeInt, Optional: true, Default: 443, Description: "Configure the https port to use(default 443) (https port(default 443))",
						},
						"timeout": {
							Type: schema.TypeInt, Optional: true, Default: 10, Description: "Configure the api execution timeout(default 10secs) (api timeout)",
						},
						"relative_uri": {
							Type: schema.TypeString, Optional: true, Default: "/", Description: "Configure the relative uri for the api(e.g /example , default /) (api relative uri)",
						},
						"disable_authentication": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable authentication to communicate to the host",
						},
						"use_mgmt_port": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use management port to send out notification",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
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
					},
				},
			},
			"debug_mode": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable debug mode",
			},
			"disable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable the notification template (Disable notification temaplate)",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "DDOS nofitication template name",
			},
			"test_connectivity": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Test connectivity to notification receiver",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"verbose": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Dump zone IPs to the payload",
			},
		},
	}
}
func resourceDdosNotificationTemplateCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNotificationTemplateCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNotificationTemplate(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosNotificationTemplateRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosNotificationTemplateUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNotificationTemplateUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNotificationTemplate(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosNotificationTemplateRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosNotificationTemplateDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNotificationTemplateDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNotificationTemplate(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosNotificationTemplateRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNotificationTemplateRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNotificationTemplate(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectDdosNotificationTemplateApi285(d []interface{}) edpt.DdosNotificationTemplateApi285 {

	count1 := len(d)
	var ret edpt.DdosNotificationTemplateApi285
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.HostIpv4Address = in["host_ipv4_address"].(string)
		ret.HostIpv6Address = in["host_ipv6_address"].(string)
		ret.Hostname = in["hostname"].(string)
		ret.HttpProtocol = in["http_protocol"].(string)
		ret.HttpPort = in["http_port"].(int)
		ret.HttpsPort = in["https_port"].(int)
		ret.Timeout = in["timeout"].(int)
		ret.RelativeUri = in["relative_uri"].(string)
		ret.DisableAuthentication = in["disable_authentication"].(int)
		ret.UseMgmtPort = in["use_mgmt_port"].(int)
		//omit uuid
		ret.Authentication = getObjectDdosNotificationTemplateApiAuthentication286(in["authentication"].([]interface{}))
	}
	return ret
}

func getObjectDdosNotificationTemplateApiAuthentication286(d []interface{}) edpt.DdosNotificationTemplateApiAuthentication286 {

	count1 := len(d)
	var ret edpt.DdosNotificationTemplateApiAuthentication286
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

func dataToEndpointDdosNotificationTemplate(d *schema.ResourceData) edpt.DdosNotificationTemplate {
	var ret edpt.DdosNotificationTemplate
	ret.Inst.Api = getObjectDdosNotificationTemplateApi285(d.Get("api").([]interface{}))
	ret.Inst.DebugMode = d.Get("debug_mode").(int)
	ret.Inst.Disable = d.Get("disable").(int)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.TestConnectivity = d.Get("test_connectivity").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.Verbose = d.Get("verbose").(int)
	return ret
}
