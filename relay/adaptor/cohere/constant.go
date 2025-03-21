package cohere

import (
	"github.com/labring/aiproxy/model"
	"github.com/labring/aiproxy/relay/relaymode"
)

var ModelList = []*model.ModelConfig{
	{
		Model: "command",
		Type:  relaymode.ChatCompletions,
		Owner: model.ModelOwnerCohere,
	},
	{
		Model: "command-nightly",
		Type:  relaymode.ChatCompletions,
		Owner: model.ModelOwnerCohere,
	},
	{
		Model: "command-light",
		Type:  relaymode.ChatCompletions,
		Owner: model.ModelOwnerCohere,
	},
	{
		Model: "command-light-nightly",
		Type:  relaymode.ChatCompletions,
		Owner: model.ModelOwnerCohere,
	},
	{
		Model: "command-r",
		Type:  relaymode.ChatCompletions,
		Owner: model.ModelOwnerCohere,
	},
	{
		Model: "command-r-plus",
		Type:  relaymode.ChatCompletions,
		Owner: model.ModelOwnerCohere,
	},
}
