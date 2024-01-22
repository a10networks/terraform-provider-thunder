package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDebugGslbProtocol() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_debug_gslb_protocol`: Debug GSLB protocol\n\n__PLACEHOLDER__",
		CreateContext: resourceDebugGslbProtocolCreate,
		UpdateContext: resourceDebugGslbProtocolUpdate,
		ReadContext:   resourceDebugGslbProtocolRead,
		DeleteContext: resourceDebugGslbProtocolDelete,

		Schema: map[string]*schema.Schema{
			"active_rdt": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "ARDT trace information",
			},
			"all": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "All trace information",
			},
			"cache": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Cache trace information",
			},
			"event": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "EVENTS trace information",
			},
			"fsm": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "FSM trace information",
			},
			"ip": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify the remote IP address",
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
			"message_ardt_query": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "GMP ardt-query messages",
			},
			"message_ardt_report": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "GMP ardt-report messages",
			},
			"message_control": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "GMP control messages",
			},
			"message_keepalive": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "GMP keepalive messages",
			},
			"message_notify": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "GMP notify messages",
			},
			"message_open": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "GMP open messages",
			},
			"message_query": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "GMP query messages",
			},
			"message_update": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "GMP update messages",
			},
			"name": {
				Type: schema.TypeString, Optional: true, Description: "Specify the slb device name",
			},
			"normal": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Normal trace information",
			},
			"peer_ipv4": {
				Type: schema.TypeString, Optional: true, Description: "Specify remote IPv4 Address",
			},
			"peer_ipv6": {
				Type: schema.TypeString, Optional: true, Description: "Specify remote IPv6 Address",
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
	}
}
func resourceDebugGslbProtocolCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugGslbProtocolCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugGslbProtocol(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugGslbProtocolRead(ctx, d, meta)
	}
	return diags
}

func resourceDebugGslbProtocolUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugGslbProtocolUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugGslbProtocol(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugGslbProtocolRead(ctx, d, meta)
	}
	return diags
}
func resourceDebugGslbProtocolDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugGslbProtocolDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugGslbProtocol(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDebugGslbProtocolRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugGslbProtocolRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugGslbProtocol(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDebugGslbProtocol(d *schema.ResourceData) edpt.DebugGslbProtocol {
	var ret edpt.DebugGslbProtocol
	ret.Inst.ActiveRdt = d.Get("active_rdt").(int)
	ret.Inst.All = d.Get("all").(int)
	ret.Inst.Cache = d.Get("cache").(int)
	ret.Inst.Event = d.Get("event").(int)
	ret.Inst.Fsm = d.Get("fsm").(int)
	ret.Inst.Ip = d.Get("ip").(int)
	ret.Inst.Ipc = d.Get("ipc").(int)
	ret.Inst.KeepAlive = d.Get("keep_alive").(int)
	ret.Inst.Message = d.Get("message").(int)
	ret.Inst.MessageAll = d.Get("message_all").(int)
	ret.Inst.MessageArdtQuery = d.Get("message_ardt_query").(int)
	ret.Inst.MessageArdtReport = d.Get("message_ardt_report").(int)
	ret.Inst.MessageControl = d.Get("message_control").(int)
	ret.Inst.MessageKeepalive = d.Get("message_keepalive").(int)
	ret.Inst.MessageNotify = d.Get("message_notify").(int)
	ret.Inst.MessageOpen = d.Get("message_open").(int)
	ret.Inst.MessageQuery = d.Get("message_query").(int)
	ret.Inst.MessageUpdate = d.Get("message_update").(int)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.Normal = d.Get("normal").(int)
	ret.Inst.PeerIpv4 = d.Get("peer_ipv4").(string)
	ret.Inst.PeerIpv6 = d.Get("peer_ipv6").(string)
	ret.Inst.Timer = d.Get("timer").(int)
	ret.Inst.Update = d.Get("update").(int)
	//omit uuid
	return ret
}
