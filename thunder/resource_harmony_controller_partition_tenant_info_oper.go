package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceHarmonyControllerPartitionTenantInfoOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_harmony_controller_partition_tenant_info_oper`: Operational Status for the object partition-tenant-info\n\n__PLACEHOLDER__",
		ReadContext: resourceHarmonyControllerPartitionTenantInfoOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"partition_name": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"tenant_name": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"tenant_id": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"cluster_name": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"cluster_id": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"log_rate_per_sec": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceHarmonyControllerPartitionTenantInfoOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHarmonyControllerPartitionTenantInfoOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHarmonyControllerPartitionTenantInfoOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		HarmonyControllerPartitionTenantInfoOperOper := setObjectHarmonyControllerPartitionTenantInfoOperOper(res)
		d.Set("oper", HarmonyControllerPartitionTenantInfoOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectHarmonyControllerPartitionTenantInfoOperOper(ret edpt.DataHarmonyControllerPartitionTenantInfoOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"partition_name":   ret.DtHarmonyControllerPartitionTenantInfoOper.Oper.PartitionName,
			"tenant_name":      ret.DtHarmonyControllerPartitionTenantInfoOper.Oper.TenantName,
			"tenant_id":        ret.DtHarmonyControllerPartitionTenantInfoOper.Oper.TenantId,
			"cluster_name":     ret.DtHarmonyControllerPartitionTenantInfoOper.Oper.ClusterName,
			"cluster_id":       ret.DtHarmonyControllerPartitionTenantInfoOper.Oper.ClusterId,
			"log_rate_per_sec": ret.DtHarmonyControllerPartitionTenantInfoOper.Oper.LogRatePerSec,
		},
	}
}

func getObjectHarmonyControllerPartitionTenantInfoOperOper(d []interface{}) edpt.HarmonyControllerPartitionTenantInfoOperOper {

	count1 := len(d)
	var ret edpt.HarmonyControllerPartitionTenantInfoOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.PartitionName = in["partition_name"].(string)
		ret.TenantName = in["tenant_name"].(string)
		ret.TenantId = in["tenant_id"].(string)
		ret.ClusterName = in["cluster_name"].(string)
		ret.ClusterId = in["cluster_id"].(string)
		ret.LogRatePerSec = in["log_rate_per_sec"].(int)
	}
	return ret
}

func dataToEndpointHarmonyControllerPartitionTenantInfoOper(d *schema.ResourceData) edpt.HarmonyControllerPartitionTenantInfoOper {
	var ret edpt.HarmonyControllerPartitionTenantInfoOper

	ret.Oper = getObjectHarmonyControllerPartitionTenantInfoOperOper(d.Get("oper").([]interface{}))
	return ret
}
