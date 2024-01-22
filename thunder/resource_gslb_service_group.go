package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGslbServiceGroup() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_gslb_service_group`: Specify GSLB Service Group\n\n__PLACEHOLDER__",
		CreateContext: resourceGslbServiceGroupCreate,
		UpdateContext: resourceGslbServiceGroupUpdate,
		ReadContext:   resourceGslbServiceGroupRead,
		DeleteContext: resourceGslbServiceGroupDelete,

		Schema: map[string]*schema.Schema{
			"dependency_site": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Dependency on site",
			},
			"disable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable all members",
			},
			"disable_site_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"disable_site": {
							Type: schema.TypeString, Optional: true, Description: "Site name",
						},
					},
				},
			},
			"member": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"member_name": {
							Type: schema.TypeString, Optional: true, Description: "Service name",
						},
					},
				},
			},
			"persistent_aging_time": {
				Type: schema.TypeInt, Optional: true, Default: 5, Description: "Specify aging-time, unit: min, default is 5 (Aging time)",
			},
			"persistent_ipv6_mask": {
				Type: schema.TypeInt, Optional: true, Default: 128, Description: "Specify IPv6 mask length, default is 128",
			},
			"persistent_mask": {
				Type: schema.TypeString, Optional: true, Default: "/32", Description: "Specify IP mask, default is /32",
			},
			"persistent_site": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Persistent based on site",
			},
			"service_group_name": {
				Type: schema.TypeString, Required: true, Description: "Specify Service Group name",
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
func resourceGslbServiceGroupCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbServiceGroupCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbServiceGroup(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGslbServiceGroupRead(ctx, d, meta)
	}
	return diags
}

func resourceGslbServiceGroupUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbServiceGroupUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbServiceGroup(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGslbServiceGroupRead(ctx, d, meta)
	}
	return diags
}
func resourceGslbServiceGroupDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbServiceGroupDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbServiceGroup(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceGslbServiceGroupRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbServiceGroupRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbServiceGroup(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceGslbServiceGroupDisableSiteList(d []interface{}) []edpt.GslbServiceGroupDisableSiteList {

	count1 := len(d)
	ret := make([]edpt.GslbServiceGroupDisableSiteList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbServiceGroupDisableSiteList
		oi.DisableSite = in["disable_site"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceGslbServiceGroupMember(d []interface{}) []edpt.GslbServiceGroupMember {

	count1 := len(d)
	ret := make([]edpt.GslbServiceGroupMember, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbServiceGroupMember
		oi.MemberName = in["member_name"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointGslbServiceGroup(d *schema.ResourceData) edpt.GslbServiceGroup {
	var ret edpt.GslbServiceGroup
	ret.Inst.DependencySite = d.Get("dependency_site").(int)
	ret.Inst.Disable = d.Get("disable").(int)
	ret.Inst.DisableSiteList = getSliceGslbServiceGroupDisableSiteList(d.Get("disable_site_list").([]interface{}))
	ret.Inst.Member = getSliceGslbServiceGroupMember(d.Get("member").([]interface{}))
	ret.Inst.PersistentAgingTime = d.Get("persistent_aging_time").(int)
	ret.Inst.PersistentIpv6Mask = d.Get("persistent_ipv6_mask").(int)
	ret.Inst.PersistentMask = d.Get("persistent_mask").(string)
	ret.Inst.PersistentSite = d.Get("persistent_site").(int)
	ret.Inst.ServiceGroupName = d.Get("service_group_name").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
