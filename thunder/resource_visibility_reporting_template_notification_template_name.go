package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityReportingTemplateNotificationTemplateName() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_visibility_reporting_template_notification_template_name`: Notification template configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceVisibilityReportingTemplateNotificationTemplateNameCreate,
		UpdateContext: resourceVisibilityReportingTemplateNotificationTemplateNameUpdate,
		ReadContext:   resourceVisibilityReportingTemplateNotificationTemplateNameRead,
		DeleteContext: resourceVisibilityReportingTemplateNotificationTemplateNameDelete,

		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable; 'disable': Disable;",
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
						"auth_password_string": {
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
			"debug_mode": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable debug mode",
			},
			"host_name": {
				Type: schema.TypeString, Optional: true, Description: "Configure the host name(e.g www.a10networks.com)",
			},
			"http_port": {
				Type: schema.TypeInt, Optional: true, Default: 80, Description: "Configure the http port to use(default 80) (http port(default 80))",
			},
			"https_port": {
				Type: schema.TypeInt, Optional: true, Default: 443, Description: "Configure the https port to use(default 443) (http port(default 443))",
			},
			"ipv4_address": {
				Type: schema.TypeString, Optional: true, Description: "Configure the host IPv4 address",
			},
			"ipv6_address": {
				Type: schema.TypeString, Optional: true, Description: "Configure the host IPv6 address",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Notification template name",
			},
			"protocol": {
				Type: schema.TypeString, Optional: true, Default: "https", Description: "'http': Use http protocol; 'https': Use https protocol(default);  (http protocol)",
			},
			"relative_uri": {
				Type: schema.TypeString, Optional: true, Default: "/", Description: "Configure the relative uri(e.g /example , default /)",
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'sent_successful': Sent successful; 'send_fail': Send failures; 'response_fail': Response failures;",
						},
					},
				},
			},
			"test_connectivity": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Test connectivity to notification receiver",
			},
			"use_mgmt_port": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use management port for notifications",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceVisibilityReportingTemplateNotificationTemplateNameCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityReportingTemplateNotificationTemplateNameCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityReportingTemplateNotificationTemplateName(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityReportingTemplateNotificationTemplateNameRead(ctx, d, meta)
	}
	return diags
}

func resourceVisibilityReportingTemplateNotificationTemplateNameUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityReportingTemplateNotificationTemplateNameUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityReportingTemplateNotificationTemplateName(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityReportingTemplateNotificationTemplateNameRead(ctx, d, meta)
	}
	return diags
}
func resourceVisibilityReportingTemplateNotificationTemplateNameDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityReportingTemplateNotificationTemplateNameDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityReportingTemplateNotificationTemplateName(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVisibilityReportingTemplateNotificationTemplateNameRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityReportingTemplateNotificationTemplateNameRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityReportingTemplateNotificationTemplateName(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectVisibilityReportingTemplateNotificationTemplateNameAuthentication3126(d []interface{}) edpt.VisibilityReportingTemplateNotificationTemplateNameAuthentication3126 {

	count1 := len(d)
	var ret edpt.VisibilityReportingTemplateNotificationTemplateNameAuthentication3126
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RelativeLoginUri = in["relative_login_uri"].(string)
		ret.RelativeLogoffUri = in["relative_logoff_uri"].(string)
		ret.AuthUsername = in["auth_username"].(string)
		ret.AuthPassword = in["auth_password"].(int)
		ret.AuthPasswordString = in["auth_password_string"].(string)
		//omit encrypted
		ret.ApiKey = in["api_key"].(int)
		ret.ApiKeyString = in["api_key_string"].(string)
		//omit api_key_encrypted
		//omit uuid
	}
	return ret
}

func getSliceVisibilityReportingTemplateNotificationTemplateNameSamplingEnable(d []interface{}) []edpt.VisibilityReportingTemplateNotificationTemplateNameSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.VisibilityReportingTemplateNotificationTemplateNameSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityReportingTemplateNotificationTemplateNameSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointVisibilityReportingTemplateNotificationTemplateName(d *schema.ResourceData) edpt.VisibilityReportingTemplateNotificationTemplateName {
	var ret edpt.VisibilityReportingTemplateNotificationTemplateName
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.Authentication = getObjectVisibilityReportingTemplateNotificationTemplateNameAuthentication3126(d.Get("authentication").([]interface{}))
	ret.Inst.DebugMode = d.Get("debug_mode").(int)
	ret.Inst.HostName = d.Get("host_name").(string)
	ret.Inst.HttpPort = d.Get("http_port").(int)
	ret.Inst.HttpsPort = d.Get("https_port").(int)
	ret.Inst.Ipv4Address = d.Get("ipv4_address").(string)
	ret.Inst.Ipv6Address = d.Get("ipv6_address").(string)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.Protocol = d.Get("protocol").(string)
	ret.Inst.RelativeUri = d.Get("relative_uri").(string)
	ret.Inst.SamplingEnable = getSliceVisibilityReportingTemplateNotificationTemplateNameSamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.TestConnectivity = d.Get("test_connectivity").(int)
	ret.Inst.UseMgmtPort = d.Get("use_mgmt_port").(int)
	//omit uuid
	return ret
}
