package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	// Order
	orderHandler "github.com/MingPV/clean-go-template/internal/order/handler/rest"
	orderRepository "github.com/MingPV/clean-go-template/internal/order/repository"
	orderUseCase "github.com/MingPV/clean-go-template/internal/order/usecase"

	// Inventory
	inventoryHandler "github.com/MingPV/clean-go-template/internal/inventory/handler/rest"
	inventoryRepository "github.com/MingPV/clean-go-template/internal/inventory/repository"
	inventoryUseCase "github.com/MingPV/clean-go-template/internal/inventory/usecase"

	// EquipmentSlot
	equipmentSlotHandler "github.com/MingPV/clean-go-template/internal/equipment_slot/handler/rest"
	equipmentSlotRepository "github.com/MingPV/clean-go-template/internal/equipment_slot/repository"
	equipmentSlotUseCase "github.com/MingPV/clean-go-template/internal/equipment_slot/usecase"

	// Character
	characterHandler "github.com/MingPV/clean-go-template/internal/character/handler/rest"
	characterRepository "github.com/MingPV/clean-go-template/internal/character/repository"
	characterUseCase "github.com/MingPV/clean-go-template/internal/character/usecase"

	// Status
	statusHandler "github.com/MingPV/clean-go-template/internal/status/handler/rest"
	statusRepository "github.com/MingPV/clean-go-template/internal/status/repository"
	statusUseCase "github.com/MingPV/clean-go-template/internal/status/usecase"

	// Class
	classHandler "github.com/MingPV/clean-go-template/internal/class/handler/rest"
	classRepository "github.com/MingPV/clean-go-template/internal/class/repository"
	classUseCase "github.com/MingPV/clean-go-template/internal/class/usecase"

	// User
	userHandler "github.com/MingPV/clean-go-template/internal/user/handler/rest"
	userRepository "github.com/MingPV/clean-go-template/internal/user/repository"
	userUseCase "github.com/MingPV/clean-go-template/internal/user/usecase"
)

func RegisterPublicRoutes(app fiber.Router, db *gorm.DB) {

	api := app.Group("/api/v1")

	// === Dependency Wiring ===

	// Order
	orderRepo := orderRepository.NewGormOrderRepository(db)
	orderService := orderUseCase.NewOrderService(orderRepo)
	orderHandler := orderHandler.NewHttpOrderHandler(orderService)

	// Inventory
	inventoryRepo := inventoryRepository.NewGormInventoryRepository(db)
	inventoryService := inventoryUseCase.NewInventoryService(inventoryRepo)
	inventoryHandler := inventoryHandler.NewHttpInventoryHandler(inventoryService)

	// EquipmentSlot
	equipmentSlotRepo := equipmentSlotRepository.NewGormEquipmentSlotRepository(db)
	equipmentSlotService := equipmentSlotUseCase.NewEquipmentSlotService(equipmentSlotRepo)
	equipmentSlotHandler := equipmentSlotHandler.NewHttpEquipmentSlotHandler(equipmentSlotService)

	// Status
	statusRepo := statusRepository.NewGormStatusRepository(db)
	statusService := statusUseCase.NewStatusService(statusRepo)
	statusHandler := statusHandler.NewHttpStatusHandler(statusService)

	// Class
	classRepo := classRepository.NewGormClassRepository(db)
	classService := classUseCase.NewClassService(classRepo)
	classHandler := classHandler.NewHttpClassHandler(classService)

	// Character
	characterRepo := characterRepository.NewGormCharacterRepository(db)
	characterService := characterUseCase.NewCharacterService(characterRepo, statusRepo, inventoryRepo, equipmentSlotRepo)
	characterHandler := characterHandler.NewHttpCharacterHandler(characterService)

	// User
	userRepo := userRepository.NewGormUserRepository(db)
	userService := userUseCase.NewUserService(userRepo)
	userHandler := userHandler.NewHttpUserHandler(userService)

	// === Public Routes ===

	// Auth routes (separated from /users)
	authGroup := api.Group("/auth")
	authGroup.Post("/signup", userHandler.Register)
	authGroup.Post("/signin", userHandler.Login)
	authGroup.Post("/signin/username", userHandler.LoginWithUsername)

	// User routes
	userGroup := api.Group("/users")
	userGroup.Get("/", userHandler.FindAllUsers)
	userGroup.Get("/:id", userHandler.FindUserByID)
	userGroup.Get("/email/:email", userHandler.FindUserByEmail)
	userGroup.Get("/username/:username", userHandler.FindUserByUsername)

	// Order routes
	orderGroup := api.Group("/orders")
	orderGroup.Get("/", orderHandler.FindAllOrders)
	orderGroup.Get("/:id", orderHandler.FindOrderByID)
	orderGroup.Post("/", orderHandler.CreateOrder)
	orderGroup.Delete("/:id", orderHandler.DeleteOrder)
	orderGroup.Patch("/:id", orderHandler.PatchOrder)

	// Character routes
	characterGroup := api.Group("/characters")
	characterGroup.Get("/", characterHandler.FindAllCharacters)
	characterGroup.Get("/:id", characterHandler.FindCharacterByID)
	characterGroup.Post("/", characterHandler.CreateCharacter)
	// characterGroup.Delete("/:id", characterHandler.DeleteCharacter)
	// characterGroup.Patch("/:id", characterHandler.PatchCharacter)

	// Status routes
	statusGroup := api.Group("/statuses")
	statusGroup.Get("/", statusHandler.FindAllStatuses)
	// statusGroup.Get("/:id", statusHandler.FindStatusByID)
	statusGroup.Post("/", statusHandler.CreateStatus)
	// statusGroup.Delete("/:id", statusHandler.DeleteStatus)
	// statusGroup.Patch("/:id", statusHandler.PatchStatus)

	// Class routes
	classGroup := api.Group("/classes")
	classGroup.Get("/", classHandler.FindAllClasses)
	// classGroup.Get("/:id", classHandler.FindClassByID)
	classGroup.Post("/", classHandler.CreateClass)
	// classGroup.Delete("/:id", classHandler.DeleteClass)
	// classGroup.Patch("/:id", classHandler.PatchClass)

	// Inventory routes
	inventoryGroup := api.Group("/inventories")
	inventoryGroup.Get("/", inventoryHandler.FindAllInventories)
	// inventoryGroup.Get("/:id", inventoryHandler.FindInventoryByID)
	inventoryGroup.Post("/", inventoryHandler.CreateInventory)
	// inventoryGroup.Delete("/:id", inventoryHandler.DeleteInventory)
	// inventoryGroup.Patch("/:id", inventoryHandler.PatchInventory)

	// EquipmentSlot routes
	equipmentSlotGroup := api.Group("/equipmentSlots")
	equipmentSlotGroup.Get("/", equipmentSlotHandler.FindAllEquipmentSlots)
	// equipmentSlotGroup.Get("/:id", equipmentSlotHandler.FindEquipmentSlotByID)
	equipmentSlotGroup.Post("/", equipmentSlotHandler.CreateEquipmentSlot)
	// equipmentSlotGroup.Delete("/:id", equipmentSlotHandler.DeleteEquipmentSlot)
	// equipmentSlotGroup.Patch("/:id", equipmentSlotHandler.PatchEquipmentSlot)

}
