package openid

type OpenIdProvider struct {
	Name string
	Auth func(string) (string, error)
}

var Providers = map[string]OpenIdProvider{
	wechatOpenIdProvider.Name: wechatOpenIdProvider,
}
