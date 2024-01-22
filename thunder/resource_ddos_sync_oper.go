package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosSyncOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_sync_oper`: Operational Status for the object sync\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosSyncOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"if_show_error_num": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"error_str": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"status": {
							Type: schema.TypeString, Optional: true, Description: "DDOS sync status",
						},
						"local_ip": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"total_message_received": {
							Type: schema.TypeInt, Optional: true, Description: "Total number of messages received on this device",
						},
						"total_message_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Total number of messages sent from this device",
						},
						"peer_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"peer_ip": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"peer_status": {
										Type: schema.TypeString, Optional: true, Description: "Status of the connection with the peer",
									},
									"peer_heartbeat_missing": {
										Type: schema.TypeInt, Optional: true, Description: "Number of heartbeat messages expected from the peer that did not arrive",
									},
									"peer_message_received": {
										Type: schema.TypeInt, Optional: true, Description: "Number of messages received from this peer",
									},
									"peer_message_sent": {
										Type: schema.TypeInt, Optional: true, Description: "Number messages sent to this peer",
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

func resourceDdosSyncOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosSyncOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosSyncOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosSyncOperOper := setObjectDdosSyncOperOper(res)
		d.Set("oper", DdosSyncOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosSyncOperOper(ret edpt.DataDdosSyncOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"if_show_error_num":      ret.DtDdosSyncOper.Oper.IfShowErrorNum,
			"error_str":              ret.DtDdosSyncOper.Oper.ErrorStr,
			"status":                 ret.DtDdosSyncOper.Oper.Status,
			"local_ip":               ret.DtDdosSyncOper.Oper.LocalIp,
			"total_message_received": ret.DtDdosSyncOper.Oper.TotalMessageReceived,
			"total_message_sent":     ret.DtDdosSyncOper.Oper.TotalMessageSent,
			"peer_list":              setSliceDdosSyncOperOperPeerList(ret.DtDdosSyncOper.Oper.PeerList),
		},
	}
}

func setSliceDdosSyncOperOperPeerList(d []edpt.DdosSyncOperOperPeerList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["peer_ip"] = item.PeerIp
		in["peer_status"] = item.PeerStatus
		in["peer_heartbeat_missing"] = item.PeerHeartbeatMissing
		in["peer_message_received"] = item.PeerMessageReceived
		in["peer_message_sent"] = item.PeerMessageSent
		result = append(result, in)
	}
	return result
}

func getObjectDdosSyncOperOper(d []interface{}) edpt.DdosSyncOperOper {

	count1 := len(d)
	var ret edpt.DdosSyncOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.IfShowErrorNum = in["if_show_error_num"].(int)
		ret.ErrorStr = in["error_str"].(string)
		ret.Status = in["status"].(string)
		ret.LocalIp = in["local_ip"].(string)
		ret.TotalMessageReceived = in["total_message_received"].(int)
		ret.TotalMessageSent = in["total_message_sent"].(int)
		ret.PeerList = getSliceDdosSyncOperOperPeerList(in["peer_list"].([]interface{}))
	}
	return ret
}

func getSliceDdosSyncOperOperPeerList(d []interface{}) []edpt.DdosSyncOperOperPeerList {

	count1 := len(d)
	ret := make([]edpt.DdosSyncOperOperPeerList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosSyncOperOperPeerList
		oi.PeerIp = in["peer_ip"].(string)
		oi.PeerStatus = in["peer_status"].(string)
		oi.PeerHeartbeatMissing = in["peer_heartbeat_missing"].(int)
		oi.PeerMessageReceived = in["peer_message_received"].(int)
		oi.PeerMessageSent = in["peer_message_sent"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDdosSyncOper(d *schema.ResourceData) edpt.DdosSyncOper {
	var ret edpt.DdosSyncOper

	ret.Oper = getObjectDdosSyncOperOper(d.Get("oper").([]interface{}))
	return ret
}
