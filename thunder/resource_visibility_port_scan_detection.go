package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityPortScanDetection() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_visibility_port_scan_detection`: Configure port scan detection\n\n__PLACEHOLDER__",
		CreateContext: resourceVisibilityPortScanDetectionCreate,
		UpdateContext: resourceVisibilityPortScanDetectionUpdate,
		ReadContext:   resourceVisibilityPortScanDetectionRead,
		DeleteContext: resourceVisibilityPortScanDetectionDelete,

		Schema: map[string]*schema.Schema{
			"events": {
				Type: schema.TypeInt, Optional: true, Default: 10, Description: "Number of scanned port events (default 10)",
			},
			"interval": {
				Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time interval for port scan detection (default 60)",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"v4_list": {
				Type: schema.TypeString, Optional: true, Description: "Class list of ipv4 addresses to be whitelisted",
			},
			"v6_list": {
				Type: schema.TypeString, Optional: true, Description: "Class list of ipv6 addresses to be whitelisted",
			},
		},
	}
}
func resourceVisibilityPortScanDetectionCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPortScanDetectionCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPortScanDetection(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPortScanDetectionRead(ctx, d, meta)
	}
	return diags
}

func resourceVisibilityPortScanDetectionUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPortScanDetectionUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPortScanDetection(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPortScanDetectionRead(ctx, d, meta)
	}
	return diags
}
func resourceVisibilityPortScanDetectionDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPortScanDetectionDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPortScanDetection(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVisibilityPortScanDetectionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPortScanDetectionRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPortScanDetection(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointVisibilityPortScanDetection(d *schema.ResourceData) edpt.VisibilityPortScanDetection {
	var ret edpt.VisibilityPortScanDetection
	ret.Inst.Events = d.Get("events").(int)
	ret.Inst.Interval = d.Get("interval").(int)
	//omit uuid
	ret.Inst.V4List = d.Get("v4_list").(string)
	ret.Inst.V6List = d.Get("v6_list").(string)
	return ret
}
