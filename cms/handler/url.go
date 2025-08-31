package handler

const (
	homePath             = "/"
	notFoundPath         = "/404"
	dashboardPath        = "/dashboard"
	loginPath            = "/login"
	loginAuthPath        = "/login/auth"
	userListPath         = "/users"
	userCreatePath       = "/users/create"
	userStorePath        = "/users/store"
	userEditPath         = "/users/edit/{user_id}"
	userUpdatePath       = "/users/update/{user_id}"
	userDeletePath       = "/users/delete/{user_id}"
	userViewPath         = "/users/view/{user_id}"
	collectionListPath   = "/collection"
	collectionCreatePath = "/collection/create"
	collectionStorePath  = "/collection/store"
	collectionEditPath   = "/collection/edit/{collection_id}"
	collectionUpdatePath = "/collection/update/{collection_id}"
	collectionDeletePath = "/collection/delete/{collection_id}"
	collectionViewPath   = "/collection/view/{collection_id}"
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
		"userDelete": userDeletePath,
		"userView":   userViewPath,
		"collList":   collectionListPath,
		"collCreate": collectionCreatePath,
		"collStore":  collectionStorePath,
		"collEdit":   collectionEditPath,
		"collUpdate": collectionUpdatePath,
		"collDelete": collectionDeletePath,
		"collView":   collectionViewPath,
	}
}
