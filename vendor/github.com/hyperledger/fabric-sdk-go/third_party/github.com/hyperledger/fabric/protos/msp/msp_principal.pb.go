/*
Notice: This file has been modified for Hyperledger Fabric SDK Go usage.
Please review third_party pinning scripts and patches for more details.
*/
// Code generated by protoc-gen-go. DO NOT EDIT.
// source: msp/msp_principal.proto

package msp

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type MSPPrincipal_Classification int32

const (
	MSPPrincipal_ROLE MSPPrincipal_Classification = 0
	// one of a member of MSP network, and the one of an
	// administrator of an MSP network
	MSPPrincipal_ORGANIZATION_UNIT MSPPrincipal_Classification = 1
	// groupping of entities, per MSP affiliation
	// E.g., this can well be represented by an MSP's
	// Organization unit
	MSPPrincipal_IDENTITY MSPPrincipal_Classification = 2
)

var MSPPrincipal_Classification_name = map[int32]string{
	0: "ROLE",
	1: "ORGANIZATION_UNIT",
	2: "IDENTITY",
}
var MSPPrincipal_Classification_value = map[string]int32{
	"ROLE":              0,
	"ORGANIZATION_UNIT": 1,
	"IDENTITY":          2,
}

func (x MSPPrincipal_Classification) String() string {
	return proto.EnumName(MSPPrincipal_Classification_name, int32(x))
}
func (MSPPrincipal_Classification) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor2, []int{0, 0}
}

type MSPRole_MSPRoleType int32

const (
	MSPRole_MEMBER  MSPRole_MSPRoleType = 0
	MSPRole_ADMIN   MSPRole_MSPRoleType = 1
	MSPRole_CLIENT  MSPRole_MSPRoleType = 2
	MSPRole_PEER    MSPRole_MSPRoleType = 3
	MSPRole_ORDERER MSPRole_MSPRoleType = 4
)

var MSPRole_MSPRoleType_name = map[int32]string{
	0: "MEMBER",
	1: "ADMIN",
	2: "CLIENT",
	3: "PEER",
	4: "ORDERER",
}
var MSPRole_MSPRoleType_value = map[string]int32{
	"MEMBER":  0,
	"ADMIN":   1,
	"CLIENT":  2,
	"PEER":    3,
	"ORDERER": 4,
}

func (x MSPRole_MSPRoleType) String() string {
	return proto.EnumName(MSPRole_MSPRoleType_name, int32(x))
}
func (MSPRole_MSPRoleType) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{2, 0} }

// MSPPrincipal aims to represent an MSP-centric set of identities.
// In particular, this structure allows for definition of
//  - a group of identities that are member of the same MSP
//  - a group of identities that are member of the same organization unit
//    in the same MSP
//  - a group of identities that are administering a specific MSP
//  - a specific identity
// Expressing these groups is done given two fields of the fields below
//  - Classification, that defines the type of classification of identities
//    in an MSP this principal would be defined on; Classification can take
//    three values:
//     (i)  ByMSPRole: that represents a classification of identities within
//          MSP based on one of the two pre-defined MSP rules, "member" and "admin"
//     (ii) ByOrganizationUnit: that represents a classification of identities
//          within MSP based on the organization unit an identity belongs to
//     (iii)ByIdentity that denotes that MSPPrincipal is mapped to a single
//          identity/certificate; this would mean that the Principal bytes
//          message
type MSPPrincipal struct {
	// Classification describes the way that one should process
	// Principal. An Classification value of "ByOrganizationUnit" reflects
	// that "Principal" contains the name of an organization this MSP
	// handles. A Classification value "ByIdentity" means that
	// "Principal" contains a specific identity. Default value
	// denotes that Principal contains one of the groups by
	// default supported by all MSPs ("admin" or "member").
	PrincipalClassification MSPPrincipal_Classification `protobuf:"varint,1,opt,name=principal_classification,json=principalClassification,enum=common.MSPPrincipal_Classification" json:"principal_classification,omitempty"`
	// Principal completes the policy principal definition. For the default
	// principal types, Principal can be either "Admin" or "Member".
	// For the ByOrganizationUnit/ByIdentity values of Classification,
	// PolicyPrincipal acquires its value from an organization unit or
	// identity, respectively.
	Principal []byte `protobuf:"bytes,2,opt,name=principal,proto3" json:"principal,omitempty"`
}

func (m *MSPPrincipal) Reset()                    { *m = MSPPrincipal{} }
func (m *MSPPrincipal) String() string            { return proto.CompactTextString(m) }
func (*MSPPrincipal) ProtoMessage()               {}
func (*MSPPrincipal) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *MSPPrincipal) GetPrincipalClassification() MSPPrincipal_Classification {
	if m != nil {
		return m.PrincipalClassification
	}
	return MSPPrincipal_ROLE
}

