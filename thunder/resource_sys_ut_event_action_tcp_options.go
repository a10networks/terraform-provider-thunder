package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSysUtEventActionTcpOptions() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_sys_ut_event_action_tcp_options`: Tcp Flags\n\n__PLACEHOLDER__",
		CreateContext: resourceSysUtEventActionTcpOptionsCreate,
		UpdateContext: resourceSysUtEventActionTcpOptionsUpdate,
		ReadContext:   resourceSysUtEventActionTcpOptionsRead,
		DeleteContext: resourceSysUtEventActionTcpOptionsDelete,

		Schema: map[string]*schema.Schema{
			"mss": {
				Type: schema.TypeInt, Optional: true, Description: "TCP MSS",
			},
			"nop": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "No Op",
			},
			"sack_type": {
				Type: schema.TypeString, Optional: true, Description: "'permitted': permitted; 'block': block;",
			},
			"time_stamp_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "adds Time Stamp to options",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"wscale": {
				Type: schema.TypeInt, Optional: true, Description: "",
			},
			"event_number": {
				Type: schema.TypeString, Required: true, Description: "EventNumber",
			},
			"direction": {
				Type: schema.TypeString, Required: true, Description: "Direction",
			},
		},
	}
}
func resourceSysUtEventActionTcpOptionsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSysUtEventActionTcpOptionsCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSysUtEventActionTcpOptions(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSysUtEventActionTcpOptionsRead(ctx, d, meta)
	}
	return diags
}

func resourceSysUtEventActionTcpOptionsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSysUtEventActionTcpOptionsUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSysUtEventActionTcpOptions(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSysUtEventActionTcpOptionsRead(ctx, d, meta)
	}
	return diags
}
func resourceSysUtEventActionTcpOptionsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSysUtEventActionTcpOptionsDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSysUtEventActionTcpOptions(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSysUtEventActionTcpOptionsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSysUtEventActionTcpOptionsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSysUtEventActionTcpOptions(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSysUtEventActionTcpOptions(d *schema.ResourceData) edpt.SysUtEventActionTcpOptions {
	var ret edpt.SysUtEventActionTcpOptions
	ret.Inst.Mss = d.Get("mss").(int)
	ret.Inst.Nop = d.Get("nop").(int)
	ret.Inst.SackType = d.Get("sack_type").(string)
	ret.Inst.TimeStampEnable = d.Get("time_stamp_enable").(int)
	//omit uuid
	ret.Inst.Wscale = d.Get("wscale").(int)
	ret.Inst.EventNumber = d.Get("event_number").(string)
	ret.Inst.Direction = d.Get("direction").(string)
	return ret
}
