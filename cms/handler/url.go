package handler

const (
	homePath       = "/"
	notFoundPath   = "/404"
	dashboardPath  = "/dashboard"
	userListPath   = "/users"
	userCreatePath = "/users/create"
	userStorePath  = "/users/store"
	userEditPath   = "/users/edit/{user_id}"
	userUpdatePath = "/users/update/{user_id}"
	deleteUserPath = "users/delete/{user_id}"
	viewUserPath   = "/users/view/{user_id}"
)

func listOfURLs() map[string]string {
	return map[string]string{
		"home":       homePath,
		"dashboard":  dashboardPath,
		"userList":   userListPath,
		"userCreate": userCreatePath,
		"userStore":  userStorePath,
		"userEdit":   userEditPath,
		"userUpdate": userUpdatePath,
		"userDelete": deleteUserPath,
		"userView":   viewUserPath,
	}
}
