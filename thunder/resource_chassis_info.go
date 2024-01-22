package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceChassisInfo() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_chassis_info`: Info regarding the chassis state\n\n__PLACEHOLDER__",
		CreateContext: resourceChassisInfoCreate,
		UpdateContext: resourceChassisInfoUpdate,
		ReadContext:   resourceChassisInfoRead,
		DeleteContext: resourceChassisInfoDelete,

		Schema: map[string]*schema.Schema{
			"brief": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Give brief Chassis State Info",
			},
		},
	}
}
func resourceChassisInfoCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceChassisInfoCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointChassisInfo(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceChassisInfoRead(ctx, d, meta)
	}
	return diags
}

func resourceChassisInfoUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceChassisInfoUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointChassisInfo(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceChassisInfoRead(ctx, d, meta)
	}
	return diags
}
func resourceChassisInfoDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceChassisInfoDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointChassisInfo(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceChassisInfoRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceChassisInfoRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointChassisInfo(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointChassisInfo(d *schema.ResourceData) edpt.ChassisInfo {
	var ret edpt.ChassisInfo
	ret.Inst.Brief = d.Get("brief").(int)
	return ret
}
