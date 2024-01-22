package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemPbslb() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_pbslb`: Configure Policy Based Response-Rate-Limiting\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemPbslbCreate,
		UpdateContext: resourceSystemPbslbUpdate,
		ReadContext:   resourceSystemPbslbRead,
		DeleteContext: resourceSystemPbslbDelete,

		Schema: map[string]*schema.Schema{
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'curr_entries': Current PBSLB Entry Count; 'total_v4_entries_created': Total V4 Entry Count Created; 'total_v4_entries_freed': Total V4 Entry Count Freed; 'total_v6_entries_created': Total V6 Entry Count Created; 'total_v6_entries_freed': Total V6 Entry Count Freed; 'total_domain_entries_created': Total Domain Entry Count Created; 'total_domain_entries_freed': Total Domain Entry Count Freed; 'total_direct_action_entries_created': Total Direct Action Entry Count Created; 'total_direct_action_entries_freed': Total Direct Action Entry Count Freed; 'curr_entries_target_global': Current Entry Target Global; 'curr_entries_target_vserver': Current Entry Target Vserver; 'curr_entries_target_vport': Current Entry Target Vport; 'curr_entries_target_LOC': Current Entry Target LOC; 'curr_entries_target_rserver': Current Entry Target Rserver; 'curr_entries_target_rport': Current Entry Target Rport; 'curr_entries_target_service': Current Entry Target Service; 'curr_entries_stats': Current Entry Stats Count;",
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
func resourceSystemPbslbCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemPbslbCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemPbslb(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemPbslbRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemPbslbUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemPbslbUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemPbslb(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemPbslbRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemPbslbDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemPbslbDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemPbslb(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemPbslbRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemPbslbRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemPbslb(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSystemPbslbSamplingEnable(d []interface{}) []edpt.SystemPbslbSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.SystemPbslbSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemPbslbSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSystemPbslb(d *schema.ResourceData) edpt.SystemPbslb {
	var ret edpt.SystemPbslb
	ret.Inst.SamplingEnable = getSliceSystemPbslbSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
