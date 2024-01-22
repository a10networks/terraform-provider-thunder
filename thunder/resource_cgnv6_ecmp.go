package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6Ecmp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_ecmp`: Configure ECMP hashing method for CGN\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6EcmpCreate,
		UpdateContext: resourceCgnv6EcmpUpdate,
		ReadContext:   resourceCgnv6EcmpRead,
		DeleteContext: resourceCgnv6EcmpDelete,

		Schema: map[string]*schema.Schema{
			"hashing_type": {
				Type: schema.TypeString, Optional: true, Description: "'4-tuple-hash': Hash on Src IP , Src Port, Dest IP and Dest Port;",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceCgnv6EcmpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6EcmpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Ecmp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6EcmpRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6EcmpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6EcmpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Ecmp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6EcmpRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6EcmpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6EcmpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Ecmp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6EcmpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6EcmpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Ecmp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointCgnv6Ecmp(d *schema.ResourceData) edpt.Cgnv6Ecmp {
	var ret edpt.Cgnv6Ecmp
	ret.Inst.HashingType = d.Get("hashing_type").(string)
	//omit uuid
	return ret
}
