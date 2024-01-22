package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDebugGslb() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_debug_gslb`: Debug GSLB\n\n__PLACEHOLDER__",
		CreateContext: resourceDebugGslbCreate,
		UpdateContext: resourceDebugGslbUpdate,
		ReadContext:   resourceDebugGslbRead,
		DeleteContext: resourceDebugGslbDelete,

		Schema: map[string]*schema.Schema{
			"glname": {
				Type: schema.TypeString, Optional: true, Description: "Debug for matched Geo-location",
			},
			"group": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cache": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Cache trace information",
						},
						"event": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "EVENTS trace information",
						},
						"all": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "All trace information",
						},
						"fsm": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "FSM trace information",
						},
						"ip": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Remote IP address of GSLB controller of the group",
						},
						"peer_ipv4": {
							Type: schema.TypeString, Optional: true, Description: "Specify remote IPv4 Address",
						},
						"peer_ipv6": {
							Type: schema.TypeString, Optional: true, Description: "Specify remote IPv6 Address",
						},
						"ipc": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "IPC trace information",
						},
						"keep_alive": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "KEEPALIVE trace information",
						},
						"message": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Debug GSLB group message",
						},
						"message_all": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "All messages",
						},
						"message_keepalive": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "GMP keepalive messages",
						},
						"message_notify": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "GMP notify messages",
						},
						"message_control": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "GMP control messages",
						},
						"message_query": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "GMP query messages",
						},
						"message_open": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "GMP open messages",
						},
						"message_group": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "GMP group messages",
						},
						"name": {
							Type: schema.TypeString, Optional: true, Description: "Name of GSLB group",
						},
						"normal": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Normal trace information",
						},
						"timer": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Timer trace information",
						},
						"update": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Update trace information",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"id1": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify ID",
			},
			"ip_addr": {
				Type: schema.TypeString, Optional: true, Description: "Debug for matched IP address",
			},
			"ipv6_addr": {
				Type: schema.TypeString, Optional: true, Description: "Debug for matched IPv6 address",
			},
			"memory": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Debug GSLB memory",
			},
			"one_shot": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Stop after get 64 states",
			},
			"protocol": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cache": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Cache trace information",
						},
						"event": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "EVENTS trace information",
						},
						"all": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "All trace information",
						},
						"active_rdt": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "ARDT trace information",
						},
						"fsm": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "FSM trace information",
						},
						"ip": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify the remote IP address",
						},
						"peer_ipv4": {
							Type: schema.TypeString, Optional: true, Description: "Specify remote IPv4 Address",
						},
						"peer_ipv6": {
							Type: schema.TypeString, Optional: true, Description: "Specify remote IPv6 Address",
						},
						"ipc": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "IPC trace information",
						},
						"keep_alive": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "KEEPALIVE trace information",
						},
						"message": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Debug GSLB protocol message",
						},
						"message_all": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "All messages",
						},
						"message_keepalive": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "GMP keepalive messages",
						},
						"message_notify": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "GMP notify messages",
						},
						"message_control": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "GMP control messages",
						},
						"message_query": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "GMP query messages",
						},
						"message_open": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "GMP open messages",
						},
						"message_update": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "GMP update messages",
						},
						"message_ardt_query": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "GMP ardt-query messages",
						},
						"message_ardt_report": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "GMP ardt-report messages",
						},
						"name": {
							Type: schema.TypeString, Optional: true, Description: "Specify the slb device name",
						},
						"normal": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Normal trace information",
						},
						"timer": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Timer trace information",
						},
						"update": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Update trace information",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"state": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Keep GSLB state information",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceDebugGslbCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugGslbCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugGslb(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugGslbRead(ctx, d, meta)
	}
	return diags
}

func resourceDebugGslbUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugGslbUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugGslb(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugGslbRead(ctx, d, meta)
	}
	return diags
}
func resourceDebugGslbDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugGslbDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugGslb(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDebugGslbRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugGslbRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugGslb(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectDebugGslbGroup323(d []interface{}) edpt.DebugGslbGroup323 {

	count1 := len(d)
	var ret edpt.DebugGslbGroup323
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Cache = in["cache"].(int)
		ret.Event = in["event"].(int)
		ret.All = in["all"].(int)
		ret.Fsm = in["fsm"].(int)
		ret.Ip = in["ip"].(int)
		ret.PeerIpv4 = in["peer_ipv4"].(string)
		ret.PeerIpv6 = in["peer_ipv6"].(string)
		ret.Ipc = in["ipc"].(int)
		ret.KeepAlive = in["keep_alive"].(int)
		ret.Message = in["message"].(int)
		ret.MessageAll = in["message_all"].(int)
		ret.MessageKeepalive = in["message_keepalive"].(int)
		ret.MessageNotify = in["message_notify"].(int)
		ret.MessageControl = in["message_control"].(int)
		ret.MessageQuery = in["message_query"].(int)
		ret.MessageOpen = in["message_open"].(int)
		ret.MessageGroup = in["message_group"].(int)
		ret.Name = in["name"].(string)
		ret.Normal = in["normal"].(int)
		ret.Timer = in["timer"].(int)
		ret.Update = in["update"].(int)
		//omit uuid
	}
	return ret
}

func getObjectDebugGslbProtocol324(d []interface{}) edpt.DebugGslbProtocol324 {

	count1 := len(d)
	var ret edpt.DebugGslbProtocol324
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Cache = in["cache"].(int)
		ret.Event = in["event"].(int)
		ret.All = in["all"].(int)
		ret.ActiveRdt = in["active_rdt"].(int)
		ret.Fsm = in["fsm"].(int)
		ret.Ip = in["ip"].(int)
		ret.PeerIpv4 = in["peer_ipv4"].(string)
		ret.PeerIpv6 = in["peer_ipv6"].(string)
		ret.Ipc = in["ipc"].(int)
		ret.KeepAlive = in["keep_alive"].(int)
		ret.Message = in["message"].(int)
		ret.MessageAll = in["message_all"].(int)
		ret.MessageKeepalive = in["message_keepalive"].(int)
		ret.MessageNotify = in["message_notify"].(int)
		ret.MessageControl = in["message_control"].(int)
		ret.MessageQuery = in["message_query"].(int)
		ret.MessageOpen = in["message_open"].(int)
		ret.MessageUpdate = in["message_update"].(int)
		ret.MessageArdtQuery = in["message_ardt_query"].(int)
		ret.MessageArdtReport = in["message_ardt_report"].(int)
		ret.Name = in["name"].(string)
		ret.Normal = in["normal"].(int)
		ret.Timer = in["timer"].(int)
		ret.Update = in["update"].(int)
		//omit uuid
	}
	return ret
}

func dataToEndpointDebugGslb(d *schema.ResourceData) edpt.DebugGslb {
	var ret edpt.DebugGslb
	ret.Inst.Glname = d.Get("glname").(string)
	ret.Inst.Group = getObjectDebugGslbGroup323(d.Get("group").([]interface{}))
	ret.Inst.Id1 = d.Get("id1").(int)
	ret.Inst.IpAddr = d.Get("ip_addr").(string)
	ret.Inst.Ipv6Addr = d.Get("ipv6_addr").(string)
	ret.Inst.Memory = d.Get("memory").(int)
	ret.Inst.OneShot = d.Get("one_shot").(int)
	ret.Inst.Protocol = getObjectDebugGslbProtocol324(d.Get("protocol").([]interface{}))
	ret.Inst.State = d.Get("state").(int)
	//omit uuid
	return ret
}
