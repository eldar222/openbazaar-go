// Code generated by protoc-gen-go.
// source: profile.proto
// DO NOT EDIT!

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Profile struct {
	Handle             string                     `protobuf:"bytes,1,opt,name=handle" json:"handle,omitempty"`
	Name               string                     `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Location           string                     `protobuf:"bytes,3,opt,name=location" json:"location,omitempty"`
	About              string                     `protobuf:"bytes,4,opt,name=about" json:"about,omitempty"`
	ShortDescription   string                     `protobuf:"bytes,5,opt,name=shortDescription" json:"shortDescription,omitempty"`
	Website            string                     `protobuf:"bytes,6,opt,name=website" json:"website,omitempty"`
	Email              string                     `protobuf:"bytes,7,opt,name=email" json:"email,omitempty"`
	PhoneNumber        string                     `protobuf:"bytes,8,opt,name=phoneNumber" json:"phoneNumber,omitempty"`
	Social             []*Profile_SocialAccount   `protobuf:"bytes,9,rep,name=social" json:"social,omitempty"`
	AvgRating          uint32                     `protobuf:"varint,10,opt,name=avgRating" json:"avgRating,omitempty"`
	NumRatings         uint32                     `protobuf:"varint,11,opt,name=numRatings" json:"numRatings,omitempty"`
	Nsfw               bool                       `protobuf:"varint,12,opt,name=nsfw" json:"nsfw,omitempty"`
	Vendor             bool                       `protobuf:"varint,13,opt,name=vendor" json:"vendor,omitempty"`
	Moderator          bool                       `protobuf:"varint,14,opt,name=moderator" json:"moderator,omitempty"`
	ModInfo            *Moderator                 `protobuf:"bytes,15,opt,name=modInfo" json:"modInfo,omitempty"`
	PrimaryColor       string                     `protobuf:"bytes,16,opt,name=primaryColor" json:"primaryColor,omitempty"`
	SecondaryColor     string                     `protobuf:"bytes,17,opt,name=secondaryColor" json:"secondaryColor,omitempty"`
	TextColor          string                     `protobuf:"bytes,18,opt,name=textColor" json:"textColor,omitempty"`
	HighlightColor     string                     `protobuf:"bytes,19,opt,name=highlightColor" json:"highlightColor,omitempty"`
	HighlightTextColor string                     `protobuf:"bytes,20,opt,name=highlightTextColor" json:"highlightTextColor,omitempty"`
	AvatarHashes       *Profile_Image             `protobuf:"bytes,21,opt,name=avatarHashes" json:"avatarHashes,omitempty"`
	HeaderHashes       *Profile_Image             `protobuf:"bytes,22,opt,name=headerHashes" json:"headerHashes,omitempty"`
	FollowerCount      uint32                     `protobuf:"varint,23,opt,name=followerCount" json:"followerCount,omitempty"`
	FollowingCount     uint32                     `protobuf:"varint,24,opt,name=followingCount" json:"followingCount,omitempty"`
	ListingCount       uint32                     `protobuf:"varint,25,opt,name=listingCount" json:"listingCount,omitempty"`
	BitcoinPubkey      string                     `protobuf:"bytes,26,opt,name=bitcoinPubkey" json:"bitcoinPubkey,omitempty"`
	LastModified       *google_protobuf.Timestamp `protobuf:"bytes,27,opt,name=lastModified" json:"lastModified,omitempty"`
}

func (m *Profile) Reset()                    { *m = Profile{} }
func (m *Profile) String() string            { return proto.CompactTextString(m) }
func (*Profile) ProtoMessage()               {}
func (*Profile) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{0} }

func (m *Profile) GetHandle() string {
	if m != nil {
		return m.Handle
	}
	return ""
}

func (m *Profile) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Profile) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func (m *Profile) GetAbout() string {
	if m != nil {
		return m.About
	}
	return ""
}

func (m *Profile) GetShortDescription() string {
	if m != nil {
		return m.ShortDescription
	}
	return ""
}

func (m *Profile) GetWebsite() string {
	if m != nil {
		return m.Website
	}
	return ""
}

func (m *Profile) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Profile) GetPhoneNumber() string {
	if m != nil {
		return m.PhoneNumber
	}
	return ""
}

func (m *Profile) GetSocial() []*Profile_SocialAccount {
	if m != nil {
		return m.Social
	}
	return nil
}

func (m *Profile) GetAvgRating() uint32 {
	if m != nil {
		return m.AvgRating
	}
	return 0
}

func (m *Profile) GetNumRatings() uint32 {
	if m != nil {
		return m.NumRatings
	}
	return 0
}

func (m *Profile) GetNsfw() bool {
	if m != nil {
		return m.Nsfw
	}
	return false
}

func (m *Profile) GetVendor() bool {
	if m != nil {
		return m.Vendor
	}
	return false
}

func (m *Profile) GetModerator() bool {
	if m != nil {
		return m.Moderator
	}
	return false
}

func (m *Profile) GetModInfo() *Moderator {
	if m != nil {
		return m.ModInfo
	}
	return nil
}

func (m *Profile) GetPrimaryColor() string {
	if m != nil {
		return m.PrimaryColor
	}
	return ""
}

func (m *Profile) GetSecondaryColor() string {
	if m != nil {
		return m.SecondaryColor
	}
	return ""
}

func (m *Profile) GetTextColor() string {
	if m != nil {
		return m.TextColor
	}
	return ""
}

func (m *Profile) GetHighlightColor() string {
	if m != nil {
		return m.HighlightColor
	}
	return ""
}

func (m *Profile) GetHighlightTextColor() string {
	if m != nil {
		return m.HighlightTextColor
	}
	return ""
}

func (m *Profile) GetAvatarHashes() *Profile_Image {
	if m != nil {
		return m.AvatarHashes
	}
	return nil
}

func (m *Profile) GetHeaderHashes() *Profile_Image {
	if m != nil {
		return m.HeaderHashes
	}
	return nil
}

func (m *Profile) GetFollowerCount() uint32 {
	if m != nil {
		return m.FollowerCount
	}
	return 0
}

func (m *Profile) GetFollowingCount() uint32 {
	if m != nil {
		return m.FollowingCount
	}
	return 0
}

func (m *Profile) GetListingCount() uint32 {
	if m != nil {
		return m.ListingCount
	}
	return 0
}

func (m *Profile) GetBitcoinPubkey() string {
	if m != nil {
		return m.BitcoinPubkey
	}
	return ""
}

func (m *Profile) GetLastModified() *google_protobuf.Timestamp {
	if m != nil {
		return m.LastModified
	}
	return nil
}

type Profile_SocialAccount struct {
	Type     string `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username" json:"username,omitempty"`
	Proof    string `protobuf:"bytes,3,opt,name=proof" json:"proof,omitempty"`
}

