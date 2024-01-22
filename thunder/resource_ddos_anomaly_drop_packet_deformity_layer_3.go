package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosAnomalyDropPacketDeformityLayer3() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_anomaly_drop_packet_deformity_layer_3`: Disable (Enable) Layer-3 Packet Deformities\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosAnomalyDropPacketDeformityLayer3Create,
		UpdateContext: resourceDdosAnomalyDropPacketDeformityLayer3Update,
		ReadContext:   resourceDdosAnomalyDropPacketDeformityLayer3Read,
		DeleteContext: resourceDdosAnomalyDropPacketDeformityLayer3Delete,

		Schema: map[string]*schema.Schema{
			"capture_config": {
				Type: schema.TypeString, Optional: true, Description: "capture-config name (Can only configure when drop-disabled)",
			},
			"log": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Log the anomaly event (Can only configure when drop-disabled)",
			},
			"toggle": {
				Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': enable; 'disable': disable;",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceDdosAnomalyDropPacketDeformityLayer3Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosAnomalyDropPacketDeformityLayer3Create()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosAnomalyDropPacketDeformityLayer3(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosAnomalyDropPacketDeformityLayer3Read(ctx, d, meta)
	}
	return diags
}

func resourceDdosAnomalyDropPacketDeformityLayer3Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosAnomalyDropPacketDeformityLayer3Update()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosAnomalyDropPacketDeformityLayer3(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosAnomalyDropPacketDeformityLayer3Read(ctx, d, meta)
	}
	return diags
}
func resourceDdosAnomalyDropPacketDeformityLayer3Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosAnomalyDropPacketDeformityLayer3Delete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosAnomalyDropPacketDeformityLayer3(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosAnomalyDropPacketDeformityLayer3Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosAnomalyDropPacketDeformityLayer3Read()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosAnomalyDropPacketDeformityLayer3(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDdosAnomalyDropPacketDeformityLayer3(d *schema.ResourceData) edpt.DdosAnomalyDropPacketDeformityLayer3 {
	var ret edpt.DdosAnomalyDropPacketDeformityLayer3
	ret.Inst.CaptureConfig = d.Get("capture_config").(string)
	ret.Inst.Log = d.Get("log").(int)
	ret.Inst.Toggle = d.Get("toggle").(string)
	//omit uuid
	return ret
}
