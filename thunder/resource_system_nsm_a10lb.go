package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemNsmA10lb() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_nsm_a10lb`: NSM to A10lb communication options\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemNsmA10lbCreate,
		UpdateContext: resourceSystemNsmA10lbUpdate,
		ReadContext:   resourceSystemNsmA10lbRead,
		DeleteContext: resourceSystemNsmA10lbDelete,

		Schema: map[string]*schema.Schema{
			"kill": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "NSM will terminate a10lb if no response received",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSystemNsmA10lbCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemNsmA10lbCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemNsmA10lb(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemNsmA10lbRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemNsmA10lbUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemNsmA10lbUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemNsmA10lb(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemNsmA10lbRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemNsmA10lbDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemNsmA10lbDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemNsmA10lb(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemNsmA10lbRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemNsmA10lbRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemNsmA10lb(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSystemNsmA10lb(d *schema.ResourceData) edpt.SystemNsmA10lb {
	var ret edpt.SystemNsmA10lb
	ret.Inst.Kill = d.Get("kill").(int)
	//omit uuid
	return ret
}
