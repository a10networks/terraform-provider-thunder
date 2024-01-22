package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpUnnumberedUseSourceIp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ip_unnumbered_use_source_ip`: Source IP address\n\n__PLACEHOLDER__",
		CreateContext: resourceIpUnnumberedUseSourceIpCreate,
		UpdateContext: resourceIpUnnumberedUseSourceIpUpdate,
		ReadContext:   resourceIpUnnumberedUseSourceIpRead,
		DeleteContext: resourceIpUnnumberedUseSourceIpDelete,

		Schema: map[string]*schema.Schema{
			"update_source_ip": {
				Type: schema.TypeString, Optional: true, Description: "IP address",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceIpUnnumberedUseSourceIpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpUnnumberedUseSourceIpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpUnnumberedUseSourceIp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpUnnumberedUseSourceIpRead(ctx, d, meta)
	}
	return diags
}

func resourceIpUnnumberedUseSourceIpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpUnnumberedUseSourceIpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpUnnumberedUseSourceIp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpUnnumberedUseSourceIpRead(ctx, d, meta)
	}
	return diags
}
func resourceIpUnnumberedUseSourceIpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpUnnumberedUseSourceIpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpUnnumberedUseSourceIp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceIpUnnumberedUseSourceIpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpUnnumberedUseSourceIpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpUnnumberedUseSourceIp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointIpUnnumberedUseSourceIp(d *schema.ResourceData) edpt.IpUnnumberedUseSourceIp {
	var ret edpt.IpUnnumberedUseSourceIp
	ret.Inst.UpdateSourceIp = d.Get("update_source_ip").(string)
	//omit uuid
	return ret
}
