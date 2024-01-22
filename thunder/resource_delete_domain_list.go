package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDeleteDomainList() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_delete_domain_list`: Delete ddos file\n\n__PLACEHOLDER__",
		CreateContext: resourceDeleteDomainListCreate,
		UpdateContext: resourceDeleteDomainListUpdate,
		ReadContext:   resourceDeleteDomainListRead,
		DeleteContext: resourceDeleteDomainListDelete,

		Schema: map[string]*schema.Schema{
			"domain_list_name": {
				Type: schema.TypeString, Optional: true, Description: "Specify file to be deleted",
			},
		},
	}
}
func resourceDeleteDomainListCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDeleteDomainListCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDeleteDomainList(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDeleteDomainListRead(ctx, d, meta)
	}
	return diags
}

func resourceDeleteDomainListUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDeleteDomainListUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDeleteDomainList(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDeleteDomainListRead(ctx, d, meta)
	}
	return diags
}
func resourceDeleteDomainListDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDeleteDomainListDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDeleteDomainList(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDeleteDomainListRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDeleteDomainListRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDeleteDomainList(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDeleteDomainList(d *schema.ResourceData) edpt.DeleteDomainList {
	var ret edpt.DeleteDomainList
	ret.Inst.DomainListName = d.Get("domain_list_name").(string)
	return ret
}
