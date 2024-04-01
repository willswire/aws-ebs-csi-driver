// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/cloud/metadata_interface.go

// Package cloud is a generated GoMock package.
package cloud

import (
	context "context"
	reflect "reflect"

	arn "github.com/aws/aws-sdk-go-v2/aws/arn"
	imds "github.com/aws/aws-sdk-go-v2/feature/ec2/imds"
	gomock "github.com/golang/mock/gomock"
)

// MockMetadataService is a mock of MetadataService interface.
type MockMetadataService struct {
	ctrl     *gomock.Controller
	recorder *MockMetadataServiceMockRecorder
}

// MockMetadataServiceMockRecorder is the mock recorder for MockMetadataService.
type MockMetadataServiceMockRecorder struct {
	mock *MockMetadataService
}

// NewMockMetadataService creates a new mock instance.
func NewMockMetadataService(ctrl *gomock.Controller) *MockMetadataService {
	mock := &MockMetadataService{ctrl: ctrl}
	mock.recorder = &MockMetadataServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMetadataService) EXPECT() *MockMetadataServiceMockRecorder {
	return m.recorder
}

// GetAvailabilityZone mocks base method.
func (m *MockMetadataService) GetAvailabilityZone() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAvailabilityZone")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetAvailabilityZone indicates an expected call of GetAvailabilityZone.
func (mr *MockMetadataServiceMockRecorder) GetAvailabilityZone() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAvailabilityZone", reflect.TypeOf((*MockMetadataService)(nil).GetAvailabilityZone))
}

// GetInstanceID mocks base method.
func (m *MockMetadataService) GetInstanceID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInstanceID")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetInstanceID indicates an expected call of GetInstanceID.
func (mr *MockMetadataServiceMockRecorder) GetInstanceID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInstanceID", reflect.TypeOf((*MockMetadataService)(nil).GetInstanceID))
}

// GetInstanceType mocks base method.
func (m *MockMetadataService) GetInstanceType() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInstanceType")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetInstanceType indicates an expected call of GetInstanceType.
func (mr *MockMetadataServiceMockRecorder) GetInstanceType() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInstanceType", reflect.TypeOf((*MockMetadataService)(nil).GetInstanceType))
}

// GetNumAttachedENIs mocks base method.
func (m *MockMetadataService) GetNumAttachedENIs() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNumAttachedENIs")
	ret0, _ := ret[0].(int)
	return ret0
}

// GetNumAttachedENIs indicates an expected call of GetNumAttachedENIs.
func (mr *MockMetadataServiceMockRecorder) GetNumAttachedENIs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNumAttachedENIs", reflect.TypeOf((*MockMetadataService)(nil).GetNumAttachedENIs))
}

// GetNumBlockDeviceMappings mocks base method.
func (m *MockMetadataService) GetNumBlockDeviceMappings() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNumBlockDeviceMappings")
	ret0, _ := ret[0].(int)
	return ret0
}

// GetNumBlockDeviceMappings indicates an expected call of GetNumBlockDeviceMappings.
func (mr *MockMetadataServiceMockRecorder) GetNumBlockDeviceMappings() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNumBlockDeviceMappings", reflect.TypeOf((*MockMetadataService)(nil).GetNumBlockDeviceMappings))
}

// GetOutpostArn mocks base method.
func (m *MockMetadataService) GetOutpostArn() arn.ARN {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOutpostArn")
	ret0, _ := ret[0].(arn.ARN)
	return ret0
}

// GetOutpostArn indicates an expected call of GetOutpostArn.
func (mr *MockMetadataServiceMockRecorder) GetOutpostArn() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOutpostArn", reflect.TypeOf((*MockMetadataService)(nil).GetOutpostArn))
}

// GetRegion mocks base method.
func (m *MockMetadataService) GetRegion() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRegion")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetRegion indicates an expected call of GetRegion.
func (mr *MockMetadataServiceMockRecorder) GetRegion() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRegion", reflect.TypeOf((*MockMetadataService)(nil).GetRegion))
}

// MockEC2Metadata is a mock of EC2Metadata interface.
type MockEC2Metadata struct {
	ctrl     *gomock.Controller
	recorder *MockEC2MetadataMockRecorder
}

// MockEC2MetadataMockRecorder is the mock recorder for MockEC2Metadata.
type MockEC2MetadataMockRecorder struct {
	mock *MockEC2Metadata
}

