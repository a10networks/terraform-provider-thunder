package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVpnIpsecGroup() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_vpn_ipsec_group`: IPsec tunnel Groupings\n\n__PLACEHOLDER__",
		CreateContext: resourceVpnIpsecGroupCreate,
		UpdateContext: resourceVpnIpsecGroupUpdate,
		ReadContext:   resourceVpnIpsecGroupRead,
		DeleteContext: resourceVpnIpsecGroupDelete,

		Schema: map[string]*schema.Schema{
			"ipsecgroup_cfg": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ipsec": {
							Type: schema.TypeString, Optional: true, Description: "specify a name to group active/backup tunnels",
						},
						"priority": {
							Type: schema.TypeInt, Optional: true, Description: "Highest priority value is the active tunnel",
						},
					},
				},
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Group name",
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
func resourceVpnIpsecGroupCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVpnIpsecGroupCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVpnIpsecGroup(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVpnIpsecGroupRead(ctx, d, meta)
	}
	return diags
}

func resourceVpnIpsecGroupUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVpnIpsecGroupUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVpnIpsecGroup(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVpnIpsecGroupRead(ctx, d, meta)
	}
	return diags
}
func resourceVpnIpsecGroupDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVpnIpsecGroupDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVpnIpsecGroup(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVpnIpsecGroupRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVpnIpsecGroupRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVpnIpsecGroup(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceVpnIpsecGroupIpsecgroupCfg(d []interface{}) []edpt.VpnIpsecGroupIpsecgroupCfg {

	count1 := len(d)
	ret := make([]edpt.VpnIpsecGroupIpsecgroupCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VpnIpsecGroupIpsecgroupCfg
		oi.Ipsec = in["ipsec"].(string)
		oi.Priority = in["priority"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointVpnIpsecGroup(d *schema.ResourceData) edpt.VpnIpsecGroup {
	var ret edpt.VpnIpsecGroup
	ret.Inst.IpsecgroupCfg = getSliceVpnIpsecGroupIpsecgroupCfg(d.Get("ipsecgroup_cfg").([]interface{}))
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
