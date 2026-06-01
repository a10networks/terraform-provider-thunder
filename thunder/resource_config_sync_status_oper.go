package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceConfigSyncStatusOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_config_sync_status_oper`: Operational Status for the object config-sync-status\n\n__PLACEHOLDER__",
		ReadContext: resourceConfigSyncStatusOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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

func resourceConfigSyncStatusOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceConfigSyncStatusOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointConfigSyncStatusOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		ConfigSyncStatusOperOper := setObjectConfigSyncStatusOperOper(res)
		d.Set("oper", ConfigSyncStatusOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectConfigSyncStatusOperOper(ret edpt.DataConfigSyncStatusOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"config_sync_list": setSliceConfigSyncStatusOperOperConfigSyncList(ret.DtConfigSyncStatusOper.Oper.ConfigSyncList),
		},
	}
}

func setSliceConfigSyncStatusOperOperConfigSyncList(d []edpt.ConfigSyncStatusOperOperConfigSyncList) []map[string]interface{} {
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

func getObjectConfigSyncStatusOperOper(d []interface{}) edpt.ConfigSyncStatusOperOper {

	count1 := len(d)
	var ret edpt.ConfigSyncStatusOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ConfigSyncList = getSliceConfigSyncStatusOperOperConfigSyncList(in["config_sync_list"].([]interface{}))
	}
	return ret
}

func getSliceConfigSyncStatusOperOperConfigSyncList(d []interface{}) []edpt.ConfigSyncStatusOperOperConfigSyncList {

	count1 := len(d)
	ret := make([]edpt.ConfigSyncStatusOperOperConfigSyncList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ConfigSyncStatusOperOperConfigSyncList
		oi.PartitionName = in["partition_name"].(string)
		oi.RunSyncStatus = in["run_sync_status"].(string)
		oi.StartupSyncStatus = in["startup_sync_status"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointConfigSyncStatusOper(d *schema.ResourceData) edpt.ConfigSyncStatusOper {
	var ret edpt.ConfigSyncStatusOper

	ret.Oper = getObjectConfigSyncStatusOperOper(d.Get("oper").([]interface{}))
	return ret
}
