package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDnsCacheConfig() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_dns_cache_config`: DNS Cache Settings\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDnsCacheConfigCreate,
		UpdateContext: resourceDdosDnsCacheConfigUpdate,
		ReadContext:   resourceDdosDnsCacheConfigRead,
		DeleteContext: resourceDdosDnsCacheConfigDelete,

		Schema: map[string]*schema.Schema{
			"disable_zone_transfer_in_oper_mode": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable operational refreshing zone transfer",
			},
			"disable_zone_transfer_in_warm_up_mode": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable warm up zone transfer",
			},
			"enable_cache_warm_up_bgp_advertise": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable route injection during cold boot",
			},
			"max_concurrent_zone_transfers": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"warm_up_mode": {
							Type: schema.TypeInt, Optional: true, Default: 65472, Description: "Number of concurrent zone transfers during cold boot (default 65472)",
						},
						"operational_mode": {
							Type: schema.TypeInt, Optional: true, Description: "Number of concurrent zone transfers after boot",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceDdosDnsCacheConfigCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDnsCacheConfigCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDnsCacheConfig(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDnsCacheConfigRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosDnsCacheConfigUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDnsCacheConfigUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDnsCacheConfig(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDnsCacheConfigRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosDnsCacheConfigDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDnsCacheConfigDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDnsCacheConfig(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDnsCacheConfigRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDnsCacheConfigRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDnsCacheConfig(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectDdosDnsCacheConfigMaxConcurrentZoneTransfers145(d []interface{}) edpt.DdosDnsCacheConfigMaxConcurrentZoneTransfers145 {

	count1 := len(d)
	var ret edpt.DdosDnsCacheConfigMaxConcurrentZoneTransfers145
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.WarmUpMode = in["warm_up_mode"].(int)
		ret.OperationalMode = in["operational_mode"].(int)
		//omit uuid
	}
	return ret
}

func dataToEndpointDdosDnsCacheConfig(d *schema.ResourceData) edpt.DdosDnsCacheConfig {
	var ret edpt.DdosDnsCacheConfig
	ret.Inst.DisableZoneTransferInOperMode = d.Get("disable_zone_transfer_in_oper_mode").(int)
	ret.Inst.DisableZoneTransferInWarmUpMode = d.Get("disable_zone_transfer_in_warm_up_mode").(int)
	ret.Inst.EnableCacheWarmUpBgpAdvertise = d.Get("enable_cache_warm_up_bgp_advertise").(int)
	ret.Inst.MaxConcurrentZoneTransfers = getObjectDdosDnsCacheConfigMaxConcurrentZoneTransfers145(d.Get("max_concurrent_zone_transfers").([]interface{}))
	//omit uuid
	return ret
}
