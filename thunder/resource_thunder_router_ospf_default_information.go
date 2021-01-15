package thunder

//Thunder resource RouterOspfDefaultInformation

import (
	"github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"util"
)

func resourceRouterOspfDefaultInformation() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterOspfDefaultInformationCreate,
		Update: resourceRouterOspfDefaultInformationUpdate,
		Read:   resourceRouterOspfDefaultInformationRead,
		Delete: resourceRouterOspfDefaultInformationDelete,
		Schema: map[string]*schema.Schema{
			"originate": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"always": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"metric": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"metric_type": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"route_map": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"as_number": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"process_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceRouterOspfDefaultInformationCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating RouterOspfDefaultInformation (Inside resourceRouterOspfDefaultInformationCreate) ")
		name1 := d.Get("process_id").(string)
		data := dataToRouterOspfDefaultInformation(d)
		logger.Println("[INFO] received formatted data from method data to RouterOspfDefaultInformation --")
		d.SetId(name1)
		go_thunder.PostRouterOspfDefaultInformation(client.Token, name1, data, client.Host)

		return resourceRouterOspfDefaultInformationRead(d, meta)

	}
	return nil
}

func resourceRouterOspfDefaultInformationRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading RouterOspfDefaultInformation (Inside resourceRouterOspfDefaultInformationRead)")

	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Fetching service Read" + name1)
		data, err := go_thunder.GetRouterOspfDefaultInformation(client.Token, name1, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name1)
			return nil
		}
		return err
	}
	return nil
}

func resourceRouterOspfDefaultInformationUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceRouterOspfDefaultInformationRead(d, meta)
}

func resourceRouterOspfDefaultInformationDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceRouterOspfDefaultInformationRead(d, meta)
}
func dataToRouterOspfDefaultInformation(d *schema.ResourceData) go_thunder.RouterOspfDefaultInformation {
	var vc go_thunder.RouterOspfDefaultInformation
	var c go_thunder.RouterOspfDefaultInformationInstance
	c.Originate = d.Get("originate").(int)
	c.Always = d.Get("always").(int)
	c.Metric = d.Get("metric").(int)
	c.MetricType = d.Get("metric_type").(int)
	c.RouteMap = d.Get("route_map").(string)

	vc.Always = c
	return vc
}
