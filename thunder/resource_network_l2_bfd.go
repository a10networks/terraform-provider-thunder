package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetworkL2Bfd() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_network_l2_bfd`: Configure Layer-2 BFD global configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceNetworkL2BfdCreate,
		UpdateContext: resourceNetworkL2BfdUpdate,
		ReadContext:   resourceNetworkL2BfdRead,
		DeleteContext: resourceNetworkL2BfdDelete,

		Schema: map[string]*schema.Schema{
			"ether_type": {
				Type: schema.TypeString, Optional: true, Description: "Ethernet paylaod type for L2BFD packets, help-val 16 bit hex value, default is hex 88B6",
			},
			"multiplier": {
				Type: schema.TypeInt, Optional: true, Default: 4, Description: "Multiplier value used to compute holddown (value used to multiply the interval (default: 4))",
			},
			"rx_interval": {
				Type: schema.TypeInt, Optional: true, Default: 800, Description: "Minimum receive interval capability (Milliseconds (default: 800))",
			},
			"tx_interval": {
				Type: schema.TypeInt, Optional: true, Default: 800, Description: "Transmit interval between BFD packets (Milliseconds (default: 800))",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceNetworkL2BfdCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkL2BfdCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkL2Bfd(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceNetworkL2BfdRead(ctx, d, meta)
	}
	return diags
}

func resourceNetworkL2BfdUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkL2BfdUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkL2Bfd(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceNetworkL2BfdRead(ctx, d, meta)
	}
	return diags
}
func resourceNetworkL2BfdDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkL2BfdDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkL2Bfd(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceNetworkL2BfdRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkL2BfdRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkL2Bfd(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointNetworkL2Bfd(d *schema.ResourceData) edpt.NetworkL2Bfd {
	var ret edpt.NetworkL2Bfd
	ret.Inst.EtherType = d.Get("ether_type").(string)
	ret.Inst.Multiplier = d.Get("multiplier").(int)
	ret.Inst.RxInterval = d.Get("rx_interval").(int)
	ret.Inst.TxInterval = d.Get("tx_interval").(int)
	//omit uuid
	return ret
}
