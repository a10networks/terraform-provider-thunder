package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpNatPoolGroup() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ip_nat_pool_group`: IP pool group name\n\n__PLACEHOLDER__",
		CreateContext: resourceIpNatPoolGroupCreate,
		UpdateContext: resourceIpNatPoolGroupUpdate,
		ReadContext:   resourceIpNatPoolGroupRead,
		DeleteContext: resourceIpNatPoolGroupDelete,

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
func resourceIpNatPoolGroupCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpNatPoolGroupCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpNatPoolGroup(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpNatPoolGroupRead(ctx, d, meta)
	}
	return diags
}

func resourceIpNatPoolGroupUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpNatPoolGroupUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpNatPoolGroup(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpNatPoolGroupRead(ctx, d, meta)
	}
	return diags
}
func resourceIpNatPoolGroupDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpNatPoolGroupDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpNatPoolGroup(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceIpNatPoolGroupRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpNatPoolGroupRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpNatPoolGroup(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceIpNatPoolGroupMemberList(d []interface{}) []edpt.IpNatPoolGroupMemberList {

	count1 := len(d)
	ret := make([]edpt.IpNatPoolGroupMemberList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.IpNatPoolGroupMemberList
		oi.PoolName = in["pool_name"].(string)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getSliceIpNatPoolGroupSamplingEnable(d []interface{}) []edpt.IpNatPoolGroupSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.IpNatPoolGroupSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.IpNatPoolGroupSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointIpNatPoolGroup(d *schema.ResourceData) edpt.IpNatPoolGroup {
	var ret edpt.IpNatPoolGroup
	ret.Inst.MemberList = getSliceIpNatPoolGroupMemberList(d.Get("member_list").([]interface{}))
	ret.Inst.PoolGroupName = d.Get("pool_group_name").(string)
	ret.Inst.SamplingEnable = getSliceIpNatPoolGroupSamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.Vrid = d.Get("vrid").(int)
	return ret
}
