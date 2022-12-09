package routes

import (
	"dot-crud-redis-go-api/configs"
	"dot-crud-redis-go-api/controllers"

	"github.com/gorilla/mux"

	repositories "dot-crud-redis-go-api/repositories/implementations"
	usecases "dot-crud-redis-go-api/usecases/implementations"
)

func PostRoute(router *mux.Router) {
	DB := configs.GetDB()
	RedisClient := configs.GetRedis()

	postRepository := repositories.CreatePostRepo(DB, RedisClient)
	postUsecase := usecases.CreatePostUsecase(postRepository)
	postController := controllers.CreatePostController(postUsecase)

	router.Handle("/posts", postController.GetPosts()).Methods("GET")
	router.Handle("/posts/{id}", postController.GetPost()).Methods("GET")
	router.Handle("/posts", postController.CreatePost()).Methods("POST")
	router.Handle("/posts/{postId}", postController.UpdatePost()).Methods("PUT")
	router.Handle("/posts/{postId}", postController.UpdatePost()).Methods("PATCH")
	router.Handle("/posts/{postId}", postController.DeletePost()).Methods("DELETE")
}