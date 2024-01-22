package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6LsnPortOverloadingGlobal() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_lsn_port_overloading_global`: Configure Port-Overloading Behavior\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6LsnPortOverloadingGlobalCreate,
		UpdateContext: resourceCgnv6LsnPortOverloadingGlobalUpdate,
		ReadContext:   resourceCgnv6LsnPortOverloadingGlobalRead,
		DeleteContext: resourceCgnv6LsnPortOverloadingGlobalDelete,

		Schema: map[string]*schema.Schema{
			"allow_different_user": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Allow different users to overload the same port (default: disabled)",
			},
			"unique": {
				Type: schema.TypeString, Optional: true, Default: "destination-address-and-port", Description: "'destination-address': Allow overloading when the destination addresses is unique; 'destination-address-and-port': Allow overloading when the destination address and port 2-tuple is unique (default);",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceCgnv6LsnPortOverloadingGlobalCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnPortOverloadingGlobalCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnPortOverloadingGlobal(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6LsnPortOverloadingGlobalRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6LsnPortOverloadingGlobalUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnPortOverloadingGlobalUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnPortOverloadingGlobal(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6LsnPortOverloadingGlobalRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6LsnPortOverloadingGlobalDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnPortOverloadingGlobalDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnPortOverloadingGlobal(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6LsnPortOverloadingGlobalRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnPortOverloadingGlobalRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnPortOverloadingGlobal(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointCgnv6LsnPortOverloadingGlobal(d *schema.ResourceData) edpt.Cgnv6LsnPortOverloadingGlobal {
	var ret edpt.Cgnv6LsnPortOverloadingGlobal
	ret.Inst.AllowDifferentUser = d.Get("allow_different_user").(int)
	ret.Inst.Unique = d.Get("unique").(string)
	//omit uuid
	return ret
}
