package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceHealthMonitorMethodSnmp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_health_monitor_method_snmp`: SNMP type\n\n__PLACEHOLDER__",
		CreateContext: resourceHealthMonitorMethodSnmpCreate,
		UpdateContext: resourceHealthMonitorMethodSnmpUpdate,
		ReadContext:   resourceHealthMonitorMethodSnmpRead,
		DeleteContext: resourceHealthMonitorMethodSnmpDelete,

		Schema: map[string]*schema.Schema{
			"community": {
				Type: schema.TypeString, Optional: true, Default: "public", Description: "Specify SNMP community, default is \"public\" (Community String)",
			},
			"oid": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"mib": {
							Type: schema.TypeString, Optional: true, Description: "'sysDescr': The MIB-2 OID of system description, 1.1.0; 'sysUpTime': The MIB-2 OID of system up time, 1.3.0; 'sysName': The MIB-2 OID of system nume, 1.5.0;",
						},
						"asn": {
							Type: schema.TypeString, Optional: true, Description: "Specify the format in ASN.1 style",
						},
					},
				},
			},
			"operation": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"oper_type": {
							Type: schema.TypeString, Optional: true, Description: "'getnext': Get-Next-Request command; 'get': Get-Request command;",
						},
					},
				},
			},
			"snmp": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "SNMP type",
			},
			"snmp_port": {
				Type: schema.TypeInt, Optional: true, Default: 161, Description: "Specify SNMP port, default is 161 (Port Number)",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}
func resourceHealthMonitorMethodSnmpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthMonitorMethodSnmpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthMonitorMethodSnmp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceHealthMonitorMethodSnmpRead(ctx, d, meta)
	}
	return diags
}

func resourceHealthMonitorMethodSnmpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthMonitorMethodSnmpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthMonitorMethodSnmp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceHealthMonitorMethodSnmpRead(ctx, d, meta)
	}
	return diags
}
func resourceHealthMonitorMethodSnmpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthMonitorMethodSnmpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthMonitorMethodSnmp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceHealthMonitorMethodSnmpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthMonitorMethodSnmpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthMonitorMethodSnmp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectHealthMonitorMethodSnmpOid(d []interface{}) edpt.HealthMonitorMethodSnmpOid {

	count1 := len(d)
	var ret edpt.HealthMonitorMethodSnmpOid
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Mib = in["mib"].(string)
		ret.Asn = in["asn"].(string)
	}
	return ret
}

func getObjectHealthMonitorMethodSnmpOperation(d []interface{}) edpt.HealthMonitorMethodSnmpOperation {

	count1 := len(d)
	var ret edpt.HealthMonitorMethodSnmpOperation
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.OperType = in["oper_type"].(string)
	}
	return ret
}

func dataToEndpointHealthMonitorMethodSnmp(d *schema.ResourceData) edpt.HealthMonitorMethodSnmp {
	var ret edpt.HealthMonitorMethodSnmp
	ret.Inst.Community = d.Get("community").(string)
	ret.Inst.Oid = getObjectHealthMonitorMethodSnmpOid(d.Get("oid").([]interface{}))
	ret.Inst.Operation = getObjectHealthMonitorMethodSnmpOperation(d.Get("operation").([]interface{}))
	ret.Inst.Snmp = d.Get("snmp").(int)
	ret.Inst.SnmpPort = d.Get("snmp_port").(int)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
