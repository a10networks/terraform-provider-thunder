package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceBgp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_bgp`: Border Gateway Protocol (BGP)\n\n__PLACEHOLDER__",
		CreateContext: resourceBgpCreate,
		UpdateContext: resourceBgpUpdate,
		ReadContext:   resourceBgpRead,
		DeleteContext: resourceBgpDelete,

		Schema: map[string]*schema.Schema{
			"asdot": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable ASDot notation",
			},
			"asdot_plus": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable ASDot+ notation",
			},
			"disable_advertisement": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"on_boot": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Upon boot-up",
						},
					},
				},
			},
			"extended_asn_cap": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable the router to send 4-octet ASN capabilities",
			},
			"nexthop_trigger": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable the nexthop tracking functionality",
						},
						"delay": {
							Type: schema.TypeInt, Optional: true, Description: "Configure nexthop trigger delay time interval (Nexthop Trigger Delay time interval between 1 and 100)",
						},
					},
				},
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceBgpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceBgpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointBgp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceBgpRead(ctx, d, meta)
	}
	return diags
}

func resourceBgpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceBgpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointBgp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceBgpRead(ctx, d, meta)
	}
	return diags
}
func resourceBgpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceBgpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointBgp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceBgpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceBgpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointBgp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectBgpDisableAdvertisement(d []interface{}) edpt.BgpDisableAdvertisement {

	count1 := len(d)
	var ret edpt.BgpDisableAdvertisement
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.OnBoot = in["on_boot"].(int)
	}
	return ret
}

func getObjectBgpNexthopTrigger(d []interface{}) edpt.BgpNexthopTrigger {

	count1 := len(d)
	var ret edpt.BgpNexthopTrigger
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Enable = in["enable"].(int)
		ret.Delay = in["delay"].(int)
	}
	return ret
}

func dataToEndpointBgp(d *schema.ResourceData) edpt.Bgp {
	var ret edpt.Bgp
	ret.Inst.Asdot = d.Get("asdot").(int)
	ret.Inst.AsdotPlus = d.Get("asdot_plus").(int)
	ret.Inst.DisableAdvertisement = getObjectBgpDisableAdvertisement(d.Get("disable_advertisement").([]interface{}))
	ret.Inst.ExtendedAsnCap = d.Get("extended_asn_cap").(int)
	ret.Inst.NexthopTrigger = getObjectBgpNexthopTrigger(d.Get("nexthop_trigger").([]interface{}))
	//omit uuid
	return ret
}
