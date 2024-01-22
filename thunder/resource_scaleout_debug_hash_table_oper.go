package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceScaleoutDebugHashTableOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_scaleout_debug_hash_table_oper`: Operational Status for the object hash-table\n\n__PLACEHOLDER__",
		ReadContext: resourceScaleoutDebugHashTableOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"mac": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"hash_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"hash": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"node": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"so_vnp_id": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"so_ip": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"so_mac": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"ref_count": {
										Type: schema.TypeInt, Optional: true, Description: "",
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

func resourceScaleoutDebugHashTableOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutDebugHashTableOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutDebugHashTableOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		ScaleoutDebugHashTableOperOper := setObjectScaleoutDebugHashTableOperOper(res)
		d.Set("oper", ScaleoutDebugHashTableOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectScaleoutDebugHashTableOperOper(ret edpt.DataScaleoutDebugHashTableOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"ip":        ret.DtScaleoutDebugHashTableOper.Oper.Ip,
			"mac":       ret.DtScaleoutDebugHashTableOper.Oper.Mac,
			"hash_list": setSliceScaleoutDebugHashTableOperOperHash_list(ret.DtScaleoutDebugHashTableOper.Oper.Hash_list),
		},
	}
}

func setSliceScaleoutDebugHashTableOperOperHash_list(d []edpt.ScaleoutDebugHashTableOperOperHash_list) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["hash"] = item.Hash
		in["node"] = item.Node
		in["so_vnp_id"] = item.So_vnp_id
		in["so_ip"] = item.So_ip
		in["so_mac"] = item.So_mac
		in["ref_count"] = item.Ref_count
		result = append(result, in)
	}
	return result
}

func getObjectScaleoutDebugHashTableOperOper(d []interface{}) edpt.ScaleoutDebugHashTableOperOper {

	count1 := len(d)
	var ret edpt.ScaleoutDebugHashTableOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ip = in["ip"].(int)
		ret.Mac = in["mac"].(int)
		ret.Hash_list = getSliceScaleoutDebugHashTableOperOperHash_list(in["hash_list"].([]interface{}))
	}
	return ret
}

func getSliceScaleoutDebugHashTableOperOperHash_list(d []interface{}) []edpt.ScaleoutDebugHashTableOperOperHash_list {

	count1 := len(d)
	ret := make([]edpt.ScaleoutDebugHashTableOperOperHash_list, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutDebugHashTableOperOperHash_list
		oi.Hash = in["hash"].(int)
		oi.Node = in["node"].(int)
		oi.So_vnp_id = in["so_vnp_id"].(int)
		oi.So_ip = in["so_ip"].(string)
		oi.So_mac = in["so_mac"].(string)
		oi.Ref_count = in["ref_count"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointScaleoutDebugHashTableOper(d *schema.ResourceData) edpt.ScaleoutDebugHashTableOper {
	var ret edpt.ScaleoutDebugHashTableOper

	ret.Oper = getObjectScaleoutDebugHashTableOperOper(d.Get("oper").([]interface{}))
	return ret
}
