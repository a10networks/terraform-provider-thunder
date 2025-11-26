package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemForcedGroupSpeed() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_forced_group_speed`: Set speed for group of interfaces\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemForcedGroupSpeedCreate,
		UpdateContext: resourceSystemForcedGroupSpeedUpdate,
		ReadContext:   resourceSystemForcedGroupSpeedRead,
		DeleteContext: resourceSystemForcedGroupSpeedDelete,

		Schema: map[string]*schema.Schema{
			"eth01_to_04": {
				Type: schema.TypeInt, Required: true, Description: "Set speed for interface ethernet  1 ~  4",
			},
			"eth05_to_08": {
				Type: schema.TypeInt, Required: true, Description: "Set speed for interface ethernet  5 ~  8",
			},
			"eth09_to_12": {
				Type: schema.TypeInt, Required: true, Description: "Set speed for interface ethernet  9 ~ 12",
			},
			"eth13_to_16": {
				Type: schema.TypeInt, Required: true, Description: "Set speed for interface ethernet 13 ~ 16",
			},
			"eth17_to_20": {
				Type: schema.TypeInt, Required: true, Description: "Set speed for interface ethernet 17 ~ 20",
			},
			"eth21_to_24": {
				Type: schema.TypeInt, Required: true, Description: "Set speed for interface ethernet 21 ~ 24",
			},
			"speed": {
				Type: schema.TypeString, Optional: true, Default: "10G", Description: "'1G': Speed 1G; '10G': Speed 10G (default); '25G': Speed 25G;",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSystemForcedGroupSpeedCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemForcedGroupSpeedCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemForcedGroupSpeed(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemForcedGroupSpeedRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemForcedGroupSpeedUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemForcedGroupSpeedUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemForcedGroupSpeed(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemForcedGroupSpeedRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemForcedGroupSpeedDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemForcedGroupSpeedDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemForcedGroupSpeed(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemForcedGroupSpeedRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemForcedGroupSpeedRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemForcedGroupSpeed(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSystemForcedGroupSpeed(d *schema.ResourceData) edpt.SystemForcedGroupSpeed {
	var ret edpt.SystemForcedGroupSpeed
	ret.Inst.Eth01_to_04 = d.Get("eth01_to_04").(int)
	ret.Inst.Eth05_to_08 = d.Get("eth05_to_08").(int)
	ret.Inst.Eth09_to_12 = d.Get("eth09_to_12").(int)
	ret.Inst.Eth13_to_16 = d.Get("eth13_to_16").(int)
	ret.Inst.Eth17_to_20 = d.Get("eth17_to_20").(int)
	ret.Inst.Eth21_to_24 = d.Get("eth21_to_24").(int)
	ret.Inst.Speed = d.Get("speed").(string)
	//omit uuid
	return ret
}
