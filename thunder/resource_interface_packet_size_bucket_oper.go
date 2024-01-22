package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterfacePacketSizeBucketOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_interface_packet_size_bucket_oper`: Operational Status for the object packet-size-bucket\n\n__PLACEHOLDER__",
		ReadContext: resourceInterfacePacketSizeBucketOperRead,

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
									"rxpkts1519to2047_counts": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rxpkts2048to4095_counts": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rxpkts4096to9216_counts": {
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
									"txpkts1519to2047_counts": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"txpkts2048to4095_counts": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"txpkts4096to9216_counts": {
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

func resourceInterfacePacketSizeBucketOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfacePacketSizeBucketOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfacePacketSizeBucketOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		InterfacePacketSizeBucketOperOper := setObjectInterfacePacketSizeBucketOperOper(res)
		d.Set("oper", InterfacePacketSizeBucketOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectInterfacePacketSizeBucketOperOper(ret edpt.DataInterfacePacketSizeBucketOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"interfaces": setSliceInterfacePacketSizeBucketOperOperInterfaces(ret.DtInterfacePacketSizeBucketOper.Oper.Interfaces),
		},
	}
}

func setSliceInterfacePacketSizeBucketOperOperInterfaces(d []edpt.InterfacePacketSizeBucketOperOperInterfaces) []map[string]interface{} {
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
		in["rxpkts1519to2047_counts"] = item.Rxpkts1519to2047_counts
		in["rxpkts2048to4095_counts"] = item.Rxpkts2048to4095_counts
		in["rxpkts4096to9216_counts"] = item.Rxpkts4096to9216_counts
		in["txpkts64_counts"] = item.Txpkts64_counts
		in["txpkts65to127_counts"] = item.Txpkts65to127_counts
		in["txpkts128to255_counts"] = item.Txpkts128to255_counts
		in["txpkts256to511_counts"] = item.Txpkts256to511_counts
		in["txpkts512to1023_counts"] = item.Txpkts512to1023_counts
		in["txpkts1024to1518_counts"] = item.Txpkts1024to1518_counts
		in["txpkts1519to2047_counts"] = item.Txpkts1519to2047_counts
		in["txpkts2048to4095_counts"] = item.Txpkts2048to4095_counts
		in["txpkts4096to9216_counts"] = item.Txpkts4096to9216_counts
		result = append(result, in)
	}
	return result
}

func getObjectInterfacePacketSizeBucketOperOper(d []interface{}) edpt.InterfacePacketSizeBucketOperOper {

	count1 := len(d)
	var ret edpt.InterfacePacketSizeBucketOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Interfaces = getSliceInterfacePacketSizeBucketOperOperInterfaces(in["interfaces"].([]interface{}))
	}
	return ret
}

func getSliceInterfacePacketSizeBucketOperOperInterfaces(d []interface{}) []edpt.InterfacePacketSizeBucketOperOperInterfaces {

	count1 := len(d)
	ret := make([]edpt.InterfacePacketSizeBucketOperOperInterfaces, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfacePacketSizeBucketOperOperInterfaces
		oi.IfType = in["if_type"].(string)
		oi.Port_num = in["port_num"].(int)
		oi.Rxpkts64_counts = in["rxpkts64_counts"].(int)
		oi.Rxpkts65to127_counts = in["rxpkts65to127_counts"].(int)
		oi.Rxpkts128to255_counts = in["rxpkts128to255_counts"].(int)
		oi.Rxpkts256to511_counts = in["rxpkts256to511_counts"].(int)
		oi.Rxpkts512to1023_counts = in["rxpkts512to1023_counts"].(int)
		oi.Rxpkts1024to1518_counts = in["rxpkts1024to1518_counts"].(int)
		oi.Rxpkts1519to2047_counts = in["rxpkts1519to2047_counts"].(int)
		oi.Rxpkts2048to4095_counts = in["rxpkts2048to4095_counts"].(int)
		oi.Rxpkts4096to9216_counts = in["rxpkts4096to9216_counts"].(int)
		oi.Txpkts64_counts = in["txpkts64_counts"].(int)
		oi.Txpkts65to127_counts = in["txpkts65to127_counts"].(int)
		oi.Txpkts128to255_counts = in["txpkts128to255_counts"].(int)
		oi.Txpkts256to511_counts = in["txpkts256to511_counts"].(int)
		oi.Txpkts512to1023_counts = in["txpkts512to1023_counts"].(int)
		oi.Txpkts1024to1518_counts = in["txpkts1024to1518_counts"].(int)
		oi.Txpkts1519to2047_counts = in["txpkts1519to2047_counts"].(int)
		oi.Txpkts2048to4095_counts = in["txpkts2048to4095_counts"].(int)
		oi.Txpkts4096to9216_counts = in["txpkts4096to9216_counts"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointInterfacePacketSizeBucketOper(d *schema.ResourceData) edpt.InterfacePacketSizeBucketOper {
	var ret edpt.InterfacePacketSizeBucketOper

	ret.Oper = getObjectInterfacePacketSizeBucketOperOper(d.Get("oper").([]interface{}))
	return ret
}
