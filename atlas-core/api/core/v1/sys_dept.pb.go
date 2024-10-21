// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.2
// source: core/v1/sys_dept.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SysDeptRep struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeptId   string `protobuf:"bytes,1,opt,name=dept_id,json=deptId,proto3" json:"dept_id"`
	ParentId string `protobuf:"bytes,2,opt,name=parentId,proto3" json:"parentId"`
	DeptName string `protobuf:"bytes,3,opt,name=deptName,proto3" json:"deptName"`
	OrderNum int64  `protobuf:"varint,4,opt,name=orderNum,proto3" json:"orderNum"`
	Leader   string `protobuf:"bytes,5,opt,name=leader,proto3" json:"leader"`
	Phone    string `protobuf:"bytes,6,opt,name=phone,proto3" json:"phone"`
	Email    string `protobuf:"bytes,7,opt,name=email,proto3" json:"email"`
	Status   string `protobuf:"bytes,8,opt,name=status,proto3" json:"status"`
}

func (x *SysDeptRep) Reset() {
	*x = SysDeptRep{}
	mi := &file_core_v1_sys_dept_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SysDeptRep) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SysDeptRep) ProtoMessage() {}

func (x *SysDeptRep) ProtoReflect() protoreflect.Message {
	mi := &file_core_v1_sys_dept_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SysDeptRep.ProtoReflect.Descriptor instead.
func (*SysDeptRep) Descriptor() ([]byte, []int) {
	return file_core_v1_sys_dept_proto_rawDescGZIP(), []int{0}
}

func (x *SysDeptRep) GetDeptId() string {
	if x != nil {
		return x.DeptId
	}
	return ""
}

func (x *SysDeptRep) GetParentId() string {
	if x != nil {
		return x.ParentId
	}
	return ""
}

func (x *SysDeptRep) GetDeptName() string {
	if x != nil {
		return x.DeptName
	}
	return ""
}

func (x *SysDeptRep) GetOrderNum() int64 {
	if x != nil {
		return x.OrderNum
	}
	return 0
}

func (x *SysDeptRep) GetLeader() string {
	if x != nil {
		return x.Leader
	}
	return ""
}

func (x *SysDeptRep) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *SysDeptRep) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *SysDeptRep) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type DeleteSysDeptRep struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
}

func (x *DeleteSysDeptRep) Reset() {
	*x = DeleteSysDeptRep{}
	mi := &file_core_v1_sys_dept_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteSysDeptRep) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSysDeptRep) ProtoMessage() {}

func (x *DeleteSysDeptRep) ProtoReflect() protoreflect.Message {
	mi := &file_core_v1_sys_dept_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSysDeptRep.ProtoReflect.Descriptor instead.
func (*DeleteSysDeptRep) Descriptor() ([]byte, []int) {
	return file_core_v1_sys_dept_proto_rawDescGZIP(), []int{1}
}

func (x *DeleteSysDeptRep) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetSysDeptRep struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
}

func (x *GetSysDeptRep) Reset() {
	*x = GetSysDeptRep{}
	mi := &file_core_v1_sys_dept_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetSysDeptRep) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSysDeptRep) ProtoMessage() {}

func (x *GetSysDeptRep) ProtoReflect() protoreflect.Message {
	mi := &file_core_v1_sys_dept_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSysDeptRep.ProtoReflect.Descriptor instead.
func (*GetSysDeptRep) Descriptor() ([]byte, []int) {
	return file_core_v1_sys_dept_proto_rawDescGZIP(), []int{2}
}

func (x *GetSysDeptRep) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetSysDeptReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dept *DeptReply `protobuf:"bytes,1,opt,name=dept,proto3" json:"dept"`
}

func (x *GetSysDeptReply) Reset() {
	*x = GetSysDeptReply{}
	mi := &file_core_v1_sys_dept_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetSysDeptReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSysDeptReply) ProtoMessage() {}

func (x *GetSysDeptReply) ProtoReflect() protoreflect.Message {
	mi := &file_core_v1_sys_dept_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSysDeptReply.ProtoReflect.Descriptor instead.
func (*GetSysDeptReply) Descriptor() ([]byte, []int) {
	return file_core_v1_sys_dept_proto_rawDescGZIP(), []int{3}
}

func (x *GetSysDeptReply) GetDept() *DeptReply {
	if x != nil {
		return x.Dept
	}
	return nil
}

type ExcludeDeptRep struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
}

