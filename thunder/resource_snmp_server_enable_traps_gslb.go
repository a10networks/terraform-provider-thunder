package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSnmpServerEnableTrapsGslb() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_snmp_server_enable_traps_gslb`: Enable GSLB group traps\n\n__PLACEHOLDER__",
		CreateContext: resourceSnmpServerEnableTrapsGslbCreate,
		UpdateContext: resourceSnmpServerEnableTrapsGslbUpdate,
		ReadContext:   resourceSnmpServerEnableTrapsGslbRead,
		DeleteContext: resourceSnmpServerEnableTrapsGslbDelete,

		Schema: map[string]*schema.Schema{
			"all": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable all GSLB traps",
			},
			"group": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable GSLB group related traps",
			},
			"service_ip": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable GSLB service-ip related traps",
			},
			"site": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable GSLB site related traps",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"zone": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable GSLB zone related traps",
			},
		},
	}
}
func resourceSnmpServerEnableTrapsGslbCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerEnableTrapsGslbCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerEnableTrapsGslb(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSnmpServerEnableTrapsGslbRead(ctx, d, meta)
	}
	return diags
}

func resourceSnmpServerEnableTrapsGslbUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerEnableTrapsGslbUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerEnableTrapsGslb(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSnmpServerEnableTrapsGslbRead(ctx, d, meta)
	}
	return diags
}
func resourceSnmpServerEnableTrapsGslbDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerEnableTrapsGslbDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerEnableTrapsGslb(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSnmpServerEnableTrapsGslbRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerEnableTrapsGslbRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerEnableTrapsGslb(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSnmpServerEnableTrapsGslb(d *schema.ResourceData) edpt.SnmpServerEnableTrapsGslb {
	var ret edpt.SnmpServerEnableTrapsGslb
	ret.Inst.All = d.Get("all").(int)
	ret.Inst.Group = d.Get("group").(int)
	ret.Inst.ServiceIp = d.Get("service_ip").(int)
	ret.Inst.Site = d.Get("site").(int)
	//omit uuid
	ret.Inst.Zone = d.Get("zone").(int)
	return ret
}
