package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVrrpAHostidOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_vrrp_a_hostid_oper`: Operational Status for the object hostid\n\n__PLACEHOLDER__",
		ReadContext: resourceVrrpAHostidOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"set_id": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"hostid_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"device_id": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"sn": {
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

func resourceVrrpAHostidOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVrrpAHostidOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVrrpAHostidOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		VrrpAHostidOperOper := setObjectVrrpAHostidOperOper(res)
		d.Set("oper", VrrpAHostidOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectVrrpAHostidOperOper(ret edpt.DataVrrpAHostidOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"set_id":      ret.DtVrrpAHostidOper.Oper.Set_id,
			"hostid_list": setSliceVrrpAHostidOperOperHostidList(ret.DtVrrpAHostidOper.Oper.HostidList),
		},
	}
}

func setSliceVrrpAHostidOperOperHostidList(d []edpt.VrrpAHostidOperOperHostidList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["device_id"] = item.Device_id
		in["sn"] = item.Sn
		result = append(result, in)
	}
	return result
}

func getObjectVrrpAHostidOperOper(d []interface{}) edpt.VrrpAHostidOperOper {

	count1 := len(d)
	var ret edpt.VrrpAHostidOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Set_id = in["set_id"].(int)
		ret.HostidList = getSliceVrrpAHostidOperOperHostidList(in["hostid_list"].([]interface{}))
	}
	return ret
}

func getSliceVrrpAHostidOperOperHostidList(d []interface{}) []edpt.VrrpAHostidOperOperHostidList {

	count1 := len(d)
	ret := make([]edpt.VrrpAHostidOperOperHostidList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VrrpAHostidOperOperHostidList
		oi.Device_id = in["device_id"].(int)
		oi.Sn = in["sn"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointVrrpAHostidOper(d *schema.ResourceData) edpt.VrrpAHostidOper {
	var ret edpt.VrrpAHostidOper

	ret.Oper = getObjectVrrpAHostidOperOper(d.Get("oper").([]interface{}))
	return ret
}
