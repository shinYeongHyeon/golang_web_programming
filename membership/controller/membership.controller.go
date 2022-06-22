package membershipController

import (
	"github.com/labstack/echo"
	membershipControllerDto "github.com/shinYeongHyeon/golang_web_programming/membership/controller/dto"
	membershipServices "github.com/shinYeongHyeon/golang_web_programming/membership/services"
	membershipServicesDto "github.com/shinYeongHyeon/golang_web_programming/membership/services/dto"
	"net/http"
)

type Controller struct {
	service membershipServices.Service
}

func NewController(service membershipServices.Service) *Controller {
	return &Controller{service: service}
}

func (c Controller) CreateMembership(echoCtx echo.Context) error {
	req := new(membershipControllerDto.CreateRequest)
	if err := echoCtx.Bind(req); err != nil {
		return respondCreateRequestFailed(echoCtx, http.StatusBadRequest)
	}

	res := getCreateResponse(c, req)
	if res.Code != membershipServices.SuccessCreate {
		return respondCreateFailed(echoCtx, res)
	}

	return echoCtx.JSON(http.StatusCreated, membershipControllerDto.CreateResponse{
		Code:           http.StatusCreated,
		Message:        http.StatusText(http.StatusCreated),
		MembershipUser: generateMembershipUserDto(res.MembershipUser),
	})
}

func getCreateResponse(c Controller, req *membershipControllerDto.CreateRequest) membershipServicesDto.CreateResponse {
	return c.service.Create(membershipServicesDto.CreateRequest{
		UserName:       req.UserName,
		MembershipType: req.MembershipType,
	})
}

func respondCreateFailed(echoCtx echo.Context, res membershipServicesDto.CreateResponse) error {
	if res.Code == membershipServices.DuplicateUserName {
		return respondCreateRequestFailed(echoCtx, http.StatusConflict)
	}

	if res.Code == membershipServices.NotAvailableMembershipType {
		return respondCreateRequestFailed(echoCtx, http.StatusBadRequest)
	}

	return respondCreateRequestFailed(echoCtx, http.StatusInternalServerError)
}

func respondCreateRequestFailed(echoCtx echo.Context, statusCode int) error {
	return echoCtx.JSON(statusCode, membershipControllerDto.CreateResponse{
		Code:    statusCode,
		Message: http.StatusText(statusCode),
	})
}

func (c Controller) FindAll(echoCtx echo.Context) error {
	foundMemberships := c.service.FindAll(membershipServicesDto.FindAllRequest{
		Offset: echoCtx.QueryParam("offset"),
		Limit:  echoCtx.QueryParam("limit"),
	})

	var memberships []*membershipControllerDto.MembershipUser
	for _, membership := range foundMemberships.MembershipUsers {
		memberships = append(memberships, generateMembershipUserDto(membership))
	}

	return echoCtx.JSON(http.StatusOK, membershipControllerDto.FindALlResponse{
		Code:            http.StatusOK,
		Message:         http.StatusText(http.StatusOK),
		MembershipUsers: memberships,
	})
}

func (c Controller) FindOne(echoCtx echo.Context) error {
	res := c.service.FineOne(echoCtx.Param("id"))

	if res.Code == membershipServices.NotFoundMembership {
		return echoCtx.JSON(http.StatusNotFound, membershipControllerDto.FindResponse{
			Code:    http.StatusNotFound,
			Message: http.StatusText(http.StatusNotFound),
		})
	}

	return echoCtx.JSON(http.StatusOK, membershipControllerDto.FindResponse{
		Code:           http.StatusOK,
		Message:        http.StatusText(http.StatusOK),
		MembershipUser: generateMembershipUserDto(res.MembershipUser),
	})
}

func (c Controller) UpdateOne(echoCtx echo.Context) error {
	req := new(membershipControllerDto.UpdateRequest)
	if err := echoCtx.Bind(req); err != nil {
		return respondUpdateRequestFailed(echoCtx, http.StatusBadRequest)
	}

	res := getUpdateResponse(c, req, echoCtx.Param("id"))
	if res.Code != membershipServices.SuccessUpdate {
		return respondUpdateFailed(echoCtx, res)
	}

	return echoCtx.JSON(http.StatusOK, membershipControllerDto.UpdateResponse{
		Code:           http.StatusOK,
		Message:        http.StatusText(http.StatusOK),
		MembershipUser: generateMembershipUserDto(res.MembershipUser),
	})
}

func respondUpdateRequestFailed(echoCtx echo.Context, statusCode int) error {
	return echoCtx.JSON(statusCode, membershipControllerDto.UpdateResponse{
		Code:    statusCode,
		Message: http.StatusText(statusCode),
	})
}

func getUpdateResponse(c Controller, req *membershipControllerDto.UpdateRequest, id string) membershipServicesDto.UpdateResponse {
	res := c.service.UpdateOne(membershipServicesDto.UpdateRequest{
		ID:             id,
		UserName:       req.UserName,
		MembershipType: req.MembershipType,
	})
	return res
}

func respondUpdateFailed(echoCtx echo.Context, res membershipServicesDto.UpdateResponse) error {
	if res.Code == membershipServices.DuplicateUserName {
		return respondUpdateRequestFailed(echoCtx, http.StatusConflict)
	}

	if res.Code == membershipServices.NotAvailableMembershipType {
		return respondUpdateRequestFailed(echoCtx, http.StatusBadRequest)
	}

	if res.Code == membershipServices.NotFoundMembership {
		return respondUpdateRequestFailed(echoCtx, http.StatusNotFound)
	}

	return respondUpdateRequestFailed(echoCtx, http.StatusInternalServerError)
}

func (c Controller) DeleteOne(echoCtx echo.Context) error {
	id := echoCtx.Param("id")
	if id == "" {
		return respondRequestFailed(echoCtx, http.StatusBadRequest)
	}

	res := c.service.DeleteOne(id)
	if res.Code == membershipServices.NotFoundMembership {
		return respondRequestFailed(echoCtx, http.StatusNotFound)
	}

	return echoCtx.JSON(http.StatusOK, membershipControllerDto.UpdateResponse{
		Code:    http.StatusOK,
		Message: http.StatusText(http.StatusOK),
	})
}

func respondRequestFailed(echoCtx echo.Context, statusCode int) error {
	return echoCtx.JSON(statusCode, membershipControllerDto.DeleteResponse{
		Code:    statusCode,
		Message: http.StatusText(statusCode),
	})
}

func generateMembershipUserDto(membership *membershipServicesDto.MembershipUser) *membershipControllerDto.MembershipUser {
	return &membershipControllerDto.MembershipUser{
		ID:             membership.ID,
		UserName:       membership.UserName,
		MembershipType: membership.MembershipType,
	}
}