func (m *MSPPrincipal) GetPrincipal() []byte {
	if m != nil {
		return m.Principal
	}
	return nil
}

// OrganizationUnit governs the organization of the Principal
// field of a policy principal when a specific organization unity members
// are to be defined within a policy principal.
type OrganizationUnit struct {
	// MSPIdentifier represents the identifier of the MSP this organization unit
	// refers to
	MspIdentifier string `protobuf:"bytes,1,opt,name=msp_identifier,json=mspIdentifier" json:"msp_identifier,omitempty"`
	// OrganizationUnitIdentifier defines the organizational unit under the
	// MSP identified with MSPIdentifier
	OrganizationalUnitIdentifier string `protobuf:"bytes,2,opt,name=organizational_unit_identifier,json=organizationalUnitIdentifier" json:"organizational_unit_identifier,omitempty"`
	// CertifiersIdentifier is the hash of certificates chain of trust
	// related to this organizational unit
	CertifiersIdentifier []byte `protobuf:"bytes,3,opt,name=certifiers_identifier,json=certifiersIdentifier,proto3" json:"certifiers_identifier,omitempty"`
}

func (m *OrganizationUnit) Reset()                    { *m = OrganizationUnit{} }
func (m *OrganizationUnit) String() string            { return proto.CompactTextString(m) }
func (*OrganizationUnit) ProtoMessage()               {}
func (*OrganizationUnit) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *OrganizationUnit) GetMspIdentifier() string {
	if m != nil {
		return m.MspIdentifier
	}
	return ""
}

func (m *OrganizationUnit) GetOrganizationalUnitIdentifier() string {
	if m != nil {
		return m.OrganizationalUnitIdentifier
	}
	return ""
}

func (m *OrganizationUnit) GetCertifiersIdentifier() []byte {
	if m != nil {
		return m.CertifiersIdentifier
	}
	return nil
}

// MSPRole governs the organization of the Principal
// field of an MSPPrincipal when it aims to define one of the
// two dedicated roles within an MSP: Admin and Members.
type MSPRole struct {
	// MSPIdentifier represents the identifier of the MSP this principal
	// refers to
	MspIdentifier string `protobuf:"bytes,1,opt,name=msp_identifier,json=mspIdentifier" json:"msp_identifier,omitempty"`
	// MSPRoleType defines which of the available, pre-defined MSP-roles
	// an identiy should posess inside the MSP with identifier MSPidentifier
	Role MSPRole_MSPRoleType `protobuf:"varint,2,opt,name=role,enum=common.MSPRole_MSPRoleType" json:"role,omitempty"`
}

func (m *MSPRole) Reset()                    { *m = MSPRole{} }
func (m *MSPRole) String() string            { return proto.CompactTextString(m) }
func (*MSPRole) ProtoMessage()               {}
func (*MSPRole) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

func (m *MSPRole) GetMspIdentifier() string {
	if m != nil {
		return m.MspIdentifier
	}
	return ""
}

func (m *MSPRole) GetRole() MSPRole_MSPRoleType {
	if m != nil {
		return m.Role
	}
	return MSPRole_MEMBER
}

func init() {
	proto.RegisterType((*MSPPrincipal)(nil), "sdk.common.MSPPrincipal")
	proto.RegisterType((*OrganizationUnit)(nil), "sdk.common.OrganizationUnit")
	proto.RegisterType((*MSPRole)(nil), "sdk.common.MSPRole")
	proto.RegisterEnum("sdk.common.MSPPrincipal_Classification", MSPPrincipal_Classification_name, MSPPrincipal_Classification_value)
	proto.RegisterEnum("sdk.common.MSPRole_MSPRoleType", MSPRole_MSPRoleType_name, MSPRole_MSPRoleType_value)
}