func (m *Profile_SocialAccount) Reset()                    { *m = Profile_SocialAccount{} }
func (m *Profile_SocialAccount) String() string            { return proto.CompactTextString(m) }
func (*Profile_SocialAccount) ProtoMessage()               {}
func (*Profile_SocialAccount) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{0, 0} }

func (m *Profile_SocialAccount) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Profile_SocialAccount) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *Profile_SocialAccount) GetProof() string {
	if m != nil {
		return m.Proof
	}
	return ""
}

type Profile_Image struct {
	Tiny     string `protobuf:"bytes,1,opt,name=tiny" json:"tiny,omitempty"`
	Small    string `protobuf:"bytes,2,opt,name=small" json:"small,omitempty"`
	Medium   string `protobuf:"bytes,3,opt,name=medium" json:"medium,omitempty"`
	Large    string `protobuf:"bytes,4,opt,name=large" json:"large,omitempty"`
	Original string `protobuf:"bytes,5,opt,name=original" json:"original,omitempty"`
}

func (m *Profile_Image) Reset()                    { *m = Profile_Image{} }
func (m *Profile_Image) String() string            { return proto.CompactTextString(m) }
func (*Profile_Image) ProtoMessage()               {}
func (*Profile_Image) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{0, 1} }

func (m *Profile_Image) GetTiny() string {
	if m != nil {
		return m.Tiny
	}
	return ""
}

func (m *Profile_Image) GetSmall() string {
	if m != nil {
		return m.Small
	}
	return ""
}

func (m *Profile_Image) GetMedium() string {
	if m != nil {
		return m.Medium
	}
	return ""
}

func (m *Profile_Image) GetLarge() string {
	if m != nil {
		return m.Large
	}
	return ""
}

func (m *Profile_Image) GetOriginal() string {
	if m != nil {
		return m.Original
	}
	return ""
}

