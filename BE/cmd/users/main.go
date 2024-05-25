package main

import (
	"thelastking/kingseafood/api"
	"thelastking/kingseafood/api/req"
	"thelastking/kingseafood/middleware"
	"thelastking/kingseafood/server"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	auth := r.Group("/auth")
	{
		auth.POST("/sign-up", api.SignUpHandler(server.Run()))
		auth.POST("/sign-in", api.SignInHandler(server.Run()))
	}
	users := r.Group("/users", middleware.RefreshJWTMiddleware(), middleware.JWTMiddleware())
	{
		users.GET("/profile/id", api.ProfileUser(server.Run()))
		users.PATCH("/update/id", api.UpdateUserHandler(server.Run()))
		users.DELETE("/delete/id", api.DeletedUserHandler(server.Run()))
	}

	menu := r.Group("/menu")
	{
		menu.POST("/", req.CreateMenuHandler(server.Run()))
		menu.GET("/:menu_id", req.HandlerGetMenu(server.Run()))
		menu.GET("/", req.HandlerGetListMenu(server.Run()))
		menu.PATCH("/:menu_id", req.HandlerUpdateMenus(server.Run()))
		menu.DELETE("/:menu_id", req.HandlerDeleteMenu(server.Run()))
	}
	table := r.Group("/table")
	{
		table.POST("/", req.HandlerCreateTables(server.Run()))
		table.GET("/:table_id", req.HandlerGetTable(server.Run()))
		table.GET("/", req.HandlerGetTables(server.Run()))
		table.PATCH("/:table_id", req.HandlerUpdateTables(server.Run()))
		table.DELETE("/:table_id", req.HandlerDeletedTable(server.Run()))
	}
	order := r.Group("/order")
	{
		order.POST("/", req.HandlerCreateOrder(server.Run()))
		order.GET("/:order_id", req.HandlerGetOrder(server.Run()))
		order.PATCH("/:order_id", req.HandlerUpdateOrder(server.Run()))
		order.DELETE("/:order_id", req.HandlerDeleteOrder(server.Run()))
	}
	invoice := r.Group("/invoice")
	{
		invoice.POST("/", req.HandlerCreateInvoice(server.Run()))
		invoice.GET("/:invoice_id", req.HandlerGetInvoice(server.Run()))

		invoice.PATCH("/:invoice_id", req.HandlerUpdateInvoices(server.Run()))
		invoice.DELETE("/:invoice_id", req.HandlerDeleteInvoice(server.Run()))
	}
	food := r.Group("/food")
	{
		food.POST("/", req.HandlerCreateProducts(server.Run()))
		food.GET("/product/:product_id", req.HandlerGetProduct(server.Run()))
		food.GET("/product/", req.HandlerGetProducts(server.Run()))
		food.PATCH("/product/:product_id", req.HandlerUpdateProducts(server.Run()))
		food.PATCH("/product/deleted/:product_id", req.HandlerDeletedProduct(server.Run()))
		food.GET("/product/searchName/:title", req.HandlerGetProductByName(server.Run()))
	}
	orderItems := r.Group("/orderItems")
	{
		orderItems.POST("/", req.HandlerCreateOrderItems(server.Run()))
		orderItems.GET("/:order_item_id", req.HandlerGetOrderItems(server.Run()))
		orderItems.GET("/order-items-by-product/:product_id", req.HandlerGetOrderItemsByProduct(server.Run()))
		orderItems.GET("/order-items-by-order/:order_id", req.HandlerGetOrderItemsByOder(server.Run()))
		orderItems.PATCH("/:order_item_id", req.HandlerUpdateOrderItems(server.Run()))
	}
	r.Run(":3250")
}
