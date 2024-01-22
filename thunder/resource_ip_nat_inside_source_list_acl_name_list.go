package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpNatInsideSourceListAclNameList() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ip_nat_inside_source_list_acl_name_list`: Apply an access list\n\n__PLACEHOLDER__",
		CreateContext: resourceIpNatInsideSourceListAclNameListCreate,
		UpdateContext: resourceIpNatInsideSourceListAclNameListUpdate,
		ReadContext:   resourceIpNatInsideSourceListAclNameListRead,
		DeleteContext: resourceIpNatInsideSourceListAclNameListDelete,

		Schema: map[string]*schema.Schema{
			"msl": {
				Type: schema.TypeInt, Optional: true, Description: "Maximum Session Life Value",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Apply an access list",
			},
			"pool": {
				Type: schema.TypeString, Optional: true, Description: "Pool or Pool Group (Pool or Pool Group Nam)",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceIpNatInsideSourceListAclNameListCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpNatInsideSourceListAclNameListCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpNatInsideSourceListAclNameList(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpNatInsideSourceListAclNameListRead(ctx, d, meta)
	}
	return diags
}

func resourceIpNatInsideSourceListAclNameListUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpNatInsideSourceListAclNameListUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpNatInsideSourceListAclNameList(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpNatInsideSourceListAclNameListRead(ctx, d, meta)
	}
	return diags
}
func resourceIpNatInsideSourceListAclNameListDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpNatInsideSourceListAclNameListDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpNatInsideSourceListAclNameList(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceIpNatInsideSourceListAclNameListRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpNatInsideSourceListAclNameListRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpNatInsideSourceListAclNameList(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointIpNatInsideSourceListAclNameList(d *schema.ResourceData) edpt.IpNatInsideSourceListAclNameList {
	var ret edpt.IpNatInsideSourceListAclNameList
	ret.Inst.Msl = d.Get("msl").(int)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.Pool = d.Get("pool").(string)
	//omit uuid
	return ret
}
