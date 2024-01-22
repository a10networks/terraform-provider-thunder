package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSysUtStateNextStateCaseActionTcpFlags() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_sys_ut_state_next_state_case_action_tcp_flags`: TCP flags\n\n__PLACEHOLDER__",
		CreateContext: resourceSysUtStateNextStateCaseActionTcpFlagsCreate,
		UpdateContext: resourceSysUtStateNextStateCaseActionTcpFlagsUpdate,
		ReadContext:   resourceSysUtStateNextStateCaseActionTcpFlagsRead,
		DeleteContext: resourceSysUtStateNextStateCaseActionTcpFlagsDelete,

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
func resourceSysUtStateNextStateCaseActionTcpFlagsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSysUtStateNextStateCaseActionTcpFlagsCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSysUtStateNextStateCaseActionTcpFlags(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSysUtStateNextStateCaseActionTcpFlagsRead(ctx, d, meta)
	}
	return diags
}

func resourceSysUtStateNextStateCaseActionTcpFlagsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSysUtStateNextStateCaseActionTcpFlagsUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSysUtStateNextStateCaseActionTcpFlags(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSysUtStateNextStateCaseActionTcpFlagsRead(ctx, d, meta)
	}
	return diags
}
func resourceSysUtStateNextStateCaseActionTcpFlagsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSysUtStateNextStateCaseActionTcpFlagsDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSysUtStateNextStateCaseActionTcpFlags(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSysUtStateNextStateCaseActionTcpFlagsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSysUtStateNextStateCaseActionTcpFlagsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSysUtStateNextStateCaseActionTcpFlags(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSysUtStateNextStateCaseActionTcpFlags(d *schema.ResourceData) edpt.SysUtStateNextStateCaseActionTcpFlags {
	var ret edpt.SysUtStateNextStateCaseActionTcpFlags
	ret.Inst.Ack = d.Get("ack").(int)
	ret.Inst.Cwr = d.Get("cwr").(int)
	ret.Inst.Ece = d.Get("ece").(int)
	ret.Inst.Fin = d.Get("fin").(int)
	ret.Inst.Psh = d.Get("psh").(int)
	ret.Inst.Rst = d.Get("rst").(int)
	ret.Inst.Syn = d.Get("syn").(int)
	ret.Inst.Urg = d.Get("urg").(int)
	//omit uuid
	ret.Inst.CaseNumber = d.Get("case_number").(string)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.Direction = d.Get("direction").(string)
	return ret
}
