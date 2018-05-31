package uptimerobot

import (
	"github.com/hashicorp/terraform/helper/schema"
)

type UptimeRobotConfig struct {
	apiKey string
}

/*
	UptimeRobotProvider integrates terraform with uptimerobot.com

	```
	provider "uptimerobot" {
	  api_key = "[YOUR MAIN API KEY]"
	}
	```
*/
func UptimeRobotProvider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_key": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
		DataSourcesMap: map[string]*schema.Resource{
			"uptimerobot_account": dataSourceAccount(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"uptimerobot_alert_contact": resourceAlertContact(),
			"uptimerobot_monitor":       resourceMonitor(),
			"uptimerobot_status_page":   resourceStatusPage(),
		},
		ConfigureFunc: func(r *schema.ResourceData) (interface{}, error) {
			config := UptimeRobotConfig{
				apiKey: r.Get("api_key").(string),
			}
			return config, nil
		},
	}
}