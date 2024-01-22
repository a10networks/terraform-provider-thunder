package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVpnIkeGatewayOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_vpn_ike_gateway_oper`: Operational Status for the object ike-gateway\n\n__PLACEHOLDER__",
		ReadContext: resourceVpnIkeGatewayOperRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString, Required: true, Description: "IKE-gateway name",
			},
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"remote_ip_filter": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"remote_id_filter": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"brief_filter": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"sa_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"initiator_spi": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"responder_spi": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"local_ip": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"remote_ip": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"encryption": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"hash": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"sign_hash": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"lifetime": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"nat_traversal": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"remote_id": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"dh_group": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"fragment_message_generated": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"fragment_message_received": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"fragmentation_error": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"fragment_reassemble_error": {
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

func resourceVpnIkeGatewayOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVpnIkeGatewayOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVpnIkeGatewayOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		VpnIkeGatewayOperOper := setObjectVpnIkeGatewayOperOper(res)
		d.Set("oper", VpnIkeGatewayOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectVpnIkeGatewayOperOper(ret edpt.DataVpnIkeGatewayOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"remote_ip_filter": ret.DtVpnIkeGatewayOper.Oper.RemoteIpFilter,
			"remote_id_filter": ret.DtVpnIkeGatewayOper.Oper.RemoteIdFilter,
			"brief_filter":     ret.DtVpnIkeGatewayOper.Oper.BriefFilter,
			"sa_list":          setSliceVpnIkeGatewayOperOperSaList(ret.DtVpnIkeGatewayOper.Oper.SaList),
		},
	}
}

func setSliceVpnIkeGatewayOperOperSaList(d []edpt.VpnIkeGatewayOperOperSaList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["initiator_spi"] = item.InitiatorSpi
		in["responder_spi"] = item.ResponderSpi
		in["local_ip"] = item.LocalIp
		in["remote_ip"] = item.RemoteIp
		in["encryption"] = item.Encryption
		in["hash"] = item.Hash
		in["sign_hash"] = item.SignHash
		in["lifetime"] = item.Lifetime
		in["status"] = item.Status
		in["nat_traversal"] = item.NatTraversal
		in["remote_id"] = item.RemoteId
		in["dh_group"] = item.DhGroup
		in["fragment_message_generated"] = item.FragmentMessageGenerated
		in["fragment_message_received"] = item.FragmentMessageReceived
		in["fragmentation_error"] = item.FragmentationError
		in["fragment_reassemble_error"] = item.FragmentReassembleError
		result = append(result, in)
	}
	return result
}

func getObjectVpnIkeGatewayOperOper(d []interface{}) edpt.VpnIkeGatewayOperOper {

	count1 := len(d)
	var ret edpt.VpnIkeGatewayOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RemoteIpFilter = in["remote_ip_filter"].(string)
		ret.RemoteIdFilter = in["remote_id_filter"].(string)
		ret.BriefFilter = in["brief_filter"].(string)
		ret.SaList = getSliceVpnIkeGatewayOperOperSaList(in["sa_list"].([]interface{}))
	}
	return ret
}

func getSliceVpnIkeGatewayOperOperSaList(d []interface{}) []edpt.VpnIkeGatewayOperOperSaList {

	count1 := len(d)
	ret := make([]edpt.VpnIkeGatewayOperOperSaList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VpnIkeGatewayOperOperSaList
		oi.InitiatorSpi = in["initiator_spi"].(string)
		oi.ResponderSpi = in["responder_spi"].(string)
		oi.LocalIp = in["local_ip"].(string)
		oi.RemoteIp = in["remote_ip"].(string)
		oi.Encryption = in["encryption"].(string)
		oi.Hash = in["hash"].(string)
		oi.SignHash = in["sign_hash"].(string)
		oi.Lifetime = in["lifetime"].(int)
		oi.Status = in["status"].(string)
		oi.NatTraversal = in["nat_traversal"].(int)
		oi.RemoteId = in["remote_id"].(string)
		oi.DhGroup = in["dh_group"].(int)
		oi.FragmentMessageGenerated = in["fragment_message_generated"].(int)
		oi.FragmentMessageReceived = in["fragment_message_received"].(int)
		oi.FragmentationError = in["fragmentation_error"].(int)
		oi.FragmentReassembleError = in["fragment_reassemble_error"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointVpnIkeGatewayOper(d *schema.ResourceData) edpt.VpnIkeGatewayOper {
	var ret edpt.VpnIkeGatewayOper

	ret.Name = d.Get("name").(string)

	ret.Oper = getObjectVpnIkeGatewayOperOper(d.Get("oper").([]interface{}))
	return ret
}
