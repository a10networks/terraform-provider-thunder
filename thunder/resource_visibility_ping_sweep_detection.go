package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityPingSweepDetection() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_visibility_ping_sweep_detection`: Configure ping sweep detection\n\n__PLACEHOLDER__",
		CreateContext: resourceVisibilityPingSweepDetectionCreate,
		UpdateContext: resourceVisibilityPingSweepDetectionUpdate,
		ReadContext:   resourceVisibilityPingSweepDetectionRead,
		DeleteContext: resourceVisibilityPingSweepDetectionDelete,

		Schema: map[string]*schema.Schema{
			"events": {
				Type: schema.TypeInt, Optional: true, Default: 10, Description: "Number of scanned ip events (default 10)",
			},
			"interval": {
				Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time interval for ip sweep detection (default 60)",
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
func resourceVisibilityPingSweepDetectionCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPingSweepDetectionCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPingSweepDetection(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPingSweepDetectionRead(ctx, d, meta)
	}
	return diags
}

func resourceVisibilityPingSweepDetectionUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPingSweepDetectionUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPingSweepDetection(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPingSweepDetectionRead(ctx, d, meta)
	}
	return diags
}
func resourceVisibilityPingSweepDetectionDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPingSweepDetectionDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPingSweepDetection(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVisibilityPingSweepDetectionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPingSweepDetectionRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPingSweepDetection(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointVisibilityPingSweepDetection(d *schema.ResourceData) edpt.VisibilityPingSweepDetection {
	var ret edpt.VisibilityPingSweepDetection
	ret.Inst.Events = d.Get("events").(int)
	ret.Inst.Interval = d.Get("interval").(int)
	//omit uuid
	ret.Inst.V4List = d.Get("v4_list").(string)
	ret.Inst.V6List = d.Get("v6_list").(string)
	return ret
}
