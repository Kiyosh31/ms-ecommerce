package handler

import (
	"net/http"

	"github.com/Kiyosh31/ms-ecommerce-common/utils"
	productPb "github.com/Kiyosh31/ms-ecommerce/gateway-api/generated/product-service"
)

func (h *GatewayApiHandler) createProduct(w http.ResponseWriter, r *http.Request) {
	var payload productPb.Product

	if err := utils.ReadJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if errs := validateProductPayload(&payload); len(errs) > 0 {
		utils.WriteErrors(w, http.StatusBadRequest, errs)
		return
	}

	res, err := h.productServiceClient.CreateProduct(r.Context(), &productPb.CreateProductRequest{
		Product: &payload,
	})
	if err != nil {
		utils.WriteRpcError(err, w)
		return
	}

	utils.WriteResponse(w, http.StatusCreated, res)
}

func (h *GatewayApiHandler) getProduct(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("productId")

	res, err := h.productServiceClient.GetProduct(r.Context(), &productPb.GetProductRequest{
		ProductId: productId,
	})
	if err != nil {
		utils.WriteRpcError(err, w)
		return
	}

	utils.WriteResponse(w, http.StatusOK, res)
}

func (h *GatewayApiHandler) getAllProducts(w http.ResponseWriter, r *http.Request) {
	res, err := h.productServiceClient.GetAllProducts(r.Context(), &productPb.GetAllProductsRequest{})
	if err != nil {
		utils.WriteRpcError(err, w)
		return
	}

	utils.WriteResponse(w, http.StatusOK, res)
}

func (h *GatewayApiHandler) updateProduct(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("productId")

	var payload productPb.Product

	if err := utils.ReadJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if errs := validateProductPayload(&payload); len(errs) > 0 {
		utils.WriteErrors(w, http.StatusBadRequest, errs)
		return
	}

	res, err := h.productServiceClient.UpdateProduct(r.Context(), &productPb.UpdateProductRequest{
		Product: &productPb.Product{
			Id:          productId,
			Name:        payload.GetName(),
			Price:       payload.GetPrice(),
			Description: payload.GetDescription(),
			SellerId:    payload.GetSellerId(),
		},
	})
	if err != nil {
		utils.WriteRpcError(err, w)
		return
	}

	utils.WriteResponse(w, http.StatusOK, res)
}

func (h *GatewayApiHandler) deleteProduct(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("productId")
	if productId == "" {
		utils.WriteError(w, http.StatusBadRequest, "productId missing")
		return
	}

	res, err := h.productServiceClient.DeleteProduct(r.Context(), &productPb.DeleteProductRequest{
		ProductId: productId,
	})
	if err != nil {
		utils.WriteRpcError(err, w)
		return
	}

	utils.WriteResponse(w, http.StatusOK, res)
}
