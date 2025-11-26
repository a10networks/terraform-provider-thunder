package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRename() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_rename`: Rename Object Name\n\n__PLACEHOLDER__",
		CreateContext: resourceRenameCreate,
		UpdateContext: resourceRenameUpdate,
		ReadContext:   resourceRenameRead,
		DeleteContext: resourceRenameDelete,

		Schema: map[string]*schema.Schema{
			"instance_name": {
				Type: schema.TypeString, Optional: true, Description: "Old Instance Name",
			},
			"new_instance_name": {
				Type: schema.TypeString, Optional: true, Description: "New Instance Name",
			},
			"object": {
				Type: schema.TypeString, Optional: true, Description: "Lineage of object being renamed e.g: slb.server, slb.service-group, slb.virtual-server",
			},
		},
	}
}
func resourceRenameCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRenameCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRename(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRenameRead(ctx, d, meta)
	}
	return diags
}

func resourceRenameUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRenameUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRename(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRenameRead(ctx, d, meta)
	}
	return diags
}
func resourceRenameDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRenameDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRename(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceRenameRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRenameRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRename(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointRename(d *schema.ResourceData) edpt.Rename {
	var ret edpt.Rename
	ret.Inst.InstanceName = d.Get("instance_name").(string)
	ret.Inst.NewInstanceName = d.Get("new_instance_name").(string)
	ret.Inst.Object = d.Get("object").(string)
	return ret
}
