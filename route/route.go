package route

import (
	//"fmt"
	//"log"
	//"strings"
	"tentativa/controllers"
	//"tentativa/repo"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	// permissions "github.com/xyproto/permissionbolt"
)

//var userRepo = services.NewMemoryUserService()

/*


func PermissionRoute(app *iris.Application) {
	// New permissions middleware.
	perm, err := permissions.New()
	if err != nil {
		log.Fatalln(err)
	}

	// Blank slate, no default permissions
	// perm.Clear()

	// Set up a middleware handler for Iris, with a custom "permission denied" message.
	permissionHandler := func(ctx iris.Context) {
		// Check if the user has the right admin/user rights
		if perm.Rejected(ctx.ResponseWriter(), ctx.Request()) {
			// Deny the request, don't call other middleware handlers
			ctx.StopWithText(iris.StatusForbidden, "Permission denied!")
			return
		}
		// Call the next middleware handler
		ctx.Next()
	}

	// Register the permissions middleware
	app.Use(permissionHandler)

	// Get the userstate, used in the handlers below
	userstate := perm.UserState()

	app.Get("/", func(ctx iris.Context) {
		msg := ""
		msg += fmt.Sprintf("Has user bob: %v\n", userstate.HasUser("bob"))
		msg += fmt.Sprintf("Logged in on server: %v\n", userstate.IsLoggedIn("bob"))
		msg += fmt.Sprintf("Is confirmed: %v\n", userstate.IsConfirmed("bob"))
		msg += fmt.Sprintf("Username stored in cookies (or blank): %v\n", userstate.Username(ctx.Request()))
		msg += fmt.Sprintf("Current user is logged in, has a valid cookie and *user rights*: %v\n", userstate.UserRights(ctx.Request()))
		msg += fmt.Sprintf("Current user is logged in, has a valid cookie and *admin rights*: %v\n", userstate.AdminRights(ctx.Request()))
		msg += fmt.Sprintln("\nTry: /register, /confirm, /remove, /login, /logout, /makeadmin, /clear, /data and /admin")
		ctx.WriteString(msg)
	})

	app.Get("/register", func(ctx iris.Context) {
		userstate.AddUser("bob", "hunter1", "bob@zombo.com")
		ctx.Writef("User bob was created: %v\n", userstate.HasUser("bob"))
	})

	app.Get("/confirm", func(ctx iris.Context) {
		userstate.MarkConfirmed("bob")
		ctx.Writef("User bob was confirmed: %v\n", userstate.IsConfirmed("bob"))
	})

	app.Get("/remove", func(ctx iris.Context) {
		userstate.RemoveUser("bob")
		ctx.Writef("User bob was removed: %v\n", !userstate.HasUser("bob"))
	})

	app.Get("/login", func(ctx iris.Context) {
		// Headers will be written, for storing a cookie
		userstate.Login(ctx.ResponseWriter(), "bob")
		ctx.Writef("bob is now logged in: %v\n", userstate.IsLoggedIn("bob"))
	})

	app.Get("/logout", func(ctx iris.Context) {
		userstate.Logout("bob")
		ctx.Writef("bob is now logged out: %v\n", !userstate.IsLoggedIn("bob"))
	})

	app.Get("/makeadmin", func(ctx iris.Context) {
		userstate.SetAdminStatus("bob")
		ctx.Writef("bob is now administrator: %v\n", userstate.IsAdmin("bob"))
	})

	app.Get("/clear", func(ctx iris.Context) {
		userstate.ClearCookie(ctx.ResponseWriter())
		ctx.WriteString("Clearing cookie")
	})

	app.Get("/data", func(ctx iris.Context) {
		ctx.WriteString("user page that only logged in users must see!")
	})

	app.Get("/admin", func(ctx iris.Context) {
		ctx.WriteString("super secret information that only logged in administrators must see!\n\n")
		if usernames, err := userstate.AllUsernames(); err == nil {
			ctx.Writef("list of all users: %s" + strings.Join(usernames, ", "))
		}
	})

	bathPath := "/api/v2"
	//mvc.New(app.Party(bathPath + "/permission")).Handle(controllers.UserController)
	mvc.New(app.Party(bathPath + "/permission")).Handle(new(controllers.UserController))
}

*/

func InitRoute(app *iris.Application) {
	/*

		if err := repository.GenerateSamples(userRepo, todoRepo); err != nil {
			panic(err)
		}

		app.PartyFunc("/", api.NewRouter(userRepo, todoRepo))



	*/
	/*
		if err := services.GenerateSamples(userRepo); err != nil {
			panic(err)
		}
	*/

	// POST http://localhost:8080/signin (Form: username, password)
	// GET  http://localhost:8080/todos
	// GET  http://localhost:8080/todos/{id}
	// POST http://localhost:8080/todos (JSON, Form or URL: title, body)
	// GET  http://localhost:8080/admin/todos

	// ________________________________________________________________________________

	bathPath := "/api/v1"
	mvc.New(app.Party(bathPath + "/auth")).Handle(controllers.NewAuthController())
	app.Use(controllers.Verify())
	mvc.New(app.Party(bathPath + "/todos")).Handle(controllers.NewTodoController())
	mvc.New(app.Party(bathPath + "/games")).Handle(controllers.NewGameController())
	mvc.New(app.Party(bathPath + "/mores")).Handle(controllers.NewMoreController())

	mvc.New(app.Party(bathPath + "/addresses")).Handle(controllers.NewAddressController())
	mvc.New(app.Party(bathPath + "/categories")).Handle(controllers.NewCategoryController())
	mvc.New(app.Party(bathPath + "/clients")).Handle(controllers.NewClientController())
	mvc.New(app.Party(bathPath + "/orders")).Handle(controllers.NewOrderController())
	mvc.New(app.Party(bathPath + "/products")).Handle(controllers.NewProductController())
	mvc.New(app.Party(bathPath + "/purchases")).Handle(controllers.NewPurchaseController())
	mvc.New(app.Party(bathPath + "/sales")).Handle(controllers.NewSaleController())
	mvc.New(app.Party(bathPath + "/stocks")).Handle(controllers.NewStockController())
	mvc.New(app.Party(bathPath + "/storeds")).Handle(controllers.NewStoredController())
}
