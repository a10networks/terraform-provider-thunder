package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpv6NatPoolGroup() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ipv6_nat_pool_group`: IPv6 pool group name\n\n__PLACEHOLDER__",
		CreateContext: resourceIpv6NatPoolGroupCreate,
		UpdateContext: resourceIpv6NatPoolGroupUpdate,
		ReadContext:   resourceIpv6NatPoolGroupRead,
		DeleteContext: resourceIpv6NatPoolGroupDelete,

		Schema: map[string]*schema.Schema{
			"member_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"pool_name": {
							Type: schema.TypeString, Required: true, Description: "Specify NAT pool name",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"pool_group_name": {
				Type: schema.TypeString, Required: true, Description: "Specify pool group name",
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'Failed': Failed;",
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
			"vrid": {
				Type: schema.TypeInt, Optional: true, Description: "Specify VRRP-A vrid (Specify ha VRRP-A vrid)",
			},
		},
	}
}
func resourceIpv6NatPoolGroupCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpv6NatPoolGroupCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpv6NatPoolGroup(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpv6NatPoolGroupRead(ctx, d, meta)
	}
	return diags
}

func resourceIpv6NatPoolGroupUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpv6NatPoolGroupUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpv6NatPoolGroup(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpv6NatPoolGroupRead(ctx, d, meta)
	}
	return diags
}
func resourceIpv6NatPoolGroupDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpv6NatPoolGroupDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpv6NatPoolGroup(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceIpv6NatPoolGroupRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpv6NatPoolGroupRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpv6NatPoolGroup(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceIpv6NatPoolGroupMemberList(d []interface{}) []edpt.Ipv6NatPoolGroupMemberList {

	count1 := len(d)
	ret := make([]edpt.Ipv6NatPoolGroupMemberList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Ipv6NatPoolGroupMemberList
		oi.PoolName = in["pool_name"].(string)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getSliceIpv6NatPoolGroupSamplingEnable(d []interface{}) []edpt.Ipv6NatPoolGroupSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.Ipv6NatPoolGroupSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Ipv6NatPoolGroupSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointIpv6NatPoolGroup(d *schema.ResourceData) edpt.Ipv6NatPoolGroup {
	var ret edpt.Ipv6NatPoolGroup
	ret.Inst.MemberList = getSliceIpv6NatPoolGroupMemberList(d.Get("member_list").([]interface{}))
	ret.Inst.PoolGroupName = d.Get("pool_group_name").(string)
	ret.Inst.SamplingEnable = getSliceIpv6NatPoolGroupSamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.Vrid = d.Get("vrid").(int)
	return ret
}
