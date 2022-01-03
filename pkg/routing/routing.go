package routing

import (
	"eduapp/pkg/handler"
	"eduapp/pkg/middleware"
	"github.com/julienschmidt/httprouter"
)

func InitRouter(router *httprouter.Router, pathName string) {
	router.POST(pathName+"/api/user/login", handler.UserLogin)
	router.POST(pathName+"/api/user/signup", handler.UserSignup)
	router.POST(pathName+"/api/wallet_signup", handler.CreateWallet)

	routerWrap := NewRouterWrap(pathName, router)

	//GET routers
	routerWrap.GET("/register", handler.RegisterUserPage)
	routerWrap.GET("/login", handler.LoginUserPage)
	routerWrap.GET("/tac", handler.TermsAndConditions)
	routerWrap.GET("/course", middleware.Middleware(handler.CoursePage))         //login required
	routerWrap.GET("/mycourses", middleware.Middleware(handler.UserCoursesPage)) //login required
	routerWrap.GET("/market", middleware.Middleware(handler.MarketPage))         //login required
	routerWrap.GET("/homepage", middleware.Middleware(handler.HomePage))         //login required
	routerWrap.GET("/verify", middleware.Middleware(handler.WalletVerifyPage))   //login required

	//POST routers
	routerWrap.POST("/course", middleware.Middleware(handler.CoursePost))         //login required
	routerWrap.POST("/mycourses", middleware.Middleware(handler.UserCoursesPost)) //login required
	routerWrap.POST("/market", middleware.Middleware(handler.MarketPost))         //login required
	routerWrap.POST("/homepage", middleware.Middleware(handler.HomePost))         //login required
}
