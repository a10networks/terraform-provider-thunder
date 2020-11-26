package thunder

//Thunder resource SnmpServerEnableTrapsRoutingIsis

import (
	"github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"util"
)

func resourceSnmpServerEnableTrapsRoutingIsis() *schema.Resource {
	return &schema.Resource{
		Create: resourceSnmpServerEnableTrapsRoutingIsisCreate,
		Update: resourceSnmpServerEnableTrapsRoutingIsisUpdate,
		Read:   resourceSnmpServerEnableTrapsRoutingIsisRead,
		Delete: resourceSnmpServerEnableTrapsRoutingIsisDelete,
		Schema: map[string]*schema.Schema{
			"isis_authentication_failure": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"isis_protocols_supported_mismatch": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"isis_rejected_adjacency": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"isis_max_area_addresses_mismatch": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"isis_corrupted_lsp_detected": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"isis_originating_lsp_buffer_size_mismatch": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"isis_area_mismatch": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"isis_lsp_too_large_to_propagate": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"isis_own_lsp_purge": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"isis_sequence_number_skip": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"isis_database_overload": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"isis_attempt_to_exceed_max_sequence": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"isis_id_len_mismatch": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"isis_authentication_type_failure": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"isis_version_skew": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"isis_manual_address_drops": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"isis_adjacency_change": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceSnmpServerEnableTrapsRoutingIsisCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating SnmpServerEnableTrapsRoutingIsis (Inside resourceSnmpServerEnableTrapsRoutingIsisCreate) ")

		data := dataToSnmpServerEnableTrapsRoutingIsis(d)
		logger.Println("[INFO] received formatted data from method data to SnmpServerEnableTrapsRoutingIsis --")
		d.SetId("1")
		go_thunder.PostSnmpServerEnableTrapsRoutingIsis(client.Token, data, client.Host)

		return resourceSnmpServerEnableTrapsRoutingIsisRead(d, meta)

	}
	return nil
}

func resourceSnmpServerEnableTrapsRoutingIsisRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading SnmpServerEnableTrapsRoutingIsis (Inside resourceSnmpServerEnableTrapsRoutingIsisRead)")

	if client.Host != "" {
		logger.Println("[INFO] Fetching service Read")
		data, err := go_thunder.GetSnmpServerEnableTrapsRoutingIsis(client.Token, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found ")
			return nil
		}
		return err
	}
	return nil
}

func resourceSnmpServerEnableTrapsRoutingIsisUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceSnmpServerEnableTrapsRoutingIsisRead(d, meta)
}

func resourceSnmpServerEnableTrapsRoutingIsisDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceSnmpServerEnableTrapsRoutingIsisRead(d, meta)
}
func dataToSnmpServerEnableTrapsRoutingIsis(d *schema.ResourceData) go_thunder.SnmpServerEnableTrapsRoutingIsis {
	var vc go_thunder.SnmpServerEnableTrapsRoutingIsis
	var c go_thunder.SnmpServerEnableTrapsRoutingIsisInstance
	c.IsisAuthenticationFailure = d.Get("isis_authentication_failure").(int)
	c.IsisProtocolsSupportedMismatch = d.Get("isis_protocols_supported_mismatch").(int)
	c.IsisRejectedAdjacency = d.Get("isis_rejected_adjacency").(int)
	c.IsisMaxAreaAddressesMismatch = d.Get("isis_max_area_addresses_mismatch").(int)
	c.IsisCorruptedLSPDetected = d.Get("isis_corrupted_lsp_detected").(int)
	c.IsisOriginatingLSPBufferSizeMismatch = d.Get("isis_originating_lsp_buffer_size_mismatch").(int)
	c.IsisAreaMismatch = d.Get("isis_area_mismatch").(int)
	c.IsisLSPTooLargeToPropagate = d.Get("isis_lsp_too_large_to_propagate").(int)
	c.IsisOwnLSPPurge = d.Get("isis_own_lsp_purge").(int)
	c.IsisSequenceNumberSkip = d.Get("isis_sequence_number_skip").(int)
	c.IsisDatabaseOverload = d.Get("isis_database_overload").(int)
	c.IsisAttemptToExceedMaxSequence = d.Get("isis_attempt_to_exceed_max_sequence").(int)
	c.IsisIDLenMismatch = d.Get("isis_id_len_mismatch").(int)
	c.IsisAuthenticationTypeFailure = d.Get("isis_authentication_type_failure").(int)
	c.IsisVersionSkew = d.Get("isis_version_skew").(int)
	c.IsisManualAddressDrops = d.Get("isis_manual_address_drops").(int)
	c.IsisAdjacencyChange = d.Get("isis_adjacency_change").(int)

	vc.IsisAuthenticationFailure = c
	return vc
}
