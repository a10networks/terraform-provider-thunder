package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstZoneCaptureConfig() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_dst_zone_capture_config`: Specify packet capture-config\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDstZoneCaptureConfigCreate,
		UpdateContext: resourceDdosDstZoneCaptureConfigUpdate,
		ReadContext:   resourceDdosDstZoneCaptureConfigRead,
		DeleteContext: resourceDdosDstZoneCaptureConfigDelete,

		Schema: map[string]*schema.Schema{
			"mode": {
				Type: schema.TypeString, Optional: true, Description: "'drop': Apply capture-config to dropped packets; 'forward': Apply capture-config to forwarded packets; 'all': Apply capture-config to both dropped and forwarded packets;",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Capture-config name",
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
func resourceDdosDstZoneCaptureConfigCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneCaptureConfigCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneCaptureConfig(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstZoneCaptureConfigRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosDstZoneCaptureConfigUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneCaptureConfigUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneCaptureConfig(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstZoneCaptureConfigRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosDstZoneCaptureConfigDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneCaptureConfigDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneCaptureConfig(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDstZoneCaptureConfigRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneCaptureConfigRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneCaptureConfig(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDdosDstZoneCaptureConfig(d *schema.ResourceData) edpt.DdosDstZoneCaptureConfig {
	var ret edpt.DdosDstZoneCaptureConfig
	ret.Inst.Mode = d.Get("mode").(string)
	ret.Inst.Name = d.Get("name").(string)
	//omit uuid
	ret.Inst.ZoneName = d.Get("zone_name").(string)
	return ret
}