func (x *ExcludeDeptRep) Reset() {
	*x = ExcludeDeptRep{}
	mi := &file_core_v1_sys_dept_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExcludeDeptRep) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExcludeDeptRep) ProtoMessage() {}

func (x *ExcludeDeptRep) ProtoReflect() protoreflect.Message {
	mi := &file_core_v1_sys_dept_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExcludeDeptRep.ProtoReflect.Descriptor instead.
func (*ExcludeDeptRep) Descriptor() ([]byte, []int) {
	return file_core_v1_sys_dept_proto_rawDescGZIP(), []int{4}
}

func (x *ExcludeDeptRep) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeptTreeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeptTreeReq) Reset() {
	*x = DeptTreeReq{}
	mi := &file_core_v1_sys_dept_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeptTreeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeptTreeReq) ProtoMessage() {}

func (x *DeptTreeReq) ProtoReflect() protoreflect.Message {
	mi := &file_core_v1_sys_dept_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeptTreeReq.ProtoReflect.Descriptor instead.
func (*DeptTreeReq) Descriptor() ([]byte, []int) {
	return file_core_v1_sys_dept_proto_rawDescGZIP(), []int{5}
}

type DeptTreeReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dept []*DeptTree `protobuf:"bytes,1,rep,name=dept,proto3" json:"dept"`
}

func (x *DeptTreeReply) Reset() {
	*x = DeptTreeReply{}
	mi := &file_core_v1_sys_dept_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeptTreeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeptTreeReply) ProtoMessage() {}

