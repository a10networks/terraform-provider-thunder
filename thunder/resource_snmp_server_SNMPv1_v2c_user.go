package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSnmpServerSnmpv1V2cUser() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_snmp_server_SNMPv1_v2c_user`: Define configuration for SNMPv1/v2c user\n\n__PLACEHOLDER__",
		CreateContext: resourceSnmpServerSnmpv1V2cUserCreate,
		UpdateContext: resourceSnmpServerSnmpv1V2cUserUpdate,
		ReadContext:   resourceSnmpServerSnmpv1V2cUserRead,
		DeleteContext: resourceSnmpServerSnmpv1V2cUserDelete,

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
			"passwd": {
				Type: schema.TypeString, Optional: true, Description: "define value of community string",
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
				Type: schema.TypeString, Required: true, Description: "SNMPv1/v2c community configuration key",
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
func resourceSnmpServerSnmpv1V2cUserCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerSnmpv1V2cUserCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerSnmpv1V2cUser(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSnmpServerSnmpv1V2cUserRead(ctx, d, meta)
	}
	return diags
}

func resourceSnmpServerSnmpv1V2cUserUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerSnmpv1V2cUserUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerSnmpv1V2cUser(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSnmpServerSnmpv1V2cUserRead(ctx, d, meta)
	}
	return diags
}
func resourceSnmpServerSnmpv1V2cUserDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerSnmpv1V2cUserDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerSnmpv1V2cUser(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSnmpServerSnmpv1V2cUserRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerSnmpv1V2cUserRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerSnmpv1V2cUser(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSnmpServerSnmpv1V2cUserOidList(d []interface{}) []edpt.SnmpServerSnmpv1V2cUserOidList {

	count1 := len(d)
	ret := make([]edpt.SnmpServerSnmpv1V2cUserOidList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SnmpServerSnmpv1V2cUserOidList
		oi.OidVal = in["oid_val"].(string)
		oi.Remote = getObjectSnmpServerSnmpv1V2cUserOidListRemote(in["remote"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSnmpServerSnmpv1V2cUserOidListRemote(d []interface{}) edpt.SnmpServerSnmpv1V2cUserOidListRemote {

	count1 := len(d)
	var ret edpt.SnmpServerSnmpv1V2cUserOidListRemote
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.HostList = getSliceSnmpServerSnmpv1V2cUserOidListRemoteHostList(in["host_list"].([]interface{}))
		ret.Ipv4List = getSliceSnmpServerSnmpv1V2cUserOidListRemoteIpv4List(in["ipv4_list"].([]interface{}))
		ret.Ipv6List = getSliceSnmpServerSnmpv1V2cUserOidListRemoteIpv6List(in["ipv6_list"].([]interface{}))
	}
	return ret
}

func getSliceSnmpServerSnmpv1V2cUserOidListRemoteHostList(d []interface{}) []edpt.SnmpServerSnmpv1V2cUserOidListRemoteHostList {

	count1 := len(d)
	ret := make([]edpt.SnmpServerSnmpv1V2cUserOidListRemoteHostList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SnmpServerSnmpv1V2cUserOidListRemoteHostList
		oi.DnsHost = in["dns_host"].(string)
		oi.Ipv4Mask = in["ipv4_mask"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSnmpServerSnmpv1V2cUserOidListRemoteIpv4List(d []interface{}) []edpt.SnmpServerSnmpv1V2cUserOidListRemoteIpv4List {

	count1 := len(d)
	ret := make([]edpt.SnmpServerSnmpv1V2cUserOidListRemoteIpv4List, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SnmpServerSnmpv1V2cUserOidListRemoteIpv4List
		oi.Ipv4Host = in["ipv4_host"].(string)
		oi.Ipv4Mask = in["ipv4_mask"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSnmpServerSnmpv1V2cUserOidListRemoteIpv6List(d []interface{}) []edpt.SnmpServerSnmpv1V2cUserOidListRemoteIpv6List {

	count1 := len(d)
	ret := make([]edpt.SnmpServerSnmpv1V2cUserOidListRemoteIpv6List, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SnmpServerSnmpv1V2cUserOidListRemoteIpv6List
		oi.Ipv6Host = in["ipv6_host"].(string)
		oi.Ipv6Mask = in["ipv6_mask"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSnmpServerSnmpv1V2cUserRemote(d []interface{}) edpt.SnmpServerSnmpv1V2cUserRemote {

	count1 := len(d)
	var ret edpt.SnmpServerSnmpv1V2cUserRemote
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.HostList = getSliceSnmpServerSnmpv1V2cUserRemoteHostList(in["host_list"].([]interface{}))
		ret.Ipv4List = getSliceSnmpServerSnmpv1V2cUserRemoteIpv4List(in["ipv4_list"].([]interface{}))
		ret.Ipv6List = getSliceSnmpServerSnmpv1V2cUserRemoteIpv6List(in["ipv6_list"].([]interface{}))
	}
	return ret
}

func getSliceSnmpServerSnmpv1V2cUserRemoteHostList(d []interface{}) []edpt.SnmpServerSnmpv1V2cUserRemoteHostList {

	count1 := len(d)
	ret := make([]edpt.SnmpServerSnmpv1V2cUserRemoteHostList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SnmpServerSnmpv1V2cUserRemoteHostList
		oi.DnsHost = in["dns_host"].(string)
		oi.Ipv4Mask = in["ipv4_mask"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSnmpServerSnmpv1V2cUserRemoteIpv4List(d []interface{}) []edpt.SnmpServerSnmpv1V2cUserRemoteIpv4List {

	count1 := len(d)
	ret := make([]edpt.SnmpServerSnmpv1V2cUserRemoteIpv4List, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SnmpServerSnmpv1V2cUserRemoteIpv4List
		oi.Ipv4Host = in["ipv4_host"].(string)
		oi.Ipv4Mask = in["ipv4_mask"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSnmpServerSnmpv1V2cUserRemoteIpv6List(d []interface{}) []edpt.SnmpServerSnmpv1V2cUserRemoteIpv6List {

	count1 := len(d)
	ret := make([]edpt.SnmpServerSnmpv1V2cUserRemoteIpv6List, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SnmpServerSnmpv1V2cUserRemoteIpv6List
		oi.Ipv6Host = in["ipv6_host"].(string)
		oi.Ipv6Mask = in["ipv6_mask"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSnmpServerSnmpv1V2cUser(d *schema.ResourceData) edpt.SnmpServerSnmpv1V2cUser {
	var ret edpt.SnmpServerSnmpv1V2cUser
	//omit encrypted
	ret.Inst.OidList = getSliceSnmpServerSnmpv1V2cUserOidList(d.Get("oid_list").([]interface{}))
	ret.Inst.Passwd = d.Get("passwd").(string)
	ret.Inst.Remote = getObjectSnmpServerSnmpv1V2cUserRemote(d.Get("remote").([]interface{}))
	ret.Inst.User = d.Get("user").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
