package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

//AxAPI cannot handle "operation" and "configuration" fields at the same time. Break /axapi/v3/slb/common to 2 resources
func resourceSlbCommonBufferThreshold() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSlbCommonBufferThresholdCreate,
		ReadContext:   resourceSlbCommonBufferThresholdRead,
		DeleteContext: resourceSlbCommonBufferThresholdDelete,
		Schema: map[string]*schema.Schema{
			"buff_thresh": {
				Type: schema.TypeInt, Required: true, ForceNew: true, Description: "Set buffer threshold",
				ValidateFunc: validation.IntBetween(0, 1),
			},
			"buff_thresh_hw_buff": {
				Type: schema.TypeInt, Optional: true, ForceNew: true, Description: "Set hardware buffer threshold",
				RequiredWith: []string{"buff_thresh_hw_buff", "buff_thresh_relieve_thresh", "buff_thresh_sys_buff_high", "buff_thresh_sys_buff_low"},
				ValidateFunc: validation.IntBetween(1, 2147483647),
			},
			"buff_thresh_relieve_thresh": {
				Type: schema.TypeInt, Optional: true, ForceNew: true, Description: "Relieve threshold",
				RequiredWith: []string{"buff_thresh_hw_buff", "buff_thresh_relieve_thresh", "buff_thresh_sys_buff_high", "buff_thresh_sys_buff_low"},
				ValidateFunc: validation.IntBetween(0, 2147483647),
			},
			"buff_thresh_sys_buff_high": {
				Type: schema.TypeInt, Optional: true, ForceNew: true, Description: "Set high water mark of system buffer",
				RequiredWith: []string{"buff_thresh_hw_buff", "buff_thresh_relieve_thresh", "buff_thresh_sys_buff_high", "buff_thresh_sys_buff_low"},
				ValidateFunc: validation.IntBetween(0, 2147483647),
			},
			"buff_thresh_sys_buff_low": {
				Type: schema.TypeInt, Optional: true, ForceNew: true, Description: "Set low water mark of system buffer",
				RequiredWith: []string{"buff_thresh_hw_buff", "buff_thresh_relieve_thresh", "buff_thresh_sys_buff_high", "buff_thresh_sys_buff_low"},
				ValidateFunc: validation.IntBetween(0, 2147483647),
			},
		},
	}
}

func resourceSlbCommonBufferThresholdCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbCommonBufferThresholdCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbCommonBufferThreshold(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbCommonBufferThresholdRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbCommonBufferThresholdRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbCommonBufferThresholdRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := edpt.SlbCommonBufferThreshold{}
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbCommonBufferThresholdDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbCommonBufferThresholdDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := edpt.SlbCommonBufferThreshold{}
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSlbCommonBufferThreshold(d *schema.ResourceData) edpt.SlbCommonBufferThreshold {
	var ret edpt.SlbCommonBufferThreshold
	ret.Inst.BuffThresh = d.Get("buff_thresh").(int)
	if ret.Inst.BuffThresh == 1 {
		ret.Inst.BuffThreshHwBuff = d.Get("buff_thresh_hw_buff").(int)
		ret.Inst.BuffThreshRelieveThresh = d.Get("buff_thresh_relieve_thresh").(int)
		ret.Inst.BuffThreshSysBuffHigh = d.Get("buff_thresh_sys_buff_high").(int)
		ret.Inst.BuffThreshSysBuffLow = d.Get("buff_thresh_sys_buff_low").(int)
	}
	return ret
}
