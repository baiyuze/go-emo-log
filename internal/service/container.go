package service

import "go.uber.org/dig"

func Provide(container *dig.Container) {
	ProvideUserService(container)
	ProvideRolesService(container)
	ProvidePermissionsService(container)
	ProvideDepartmentService(container)
	ProvideDictService(container)
	ProvideEmoService(container)
	ProvideVersionService(container)
	ProvideFeedbackService(container)
	ProvideDevicesService(container)
	ProvideAiService(container)
}
