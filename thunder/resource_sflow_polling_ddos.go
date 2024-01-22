package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSflowPollingDdos() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_sflow_polling_ddos`: Poll DDoS counters\n\n__PLACEHOLDER__",
		CreateContext: resourceSflowPollingDdosCreate,
		UpdateContext: resourceSflowPollingDdosUpdate,
		ReadContext:   resourceSflowPollingDdosRead,
		DeleteContext: resourceSflowPollingDdosDelete,

		Schema: map[string]*schema.Schema{
			"address_byte_order_host": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Export sflow address field in host byte order",
			},
			"compatibility2_9": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable DDOS sflow polling 2.9 compatibility mode",
			},
			"compatibility3_0": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable DDOS sflow polling 3.0/3.1 compatibility mode",
			},
			"dns_cache_zone_stats": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable polling for dns cache per instance and per zone statistics",
			},
			"dyn_entry_stats": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable polling for dynamic entry statistics",
			},
			"enable_anomaly_stats": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable Polling for system wide anomaly statistics",
			},
			"toggle": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable sflow polling for DDOS statistics; 'disable': Disable sflow polling for DDOS statistics;",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSflowPollingDdosCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSflowPollingDdosCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSflowPollingDdos(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSflowPollingDdosRead(ctx, d, meta)
	}
	return diags
}

func resourceSflowPollingDdosUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSflowPollingDdosUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSflowPollingDdos(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSflowPollingDdosRead(ctx, d, meta)
	}
	return diags
}
func resourceSflowPollingDdosDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSflowPollingDdosDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSflowPollingDdos(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSflowPollingDdosRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSflowPollingDdosRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSflowPollingDdos(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSflowPollingDdos(d *schema.ResourceData) edpt.SflowPollingDdos {
	var ret edpt.SflowPollingDdos
	ret.Inst.AddressByteOrderHost = d.Get("address_byte_order_host").(int)
	ret.Inst.Compatibility2_9 = d.Get("compatibility2_9").(int)
	ret.Inst.Compatibility3_0 = d.Get("compatibility3_0").(int)
	ret.Inst.DnsCacheZoneStats = d.Get("dns_cache_zone_stats").(int)
	ret.Inst.DynEntryStats = d.Get("dyn_entry_stats").(int)
	ret.Inst.EnableAnomalyStats = d.Get("enable_anomaly_stats").(int)
	ret.Inst.Toggle = d.Get("toggle").(string)
	//omit uuid
	return ret
}
