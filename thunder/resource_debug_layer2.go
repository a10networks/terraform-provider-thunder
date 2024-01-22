package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDebugLayer2() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_debug_layer2`: Layer2 module parameters\n\n__PLACEHOLDER__",
		CreateContext: resourceDebugLayer2Create,
		UpdateContext: resourceDebugLayer2Update,
		ReadContext:   resourceDebugLayer2Read,
		DeleteContext: resourceDebugLayer2Delete,

		Schema: map[string]*schema.Schema{
			"all": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Layer2 all debug switch",
			},
			"arp": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Layer2 debug arp switch",
			},
			"interface": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Layer2 debug interface switch",
			},
			"misc": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Layer2 debug misc switch",
			},
			"trace": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Layer2 debug trace switch",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"vlan": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Layer2 debug vlan switch",
			},
		},
	}
}
func resourceDebugLayer2Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugLayer2Create()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugLayer2(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugLayer2Read(ctx, d, meta)
	}
	return diags
}

func resourceDebugLayer2Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugLayer2Update()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugLayer2(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugLayer2Read(ctx, d, meta)
	}
	return diags
}
func resourceDebugLayer2Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugLayer2Delete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugLayer2(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDebugLayer2Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugLayer2Read()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugLayer2(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDebugLayer2(d *schema.ResourceData) edpt.DebugLayer2 {
	var ret edpt.DebugLayer2
	ret.Inst.All = d.Get("all").(int)
	ret.Inst.Arp = d.Get("arp").(int)
	ret.Inst.Interface = d.Get("interface").(int)
	ret.Inst.Misc = d.Get("misc").(int)
	ret.Inst.Trace = d.Get("trace").(int)
	//omit uuid
	ret.Inst.Vlan = d.Get("vlan").(int)
	return ret
}
