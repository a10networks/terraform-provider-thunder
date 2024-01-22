package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSysUtEventActionTcpFlags() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_sys_ut_event_action_tcp_flags`: TCP flags\n\n__PLACEHOLDER__",
		CreateContext: resourceSysUtEventActionTcpFlagsCreate,
		UpdateContext: resourceSysUtEventActionTcpFlagsUpdate,
		ReadContext:   resourceSysUtEventActionTcpFlagsRead,
		DeleteContext: resourceSysUtEventActionTcpFlagsDelete,

		Schema: map[string]*schema.Schema{
			"ack": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Ack",
			},
			"cwr": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Cwr",
			},
			"ece": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Ece",
			},
			"fin": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Fin",
			},
			"psh": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Psh",
			},
			"rst": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Rst",
			},
			"syn": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Syn",
			},
			"urg": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Urg",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
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
func resourceSysUtEventActionTcpFlagsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSysUtEventActionTcpFlagsCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSysUtEventActionTcpFlags(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSysUtEventActionTcpFlagsRead(ctx, d, meta)
	}
	return diags
}

func resourceSysUtEventActionTcpFlagsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSysUtEventActionTcpFlagsUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSysUtEventActionTcpFlags(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSysUtEventActionTcpFlagsRead(ctx, d, meta)
	}
	return diags
}
func resourceSysUtEventActionTcpFlagsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSysUtEventActionTcpFlagsDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSysUtEventActionTcpFlags(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSysUtEventActionTcpFlagsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSysUtEventActionTcpFlagsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSysUtEventActionTcpFlags(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSysUtEventActionTcpFlags(d *schema.ResourceData) edpt.SysUtEventActionTcpFlags {
	var ret edpt.SysUtEventActionTcpFlags
	ret.Inst.Ack = d.Get("ack").(int)
	ret.Inst.Cwr = d.Get("cwr").(int)
	ret.Inst.Ece = d.Get("ece").(int)
	ret.Inst.Fin = d.Get("fin").(int)
	ret.Inst.Psh = d.Get("psh").(int)
	ret.Inst.Rst = d.Get("rst").(int)
	ret.Inst.Syn = d.Get("syn").(int)
	ret.Inst.Urg = d.Get("urg").(int)
	//omit uuid
	ret.Inst.EventNumber = d.Get("event_number").(string)
	ret.Inst.Direction = d.Get("direction").(string)
	return ret
}
