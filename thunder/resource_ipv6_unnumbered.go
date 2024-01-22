package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpv6Unnumbered() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ipv6_unnumbered`: unnumbered feature\n\n__PLACEHOLDER__",
		CreateContext: resourceIpv6UnnumberedCreate,
		UpdateContext: resourceIpv6UnnumberedUpdate,
		ReadContext:   resourceIpv6UnnumberedRead,
		DeleteContext: resourceIpv6UnnumberedDelete,

		Schema: map[string]*schema.Schema{
			"use_source_acl": {
				Type: schema.TypeString, Optional: true, Description: "Access List, Upon deny packet's source address will be NATed (ACL Name)",
			},
			"use_source_ipv6": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"update_source_ipv6": {
							Type: schema.TypeString, Optional: true, Description: "IPv6 address",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceIpv6UnnumberedCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpv6UnnumberedCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpv6Unnumbered(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpv6UnnumberedRead(ctx, d, meta)
	}
	return diags
}

func resourceIpv6UnnumberedUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpv6UnnumberedUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpv6Unnumbered(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpv6UnnumberedRead(ctx, d, meta)
	}
	return diags
}
func resourceIpv6UnnumberedDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpv6UnnumberedDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpv6Unnumbered(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceIpv6UnnumberedRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpv6UnnumberedRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpv6Unnumbered(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectIpv6UnnumberedUseSourceIpv61045(d []interface{}) edpt.Ipv6UnnumberedUseSourceIpv61045 {

	count1 := len(d)
	var ret edpt.Ipv6UnnumberedUseSourceIpv61045
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.UpdateSourceIpv6 = in["update_source_ipv6"].(string)
		//omit uuid
	}
	return ret
}

func dataToEndpointIpv6Unnumbered(d *schema.ResourceData) edpt.Ipv6Unnumbered {
	var ret edpt.Ipv6Unnumbered
	ret.Inst.UseSourceAcl = d.Get("use_source_acl").(string)
	ret.Inst.UseSourceIpv6 = getObjectIpv6UnnumberedUseSourceIpv61045(d.Get("use_source_ipv6").([]interface{}))
	//omit uuid
	return ret
}
