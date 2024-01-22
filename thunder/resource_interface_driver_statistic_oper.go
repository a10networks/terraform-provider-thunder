package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterfaceDriverStatisticOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_interface_driver_statistic_oper`: Operational Status for the object driver-statistic\n\n__PLACEHOLDER__",
		ReadContext: resourceInterfaceDriverStatisticOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"interfaces": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"if_type": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"port_num": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rxpkts64_counts": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rxpkts65to127_counts": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rxpkts128to255_counts": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rxpkts256to511_counts": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rxpkts512to1023_counts": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rxpkts1024to1518_counts": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rxpkts1519tomax_counts": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"txpkts64_counts": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"txpkts65to127_counts": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"txpkts128to255_counts": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"txpkts256to511_counts": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"txpkts512to1023_counts": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"txpkts1024to1518_counts": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"txpkts1519tomax_counts": {
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

func resourceInterfaceDriverStatisticOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceDriverStatisticOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceDriverStatisticOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		InterfaceDriverStatisticOperOper := setObjectInterfaceDriverStatisticOperOper(res)
		d.Set("oper", InterfaceDriverStatisticOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectInterfaceDriverStatisticOperOper(ret edpt.DataInterfaceDriverStatisticOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"interfaces": setSliceInterfaceDriverStatisticOperOperInterfaces(ret.DtInterfaceDriverStatisticOper.Oper.Interfaces),
		},
	}
}

func setSliceInterfaceDriverStatisticOperOperInterfaces(d []edpt.InterfaceDriverStatisticOperOperInterfaces) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["if_type"] = item.IfType
		in["port_num"] = item.Port_num
		in["rxpkts64_counts"] = item.Rxpkts64_counts
		in["rxpkts65to127_counts"] = item.Rxpkts65to127_counts
		in["rxpkts128to255_counts"] = item.Rxpkts128to255_counts
		in["rxpkts256to511_counts"] = item.Rxpkts256to511_counts
		in["rxpkts512to1023_counts"] = item.Rxpkts512to1023_counts
		in["rxpkts1024to1518_counts"] = item.Rxpkts1024to1518_counts
		in["rxpkts1519tomax_counts"] = item.Rxpkts1519tomax_counts
		in["txpkts64_counts"] = item.Txpkts64_counts
		in["txpkts65to127_counts"] = item.Txpkts65to127_counts
		in["txpkts128to255_counts"] = item.Txpkts128to255_counts
		in["txpkts256to511_counts"] = item.Txpkts256to511_counts
		in["txpkts512to1023_counts"] = item.Txpkts512to1023_counts
		in["txpkts1024to1518_counts"] = item.Txpkts1024to1518_counts
		in["txpkts1519tomax_counts"] = item.Txpkts1519tomax_counts
		result = append(result, in)
	}
	return result
}

func getObjectInterfaceDriverStatisticOperOper(d []interface{}) edpt.InterfaceDriverStatisticOperOper {

	count1 := len(d)
	var ret edpt.InterfaceDriverStatisticOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Interfaces = getSliceInterfaceDriverStatisticOperOperInterfaces(in["interfaces"].([]interface{}))
	}
	return ret
}

func getSliceInterfaceDriverStatisticOperOperInterfaces(d []interface{}) []edpt.InterfaceDriverStatisticOperOperInterfaces {

	count1 := len(d)
	ret := make([]edpt.InterfaceDriverStatisticOperOperInterfaces, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceDriverStatisticOperOperInterfaces
		oi.IfType = in["if_type"].(string)
		oi.Port_num = in["port_num"].(int)
		oi.Rxpkts64_counts = in["rxpkts64_counts"].(int)
		oi.Rxpkts65to127_counts = in["rxpkts65to127_counts"].(int)
		oi.Rxpkts128to255_counts = in["rxpkts128to255_counts"].(int)
		oi.Rxpkts256to511_counts = in["rxpkts256to511_counts"].(int)
		oi.Rxpkts512to1023_counts = in["rxpkts512to1023_counts"].(int)
		oi.Rxpkts1024to1518_counts = in["rxpkts1024to1518_counts"].(int)
		oi.Rxpkts1519tomax_counts = in["rxpkts1519tomax_counts"].(int)
		oi.Txpkts64_counts = in["txpkts64_counts"].(int)
		oi.Txpkts65to127_counts = in["txpkts65to127_counts"].(int)
		oi.Txpkts128to255_counts = in["txpkts128to255_counts"].(int)
		oi.Txpkts256to511_counts = in["txpkts256to511_counts"].(int)
		oi.Txpkts512to1023_counts = in["txpkts512to1023_counts"].(int)
		oi.Txpkts1024to1518_counts = in["txpkts1024to1518_counts"].(int)
		oi.Txpkts1519tomax_counts = in["txpkts1519tomax_counts"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointInterfaceDriverStatisticOper(d *schema.ResourceData) edpt.InterfaceDriverStatisticOper {
	var ret edpt.InterfaceDriverStatisticOper

	ret.Oper = getObjectInterfaceDriverStatisticOperOper(d.Get("oper").([]interface{}))
	return ret
}
