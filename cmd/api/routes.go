package main

import (
	"database/sql"

	"github.com/Raihanki/checklisters/cmd/api/handlers"
	"github.com/Raihanki/checklisters/internal/middleware"
	"github.com/Raihanki/checklisters/internal/repositories"
	"github.com/Raihanki/checklisters/internal/services"
	"github.com/gofiber/fiber/v2"
)

func Routes(route *fiber.App, db *sql.DB) {
	api := route.Group("/api/v1")

	userHandler := handlers.UserHandler{
		UserService: services.NewUserService(repositories.NewUserRepository(), db),
	}
	api.Post("/register", userHandler.Register)
	api.Post("/login", userHandler.Login)

	checklistRoute := api.Group("/checklists")
	ChecklistHandler := handlers.ChecklistHandler{
		ChecklistService: services.NewChecklistService(repositories.NewChecklistRepository(), db),
	}
	checklistRoute.Post("", middleware.Authenticate, ChecklistHandler.Store)
	checklistRoute.Get("", middleware.Authenticate, ChecklistHandler.Index)
	checklistRoute.Delete("/:checklistId", middleware.Authenticate, ChecklistHandler.Destroy)

	checklistItemRoute := checklistRoute.Group("/:checklistId/items")
	checklistItemHandler := handlers.ChecklistItemHandler{
		ChecklistItemService: services.NewChecklistItemService(db, repositories.NewChecklistItemRepository(), repositories.NewChecklistRepository()),
	}
	checklistItemRoute.Post("", middleware.Authenticate, checklistItemHandler.Store)
	checklistItemRoute.Get("", middleware.Authenticate, checklistItemHandler.Index)
	checklistItemRoute.Get("/:checklistItemId", middleware.Authenticate, checklistItemHandler.Show)
	checklistItemRoute.Patch("/:checklistItemId", middleware.Authenticate, checklistItemHandler.Rename)
	checklistItemRoute.Delete("/:checklistItemId", middleware.Authenticate, checklistItemHandler.Destroy)
	checklistItemRoute.Patch("/:checklistItemId/complete", middleware.Authenticate, checklistItemHandler.Complete)
}
