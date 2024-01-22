package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosOutboundPolicyCountryBasedTracking() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_outbound_policy_country_based_tracking`: Configure country based tracking\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosOutboundPolicyCountryBasedTrackingCreate,
		UpdateContext: resourceDdosOutboundPolicyCountryBasedTrackingUpdate,
		ReadContext:   resourceDdosOutboundPolicyCountryBasedTrackingRead,
		DeleteContext: resourceDdosOutboundPolicyCountryBasedTrackingDelete,

		Schema: map[string]*schema.Schema{
			"configuration": {
				Type: schema.TypeString, Required: true, Description: "'configuration': Configure country based tracking;",
			},
			"packet_rate_triggered": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Triggered by 1/2 packet rate limitation in per-country-glid.",
			},
			"per_country_glid": {
				Type: schema.TypeString, Optional: true, Description: "Global limit ID",
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
func resourceDdosOutboundPolicyCountryBasedTrackingCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosOutboundPolicyCountryBasedTrackingCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosOutboundPolicyCountryBasedTracking(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosOutboundPolicyCountryBasedTrackingRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosOutboundPolicyCountryBasedTrackingUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosOutboundPolicyCountryBasedTrackingUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosOutboundPolicyCountryBasedTracking(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosOutboundPolicyCountryBasedTrackingRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosOutboundPolicyCountryBasedTrackingDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosOutboundPolicyCountryBasedTrackingDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosOutboundPolicyCountryBasedTracking(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosOutboundPolicyCountryBasedTrackingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosOutboundPolicyCountryBasedTrackingRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosOutboundPolicyCountryBasedTracking(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDdosOutboundPolicyCountryBasedTracking(d *schema.ResourceData) edpt.DdosOutboundPolicyCountryBasedTracking {
	var ret edpt.DdosOutboundPolicyCountryBasedTracking
	ret.Inst.Configuration = d.Get("configuration").(string)
	ret.Inst.PacketRateTriggered = d.Get("packet_rate_triggered").(int)
	ret.Inst.PerCountryGlid = d.Get("per_country_glid").(string)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
