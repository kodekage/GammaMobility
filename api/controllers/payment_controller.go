package paymentcontroller

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/kodekage/gamma_mobility/dto"
	"github.com/kodekage/gamma_mobility/internal/logger"
	"github.com/kodekage/gamma_mobility/internal/queue"
	"github.com/kodekage/gamma_mobility/utils"
)

var producer *queue.Producer

type paymentController struct{}

func Mount(r *mux.Router) {
	controller := paymentController{}
	producer = queue.NewProducer([]string{"localhost:9092"}, "payments")

	r.HandleFunc("/payments", controller.ProcessCustomerPayment).Methods(http.MethodPost)
}

func (controller paymentController) ProcessCustomerPayment(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	defer r.Body.Close()

	var requestBody dto.CreateCustomerPaymentRequest
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		logger.Error("Invalid Request Body: " + err.Error())
		utils.WriteResponse(w, http.StatusBadRequest, err.Error())

		return
	}

	// TODO: Validate requestBody fields

	// Publish to Kafka
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	if err := producer.PublishMessage(ctx, requestBody.TransactionReference, requestBody); err != nil {
		logger.Error("Kafka Error: " + err.Error())
		utils.WriteResponse(w, http.StatusInternalServerError, "Failed to enqueue payment")
		return
	}

	// Return status ok
	utils.WriteResponse(w, http.StatusOK, "Payment Processed")
}
