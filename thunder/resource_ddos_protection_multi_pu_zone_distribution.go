package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosProtectionMultiPuZoneDistribution() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_protection_multi_pu_zone_distribution`: Multi pu traffic distribution mode configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosProtectionMultiPuZoneDistributionCreate,
		UpdateContext: resourceDdosProtectionMultiPuZoneDistributionUpdate,
		ReadContext:   resourceDdosProtectionMultiPuZoneDistributionRead,
		DeleteContext: resourceDdosProtectionMultiPuZoneDistributionDelete,

		Schema: map[string]*schema.Schema{
			"regular_rebalance": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': enable; 'disable': disable;",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceDdosProtectionMultiPuZoneDistributionCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosProtectionMultiPuZoneDistributionCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosProtectionMultiPuZoneDistribution(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosProtectionMultiPuZoneDistributionRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosProtectionMultiPuZoneDistributionUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosProtectionMultiPuZoneDistributionUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosProtectionMultiPuZoneDistribution(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosProtectionMultiPuZoneDistributionRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosProtectionMultiPuZoneDistributionDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosProtectionMultiPuZoneDistributionDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosProtectionMultiPuZoneDistribution(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosProtectionMultiPuZoneDistributionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosProtectionMultiPuZoneDistributionRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosProtectionMultiPuZoneDistribution(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDdosProtectionMultiPuZoneDistribution(d *schema.ResourceData) edpt.DdosProtectionMultiPuZoneDistribution {
	var ret edpt.DdosProtectionMultiPuZoneDistribution
	ret.Inst.RegularRebalance = d.Get("regular_rebalance").(string)
	//omit uuid
	return ret
}
