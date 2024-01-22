package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbPlayerIdEpOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_player_id_ep_oper`: Operational Status for the object player-id-ep\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbPlayerIdEpOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"player_id_ep_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"player_id": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"game_server_address": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"game_server_port": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"age": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"user_session_count": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"idle_time": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
						"total_players": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"filter_type": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"player_id": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceSlbPlayerIdEpOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbPlayerIdEpOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbPlayerIdEpOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbPlayerIdEpOperOper := setObjectSlbPlayerIdEpOperOper(res)
		d.Set("oper", SlbPlayerIdEpOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbPlayerIdEpOperOper(ret edpt.DataSlbPlayerIdEpOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"player_id_ep_list": setSliceSlbPlayerIdEpOperOperPlayerIdEpList(ret.DtSlbPlayerIdEpOper.Oper.PlayerIdEpList),
			"total_players":     ret.DtSlbPlayerIdEpOper.Oper.TotalPlayers,
			"filter_type":       ret.DtSlbPlayerIdEpOper.Oper.Filter_type,
			"player_id":         ret.DtSlbPlayerIdEpOper.Oper.PlayerId,
		},
	}
}

func setSliceSlbPlayerIdEpOperOperPlayerIdEpList(d []edpt.SlbPlayerIdEpOperOperPlayerIdEpList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["player_id"] = item.PlayerId
		in["game_server_address"] = item.GameServerAddress
		in["game_server_port"] = item.GameServerPort
		in["age"] = item.Age
		in["user_session_count"] = item.UserSessionCount
		in["idle_time"] = item.IdleTime
		result = append(result, in)
	}
	return result
}

func getObjectSlbPlayerIdEpOperOper(d []interface{}) edpt.SlbPlayerIdEpOperOper {

	count1 := len(d)
	var ret edpt.SlbPlayerIdEpOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.PlayerIdEpList = getSliceSlbPlayerIdEpOperOperPlayerIdEpList(in["player_id_ep_list"].([]interface{}))
		ret.TotalPlayers = in["total_players"].(int)
		ret.Filter_type = in["filter_type"].(string)
		ret.PlayerId = in["player_id"].(int)
	}
	return ret
}

func getSliceSlbPlayerIdEpOperOperPlayerIdEpList(d []interface{}) []edpt.SlbPlayerIdEpOperOperPlayerIdEpList {

	count1 := len(d)
	ret := make([]edpt.SlbPlayerIdEpOperOperPlayerIdEpList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbPlayerIdEpOperOperPlayerIdEpList
		oi.PlayerId = in["player_id"].(int)
		oi.GameServerAddress = in["game_server_address"].(string)
		oi.GameServerPort = in["game_server_port"].(int)
		oi.Age = in["age"].(int)
		oi.UserSessionCount = in["user_session_count"].(int)
		oi.IdleTime = in["idle_time"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbPlayerIdEpOper(d *schema.ResourceData) edpt.SlbPlayerIdEpOper {
	var ret edpt.SlbPlayerIdEpOper

	ret.Oper = getObjectSlbPlayerIdEpOperOper(d.Get("oper").([]interface{}))
	return ret
}
