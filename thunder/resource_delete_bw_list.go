package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDeleteBwList() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_delete_bw_list`: Blacklist/whitelist\n\n__PLACEHOLDER__",
		CreateContext: resourceDeleteBwListCreate,
		UpdateContext: resourceDeleteBwListUpdate,
		ReadContext:   resourceDeleteBwListRead,
		DeleteContext: resourceDeleteBwListDelete,

		Schema: map[string]*schema.Schema{
			"file_name": {
				Type: schema.TypeString, Optional: true, Description: "Blacklist/whitelist Name",
			},
		},
	}
}
func resourceDeleteBwListCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDeleteBwListCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDeleteBwList(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDeleteBwListRead(ctx, d, meta)
	}
	return diags
}

func resourceDeleteBwListUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDeleteBwListUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDeleteBwList(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDeleteBwListRead(ctx, d, meta)
	}
	return diags
}
func resourceDeleteBwListDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDeleteBwListDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDeleteBwList(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDeleteBwListRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDeleteBwListRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDeleteBwList(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDeleteBwList(d *schema.ResourceData) edpt.DeleteBwList {
	var ret edpt.DeleteBwList
	ret.Inst.FileName = d.Get("file_name").(string)
	return ret
}
