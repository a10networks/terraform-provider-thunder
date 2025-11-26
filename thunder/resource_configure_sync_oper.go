package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceConfigureSyncOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_configure_sync_oper`: Operational Status for the object sync\n\n__PLACEHOLDER__",
		ReadContext: resourceConfigureSyncOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"all_partitions": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"config_sync_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"partition_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"run_sync_status": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"startup_sync_status": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func resourceConfigureSyncOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceConfigureSyncOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointConfigureSyncOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		ConfigureSyncOperOper := setObjectConfigureSyncOperOper(res)
		d.Set("oper", ConfigureSyncOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectConfigureSyncOperOper(ret edpt.DataConfigureSyncOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"all_partitions":   ret.DtConfigureSyncOper.Oper.AllPartitions,
			"config_sync_list": setSliceConfigureSyncOperOperConfigSyncList(ret.DtConfigureSyncOper.Oper.ConfigSyncList),
		},
	}
}

func setSliceConfigureSyncOperOperConfigSyncList(d []edpt.ConfigureSyncOperOperConfigSyncList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["partition_name"] = item.PartitionName
		in["run_sync_status"] = item.RunSyncStatus
		in["startup_sync_status"] = item.StartupSyncStatus
		result = append(result, in)
	}
	return result
}

func getObjectConfigureSyncOperOper(d []interface{}) edpt.ConfigureSyncOperOper {

	count1 := len(d)
	var ret edpt.ConfigureSyncOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AllPartitions = in["all_partitions"].(int)
		ret.ConfigSyncList = getSliceConfigureSyncOperOperConfigSyncList(in["config_sync_list"].([]interface{}))
	}
	return ret
}

func getSliceConfigureSyncOperOperConfigSyncList(d []interface{}) []edpt.ConfigureSyncOperOperConfigSyncList {

	count1 := len(d)
	ret := make([]edpt.ConfigureSyncOperOperConfigSyncList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ConfigureSyncOperOperConfigSyncList
		oi.PartitionName = in["partition_name"].(string)
		oi.RunSyncStatus = in["run_sync_status"].(string)
		oi.StartupSyncStatus = in["startup_sync_status"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointConfigureSyncOper(d *schema.ResourceData) edpt.ConfigureSyncOper {
	var ret edpt.ConfigureSyncOper

	ret.Oper = getObjectConfigureSyncOperOper(d.Get("oper").([]interface{}))
	return ret
}
