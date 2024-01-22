package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGslbSiteSlbDevVipServerVipServerV4() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_gslb_site_slb_dev_vip_server_vip_server_v4`: Specify a VIP for the SLB device\n\n__PLACEHOLDER__",
		CreateContext: resourceGslbSiteSlbDevVipServerVipServerV4Create,
		UpdateContext: resourceGslbSiteSlbDevVipServerVipServerV4Update,
		ReadContext:   resourceGslbSiteSlbDevVipServerVipServerV4Read,
		DeleteContext: resourceGslbSiteSlbDevVipServerVipServerV4Delete,

		Schema: map[string]*schema.Schema{
			"ipv4": {
				Type: schema.TypeString, Required: true, Description: "Specify IP address",
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
func resourceGslbSiteSlbDevVipServerVipServerV4Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbSiteSlbDevVipServerVipServerV4Create()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbSiteSlbDevVipServerVipServerV4(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGslbSiteSlbDevVipServerVipServerV4Read(ctx, d, meta)
	}
	return diags
}

func resourceGslbSiteSlbDevVipServerVipServerV4Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbSiteSlbDevVipServerVipServerV4Update()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbSiteSlbDevVipServerVipServerV4(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGslbSiteSlbDevVipServerVipServerV4Read(ctx, d, meta)
	}
	return diags
}
func resourceGslbSiteSlbDevVipServerVipServerV4Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbSiteSlbDevVipServerVipServerV4Delete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbSiteSlbDevVipServerVipServerV4(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceGslbSiteSlbDevVipServerVipServerV4Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbSiteSlbDevVipServerVipServerV4Read()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbSiteSlbDevVipServerVipServerV4(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceGslbSiteSlbDevVipServerVipServerV4SamplingEnable(d []interface{}) []edpt.GslbSiteSlbDevVipServerVipServerV4SamplingEnable {

	count1 := len(d)
	ret := make([]edpt.GslbSiteSlbDevVipServerVipServerV4SamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbSiteSlbDevVipServerVipServerV4SamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointGslbSiteSlbDevVipServerVipServerV4(d *schema.ResourceData) edpt.GslbSiteSlbDevVipServerVipServerV4 {
	var ret edpt.GslbSiteSlbDevVipServerVipServerV4
	ret.Inst.Ipv4 = d.Get("ipv4").(string)
	ret.Inst.SamplingEnable = getSliceGslbSiteSlbDevVipServerVipServerV4SamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	ret.Inst.SiteName = d.Get("site_name").(string)
	ret.Inst.DeviceName = d.Get("device_name").(string)
	return ret
}
