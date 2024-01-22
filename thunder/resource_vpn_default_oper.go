package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVpnDefaultOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_vpn_default_oper`: Operational Status for the object default\n\n__PLACEHOLDER__",
		ReadContext: resourceVpnDefaultOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ike_version": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"ike_mode": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"ike_dh_group": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"ike_auth_method": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"ike_encryption": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"ike_hash": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"ike_priority": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"ike_lifetime": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"ike_nat_traversal": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"ike_local_address": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"ike_remote_address": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"ike_dpd_interval": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"ipsec_mode": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"ipsec_protocol": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"ipsec_dh_group": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"ipsec_encryption": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"ipsec_hash": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"ipsec_priority": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"ipsec_lifetime": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"ipsec_lifebytes": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"ipsec_traffic_selector": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"ipsec_local_subnet": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"ipsec_local_port": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"ipsec_local_protocol": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"ipsec_remote_subnet": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"ipsec_remote_port": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"ipsec_remote_protocol": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"ipsec_anti_replay_window": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceVpnDefaultOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVpnDefaultOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVpnDefaultOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		VpnDefaultOperOper := setObjectVpnDefaultOperOper(res)
		d.Set("oper", VpnDefaultOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectVpnDefaultOperOper(ret edpt.DataVpnDefaultOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"ike_version":              ret.DtVpnDefaultOper.Oper.IkeVersion,
			"ike_mode":                 ret.DtVpnDefaultOper.Oper.IkeMode,
			"ike_dh_group":             ret.DtVpnDefaultOper.Oper.IkeDhGroup,
			"ike_auth_method":          ret.DtVpnDefaultOper.Oper.IkeAuthMethod,
			"ike_encryption":           ret.DtVpnDefaultOper.Oper.IkeEncryption,
			"ike_hash":                 ret.DtVpnDefaultOper.Oper.IkeHash,
			"ike_priority":             ret.DtVpnDefaultOper.Oper.IkePriority,
			"ike_lifetime":             ret.DtVpnDefaultOper.Oper.IkeLifetime,
			"ike_nat_traversal":        ret.DtVpnDefaultOper.Oper.IkeNatTraversal,
			"ike_local_address":        ret.DtVpnDefaultOper.Oper.IkeLocalAddress,
			"ike_remote_address":       ret.DtVpnDefaultOper.Oper.IkeRemoteAddress,
			"ike_dpd_interval":         ret.DtVpnDefaultOper.Oper.IkeDpdInterval,
			"ipsec_mode":               ret.DtVpnDefaultOper.Oper.IpsecMode,
			"ipsec_protocol":           ret.DtVpnDefaultOper.Oper.IpsecProtocol,
			"ipsec_dh_group":           ret.DtVpnDefaultOper.Oper.IpsecDhGroup,
			"ipsec_encryption":         ret.DtVpnDefaultOper.Oper.IpsecEncryption,
			"ipsec_hash":               ret.DtVpnDefaultOper.Oper.IpsecHash,
			"ipsec_priority":           ret.DtVpnDefaultOper.Oper.IpsecPriority,
			"ipsec_lifetime":           ret.DtVpnDefaultOper.Oper.IpsecLifetime,
			"ipsec_lifebytes":          ret.DtVpnDefaultOper.Oper.IpsecLifebytes,
			"ipsec_traffic_selector":   ret.DtVpnDefaultOper.Oper.IpsecTrafficSelector,
			"ipsec_local_subnet":       ret.DtVpnDefaultOper.Oper.IpsecLocalSubnet,
			"ipsec_local_port":         ret.DtVpnDefaultOper.Oper.IpsecLocalPort,
			"ipsec_local_protocol":     ret.DtVpnDefaultOper.Oper.IpsecLocalProtocol,
			"ipsec_remote_subnet":      ret.DtVpnDefaultOper.Oper.IpsecRemoteSubnet,
			"ipsec_remote_port":        ret.DtVpnDefaultOper.Oper.IpsecRemotePort,
			"ipsec_remote_protocol":    ret.DtVpnDefaultOper.Oper.IpsecRemoteProtocol,
			"ipsec_anti_replay_window": ret.DtVpnDefaultOper.Oper.IpsecAntiReplayWindow,
		},
	}
}

func getObjectVpnDefaultOperOper(d []interface{}) edpt.VpnDefaultOperOper {

	count1 := len(d)
	var ret edpt.VpnDefaultOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.IkeVersion = in["ike_version"].(string)
		ret.IkeMode = in["ike_mode"].(string)
		ret.IkeDhGroup = in["ike_dh_group"].(string)
		ret.IkeAuthMethod = in["ike_auth_method"].(string)
		ret.IkeEncryption = in["ike_encryption"].(string)
		ret.IkeHash = in["ike_hash"].(string)
		ret.IkePriority = in["ike_priority"].(int)
		ret.IkeLifetime = in["ike_lifetime"].(int)
		ret.IkeNatTraversal = in["ike_nat_traversal"].(string)
		ret.IkeLocalAddress = in["ike_local_address"].(string)
		ret.IkeRemoteAddress = in["ike_remote_address"].(string)
		ret.IkeDpdInterval = in["ike_dpd_interval"].(int)
		ret.IpsecMode = in["ipsec_mode"].(string)
		ret.IpsecProtocol = in["ipsec_protocol"].(string)
		ret.IpsecDhGroup = in["ipsec_dh_group"].(string)
		ret.IpsecEncryption = in["ipsec_encryption"].(string)
		ret.IpsecHash = in["ipsec_hash"].(string)
		ret.IpsecPriority = in["ipsec_priority"].(int)
		ret.IpsecLifetime = in["ipsec_lifetime"].(int)
		ret.IpsecLifebytes = in["ipsec_lifebytes"].(int)
		ret.IpsecTrafficSelector = in["ipsec_traffic_selector"].(string)
		ret.IpsecLocalSubnet = in["ipsec_local_subnet"].(string)
		ret.IpsecLocalPort = in["ipsec_local_port"].(int)
		ret.IpsecLocalProtocol = in["ipsec_local_protocol"].(int)
		ret.IpsecRemoteSubnet = in["ipsec_remote_subnet"].(string)
		ret.IpsecRemotePort = in["ipsec_remote_port"].(int)
		ret.IpsecRemoteProtocol = in["ipsec_remote_protocol"].(int)
		ret.IpsecAntiReplayWindow = in["ipsec_anti_replay_window"].(int)
	}
	return ret
}

func dataToEndpointVpnDefaultOper(d *schema.ResourceData) edpt.VpnDefaultOper {
	var ret edpt.VpnDefaultOper

	ret.Oper = getObjectVpnDefaultOperOper(d.Get("oper").([]interface{}))
	return ret
}
