package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceScaleoutDistributedForwardingFw() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_scaleout_distributed_forwarding_fw`: Enable Scaleout distributed-forwarding for fw\n\n__PLACEHOLDER__",
		CreateContext: resourceScaleoutDistributedForwardingFwCreate,
		UpdateContext: resourceScaleoutDistributedForwardingFwUpdate,
		ReadContext:   resourceScaleoutDistributedForwardingFwRead,
		DeleteContext: resourceScaleoutDistributedForwardingFwDelete,

		Schema: map[string]*schema.Schema{
			"fw_value": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable FW; 'disable': Disable FW;",
			},
			"session_offload_direction": {
				Type: schema.TypeString, Optional: true, Default: "both", Description: "'uplink': Enable session offload only in uplink direction; 'downlink': Enable session offload in downlink direction; 'both': Enable session offload in both direction;",
			},
			"threshold": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"threshold_value": {
							Type: schema.TypeInt, Optional: true, Default: 5, Description: "configure packet threshold value to offload sessions(default 5)",
						},
						"protocol_value": {
							Type: schema.TypeString, Optional: true, Description: "'UDP': configure threshold for udp session offload; 'TCP': configure threshold for tcp session offload;",
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
func resourceScaleoutDistributedForwardingFwCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutDistributedForwardingFwCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutDistributedForwardingFw(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceScaleoutDistributedForwardingFwRead(ctx, d, meta)
	}
	return diags
}

func resourceScaleoutDistributedForwardingFwUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutDistributedForwardingFwUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutDistributedForwardingFw(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceScaleoutDistributedForwardingFwRead(ctx, d, meta)
	}
	return diags
}
func resourceScaleoutDistributedForwardingFwDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutDistributedForwardingFwDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutDistributedForwardingFw(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceScaleoutDistributedForwardingFwRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutDistributedForwardingFwRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutDistributedForwardingFw(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceScaleoutDistributedForwardingFwThreshold(d []interface{}) []edpt.ScaleoutDistributedForwardingFwThreshold {

	count1 := len(d)
	ret := make([]edpt.ScaleoutDistributedForwardingFwThreshold, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutDistributedForwardingFwThreshold
		oi.ThresholdValue = in["threshold_value"].(int)
		oi.ProtocolValue = in["protocol_value"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointScaleoutDistributedForwardingFw(d *schema.ResourceData) edpt.ScaleoutDistributedForwardingFw {
	var ret edpt.ScaleoutDistributedForwardingFw
	ret.Inst.FwValue = d.Get("fw_value").(string)
	ret.Inst.SessionOffloadDirection = d.Get("session_offload_direction").(string)
	ret.Inst.Threshold = getSliceScaleoutDistributedForwardingFwThreshold(d.Get("threshold").([]interface{}))
	//omit uuid
	return ret
}
