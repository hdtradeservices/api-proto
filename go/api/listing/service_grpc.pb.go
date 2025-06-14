// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package listing

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ListingServiceClient is the client API for ListingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ListingServiceClient interface {
	// Get retrieves a single listing by its ID
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Listing, error)
	// GetBySKU retrieves a single listing which contains a variant with the given
	// SKU
	GetBySKU(ctx context.Context, in *GetBySKURequest, opts ...grpc.CallOption) (*Listing, error)
	// CategoryForSKU retusns the category for a given SKU
	// Since the category is only available on the listing level, this will return
	// the category and potentially save a call to GetBySKU if only the variant is
	// needed
	CategoryForSKU(ctx context.Context, in *CategoryForSKURequest, opts ...grpc.CallOption) (*CategoryForSKUResponse, error)
	// GetVariant retrieves a single variant by its SKU
	GetVariant(ctx context.Context, in *GetVariantRequest, opts ...grpc.CallOption) (*Variant, error)
	// ListNewListings will list any listing created or updated
	// since the given timestamp where:
	//
	// 1. Product data is enabled for at least one Variant in the Listing
	//
	// 2. No variants have a channel ID
	ListNewListings(ctx context.Context, in *ListSinceRequest, opts ...grpc.CallOption) (*ListListingsResponse, error)
	// ListUpdateListings will return any listing that:
	//
	// 1. Has at least one Variant with a channel ID
	//
	// 2. Has a Product Data change since the last timestamp (including Variants)
	//
	// 3. Product Data is enabled for the Listing
	ListUpdatedListings(ctx context.Context, in *ListSinceRequest, opts ...grpc.CallOption) (*ListListingsResponse, error)
	// ListVariantsWithUpdatedInventory will return any variant that:
	//
	// 1. Has an inventory change since the last timestamp
	//
	// 2. Inventory Data is enabled for the Variant
	ListVariantsWithUpdatedInventory(ctx context.Context, in *ListInventorySinceRequest, opts ...grpc.CallOption) (*ListVariantsResponse, error)
	// ListVariantsWithUpdatedPricing will return any variant that:
	//
	// 1. Has a pricing change since the last timestamp
	//
	// 2. Pricing Data is enabled for the Variant
	ListVariantsWithUpdatedPricing(ctx context.Context, in *ListSinceRequest, opts ...grpc.CallOption) (*ListVariantsResponse, error)
	// UpdateStatus updates the status of a listing
	UpdateStatus(ctx context.Context, in *UpdateStatusRequest, opts ...grpc.CallOption) (*UpdateStatusResponse, error)
	// UpdateChannelListingID updates the channel ID for the listing
	UpdateChannelListingID(ctx context.Context, in *UpdateChannelListingIDRequest, opts ...grpc.CallOption) (*UpdateChannelListingIDResponse, error)
	// ReplaceErrors replaces the errors for a variant
	ReplaceErrors(ctx context.Context, in *ReplaceErrorsRequest, opts ...grpc.CallOption) (*ReplaceErrorsResponse, error)
	// CreateSubmissions replaces the submissions for a variant
	// This can be used at any stage, but is most recommended for the intitial
	// creation of a new submission. For updating submissions after creation
	// see UpdateSubmission
	CreateSubmissions(ctx context.Context, in *CreateSubmissionsRequest, opts ...grpc.CallOption) (*CreateSubmissionsResponse, error)
	// UpdateSubmission is used to move a submission to another status
	// this can be done when its actually submitted for ingestion and/or when the
	// ingestion is complete.
	UpdateSubmission(ctx context.Context, in *UpdateSubmissionRequest, opts ...grpc.CallOption) (*Submission, error)
	// SetInventorySubmissionDetails is used to set the inventory details for a
	// given submission
	SetInventorySubmissionDetails(ctx context.Context, in *SetInventorySubmissionDetailsRequest, opts ...grpc.CallOption) (*SetInventorySubmissionDetailsResponse, error)
	// BeginIngestion is used to initiate the ingestion of listings
	// for a given storefront.
	// Needs to be called before RequestIngestion.
	BeginIngestion(ctx context.Context, in *BeginIngestionRequest, opts ...grpc.CallOption) (*BeginIngestionResponse, error)
	// RequestIngestion is used to request ingestion of a listing into Zentail.
	// Need to call BeginIngestion before calling this method.
	RequestIngestion(ctx context.Context, in *RequestIngestionRequest, opts ...grpc.CallOption) (*RequestIngestionResponse, error)
	// EndIngestion is used to end the ingestion of Listings
	// Needs to be called after all Listings requiring ingestion
	// have been requested so that the generated Ingestion Plan can be applied.
	EndIngestion(ctx context.Context, in *EndIngestionRequest, opts ...grpc.CallOption) (*EndIngestionResponse, error)
}

type listingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewListingServiceClient(cc grpc.ClientConnInterface) ListingServiceClient {
	return &listingServiceClient{cc}
}

func (c *listingServiceClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Listing, error) {
	out := new(Listing)
	err := c.cc.Invoke(ctx, "/listing_api.ListingService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *listingServiceClient) GetBySKU(ctx context.Context, in *GetBySKURequest, opts ...grpc.CallOption) (*Listing, error) {
	out := new(Listing)
	err := c.cc.Invoke(ctx, "/listing_api.ListingService/GetBySKU", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *listingServiceClient) CategoryForSKU(ctx context.Context, in *CategoryForSKURequest, opts ...grpc.CallOption) (*CategoryForSKUResponse, error) {
	out := new(CategoryForSKUResponse)
	err := c.cc.Invoke(ctx, "/listing_api.ListingService/CategoryForSKU", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *listingServiceClient) GetVariant(ctx context.Context, in *GetVariantRequest, opts ...grpc.CallOption) (*Variant, error) {
	out := new(Variant)
	err := c.cc.Invoke(ctx, "/listing_api.ListingService/GetVariant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *listingServiceClient) ListNewListings(ctx context.Context, in *ListSinceRequest, opts ...grpc.CallOption) (*ListListingsResponse, error) {
	out := new(ListListingsResponse)
	err := c.cc.Invoke(ctx, "/listing_api.ListingService/ListNewListings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *listingServiceClient) ListUpdatedListings(ctx context.Context, in *ListSinceRequest, opts ...grpc.CallOption) (*ListListingsResponse, error) {
	out := new(ListListingsResponse)
	err := c.cc.Invoke(ctx, "/listing_api.ListingService/ListUpdatedListings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *listingServiceClient) ListVariantsWithUpdatedInventory(ctx context.Context, in *ListInventorySinceRequest, opts ...grpc.CallOption) (*ListVariantsResponse, error) {
	out := new(ListVariantsResponse)
	err := c.cc.Invoke(ctx, "/listing_api.ListingService/ListVariantsWithUpdatedInventory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *listingServiceClient) ListVariantsWithUpdatedPricing(ctx context.Context, in *ListSinceRequest, opts ...grpc.CallOption) (*ListVariantsResponse, error) {
	out := new(ListVariantsResponse)
	err := c.cc.Invoke(ctx, "/listing_api.ListingService/ListVariantsWithUpdatedPricing", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *listingServiceClient) UpdateStatus(ctx context.Context, in *UpdateStatusRequest, opts ...grpc.CallOption) (*UpdateStatusResponse, error) {
	out := new(UpdateStatusResponse)
	err := c.cc.Invoke(ctx, "/listing_api.ListingService/UpdateStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *listingServiceClient) UpdateChannelListingID(ctx context.Context, in *UpdateChannelListingIDRequest, opts ...grpc.CallOption) (*UpdateChannelListingIDResponse, error) {
	out := new(UpdateChannelListingIDResponse)
	err := c.cc.Invoke(ctx, "/listing_api.ListingService/UpdateChannelListingID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *listingServiceClient) ReplaceErrors(ctx context.Context, in *ReplaceErrorsRequest, opts ...grpc.CallOption) (*ReplaceErrorsResponse, error) {
	out := new(ReplaceErrorsResponse)
	err := c.cc.Invoke(ctx, "/listing_api.ListingService/ReplaceErrors", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *listingServiceClient) CreateSubmissions(ctx context.Context, in *CreateSubmissionsRequest, opts ...grpc.CallOption) (*CreateSubmissionsResponse, error) {
	out := new(CreateSubmissionsResponse)
	err := c.cc.Invoke(ctx, "/listing_api.ListingService/CreateSubmissions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *listingServiceClient) UpdateSubmission(ctx context.Context, in *UpdateSubmissionRequest, opts ...grpc.CallOption) (*Submission, error) {
	out := new(Submission)
	err := c.cc.Invoke(ctx, "/listing_api.ListingService/UpdateSubmission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *listingServiceClient) SetInventorySubmissionDetails(ctx context.Context, in *SetInventorySubmissionDetailsRequest, opts ...grpc.CallOption) (*SetInventorySubmissionDetailsResponse, error) {
	out := new(SetInventorySubmissionDetailsResponse)
	err := c.cc.Invoke(ctx, "/listing_api.ListingService/SetInventorySubmissionDetails", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *listingServiceClient) BeginIngestion(ctx context.Context, in *BeginIngestionRequest, opts ...grpc.CallOption) (*BeginIngestionResponse, error) {
	out := new(BeginIngestionResponse)
	err := c.cc.Invoke(ctx, "/listing_api.ListingService/BeginIngestion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *listingServiceClient) RequestIngestion(ctx context.Context, in *RequestIngestionRequest, opts ...grpc.CallOption) (*RequestIngestionResponse, error) {
	out := new(RequestIngestionResponse)
	err := c.cc.Invoke(ctx, "/listing_api.ListingService/RequestIngestion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *listingServiceClient) EndIngestion(ctx context.Context, in *EndIngestionRequest, opts ...grpc.CallOption) (*EndIngestionResponse, error) {
	out := new(EndIngestionResponse)
	err := c.cc.Invoke(ctx, "/listing_api.ListingService/EndIngestion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ListingServiceServer is the server API for ListingService service.
// All implementations should embed UnimplementedListingServiceServer
// for forward compatibility
type ListingServiceServer interface {
	// Get retrieves a single listing by its ID
	Get(context.Context, *GetRequest) (*Listing, error)
	// GetBySKU retrieves a single listing which contains a variant with the given
	// SKU
	GetBySKU(context.Context, *GetBySKURequest) (*Listing, error)
	// CategoryForSKU retusns the category for a given SKU
	// Since the category is only available on the listing level, this will return
	// the category and potentially save a call to GetBySKU if only the variant is
	// needed
	CategoryForSKU(context.Context, *CategoryForSKURequest) (*CategoryForSKUResponse, error)
	// GetVariant retrieves a single variant by its SKU
	GetVariant(context.Context, *GetVariantRequest) (*Variant, error)
	// ListNewListings will list any listing created or updated
	// since the given timestamp where:
	//
	// 1. Product data is enabled for at least one Variant in the Listing
	//
	// 2. No variants have a channel ID
	ListNewListings(context.Context, *ListSinceRequest) (*ListListingsResponse, error)
	// ListUpdateListings will return any listing that:
	//
	// 1. Has at least one Variant with a channel ID
	//
	// 2. Has a Product Data change since the last timestamp (including Variants)
	//
	// 3. Product Data is enabled for the Listing
	ListUpdatedListings(context.Context, *ListSinceRequest) (*ListListingsResponse, error)
	// ListVariantsWithUpdatedInventory will return any variant that:
	//
	// 1. Has an inventory change since the last timestamp
	//
	// 2. Inventory Data is enabled for the Variant
	ListVariantsWithUpdatedInventory(context.Context, *ListInventorySinceRequest) (*ListVariantsResponse, error)
	// ListVariantsWithUpdatedPricing will return any variant that:
	//
	// 1. Has a pricing change since the last timestamp
	//
	// 2. Pricing Data is enabled for the Variant
	ListVariantsWithUpdatedPricing(context.Context, *ListSinceRequest) (*ListVariantsResponse, error)
	// UpdateStatus updates the status of a listing
	UpdateStatus(context.Context, *UpdateStatusRequest) (*UpdateStatusResponse, error)
	// UpdateChannelListingID updates the channel ID for the listing
	UpdateChannelListingID(context.Context, *UpdateChannelListingIDRequest) (*UpdateChannelListingIDResponse, error)
	// ReplaceErrors replaces the errors for a variant
	ReplaceErrors(context.Context, *ReplaceErrorsRequest) (*ReplaceErrorsResponse, error)
	// CreateSubmissions replaces the submissions for a variant
	// This can be used at any stage, but is most recommended for the intitial
	// creation of a new submission. For updating submissions after creation
	// see UpdateSubmission
	CreateSubmissions(context.Context, *CreateSubmissionsRequest) (*CreateSubmissionsResponse, error)
	// UpdateSubmission is used to move a submission to another status
	// this can be done when its actually submitted for ingestion and/or when the
	// ingestion is complete.
	UpdateSubmission(context.Context, *UpdateSubmissionRequest) (*Submission, error)
	// SetInventorySubmissionDetails is used to set the inventory details for a
	// given submission
	SetInventorySubmissionDetails(context.Context, *SetInventorySubmissionDetailsRequest) (*SetInventorySubmissionDetailsResponse, error)
	// BeginIngestion is used to initiate the ingestion of listings
	// for a given storefront.
	// Needs to be called before RequestIngestion.
	BeginIngestion(context.Context, *BeginIngestionRequest) (*BeginIngestionResponse, error)
	// RequestIngestion is used to request ingestion of a listing into Zentail.
	// Need to call BeginIngestion before calling this method.
	RequestIngestion(context.Context, *RequestIngestionRequest) (*RequestIngestionResponse, error)
	// EndIngestion is used to end the ingestion of Listings
	// Needs to be called after all Listings requiring ingestion
	// have been requested so that the generated Ingestion Plan can be applied.
	EndIngestion(context.Context, *EndIngestionRequest) (*EndIngestionResponse, error)
}

// UnimplementedListingServiceServer should be embedded to have forward compatible implementations.
type UnimplementedListingServiceServer struct {
}

func (UnimplementedListingServiceServer) Get(context.Context, *GetRequest) (*Listing, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedListingServiceServer) GetBySKU(context.Context, *GetBySKURequest) (*Listing, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBySKU not implemented")
}
func (UnimplementedListingServiceServer) CategoryForSKU(context.Context, *CategoryForSKURequest) (*CategoryForSKUResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CategoryForSKU not implemented")
}
func (UnimplementedListingServiceServer) GetVariant(context.Context, *GetVariantRequest) (*Variant, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVariant not implemented")
}
func (UnimplementedListingServiceServer) ListNewListings(context.Context, *ListSinceRequest) (*ListListingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListNewListings not implemented")
}
func (UnimplementedListingServiceServer) ListUpdatedListings(context.Context, *ListSinceRequest) (*ListListingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUpdatedListings not implemented")
}
func (UnimplementedListingServiceServer) ListVariantsWithUpdatedInventory(context.Context, *ListInventorySinceRequest) (*ListVariantsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListVariantsWithUpdatedInventory not implemented")
}
func (UnimplementedListingServiceServer) ListVariantsWithUpdatedPricing(context.Context, *ListSinceRequest) (*ListVariantsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListVariantsWithUpdatedPricing not implemented")
}
func (UnimplementedListingServiceServer) UpdateStatus(context.Context, *UpdateStatusRequest) (*UpdateStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateStatus not implemented")
}
func (UnimplementedListingServiceServer) UpdateChannelListingID(context.Context, *UpdateChannelListingIDRequest) (*UpdateChannelListingIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateChannelListingID not implemented")
}
func (UnimplementedListingServiceServer) ReplaceErrors(context.Context, *ReplaceErrorsRequest) (*ReplaceErrorsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReplaceErrors not implemented")
}
func (UnimplementedListingServiceServer) CreateSubmissions(context.Context, *CreateSubmissionsRequest) (*CreateSubmissionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSubmissions not implemented")
}
func (UnimplementedListingServiceServer) UpdateSubmission(context.Context, *UpdateSubmissionRequest) (*Submission, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSubmission not implemented")
}
func (UnimplementedListingServiceServer) SetInventorySubmissionDetails(context.Context, *SetInventorySubmissionDetailsRequest) (*SetInventorySubmissionDetailsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetInventorySubmissionDetails not implemented")
}
func (UnimplementedListingServiceServer) BeginIngestion(context.Context, *BeginIngestionRequest) (*BeginIngestionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BeginIngestion not implemented")
}
func (UnimplementedListingServiceServer) RequestIngestion(context.Context, *RequestIngestionRequest) (*RequestIngestionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestIngestion not implemented")
}
func (UnimplementedListingServiceServer) EndIngestion(context.Context, *EndIngestionRequest) (*EndIngestionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EndIngestion not implemented")
}

// UnsafeListingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ListingServiceServer will
// result in compilation errors.
type UnsafeListingServiceServer interface {
	mustEmbedUnimplementedListingServiceServer()
}

func RegisterListingServiceServer(s grpc.ServiceRegistrar, srv ListingServiceServer) {
	s.RegisterService(&ListingService_ServiceDesc, srv)
}

func _ListingService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListingServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/listing_api.ListingService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListingServiceServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ListingService_GetBySKU_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBySKURequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListingServiceServer).GetBySKU(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/listing_api.ListingService/GetBySKU",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListingServiceServer).GetBySKU(ctx, req.(*GetBySKURequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ListingService_CategoryForSKU_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CategoryForSKURequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListingServiceServer).CategoryForSKU(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/listing_api.ListingService/CategoryForSKU",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListingServiceServer).CategoryForSKU(ctx, req.(*CategoryForSKURequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ListingService_GetVariant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVariantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListingServiceServer).GetVariant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/listing_api.ListingService/GetVariant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListingServiceServer).GetVariant(ctx, req.(*GetVariantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ListingService_ListNewListings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSinceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListingServiceServer).ListNewListings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/listing_api.ListingService/ListNewListings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListingServiceServer).ListNewListings(ctx, req.(*ListSinceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ListingService_ListUpdatedListings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSinceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListingServiceServer).ListUpdatedListings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/listing_api.ListingService/ListUpdatedListings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListingServiceServer).ListUpdatedListings(ctx, req.(*ListSinceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ListingService_ListVariantsWithUpdatedInventory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListInventorySinceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListingServiceServer).ListVariantsWithUpdatedInventory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/listing_api.ListingService/ListVariantsWithUpdatedInventory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListingServiceServer).ListVariantsWithUpdatedInventory(ctx, req.(*ListInventorySinceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ListingService_ListVariantsWithUpdatedPricing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSinceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListingServiceServer).ListVariantsWithUpdatedPricing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/listing_api.ListingService/ListVariantsWithUpdatedPricing",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListingServiceServer).ListVariantsWithUpdatedPricing(ctx, req.(*ListSinceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ListingService_UpdateStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListingServiceServer).UpdateStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/listing_api.ListingService/UpdateStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListingServiceServer).UpdateStatus(ctx, req.(*UpdateStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ListingService_UpdateChannelListingID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateChannelListingIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListingServiceServer).UpdateChannelListingID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/listing_api.ListingService/UpdateChannelListingID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListingServiceServer).UpdateChannelListingID(ctx, req.(*UpdateChannelListingIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ListingService_ReplaceErrors_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReplaceErrorsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListingServiceServer).ReplaceErrors(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/listing_api.ListingService/ReplaceErrors",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListingServiceServer).ReplaceErrors(ctx, req.(*ReplaceErrorsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ListingService_CreateSubmissions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSubmissionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListingServiceServer).CreateSubmissions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/listing_api.ListingService/CreateSubmissions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListingServiceServer).CreateSubmissions(ctx, req.(*CreateSubmissionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ListingService_UpdateSubmission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSubmissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListingServiceServer).UpdateSubmission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/listing_api.ListingService/UpdateSubmission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListingServiceServer).UpdateSubmission(ctx, req.(*UpdateSubmissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ListingService_SetInventorySubmissionDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetInventorySubmissionDetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListingServiceServer).SetInventorySubmissionDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/listing_api.ListingService/SetInventorySubmissionDetails",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListingServiceServer).SetInventorySubmissionDetails(ctx, req.(*SetInventorySubmissionDetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ListingService_BeginIngestion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BeginIngestionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListingServiceServer).BeginIngestion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/listing_api.ListingService/BeginIngestion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListingServiceServer).BeginIngestion(ctx, req.(*BeginIngestionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ListingService_RequestIngestion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestIngestionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListingServiceServer).RequestIngestion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/listing_api.ListingService/RequestIngestion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListingServiceServer).RequestIngestion(ctx, req.(*RequestIngestionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ListingService_EndIngestion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EndIngestionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListingServiceServer).EndIngestion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/listing_api.ListingService/EndIngestion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListingServiceServer).EndIngestion(ctx, req.(*EndIngestionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ListingService_ServiceDesc is the grpc.ServiceDesc for ListingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ListingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "listing_api.ListingService",
	HandlerType: (*ListingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _ListingService_Get_Handler,
		},
		{
			MethodName: "GetBySKU",
			Handler:    _ListingService_GetBySKU_Handler,
		},
		{
			MethodName: "CategoryForSKU",
			Handler:    _ListingService_CategoryForSKU_Handler,
		},
		{
			MethodName: "GetVariant",
			Handler:    _ListingService_GetVariant_Handler,
		},
		{
			MethodName: "ListNewListings",
			Handler:    _ListingService_ListNewListings_Handler,
		},
		{
			MethodName: "ListUpdatedListings",
			Handler:    _ListingService_ListUpdatedListings_Handler,
		},
		{
			MethodName: "ListVariantsWithUpdatedInventory",
			Handler:    _ListingService_ListVariantsWithUpdatedInventory_Handler,
		},
		{
			MethodName: "ListVariantsWithUpdatedPricing",
			Handler:    _ListingService_ListVariantsWithUpdatedPricing_Handler,
		},
		{
			MethodName: "UpdateStatus",
			Handler:    _ListingService_UpdateStatus_Handler,
		},
		{
			MethodName: "UpdateChannelListingID",
			Handler:    _ListingService_UpdateChannelListingID_Handler,
		},
		{
			MethodName: "ReplaceErrors",
			Handler:    _ListingService_ReplaceErrors_Handler,
		},
		{
			MethodName: "CreateSubmissions",
			Handler:    _ListingService_CreateSubmissions_Handler,
		},
		{
			MethodName: "UpdateSubmission",
			Handler:    _ListingService_UpdateSubmission_Handler,
		},
		{
			MethodName: "SetInventorySubmissionDetails",
			Handler:    _ListingService_SetInventorySubmissionDetails_Handler,
		},
		{
			MethodName: "BeginIngestion",
			Handler:    _ListingService_BeginIngestion_Handler,
		},
		{
			MethodName: "RequestIngestion",
			Handler:    _ListingService_RequestIngestion_Handler,
		},
		{
			MethodName: "EndIngestion",
			Handler:    _ListingService_EndIngestion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/listing/service.proto",
}
