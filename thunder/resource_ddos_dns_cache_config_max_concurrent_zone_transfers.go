package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDnsCacheConfigMaxConcurrentZoneTransfers() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_dns_cache_config_max_concurrent_zone_transfers`: Cold-Boot configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDnsCacheConfigMaxConcurrentZoneTransfersCreate,
		UpdateContext: resourceDdosDnsCacheConfigMaxConcurrentZoneTransfersUpdate,
		ReadContext:   resourceDdosDnsCacheConfigMaxConcurrentZoneTransfersRead,
		DeleteContext: resourceDdosDnsCacheConfigMaxConcurrentZoneTransfersDelete,

		Schema: map[string]*schema.Schema{
			"operational_mode": {
				Type: schema.TypeInt, Optional: true, Description: "Number of concurrent zone transfers after boot",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"warm_up_mode": {
				Type: schema.TypeInt, Optional: true, Default: 65472, Description: "Number of concurrent zone transfers during cold boot (default 65472)",
			},
		},
	}
}
func resourceDdosDnsCacheConfigMaxConcurrentZoneTransfersCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDnsCacheConfigMaxConcurrentZoneTransfersCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDnsCacheConfigMaxConcurrentZoneTransfers(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDnsCacheConfigMaxConcurrentZoneTransfersRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosDnsCacheConfigMaxConcurrentZoneTransfersUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDnsCacheConfigMaxConcurrentZoneTransfersUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDnsCacheConfigMaxConcurrentZoneTransfers(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDnsCacheConfigMaxConcurrentZoneTransfersRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosDnsCacheConfigMaxConcurrentZoneTransfersDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDnsCacheConfigMaxConcurrentZoneTransfersDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDnsCacheConfigMaxConcurrentZoneTransfers(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDnsCacheConfigMaxConcurrentZoneTransfersRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDnsCacheConfigMaxConcurrentZoneTransfersRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDnsCacheConfigMaxConcurrentZoneTransfers(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDdosDnsCacheConfigMaxConcurrentZoneTransfers(d *schema.ResourceData) edpt.DdosDnsCacheConfigMaxConcurrentZoneTransfers {
	var ret edpt.DdosDnsCacheConfigMaxConcurrentZoneTransfers
	ret.Inst.OperationalMode = d.Get("operational_mode").(int)
	//omit uuid
	ret.Inst.WarmUpMode = d.Get("warm_up_mode").(int)
	return ret
}
