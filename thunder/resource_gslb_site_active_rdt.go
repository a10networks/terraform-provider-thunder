package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGslbSiteActiveRdt() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_gslb_site_active_rdt`: Active RDT options\n\n__PLACEHOLDER__",
		CreateContext: resourceGslbSiteActiveRdtCreate,
		UpdateContext: resourceGslbSiteActiveRdtUpdate,
		ReadContext:   resourceGslbSiteActiveRdtRead,
		DeleteContext: resourceGslbSiteActiveRdtDelete,

		Schema: map[string]*schema.Schema{
			"aging_time": {
				Type: schema.TypeInt, Optional: true, Default: 10, Description: "Aging Time, Unit: min, default is 10",
			},
			"bind_geoloc": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Bind RDT to geo-location",
			},
			"ignore_count": {
				Type: schema.TypeInt, Optional: true, Default: 5, Description: "Ignore count if RDT is out of range, default is 5",
			},
			"ipv6_mask": {
				Type: schema.TypeInt, Optional: true, Default: 128, Description: "Client IPv6 subnet mask, default is 128",
			},
			"limit": {
				Type: schema.TypeInt, Optional: true, Default: 16383, Description: "Limit of valid RDT, default is 16383 (Limit, unit: millisecond)",
			},
			"mask": {
				Type: schema.TypeString, Optional: true, Default: "/32", Description: "Client IP subnet mask, default is 32",
			},
			"overlap": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable overlap for geo-location to do longest match",
			},
			"range_factor": {
				Type: schema.TypeInt, Optional: true, Default: 25, Description: "Factor of RDT Range, default is 25 (Range Factor of Smooth RDT)",
			},
			"smooth_factor": {
				Type: schema.TypeInt, Optional: true, Default: 10, Description: "Factor of Smooth RDT, default is 10",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"site_name": {
				Type: schema.TypeString, Required: true, Description: "SiteName",
			},
		},
	}
}
func resourceGslbSiteActiveRdtCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbSiteActiveRdtCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbSiteActiveRdt(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGslbSiteActiveRdtRead(ctx, d, meta)
	}
	return diags
}

func resourceGslbSiteActiveRdtUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbSiteActiveRdtUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbSiteActiveRdt(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGslbSiteActiveRdtRead(ctx, d, meta)
	}
	return diags
}
func resourceGslbSiteActiveRdtDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbSiteActiveRdtDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbSiteActiveRdt(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceGslbSiteActiveRdtRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbSiteActiveRdtRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbSiteActiveRdt(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointGslbSiteActiveRdt(d *schema.ResourceData) edpt.GslbSiteActiveRdt {
	var ret edpt.GslbSiteActiveRdt
	ret.Inst.AgingTime = d.Get("aging_time").(int)
	ret.Inst.BindGeoloc = d.Get("bind_geoloc").(int)
	ret.Inst.IgnoreCount = d.Get("ignore_count").(int)
	ret.Inst.Ipv6Mask = d.Get("ipv6_mask").(int)
	ret.Inst.Limit = d.Get("limit").(int)
	ret.Inst.Mask = d.Get("mask").(string)
	ret.Inst.Overlap = d.Get("overlap").(int)
	ret.Inst.RangeFactor = d.Get("range_factor").(int)
	ret.Inst.SmoothFactor = d.Get("smooth_factor").(int)
	//omit uuid
	ret.Inst.SiteName = d.Get("site_name").(string)
	return ret
}
