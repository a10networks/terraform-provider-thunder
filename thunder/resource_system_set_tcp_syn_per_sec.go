package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemSetTcpSynPerSec() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_set_tcp_syn_per_sec`: Configure tcp syn per sec\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemSetTcpSynPerSecCreate,
		UpdateContext: resourceSystemSetTcpSynPerSecUpdate,
		ReadContext:   resourceSystemSetTcpSynPerSecRead,
		DeleteContext: resourceSystemSetTcpSynPerSecDelete,

		Schema: map[string]*schema.Schema{
			"tcp_syn_value": {
				Type: schema.TypeInt, Optional: true, Default: 70, Description: "Configure Tcp SYN's per sec, default 70",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSystemSetTcpSynPerSecCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemSetTcpSynPerSecCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemSetTcpSynPerSec(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemSetTcpSynPerSecRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemSetTcpSynPerSecUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemSetTcpSynPerSecUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemSetTcpSynPerSec(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemSetTcpSynPerSecRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemSetTcpSynPerSecDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemSetTcpSynPerSecDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemSetTcpSynPerSec(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemSetTcpSynPerSecRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemSetTcpSynPerSecRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemSetTcpSynPerSec(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSystemSetTcpSynPerSec(d *schema.ResourceData) edpt.SystemSetTcpSynPerSec {
	var ret edpt.SystemSetTcpSynPerSec
	ret.Inst.TcpSynValue = d.Get("tcp_syn_value").(int)
	//omit uuid
	return ret
}
