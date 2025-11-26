package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemGlid() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_glid`: Apply global limiter to the whole system\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemGlidCreate,
		UpdateContext: resourceSystemGlidUpdate,
		ReadContext:   resourceSystemGlidRead,
		DeleteContext: resourceSystemGlidDelete,

		Schema: map[string]*schema.Schema{
			"glid_id": {
				Type: schema.TypeString, Optional: true, Description: "Apply limits to the whole system",
			},
			"non_shared": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Apply global limit ID to the whole system at per data cpu level (default disabled)",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSystemGlidCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemGlidCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemGlid(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemGlidRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemGlidUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemGlidUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemGlid(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemGlidRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemGlidDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemGlidDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemGlid(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemGlidRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemGlidRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemGlid(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSystemGlid(d *schema.ResourceData) edpt.SystemGlid {
	var ret edpt.SystemGlid
	ret.Inst.GlidId = d.Get("glid_id").(string)
	ret.Inst.NonShared = d.Get("non_shared").(int)
	//omit uuid
	return ret
}
