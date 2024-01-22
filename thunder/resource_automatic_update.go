package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAutomaticUpdate() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_automatic_update`: Automatic update configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceAutomaticUpdateCreate,
		UpdateContext: resourceAutomaticUpdateUpdate,
		ReadContext:   resourceAutomaticUpdateRead,
		DeleteContext: resourceAutomaticUpdateDelete,

		Schema: map[string]*schema.Schema{
			"check_now": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"feature_name": {
							Type: schema.TypeString, Optional: true, Description: "'app-fw': Application Firewall; 'ca-bundle': CA Certificate Bundle; 'a10-threat-intel': A10 Threat intel class list; 'central-cert-pin-list': Central updated cert pinning list;",
						},
						"prod_ver": {
							Type: schema.TypeString, Optional: true, Description: "update to this specific version, if this option is not configured, update to the latest version",
						},
						"from_staging_server": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Get files from GLM Staging storage",
						},
						"stage_ver": {
							Type: schema.TypeString, Optional: true, Description: "update this specific version",
						},
					},
				},
			},
			"checknow": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"config_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"feature_name": {
							Type: schema.TypeString, Required: true, Description: "'app-fw': Application Firewall Configuration; 'ca-bundle': CA Certificate Bundle; 'a10-threat-intel': A10 Threat intel class list; 'central-cert-pin-list': Central updated cert pinning list;",
						},
						"debug": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable libcurl debug option",
						},
						"disable_ssl_verify": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable peer server certificate verification",
						},
						"schedule": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "",
						},
						"weekly": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Every week",
						},
						"week_day": {
							Type: schema.TypeString, Optional: true, Description: "'Monday': Monday; 'Tuesday': Tuesday; 'Wednesday': Wednesday; 'Thursday': Thursday; 'Friday': Friday; 'Saturday': Saturday; 'Sunday': Sunday;",
						},
						"week_time": {
							Type: schema.TypeString, Optional: true, Description: "Time of day to update (hh:mm) in 24 hour local time",
						},
						"daily": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Every day",
						},
						"day_time": {
							Type: schema.TypeString, Optional: true, Description: "Time of day to update (hh:mm) in 24 hour local time",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"glm_source_url": {
				Type: schema.TypeString, Optional: true, Description: "Change GLM source url",
			},
			"info": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"proxy_server": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"proxy_host": {
							Type: schema.TypeString, Optional: true, Description: "Proxy server hostname or IP address",
						},
						"https_port": {
							Type: schema.TypeInt, Optional: true, Description: "Proxy server HTTPs port",
						},
						"auth_type": {
							Type: schema.TypeString, Optional: true, Default: "ntlm", Description: "'ntlm': NTLM authentication(default); 'basic': Basic authentication;",
						},
						"domain": {
							Type: schema.TypeString, Optional: true, Description: "Realm for NTLM authentication",
						},
						"username": {
							Type: schema.TypeString, Optional: true, Description: "Username for proxy authentication",
						},
						"password": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Password for proxy authentication",
						},
						"secret_string": {
							Type: schema.TypeString, Optional: true, Description: "password value",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"reset": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"feature_name": {
							Type: schema.TypeString, Optional: true, Description: "'app-fw': Application Firewall; 'ca-bundle': CA Certificate Bundle;",
						},
					},
				},
			},
			"revert": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"feature_name": {
							Type: schema.TypeString, Optional: true, Description: "'app-fw': Application Firewall; 'a10-threat-intel': A10 Threat intel class list; 'central-cert-pin-list': Central updated cert pinning list;",
						},
					},
				},
			},
			"use_mgmt_port": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use management port to connect",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceAutomaticUpdateCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAutomaticUpdateCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAutomaticUpdate(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAutomaticUpdateRead(ctx, d, meta)
	}
	return diags
}

func resourceAutomaticUpdateUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAutomaticUpdateUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAutomaticUpdate(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAutomaticUpdateRead(ctx, d, meta)
	}
	return diags
}
func resourceAutomaticUpdateDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAutomaticUpdateDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAutomaticUpdate(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceAutomaticUpdateRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAutomaticUpdateRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAutomaticUpdate(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectAutomaticUpdateCheckNow69(d []interface{}) edpt.AutomaticUpdateCheckNow69 {

	count1 := len(d)
	var ret edpt.AutomaticUpdateCheckNow69
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.FeatureName = in["feature_name"].(string)
		ret.ProdVer = in["prod_ver"].(string)
		ret.FromStagingServer = in["from_staging_server"].(int)
		ret.StageVer = in["stage_ver"].(string)
	}
	return ret
}

func getObjectAutomaticUpdateChecknow70(d []interface{}) edpt.AutomaticUpdateChecknow70 {

	var ret edpt.AutomaticUpdateChecknow70
	return ret
}

func getSliceAutomaticUpdateConfigList(d []interface{}) []edpt.AutomaticUpdateConfigList {

	count1 := len(d)
	ret := make([]edpt.AutomaticUpdateConfigList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AutomaticUpdateConfigList
		oi.FeatureName = in["feature_name"].(string)
		oi.Debug = in["debug"].(int)
		oi.DisableSslVerify = in["disable_ssl_verify"].(int)
		oi.Schedule = in["schedule"].(int)
		oi.Weekly = in["weekly"].(int)
		oi.WeekDay = in["week_day"].(string)
		oi.WeekTime = in["week_time"].(string)
		oi.Daily = in["daily"].(int)
		oi.DayTime = in["day_time"].(string)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getObjectAutomaticUpdateInfo71(d []interface{}) edpt.AutomaticUpdateInfo71 {

	var ret edpt.AutomaticUpdateInfo71
	return ret
}

func getObjectAutomaticUpdateProxyServer72(d []interface{}) edpt.AutomaticUpdateProxyServer72 {

	count1 := len(d)
	var ret edpt.AutomaticUpdateProxyServer72
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ProxyHost = in["proxy_host"].(string)
		ret.HttpsPort = in["https_port"].(int)
		ret.AuthType = in["auth_type"].(string)
		ret.Domain = in["domain"].(string)
		ret.Username = in["username"].(string)
		ret.Password = in["password"].(int)
		ret.SecretString = in["secret_string"].(string)
		//omit encrypted
		//omit uuid
	}
	return ret
}

func getObjectAutomaticUpdateReset73(d []interface{}) edpt.AutomaticUpdateReset73 {

	count1 := len(d)
	var ret edpt.AutomaticUpdateReset73
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.FeatureName = in["feature_name"].(string)
	}
	return ret
}

func getObjectAutomaticUpdateRevert74(d []interface{}) edpt.AutomaticUpdateRevert74 {

	count1 := len(d)
	var ret edpt.AutomaticUpdateRevert74
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.FeatureName = in["feature_name"].(string)
	}
	return ret
}

func dataToEndpointAutomaticUpdate(d *schema.ResourceData) edpt.AutomaticUpdate {
	var ret edpt.AutomaticUpdate
	ret.Inst.CheckNow = getObjectAutomaticUpdateCheckNow69(d.Get("check_now").([]interface{}))
	ret.Inst.Checknow = getObjectAutomaticUpdateChecknow70(d.Get("checknow").([]interface{}))
	ret.Inst.ConfigList = getSliceAutomaticUpdateConfigList(d.Get("config_list").([]interface{}))
	ret.Inst.GlmSourceUrl = d.Get("glm_source_url").(string)
	ret.Inst.Info = getObjectAutomaticUpdateInfo71(d.Get("info").([]interface{}))
	ret.Inst.ProxyServer = getObjectAutomaticUpdateProxyServer72(d.Get("proxy_server").([]interface{}))
	ret.Inst.Reset = getObjectAutomaticUpdateReset73(d.Get("reset").([]interface{}))
	ret.Inst.Revert = getObjectAutomaticUpdateRevert74(d.Get("revert").([]interface{}))
	ret.Inst.UseMgmtPort = d.Get("use_mgmt_port").(int)
	//omit uuid
	return ret
}
