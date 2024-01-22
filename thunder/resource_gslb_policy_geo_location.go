package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGslbPolicyGeoLocation() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_gslb_policy_geo_location`: Specify geo-location\n\n__PLACEHOLDER__",
		CreateContext: resourceGslbPolicyGeoLocationCreate,
		UpdateContext: resourceGslbPolicyGeoLocationUpdate,
		ReadContext:   resourceGslbPolicyGeoLocationRead,
		DeleteContext: resourceGslbPolicyGeoLocationDelete,

		Schema: map[string]*schema.Schema{
			"ip_multiple_fields": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip_sub": {
							Type: schema.TypeString, Optional: true, Description: "Specify IP information",
						},
						"ip_mask_sub": {
							Type: schema.TypeString, Optional: true, Description: "Specify IP/mask format (Specify IP address mask)",
						},
						"ip_addr2_sub": {
							Type: schema.TypeString, Optional: true, Description: "Specify IP address range",
						},
					},
				},
			},
			"ipv6_multiple_fields": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ipv6_sub": {
							Type: schema.TypeString, Optional: true, Description: "Specify IPv6 information",
						},
						"ipv6_mask_sub": {
							Type: schema.TypeInt, Optional: true, Description: "Specify IPv6/mask format (Specify IP address mask)",
						},
						"ipv6_addr2_sub": {
							Type: schema.TypeString, Optional: true, Description: "Specify IPv6 address range",
						},
					},
				},
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Specify geo-location name, section range is (1-15)",
			},
			"policy_name": {
				Type: schema.TypeString, Required: true, Description: "Specify policy-name",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceGslbPolicyGeoLocationCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbPolicyGeoLocationCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbPolicyGeoLocation(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGslbPolicyGeoLocationRead(ctx, d, meta)
	}
	return diags
}

func resourceGslbPolicyGeoLocationUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbPolicyGeoLocationUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbPolicyGeoLocation(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGslbPolicyGeoLocationRead(ctx, d, meta)
	}
	return diags
}
func resourceGslbPolicyGeoLocationDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbPolicyGeoLocationDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbPolicyGeoLocation(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceGslbPolicyGeoLocationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbPolicyGeoLocationRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbPolicyGeoLocation(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceGslbPolicyGeoLocationIpMultipleFields(d []interface{}) []edpt.GslbPolicyGeoLocationIpMultipleFields {

	count1 := len(d)
	ret := make([]edpt.GslbPolicyGeoLocationIpMultipleFields, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbPolicyGeoLocationIpMultipleFields
		oi.IpSub = in["ip_sub"].(string)
		oi.IpMaskSub = in["ip_mask_sub"].(string)
		oi.IpAddr2Sub = in["ip_addr2_sub"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceGslbPolicyGeoLocationIpv6MultipleFields(d []interface{}) []edpt.GslbPolicyGeoLocationIpv6MultipleFields {

	count1 := len(d)
	ret := make([]edpt.GslbPolicyGeoLocationIpv6MultipleFields, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbPolicyGeoLocationIpv6MultipleFields
		oi.Ipv6Sub = in["ipv6_sub"].(string)
		oi.Ipv6MaskSub = in["ipv6_mask_sub"].(int)
		oi.Ipv6Addr2Sub = in["ipv6_addr2_sub"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointGslbPolicyGeoLocation(d *schema.ResourceData) edpt.GslbPolicyGeoLocation {
	var ret edpt.GslbPolicyGeoLocation
	ret.Inst.IpMultipleFields = getSliceGslbPolicyGeoLocationIpMultipleFields(d.Get("ip_multiple_fields").([]interface{}))
	ret.Inst.Ipv6MultipleFields = getSliceGslbPolicyGeoLocationIpv6MultipleFields(d.Get("ipv6_multiple_fields").([]interface{}))
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.PolicyName = d.Get("policy_name").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
