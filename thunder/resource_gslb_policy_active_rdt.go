package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGslbPolicyActiveRdt() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_gslb_policy_active_rdt`: Select SLB device with the shortest round delay time to local DNS\n\n__PLACEHOLDER__",
		CreateContext: resourceGslbPolicyActiveRdtCreate,
		UpdateContext: resourceGslbPolicyActiveRdtUpdate,
		ReadContext:   resourceGslbPolicyActiveRdtRead,
		DeleteContext: resourceGslbPolicyActiveRdtDelete,

		Schema: map[string]*schema.Schema{
			"controller": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Active round-delay-time by controller",
			},
			"difference": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "The difference between the round-delay-time, default is 0",
			},
			"enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable the active rdt",
			},
			"fail_break": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Break when no valid RDT",
			},
			"ignore_id": {
				Type: schema.TypeInt, Optional: true, Description: "Ignore IP Address specified in IP List by ID",
			},
			"keep_tracking": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Keep tracking client even round-delay-time samples are ready",
			},
			"limit": {
				Type: schema.TypeInt, Optional: true, Default: 16383, Description: "Limit of allowed RDT, default is 16383 (Limit, unit: millisecond)",
			},
			"proto_rdt_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable the round-delay-time to the controller",
			},
			"samples": {
				Type: schema.TypeInt, Optional: true, Default: 5, Description: "Specify samples number for round-delay-time (Number of samples,default is 5)",
			},
			"single_shot": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Single Shot RDT",
			},
			"skip": {
				Type: schema.TypeInt, Optional: true, Default: 3, Description: "Skip query if round-delay-time samples are not ready (Specify maximum skip count,default is 3)",
			},
			"timeout": {
				Type: schema.TypeInt, Optional: true, Default: 3, Description: "Specify timeout if round-delay-time samples are not ready (Specify timeout, unit:sec,default is 3)",
			},
			"tolerance": {
				Type: schema.TypeInt, Optional: true, Default: 10, Description: "The difference percentage between the round-delay-time, default is 10 (Tolerance)",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}
func resourceGslbPolicyActiveRdtCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbPolicyActiveRdtCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbPolicyActiveRdt(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGslbPolicyActiveRdtRead(ctx, d, meta)
	}
	return diags
}

func resourceGslbPolicyActiveRdtUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbPolicyActiveRdtUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbPolicyActiveRdt(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGslbPolicyActiveRdtRead(ctx, d, meta)
	}
	return diags
}
func resourceGslbPolicyActiveRdtDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbPolicyActiveRdtDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbPolicyActiveRdt(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceGslbPolicyActiveRdtRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbPolicyActiveRdtRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbPolicyActiveRdt(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointGslbPolicyActiveRdt(d *schema.ResourceData) edpt.GslbPolicyActiveRdt {
	var ret edpt.GslbPolicyActiveRdt
	ret.Inst.Controller = d.Get("controller").(int)
	ret.Inst.Difference = d.Get("difference").(int)
	ret.Inst.Enable = d.Get("enable").(int)
	ret.Inst.FailBreak = d.Get("fail_break").(int)
	ret.Inst.IgnoreId = d.Get("ignore_id").(int)
	ret.Inst.KeepTracking = d.Get("keep_tracking").(int)
	ret.Inst.Limit = d.Get("limit").(int)
	ret.Inst.ProtoRdtEnable = d.Get("proto_rdt_enable").(int)
	ret.Inst.Samples = d.Get("samples").(int)
	ret.Inst.SingleShot = d.Get("single_shot").(int)
	ret.Inst.Skip = d.Get("skip").(int)
	ret.Inst.Timeout = d.Get("timeout").(int)
	ret.Inst.Tolerance = d.Get("tolerance").(int)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
