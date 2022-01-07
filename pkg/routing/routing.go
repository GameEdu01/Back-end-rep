package routing

import (
	"eduapp/pkg/auth"
	"eduapp/pkg/course"
	"eduapp/pkg/future"
	"eduapp/pkg/middleware"
	"eduapp/pkg/wallet"
	"github.com/julienschmidt/httprouter"
)

func InitRouter(router *httprouter.Router, pathName string) {

	routerWrap := NewRouterWrap(pathName, router)

	//GET routers
	routerWrap.GET("/", auth.ForwardToNewsFeed)
	routerWrap.GET("/register", auth.RegisterUserPage)
	routerWrap.GET("/login", auth.LoginUserPage)
	routerWrap.GET("/tac", auth.TermsAndConditions)
	routerWrap.GET("/tac_for_wallet", auth.TermsAndConditionsForWallet)
	routerWrap.GET("/newsfeed", middleware.AuthMiddleware(course.NewsFeedPage))
	routerWrap.GET("/course", middleware.AuthMiddleware(course.CoursePage))         //login required
	routerWrap.GET("/mycourses", middleware.AuthMiddleware(course.UserCoursesPage)) //login required
	routerWrap.GET("/leaderboard", middleware.AuthMiddleware(future.Leaderboard))   //login required
	routerWrap.GET("/verify", middleware.AuthMiddleware(wallet.WalletVerifyPage))   //login required
	routerWrap.GET(pathName+"/api/newsfeed", course.SendNewsFeed)

	//POST routers
	routerWrap.POST(pathName+"/api/user/login", auth.UserLogin)
	routerWrap.POST(pathName+"/api/user/signup", auth.UserSignup)
	routerWrap.POST(pathName+"/api/wallet_signup", wallet.CreateWallet)
	routerWrap.POST("/course", course.CoursePost)                                    //login required
	routerWrap.POST("/mycourses", middleware.AuthMiddleware(course.UserCoursesPost)) //login required
}
