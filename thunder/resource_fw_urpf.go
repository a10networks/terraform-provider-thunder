package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFwUrpf() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_fw_urpf`: Unicast Reverse Path Forwarding (Default: loose)\n\n__PLACEHOLDER__",
		CreateContext: resourceFwUrpfCreate,
		UpdateContext: resourceFwUrpfUpdate,
		ReadContext:   resourceFwUrpfRead,
		DeleteContext: resourceFwUrpfDelete,

		Schema: map[string]*schema.Schema{
			"status": {
				Type: schema.TypeString, Optional: true, Default: "loose", Description: "'loose': Perform loose check; 'strict': Perform strict check; 'disable': Disable check;",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceFwUrpfCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwUrpfCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwUrpf(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFwUrpfRead(ctx, d, meta)
	}
	return diags
}

func resourceFwUrpfUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwUrpfUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwUrpf(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFwUrpfRead(ctx, d, meta)
	}
	return diags
}
func resourceFwUrpfDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwUrpfDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwUrpf(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFwUrpfRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwUrpfRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwUrpf(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointFwUrpf(d *schema.ResourceData) edpt.FwUrpf {
	var ret edpt.FwUrpf
	ret.Inst.Status = d.Get("status").(string)
	//omit uuid
	return ret
}
