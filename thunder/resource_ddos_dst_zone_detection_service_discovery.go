package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstZoneDetectionServiceDiscovery() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_dst_zone_detection_service_discovery`: Service discovery Configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDstZoneDetectionServiceDiscoveryCreate,
		UpdateContext: resourceDdosDstZoneDetectionServiceDiscoveryUpdate,
		ReadContext:   resourceDdosDstZoneDetectionServiceDiscoveryRead,
		DeleteContext: resourceDdosDstZoneDetectionServiceDiscoveryDelete,

		Schema: map[string]*schema.Schema{
			"configuration": {
				Type: schema.TypeString, Required: true, Description: "'configuration': configuration;",
			},
			"pkt_rate_threshold": {
				Type: schema.TypeInt, Optional: true, Default: 10, Description: "packet rate threshold for discovery (default 10 packets per second)",
			},
			"toggle": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable service discovery; 'disable': Disable service discovery;",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"zone_name": {
				Type: schema.TypeString, Required: true, Description: "ZoneName",
			},
		},
	}
}
func resourceDdosDstZoneDetectionServiceDiscoveryCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneDetectionServiceDiscoveryCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneDetectionServiceDiscovery(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstZoneDetectionServiceDiscoveryRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosDstZoneDetectionServiceDiscoveryUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneDetectionServiceDiscoveryUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneDetectionServiceDiscovery(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstZoneDetectionServiceDiscoveryRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosDstZoneDetectionServiceDiscoveryDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneDetectionServiceDiscoveryDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneDetectionServiceDiscovery(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDstZoneDetectionServiceDiscoveryRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneDetectionServiceDiscoveryRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneDetectionServiceDiscovery(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDdosDstZoneDetectionServiceDiscovery(d *schema.ResourceData) edpt.DdosDstZoneDetectionServiceDiscovery {
	var ret edpt.DdosDstZoneDetectionServiceDiscovery
	ret.Inst.Configuration = d.Get("configuration").(string)
	ret.Inst.PktRateThreshold = d.Get("pkt_rate_threshold").(int)
	ret.Inst.Toggle = d.Get("toggle").(string)
	//omit uuid
	ret.Inst.ZoneName = d.Get("zone_name").(string)
	return ret
}
