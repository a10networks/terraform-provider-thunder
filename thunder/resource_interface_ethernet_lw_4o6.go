package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterfaceEthernetLw4o6() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_interface_ethernet_lw_4o6`: Configure LW-4over6 interface\n\n__PLACEHOLDER__",
		CreateContext: resourceInterfaceEthernetLw4o6Create,
		UpdateContext: resourceInterfaceEthernetLw4o6Update,
		ReadContext:   resourceInterfaceEthernetLw4o6Read,
		DeleteContext: resourceInterfaceEthernetLw4o6Delete,

		Schema: map[string]*schema.Schema{
			"inside": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure LW-4over6 outside interface",
			},
			"outside": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure LW-4over6 inside interface",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"ifnum": {
				Type: schema.TypeString, Required: true, Description: "Ifnum",
			},
		},
	}
}
func resourceInterfaceEthernetLw4o6Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceEthernetLw4o6Create()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceEthernetLw4o6(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceEthernetLw4o6Read(ctx, d, meta)
	}
	return diags
}

func resourceInterfaceEthernetLw4o6Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceEthernetLw4o6Update()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceEthernetLw4o6(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceEthernetLw4o6Read(ctx, d, meta)
	}
	return diags
}
func resourceInterfaceEthernetLw4o6Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceEthernetLw4o6Delete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceEthernetLw4o6(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceInterfaceEthernetLw4o6Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceEthernetLw4o6Read()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceEthernetLw4o6(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointInterfaceEthernetLw4o6(d *schema.ResourceData) edpt.InterfaceEthernetLw4o6 {
	var ret edpt.InterfaceEthernetLw4o6
	ret.Inst.Inside = d.Get("inside").(int)
	ret.Inst.Outside = d.Get("outside").(int)
	//omit uuid
	ret.Inst.Ifnum = d.Get("ifnum").(string)
	return ret
}
