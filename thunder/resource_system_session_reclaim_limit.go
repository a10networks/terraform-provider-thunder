package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemSessionReclaimLimit() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_session_reclaim_limit`: Set limits for smp session reclaim\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemSessionReclaimLimitCreate,
		UpdateContext: resourceSystemSessionReclaimLimitUpdate,
		ReadContext:   resourceSystemSessionReclaimLimitRead,
		DeleteContext: resourceSystemSessionReclaimLimitDelete,

		Schema: map[string]*schema.Schema{
			"nscan_limit": {
				Type: schema.TypeInt, Optional: true, Default: 4096, Description: "smp session scan limit (number of smp sessions per scan)",
			},
			"scan_freq": {
				Type: schema.TypeInt, Optional: true, Default: 5, Description: "smp session scan frequency (scan per second)",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSystemSessionReclaimLimitCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemSessionReclaimLimitCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemSessionReclaimLimit(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemSessionReclaimLimitRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemSessionReclaimLimitUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemSessionReclaimLimitUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemSessionReclaimLimit(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemSessionReclaimLimitRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemSessionReclaimLimitDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemSessionReclaimLimitDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemSessionReclaimLimit(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemSessionReclaimLimitRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemSessionReclaimLimitRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemSessionReclaimLimit(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSystemSessionReclaimLimit(d *schema.ResourceData) edpt.SystemSessionReclaimLimit {
	var ret edpt.SystemSessionReclaimLimit
	ret.Inst.NscanLimit = d.Get("nscan_limit").(int)
	ret.Inst.ScanFreq = d.Get("scan_freq").(int)
	//omit uuid
	return ret
}
