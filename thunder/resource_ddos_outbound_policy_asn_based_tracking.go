package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosOutboundPolicyAsnBasedTracking() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_outbound_policy_asn_based_tracking`: Configure asn based tracking\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosOutboundPolicyAsnBasedTrackingCreate,
		UpdateContext: resourceDdosOutboundPolicyAsnBasedTrackingUpdate,
		ReadContext:   resourceDdosOutboundPolicyAsnBasedTrackingRead,
		DeleteContext: resourceDdosOutboundPolicyAsnBasedTrackingDelete,

		Schema: map[string]*schema.Schema{
			"configuration": {
				Type: schema.TypeString, Required: true, Description: "'configuration': Configure asn based tracking;",
			},
			"packet_rate_triggered": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Triggered by 1/2 packet rate limitation in per-asn-glid.",
			},
			"per_asn_glid": {
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
func resourceDdosOutboundPolicyAsnBasedTrackingCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosOutboundPolicyAsnBasedTrackingCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosOutboundPolicyAsnBasedTracking(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosOutboundPolicyAsnBasedTrackingRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosOutboundPolicyAsnBasedTrackingUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosOutboundPolicyAsnBasedTrackingUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosOutboundPolicyAsnBasedTracking(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosOutboundPolicyAsnBasedTrackingRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosOutboundPolicyAsnBasedTrackingDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosOutboundPolicyAsnBasedTrackingDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosOutboundPolicyAsnBasedTracking(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosOutboundPolicyAsnBasedTrackingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosOutboundPolicyAsnBasedTrackingRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosOutboundPolicyAsnBasedTracking(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDdosOutboundPolicyAsnBasedTracking(d *schema.ResourceData) edpt.DdosOutboundPolicyAsnBasedTracking {
	var ret edpt.DdosOutboundPolicyAsnBasedTracking
	ret.Inst.Configuration = d.Get("configuration").(string)
	ret.Inst.PacketRateTriggered = d.Get("packet_rate_triggered").(int)
	ret.Inst.PerAsnGlid = d.Get("per_asn_glid").(string)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