// NewMockEC2Metadata creates a new mock instance.
func NewMockEC2Metadata(ctrl *gomock.Controller) *MockEC2Metadata {
	mock := &MockEC2Metadata{ctrl: ctrl}
	mock.recorder = &MockEC2MetadataMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEC2Metadata) EXPECT() *MockEC2MetadataMockRecorder {
	return m.recorder
}

// GetDynamicData mocks base method.
func (m *MockEC2Metadata) GetDynamicData(ctx context.Context, params *imds.GetDynamicDataInput, optFns ...func(*imds.Options)) (*imds.GetDynamicDataOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, params}
	for _, a := range optFns {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetDynamicData", varargs...)
	ret0, _ := ret[0].(*imds.GetDynamicDataOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDynamicData indicates an expected call of GetDynamicData.
func (mr *MockEC2MetadataMockRecorder) GetDynamicData(ctx, params interface{}, optFns ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, params}, optFns...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDynamicData", reflect.TypeOf((*MockEC2Metadata)(nil).GetDynamicData), varargs...)
}

// GetIAMInfo mocks base method.
func (m *MockEC2Metadata) GetIAMInfo(ctx context.Context, params *imds.GetIAMInfoInput, optFns ...func(*imds.Options)) (*imds.GetIAMInfoOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, params}
	for _, a := range optFns {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetIAMInfo", varargs...)
	ret0, _ := ret[0].(*imds.GetIAMInfoOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIAMInfo indicates an expected call of GetIAMInfo.
func (mr *MockEC2MetadataMockRecorder) GetIAMInfo(ctx, params interface{}, optFns ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, params}, optFns...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIAMInfo", reflect.TypeOf((*MockEC2Metadata)(nil).GetIAMInfo), varargs...)
}

// GetInstanceIdentityDocument mocks base method.
func (m *MockEC2Metadata) GetInstanceIdentityDocument(ctx context.Context, params *imds.GetInstanceIdentityDocumentInput, optFns ...func(*imds.Options)) (*imds.GetInstanceIdentityDocumentOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, params}
	for _, a := range optFns {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetInstanceIdentityDocument", varargs...)
	ret0, _ := ret[0].(*imds.GetInstanceIdentityDocumentOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInstanceIdentityDocument indicates an expected call of GetInstanceIdentityDocument.
func (mr *MockEC2MetadataMockRecorder) GetInstanceIdentityDocument(ctx, params interface{}, optFns ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, params}, optFns...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInstanceIdentityDocument", reflect.TypeOf((*MockEC2Metadata)(nil).GetInstanceIdentityDocument), varargs...)
}

// GetMetadata mocks base method.
func (m *MockEC2Metadata) GetMetadata(ctx context.Context, params *imds.GetMetadataInput, optFns ...func(*imds.Options)) (*imds.GetMetadataOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, params}
	for _, a := range optFns {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetMetadata", varargs...)
	ret0, _ := ret[0].(*imds.GetMetadataOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMetadata indicates an expected call of GetMetadata.
func (mr *MockEC2MetadataMockRecorder) GetMetadata(ctx, params interface{}, optFns ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, params}, optFns...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMetadata", reflect.TypeOf((*MockEC2Metadata)(nil).GetMetadata), varargs...)
}

// GetRegion mocks base method.
func (m *MockEC2Metadata) GetRegion(ctx context.Context, params *imds.GetRegionInput, optFns ...func(*imds.Options)) (*imds.GetRegionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, params}
	for _, a := range optFns {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetRegion", varargs...)
	ret0, _ := ret[0].(*imds.GetRegionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRegion indicates an expected call of GetRegion.
func (mr *MockEC2MetadataMockRecorder) GetRegion(ctx, params interface{}, optFns ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, params}, optFns...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRegion", reflect.TypeOf((*MockEC2Metadata)(nil).GetRegion), varargs...)
}

// GetUserData mocks base method.
func (m *MockEC2Metadata) GetUserData(ctx context.Context, params *imds.GetUserDataInput, optFns ...func(*imds.Options)) (*imds.GetUserDataOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, params}
	for _, a := range optFns {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetUserData", varargs...)
	ret0, _ := ret[0].(*imds.GetUserDataOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserData indicates an expected call of GetUserData.
func (mr *MockEC2MetadataMockRecorder) GetUserData(ctx, params interface{}, optFns ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, params}, optFns...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserData", reflect.TypeOf((*MockEC2Metadata)(nil).GetUserData), varargs...)
}
