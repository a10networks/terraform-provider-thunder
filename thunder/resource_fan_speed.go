package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFanSpeed() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_fan_speed`: Control FAN Speed setting\n\n__PLACEHOLDER__",
		CreateContext: resourceFanSpeedCreate,
		UpdateContext: resourceFanSpeedUpdate,
		ReadContext:   resourceFanSpeedRead,
		DeleteContext: resourceFanSpeedDelete,

		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeString, Optional: true, Default: "auto", Description: "'auto': Set FAN Speed to automatic based on system temperature; '50': Set FAN Speed to 50% of maximum speed (approx); '75': Set FAN Speed to 75% of maximum speed (approx); '100': Set FAN Speed to 100% full hardware speed;",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceFanSpeedCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFanSpeedCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFanSpeed(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFanSpeedRead(ctx, d, meta)
	}
	return diags
}

func resourceFanSpeedUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFanSpeedUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFanSpeed(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFanSpeedRead(ctx, d, meta)
	}
	return diags
}
func resourceFanSpeedDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFanSpeedDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFanSpeed(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFanSpeedRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFanSpeedRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFanSpeed(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointFanSpeed(d *schema.ResourceData) edpt.FanSpeed {
	var ret edpt.FanSpeed
	ret.Inst.Action = d.Get("action").(string)
	//omit uuid
	return ret
}
