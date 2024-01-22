package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFlowspecOperationalMode() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_flowspec_operational_mode`: Set the operational mode for the Flowspec\n\n__PLACEHOLDER__",
		CreateContext: resourceFlowspecOperationalModeCreate,
		UpdateContext: resourceFlowspecOperationalModeUpdate,
		ReadContext:   resourceFlowspecOperationalModeRead,
		DeleteContext: resourceFlowspecOperationalModeDelete,

		Schema: map[string]*schema.Schema{
			"mode": {
				Type: schema.TypeString, Optional: true, Default: "disabled", Description: "'enabled': Enable the flowspec and send the prefix to BGP; 'disabled': Disable the flowspec and remove the prefix from BGP;",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}
func resourceFlowspecOperationalModeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFlowspecOperationalModeCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFlowspecOperationalMode(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFlowspecOperationalModeRead(ctx, d, meta)
	}
	return diags
}

func resourceFlowspecOperationalModeUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFlowspecOperationalModeUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFlowspecOperationalMode(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFlowspecOperationalModeRead(ctx, d, meta)
	}
	return diags
}
func resourceFlowspecOperationalModeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFlowspecOperationalModeDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFlowspecOperationalMode(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFlowspecOperationalModeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFlowspecOperationalModeRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFlowspecOperationalMode(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointFlowspecOperationalMode(d *schema.ResourceData) edpt.FlowspecOperationalMode {
	var ret edpt.FlowspecOperationalMode
	ret.Inst.Mode = d.Get("mode").(string)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
