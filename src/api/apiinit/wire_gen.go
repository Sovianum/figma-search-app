// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package apiinit

import (
	"github.com/Sovianum/figma-search-app/src/api"
	"github.com/Sovianum/figma-search-app/src/client/clienttag"
	"github.com/Sovianum/figma-search-app/src/domain/tag/tagimpl"
)

// Injectors from init.go:

func InitializeAPI() *api.API {
	manager := tagimpl.NewManager()
	converter := clienttag.NewConverter()
	tagger := tagimpl.NewTagger()
	tagEndpoints := api.NewTagEndpoints(manager, converter, tagger)
	apiAPI := api.NewAPI(tagEndpoints)
	return apiAPI
}
