package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGslbSiteSlbDevVipServerVipServerV6() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_gslb_site_slb_dev_vip_server_vip_server_v6`: Specify a VIP for the SLB device\n\n__PLACEHOLDER__",
		CreateContext: resourceGslbSiteSlbDevVipServerVipServerV6Create,
		UpdateContext: resourceGslbSiteSlbDevVipServerVipServerV6Update,
		ReadContext:   resourceGslbSiteSlbDevVipServerVipServerV6Read,
		DeleteContext: resourceGslbSiteSlbDevVipServerVipServerV6Delete,

		Schema: map[string]*schema.Schema{
			"ipv6": {
				Type: schema.TypeString, Required: true, Description: "Specify IP address (IPv6 address)",
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'dev_vip_hits': Number of times the service-ip was selected; 'dev_vip_recent': Recent hits;",
						},
					},
				},
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"site_name": {
				Type: schema.TypeString, Required: true, Description: "SiteName",
			},
			"device_name": {
				Type: schema.TypeString, Required: true, Description: "DeviceName",
			},
		},
	}
}
func resourceGslbSiteSlbDevVipServerVipServerV6Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbSiteSlbDevVipServerVipServerV6Create()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbSiteSlbDevVipServerVipServerV6(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGslbSiteSlbDevVipServerVipServerV6Read(ctx, d, meta)
	}
	return diags
}

func resourceGslbSiteSlbDevVipServerVipServerV6Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbSiteSlbDevVipServerVipServerV6Update()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbSiteSlbDevVipServerVipServerV6(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGslbSiteSlbDevVipServerVipServerV6Read(ctx, d, meta)
	}
	return diags
}
func resourceGslbSiteSlbDevVipServerVipServerV6Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbSiteSlbDevVipServerVipServerV6Delete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbSiteSlbDevVipServerVipServerV6(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceGslbSiteSlbDevVipServerVipServerV6Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbSiteSlbDevVipServerVipServerV6Read()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbSiteSlbDevVipServerVipServerV6(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceGslbSiteSlbDevVipServerVipServerV6SamplingEnable(d []interface{}) []edpt.GslbSiteSlbDevVipServerVipServerV6SamplingEnable {

	count1 := len(d)
	ret := make([]edpt.GslbSiteSlbDevVipServerVipServerV6SamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbSiteSlbDevVipServerVipServerV6SamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointGslbSiteSlbDevVipServerVipServerV6(d *schema.ResourceData) edpt.GslbSiteSlbDevVipServerVipServerV6 {
	var ret edpt.GslbSiteSlbDevVipServerVipServerV6
	ret.Inst.Ipv6 = d.Get("ipv6").(string)
	ret.Inst.SamplingEnable = getSliceGslbSiteSlbDevVipServerVipServerV6SamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	ret.Inst.SiteName = d.Get("site_name").(string)
	ret.Inst.DeviceName = d.Get("device_name").(string)
	return ret
}
