package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplateTcp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_template_tcp`: L4 TCP switch config\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbTemplateTcpCreate,
		UpdateContext: resourceSlbTemplateTcpUpdate,
		ReadContext:   resourceSlbTemplateTcpRead,
		DeleteContext: resourceSlbTemplateTcpDelete,

		Schema: map[string]*schema.Schema{
			"alive_if_active": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "keep connection alive if active traffic",
			},
			"del_session_on_server_down": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Delete session if the server/port goes down (either disabled/hm down)",
			},
			"disable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "send reset to client when server is disabled",
			},
			"down": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "send reset to client when server is down",
			},
			"force_delete_timeout": {
				Type: schema.TypeInt, Optional: true, Description: "The maximum time that a session can stay in the system before being delete (number (second))",
			},
			"force_delete_timeout_100ms": {
				Type: schema.TypeInt, Optional: true, Description: "The maximum time that a session can stay in the system before being delete (number in 100ms)",
			},
			"half_close_idle_timeout": {
				Type: schema.TypeInt, Optional: true, Description: "TCP Half Close Idle Timeout (sec), default off (half close idle timeout in second, default off)",
			},
			"half_open_idle_timeout": {
				Type: schema.TypeInt, Optional: true, Description: "TCP Half Open Idle Timeout (sec), default off (half open idle timeout in second, default off)",
			},
			"idle_timeout": {
				Type: schema.TypeInt, Optional: true, Default: 120, Description: "Idle Timeout value (Interval of 60 seconds), default 120 seconds (idle timeout in second, default 120)",
			},
			"initial_window_size": {
				Type: schema.TypeInt, Optional: true, Description: "Set the initial window size (number)",
			},
			"insert_client_ip": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Insert client ip into TCP option",
			},
			"lan_fast_ack": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable fast TCP ack on LAN",
			},
			"logging": {
				Type: schema.TypeString, Optional: true, Description: "'init': init only log; 'term': termination only log; 'both': both initial and termination log;",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Fast TCP Template Name",
			},
			"proxy_header": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"proxy_header_action": {
							Type: schema.TypeString, Optional: true, Description: "'insert': Insert proxy header;",
						},
						"proxy_header_version": {
							Type: schema.TypeString, Optional: true, Description: "'v1': version 1; 'v2': version 2;",
						},
					},
				},
			},
			"qos": {
				Type: schema.TypeInt, Optional: true, Description: "QOS level (number)",
			},
			"re_select_if_server_down": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "re-select another server if service port is down",
			},
			"reset_follow_fin": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "send reset to client or server upon receiving first fin",
			},
			"reset_fwd": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "send reset to server if error happens",
			},
			"reset_rev": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "send reset to client if error happens",
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
func resourceSlbTemplateTcpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateTcpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateTcp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateTcpRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbTemplateTcpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateTcpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateTcp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateTcpRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbTemplateTcpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateTcpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateTcp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbTemplateTcpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateTcpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateTcp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectSlbTemplateTcpProxyHeader(d []interface{}) edpt.SlbTemplateTcpProxyHeader {

	count1 := len(d)
	var ret edpt.SlbTemplateTcpProxyHeader
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ProxyHeaderAction = in["proxy_header_action"].(string)
		ret.ProxyHeaderVersion = in["proxy_header_version"].(string)
	}
	return ret
}

func dataToEndpointSlbTemplateTcp(d *schema.ResourceData) edpt.SlbTemplateTcp {
	var ret edpt.SlbTemplateTcp
	ret.Inst.AliveIfActive = d.Get("alive_if_active").(int)
	ret.Inst.DelSessionOnServerDown = d.Get("del_session_on_server_down").(int)
	ret.Inst.Disable = d.Get("disable").(int)
	ret.Inst.Down = d.Get("down").(int)
	ret.Inst.ForceDeleteTimeout = d.Get("force_delete_timeout").(int)
	ret.Inst.ForceDeleteTimeout100ms = d.Get("force_delete_timeout_100ms").(int)
	ret.Inst.HalfCloseIdleTimeout = d.Get("half_close_idle_timeout").(int)
	ret.Inst.HalfOpenIdleTimeout = d.Get("half_open_idle_timeout").(int)
	ret.Inst.IdleTimeout = d.Get("idle_timeout").(int)
	ret.Inst.InitialWindowSize = d.Get("initial_window_size").(int)
	ret.Inst.InsertClientIp = d.Get("insert_client_ip").(int)
	ret.Inst.LanFastAck = d.Get("lan_fast_ack").(int)
	ret.Inst.Logging = d.Get("logging").(string)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.ProxyHeader = getObjectSlbTemplateTcpProxyHeader(d.Get("proxy_header").([]interface{}))
	ret.Inst.Qos = d.Get("qos").(int)
	ret.Inst.ReSelectIfServerDown = d.Get("re_select_if_server_down").(int)
	ret.Inst.ResetFollowFin = d.Get("reset_follow_fin").(int)
	ret.Inst.ResetFwd = d.Get("reset_fwd").(int)
	ret.Inst.ResetRev = d.Get("reset_rev").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
