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
			"cpu_threshold_per_entry": {
				Type: schema.TypeInt, Optional: true, Default: 60, Description: "Entry/zone percentage threshold of CPU usage for source hash mode. Requires distribution-method cpu-usage. Default:60",
			},
			"cpu_threshold_per_pu": {
				Type: schema.TypeInt, Optional: true, Default: 80, Description: "Per PU percentage threshold of average CPU usage to start check entry usage. Requires distribution-method cpu-usage. Default:80",
			},
			"distribution_method": {
				Type: schema.TypeString, Optional: true, Default: "traffic-rate", Description: "'cpu-usage': Entry/Zone distribution based on CPU usage percentage; 'traffic-rate': Entry/Zone distribution based on traffic kbit/pkt rate (Default);",
			},
			"rate_kbit_threshold": {
				Type: schema.TypeInt, Optional: true, Default: 150000000, Description: "DDOS DST Entry/Zone kbit rate threshold for source hash mode",
			},
			"rate_pkt_threshold": {
				Type: schema.TypeInt, Optional: true, Default: 55000000, Description: "DDOS DST Entry/Zone packet rate threshold for source hash mode",
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
	ret.Inst.CpuThresholdPerEntry = d.Get("cpu_threshold_per_entry").(int)
	ret.Inst.CpuThresholdPerPu = d.Get("cpu_threshold_per_pu").(int)
	ret.Inst.DistributionMethod = d.Get("distribution_method").(string)
	ret.Inst.RateKbitThreshold = d.Get("rate_kbit_threshold").(int)
	ret.Inst.RatePktThreshold = d.Get("rate_pkt_threshold").(int)
	//omit uuid
	return ret
}
