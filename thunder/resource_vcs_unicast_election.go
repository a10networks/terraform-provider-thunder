package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVcsUnicastElection() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_vcs_unicast_election`: VCS Device\n\n__PLACEHOLDER__",
		CreateContext: resourceVcsUnicastElectionCreate,
		UpdateContext: resourceVcsUnicastElectionUpdate,
		ReadContext:   resourceVcsUnicastElectionRead,
		DeleteContext: resourceVcsUnicastElectionDelete,

		Schema: map[string]*schema.Schema{
			"members": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip_address_cfg": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ip_address": {
										Type: schema.TypeString, Optional: true, Description: "IP Address of the member",
									},
									"use_mgmt_port": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use management port for connections in different subnet",
									},
								},
							},
						},
						"ipv6_address_cfg": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ipv6_address": {
										Type: schema.TypeString, Optional: true, Description: "IPV6 address of the member",
									},
									"use_mgmt_port": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use management port for connections in different subnet",
									},
								},
							},
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"port": {
				Type: schema.TypeInt, Optional: true, Default: 41217, Description: "Destination port for unicast election packet (Destination port for unicast election packet (default 41217))",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceVcsUnicastElectionCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVcsUnicastElectionCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVcsUnicastElection(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVcsUnicastElectionRead(ctx, d, meta)
	}
	return diags
}

func resourceVcsUnicastElectionUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVcsUnicastElectionUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVcsUnicastElection(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVcsUnicastElectionRead(ctx, d, meta)
	}
	return diags
}
func resourceVcsUnicastElectionDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVcsUnicastElectionDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVcsUnicastElection(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVcsUnicastElectionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVcsUnicastElectionRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVcsUnicastElection(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectVcsUnicastElectionMembers1907(d []interface{}) edpt.VcsUnicastElectionMembers1907 {

	count1 := len(d)
	var ret edpt.VcsUnicastElectionMembers1907
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.IpAddressCfg = getSliceVcsUnicastElectionMembersIpAddressCfg1908(in["ip_address_cfg"].([]interface{}))
		ret.Ipv6AddressCfg = getSliceVcsUnicastElectionMembersIpv6AddressCfg1909(in["ipv6_address_cfg"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceVcsUnicastElectionMembersIpAddressCfg1908(d []interface{}) []edpt.VcsUnicastElectionMembersIpAddressCfg1908 {

	count1 := len(d)
	ret := make([]edpt.VcsUnicastElectionMembersIpAddressCfg1908, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VcsUnicastElectionMembersIpAddressCfg1908
		oi.IpAddress = in["ip_address"].(string)
		oi.UseMgmtPort = in["use_mgmt_port"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceVcsUnicastElectionMembersIpv6AddressCfg1909(d []interface{}) []edpt.VcsUnicastElectionMembersIpv6AddressCfg1909 {

	count1 := len(d)
	ret := make([]edpt.VcsUnicastElectionMembersIpv6AddressCfg1909, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VcsUnicastElectionMembersIpv6AddressCfg1909
		oi.Ipv6Address = in["ipv6_address"].(string)
		oi.UseMgmtPort = in["use_mgmt_port"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointVcsUnicastElection(d *schema.ResourceData) edpt.VcsUnicastElection {
	var ret edpt.VcsUnicastElection
	ret.Inst.Members = getObjectVcsUnicastElectionMembers1907(d.Get("members").([]interface{}))
	ret.Inst.Port = d.Get("port").(int)
	//omit uuid
	return ret
}
