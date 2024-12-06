package service

import (
	eventBus "github.com/asaskevich/EventBus"
)

var localEventBus eventBus.Bus

func EventBus() eventBus.Bus {
	if localEventBus == nil {
		panic("implement not found for interface EventBus, forgot register?")
	}
	return localEventBus
}

func RegisterEventBus(i eventBus.Bus) {
	localEventBus = i
}
