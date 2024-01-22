package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetworkTwampResponderIp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_network_twamp_responder_ip`: Configure TWAMP responder\n\n__PLACEHOLDER__",
		CreateContext: resourceNetworkTwampResponderIpCreate,
		UpdateContext: resourceNetworkTwampResponderIpUpdate,
		ReadContext:   resourceNetworkTwampResponderIpRead,
		DeleteContext: resourceNetworkTwampResponderIpDelete,

		Schema: map[string]*schema.Schema{
			"acl_id": {
				Type: schema.TypeInt, Optional: true, Description: "ACL id",
			},
			"acl_name": {
				Type: schema.TypeString, Optional: true, Description: "Apply a named access list (Access List name)",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceNetworkTwampResponderIpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkTwampResponderIpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkTwampResponderIp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceNetworkTwampResponderIpRead(ctx, d, meta)
	}
	return diags
}

func resourceNetworkTwampResponderIpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkTwampResponderIpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkTwampResponderIp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceNetworkTwampResponderIpRead(ctx, d, meta)
	}
	return diags
}
func resourceNetworkTwampResponderIpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkTwampResponderIpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkTwampResponderIp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceNetworkTwampResponderIpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkTwampResponderIpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkTwampResponderIp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointNetworkTwampResponderIp(d *schema.ResourceData) edpt.NetworkTwampResponderIp {
	var ret edpt.NetworkTwampResponderIp
	ret.Inst.AclId = d.Get("acl_id").(int)
	ret.Inst.AclName = d.Get("acl_name").(string)
	//omit uuid
	return ret
}
