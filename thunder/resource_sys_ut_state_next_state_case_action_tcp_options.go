package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSysUtStateNextStateCaseActionTcpOptions() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_sys_ut_state_next_state_case_action_tcp_options`: Tcp Flags\n\n__PLACEHOLDER__",
		CreateContext: resourceSysUtStateNextStateCaseActionTcpOptionsCreate,
		UpdateContext: resourceSysUtStateNextStateCaseActionTcpOptionsUpdate,
		ReadContext:   resourceSysUtStateNextStateCaseActionTcpOptionsRead,
		DeleteContext: resourceSysUtStateNextStateCaseActionTcpOptionsDelete,

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
			"case_number": {
				Type: schema.TypeString, Required: true, Description: "CaseNumber",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
			"direction": {
				Type: schema.TypeString, Required: true, Description: "Direction",
			},
		},
	}
}
func resourceSysUtStateNextStateCaseActionTcpOptionsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSysUtStateNextStateCaseActionTcpOptionsCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSysUtStateNextStateCaseActionTcpOptions(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSysUtStateNextStateCaseActionTcpOptionsRead(ctx, d, meta)
	}
	return diags
}

func resourceSysUtStateNextStateCaseActionTcpOptionsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSysUtStateNextStateCaseActionTcpOptionsUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSysUtStateNextStateCaseActionTcpOptions(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSysUtStateNextStateCaseActionTcpOptionsRead(ctx, d, meta)
	}
	return diags
}
func resourceSysUtStateNextStateCaseActionTcpOptionsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSysUtStateNextStateCaseActionTcpOptionsDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSysUtStateNextStateCaseActionTcpOptions(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSysUtStateNextStateCaseActionTcpOptionsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSysUtStateNextStateCaseActionTcpOptionsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSysUtStateNextStateCaseActionTcpOptions(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSysUtStateNextStateCaseActionTcpOptions(d *schema.ResourceData) edpt.SysUtStateNextStateCaseActionTcpOptions {
	var ret edpt.SysUtStateNextStateCaseActionTcpOptions
	ret.Inst.Mss = d.Get("mss").(int)
	ret.Inst.Nop = d.Get("nop").(int)
	ret.Inst.SackType = d.Get("sack_type").(string)
	ret.Inst.TimeStampEnable = d.Get("time_stamp_enable").(int)
	//omit uuid
	ret.Inst.Wscale = d.Get("wscale").(int)
	ret.Inst.CaseNumber = d.Get("case_number").(string)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.Direction = d.Get("direction").(string)
	return ret
}
