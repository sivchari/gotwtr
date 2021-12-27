package gotwtr

func (c *client) ExportClient() map[string]string {
	return map[string]string{
		"bearerToken":    c.bearerToken,
		"consumerKey":    c.consumerKey,
		"consumerSecret": c.consumerSecret,
	}
}
