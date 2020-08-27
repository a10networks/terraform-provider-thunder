package vthunder

import (
	"github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"util"
)

func resourceReboot() *schema.Resource {
	return &schema.Resource{
		Create: resourceRebootCreate,
		Update: resourceRebootUpdate,
		Read:   resourceRebootRead,
		Delete: resourceRebootDelete,

		Schema: map[string]*schema.Schema{
			"all": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"reason_3": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"cancel": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"reason": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"reason_2": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"in": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"month_2": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"at": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"month": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"day_of_month_2": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"time": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"day_of_month": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"device": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
		},
	}

}

func resourceRebootCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	logger.Println("[INFO] Creating Reboot (Inside resourceRebootCreate)")

	if client.Host != "" {
		vc := dataToReboot(d)
		go_vthunder.PostReboot(client.Token, vc, client.Host)
		d.SetId("1")

		return resourceRebootRead(d, meta)
	}
	return nil
}

func resourceRebootRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	logger.Println("[INFO] Reading Reboot (Inside resourceRebootRead)")

	client := meta.(vThunder)

	if client.Host != "" {

		vc, err := go_vthunder.GetReboot(client.Token, client.Host)

		if vc == nil {
			logger.Println("[INFO] No Reboot found")
			//d.SetId("")
			return nil
		}

		return err
	}
	return nil
}

func resourceRebootUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceRebootRead(d, meta)
}

func resourceRebootDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceRebootRead(d, meta)
}

//Utility method to instantiate Reboot Structure
//Configuration for basic required parameters
//TODO: Conf. for all the parameters

func dataToReboot(d *schema.ResourceData) go_vthunder.Reboot {
	var vc go_vthunder.Reboot

	var c go_vthunder.RebootInstance

	c.All = d.Get("all").(int)

	vc.All = c

	return vc
}
