package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpNatInsideSourceListAclIdList() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ip_nat_inside_source_list_acl_id_list`: Acl id\n\n__PLACEHOLDER__",
		CreateContext: resourceIpNatInsideSourceListAclIdListCreate,
		UpdateContext: resourceIpNatInsideSourceListAclIdListUpdate,
		ReadContext:   resourceIpNatInsideSourceListAclIdListRead,
		DeleteContext: resourceIpNatInsideSourceListAclIdListDelete,

		Schema: map[string]*schema.Schema{
			"acl_id": {
				Type: schema.TypeInt, Required: true, Description: "Acl id",
			},
			"msl": {
				Type: schema.TypeInt, Optional: true, Description: "Maximum Session Life Value",
			},
			"pool": {
				Type: schema.TypeString, Optional: true, Description: "Pool or Pool Group (Pool or Pool Group Name)",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceIpNatInsideSourceListAclIdListCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpNatInsideSourceListAclIdListCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpNatInsideSourceListAclIdList(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpNatInsideSourceListAclIdListRead(ctx, d, meta)
	}
	return diags
}

func resourceIpNatInsideSourceListAclIdListUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpNatInsideSourceListAclIdListUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpNatInsideSourceListAclIdList(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpNatInsideSourceListAclIdListRead(ctx, d, meta)
	}
	return diags
}
func resourceIpNatInsideSourceListAclIdListDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpNatInsideSourceListAclIdListDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpNatInsideSourceListAclIdList(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceIpNatInsideSourceListAclIdListRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpNatInsideSourceListAclIdListRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpNatInsideSourceListAclIdList(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointIpNatInsideSourceListAclIdList(d *schema.ResourceData) edpt.IpNatInsideSourceListAclIdList {
	var ret edpt.IpNatInsideSourceListAclIdList
	ret.Inst.AclId = d.Get("acl_id").(int)
	ret.Inst.Msl = d.Get("msl").(int)
	ret.Inst.Pool = d.Get("pool").(string)
	//omit uuid
	return ret
}
