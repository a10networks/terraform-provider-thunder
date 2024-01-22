package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSnmpServerEnableTrapsRoutingOspf() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_snmp_server_enable_traps_routing_ospf`: Enable OSPFv2 traps\n\n__PLACEHOLDER__",
		CreateContext: resourceSnmpServerEnableTrapsRoutingOspfCreate,
		UpdateContext: resourceSnmpServerEnableTrapsRoutingOspfUpdate,
		ReadContext:   resourceSnmpServerEnableTrapsRoutingOspfRead,
		DeleteContext: resourceSnmpServerEnableTrapsRoutingOspfDelete,

		Schema: map[string]*schema.Schema{
			"ospfifauthfailure": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable ospfIfAuthFailure traps",
			},
			"ospfifconfigerror": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable ospfIfConfigError traps",
			},
			"ospfifrxbadpacket": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable ospfIfRxBadPacket traps",
			},
			"ospfifstatechange": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable ospfIfStateChange traps",
			},
			"ospflsdbapproachingoverflow": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable ospfLsdbApproachingOverflow traps",
			},
			"ospflsdboverflow": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable ospfLsdbOverflow traps",
			},
			"ospfmaxagelsa": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable ospfMaxAgeLsa traps",
			},
			"ospfnbrstatechange": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable ospfNbrStateChange traps",
			},
			"ospforiginatelsa": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable ospfOriginateLsa traps",
			},
			"ospftxretransmit": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable ospfTxRetransmit traps",
			},
			"ospfvirtifauthfailure": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable ospfVirtIfAuthFailure traps",
			},
			"ospfvirtifconfigerror": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable ospfVirtIfConfigError traps",
			},
			"ospfvirtifrxbadpacket": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable ospfVirtIfRxBadPacket traps",
			},
			"ospfvirtifstatechange": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable ospfVirtIfStateChange traps",
			},
			"ospfvirtiftxretransmit": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable ospfVirtIfTxRetransmit traps",
			},
			"ospfvirtnbrstatechange": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable ospfVirtNbrStateChange traps",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSnmpServerEnableTrapsRoutingOspfCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerEnableTrapsRoutingOspfCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerEnableTrapsRoutingOspf(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSnmpServerEnableTrapsRoutingOspfRead(ctx, d, meta)
	}
	return diags
}

func resourceSnmpServerEnableTrapsRoutingOspfUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerEnableTrapsRoutingOspfUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerEnableTrapsRoutingOspf(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSnmpServerEnableTrapsRoutingOspfRead(ctx, d, meta)
	}
	return diags
}
func resourceSnmpServerEnableTrapsRoutingOspfDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerEnableTrapsRoutingOspfDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerEnableTrapsRoutingOspf(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSnmpServerEnableTrapsRoutingOspfRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerEnableTrapsRoutingOspfRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerEnableTrapsRoutingOspf(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSnmpServerEnableTrapsRoutingOspf(d *schema.ResourceData) edpt.SnmpServerEnableTrapsRoutingOspf {
	var ret edpt.SnmpServerEnableTrapsRoutingOspf
	ret.Inst.Ospfifauthfailure = d.Get("ospfifauthfailure").(int)
	ret.Inst.Ospfifconfigerror = d.Get("ospfifconfigerror").(int)
	ret.Inst.Ospfifrxbadpacket = d.Get("ospfifrxbadpacket").(int)
	ret.Inst.Ospfifstatechange = d.Get("ospfifstatechange").(int)
	ret.Inst.Ospflsdbapproachingoverflow = d.Get("ospflsdbapproachingoverflow").(int)
	ret.Inst.Ospflsdboverflow = d.Get("ospflsdboverflow").(int)
	ret.Inst.Ospfmaxagelsa = d.Get("ospfmaxagelsa").(int)
	ret.Inst.Ospfnbrstatechange = d.Get("ospfnbrstatechange").(int)
	ret.Inst.Ospforiginatelsa = d.Get("ospforiginatelsa").(int)
	ret.Inst.Ospftxretransmit = d.Get("ospftxretransmit").(int)
	ret.Inst.Ospfvirtifauthfailure = d.Get("ospfvirtifauthfailure").(int)
	ret.Inst.Ospfvirtifconfigerror = d.Get("ospfvirtifconfigerror").(int)
	ret.Inst.Ospfvirtifrxbadpacket = d.Get("ospfvirtifrxbadpacket").(int)
	ret.Inst.Ospfvirtifstatechange = d.Get("ospfvirtifstatechange").(int)
	ret.Inst.Ospfvirtiftxretransmit = d.Get("ospfvirtiftxretransmit").(int)
	ret.Inst.Ospfvirtnbrstatechange = d.Get("ospfvirtnbrstatechange").(int)
	//omit uuid
	return ret
}
