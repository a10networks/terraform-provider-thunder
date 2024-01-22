package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSnmpServerSlbDataCacheTimeout() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_snmp_server_slb_data_cache_timeout`: Configure the SLB data cache time-out in seconds\n\n__PLACEHOLDER__",
		CreateContext: resourceSnmpServerSlbDataCacheTimeoutCreate,
		UpdateContext: resourceSnmpServerSlbDataCacheTimeoutUpdate,
		ReadContext:   resourceSnmpServerSlbDataCacheTimeoutRead,
		DeleteContext: resourceSnmpServerSlbDataCacheTimeoutDelete,

		Schema: map[string]*schema.Schema{
			"slblimit": {
				Type: schema.TypeInt, Optional: true, Description: "Cache time-out in seconds, default as 60 seconds",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSnmpServerSlbDataCacheTimeoutCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerSlbDataCacheTimeoutCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerSlbDataCacheTimeout(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSnmpServerSlbDataCacheTimeoutRead(ctx, d, meta)
	}
	return diags
}

func resourceSnmpServerSlbDataCacheTimeoutUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerSlbDataCacheTimeoutUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerSlbDataCacheTimeout(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSnmpServerSlbDataCacheTimeoutRead(ctx, d, meta)
	}
	return diags
}
func resourceSnmpServerSlbDataCacheTimeoutDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerSlbDataCacheTimeoutDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerSlbDataCacheTimeout(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSnmpServerSlbDataCacheTimeoutRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerSlbDataCacheTimeoutRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerSlbDataCacheTimeout(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSnmpServerSlbDataCacheTimeout(d *schema.ResourceData) edpt.SnmpServerSlbDataCacheTimeout {
	var ret edpt.SnmpServerSlbDataCacheTimeout
	ret.Inst.Slblimit = d.Get("slblimit").(int)
	//omit uuid
	return ret
}
