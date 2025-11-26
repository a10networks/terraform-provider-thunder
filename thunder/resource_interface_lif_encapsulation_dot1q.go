package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterfaceLifEncapsulationDot1q() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_interface_lif_encapsulation_dot1q`: dot1q encapsulation type\n\n__PLACEHOLDER__",
		CreateContext: resourceInterfaceLifEncapsulationDot1qCreate,
		UpdateContext: resourceInterfaceLifEncapsulationDot1qUpdate,
		ReadContext:   resourceInterfaceLifEncapsulationDot1qRead,
		DeleteContext: resourceInterfaceLifEncapsulationDot1qDelete,

		Schema: map[string]*schema.Schema{
			"ethernet": {
				Type: schema.TypeInt, Optional: true, Description: "Ethernet Interface (Ethernet port number)",
			},
			"tag": {
				Type: schema.TypeInt, Required: true, Description: "Tag ID",
			},
			"trunk": {
				Type: schema.TypeInt, Optional: true, Description: "Trunk Interface (Trunk number)",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"ifname": {
				Type: schema.TypeString, Required: true, Description: "Ifname",
			},
		},
	}
}
func resourceInterfaceLifEncapsulationDot1qCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceLifEncapsulationDot1qCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceLifEncapsulationDot1q(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceLifEncapsulationDot1qRead(ctx, d, meta)
	}
	return diags
}

func resourceInterfaceLifEncapsulationDot1qUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceLifEncapsulationDot1qUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceLifEncapsulationDot1q(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceLifEncapsulationDot1qRead(ctx, d, meta)
	}
	return diags
}
func resourceInterfaceLifEncapsulationDot1qDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceLifEncapsulationDot1qDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceLifEncapsulationDot1q(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceInterfaceLifEncapsulationDot1qRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceLifEncapsulationDot1qRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceLifEncapsulationDot1q(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointInterfaceLifEncapsulationDot1q(d *schema.ResourceData) edpt.InterfaceLifEncapsulationDot1q {
	var ret edpt.InterfaceLifEncapsulationDot1q
	ret.Inst.Ethernet = d.Get("ethernet").(int)
	ret.Inst.Tag = d.Get("tag").(int)
	ret.Inst.Trunk = d.Get("trunk").(int)
	//omit uuid
	ret.Inst.Ifname = d.Get("ifname").(string)
	return ret
}
