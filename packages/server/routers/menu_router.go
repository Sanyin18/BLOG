package routers

import (
	"web/api"
)

func (router RouterGroup) MenuRouter(){
	app := api.ApiGroupApp.MenuApi
	router.POST("menus", app.MenuCreateView)
	router.GET("menus", app.MenuListView)
	router.GET("menu_names", app.MenuNameList)
	router.PUT("menus/:id", app.MenuUpdateView)
	router.DELETE("menus", app.MenuRemoveView)
	router.GET("menus/:id", app.MenuDetailView)
}
