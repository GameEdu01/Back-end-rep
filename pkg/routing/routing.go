package routing

import (
	"eduapp/pkg/handler"
	"eduapp/pkg/middleware"
	"github.com/julienschmidt/httprouter"
)

func InitRouter(router *httprouter.Router, pathName string) {

	routerWrap := NewRouterWrap(pathName, router)

	//GET routers
	routerWrap.GET("/register", handler.RegisterUserPage)
	routerWrap.GET("/login", handler.LoginUserPage)
	routerWrap.GET("/tac", handler.TermsAndConditions)
	routerWrap.GET("/course", middleware.AuthMiddleware(handler.CoursePage))         //login required
	routerWrap.GET("/mycourses", middleware.AuthMiddleware(handler.UserCoursesPage)) //login required
	routerWrap.GET("/leaderboard", middleware.AuthMiddleware(handler.Leaderboard))   //login required
	routerWrap.GET("/verify", middleware.AuthMiddleware(handler.WalletVerifyPage))   //login required
	routerWrap.GET(pathName+"/api/newsfeed", handler.SendNewsFeed)

	//POST routers
	routerWrap.POST(pathName+"/api/user/login", handler.UserLogin)
	routerWrap.POST(pathName+"/api/user/signup", handler.UserSignup)
	routerWrap.POST(pathName+"/api/wallet_signup", handler.CreateWallet)
	routerWrap.POST("/course", middleware.AuthMiddleware(handler.CoursePost))         //login required
	routerWrap.POST("/mycourses", middleware.AuthMiddleware(handler.UserCoursesPost)) //login required
}
