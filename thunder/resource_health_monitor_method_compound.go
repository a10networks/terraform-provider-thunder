package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceHealthMonitorMethodCompound() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_health_monitor_method_compound`: Compound type\n\n__PLACEHOLDER__",
		CreateContext: resourceHealthMonitorMethodCompoundCreate,
		UpdateContext: resourceHealthMonitorMethodCompoundUpdate,
		ReadContext:   resourceHealthMonitorMethodCompoundRead,
		DeleteContext: resourceHealthMonitorMethodCompoundDelete,

		Schema: map[string]*schema.Schema{
			"compound": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Compound type",
			},
			"rpn_string": {
				Type: schema.TypeString, Optional: true, Description: "Enter Reverse Polish Notation, example: sub hm1 sub hm2 and",
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
func resourceHealthMonitorMethodCompoundCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthMonitorMethodCompoundCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthMonitorMethodCompound(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceHealthMonitorMethodCompoundRead(ctx, d, meta)
	}
	return diags
}

func resourceHealthMonitorMethodCompoundUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthMonitorMethodCompoundUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthMonitorMethodCompound(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceHealthMonitorMethodCompoundRead(ctx, d, meta)
	}
	return diags
}
func resourceHealthMonitorMethodCompoundDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthMonitorMethodCompoundDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthMonitorMethodCompound(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceHealthMonitorMethodCompoundRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthMonitorMethodCompoundRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthMonitorMethodCompound(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointHealthMonitorMethodCompound(d *schema.ResourceData) edpt.HealthMonitorMethodCompound {
	var ret edpt.HealthMonitorMethodCompound
	ret.Inst.Compound = d.Get("compound").(int)
	ret.Inst.RpnString = d.Get("rpn_string").(string)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
