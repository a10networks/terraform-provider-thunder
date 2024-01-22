package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGslbZoneServiceGeoLocation() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_gslb_zone_service_geo_location`: Geo location settings\n\n__PLACEHOLDER__",
		CreateContext: resourceGslbZoneServiceGeoLocationCreate,
		UpdateContext: resourceGslbZoneServiceGeoLocationUpdate,
		ReadContext:   resourceGslbZoneServiceGeoLocationRead,
		DeleteContext: resourceGslbZoneServiceGeoLocationDelete,

		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Action for this geo-location",
			},
			"action_type": {
				Type: schema.TypeString, Optional: true, Description: "'allow': Allow query from this geo-location; 'drop': Drop query from this geo-location; 'forward': Forward packet for this geo-location; 'ignore': Send empty response to this geo-location; 'reject': Send refuse response to this geo-location;",
			},
			"alias": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"alias": {
							Type: schema.TypeString, Optional: true, Description: "Send CNAME response for this geo-location (Specify a CNAME record)",
						},
					},
				},
			},
			"forward_type": {
				Type: schema.TypeString, Optional: true, Description: "'both': Forward both query and response; 'query': Forward query from this geo-location; 'response': Forward response to this geo-location;",
			},
			"geo_name": {
				Type: schema.TypeString, Required: true, Description: "Specify the geo-location",
			},
			"policy": {
				Type: schema.TypeString, Optional: true, Description: "Policy for this geo-location (Specify the policy name)",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"service_name": {
				Type: schema.TypeString, Required: true, Description: "ServiceName",
			},
			"service_port": {
				Type: schema.TypeString, Required: true, Description: "ServicePort",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}
func resourceGslbZoneServiceGeoLocationCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbZoneServiceGeoLocationCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbZoneServiceGeoLocation(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGslbZoneServiceGeoLocationRead(ctx, d, meta)
	}
	return diags
}

func resourceGslbZoneServiceGeoLocationUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbZoneServiceGeoLocationUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbZoneServiceGeoLocation(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGslbZoneServiceGeoLocationRead(ctx, d, meta)
	}
	return diags
}
func resourceGslbZoneServiceGeoLocationDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbZoneServiceGeoLocationDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbZoneServiceGeoLocation(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceGslbZoneServiceGeoLocationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbZoneServiceGeoLocationRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbZoneServiceGeoLocation(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceGslbZoneServiceGeoLocationAlias(d []interface{}) []edpt.GslbZoneServiceGeoLocationAlias {

	count1 := len(d)
	ret := make([]edpt.GslbZoneServiceGeoLocationAlias, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbZoneServiceGeoLocationAlias
		oi.Alias = in["alias"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointGslbZoneServiceGeoLocation(d *schema.ResourceData) edpt.GslbZoneServiceGeoLocation {
	var ret edpt.GslbZoneServiceGeoLocation
	ret.Inst.Action = d.Get("action").(int)
	ret.Inst.ActionType = d.Get("action_type").(string)
	ret.Inst.Alias = getSliceGslbZoneServiceGeoLocationAlias(d.Get("alias").([]interface{}))
	ret.Inst.ForwardType = d.Get("forward_type").(string)
	ret.Inst.GeoName = d.Get("geo_name").(string)
	ret.Inst.Policy = d.Get("policy").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.ServiceName = d.Get("service_name").(string)
	ret.Inst.ServicePort = d.Get("service_port").(string)
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
