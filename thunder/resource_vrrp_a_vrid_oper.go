package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVrrpAVridOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_vrrp_a_vrid_oper`: Operational Status for the object vrid\n\n__PLACEHOLDER__",
		ReadContext: resourceVrrpAVridOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"unit": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"state": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"weight": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"priority": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"force_standby": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"init_status": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"became_active": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"vrid_lead": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"active_standby_local": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"peer_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"peer_state": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"peer_unit": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"peer_weight": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"peer_priority": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"active_standby_peer": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"peer_vrid": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
					},
				},
			},
			"vrid_val": {
				Type: schema.TypeInt, Required: true, Description: "Specify ha VRRP-A vrid",
			},
		},
	}
}

func resourceVrrpAVridOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVrrpAVridOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVrrpAVridOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		VrrpAVridOperOper := setObjectVrrpAVridOperOper(res)
		d.Set("oper", VrrpAVridOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectVrrpAVridOperOper(ret edpt.DataVrrpAVridOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"unit":                 ret.DtVrrpAVridOper.Oper.Unit,
			"state":                ret.DtVrrpAVridOper.Oper.State,
			"weight":               ret.DtVrrpAVridOper.Oper.Weight,
			"priority":             ret.DtVrrpAVridOper.Oper.Priority,
			"force_standby":        ret.DtVrrpAVridOper.Oper.Force_standby,
			"init_status":          ret.DtVrrpAVridOper.Oper.Init_status,
			"became_active":        ret.DtVrrpAVridOper.Oper.Became_active,
			"vrid_lead":            ret.DtVrrpAVridOper.Oper.Vrid_lead,
			"active_standby_local": ret.DtVrrpAVridOper.Oper.Active_standby_local,
			"peer_list":            setSliceVrrpAVridOperOperPeerList(ret.DtVrrpAVridOper.Oper.PeerList),
		},
	}
}

func setSliceVrrpAVridOperOperPeerList(d []edpt.VrrpAVridOperOperPeerList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["peer_state"] = item.Peer_state
		in["peer_unit"] = item.Peer_unit
		in["peer_weight"] = item.Peer_weight
		in["peer_priority"] = item.Peer_priority
		in["active_standby_peer"] = item.Active_standby_peer
		in["peer_vrid"] = item.Peer_vrid
		result = append(result, in)
	}
	return result
}

func getObjectVrrpAVridOperOper(d []interface{}) edpt.VrrpAVridOperOper {

	count1 := len(d)
	var ret edpt.VrrpAVridOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Unit = in["unit"].(int)
		ret.State = in["state"].(string)
		ret.Weight = in["weight"].(int)
		ret.Priority = in["priority"].(int)
		ret.Force_standby = in["force_standby"].(string)
		ret.Init_status = in["init_status"].(string)
		ret.Became_active = in["became_active"].(string)
		ret.Vrid_lead = in["vrid_lead"].(string)
		ret.Active_standby_local = in["active_standby_local"].(string)
		ret.PeerList = getSliceVrrpAVridOperOperPeerList(in["peer_list"].([]interface{}))
	}
	return ret
}

func getSliceVrrpAVridOperOperPeerList(d []interface{}) []edpt.VrrpAVridOperOperPeerList {

	count1 := len(d)
	ret := make([]edpt.VrrpAVridOperOperPeerList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VrrpAVridOperOperPeerList
		oi.Peer_state = in["peer_state"].(string)
		oi.Peer_unit = in["peer_unit"].(int)
		oi.Peer_weight = in["peer_weight"].(int)
		oi.Peer_priority = in["peer_priority"].(int)
		oi.Active_standby_peer = in["active_standby_peer"].(string)
		oi.Peer_vrid = in["peer_vrid"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointVrrpAVridOper(d *schema.ResourceData) edpt.VrrpAVridOper {
	var ret edpt.VrrpAVridOper

	ret.Oper = getObjectVrrpAVridOperOper(d.Get("oper").([]interface{}))

	ret.VridVal = d.Get("vrid_val").(int)
	return ret
}
