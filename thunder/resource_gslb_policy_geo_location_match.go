package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGslbPolicyGeoLocationMatch() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_gslb_policy_geo_location_match`: Specify match order of geographic\n\n__PLACEHOLDER__",
		CreateContext: resourceGslbPolicyGeoLocationMatchCreate,
		UpdateContext: resourceGslbPolicyGeoLocationMatchUpdate,
		ReadContext:   resourceGslbPolicyGeoLocationMatchRead,
		DeleteContext: resourceGslbPolicyGeoLocationMatchDelete,

		Schema: map[string]*schema.Schema{
			"geo_type_overlap": {
				Type: schema.TypeString, Optional: true, Description: "'global': Global Geo-location; 'policy': Policy Geo-location;",
			},
			"match_first": {
				Type: schema.TypeString, Optional: true, Default: "global", Description: "'global': Global Geo-location; 'policy': Policy Geo-location;",
			},
			"overlap": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable overlap mode to do longest match",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}
func resourceGslbPolicyGeoLocationMatchCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbPolicyGeoLocationMatchCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbPolicyGeoLocationMatch(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGslbPolicyGeoLocationMatchRead(ctx, d, meta)
	}
	return diags
}

func resourceGslbPolicyGeoLocationMatchUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbPolicyGeoLocationMatchUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbPolicyGeoLocationMatch(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGslbPolicyGeoLocationMatchRead(ctx, d, meta)
	}
	return diags
}
func resourceGslbPolicyGeoLocationMatchDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbPolicyGeoLocationMatchDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbPolicyGeoLocationMatch(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceGslbPolicyGeoLocationMatchRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbPolicyGeoLocationMatchRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbPolicyGeoLocationMatch(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointGslbPolicyGeoLocationMatch(d *schema.ResourceData) edpt.GslbPolicyGeoLocationMatch {
	var ret edpt.GslbPolicyGeoLocationMatch
	ret.Inst.GeoTypeOverlap = d.Get("geo_type_overlap").(string)
	ret.Inst.MatchFirst = d.Get("match_first").(string)
	ret.Inst.Overlap = d.Get("overlap").(int)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
