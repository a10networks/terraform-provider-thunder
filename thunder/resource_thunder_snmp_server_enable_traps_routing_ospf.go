package thunder

//Thunder resource SnmpServerEnableTrapsRoutingOspf

import (
	"context"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSnmpServerEnableTrapsRoutingOspf() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSnmpServerEnableTrapsRoutingOspfCreate,
		UpdateContext: resourceSnmpServerEnableTrapsRoutingOspfUpdate,
		ReadContext:   resourceSnmpServerEnableTrapsRoutingOspfRead,
		DeleteContext: resourceSnmpServerEnableTrapsRoutingOspfDelete,
		Schema: map[string]*schema.Schema{
			"ospf_lsdb_overflow": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"ospf_nbr_state_change": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"ospf_if_state_change": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"ospf_virt_nbr_state_change": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"ospf_lsdb_approaching_overflow": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"ospf_if_auth_failure": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"ospf_virt_if_auth_failure": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"ospf_virt_if_config_error": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"ospf_virt_if_rx_bad_packet": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"ospf_tx_retransmit": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"ospf_virt_if_state_change": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"ospf_if_config_error": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"ospf_max_age_lsa": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"ospf_if_rx_bad_packet": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"ospf_virt_if_tx_retransmit": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"ospf_originate_lsa": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceSnmpServerEnableTrapsRoutingOspfCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating SnmpServerEnableTrapsRoutingOspf (Inside resourceSnmpServerEnableTrapsRoutingOspfCreate) ")

		data := dataToSnmpServerEnableTrapsRoutingOspf(d)
		logger.Println("[INFO] received formatted data from method data to SnmpServerEnableTrapsRoutingOspf --")
		d.SetId("1")
		err := go_thunder.PostSnmpServerEnableTrapsRoutingOspf(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceSnmpServerEnableTrapsRoutingOspfRead(ctx, d, meta)

	}
	return diags
}

func resourceSnmpServerEnableTrapsRoutingOspfRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading SnmpServerEnableTrapsRoutingOspf (Inside resourceSnmpServerEnableTrapsRoutingOspfRead)")

	if client.Host != "" {
		logger.Println("[INFO] Fetching service Read")
		data, err := go_thunder.GetSnmpServerEnableTrapsRoutingOspf(client.Token, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		if data == nil {
			logger.Println("[INFO] No data found ")
			return nil
		}
		return diags
	}
	return nil
}

func resourceSnmpServerEnableTrapsRoutingOspfUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSnmpServerEnableTrapsRoutingOspfRead(ctx, d, meta)
}

func resourceSnmpServerEnableTrapsRoutingOspfDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSnmpServerEnableTrapsRoutingOspfRead(ctx, d, meta)
}
func dataToSnmpServerEnableTrapsRoutingOspf(d *schema.ResourceData) go_thunder.SnmpServerEnableTrapsRoutingOspf {
	var vc go_thunder.SnmpServerEnableTrapsRoutingOspf
	var c go_thunder.SnmpServerEnableTrapsRoutingOspfInstance
	c.OspfLsdbOverflow = d.Get("ospf_lsdb_overflow").(int)
	c.OspfNbrStateChange = d.Get("ospf_nbr_state_change").(int)
	c.OspfIfStateChange = d.Get("ospf_if_state_change").(int)
	c.OspfVirtNbrStateChange = d.Get("ospf_virt_nbr_state_change").(int)
	c.OspfLsdbApproachingOverflow = d.Get("ospf_lsdb_approaching_overflow").(int)
	c.OspfIfAuthFailure = d.Get("ospf_if_auth_failure").(int)
	c.OspfVirtIfAuthFailure = d.Get("ospf_virt_if_auth_failure").(int)
	c.OspfVirtIfConfigError = d.Get("ospf_virt_if_config_error").(int)
	c.OspfVirtIfRxBadPacket = d.Get("ospf_virt_if_rx_bad_packet").(int)
	c.OspfTxRetransmit = d.Get("ospf_tx_retransmit").(int)
	c.OspfVirtIfStateChange = d.Get("ospf_virt_if_state_change").(int)
	c.OspfIfConfigError = d.Get("ospf_if_config_error").(int)
	c.OspfMaxAgeLsa = d.Get("ospf_max_age_lsa").(int)
	c.OspfIfRxBadPacket = d.Get("ospf_if_rx_bad_packet").(int)
	c.OspfVirtIfTxRetransmit = d.Get("ospf_virt_if_tx_retransmit").(int)
	c.OspfOriginateLsa = d.Get("ospf_originate_lsa").(int)

	vc.OspfLsdbOverflow = c
	return vc
}