func (x *DeptTreeReply) ProtoReflect() protoreflect.Message {
	mi := &file_core_v1_sys_dept_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeptTreeReply.ProtoReflect.Descriptor instead.
func (*DeptTreeReply) Descriptor() ([]byte, []int) {
	return file_core_v1_sys_dept_proto_rawDescGZIP(), []int{6}
}

func (x *DeptTreeReply) GetDept() []*DeptTree {
	if x != nil {
		return x.Dept
	}
	return nil
}

type DeptTree struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string      `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	Pid      string      `protobuf:"bytes,2,opt,name=pid,proto3" json:"pid"`
	Label    string      `protobuf:"bytes,3,opt,name=label,proto3" json:"label"`
	Children []*DeptTree `protobuf:"bytes,4,rep,name=children,proto3" json:"children"`
}

func (x *DeptTree) Reset() {
	*x = DeptTree{}
	mi := &file_core_v1_sys_dept_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeptTree) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeptTree) ProtoMessage() {}

func (x *DeptTree) ProtoReflect() protoreflect.Message {
	mi := &file_core_v1_sys_dept_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeptTree.ProtoReflect.Descriptor instead.
func (*DeptTree) Descriptor() ([]byte, []int) {
	return file_core_v1_sys_dept_proto_rawDescGZIP(), []int{7}
}

func (x *DeptTree) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DeptTree) GetPid() string {
	if x != nil {
		return x.Pid
	}
	return ""
}

func (x *DeptTree) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *DeptTree) GetChildren() []*DeptTree {
	if x != nil {
		return x.Children
	}
	return nil
}

type ListSysDeptRep struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListSysDeptRep) Reset() {
	*x = ListSysDeptRep{}
	mi := &file_core_v1_sys_dept_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListSysDeptRep) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSysDeptRep) ProtoMessage() {}

func (x *ListSysDeptRep) ProtoReflect() protoreflect.Message {
	mi := &file_core_v1_sys_dept_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSysDeptRep.ProtoReflect.Descriptor instead.
func (*ListSysDeptRep) Descriptor() ([]byte, []int) {
	return file_core_v1_sys_dept_proto_rawDescGZIP(), []int{8}
}

type ListSysDeptReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dept []*DeptReply `protobuf:"bytes,1,rep,name=dept,proto3" json:"dept"`
}

func (x *ListSysDeptReply) Reset() {
	*x = ListSysDeptReply{}
	mi := &file_core_v1_sys_dept_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListSysDeptReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSysDeptReply) ProtoMessage() {}

func (x *ListSysDeptReply) ProtoReflect() protoreflect.Message {
	mi := &file_core_v1_sys_dept_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSysDeptReply.ProtoReflect.Descriptor instead.
func (*ListSysDeptReply) Descriptor() ([]byte, []int) {
	return file_core_v1_sys_dept_proto_rawDescGZIP(), []int{9}
}

func (x *ListSysDeptReply) GetDept() []*DeptReply {
	if x != nil {
		return x.Dept
	}
	return nil
}

type GetSysRoleDeptRep struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoleId string `protobuf:"bytes,1,opt,name=roleId,proto3" json:"roleId"`
}

func (x *GetSysRoleDeptRep) Reset() {
	*x = GetSysRoleDeptRep{}
	mi := &file_core_v1_sys_dept_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetSysRoleDeptRep) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSysRoleDeptRep) ProtoMessage() {}

func (x *GetSysRoleDeptRep) ProtoReflect() protoreflect.Message {
	mi := &file_core_v1_sys_dept_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSysRoleDeptRep.ProtoReflect.Descriptor instead.
func (*GetSysRoleDeptRep) Descriptor() ([]byte, []int) {
	return file_core_v1_sys_dept_proto_rawDescGZIP(), []int{10}
}

func (x *GetSysRoleDeptRep) GetRoleId() string {
	if x != nil {
		return x.RoleId
	}
	return ""
}

type GetSysRoleDeptReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dept        []*DeptTree `protobuf:"bytes,1,rep,name=dept,proto3" json:"dept"`
	CheckedKeys []string    `protobuf:"bytes,2,rep,name=checkedKeys,proto3" json:"checkedKeys"`
}

func (x *GetSysRoleDeptReply) Reset() {
	*x = GetSysRoleDeptReply{}
	mi := &file_core_v1_sys_dept_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetSysRoleDeptReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSysRoleDeptReply) ProtoMessage() {}

func (x *GetSysRoleDeptReply) ProtoReflect() protoreflect.Message {
	mi := &file_core_v1_sys_dept_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSysRoleDeptReply.ProtoReflect.Descriptor instead.
func (*GetSysRoleDeptReply) Descriptor() ([]byte, []int) {
	return file_core_v1_sys_dept_proto_rawDescGZIP(), []int{11}
}

func (x *GetSysRoleDeptReply) GetDept() []*DeptTree {
	if x != nil {
		return x.Dept
	}
	return nil
}

func (x *GetSysRoleDeptReply) GetCheckedKeys() []string {
	if x != nil {
		return x.CheckedKeys
	}
	return nil
}

var File_core_v1_sys_dept_proto protoreflect.FileDescriptor

var file_core_v1_sys_dept_proto_rawDesc = []byte{
	0x0a, 0x16, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x79, 0x73, 0x5f, 0x64, 0x65,
	0x70, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12,
	0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xd5, 0x01, 0x0a, 0x0a, 0x53, 0x79, 0x73, 0x44, 0x65, 0x70, 0x74, 0x52, 0x65,
	0x70, 0x12, 0x17, 0x0a, 0x07, 0x64, 0x65, 0x70, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x64, 0x65, 0x70, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x70, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x70, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x4e, 0x75, 0x6d, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x4e, 0x75, 0x6d, 0x12, 0x16,
	0x0a, 0x06, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x22, 0x0a, 0x10, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x53, 0x79, 0x73, 0x44, 0x65, 0x70, 0x74, 0x52, 0x65, 0x70, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x28,
	0x0a, 0x0d, 0x47, 0x65, 0x74, 0x53, 0x79, 0x73, 0x44, 0x65, 0x70, 0x74, 0x52, 0x65, 0x70, 0x12,
	0x17, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x72, 0x02, 0x10, 0x01, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3f, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x53,
	0x79, 0x73, 0x44, 0x65, 0x70, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2c, 0x0a, 0x04, 0x64,
	0x65, 0x70, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x70, 0x74, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x52, 0x04, 0x64, 0x65, 0x70, 0x74, 0x22, 0x29, 0x0a, 0x0e, 0x45, 0x78, 0x63,
	0x6c, 0x75, 0x64, 0x65, 0x44, 0x65, 0x70, 0x74, 0x52, 0x65, 0x70, 0x12, 0x17, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x0d, 0x0a, 0x0b, 0x44, 0x65, 0x70, 0x74, 0x54, 0x72, 0x65, 0x65,
	0x52, 0x65, 0x71, 0x22, 0x3c, 0x0a, 0x0d, 0x44, 0x65, 0x70, 0x74, 0x54, 0x72, 0x65, 0x65, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x2b, 0x0a, 0x04, 0x64, 0x65, 0x70, 0x74, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x65, 0x70, 0x74, 0x54, 0x72, 0x65, 0x65, 0x52, 0x04, 0x64, 0x65, 0x70,
	0x74, 0x22, 0x77, 0x0a, 0x08, 0x44, 0x65, 0x70, 0x74, 0x54, 0x72, 0x65, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a,
	0x03, 0x70, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x70, 0x69, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x33, 0x0a, 0x08, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65,
	0x6e, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x70, 0x74, 0x54, 0x72, 0x65, 0x65,
	0x52, 0x08, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x22, 0x10, 0x0a, 0x0e, 0x4c, 0x69,
	0x73, 0x74, 0x53, 0x79, 0x73, 0x44, 0x65, 0x70, 0x74, 0x52, 0x65, 0x70, 0x22, 0x40, 0x0a, 0x10,
	0x4c, 0x69, 0x73, 0x74, 0x53, 0x79, 0x73, 0x44, 0x65, 0x70, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x2c, 0x0a, 0x04, 0x64, 0x65, 0x70, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x44,
	0x65, 0x70, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x52, 0x04, 0x64, 0x65, 0x70, 0x74, 0x22, 0x2b,
	0x0a, 0x11, 0x47, 0x65, 0x74, 0x53, 0x79, 0x73, 0x52, 0x6f, 0x6c, 0x65, 0x44, 0x65, 0x70, 0x74,
	0x52, 0x65, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x22, 0x64, 0x0a, 0x13, 0x47,
	0x65, 0x74, 0x53, 0x79, 0x73, 0x52, 0x6f, 0x6c, 0x65, 0x44, 0x65, 0x70, 0x74, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x2b, 0x0a, 0x04, 0x64, 0x65, 0x70, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31,
	0x2e, 0x44, 0x65, 0x70, 0x74, 0x54, 0x72, 0x65, 0x65, 0x52, 0x04, 0x64, 0x65, 0x70, 0x74, 0x12,
	0x20, 0x0a, 0x0b, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x64, 0x4b, 0x65, 0x79, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x64, 0x4b, 0x65, 0x79,
	0x73, 0x32, 0xe2, 0x06, 0x0a, 0x07, 0x53, 0x79, 0x73, 0x44, 0x65, 0x70, 0x74, 0x12, 0x5f, 0x0a,
	0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x79, 0x73, 0x44, 0x65, 0x70, 0x74, 0x12, 0x19,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x79, 0x73, 0x44, 0x65, 0x70, 0x74, 0x52, 0x65, 0x70, 0x1a, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x3a, 0x01, 0x2a, 0x22,
	0x0d, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x70, 0x74, 0x2f, 0x73, 0x61, 0x76, 0x65, 0x12, 0x61,
	0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x79, 0x73, 0x44, 0x65, 0x70, 0x74, 0x12,
	0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x79, 0x73, 0x44, 0x65, 0x70, 0x74, 0x52, 0x65, 0x70, 0x1a, 0x19, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x3a, 0x01, 0x2a,
	0x1a, 0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x70, 0x74, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x12, 0x69, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x79, 0x73, 0x44, 0x65,
	0x70, 0x74, 0x12, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x79, 0x73, 0x44, 0x65, 0x70, 0x74,
	0x52, 0x65, 0x70, 0x1a, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1c,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x2a, 0x14, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x70, 0x74,
	0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x66, 0x0a, 0x0a,
	0x47, 0x65, 0x74, 0x53, 0x79, 0x73, 0x44, 0x65, 0x70, 0x74, 0x12, 0x1c, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x79,
	0x73, 0x44, 0x65, 0x70, 0x74, 0x52, 0x65, 0x70, 0x1a, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x79, 0x73, 0x44,
	0x65, 0x70, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14,
	0x12, 0x12, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x70, 0x74, 0x2f, 0x69, 0x6e, 0x66, 0x6f, 0x2f,
	0x7b, 0x69, 0x64, 0x7d, 0x12, 0x67, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x79, 0x73, 0x44,
	0x65, 0x70, 0x74, 0x12, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x79, 0x73, 0x44, 0x65, 0x70, 0x74, 0x52,
	0x65, 0x70, 0x1a, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x79, 0x73, 0x44, 0x65, 0x70, 0x74, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x3a, 0x01, 0x2a, 0x22, 0x0d,
	0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x70, 0x74, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x71, 0x0a,
	0x0b, 0x45, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x44, 0x65, 0x70, 0x74, 0x12, 0x1d, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x63,
	0x6c, 0x75, 0x64, 0x65, 0x44, 0x65, 0x70, 0x74, 0x52, 0x65, 0x70, 0x1a, 0x1f, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x53, 0x79, 0x73, 0x44, 0x65, 0x70, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x22, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x1c, 0x12, 0x1a, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x70, 0x74, 0x2f, 0x6c,
	0x69, 0x73, 0x74, 0x2f, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x2f, 0x7b, 0x69, 0x64, 0x7d,
	0x12, 0x5f, 0x0a, 0x08, 0x44, 0x65, 0x70, 0x74, 0x54, 0x72, 0x65, 0x65, 0x12, 0x1a, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x70,
	0x74, 0x54, 0x72, 0x65, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x70, 0x74, 0x54, 0x72, 0x65,
	0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x12, 0x11,
	0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x70, 0x74, 0x2f, 0x64, 0x65, 0x70, 0x74, 0x54, 0x72, 0x65,
	0x65, 0x12, 0x82, 0x01, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x53, 0x79, 0x73, 0x52, 0x6f, 0x6c, 0x65,
	0x44, 0x65, 0x70, 0x74, 0x12, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x79, 0x73, 0x52, 0x6f, 0x6c, 0x65, 0x44,
	0x65, 0x70, 0x74, 0x52, 0x65, 0x70, 0x1a, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x79, 0x73, 0x52, 0x6f, 0x6c,
	0x65, 0x44, 0x65, 0x70, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x2a, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x24, 0x12, 0x22, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x70, 0x74, 0x2f, 0x64, 0x65, 0x70,
	0x74, 0x54, 0x72, 0x65, 0x65, 0x42, 0x79, 0x52, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x2f, 0x7b, 0x72,
	0x6f, 0x6c, 0x65, 0x49, 0x64, 0x7d, 0x42, 0x2a, 0x0a, 0x0b, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x19, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2d, 0x63,
	0x6f, 0x72, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x3b,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_core_v1_sys_dept_proto_rawDescOnce sync.Once
	file_core_v1_sys_dept_proto_rawDescData = file_core_v1_sys_dept_proto_rawDesc
)

func file_core_v1_sys_dept_proto_rawDescGZIP() []byte {
	file_core_v1_sys_dept_proto_rawDescOnce.Do(func() {
		file_core_v1_sys_dept_proto_rawDescData = protoimpl.X.CompressGZIP(file_core_v1_sys_dept_proto_rawDescData)
	})
	return file_core_v1_sys_dept_proto_rawDescData
}

var file_core_v1_sys_dept_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_core_v1_sys_dept_proto_goTypes = []any{
	(*SysDeptRep)(nil),          // 0: api.system.v1.SysDeptRep
	(*DeleteSysDeptRep)(nil),    // 1: api.system.v1.DeleteSysDeptRep
	(*GetSysDeptRep)(nil),       // 2: api.system.v1.GetSysDeptRep
	(*GetSysDeptReply)(nil),     // 3: api.system.v1.GetSysDeptReply
	(*ExcludeDeptRep)(nil),      // 4: api.system.v1.ExcludeDeptRep
	(*DeptTreeReq)(nil),         // 5: api.system.v1.DeptTreeReq
	(*DeptTreeReply)(nil),       // 6: api.system.v1.DeptTreeReply
	(*DeptTree)(nil),            // 7: api.system.v1.DeptTree
	(*ListSysDeptRep)(nil),      // 8: api.system.v1.ListSysDeptRep
	(*ListSysDeptReply)(nil),    // 9: api.system.v1.ListSysDeptReply
	(*GetSysRoleDeptRep)(nil),   // 10: api.system.v1.GetSysRoleDeptRep
	(*GetSysRoleDeptReply)(nil), // 11: api.system.v1.GetSysRoleDeptReply
	(*DeptReply)(nil),           // 12: api.system.v1.DeptReply
	(*EmptyReply)(nil),          // 13: api.system.v1.EmptyReply
}
var file_core_v1_sys_dept_proto_depIdxs = []int32{
	12, // 0: api.system.v1.GetSysDeptReply.dept:type_name -> api.system.v1.DeptReply
	7,  // 1: api.system.v1.DeptTreeReply.dept:type_name -> api.system.v1.DeptTree
	7,  // 2: api.system.v1.DeptTree.children:type_name -> api.system.v1.DeptTree
	12, // 3: api.system.v1.ListSysDeptReply.dept:type_name -> api.system.v1.DeptReply
	7,  // 4: api.system.v1.GetSysRoleDeptReply.dept:type_name -> api.system.v1.DeptTree
	0,  // 5: api.system.v1.SysDept.CreateSysDept:input_type -> api.system.v1.SysDeptRep
	0,  // 6: api.system.v1.SysDept.UpdateSysDept:input_type -> api.system.v1.SysDeptRep
	1,  // 7: api.system.v1.SysDept.DeleteSysDept:input_type -> api.system.v1.DeleteSysDeptRep
	2,  // 8: api.system.v1.SysDept.GetSysDept:input_type -> api.system.v1.GetSysDeptRep
	8,  // 9: api.system.v1.SysDept.ListSysDept:input_type -> api.system.v1.ListSysDeptRep
	4,  // 10: api.system.v1.SysDept.ExcludeDept:input_type -> api.system.v1.ExcludeDeptRep
	5,  // 11: api.system.v1.SysDept.DeptTree:input_type -> api.system.v1.DeptTreeReq
	10, // 12: api.system.v1.SysDept.GetSysRoleDept:input_type -> api.system.v1.GetSysRoleDeptRep
	13, // 13: api.system.v1.SysDept.CreateSysDept:output_type -> api.system.v1.EmptyReply
	13, // 14: api.system.v1.SysDept.UpdateSysDept:output_type -> api.system.v1.EmptyReply
	13, // 15: api.system.v1.SysDept.DeleteSysDept:output_type -> api.system.v1.EmptyReply
	3,  // 16: api.system.v1.SysDept.GetSysDept:output_type -> api.system.v1.GetSysDeptReply
	9,  // 17: api.system.v1.SysDept.ListSysDept:output_type -> api.system.v1.ListSysDeptReply
	9,  // 18: api.system.v1.SysDept.ExcludeDept:output_type -> api.system.v1.ListSysDeptReply
	6,  // 19: api.system.v1.SysDept.DeptTree:output_type -> api.system.v1.DeptTreeReply
	11, // 20: api.system.v1.SysDept.GetSysRoleDept:output_type -> api.system.v1.GetSysRoleDeptReply
	13, // [13:21] is the sub-list for method output_type
	5,  // [5:13] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_core_v1_sys_dept_proto_init() }
func file_core_v1_sys_dept_proto_init() {
	if File_core_v1_sys_dept_proto != nil {
		return
	}
	file_core_v1_base_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_core_v1_sys_dept_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_core_v1_sys_dept_proto_goTypes,
		DependencyIndexes: file_core_v1_sys_dept_proto_depIdxs,
		MessageInfos:      file_core_v1_sys_dept_proto_msgTypes,
	}.Build()
	File_core_v1_sys_dept_proto = out.File
	file_core_v1_sys_dept_proto_rawDesc = nil
	file_core_v1_sys_dept_proto_goTypes = nil
	file_core_v1_sys_dept_proto_depIdxs = nil
}