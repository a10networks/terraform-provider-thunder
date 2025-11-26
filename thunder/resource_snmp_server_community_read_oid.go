package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSnmpServerCommunityReadOid() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_snmp_server_community_read_oid`: Define a remote entity to which user belongs\n\n__PLACEHOLDER__",
		CreateContext: resourceSnmpServerCommunityReadOidCreate,
		UpdateContext: resourceSnmpServerCommunityReadOidUpdate,
		ReadContext:   resourceSnmpServerCommunityReadOidRead,
		DeleteContext: resourceSnmpServerCommunityReadOidDelete,

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
func resourceSnmpServerCommunityReadOidCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerCommunityReadOidCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerCommunityReadOid(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSnmpServerCommunityReadOidRead(ctx, d, meta)
	}
	return diags
}

func resourceSnmpServerCommunityReadOidUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerCommunityReadOidUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerCommunityReadOid(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSnmpServerCommunityReadOidRead(ctx, d, meta)
	}
	return diags
}
func resourceSnmpServerCommunityReadOidDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerCommunityReadOidDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerCommunityReadOid(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSnmpServerCommunityReadOidRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerCommunityReadOidRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerCommunityReadOid(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectSnmpServerCommunityReadOidRemote(d []interface{}) edpt.SnmpServerCommunityReadOidRemote {

	count1 := len(d)
	var ret edpt.SnmpServerCommunityReadOidRemote
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.HostList = getSliceSnmpServerCommunityReadOidRemoteHostList(in["host_list"].([]interface{}))
		ret.Ipv4List = getSliceSnmpServerCommunityReadOidRemoteIpv4List(in["ipv4_list"].([]interface{}))
		ret.Ipv6List = getSliceSnmpServerCommunityReadOidRemoteIpv6List(in["ipv6_list"].([]interface{}))
	}
	return ret
}

func getSliceSnmpServerCommunityReadOidRemoteHostList(d []interface{}) []edpt.SnmpServerCommunityReadOidRemoteHostList {

	count1 := len(d)
	ret := make([]edpt.SnmpServerCommunityReadOidRemoteHostList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SnmpServerCommunityReadOidRemoteHostList
		oi.DnsHost = in["dns_host"].(string)
		oi.Ipv4Mask = in["ipv4_mask"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSnmpServerCommunityReadOidRemoteIpv4List(d []interface{}) []edpt.SnmpServerCommunityReadOidRemoteIpv4List {

	count1 := len(d)
	ret := make([]edpt.SnmpServerCommunityReadOidRemoteIpv4List, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SnmpServerCommunityReadOidRemoteIpv4List
		oi.Ipv4Host = in["ipv4_host"].(string)
		oi.Ipv4Mask = in["ipv4_mask"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSnmpServerCommunityReadOidRemoteIpv6List(d []interface{}) []edpt.SnmpServerCommunityReadOidRemoteIpv6List {

	count1 := len(d)
	ret := make([]edpt.SnmpServerCommunityReadOidRemoteIpv6List, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SnmpServerCommunityReadOidRemoteIpv6List
		oi.Ipv6Host = in["ipv6_host"].(string)
		oi.Ipv6Mask = in["ipv6_mask"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSnmpServerCommunityReadOid(d *schema.ResourceData) edpt.SnmpServerCommunityReadOid {
	var ret edpt.SnmpServerCommunityReadOid
	ret.Inst.OidVal = d.Get("oid_val").(string)
	ret.Inst.Remote = getObjectSnmpServerCommunityReadOidRemote(d.Get("remote").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.User = d.Get("user").(string)
	return ret
}
