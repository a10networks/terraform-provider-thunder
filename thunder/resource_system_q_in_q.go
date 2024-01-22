package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemQInQ() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_q_in_q`: 802.1QinQ system settings\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemQInQCreate,
		UpdateContext: resourceSystemQInQUpdate,
		ReadContext:   resourceSystemQInQRead,
		DeleteContext: resourceSystemQInQDelete,

		Schema: map[string]*schema.Schema{
			"enable_all_ports": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable 802.1QinQ on all physical ports",
			},
			"inner_tpid": {
				Type: schema.TypeString, Optional: true, Description: "TPID for inner VLAN (Inner TPID, 16 bit hex value, default is 8100)",
			},
			"outer_tpid": {
				Type: schema.TypeString, Optional: true, Description: "TPID for outer VLAN (Outer TPID, 16 bit hex value, default is 8100)",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSystemQInQCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemQInQCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemQInQ(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemQInQRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemQInQUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemQInQUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemQInQ(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemQInQRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemQInQDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemQInQDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemQInQ(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemQInQRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemQInQRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemQInQ(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSystemQInQ(d *schema.ResourceData) edpt.SystemQInQ {
	var ret edpt.SystemQInQ
	ret.Inst.EnableAllPorts = d.Get("enable_all_ports").(int)
	ret.Inst.InnerTpid = d.Get("inner_tpid").(string)
	ret.Inst.OuterTpid = d.Get("outer_tpid").(string)
	//omit uuid
	return ret
}
