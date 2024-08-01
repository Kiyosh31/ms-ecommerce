package handler

import (
	"net/http"

	customlogger "github.com/Kiyosh31/ms-ecommerce-common/custom_logger"
	"github.com/Kiyosh31/ms-ecommerce-common/utils"
	productPb "github.com/Kiyosh31/ms-ecommerce/gateway-api/generated/product-service"
)

func (h *GatewayApiHandler) createProduct(w http.ResponseWriter, r *http.Request) {
	h.logger.Infof("create product request incoming: %v", customlogger.ReadRequestPayload(r))
	var payload productPb.Product

	if err := utils.ReadJSON(r, &payload); err != nil {
		h.logger.Errorf("failed to read product payload: %v", err)
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if errs := validateProductPayload(&payload); len(errs) > 0 {
		h.logger.Errorf("failed to validate product payload: %v", errs)
		utils.WriteErrors(w, http.StatusBadRequest, errs)
		return
	}

	res, err := h.productServiceClient.CreateProduct(r.Context(), &productPb.CreateProductRequest{
		Product: &payload,
	})
	if err != nil {
		h.logger.Errorf("failed to create product: %v", err)
		utils.WriteRpcError(err, w)
		return
	}

	h.logger.Infof("create product finished: %v", res)
	utils.WriteResponse(w, http.StatusCreated, res)
}

func (h *GatewayApiHandler) getProduct(w http.ResponseWriter, r *http.Request) {
	h.logger.Infof("get product request incoming: %v", customlogger.ReadRequestPayload(r))
	productId := r.PathValue("productId")

	res, err := h.productServiceClient.GetProduct(r.Context(), &productPb.GetProductRequest{
		ProductId: productId,
	})
	if err != nil {
		h.logger.Errorf("failed to get product: %v", err)
		utils.WriteRpcError(err, w)
		return
	}

	h.logger.Infof("get product request finished: %v", res)
	utils.WriteResponse(w, http.StatusOK, res)
}

func (h *GatewayApiHandler) getAllProducts(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("get all products request incoming")

	res, err := h.productServiceClient.GetAllProducts(r.Context(), &productPb.GetAllProductsRequest{})
	if err != nil {
		h.logger.Errorf("failed to get all products: %v", err)
		utils.WriteRpcError(err, w)
		return
	}

	h.logger.Infof("get all products finished: %v", res)
	utils.WriteResponse(w, http.StatusOK, res)
}

func (h *GatewayApiHandler) updateProduct(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("productId")
	h.logger.Infof("update product request incoming with id: %v and body: %v", productId, customlogger.ReadRequestPayload(r))

	var payload productPb.Product

	if err := utils.ReadJSON(r, &payload); err != nil {
		h.logger.Errorf("failed to read product payload: %v", err)
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if errs := validateProductPayload(&payload); len(errs) > 0 {
		h.logger.Errorf("failed to validate product payload: %v", errs)
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
		h.logger.Errorf("failed to update product: %v", err)
		utils.WriteRpcError(err, w)
		return
	}

	h.logger.Infof("update request input finished: %v", res)
	utils.WriteResponse(w, http.StatusOK, res)
}

func (h *GatewayApiHandler) deleteProduct(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("productId")
	h.logger.Infof("delete product request incoming: %v", productId)
	if productId == "" {
		h.logger.Error("missing productId")
		utils.WriteError(w, http.StatusBadRequest, "productId missing")
		return
	}

	res, err := h.productServiceClient.DeleteProduct(r.Context(), &productPb.DeleteProductRequest{
		ProductId: productId,
	})
	if err != nil {
		h.logger.Errorf("failed to delete product: %v", err)
		utils.WriteRpcError(err, w)
		return
	}

	h.logger.Infof("delete product request finished: %v", res)
	utils.WriteResponse(w, http.StatusOK, res)
}
