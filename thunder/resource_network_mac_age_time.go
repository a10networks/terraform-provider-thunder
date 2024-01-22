package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetworkMacAgeTime() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_network_mac_age_time`: Set Aging period for all MAC Interfaces\n\n__PLACEHOLDER__",
		CreateContext: resourceNetworkMacAgeTimeCreate,
		UpdateContext: resourceNetworkMacAgeTimeUpdate,
		ReadContext:   resourceNetworkMacAgeTimeRead,
		DeleteContext: resourceNetworkMacAgeTimeDelete,

		Schema: map[string]*schema.Schema{
			"aging_time": {
				Type: schema.TypeInt, Optional: true, Default: 300, Description: "Set aging period in seconds for all MAC interfaces (default 300 seconds)",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceNetworkMacAgeTimeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkMacAgeTimeCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkMacAgeTime(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceNetworkMacAgeTimeRead(ctx, d, meta)
	}
	return diags
}

func resourceNetworkMacAgeTimeUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkMacAgeTimeUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkMacAgeTime(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceNetworkMacAgeTimeRead(ctx, d, meta)
	}
	return diags
}
func resourceNetworkMacAgeTimeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkMacAgeTimeDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkMacAgeTime(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceNetworkMacAgeTimeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkMacAgeTimeRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkMacAgeTime(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointNetworkMacAgeTime(d *schema.ResourceData) edpt.NetworkMacAgeTime {
	var ret edpt.NetworkMacAgeTime
	ret.Inst.AgingTime = d.Get("aging_time").(int)
	//omit uuid
	return ret
}
