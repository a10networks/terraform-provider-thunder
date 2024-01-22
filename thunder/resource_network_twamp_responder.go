package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetworkTwampResponder() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_network_twamp_responder`: Configure TWAMP responder\n\n__PLACEHOLDER__",
		CreateContext: resourceNetworkTwampResponderCreate,
		UpdateContext: resourceNetworkTwampResponderUpdate,
		ReadContext:   resourceNetworkTwampResponderRead,
		DeleteContext: resourceNetworkTwampResponderDelete,

		Schema: map[string]*schema.Schema{
			"enable_both_ip_ipv6": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable both IP and IPv6 TWAMP packets",
			},
			"enable_ip": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable IP TWAMP packets",
			},
			"enable_ipv6": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable IPv6 TWAMP packets",
			},
			"ip": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"acl_id": {
							Type: schema.TypeInt, Optional: true, Description: "ACL id",
						},
						"acl_name": {
							Type: schema.TypeString, Optional: true, Description: "Apply a named access list (Access List name)",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"ipv6": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"v6_acl_name": {
							Type: schema.TypeString, Optional: true, Description: "Apply an access list (Named Access List)",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"port": {
				Type: schema.TypeInt, Optional: true, Description: "Port number",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceNetworkTwampResponderCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkTwampResponderCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkTwampResponder(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceNetworkTwampResponderRead(ctx, d, meta)
	}
	return diags
}

func resourceNetworkTwampResponderUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkTwampResponderUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkTwampResponder(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceNetworkTwampResponderRead(ctx, d, meta)
	}
	return diags
}
func resourceNetworkTwampResponderDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkTwampResponderDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkTwampResponder(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceNetworkTwampResponderRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkTwampResponderRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkTwampResponder(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectNetworkTwampResponderIp1074(d []interface{}) edpt.NetworkTwampResponderIp1074 {

	count1 := len(d)
	var ret edpt.NetworkTwampResponderIp1074
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AclId = in["acl_id"].(int)
		ret.AclName = in["acl_name"].(string)
		//omit uuid
	}
	return ret
}

func getObjectNetworkTwampResponderIpv61075(d []interface{}) edpt.NetworkTwampResponderIpv61075 {

	count1 := len(d)
	var ret edpt.NetworkTwampResponderIpv61075
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.V6AclName = in["v6_acl_name"].(string)
		//omit uuid
	}
	return ret
}

func dataToEndpointNetworkTwampResponder(d *schema.ResourceData) edpt.NetworkTwampResponder {
	var ret edpt.NetworkTwampResponder
	ret.Inst.EnableBothIpIpv6 = d.Get("enable_both_ip_ipv6").(int)
	ret.Inst.EnableIp = d.Get("enable_ip").(int)
	ret.Inst.EnableIpv6 = d.Get("enable_ipv6").(int)
	ret.Inst.Ip = getObjectNetworkTwampResponderIp1074(d.Get("ip").([]interface{}))
	ret.Inst.Ipv6 = getObjectNetworkTwampResponderIpv61075(d.Get("ipv6").([]interface{}))
	ret.Inst.Port = d.Get("port").(int)
	//omit uuid
	return ret
}
