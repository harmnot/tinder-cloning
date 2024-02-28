package main

import (
	"database/sql"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"tinder-cloning/pkg/driver"
	"tinder-cloning/pkg/util"
	accountHandler "tinder-cloning/services/account/handler"
	_accountRepo "tinder-cloning/services/account/repository"
	_accountUsecase "tinder-cloning/services/account/usecase"
	membershipHandler "tinder-cloning/services/membership/handler"
	_membershipRepo "tinder-cloning/services/membership/repository"
	_membershipUsecase "tinder-cloning/services/membership/usecase"
	swipeHandler "tinder-cloning/services/swipe/handler"
	_swipeRepo "tinder-cloning/services/swipe/repository"
	_swipeUsecase "tinder-cloning/services/swipe/usecase"
)

func main() {
	config, err := driver.LoadGlobal("./.env")
	if err != nil {
		log.Fatalln("Failed To Load Config")
	}
	db := driver.NewSQL(config)

	defer func() {
		err := db.Close()
		if err != nil {
			log.Fatalln("Failed To Close Database Connection")
			return
		}
	}()

	r := chi.NewRouter()
	r.Mount("/", Api(db))
	serverAddress := os.Getenv("SERVER_ADDRESS")
	log.Println("Server Running On Port", serverAddress)
	printLog(serverAddress)
	err = http.ListenAndServe(":9191", r)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func printLog(serverAddress string) {
	log.Println("=================================================")
	log.Println("Server start at http://localhost:" + serverAddress)
	log.Println("run in " + os.Getenv("ENV_APP") + "  mode")
	log.Println("=================================================")
}

func Api(db *sql.DB) chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		util.RenderJSON(w, http.StatusNotFound, "Route Not Found")
	})
	r.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		util.RenderJSON(w, http.StatusMethodNotAllowed, "Method Not Allowed")
	})
	r.Use(cors.Handler(cors.Options{
		AllowedHeaders: []string{"X-Requested-With", "Content-Type", "Authorization"},
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
	}))

	membershipRepo := _membershipRepo.NewMembershipRepositoryImpl(db)
	membershipService := _membershipUsecase.NewMembershipUseCase(membershipRepo)
	membershipHandler.NewMembershipHandler(membershipService).RegisterRoute(r)

	accountRepo := _accountRepo.NewAccountRepositoryImpl(db)
	accountService := _accountUsecase.NewAccountUseCase(accountRepo, membershipService)
	accountHandler.NewAccountHandler(accountService).RegisterRoute(r)

	swipeRepo := _swipeRepo.NewSwipesRepositoryImpl(db)
	swipeService := _swipeUsecase.NewSwipesUseCase(swipeRepo, membershipService, accountService)
	swipeHandler.NewSwipeHandler(swipeService).RegisterRoute(r)

	return r
}
