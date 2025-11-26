package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceControllerPartitionTenantInfoOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_controller_partition_tenant_info_oper`: Operational Status for the object partition-tenant-info\n\n__PLACEHOLDER__",
		ReadContext: resourceControllerPartitionTenantInfoOperRead,

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

func resourceControllerPartitionTenantInfoOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceControllerPartitionTenantInfoOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointControllerPartitionTenantInfoOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		ControllerPartitionTenantInfoOperOper := setObjectControllerPartitionTenantInfoOperOper(res)
		d.Set("oper", ControllerPartitionTenantInfoOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectControllerPartitionTenantInfoOperOper(ret edpt.DataControllerPartitionTenantInfoOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"partition_name":   ret.DtControllerPartitionTenantInfoOper.Oper.PartitionName,
			"tenant_name":      ret.DtControllerPartitionTenantInfoOper.Oper.TenantName,
			"tenant_id":        ret.DtControllerPartitionTenantInfoOper.Oper.TenantId,
			"cluster_name":     ret.DtControllerPartitionTenantInfoOper.Oper.ClusterName,
			"cluster_id":       ret.DtControllerPartitionTenantInfoOper.Oper.ClusterId,
			"log_rate_per_sec": ret.DtControllerPartitionTenantInfoOper.Oper.LogRatePerSec,
		},
	}
}

func getObjectControllerPartitionTenantInfoOperOper(d []interface{}) edpt.ControllerPartitionTenantInfoOperOper {

	count1 := len(d)
	var ret edpt.ControllerPartitionTenantInfoOperOper
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

func dataToEndpointControllerPartitionTenantInfoOper(d *schema.ResourceData) edpt.ControllerPartitionTenantInfoOper {
	var ret edpt.ControllerPartitionTenantInfoOper

	ret.Oper = getObjectControllerPartitionTenantInfoOperOper(d.Get("oper").([]interface{}))
	return ret
}
