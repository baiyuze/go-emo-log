package handler

import "go.uber.org/dig"

func Provide(container *dig.Container) {
	ProviderUserHandler(container)
	ProviderRpcHandler(container)
	ProviderRolesHandler(container)
	ProviderPermissionsHandler(container)
	ProviderDepartmentHandler(container)
	ProviderDictHandler(container)
	ProviderEmoHandler(container)
	ProviderVersionHandler(container)
	ProviderFeedbackHandler(container)
	ProviderDevicesHandler(container)
	ProviderAiHandler(container)
}
