package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDebugVector() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_debug_vector`: Debug vector flow\n\n__PLACEHOLDER__",
		CreateContext: resourceDebugVectorCreate,
		UpdateContext: resourceDebugVectorUpdate,
		ReadContext:   resourceDebugVectorRead,
		DeleteContext: resourceDebugVectorDelete,

		Schema: map[string]*schema.Schema{
			"error": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Debug logs for vector errors",
			},
			"packet": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Debug logs for vector packet flow",
			},
			"trace": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Debug logs for vector callbacks",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceDebugVectorCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugVectorCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugVector(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugVectorRead(ctx, d, meta)
	}
	return diags
}

func resourceDebugVectorUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugVectorUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugVector(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugVectorRead(ctx, d, meta)
	}
	return diags
}
func resourceDebugVectorDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugVectorDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugVector(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDebugVectorRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugVectorRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugVector(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDebugVector(d *schema.ResourceData) edpt.DebugVector {
	var ret edpt.DebugVector
	ret.Inst.Error = d.Get("error").(int)
	ret.Inst.Packet = d.Get("packet").(int)
	ret.Inst.Trace = d.Get("trace").(int)
	//omit uuid
	return ret
}
