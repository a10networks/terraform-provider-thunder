package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemGeoLocationEntry() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_geo_location_entry`: Manually configure geo-location entry\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemGeoLocationEntryCreate,
		UpdateContext: resourceSystemGeoLocationEntryUpdate,
		ReadContext:   resourceSystemGeoLocationEntryRead,
		DeleteContext: resourceSystemGeoLocationEntryDelete,

		Schema: map[string]*schema.Schema{
			"geo_locn_multiple_addresses": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"first_ip_address": {
							Type: schema.TypeString, Optional: true, Description: "Specify IP information (Specify IP address)",
						},
						"geol_ipv4_mask": {
							Type: schema.TypeString, Optional: true, Description: "Specify IPv4 mask",
						},
						"ip_addr2": {
							Type: schema.TypeString, Optional: true, Description: "Specify IP address range",
						},
						"first_ipv6_address": {
							Type: schema.TypeString, Optional: true, Description: "Specify IPv6 address",
						},
						"geol_ipv6_mask": {
							Type: schema.TypeInt, Optional: true, Description: "Specify IPv6 mask",
						},
						"ipv6_addr2": {
							Type: schema.TypeString, Optional: true, Description: "Specify IPv6 address range",
						},
					},
				},
			},
			"geo_locn_obj_name": {
				Type: schema.TypeString, Required: true, Description: "Specify geo-location name, section range is (1-15)",
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
func resourceSystemGeoLocationEntryCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemGeoLocationEntryCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemGeoLocationEntry(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemGeoLocationEntryRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemGeoLocationEntryUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemGeoLocationEntryUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemGeoLocationEntry(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemGeoLocationEntryRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemGeoLocationEntryDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemGeoLocationEntryDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemGeoLocationEntry(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemGeoLocationEntryRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemGeoLocationEntryRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemGeoLocationEntry(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSystemGeoLocationEntryGeoLocnMultipleAddresses(d []interface{}) []edpt.SystemGeoLocationEntryGeoLocnMultipleAddresses {

	count1 := len(d)
	ret := make([]edpt.SystemGeoLocationEntryGeoLocnMultipleAddresses, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemGeoLocationEntryGeoLocnMultipleAddresses
		oi.FirstIpAddress = in["first_ip_address"].(string)
		oi.GeolIpv4Mask = in["geol_ipv4_mask"].(string)
		oi.IpAddr2 = in["ip_addr2"].(string)
		oi.FirstIpv6Address = in["first_ipv6_address"].(string)
		oi.GeolIpv6Mask = in["geol_ipv6_mask"].(int)
		oi.Ipv6Addr2 = in["ipv6_addr2"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSystemGeoLocationEntry(d *schema.ResourceData) edpt.SystemGeoLocationEntry {
	var ret edpt.SystemGeoLocationEntry
	ret.Inst.GeoLocnMultipleAddresses = getSliceSystemGeoLocationEntryGeoLocnMultipleAddresses(d.Get("geo_locn_multiple_addresses").([]interface{}))
	ret.Inst.GeoLocnObjName = d.Get("geo_locn_obj_name").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
