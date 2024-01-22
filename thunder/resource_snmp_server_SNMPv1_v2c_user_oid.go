package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSnmpServerSnmpv1V2cUserOid() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_snmp_server_SNMPv1_v2c_user_oid`: Define a remote entity to which user belongs\n\n__PLACEHOLDER__",
		CreateContext: resourceSnmpServerSnmpv1V2cUserOidCreate,
		UpdateContext: resourceSnmpServerSnmpv1V2cUserOidUpdate,
		ReadContext:   resourceSnmpServerSnmpv1V2cUserOidRead,
		DeleteContext: resourceSnmpServerSnmpv1V2cUserOidDelete,

		Schema: map[string]*schema.Schema{
			"oid_val": {
				Type: schema.TypeString, Required: true, Description: "specific the oid (The oid value, object-key)",
			},
			"remote": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"host_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"dns_host": {
										Type: schema.TypeString, Optional: true, Description: "DNS remote host",
									},
									"ipv4_mask": {
										Type: schema.TypeString, Optional: true, Description: "IPV4 mask",
									},
								},
							},
						},
						"ipv4_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ipv4_host": {
										Type: schema.TypeString, Optional: true, Description: "IPV4 remote host",
									},
									"ipv4_mask": {
										Type: schema.TypeString, Optional: true, Description: "IPV4 mask",
									},
								},
							},
						},
						"ipv6_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ipv6_host": {
										Type: schema.TypeString, Optional: true, Description: "IPV6 remote host",
									},
									"ipv6_mask": {
										Type: schema.TypeInt, Optional: true, Description: "IPV6 mask",
									},
								},
							},
						},
					},
				},
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"user": {
				Type: schema.TypeString, Required: true, Description: "User",
			},
		},
	}
}
func resourceSnmpServerSnmpv1V2cUserOidCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerSnmpv1V2cUserOidCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerSnmpv1V2cUserOid(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSnmpServerSnmpv1V2cUserOidRead(ctx, d, meta)
	}
	return diags
}

func resourceSnmpServerSnmpv1V2cUserOidUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerSnmpv1V2cUserOidUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerSnmpv1V2cUserOid(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSnmpServerSnmpv1V2cUserOidRead(ctx, d, meta)
	}
	return diags
}
func resourceSnmpServerSnmpv1V2cUserOidDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerSnmpv1V2cUserOidDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerSnmpv1V2cUserOid(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSnmpServerSnmpv1V2cUserOidRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerSnmpv1V2cUserOidRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerSnmpv1V2cUserOid(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectSnmpServerSnmpv1V2cUserOidRemote(d []interface{}) edpt.SnmpServerSnmpv1V2cUserOidRemote {

	count1 := len(d)
	var ret edpt.SnmpServerSnmpv1V2cUserOidRemote
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.HostList = getSliceSnmpServerSnmpv1V2cUserOidRemoteHostList(in["host_list"].([]interface{}))
		ret.Ipv4List = getSliceSnmpServerSnmpv1V2cUserOidRemoteIpv4List(in["ipv4_list"].([]interface{}))
		ret.Ipv6List = getSliceSnmpServerSnmpv1V2cUserOidRemoteIpv6List(in["ipv6_list"].([]interface{}))
	}
	return ret
}

func getSliceSnmpServerSnmpv1V2cUserOidRemoteHostList(d []interface{}) []edpt.SnmpServerSnmpv1V2cUserOidRemoteHostList {

	count1 := len(d)
	ret := make([]edpt.SnmpServerSnmpv1V2cUserOidRemoteHostList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SnmpServerSnmpv1V2cUserOidRemoteHostList
		oi.DnsHost = in["dns_host"].(string)
		oi.Ipv4Mask = in["ipv4_mask"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSnmpServerSnmpv1V2cUserOidRemoteIpv4List(d []interface{}) []edpt.SnmpServerSnmpv1V2cUserOidRemoteIpv4List {

	count1 := len(d)
	ret := make([]edpt.SnmpServerSnmpv1V2cUserOidRemoteIpv4List, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SnmpServerSnmpv1V2cUserOidRemoteIpv4List
		oi.Ipv4Host = in["ipv4_host"].(string)
		oi.Ipv4Mask = in["ipv4_mask"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSnmpServerSnmpv1V2cUserOidRemoteIpv6List(d []interface{}) []edpt.SnmpServerSnmpv1V2cUserOidRemoteIpv6List {

	count1 := len(d)
	ret := make([]edpt.SnmpServerSnmpv1V2cUserOidRemoteIpv6List, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SnmpServerSnmpv1V2cUserOidRemoteIpv6List
		oi.Ipv6Host = in["ipv6_host"].(string)
		oi.Ipv6Mask = in["ipv6_mask"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSnmpServerSnmpv1V2cUserOid(d *schema.ResourceData) edpt.SnmpServerSnmpv1V2cUserOid {
	var ret edpt.SnmpServerSnmpv1V2cUserOid
	ret.Inst.OidVal = d.Get("oid_val").(string)
	ret.Inst.Remote = getObjectSnmpServerSnmpv1V2cUserOidRemote(d.Get("remote").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.User = d.Get("user").(string)
	return ret
}
