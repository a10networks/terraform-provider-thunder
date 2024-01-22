package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSnmpServerCommunityRead() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_snmp_server_community_read`: Define a read only community string\n\n__PLACEHOLDER__",
		CreateContext: resourceSnmpServerCommunityReadCreate,
		UpdateContext: resourceSnmpServerCommunityReadUpdate,
		ReadContext:   resourceSnmpServerCommunityReadRead,
		DeleteContext: resourceSnmpServerCommunityReadDelete,

		Schema: map[string]*schema.Schema{
			"oid_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
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
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
						},
					},
				},
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
			"user": {
				Type: schema.TypeString, Required: true, Description: "SNMPv1/v2c community string",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSnmpServerCommunityReadCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerCommunityReadCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerCommunityRead(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSnmpServerCommunityReadRead(ctx, d, meta)
	}
	return diags
}

func resourceSnmpServerCommunityReadUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerCommunityReadUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerCommunityRead(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSnmpServerCommunityReadRead(ctx, d, meta)
	}
	return diags
}
func resourceSnmpServerCommunityReadDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerCommunityReadDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerCommunityRead(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSnmpServerCommunityReadRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerCommunityReadRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerCommunityRead(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSnmpServerCommunityReadOidList(d []interface{}) []edpt.SnmpServerCommunityReadOidList {

	count1 := len(d)
	ret := make([]edpt.SnmpServerCommunityReadOidList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SnmpServerCommunityReadOidList
		oi.OidVal = in["oid_val"].(string)
		oi.Remote = getObjectSnmpServerCommunityReadOidListRemote(in["remote"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSnmpServerCommunityReadOidListRemote(d []interface{}) edpt.SnmpServerCommunityReadOidListRemote {

	count1 := len(d)
	var ret edpt.SnmpServerCommunityReadOidListRemote
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.HostList = getSliceSnmpServerCommunityReadOidListRemoteHostList(in["host_list"].([]interface{}))
		ret.Ipv4List = getSliceSnmpServerCommunityReadOidListRemoteIpv4List(in["ipv4_list"].([]interface{}))
		ret.Ipv6List = getSliceSnmpServerCommunityReadOidListRemoteIpv6List(in["ipv6_list"].([]interface{}))
	}
	return ret
}

func getSliceSnmpServerCommunityReadOidListRemoteHostList(d []interface{}) []edpt.SnmpServerCommunityReadOidListRemoteHostList {

	count1 := len(d)
	ret := make([]edpt.SnmpServerCommunityReadOidListRemoteHostList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SnmpServerCommunityReadOidListRemoteHostList
		oi.DnsHost = in["dns_host"].(string)
		oi.Ipv4Mask = in["ipv4_mask"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSnmpServerCommunityReadOidListRemoteIpv4List(d []interface{}) []edpt.SnmpServerCommunityReadOidListRemoteIpv4List {

	count1 := len(d)
	ret := make([]edpt.SnmpServerCommunityReadOidListRemoteIpv4List, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SnmpServerCommunityReadOidListRemoteIpv4List
		oi.Ipv4Host = in["ipv4_host"].(string)
		oi.Ipv4Mask = in["ipv4_mask"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSnmpServerCommunityReadOidListRemoteIpv6List(d []interface{}) []edpt.SnmpServerCommunityReadOidListRemoteIpv6List {

	count1 := len(d)
	ret := make([]edpt.SnmpServerCommunityReadOidListRemoteIpv6List, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SnmpServerCommunityReadOidListRemoteIpv6List
		oi.Ipv6Host = in["ipv6_host"].(string)
		oi.Ipv6Mask = in["ipv6_mask"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSnmpServerCommunityReadRemote(d []interface{}) edpt.SnmpServerCommunityReadRemote {

	count1 := len(d)
	var ret edpt.SnmpServerCommunityReadRemote
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.HostList = getSliceSnmpServerCommunityReadRemoteHostList(in["host_list"].([]interface{}))
		ret.Ipv4List = getSliceSnmpServerCommunityReadRemoteIpv4List(in["ipv4_list"].([]interface{}))
		ret.Ipv6List = getSliceSnmpServerCommunityReadRemoteIpv6List(in["ipv6_list"].([]interface{}))
	}
	return ret
}

func getSliceSnmpServerCommunityReadRemoteHostList(d []interface{}) []edpt.SnmpServerCommunityReadRemoteHostList {

	count1 := len(d)
	ret := make([]edpt.SnmpServerCommunityReadRemoteHostList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SnmpServerCommunityReadRemoteHostList
		oi.DnsHost = in["dns_host"].(string)
		oi.Ipv4Mask = in["ipv4_mask"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSnmpServerCommunityReadRemoteIpv4List(d []interface{}) []edpt.SnmpServerCommunityReadRemoteIpv4List {

	count1 := len(d)
	ret := make([]edpt.SnmpServerCommunityReadRemoteIpv4List, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SnmpServerCommunityReadRemoteIpv4List
		oi.Ipv4Host = in["ipv4_host"].(string)
		oi.Ipv4Mask = in["ipv4_mask"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSnmpServerCommunityReadRemoteIpv6List(d []interface{}) []edpt.SnmpServerCommunityReadRemoteIpv6List {

	count1 := len(d)
	ret := make([]edpt.SnmpServerCommunityReadRemoteIpv6List, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SnmpServerCommunityReadRemoteIpv6List
		oi.Ipv6Host = in["ipv6_host"].(string)
		oi.Ipv6Mask = in["ipv6_mask"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSnmpServerCommunityRead(d *schema.ResourceData) edpt.SnmpServerCommunityRead {
	var ret edpt.SnmpServerCommunityRead
	ret.Inst.OidList = getSliceSnmpServerCommunityReadOidList(d.Get("oid_list").([]interface{}))
	ret.Inst.Remote = getObjectSnmpServerCommunityReadRemote(d.Get("remote").([]interface{}))
	ret.Inst.User = d.Get("user").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