func init() { proto.RegisterFile("msp/msp_principal.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 411 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xdf, 0x6a, 0xdb, 0x30,
	0x14, 0x87, 0xab, 0x34, 0x4b, 0x9b, 0xd3, 0x2c, 0x68, 0x62, 0xa5, 0x81, 0x95, 0x51, 0xbc, 0x0d,
	0x72, 0x65, 0x43, 0xfb, 0x00, 0x23, 0x6d, 0x44, 0x11, 0xd4, 0x7f, 0x50, 0xdd, 0x8b, 0xf5, 0x62,
	0xc1, 0x71, 0x15, 0x57, 0x60, 0x5b, 0x42, 0x76, 0x2f, 0xba, 0x47, 0xda, 0xf5, 0x5e, 0x63, 0xef,
	0x34, 0x6c, 0x2f, 0x8e, 0xb2, 0xab, 0x5d, 0xd9, 0x3a, 0xbf, 0xef, 0x3b, 0x3e, 0x92, 0x05, 0x67,
	0x45, 0xa5, 0xbd, 0xa2, 0xd2, 0x2b, 0x6d, 0x64, 0x99, 0x4a, 0x9d, 0xe4, 0xae, 0x36, 0xaa, 0x56,
	0x64, 0x94, 0xaa, 0xa2, 0x50, 0xa5, 0xf3, 0x1b, 0xc1, 0xc4, 0xbf, 0x8f, 0xa2, 0x6d, 0x4c, 0xbe,
	0xc3, 0xac, 0x67, 0x57, 0x69, 0x9e, 0x54, 0x95, 0xdc, 0xc8, 0x34, 0xa9, 0xa5, 0x2a, 0x67, 0xe8,
	0x02, 0xcd, 0xa7, 0x97, 0x9f, 0xdc, 0xce, 0x75, 0x6d, 0xcf, 0xbd, 0xd9, 0x43, 0xf9, 0x59, 0xdf,
	0x64, 0x3f, 0x20, 0xe7, 0x30, 0xee, 0xa3, 0xd9, 0xe0, 0x02, 0xcd, 0x27, 0x7c, 0x57, 0x70, 0xbe,
	0xc2, 0xf4, 0x1f, 0xfe, 0x18, 0x86, 0x3c, 0xbc, 0xa3, 0xf8, 0x80, 0x9c, 0xc2, 0xbb, 0x90, 0xdf,
	0x2e, 0x02, 0xf6, 0xb8, 0x88, 0x59, 0x18, 0xac, 0x1e, 0x02, 0x16, 0x63, 0x44, 0x26, 0x70, 0xcc,
	0x96, 0x34, 0x88, 0x59, 0xfc, 0x0d, 0x0f, 0x9c, 0x5f, 0x08, 0x70, 0x68, 0xb2, 0xa4, 0x94, 0x3f,
	0x5a, 0xff, 0xa1, 0x94, 0x35, 0xf9, 0x02, 0xd3, 0xe6, 0x0c, 0xe4, 0x93, 0x28, 0x6b, 0xb9, 0x91,
	0xc2, 0xb4, 0x3b, 0x19, 0xf3, 0xb7, 0x45, 0xa5, 0x59, 0x5f, 0x24, 0x4b, 0xf8, 0xa8, 0x2c, 0x35,
	0xc9, 0x57, 0x2f, 0xa5, 0xac, 0x6d, 0x6d, 0xd0, 0x6a, 0xe7, 0xfb, 0x54, 0xf3, 0x09, 0xab, 0xcb,
	0x15, 0x9c, 0xa6, 0xc2, 0x74, 0x8b, 0xca, 0x96, 0x0f, 0xdb, 0xcd, 0xbe, 0xdf, 0x85, 0x3b, 0xc9,
	0xf9, 0x89, 0xe0, 0xc8, 0xbf, 0x8f, 0xb8, 0xca, 0xc5, 0xff, 0x4e, 0xeb, 0xc1, 0xd0, 0xa8, 0x5c,
	0xb4, 0x33, 0x4d, 0x2f, 0x3f, 0x58, 0x3f, 0xa5, 0xe9, 0xb2, 0x7d, 0xc6, 0xaf, 0x5a, 0xf0, 0x16,
	0x74, 0x6e, 0xe1, 0xc4, 0x2a, 0x12, 0x80, 0x91, 0x4f, 0xfd, 0x6b, 0xca, 0xf1, 0x01, 0x19, 0xc3,
	0x9b, 0xc5, 0xd2, 0x67, 0x01, 0x46, 0x4d, 0xf9, 0xe6, 0x8e, 0xd1, 0x20, 0xc6, 0x83, 0xe6, 0xec,
	0x23, 0x4a, 0x39, 0x3e, 0x24, 0x27, 0x70, 0x14, 0xf2, 0x25, 0xe5, 0x94, 0xe3, 0xe1, 0x75, 0x04,
	0x9f, 0x95, 0xc9, 0xdc, 0xe7, 0x57, 0x2d, 0x4c, 0x2e, 0x9e, 0x32, 0x61, 0xdc, 0x4d, 0xb2, 0x36,
	0x32, 0xed, 0xee, 0x56, 0xf5, 0x77, 0x94, 0xc7, 0x79, 0x26, 0xeb, 0xe7, 0x97, 0x75, 0xb3, 0xf4,
	0x2c, 0xd8, 0xeb, 0x60, 0xaf, 0x83, 0x9b, 0xdb, 0xb9, 0x1e, 0xb5, 0xef, 0x57, 0x7f, 0x02, 0x00,
	0x00, 0xff, 0xff, 0x95, 0x4b, 0xcf, 0xee, 0xaf, 0x02, 0x00, 0x00,
}