func init() {
	proto.RegisterType((*Profile)(nil), "Profile")
	proto.RegisterType((*Profile_SocialAccount)(nil), "Profile.SocialAccount")
	proto.RegisterType((*Profile_Image)(nil), "Profile.Image")
}

func init() { proto.RegisterFile("profile.proto", fileDescriptor6) }

var fileDescriptor6 = []byte{
	// 618 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x94, 0x4f, 0x6f, 0xdb, 0x38,
	0x10, 0xc5, 0xe1, 0xfc, 0xb1, 0x63, 0xda, 0x4e, 0xb2, 0xdc, 0x6c, 0x96, 0xeb, 0x2d, 0x5a, 0x23,
	0x08, 0x0a, 0xa3, 0x07, 0x05, 0x48, 0xef, 0x05, 0xda, 0xf4, 0xd0, 0x1c, 0x52, 0x04, 0x6a, 0x7a,
	0xe9, 0x8d, 0x92, 0x28, 0x89, 0x28, 0xc9, 0x11, 0x48, 0x2a, 0xa9, 0xd1, 0x4f, 0xd2, 0x6f, 0x5b,
	0x70, 0xf4, 0x27, 0x56, 0x9a, 0x9b, 0xde, 0x6f, 0x1e, 0x9f, 0xc4, 0x21, 0x47, 0x64, 0x51, 0x59,
	0xc8, 0xa5, 0x12, 0x51, 0x65, 0xc1, 0xc3, 0xf2, 0x55, 0x01, 0x50, 0x28, 0x71, 0x81, 0x2a, 0xa9,
	0xf3, 0x0b, 0x2f, 0xb5, 0x70, 0x9e, 0xeb, 0xaa, 0x35, 0x1c, 0x69, 0xc8, 0x84, 0xe5, 0x1e, 0x6c,
	0x03, 0xce, 0x7e, 0x4d, 0xc9, 0xe4, 0xb6, 0xc9, 0xa0, 0xa7, 0x64, 0x5c, 0x72, 0x93, 0x29, 0xc1,
	0x46, 0xab, 0xd1, 0x7a, 0x1a, 0xb7, 0x8a, 0x52, 0xb2, 0x67, 0xb8, 0x16, 0x6c, 0x07, 0x29, 0x3e,
	0xd3, 0x25, 0x39, 0x50, 0x90, 0x72, 0x2f, 0xc1, 0xb0, 0x5d, 0xe4, 0xbd, 0xa6, 0x27, 0x64, 0x9f,
	0x27, 0x50, 0x7b, 0xb6, 0x87, 0x85, 0x46, 0xd0, 0x37, 0xe4, 0xd8, 0x95, 0x60, 0xfd, 0x47, 0xe1,
	0x52, 0x2b, 0x2b, 0x5c, 0xb9, 0x8f, 0x86, 0x3f, 0x38, 0x65, 0x64, 0xf2, 0x20, 0x12, 0x27, 0xbd,
	0x60, 0x63, 0xb4, 0x74, 0x32, 0x64, 0x0b, 0xcd, 0xa5, 0x62, 0x93, 0x26, 0x1b, 0x05, 0x5d, 0x91,
	0x59, 0x55, 0x82, 0x11, 0x9f, 0x6b, 0x9d, 0x08, 0xcb, 0x0e, 0xb0, 0xb6, 0x8d, 0x68, 0x44, 0xc6,
	0x0e, 0x52, 0xc9, 0x15, 0x9b, 0xae, 0x76, 0xd7, 0xb3, 0xcb, 0xd3, 0xa8, 0xdd, 0x75, 0xf4, 0x05,
	0xf1, 0xfb, 0x34, 0x85, 0xda, 0xf8, 0xb8, 0x75, 0xd1, 0x17, 0x64, 0xca, 0xef, 0x8b, 0x98, 0x7b,
	0x69, 0x0a, 0x46, 0x56, 0xa3, 0xf5, 0x22, 0x7e, 0x04, 0xf4, 0x25, 0x21, 0xa6, 0xd6, 0x8d, 0x70,
	0x6c, 0x86, 0xe5, 0x2d, 0x82, 0x1d, 0x73, 0xf9, 0x03, 0x9b, 0xaf, 0x46, 0xeb, 0x83, 0x18, 0x9f,
	0x43, 0x77, 0xef, 0x85, 0xc9, 0xc0, 0xb2, 0x05, 0xd2, 0x56, 0x85, 0x37, 0xf5, 0x87, 0xc2, 0x0e,
	0xb1, 0xf4, 0x08, 0xe8, 0x39, 0x99, 0x68, 0xc8, 0xae, 0x4d, 0x0e, 0xec, 0x68, 0x35, 0x5a, 0xcf,
	0x2e, 0x49, 0x74, 0xd3, 0x15, 0xe3, 0xae, 0x44, 0xcf, 0xc8, 0xbc, 0xb2, 0x52, 0x73, 0xbb, 0xb9,
	0x02, 0x05, 0x96, 0x1d, 0x63, 0x03, 0x06, 0x8c, 0xbe, 0x26, 0x87, 0x4e, 0xa4, 0x60, 0xb2, 0xde,
	0xf5, 0x17, 0xba, 0x9e, 0xd0, 0xf0, 0x3d, 0x5e, 0xfc, 0xf0, 0x8d, 0x85, 0xa2, 0xe5, 0x11, 0x84,
	0x94, 0x52, 0x16, 0xa5, 0x92, 0x45, 0xd9, 0x5a, 0xfe, 0x6e, 0x52, 0x86, 0x94, 0x46, 0x84, 0xf6,
	0xe4, 0xae, 0x8f, 0x3b, 0x41, 0xef, 0x33, 0x15, 0x7a, 0x49, 0xe6, 0xfc, 0x9e, 0x7b, 0x6e, 0x3f,
	0x71, 0x57, 0x0a, 0xc7, 0xfe, 0xc1, 0xcd, 0x1e, 0xf6, 0xa7, 0x74, 0xad, 0x79, 0x21, 0xe2, 0x81,
	0x27, 0xac, 0x29, 0x05, 0xcf, 0x44, 0xb7, 0xe6, 0xf4, 0xf9, 0x35, 0xdb, 0x1e, 0x7a, 0x4e, 0x16,
	0x39, 0x28, 0x05, 0x0f, 0xc2, 0x5e, 0x85, 0x03, 0x67, 0xff, 0xe2, 0xe1, 0x0d, 0x61, 0xd8, 0x65,
	0x03, 0xa4, 0x29, 0x1a, 0x1b, 0x43, 0xdb, 0x13, 0x1a, 0xfa, 0xae, 0xa4, 0xf3, 0xbd, 0xeb, 0x3f,
	0x74, 0x0d, 0x58, 0x78, 0x63, 0x22, 0x7d, 0x0a, 0xd2, 0xdc, 0xd6, 0xc9, 0x77, 0xb1, 0x61, 0x4b,
	0x6c, 0xc2, 0x10, 0xd2, 0x77, 0x64, 0xae, 0xb8, 0xf3, 0x37, 0x90, 0xc9, 0x5c, 0x8a, 0x8c, 0xfd,
	0x8f, 0x7b, 0x59, 0x46, 0xcd, 0x40, 0x47, 0xdd, 0x40, 0x47, 0x77, 0xdd, 0x40, 0xc7, 0x03, 0xff,
	0xf2, 0x2b, 0x59, 0x0c, 0x2e, 0x72, 0xb8, 0x82, 0x7e, 0x53, 0x75, 0xa3, 0x8c, 0xcf, 0x61, 0x68,
	0x6b, 0x27, 0xec, 0xd6, 0x30, 0xf7, 0x3a, 0x0c, 0x56, 0x65, 0x01, 0xf2, 0x76, 0x9a, 0x1b, 0xb1,
	0xfc, 0x49, 0xf6, 0xb1, 0x8b, 0x18, 0x27, 0xcd, 0xa6, 0x8f, 0x93, 0x66, 0x13, 0x96, 0x38, 0xcd,
	0x95, 0x6a, 0xb3, 0x1a, 0x11, 0xee, 0xb9, 0x16, 0x99, 0xac, 0x75, 0x9b, 0xd4, 0xaa, 0xe0, 0x56,
	0xdc, 0x16, 0xa2, 0xfb, 0x2b, 0xa0, 0x08, 0x9f, 0x04, 0x56, 0x16, 0xd2, 0x70, 0xd5, 0xfe, 0x0d,
	0x7a, 0xfd, 0x61, 0xef, 0xdb, 0x4e, 0x95, 0x24, 0x63, 0xdc, 0xfb, 0xdb, 0xdf, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x20, 0x67, 0x27, 0xda, 0xeb, 0x04, 0x00, 0x00,
}
