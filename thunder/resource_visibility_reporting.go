package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityReporting() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_visibility_reporting`: Configure reporting framework\n\n__PLACEHOLDER__",
		CreateContext: resourceVisibilityReportingCreate,
		UpdateContext: resourceVisibilityReportingUpdate,
		ReadContext:   resourceVisibilityReportingRead,
		DeleteContext: resourceVisibilityReportingDelete,

		Schema: map[string]*schema.Schema{
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'log-transmit-failure': Total log transmit failures; 'buffer-alloc-failure': Total reporting buffer allocation failures; 'notif-jobs-in-queue': Total notification jobs in queue; 'enqueue-fail': Total enqueue jobs failed; 'enqueue-pass': Total enqueue jobs passed; 'dequeued': Total jobs dequeued;",
						},
					},
				},
			},
			"session_logging": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable session logging; 'disable': Disable session logging(default);",
			},
			"telemetry_export_interval": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"value": {
							Type: schema.TypeInt, Optional: true, Default: 5, Description: "Monitored entity telemetry data export interval in mins (Default 5 mins)",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"template": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"notification": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"template_name_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": {
													Type: schema.TypeString, Required: true, Description: "Notification template name",
												},
												"ipv4_address": {
													Type: schema.TypeString, Optional: true, Description: "Configure the host IPv4 address",
												},
												"ipv6_address": {
													Type: schema.TypeString, Optional: true, Description: "Configure the host IPv6 address",
												},
												"host_name": {
													Type: schema.TypeString, Optional: true, Description: "Configure the host name(e.g www.a10networks.com)",
												},
												"use_mgmt_port": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use management port for notifications",
												},
												"protocol": {
													Type: schema.TypeString, Optional: true, Default: "https", Description: "'http': Use http protocol; 'https': Use https protocol(default);  (http protocol)",
												},
												"http_port": {
													Type: schema.TypeInt, Optional: true, Default: 80, Description: "Configure the http port to use(default 80) (http port(default 80))",
												},
												"https_port": {
													Type: schema.TypeInt, Optional: true, Default: 443, Description: "Configure the https port to use(default 443) (http port(default 443))",
												},
												"relative_uri": {
													Type: schema.TypeString, Optional: true, Default: "/", Description: "Configure the relative uri(e.g /example , default /)",
												},
												"action": {
													Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable; 'disable': Disable;",
												},
												"debug_mode": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable debug mode",
												},
												"test_connectivity": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Test connectivity to notification receiver",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
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
											},
										},
									},
									"debug": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
								},
							},
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
func resourceVisibilityReportingCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityReportingCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityReporting(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityReportingRead(ctx, d, meta)
	}
	return diags
}

func resourceVisibilityReportingUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityReportingUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityReporting(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityReportingRead(ctx, d, meta)
	}
	return diags
}
func resourceVisibilityReportingDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityReportingDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityReporting(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVisibilityReportingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityReportingRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityReporting(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceVisibilityReportingSamplingEnable(d []interface{}) []edpt.VisibilityReportingSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.VisibilityReportingSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityReportingSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectVisibilityReportingTelemetryExportInterval3127(d []interface{}) edpt.VisibilityReportingTelemetryExportInterval3127 {

	count1 := len(d)
	var ret edpt.VisibilityReportingTelemetryExportInterval3127
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Value = in["value"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityReportingTemplate3128(d []interface{}) edpt.VisibilityReportingTemplate3128 {

	count1 := len(d)
	var ret edpt.VisibilityReportingTemplate3128
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Notification = getObjectVisibilityReportingTemplateNotification3129(in["notification"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityReportingTemplateNotification3129(d []interface{}) edpt.VisibilityReportingTemplateNotification3129 {

	count1 := len(d)
	var ret edpt.VisibilityReportingTemplateNotification3129
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.TemplateNameList = getSliceVisibilityReportingTemplateNotificationTemplateNameList(in["template_name_list"].([]interface{}))
		ret.Debug = getObjectVisibilityReportingTemplateNotificationDebug3130(in["debug"].([]interface{}))
	}
	return ret
}

func getSliceVisibilityReportingTemplateNotificationTemplateNameList(d []interface{}) []edpt.VisibilityReportingTemplateNotificationTemplateNameList {

	count1 := len(d)
	ret := make([]edpt.VisibilityReportingTemplateNotificationTemplateNameList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityReportingTemplateNotificationTemplateNameList
		oi.Name = in["name"].(string)
		oi.Ipv4Address = in["ipv4_address"].(string)
		oi.Ipv6Address = in["ipv6_address"].(string)
		oi.HostName = in["host_name"].(string)
		oi.UseMgmtPort = in["use_mgmt_port"].(int)
		oi.Protocol = in["protocol"].(string)
		oi.HttpPort = in["http_port"].(int)
		oi.HttpsPort = in["https_port"].(int)
		oi.RelativeUri = in["relative_uri"].(string)
		oi.Action = in["action"].(string)
		oi.DebugMode = in["debug_mode"].(int)
		oi.TestConnectivity = in["test_connectivity"].(int)
		//omit uuid
		oi.SamplingEnable = getSliceVisibilityReportingTemplateNotificationTemplateNameListSamplingEnable(in["sampling_enable"].([]interface{}))
		oi.Authentication = getObjectVisibilityReportingTemplateNotificationTemplateNameListAuthentication(in["authentication"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceVisibilityReportingTemplateNotificationTemplateNameListSamplingEnable(d []interface{}) []edpt.VisibilityReportingTemplateNotificationTemplateNameListSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.VisibilityReportingTemplateNotificationTemplateNameListSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityReportingTemplateNotificationTemplateNameListSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectVisibilityReportingTemplateNotificationTemplateNameListAuthentication(d []interface{}) edpt.VisibilityReportingTemplateNotificationTemplateNameListAuthentication {

	count1 := len(d)
	var ret edpt.VisibilityReportingTemplateNotificationTemplateNameListAuthentication
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

func getObjectVisibilityReportingTemplateNotificationDebug3130(d []interface{}) edpt.VisibilityReportingTemplateNotificationDebug3130 {

	var ret edpt.VisibilityReportingTemplateNotificationDebug3130
	return ret
}

func dataToEndpointVisibilityReporting(d *schema.ResourceData) edpt.VisibilityReporting {
	var ret edpt.VisibilityReporting
	ret.Inst.SamplingEnable = getSliceVisibilityReportingSamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.SessionLogging = d.Get("session_logging").(string)
	ret.Inst.TelemetryExportInterval = getObjectVisibilityReportingTelemetryExportInterval3127(d.Get("telemetry_export_interval").([]interface{}))
	ret.Inst.Template = getObjectVisibilityReportingTemplate3128(d.Get("template").([]interface{}))
	//omit uuid
	return ret
}
