package main

import "fmt"

type PlatformConfigI interface {
	GetAds() string
	GetID() string
	SetAds() string
	SetChannels() string
}

type CommonConfig struct {
	ID string
}

func (c *CommonConfig) GetID() string {
	return ""
}

func (c *CommonConfig) SetAds() string {
	fmt.Println("this is default set ads")
	return "this is default set ads"
}

func (c *CommonConfig) SetChannels() string {
	return "this is default"
}

func (c *CommonConfig) GetAds() string {
	return "this is default"
}

type TvCommonConfig struct {
	TvCommon string
}

type AndroidTvConfig struct {
	CommonConfig
	TvCommonConfig
	AndroidTvConfigOnly
}

type AndroidTvConfigOnly struct {
	Abc string
}

func (a *AndroidTvConfig) GetAds() string {
	return ""
}

func (c *AndroidTvConfig) SetChannels() string {
	fmt.Println("setting android channel")
	return ""
}

type IosTvConfig struct {
	CommonConfig
	TvCommonConfig
	IosTvConfigOnly
}

type IosTvConfigOnly struct {
	Xyz string
}

func (i *IosTvConfig) GetAds() string {
	return ""
}

func (c *IosTvConfig) SetChannels() string {
	fmt.Println("setting iostv channels")
	return ""
}

type WebConfig struct {
	CommonConfig
	WebConfigOnly
}

type WebConfigOnly struct {
	Xyz string
}

func (c *WebConfig) SetAds() string {
	fmt.Println("this is web set ads")
	return "this is web set ads"
}

//configRepo.FetchConfig
func NewConfig(configType string) PlatformConfigI {

	switch configType {
	case "android":
		return &AndroidTvConfig{
			GetCommonConfig(),
			GetTvCommonConfig(),
			GetAndroidTvConfigOnly(),
		}
	case "iostv":
		return &IosTvConfig{
			GetCommonConfig(),
			GetTvCommonConfig(),
			GetIosConfigOnly(),
		}
	case "web":
		return &WebConfig{
			GetCommonConfig(),
			GetWebOnlyConfig(),
		}
	}

	return &AndroidTvConfig{}
}

func GetCommonConfig() CommonConfig {
	return CommonConfig{}
}

//dbcall
func GetAndroidTvConfigOnly() AndroidTvConfigOnly {
	// collection := s.db.Collection(types.PlatformConfigCollection)
	// 	filter := bson.M{
	// 		"version":  version,
	// 		"name": "android-tv-only-config",
	// 	}
	// 	findOptions := options.FindOne()
	// 	var config types.AndroidTvConfigOnly
	// 	err := collection.FindOne(ctx, filter, findOptions).Decode(&config)
	return AndroidTvConfigOnly{}
}

func GetIosConfigOnly() IosTvConfigOnly {
	return IosTvConfigOnly{}
}

func GetTvCommonConfig() TvCommonConfig {
	// collection := s.db.Collection(types.PlatformConfigCollection)
	// 	filter := bson.M{
	// 		"version":  version,
	// 		"name": "tv-common-config",
	// 	}
	// 	findOptions := options.FindOne()
	// 	var config types.TvCommonConfig
	// 	err := collection.FindOne(ctx, filter, findOptions).Decode(&config)
	return TvCommonConfig{}
}

func GetWebOnlyConfig() WebConfigOnly {
	return WebConfigOnly{}
}

func main() {
	config := NewConfig("iostv")
	config.SetChannels()
	config.SetAds()
}
