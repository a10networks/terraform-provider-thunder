package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityZbarDest() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_visibility_zbar_dest`: Configure zbar destinations\n\n__PLACEHOLDER__",
		CreateContext: resourceVisibilityZbarDestCreate,
		UpdateContext: resourceVisibilityZbarDestUpdate,
		ReadContext:   resourceVisibilityZbarDestRead,
		DeleteContext: resourceVisibilityZbarDestDelete,

		Schema: map[string]*schema.Schema{
			"bad_sources": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
func resourceVisibilityZbarDestCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityZbarDestCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityZbarDest(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityZbarDestRead(ctx, d, meta)
	}
	return diags
}

func resourceVisibilityZbarDestUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityZbarDestUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityZbarDest(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityZbarDestRead(ctx, d, meta)
	}
	return diags
}
func resourceVisibilityZbarDestDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityZbarDestDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityZbarDest(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVisibilityZbarDestRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityZbarDestRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityZbarDest(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectVisibilityZbarDestBadSources3137(d []interface{}) edpt.VisibilityZbarDestBadSources3137 {

	var ret edpt.VisibilityZbarDestBadSources3137
	return ret
}

func dataToEndpointVisibilityZbarDest(d *schema.ResourceData) edpt.VisibilityZbarDest {
	var ret edpt.VisibilityZbarDest
	ret.Inst.BadSources = getObjectVisibilityZbarDestBadSources3137(d.Get("bad_sources").([]interface{}))
	//omit uuid
	return ret
}
