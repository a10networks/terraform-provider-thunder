package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6NatHistogram() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_nat_histogram`: Configure CGNv6 NAT Histogram\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6NatHistogramCreate,
		UpdateContext: resourceCgnv6NatHistogramUpdate,
		ReadContext:   resourceCgnv6NatHistogramRead,
		DeleteContext: resourceCgnv6NatHistogramDelete,

		Schema: map[string]*schema.Schema{
			"bin_count": {
				Type: schema.TypeInt, Optional: true, Default: 50, Description: "Number of bins in the histogram (default: 50)",
			},
			"bin_skew": {
				Type: schema.TypeInt, Optional: true, Default: 75, Description: "Percentage of bins that represent the upper bound (default: 75)",
			},
			"data_skew": {
				Type: schema.TypeInt, Optional: true, Default: 25, Description: "Percentage of data that represents the upper bound (default: 25)",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceCgnv6NatHistogramCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6NatHistogramCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6NatHistogram(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6NatHistogramRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6NatHistogramUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6NatHistogramUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6NatHistogram(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6NatHistogramRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6NatHistogramDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6NatHistogramDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6NatHistogram(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6NatHistogramRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6NatHistogramRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6NatHistogram(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointCgnv6NatHistogram(d *schema.ResourceData) edpt.Cgnv6NatHistogram {
	var ret edpt.Cgnv6NatHistogram
	ret.Inst.BinCount = d.Get("bin_count").(int)
	ret.Inst.BinSkew = d.Get("bin_skew").(int)
	ret.Inst.DataSkew = d.Get("data_skew").(int)
	//omit uuid
	return ret
}
