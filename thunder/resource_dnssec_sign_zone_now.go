package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDnssecSignZoneNow() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_dnssec_sign_zone_now`: sign zone right now\n\n__PLACEHOLDER__",
		CreateContext: resourceDnssecSignZoneNowCreate,
		UpdateContext: resourceDnssecSignZoneNowUpdate,
		ReadContext:   resourceDnssecSignZoneNowRead,
		DeleteContext: resourceDnssecSignZoneNowDelete,

		Schema: map[string]*schema.Schema{
			"zone_name": {
				Type: schema.TypeString, Optional: true, Description: "Specify the name for the DNS zone, empty means sign all zones",
			},
		},
	}
}
func resourceDnssecSignZoneNowCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDnssecSignZoneNowCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDnssecSignZoneNow(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDnssecSignZoneNowRead(ctx, d, meta)
	}
	return diags
}

func resourceDnssecSignZoneNowUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDnssecSignZoneNowUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDnssecSignZoneNow(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDnssecSignZoneNowRead(ctx, d, meta)
	}
	return diags
}
func resourceDnssecSignZoneNowDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDnssecSignZoneNowDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDnssecSignZoneNow(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDnssecSignZoneNowRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDnssecSignZoneNowRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDnssecSignZoneNow(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDnssecSignZoneNow(d *schema.ResourceData) edpt.DnssecSignZoneNow {
	var ret edpt.DnssecSignZoneNow
	ret.Inst.ZoneName = d.Get("zone_name").(string)
	return ret
}
