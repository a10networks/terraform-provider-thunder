package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbPlayerIdList() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_player_id_list`: Player id records config\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbPlayerIdListCreate,
		UpdateContext: resourceSlbPlayerIdListUpdate,
		ReadContext:   resourceSlbPlayerIdListRead,
		DeleteContext: resourceSlbPlayerIdListDelete,

		Schema: map[string]*schema.Schema{
			"player_record": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"player_id": {
							Type: schema.TypeInt, Optional: true, Description: "64/32 bit player id based on config",
						},
						"game_server_ipv4": {
							Type: schema.TypeString, Optional: true, Description: "Specify IP address",
						},
						"game_server_ipv6": {
							Type: schema.TypeString, Optional: true, Description: "Specify IPv6 address",
						},
						"game_server_port_v4": {
							Type: schema.TypeInt, Optional: true, Description: "Port",
						},
						"game_server_port_v6": {
							Type: schema.TypeInt, Optional: true, Description: "Port",
						},
					},
				},
			},
		},
	}
}
func resourceSlbPlayerIdListCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbPlayerIdListCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbPlayerIdList(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbPlayerIdListRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbPlayerIdListUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbPlayerIdListUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbPlayerIdList(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbPlayerIdListRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbPlayerIdListDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbPlayerIdListDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbPlayerIdList(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbPlayerIdListRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbPlayerIdListRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbPlayerIdList(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSlbPlayerIdListPlayerRecord(d []interface{}) []edpt.SlbPlayerIdListPlayerRecord {

	count1 := len(d)
	ret := make([]edpt.SlbPlayerIdListPlayerRecord, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbPlayerIdListPlayerRecord
		oi.PlayerId = in["player_id"].(int)
		oi.GameServerIpv4 = in["game_server_ipv4"].(string)
		oi.GameServerIpv6 = in["game_server_ipv6"].(string)
		oi.GameServerPortV4 = in["game_server_port_v4"].(int)
		oi.GameServerPortV6 = in["game_server_port_v6"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbPlayerIdList(d *schema.ResourceData) edpt.SlbPlayerIdList {
	var ret edpt.SlbPlayerIdList
	ret.Inst.PlayerRecord = getSliceSlbPlayerIdListPlayerRecord(d.Get("player_record").([]interface{}))
	return ret
}
