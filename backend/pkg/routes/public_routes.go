package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	// Order
	orderHandler "github.com/MingPV/clean-go-template/internal/order/handler/rest"
	orderRepository "github.com/MingPV/clean-go-template/internal/order/repository"
	orderUseCase "github.com/MingPV/clean-go-template/internal/order/usecase"

	// Character
	characterHandler "github.com/MingPV/clean-go-template/internal/character/handler/rest"
	characterRepository "github.com/MingPV/clean-go-template/internal/character/repository"
	characterUseCase "github.com/MingPV/clean-go-template/internal/character/usecase"

	// Class
	classHandler "github.com/MingPV/clean-go-template/internal/class/handler/rest"
	classRepository "github.com/MingPV/clean-go-template/internal/class/repository"
	classUseCase "github.com/MingPV/clean-go-template/internal/class/usecase"

	// EquipmentSlot
	equipmentSlotHandler "github.com/MingPV/clean-go-template/internal/equipment_slot/handler/rest"
	equipmentSlotRepository "github.com/MingPV/clean-go-template/internal/equipment_slot/repository"
	equipmentSlotUseCase "github.com/MingPV/clean-go-template/internal/equipment_slot/usecase"

	// Inventory
	inventoryHandler "github.com/MingPV/clean-go-template/internal/inventory/handler/rest"
	inventoryRepository "github.com/MingPV/clean-go-template/internal/inventory/repository"
	inventoryUseCase "github.com/MingPV/clean-go-template/internal/inventory/usecase"

	// Item
	itemHandler "github.com/MingPV/clean-go-template/internal/item/handler/rest"
	itemRepository "github.com/MingPV/clean-go-template/internal/item/repository"
	itemUseCase "github.com/MingPV/clean-go-template/internal/item/usecase"

	// ItemInstance
	itemInstanceHandler "github.com/MingPV/clean-go-template/internal/item_instance/handler/rest"
	itemInstanceRepository "github.com/MingPV/clean-go-template/internal/item_instance/repository"
	itemInstanceUseCase "github.com/MingPV/clean-go-template/internal/item_instance/usecase"

	// ItemLevelStat
	itemLevelStatHandler "github.com/MingPV/clean-go-template/internal/item_level_stat/handler/rest"
	itemLevelStatRepository "github.com/MingPV/clean-go-template/internal/item_level_stat/repository"
	itemLevelStatUseCase "github.com/MingPV/clean-go-template/internal/item_level_stat/usecase"

	// ItemType
	itemTypeHandler "github.com/MingPV/clean-go-template/internal/item_type/handler/rest"
	itemTypeRepository "github.com/MingPV/clean-go-template/internal/item_type/repository"
	itemTypeUseCase "github.com/MingPV/clean-go-template/internal/item_type/usecase"

	// Rarity
	rarityHandler "github.com/MingPV/clean-go-template/internal/rarity/handler/rest"
	rarityRepository "github.com/MingPV/clean-go-template/internal/rarity/repository"
	rarityUseCase "github.com/MingPV/clean-go-template/internal/rarity/usecase"

	// LevelProgress
	levelProgressHandler "github.com/MingPV/clean-go-template/internal/level_progress/handler/rest"
	levelProgressRepository "github.com/MingPV/clean-go-template/internal/level_progress/repository"
	levelProgressUseCase "github.com/MingPV/clean-go-template/internal/level_progress/usecase"

	// Monster
	monsterHandler "github.com/MingPV/clean-go-template/internal/monster/handler/rest"
	monsterRepository "github.com/MingPV/clean-go-template/internal/monster/repository"
	monsterUseCase "github.com/MingPV/clean-go-template/internal/monster/usecase"

	// MonsterType
	monsterTypeHandler "github.com/MingPV/clean-go-template/internal/monster_type/handler/rest"
	monsterTypeRepository "github.com/MingPV/clean-go-template/internal/monster_type/repository"
	monsterTypeUseCase "github.com/MingPV/clean-go-template/internal/monster_type/usecase"

	// MonsterLoot
	monsterLootHandler "github.com/MingPV/clean-go-template/internal/monster_loot/handler/rest"
	monsterLootRepository "github.com/MingPV/clean-go-template/internal/monster_loot/repository"
	monsterLootUseCase "github.com/MingPV/clean-go-template/internal/monster_loot/usecase"

	// Status
	statusHandler "github.com/MingPV/clean-go-template/internal/status/handler/rest"
	statusRepository "github.com/MingPV/clean-go-template/internal/status/repository"
	statusUseCase "github.com/MingPV/clean-go-template/internal/status/usecase"

	// User
	userHandler "github.com/MingPV/clean-go-template/internal/user/handler/rest"
	userRepository "github.com/MingPV/clean-go-template/internal/user/repository"
	userUseCase "github.com/MingPV/clean-go-template/internal/user/usecase"

	// Setting
	settingHandler "github.com/MingPV/clean-go-template/internal/setting/handler/rest"
	settingRepository "github.com/MingPV/clean-go-template/internal/setting/repository"
	settingUseCase "github.com/MingPV/clean-go-template/internal/setting/usecase"
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

	// Character
	characterRepo := characterRepository.NewGormCharacterRepository(db)
	characterService := characterUseCase.NewCharacterService(characterRepo, statusRepo, inventoryRepo, equipmentSlotRepo)
	characterHandler := characterHandler.NewHttpCharacterHandler(characterService)

	// Class
	classRepo := classRepository.NewGormClassRepository(db)
	classService := classUseCase.NewClassService(classRepo)
	classHandler := classHandler.NewHttpClassHandler(classService)

	// ItemLevelStat
	itemLevelStatRepo := itemLevelStatRepository.NewGormItemLevelStatRepository(db)
	itemLevelStatService := itemLevelStatUseCase.NewItemLevelStatService(itemLevelStatRepo)
	itemLevelStatHandler := itemLevelStatHandler.NewHttpItemLevelStatHandler(itemLevelStatService)

	// Item
	itemRepo := itemRepository.NewGormItemRepository(db)
	itemService := itemUseCase.NewItemService(itemRepo, itemLevelStatRepo)
	itemHandler := itemHandler.NewHttpItemHandler(itemService)

	// ItemInstance
	itemInstanceRepo := itemInstanceRepository.NewGormItemInstanceRepository(db)
	itemInstanceService := itemInstanceUseCase.NewItemInstanceService(itemInstanceRepo)
	itemInstanceHandler := itemInstanceHandler.NewHttpItemInstanceHandler(itemInstanceService)

	// ItemType
	itemTypeRepo := itemTypeRepository.NewGormItemTypeRepository(db)
	itemTypeService := itemTypeUseCase.NewItemTypeService(itemTypeRepo)
	itemTypeHandler := itemTypeHandler.NewHttpItemTypeHandler(itemTypeService)

	// Rarity
	rarityRepo := rarityRepository.NewGormRarityRepository(db)
	rarityService := rarityUseCase.NewRarityService(rarityRepo)
	rarityHandler := rarityHandler.NewHttpRarityHandler(rarityService)

	// LevelProgress
	levelProgressRepo := levelProgressRepository.NewGormLevelProgressRepository(db)
	levelProgressService := levelProgressUseCase.NewLevelProgressService(levelProgressRepo)
	levelProgressHandler := levelProgressHandler.NewHttpLevelProgressHandler(levelProgressService)

	// Monster
	monsterRepo := monsterRepository.NewGormMonsterRepository(db)
	monsterService := monsterUseCase.NewMonsterService(monsterRepo)
	monsterHandler := monsterHandler.NewHttpMonsterHandler(monsterService)

	// MonsterLoot
	monsterLootRepo := monsterLootRepository.NewGormMonsterLootRepository(db)
	monsterLootService := monsterLootUseCase.NewMonsterLootService(monsterLootRepo)
	monsterLootHandler := monsterLootHandler.NewHttpMonsterLootHandler(monsterLootService)

	// MonsterType
	monsterTypeRepo := monsterTypeRepository.NewGormMonsterTypeRepository(db)
	monsterTypeService := monsterTypeUseCase.NewMonsterTypeService(monsterTypeRepo)
	monsterTypeHandler := monsterTypeHandler.NewHttpMonsterTypeHandler(monsterTypeService)

	// Setting
	settingRepo := settingRepository.NewGormSettingRepository(db)
	settingService := settingUseCase.NewSettingService(settingRepo)
	settingHandler := settingHandler.NewHttpSettingHandler(settingService)

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
	characterGroup.Get("/userid/:userID", characterHandler.FindCharacterByUserID)
	characterGroup.Post("/", characterHandler.CreateCharacter)
	characterGroup.Delete("/:id", characterHandler.DeleteCharacter)
	characterGroup.Patch("/:id", characterHandler.PatchCharacter)

	// Status routes
	statusGroup := api.Group("/statuses")
	statusGroup.Get("/", statusHandler.FindAllStatuses)
	statusGroup.Get("/:character_id", statusHandler.FindStatusByCharacterID)
	statusGroup.Post("/", statusHandler.CreateStatus)
	statusGroup.Delete("/:character_id", statusHandler.DeleteStatus)
	statusGroup.Patch("/:character_id", statusHandler.PatchStatus)

	// Class routes
	classGroup := api.Group("/classes")
	classGroup.Get("/", classHandler.FindAllClasses)
	classGroup.Get("/:id", classHandler.FindClassByID)
	classGroup.Post("/", classHandler.CreateClass)
	classGroup.Delete("/:id", classHandler.DeleteClass)
	classGroup.Patch("/:id", classHandler.PatchClass)

	// Inventory routes
	inventoryGroup := api.Group("/inventories")
	inventoryGroup.Get("/", inventoryHandler.FindAllInventories)
	inventoryGroup.Get("/:id", inventoryHandler.FindInventoryByID)
	inventoryGroup.Post("/", inventoryHandler.CreateInventory)
	inventoryGroup.Delete("/:id", inventoryHandler.DeleteInventory)
	inventoryGroup.Patch("/:id", inventoryHandler.PatchInventory)

	// EquipmentSlot routes
	equipmentSlotGroup := api.Group("/equipmentSlots")
	equipmentSlotGroup.Get("/", equipmentSlotHandler.FindAllEquipmentSlots)
	equipmentSlotGroup.Get("/:id", equipmentSlotHandler.FindEquipmentSlotByID)
	equipmentSlotGroup.Post("/", equipmentSlotHandler.CreateEquipmentSlot)
	equipmentSlotGroup.Delete("/:id", equipmentSlotHandler.DeleteEquipmentSlot)
	equipmentSlotGroup.Patch("/:id", equipmentSlotHandler.PatchEquipmentSlot)

	// Item routes
	itemGroup := api.Group("/items")
	itemGroup.Get("/", itemHandler.FindAllItems)
	itemGroup.Get("/:id", itemHandler.FindItemByID)
	itemGroup.Post("/", itemHandler.CreateItem)
	itemGroup.Delete("/:id", itemHandler.DeleteItem)
	itemGroup.Patch("/:id", itemHandler.PatchItem)

	// ItemInstance routes
	itemInstanceGroup := api.Group("/itemInstances")
	itemInstanceGroup.Get("/", itemInstanceHandler.FindAllItemInstances)
	itemInstanceGroup.Get("/:id", itemInstanceHandler.FindItemInstanceByID)
	itemInstanceGroup.Post("/", itemInstanceHandler.CreateItemInstance)
	itemInstanceGroup.Delete("/:id", itemInstanceHandler.DeleteItemInstance)
	itemInstanceGroup.Patch("/:id", itemInstanceHandler.PatchItemInstance)

	// ItemLevelStat routes
	itemLevelStatGroup := api.Group("/itemLevelStats")
	itemLevelStatGroup.Get("/", itemLevelStatHandler.FindAllItemLevelStats)
	itemLevelStatGroup.Get("/:id", itemLevelStatHandler.FindItemLevelStatByID)
	itemLevelStatGroup.Post("/", itemLevelStatHandler.CreateItemLevelStat)
	itemLevelStatGroup.Delete("/:id", itemLevelStatHandler.DeleteItemLevelStat)
	itemLevelStatGroup.Patch("/:id", itemLevelStatHandler.PatchItemLevelStat)

	// ItemType routes
	itemTypeGroup := api.Group("/itemTypes")
	itemTypeGroup.Get("/", itemTypeHandler.FindAllItemTypes)
	itemTypeGroup.Get("/:id", itemTypeHandler.FindItemTypeByID)
	itemTypeGroup.Post("/", itemTypeHandler.CreateItemType)
	itemTypeGroup.Delete("/:id", itemTypeHandler.DeleteItemType)
	itemTypeGroup.Patch("/:id", itemTypeHandler.PatchItemType)

	// Rarity routes
	rarityGroup := api.Group("/rarities")
	rarityGroup.Get("/", rarityHandler.FindAllRarities)
	rarityGroup.Get("/:id", rarityHandler.FindRarityByID)
	rarityGroup.Post("/", rarityHandler.CreateRarity)
	rarityGroup.Delete("/:id", rarityHandler.DeleteRarity)
	rarityGroup.Patch("/:id", rarityHandler.PatchRarity)

	// LevelProgress routes
	levelProgressGroup := api.Group("/levelProgresses")
	levelProgressGroup.Get("/", levelProgressHandler.FindAllLevelProgresses)
	levelProgressGroup.Get("/:level", levelProgressHandler.FindLevelProgressByLevel)
	levelProgressGroup.Post("/", levelProgressHandler.CreateLevelProgress)
	levelProgressGroup.Delete("/:level", levelProgressHandler.DeleteLevelProgress)
	levelProgressGroup.Patch("/:level", levelProgressHandler.PatchLevelProgress)

	// Monster routes
	monsterGroup := api.Group("/monsters")
	monsterGroup.Get("/", monsterHandler.FindAllMonsters)
	monsterGroup.Get("/:id", monsterHandler.FindMonsterByID)
	monsterGroup.Post("/", monsterHandler.CreateMonster)
	monsterGroup.Delete("/:id", monsterHandler.DeleteMonster)
	monsterGroup.Patch("/:id", monsterHandler.PatchMonster)

	// MonsterLoot routes
	monsterLootGroup := api.Group("/monsterLoots")
	monsterLootGroup.Get("/", monsterLootHandler.FindAllMonsterLoots)
	monsterLootGroup.Get("/monsterID/:monsterID", monsterLootHandler.FindMonsterLootByMonsterID)
	monsterLootGroup.Get("/itemID/:itemID", monsterLootHandler.FindMonsterLootByItemID)
	monsterLootGroup.Get("/:monsterID/:itemID", monsterLootHandler.FindMonsterLootByMonsterIDAndItemID)
	monsterLootGroup.Post("/", monsterLootHandler.CreateMonsterLoot)
	monsterLootGroup.Delete("/:monsterID/:itemID", monsterLootHandler.DeleteMonsterLoot)
	monsterLootGroup.Patch("/:monsterID/:itemID", monsterLootHandler.PatchMonsterLoot)

	// MonsterType routes
	monsterTypeGroup := api.Group("/monsterTypes")
	monsterTypeGroup.Get("/", monsterTypeHandler.FindAllMonsterTypes)
	monsterTypeGroup.Get("/:id", monsterTypeHandler.FindMonsterTypeByID)
	monsterTypeGroup.Post("/", monsterTypeHandler.CreateMonsterType)
	monsterTypeGroup.Delete("/:id", monsterTypeHandler.DeleteMonsterType)
	monsterTypeGroup.Patch("/:id", monsterTypeHandler.PatchMonsterType)

	// Setting routes
	settingGroup := api.Group("/settings")
	settingGroup.Get("/", settingHandler.FindAllSettings)
	settingGroup.Get("/:id", settingHandler.FindSettingByID)
	settingGroup.Post("/", settingHandler.CreateSetting)
	settingGroup.Delete("/:id", settingHandler.DeleteSetting)
	settingGroup.Patch("/:id", settingHandler.PatchSetting)

}
