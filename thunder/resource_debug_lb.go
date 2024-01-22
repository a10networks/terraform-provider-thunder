package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDebugLb() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_debug_lb`: Load Balancing debug switch parameter\n\n__PLACEHOLDER__",
		CreateContext: resourceDebugLbCreate,
		UpdateContext: resourceDebugLbUpdate,
		ReadContext:   resourceDebugLbRead,
		DeleteContext: resourceDebugLbDelete,

		Schema: map[string]*schema.Schema{
			"all": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Load Balancing debug all switch",
			},
			"cfg": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Load Balancing configuration debug switch",
			},
			"clb": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Cache Load Balancing debug switch",
			},
			"flow": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Packet flow debug switch",
			},
			"fwlb": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Firewall Load Balancing debug switch",
			},
			"llb": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Link Load Balancing debug switch",
			},
			"slb": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Server Load Balancing debug switch",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceDebugLbCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugLbCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugLb(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugLbRead(ctx, d, meta)
	}
	return diags
}

func resourceDebugLbUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugLbUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugLb(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugLbRead(ctx, d, meta)
	}
	return diags
}
func resourceDebugLbDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugLbDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugLb(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDebugLbRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugLbRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugLb(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDebugLb(d *schema.ResourceData) edpt.DebugLb {
	var ret edpt.DebugLb
	ret.Inst.All = d.Get("all").(int)
	ret.Inst.Cfg = d.Get("cfg").(int)
	ret.Inst.Clb = d.Get("clb").(int)
	ret.Inst.Flow = d.Get("flow").(int)
	ret.Inst.Fwlb = d.Get("fwlb").(int)
	ret.Inst.Llb = d.Get("llb").(int)
	ret.Inst.Slb = d.Get("slb").(int)
	//omit uuid
	return ret
}
