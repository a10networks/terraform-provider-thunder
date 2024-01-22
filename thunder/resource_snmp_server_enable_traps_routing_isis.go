package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSnmpServerEnableTrapsRoutingIsis() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_snmp_server_enable_traps_routing_isis`: Enable isis traps\n\n__PLACEHOLDER__",
		CreateContext: resourceSnmpServerEnableTrapsRoutingIsisCreate,
		UpdateContext: resourceSnmpServerEnableTrapsRoutingIsisUpdate,
		ReadContext:   resourceSnmpServerEnableTrapsRoutingIsisRead,
		DeleteContext: resourceSnmpServerEnableTrapsRoutingIsisDelete,

		Schema: map[string]*schema.Schema{
			"isisadjacencychange": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable isisAdjacencyChange traps",
			},
			"isisareamismatch": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable isisAreaMismatch traps",
			},
			"isisattempttoexceedmaxsequence": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable isisAttemptToExceedMaxSequence traps",
			},
			"isisauthenticationfailure": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable isisAuthenticationFailure traps",
			},
			"isisauthenticationtypefailure": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable isisAuthenticationTypeFailure traps",
			},
			"isiscorruptedlspdetected": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable isisCorruptedLSPDetected traps",
			},
			"isisdatabaseoverload": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable isisDatabaseOverload traps",
			},
			"isisidlenmismatch": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable isisIDLenMismatch traps",
			},
			"isislsperrordetected": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable isisLSPErrorDetected traps",
			},
			"isislsptoolargetopropagate": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable isisLSPTooLargeToPropagate traps",
			},
			"isismanualaddressdrops": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable isisManualAddressDrops traps",
			},
			"isismaxareaaddressesmismatch": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable isisMaxAreaAddressesMismatch traps",
			},
			"isisoriginatinglspbuffersizemismatch": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable isisOriginatingLSPBufferSizeMismatch traps",
			},
			"isisownlsppurge": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable isisOwnLSPPurge traps",
			},
			"isisprotocolssupportedmismatch": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable isisProtocolsSupportedMismatch traps",
			},
			"isisrejectedadjacency": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable isisRejectedAdjacency traps",
			},
			"isissequencenumberskip": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable isisSequenceNumberSkip traps",
			},
			"isisversionskew": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable isisVersionSkew traps",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSnmpServerEnableTrapsRoutingIsisCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerEnableTrapsRoutingIsisCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerEnableTrapsRoutingIsis(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSnmpServerEnableTrapsRoutingIsisRead(ctx, d, meta)
	}
	return diags
}

func resourceSnmpServerEnableTrapsRoutingIsisUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerEnableTrapsRoutingIsisUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerEnableTrapsRoutingIsis(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSnmpServerEnableTrapsRoutingIsisRead(ctx, d, meta)
	}
	return diags
}
func resourceSnmpServerEnableTrapsRoutingIsisDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerEnableTrapsRoutingIsisDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerEnableTrapsRoutingIsis(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSnmpServerEnableTrapsRoutingIsisRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerEnableTrapsRoutingIsisRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerEnableTrapsRoutingIsis(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSnmpServerEnableTrapsRoutingIsis(d *schema.ResourceData) edpt.SnmpServerEnableTrapsRoutingIsis {
	var ret edpt.SnmpServerEnableTrapsRoutingIsis
	ret.Inst.Isisadjacencychange = d.Get("isisadjacencychange").(int)
	ret.Inst.Isisareamismatch = d.Get("isisareamismatch").(int)
	ret.Inst.Isisattempttoexceedmaxsequence = d.Get("isisattempttoexceedmaxsequence").(int)
	ret.Inst.Isisauthenticationfailure = d.Get("isisauthenticationfailure").(int)
	ret.Inst.Isisauthenticationtypefailure = d.Get("isisauthenticationtypefailure").(int)
	ret.Inst.Isiscorruptedlspdetected = d.Get("isiscorruptedlspdetected").(int)
	ret.Inst.Isisdatabaseoverload = d.Get("isisdatabaseoverload").(int)
	ret.Inst.Isisidlenmismatch = d.Get("isisidlenmismatch").(int)
	ret.Inst.Isislsperrordetected = d.Get("isislsperrordetected").(int)
	ret.Inst.Isislsptoolargetopropagate = d.Get("isislsptoolargetopropagate").(int)
	ret.Inst.Isismanualaddressdrops = d.Get("isismanualaddressdrops").(int)
	ret.Inst.Isismaxareaaddressesmismatch = d.Get("isismaxareaaddressesmismatch").(int)
	ret.Inst.Isisoriginatinglspbuffersizemismatch = d.Get("isisoriginatinglspbuffersizemismatch").(int)
	ret.Inst.Isisownlsppurge = d.Get("isisownlsppurge").(int)
	ret.Inst.Isisprotocolssupportedmismatch = d.Get("isisprotocolssupportedmismatch").(int)
	ret.Inst.Isisrejectedadjacency = d.Get("isisrejectedadjacency").(int)
	ret.Inst.Isissequencenumberskip = d.Get("isissequencenumberskip").(int)
	ret.Inst.Isisversionskew = d.Get("isisversionskew").(int)
	//omit uuid
	return ret
}
