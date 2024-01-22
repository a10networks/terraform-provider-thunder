package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceMaximumPaths() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_maximum_paths`: Set maximum number of route multipaths installed into FIB\n\n__PLACEHOLDER__",
		CreateContext: resourceMaximumPathsCreate,
		UpdateContext: resourceMaximumPathsUpdate,
		ReadContext:   resourceMaximumPathsRead,
		DeleteContext: resourceMaximumPathsDelete,

		Schema: map[string]*schema.Schema{
			"path": {
				Type: schema.TypeInt, Optional: true, Default: 4, Description: "supported multipath numbers",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceMaximumPathsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceMaximumPathsCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointMaximumPaths(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceMaximumPathsRead(ctx, d, meta)
	}
	return diags
}

func resourceMaximumPathsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceMaximumPathsUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointMaximumPaths(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceMaximumPathsRead(ctx, d, meta)
	}
	return diags
}
func resourceMaximumPathsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceMaximumPathsDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointMaximumPaths(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceMaximumPathsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceMaximumPathsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointMaximumPaths(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointMaximumPaths(d *schema.ResourceData) edpt.MaximumPaths {
	var ret edpt.MaximumPaths
	ret.Inst.Path = d.Get("path").(int)
	//omit uuid
	return ret
}
