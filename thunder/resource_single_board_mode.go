package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSingleBoardMode() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_single_board_mode`: chassis single board mode config\n\n__PLACEHOLDER__",
		CreateContext: resourceSingleBoardModeCreate,
		UpdateContext: resourceSingleBoardModeUpdate,
		ReadContext:   resourceSingleBoardModeRead,
		DeleteContext: resourceSingleBoardModeDelete,

		Schema: map[string]*schema.Schema{
			"fallback": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Run the chasis in single board mode if no communication with the blade board.",
			},
			"forced": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Run the chasis in single board mode only.",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSingleBoardModeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSingleBoardModeCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSingleBoardMode(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSingleBoardModeRead(ctx, d, meta)
	}
	return diags
}

func resourceSingleBoardModeUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSingleBoardModeUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSingleBoardMode(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSingleBoardModeRead(ctx, d, meta)
	}
	return diags
}
func resourceSingleBoardModeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSingleBoardModeDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSingleBoardMode(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSingleBoardModeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSingleBoardModeRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSingleBoardMode(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSingleBoardMode(d *schema.ResourceData) edpt.SingleBoardMode {
	var ret edpt.SingleBoardMode
	ret.Inst.Fallback = d.Get("fallback").(int)
	ret.Inst.Forced = d.Get("forced").(int)
	//omit uuid
	return ret
}
