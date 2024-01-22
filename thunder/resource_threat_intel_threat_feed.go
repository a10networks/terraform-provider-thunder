package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceThreatIntelThreatFeed() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_threat_intel_threat_feed`: Configure vendor specific module control options\n\n__PLACEHOLDER__",
		CreateContext: resourceThreatIntelThreatFeedCreate,
		UpdateContext: resourceThreatIntelThreatFeedUpdate,
		ReadContext:   resourceThreatIntelThreatFeedRead,
		DeleteContext: resourceThreatIntelThreatFeedDelete,

		Schema: map[string]*schema.Schema{
			"domain": {
				Type: schema.TypeString, Optional: true, Description: "Realm for NTLM authentication",
			},
			"enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable module",
			},
			"log_level": {
				Type: schema.TypeString, Optional: true, Default: "warning", Description: "'disable': Disable all logging; 'error': Log error events; 'warning': Log warning events and above; 'info': Log info events and above; 'debug': Log debug events and above; 'trace': enable all logs;",
			},
			"port": {
				Type: schema.TypeInt, Optional: true, Default: 443, Description: "Port to query server(default 443)",
			},
			"proxy_auth_type": {
				Type: schema.TypeString, Optional: true, Default: "ntlm", Description: "'ntlm': NTLM authentication(default); 'basic': Basic authentication;",
			},
			"proxy_host": {
				Type: schema.TypeString, Optional: true, Description: "Proxy server hostname or IP address",
			},
			"proxy_password": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Password for proxy authentication",
			},
			"proxy_port": {
				Type: schema.TypeInt, Optional: true, Description: "Port to connect on proxy server",
			},
			"proxy_username": {
				Type: schema.TypeString, Optional: true, Description: "Username for proxy authentication",
			},
			"rtu_update_disable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disables real time updates(default enable)",
			},
			"secret_string": {
				Type: schema.TypeString, Optional: true, Description: "password value",
			},
			"server": {
				Type: schema.TypeString, Optional: true, Description: "Server IP or Hostname",
			},
			"server_timeout": {
				Type: schema.TypeInt, Optional: true, Default: 15, Description: "Server Timeout in seconds (default: 15s)",
			},
			"type": {
				Type: schema.TypeString, Required: true, Description: "'webroot': Configure Webroot module options;",
			},
			"update_interval": {
				Type: schema.TypeInt, Optional: true, Default: 120, Description: "Interval to check for database or RTU updates(default 120 mins)",
			},
			"use_mgmt_port": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use management interface for all communication with threat-intel server",
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
func resourceThreatIntelThreatFeedCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceThreatIntelThreatFeedCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointThreatIntelThreatFeed(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceThreatIntelThreatFeedRead(ctx, d, meta)
	}
	return diags
}

func resourceThreatIntelThreatFeedUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceThreatIntelThreatFeedUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointThreatIntelThreatFeed(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceThreatIntelThreatFeedRead(ctx, d, meta)
	}
	return diags
}
func resourceThreatIntelThreatFeedDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceThreatIntelThreatFeedDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointThreatIntelThreatFeed(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceThreatIntelThreatFeedRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceThreatIntelThreatFeedRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointThreatIntelThreatFeed(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointThreatIntelThreatFeed(d *schema.ResourceData) edpt.ThreatIntelThreatFeed {
	var ret edpt.ThreatIntelThreatFeed
	ret.Inst.Domain = d.Get("domain").(string)
	ret.Inst.Enable = d.Get("enable").(int)
	//omit encrypted
	ret.Inst.LogLevel = d.Get("log_level").(string)
	ret.Inst.Port = d.Get("port").(int)
	ret.Inst.ProxyAuthType = d.Get("proxy_auth_type").(string)
	ret.Inst.ProxyHost = d.Get("proxy_host").(string)
	ret.Inst.ProxyPassword = d.Get("proxy_password").(int)
	ret.Inst.ProxyPort = d.Get("proxy_port").(int)
	ret.Inst.ProxyUsername = d.Get("proxy_username").(string)
	ret.Inst.RtuUpdateDisable = d.Get("rtu_update_disable").(int)
	ret.Inst.SecretString = d.Get("secret_string").(string)
	ret.Inst.Server = d.Get("server").(string)
	ret.Inst.ServerTimeout = d.Get("server_timeout").(int)
	ret.Inst.Type = d.Get("type").(string)
	ret.Inst.UpdateInterval = d.Get("update_interval").(int)
	ret.Inst.UseMgmtPort = d.Get("use_mgmt_port").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
