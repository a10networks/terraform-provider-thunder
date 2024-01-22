package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGslbPolicyCapacity() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_gslb_policy_capacity`: Select Service-IP for the device having highest available connection capacity\n\n__PLACEHOLDER__",
		CreateContext: resourceGslbPolicyCapacityCreate,
		UpdateContext: resourceGslbPolicyCapacityUpdate,
		ReadContext:   resourceGslbPolicyCapacityRead,
		DeleteContext: resourceGslbPolicyCapacityDelete,

		Schema: map[string]*schema.Schema{
			"capacity_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable capacity",
			},
			"capacity_fail_break": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Break when exceed threshold",
			},
			"threshold": {
				Type: schema.TypeInt, Optional: true, Default: 90, Description: "Specify capacity threshold, default is 90",
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
func resourceGslbPolicyCapacityCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbPolicyCapacityCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbPolicyCapacity(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGslbPolicyCapacityRead(ctx, d, meta)
	}
	return diags
}

func resourceGslbPolicyCapacityUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbPolicyCapacityUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbPolicyCapacity(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGslbPolicyCapacityRead(ctx, d, meta)
	}
	return diags
}
func resourceGslbPolicyCapacityDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbPolicyCapacityDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbPolicyCapacity(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceGslbPolicyCapacityRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbPolicyCapacityRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbPolicyCapacity(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointGslbPolicyCapacity(d *schema.ResourceData) edpt.GslbPolicyCapacity {
	var ret edpt.GslbPolicyCapacity
	ret.Inst.CapacityEnable = d.Get("capacity_enable").(int)
	ret.Inst.CapacityFailBreak = d.Get("capacity_fail_break").(int)
	ret.Inst.Threshold = d.Get("threshold").(int)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
