package main

func main() {
	//// Load Config
	//err := config.Global.Load("./config-dev.yml")
	//if err != nil {
	//	log.Fatal("Config load fail", err)
	//}
	//// Load DB
	//db, err := model.LoadFromConfig()
	//if err != nil {
	//	log.Fatal("Database can't be loaded from config", err)
	//}
	//
	//// Init echo
	//e := echo.New()
	//e.Validator = handler.NewValidator()
	//
	//e.Use(middleware.Logger())
	//e.Use(middleware.Recover())
	//e.Use(xecho.BreakErrorRecover())
	//e.Use(myMiddleware.DbContext(db))
	//e.Use(myMiddleware.Jwt())
	//
	//InitRouter(e)
	//
	//log.Fatal(e.Start(":80"))
}
