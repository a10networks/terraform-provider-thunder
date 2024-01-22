package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceScaleoutDistributedForwardingCgn() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_scaleout_distributed_forwarding_cgn`: Enable Scaleout distributed-forwarding for cgn\n\n__PLACEHOLDER__",
		CreateContext: resourceScaleoutDistributedForwardingCgnCreate,
		UpdateContext: resourceScaleoutDistributedForwardingCgnUpdate,
		ReadContext:   resourceScaleoutDistributedForwardingCgnRead,
		DeleteContext: resourceScaleoutDistributedForwardingCgnDelete,

		Schema: map[string]*schema.Schema{
			"cgn_value": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable CGN; 'disable': Disable CGN;",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceScaleoutDistributedForwardingCgnCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutDistributedForwardingCgnCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutDistributedForwardingCgn(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceScaleoutDistributedForwardingCgnRead(ctx, d, meta)
	}
	return diags
}

func resourceScaleoutDistributedForwardingCgnUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutDistributedForwardingCgnUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutDistributedForwardingCgn(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceScaleoutDistributedForwardingCgnRead(ctx, d, meta)
	}
	return diags
}
func resourceScaleoutDistributedForwardingCgnDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutDistributedForwardingCgnDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutDistributedForwardingCgn(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceScaleoutDistributedForwardingCgnRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutDistributedForwardingCgnRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutDistributedForwardingCgn(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointScaleoutDistributedForwardingCgn(d *schema.ResourceData) edpt.ScaleoutDistributedForwardingCgn {
	var ret edpt.ScaleoutDistributedForwardingCgn
	ret.Inst.CgnValue = d.Get("cgn_value").(string)
	//omit uuid
	return ret
}
