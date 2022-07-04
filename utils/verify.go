package utils

var (
	RegisterVerify       = Rules{"Mobile": {NotEmpty()}, "Password": {NotEmpty()}}
	LoginVerify          = Rules{"Mobile": {NotEmpty()}, "Password": {NotEmpty()}}
	ChangePasswordVerify = Rules{"Mobile": {NotEmpty()}, "Password": {NotEmpty()}, "NewPassword": {NotEmpty()}}
	PageInfoVerify       = Rules{"Limit": {NotEmpty()}, "Offset": {NotEmpty()}}
	IdVerify             = Rules{"ID": {NotEmpty()}}
	RoleVerify           = Rules{"Name": {NotEmpty()}, "Describe": {NotEmpty()}}
	RolePermissionVerify = Rules{"Role": {NotEmpty()}, "Permission": {NotEmpty()}}
	UserRoleVerify       = Rules{"Role": {NotEmpty()}, "UserID": {NotEmpty()}}
	SMSVerify            = Rules{"Mobile": {NotEmpty()}, "Captcha": {NotEmpty()}, "CaptchaID": {NotEmpty()}}
	CreateTaskVerify     = Rules{"Name": {NotEmpty()}, "FetchURL": {NotEmpty()}, "LimitTimes": {NotEmpty()}, "Status": {NotEmpty()}, "MaxExcuteTime": {NotEmpty()}, "ExcuteStart": {NotEmpty()}, "ExcuteEnd": {NotEmpty()}}
	StartTaskVerify      = Rules{"TaskID": {NotEmpty()}, "UserID": {NotEmpty()}}
	CreateAgentVerify    = Rules{"Name": {NotEmpty()}, "UserID": {NotEmpty()}, "Mobile": {NotEmpty()}}
	CtTaskPushVerify     = Rules{"ID": {NotEmpty()}, "poiName": {NotEmpty()}, "checkIn": {NotEmpty()}, "checkOut": {NotEmpty()}}
	CreatePushVerify     = Rules{"Title": {NotEmpty()}, "Content": {NotEmpty()}}
	SetPushVerify        = Rules{"TaskID": {NotEmpty()}, "PushID": {NotEmpty()}}
)
