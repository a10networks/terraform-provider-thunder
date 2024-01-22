package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceScaleoutDebugClusterDatabaseOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_scaleout_debug_cluster_database_oper`: Operational Status for the object database\n\n__PLACEHOLDER__",
		ReadContext: resourceScaleoutDebugClusterDatabaseOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"database_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"root": {
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

func resourceScaleoutDebugClusterDatabaseOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutDebugClusterDatabaseOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutDebugClusterDatabaseOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		ScaleoutDebugClusterDatabaseOperOper := setObjectScaleoutDebugClusterDatabaseOperOper(res)
		d.Set("oper", ScaleoutDebugClusterDatabaseOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectScaleoutDebugClusterDatabaseOperOper(ret edpt.DataScaleoutDebugClusterDatabaseOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"database_list": setSliceScaleoutDebugClusterDatabaseOperOperDatabaseList(ret.DtScaleoutDebugClusterDatabaseOper.Oper.DatabaseList),
		},
	}
}

func setSliceScaleoutDebugClusterDatabaseOperOperDatabaseList(d []edpt.ScaleoutDebugClusterDatabaseOperOperDatabaseList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["root"] = item.Root
		result = append(result, in)
	}
	return result
}

func getObjectScaleoutDebugClusterDatabaseOperOper(d []interface{}) edpt.ScaleoutDebugClusterDatabaseOperOper {

	count1 := len(d)
	var ret edpt.ScaleoutDebugClusterDatabaseOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DatabaseList = getSliceScaleoutDebugClusterDatabaseOperOperDatabaseList(in["database_list"].([]interface{}))
	}
	return ret
}

func getSliceScaleoutDebugClusterDatabaseOperOperDatabaseList(d []interface{}) []edpt.ScaleoutDebugClusterDatabaseOperOperDatabaseList {

	count1 := len(d)
	ret := make([]edpt.ScaleoutDebugClusterDatabaseOperOperDatabaseList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutDebugClusterDatabaseOperOperDatabaseList
		oi.Root = in["root"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointScaleoutDebugClusterDatabaseOper(d *schema.ResourceData) edpt.ScaleoutDebugClusterDatabaseOper {
	var ret edpt.ScaleoutDebugClusterDatabaseOper

	ret.Oper = getObjectScaleoutDebugClusterDatabaseOperOper(d.Get("oper").([]interface{}))
	return ret
}
