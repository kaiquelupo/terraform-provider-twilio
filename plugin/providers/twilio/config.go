package twilio

import (
	log "github.com/sirupsen/logrus"
	"context"
	twiclient "github.com/kaiquelupo/twilio-go"
	twiclientServerless "github.com/kaiquelupo/twilio-go-serverless"
)

// Config contains our different configuration attributes and instantiates our Twilio client.
type Config struct {
	AccountSID string
	AuthToken  string
	Endpoint   string
}

// TerraformTwilioContext is our Terraform context that will contain both our Twilio client and configuration for access downstream.
type TerraformTwilioContext struct {
	client        		*twiclient.Client
	clientServerless 	*twiclientServerless.APIClient
	configuration Config
	auth	*context.Context
}

// Client creates a Twilio client and prepares it for use with Terraform.
func (config *Config) Client() (interface{}, error) {
	log.WithFields(
		log.Fields{
			"account_sid": config.AccountSID,
		},
	).Debug("Initializing Twilio client")

	// TODO Support unique endpoints

	client := twiclient.NewClient(config.AccountSID, config.AuthToken, nil)


	//Twilio Serverless API
	cfg := sw.NewConfiguration()
	cfg.Host = "serverless.twilio.com"
	cfg.Scheme = "https"
	clientServerless = sw.NewAPIClient(cfg)
	auth := context.WithValue(context.Background(), sw.ContextBasicAuth, sw.BasicAuth{
		UserName: config.AccountSID,
		Password: config.AuthToken,
	})
	// ---

	context := TerraformTwilioContext{
		client:        client,
		clientServerless: clientServerless,
		auth: auth,
		configuration: *config,
	}

	return &context, nil
}
