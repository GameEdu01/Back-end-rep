package routing

import (
	"eduapp/pkg/handler"
	"github.com/julienschmidt/httprouter"
)

func InitRouter(router *httprouter.Router, pathName string) {
	router.POST(pathName+"/api/user/login", handler.UserLogin)
	router.POST(pathName+"/api/user/signup", handler.UserSignup)

	routerWrap := NewRouterWrap(pathName, router)

	//GET routers
	routerWrap.GET("/register", handler.RegisterUserPage)
	routerWrap.GET("/login", handler.LoginUserPage)
	routerWrap.GET("/tac", handler.TermsAndConditions)
	routerWrap.GET("/course", handler.CoursePage)         //login required
	routerWrap.GET("/mycourses", handler.UserCoursesPage) //login required
	routerWrap.GET("/market", handler.MarketPage)         //login required
	routerWrap.GET("/homepage", handler.HomePage)         //login required

	//POST routers
	routerWrap.POST("/course", handler.CoursePost)              //login required
	routerWrap.POST("/mycourses", handler.UserCoursesPost)      //login required
	routerWrap.POST("/market", handler.MarketPost)              //login required
	routerWrap.POST("/homepage", handler.HomePost)              //login required
	routerWrap.POST("/api/create-wallet", handler.CreateWallet) //login required
}
