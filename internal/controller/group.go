package controller

import transfer "github.com/utf2/utf-account-service/internal/controller/model"

type GroupControllerAPI interface {
	Search(transfer.GroupSearchRequest) transfer.GroupSearchResponse
}
