package thunder

//Thunder resource FwGlobal

import (
	"fmt"
	"github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"strconv"
	"util"
)

func resourceFwGlobal() *schema.Resource {
	return &schema.Resource{
		Create: resourceFwGlobalCreate,
		Update: resourceFwGlobalUpdate,
		Read:   resourceFwGlobalRead,
		Delete: resourceFwGlobalDelete,
		Schema: map[string]*schema.Schema{
			"alg_processing": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"listen_on_port_timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"disable_ip_fw_sessions": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"disable_app_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"disable_application_protocol": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"disable_application_category": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"extended_matching": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"sampling_enable": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"respond_to_user_mac": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"permit_default_action": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"natip_ddos_protection": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceFwGlobalCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating FwGlobal (Inside resourceFwGlobalCreate) ")

		data := dataToFwGlobal(d)
		logger.Println("[INFO] received formatted data from method data to FwGlobal --")
		d.SetId(strconv.Itoa('1'))
		go_thunder.PostFwGlobal(client.Token, data, client.Host)

		return resourceFwGlobalRead(d, meta)

	}
	return nil
}

func resourceFwGlobalRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading FwGlobal (Inside resourceFwGlobalRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_thunder.GetFwGlobal(client.Token, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name)
			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceFwGlobalUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceFwGlobalRead(d, meta)
}

func resourceFwGlobalDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceFwGlobalRead(d, meta)
}
func dataToFwGlobal(d *schema.ResourceData) go_thunder.FwGlobal {
	var vc go_thunder.FwGlobal
	var c go_thunder.FwGlobalInstance
	c.AlgProcessing = d.Get("alg_processing").(string)
	c.ListenOnPortTimeout = d.Get("listen_on_port_timeout").(int)
	c.DisableIPFwSessions = d.Get("disable_ip_fw_sessions").(int)

	DisableAppListCount := d.Get("disable_app_list.#").(int)
	c.DisableApplicationProtocol = make([]go_thunder.FwGlobalDisableAppList, 0, DisableAppListCount)

	for i := 0; i < DisableAppListCount; i++ {
		var obj1 go_thunder.FwGlobalDisableAppList
		prefix := fmt.Sprintf("disable_app_list.%d.", i)
		obj1.DisableApplicationProtocol = d.Get(prefix + "disable_application_protocol").(string)
		obj1.DisableApplicationCategory = d.Get(prefix + "disable_application_category").(string)
		c.DisableApplicationProtocol = append(c.DisableApplicationProtocol, obj1)
	}

	c.ExtendedMatching = d.Get("extended_matching").(string)

	SamplingEnableCount := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_thunder.FwGlobalSamplingEnable, 0, SamplingEnableCount)

	for i := 0; i < SamplingEnableCount; i++ {
		var obj2 go_thunder.FwGlobalSamplingEnable
		prefix := fmt.Sprintf("sampling_enable.%d.", i)
		obj2.Counters1 = d.Get(prefix + "counters1").(string)
		c.Counters1 = append(c.Counters1, obj2)
	}

	c.RespondToUserMac = d.Get("respond_to_user_mac").(int)
	c.PermitDefaultAction = d.Get("permit_default_action").(string)
	c.NatipDdosProtection = d.Get("natip_ddos_protection").(string)

	vc.AlgProcessing = c
	return vc
}
