package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceScaleoutDistributedForwardingFw() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_scaleout_distributed_forwarding_fw`: Enable Scaleout distributed-forwarding for firewall sessions\n\n__PLACEHOLDER__",
		CreateContext: resourceScaleoutDistributedForwardingFwCreate,
		UpdateContext: resourceScaleoutDistributedForwardingFwUpdate,
		ReadContext:   resourceScaleoutDistributedForwardingFwRead,
		DeleteContext: resourceScaleoutDistributedForwardingFwDelete,

		Schema: map[string]*schema.Schema{
			"enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable distributed-forwarding for Firewall",
			},
			"session_offload_direction": {
				Type: schema.TypeString, Optional: true, Default: "both", Description: "'uplink': Enable session offload only in uplink direction; 'downlink': Enable session offload in downlink direction; 'both': Enable session offload in both direction;",
			},
			"threshold": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"protocol_threshold": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"protocol_value": {
										Type: schema.TypeString, Optional: true, Description: "'UDP': configure threshold for udp session offload; 'TCP': configure threshold for tcp session offload;",
									},
									"threshold_value": {
										Type: schema.TypeInt, Optional: true, Default: 5, Description: "configure packet threshold value to offload sessions(default 5)",
									},
								},
							},
						},
						"global_threshold": {
							Type: schema.TypeInt, Optional: true, Default: 5, Description: "configure packet threshold value to offload sessions of any(TCP and UDP) protocol(default 5)",
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

func getObjectScaleoutDistributedForwardingFwThreshold(d []interface{}) edpt.ScaleoutDistributedForwardingFwThreshold {

	count1 := len(d)
	var ret edpt.ScaleoutDistributedForwardingFwThreshold
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ProtocolThreshold = getSliceScaleoutDistributedForwardingFwThresholdProtocolThreshold(in["protocol_threshold"].([]interface{}))
		ret.GlobalThreshold = in["global_threshold"].(int)
	}
	return ret
}

func getSliceScaleoutDistributedForwardingFwThresholdProtocolThreshold(d []interface{}) []edpt.ScaleoutDistributedForwardingFwThresholdProtocolThreshold {

	count1 := len(d)
	ret := make([]edpt.ScaleoutDistributedForwardingFwThresholdProtocolThreshold, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutDistributedForwardingFwThresholdProtocolThreshold
		oi.ProtocolValue = in["protocol_value"].(string)
		oi.ThresholdValue = in["threshold_value"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointScaleoutDistributedForwardingFw(d *schema.ResourceData) edpt.ScaleoutDistributedForwardingFw {
	var ret edpt.ScaleoutDistributedForwardingFw
	ret.Inst.Enable = d.Get("enable").(int)
	ret.Inst.SessionOffloadDirection = d.Get("session_offload_direction").(string)
	ret.Inst.Threshold = getObjectScaleoutDistributedForwardingFwThreshold(d.Get("threshold").([]interface{}))
	//omit uuid
	return ret
}
