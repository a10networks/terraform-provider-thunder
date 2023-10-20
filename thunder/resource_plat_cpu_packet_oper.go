package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourcePlatCpuPacketOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_plat_cpu_packet_oper`: Operational Status for the object plat-cpu-packet\n\n__PLACEHOLDER__",
		ReadContext: resourcePlatCpuPacketOperRead,
		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"pkt_stats": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"cpu_num": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"pkt_sent": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"pkt_rcvd": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"pkt_drop": {
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

func resourcePlatCpuPacketOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourcePlatCpuPacketOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointPlatCpuPacketOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		items := setObjectPlatCpuPacketOperOper(res)
		d.SetId(obj.GetId())
		d.Set("oper", items)

		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectPlatCpuPacketOperOper(res edpt.PlatCpuPackett) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["pkt_stats"] = setSlicePlatCpuPacketOperOperPktStats(res.DataPlatCpuPacket.Oper.PktStats)
	result = append(result, in)
	return result
}

func setSlicePlatCpuPacketOperOperPktStats(d []edpt.PlatCpuPacketOperOperPktStats) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["cpu_num"] = item.CpuNum
		in["pkt_sent"] = item.PktSent
		in["pkt_rcvd"] = item.PktRcvd
		in["pkt_drop"] = item.PktDrop
		result = append(result, in)
	}
	return result
}

func getObjectPlatCpuPacketOperOper(d []interface{}) edpt.PlatCpuPacketOperOper {
	count := len(d)
	var ret edpt.PlatCpuPacketOperOper
	if count > 0 {
		in := d[0].(map[string]interface{})
		ret.PktStats = getSlicePlatCpuPacketOperOperPktStats(in["pkt_stats"].([]interface{}))
	}
	return ret
}

func getSlicePlatCpuPacketOperOperPktStats(d []interface{}) []edpt.PlatCpuPacketOperOperPktStats {
	count := len(d)
	ret := make([]edpt.PlatCpuPacketOperOperPktStats, 0, count)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.PlatCpuPacketOperOperPktStats
		oi.CpuNum = in["cpu_num"].(int)
		oi.PktSent = in["pkt_sent"].(int)
		oi.PktRcvd = in["pkt_rcvd"].(int)
		oi.PktDrop = in["pkt_drop"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointPlatCpuPacketOper(d *schema.ResourceData) edpt.PlatCpuPacketOper {
	var ret edpt.PlatCpuPacketOper
	ret.Oper = getObjectPlatCpuPacketOperOper(d.Get("oper").([]interface{}))
	return ret
}
