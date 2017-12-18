// Code generated by protoc-gen-gogo.
// source: github.com/containerd/containerd/api/types/task/task.proto
// DO NOT EDIT!

/*
	Package task is a generated protocol buffer package.

	It is generated from these files:
		github.com/containerd/containerd/api/types/task/task.proto

	It has these top-level messages:
		Process
		ProcessInfo
*/
package task

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "github.com/gogo/protobuf/types"
import google_protobuf2 "github.com/gogo/protobuf/types"

import time "time"

import github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Status int32

const (
	StatusUnknown Status = 0
	StatusCreated Status = 1
	StatusRunning Status = 2
	StatusStopped Status = 3
	StatusPaused  Status = 4
	StatusPausing Status = 5
)

var Status_name = map[int32]string{
	0: "UNKNOWN",
	1: "CREATED",
	2: "RUNNING",
	3: "STOPPED",
	4: "PAUSED",
	5: "PAUSING",
}
var Status_value = map[string]int32{
	"UNKNOWN": 0,
	"CREATED": 1,
	"RUNNING": 2,
	"STOPPED": 3,
	"PAUSED":  4,
	"PAUSING": 5,
}

func (x Status) String() string {
	return proto.EnumName(Status_name, int32(x))
}
func (Status) EnumDescriptor() ([]byte, []int) { return fileDescriptorTask, []int{0} }

type Process struct {
	ContainerID string    `protobuf:"bytes,1,opt,name=container_id,json=containerId,proto3" json:"container_id,omitempty"`
	ID          string    `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Pid         uint32    `protobuf:"varint,3,opt,name=pid,proto3" json:"pid,omitempty"`
	Status      Status    `protobuf:"varint,4,opt,name=status,proto3,enum=containerd.v1.types.Status" json:"status,omitempty"`
	Stdin       string    `protobuf:"bytes,5,opt,name=stdin,proto3" json:"stdin,omitempty"`
	Stdout      string    `protobuf:"bytes,6,opt,name=stdout,proto3" json:"stdout,omitempty"`
	Stderr      string    `protobuf:"bytes,7,opt,name=stderr,proto3" json:"stderr,omitempty"`
	Terminal    bool      `protobuf:"varint,8,opt,name=terminal,proto3" json:"terminal,omitempty"`
	ExitStatus  uint32    `protobuf:"varint,9,opt,name=exit_status,json=exitStatus,proto3" json:"exit_status,omitempty"`
	ExitedAt    time.Time `protobuf:"bytes,10,opt,name=exited_at,json=exitedAt,stdtime" json:"exited_at"`
}

func (m *Process) Reset()                    { *m = Process{} }
func (*Process) ProtoMessage()               {}
func (*Process) Descriptor() ([]byte, []int) { return fileDescriptorTask, []int{0} }

type ProcessInfo struct {
	// PID is the process ID.
	Pid uint32 `protobuf:"varint,1,opt,name=pid,proto3" json:"pid,omitempty"`
	// Info contains additional process information.
	//
	// Info varies by platform.
	Info *google_protobuf2.Any `protobuf:"bytes,2,opt,name=info" json:"info,omitempty"`
}

func (m *ProcessInfo) Reset()                    { *m = ProcessInfo{} }
func (*ProcessInfo) ProtoMessage()               {}
func (*ProcessInfo) Descriptor() ([]byte, []int) { return fileDescriptorTask, []int{1} }

func init() {
	proto.RegisterType((*Process)(nil), "containerd.v1.types.Process")
	proto.RegisterType((*ProcessInfo)(nil), "containerd.v1.types.ProcessInfo")
	proto.RegisterEnum("containerd.v1.types.Status", Status_name, Status_value)
}
func (m *Process) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Process) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ContainerID) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintTask(dAtA, i, uint64(len(m.ContainerID)))
		i += copy(dAtA[i:], m.ContainerID)
	}
	if len(m.ID) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintTask(dAtA, i, uint64(len(m.ID)))
		i += copy(dAtA[i:], m.ID)
	}
	if m.Pid != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintTask(dAtA, i, uint64(m.Pid))
	}
	if m.Status != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintTask(dAtA, i, uint64(m.Status))
	}
	if len(m.Stdin) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintTask(dAtA, i, uint64(len(m.Stdin)))
		i += copy(dAtA[i:], m.Stdin)
	}
	if len(m.Stdout) > 0 {
		dAtA[i] = 0x32
		i++
		i = encodeVarintTask(dAtA, i, uint64(len(m.Stdout)))
		i += copy(dAtA[i:], m.Stdout)
	}
	if len(m.Stderr) > 0 {
		dAtA[i] = 0x3a
		i++
		i = encodeVarintTask(dAtA, i, uint64(len(m.Stderr)))
		i += copy(dAtA[i:], m.Stderr)
	}
	if m.Terminal {
		dAtA[i] = 0x40
		i++
		if m.Terminal {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.ExitStatus != 0 {
		dAtA[i] = 0x48
		i++
		i = encodeVarintTask(dAtA, i, uint64(m.ExitStatus))
	}
	dAtA[i] = 0x52
	i++
	i = encodeVarintTask(dAtA, i, uint64(github_com_gogo_protobuf_types.SizeOfStdTime(m.ExitedAt)))
	n1, err := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.ExitedAt, dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	return i, nil
}

func (m *ProcessInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProcessInfo) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Pid != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintTask(dAtA, i, uint64(m.Pid))
	}
	if m.Info != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintTask(dAtA, i, uint64(m.Info.Size()))
		n2, err := m.Info.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}

func encodeFixed64Task(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Task(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintTask(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Process) Size() (n int) {
	var l int
	_ = l
	l = len(m.ContainerID)
	if l > 0 {
		n += 1 + l + sovTask(uint64(l))
	}
	l = len(m.ID)
	if l > 0 {
		n += 1 + l + sovTask(uint64(l))
	}
	if m.Pid != 0 {
		n += 1 + sovTask(uint64(m.Pid))
	}
	if m.Status != 0 {
		n += 1 + sovTask(uint64(m.Status))
	}
	l = len(m.Stdin)
	if l > 0 {
		n += 1 + l + sovTask(uint64(l))
	}
	l = len(m.Stdout)
	if l > 0 {
		n += 1 + l + sovTask(uint64(l))
	}
	l = len(m.Stderr)
	if l > 0 {
		n += 1 + l + sovTask(uint64(l))
	}
	if m.Terminal {
		n += 2
	}
	if m.ExitStatus != 0 {
		n += 1 + sovTask(uint64(m.ExitStatus))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.ExitedAt)
	n += 1 + l + sovTask(uint64(l))
	return n
}

func (m *ProcessInfo) Size() (n int) {
	var l int
	_ = l
	if m.Pid != 0 {
		n += 1 + sovTask(uint64(m.Pid))
	}
	if m.Info != nil {
		l = m.Info.Size()
		n += 1 + l + sovTask(uint64(l))
	}
	return n
}

func sovTask(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozTask(x uint64) (n int) {
	return sovTask(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Process) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Process{`,
		`ContainerID:` + fmt.Sprintf("%v", this.ContainerID) + `,`,
		`ID:` + fmt.Sprintf("%v", this.ID) + `,`,
		`Pid:` + fmt.Sprintf("%v", this.Pid) + `,`,
		`Status:` + fmt.Sprintf("%v", this.Status) + `,`,
		`Stdin:` + fmt.Sprintf("%v", this.Stdin) + `,`,
		`Stdout:` + fmt.Sprintf("%v", this.Stdout) + `,`,
		`Stderr:` + fmt.Sprintf("%v", this.Stderr) + `,`,
		`Terminal:` + fmt.Sprintf("%v", this.Terminal) + `,`,
		`ExitStatus:` + fmt.Sprintf("%v", this.ExitStatus) + `,`,
		`ExitedAt:` + strings.Replace(strings.Replace(this.ExitedAt.String(), "Timestamp", "google_protobuf1.Timestamp", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ProcessInfo) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ProcessInfo{`,
		`Pid:` + fmt.Sprintf("%v", this.Pid) + `,`,
		`Info:` + strings.Replace(fmt.Sprintf("%v", this.Info), "Any", "google_protobuf2.Any", 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringTask(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Process) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTask
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Process: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Process: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContainerID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTask
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContainerID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTask
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pid", wireType)
			}
			m.Pid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Pid |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= (Status(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Stdin", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTask
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Stdin = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Stdout", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTask
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Stdout = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Stderr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTask
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Stderr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Terminal", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Terminal = bool(v != 0)
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExitStatus", wireType)
			}
			m.ExitStatus = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ExitStatus |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExitedAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTask
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.ExitedAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTask(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTask
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ProcessInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTask
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ProcessInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProcessInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pid", wireType)
			}
			m.Pid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Pid |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Info", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTask
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Info == nil {
				m.Info = &google_protobuf2.Any{}
			}
			if err := m.Info.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTask(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTask
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTask(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTask
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTask
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTask
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthTask
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowTask
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipTask(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthTask = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTask   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("github.com/containerd/containerd/api/types/task/task.proto", fileDescriptorTask)
}

var fileDescriptorTask = []byte{
	// 543 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xcf, 0x6e, 0xd3, 0x4c,
	0x14, 0xc5, 0x33, 0x6e, 0xe3, 0x24, 0xe3, 0xb6, 0x9f, 0x3f, 0x13, 0x55, 0xc6, 0x20, 0xdb, 0xea,
	0xca, 0x62, 0x61, 0x8b, 0x74, 0xc7, 0x2e, 0xff, 0x84, 0x2c, 0x24, 0x37, 0x72, 0x12, 0xb1, 0x8c,
	0x9c, 0x78, 0x62, 0x46, 0x6d, 0x66, 0x2c, 0x7b, 0x0c, 0x64, 0xc7, 0x12, 0x75, 0xc5, 0x0b, 0x74,
	0x05, 0x4f, 0xc1, 0x13, 0x64, 0xc9, 0x0a, 0xb1, 0x0a, 0xd4, 0x4f, 0x82, 0xc6, 0x76, 0xd2, 0x08,
	0xd8, 0x8c, 0xee, 0x3d, 0xbf, 0x33, 0x77, 0xee, 0x1c, 0xf8, 0x22, 0xc2, 0xec, 0x4d, 0x36, 0xb7,
	0x17, 0x74, 0xe5, 0x2c, 0x28, 0x61, 0x01, 0x26, 0x28, 0x09, 0x0f, 0xcb, 0x20, 0xc6, 0x0e, 0x5b,
	0xc7, 0x28, 0x75, 0x58, 0x90, 0x5e, 0x17, 0x87, 0x1d, 0x27, 0x94, 0x51, 0xe5, 0xd1, 0x83, 0xcb,
	0x7e, 0xfb, 0xdc, 0x2e, 0x4c, 0x5a, 0x3b, 0xa2, 0x11, 0x2d, 0xb8, 0xc3, 0xab, 0xd2, 0xaa, 0x19,
	0x11, 0xa5, 0xd1, 0x0d, 0x72, 0x8a, 0x6e, 0x9e, 0x2d, 0x1d, 0x86, 0x57, 0x28, 0x65, 0xc1, 0x2a,
	0xae, 0x0c, 0x8f, 0xff, 0x34, 0x04, 0x64, 0x5d, 0xa2, 0x8b, 0x5c, 0x80, 0x8d, 0x51, 0x42, 0x17,
	0x28, 0x4d, 0x95, 0x0e, 0x3c, 0xd9, 0x3f, 0x3a, 0xc3, 0xa1, 0x0a, 0x4c, 0x60, 0xb5, 0x7a, 0xff,
	0xe5, 0x5b, 0x43, 0xea, 0xef, 0x74, 0x77, 0xe0, 0x4b, 0x7b, 0x93, 0x1b, 0x2a, 0xe7, 0x50, 0xc0,
	0xa1, 0x2a, 0x14, 0x4e, 0x31, 0xdf, 0x1a, 0x82, 0x3b, 0xf0, 0x05, 0x1c, 0x2a, 0x32, 0x3c, 0x8a,
	0x71, 0xa8, 0x1e, 0x99, 0xc0, 0x3a, 0xf5, 0x79, 0xa9, 0x5c, 0x42, 0x31, 0x65, 0x01, 0xcb, 0x52,
	0xf5, 0xd8, 0x04, 0xd6, 0x59, 0xe7, 0x89, 0xfd, 0x8f, 0x1f, 0xda, 0xe3, 0xc2, 0xe2, 0x57, 0x56,
	0xa5, 0x0d, 0xeb, 0x29, 0x0b, 0x31, 0x51, 0xeb, 0xfc, 0x05, 0xbf, 0x6c, 0x94, 0x73, 0x3e, 0x2a,
	0xa4, 0x19, 0x53, 0xc5, 0x42, 0xae, 0xba, 0x4a, 0x47, 0x49, 0xa2, 0x36, 0xf6, 0x3a, 0x4a, 0x12,
	0x45, 0x83, 0x4d, 0x86, 0x92, 0x15, 0x26, 0xc1, 0x8d, 0xda, 0x34, 0x81, 0xd5, 0xf4, 0xf7, 0xbd,
	0x62, 0x40, 0x09, 0xbd, 0xc7, 0x6c, 0x56, 0xed, 0xd6, 0x2a, 0x16, 0x86, 0x5c, 0x2a, 0x57, 0x51,
	0xba, 0xb0, 0xc5, 0x3b, 0x14, 0xce, 0x02, 0xa6, 0x42, 0x13, 0x58, 0x52, 0x47, 0xb3, 0xcb, 0x40,
	0xed, 0x5d, 0xa0, 0xf6, 0x64, 0x97, 0x78, 0xaf, 0xb9, 0xd9, 0x1a, 0xb5, 0x4f, 0x3f, 0x0d, 0xe0,
	0x37, 0xcb, 0x6b, 0x5d, 0x76, 0xe1, 0x42, 0xa9, 0xca, 0xd8, 0x25, 0x4b, 0xba, 0xcb, 0x06, 0x3c,
	0x64, 0x63, 0xc1, 0x63, 0x4c, 0x96, 0xb4, 0xc8, 0x51, 0xea, 0xb4, 0xff, 0x1a, 0xdf, 0x25, 0x6b,
	0xbf, 0x70, 0x3c, 0xfb, 0x0e, 0xa0, 0x58, 0x2d, 0xa6, 0xc3, 0xc6, 0xd4, 0x7b, 0xe5, 0x5d, 0xbd,
	0xf6, 0xe4, 0x9a, 0xf6, 0xff, 0xed, 0x9d, 0x79, 0x5a, 0x82, 0x29, 0xb9, 0x26, 0xf4, 0x1d, 0xe1,
	0xbc, 0xef, 0x0f, 0xbb, 0x93, 0xe1, 0x40, 0x06, 0x87, 0xbc, 0x9f, 0xa0, 0x80, 0xa1, 0x90, 0x73,
	0x7f, 0xea, 0x79, 0xae, 0xf7, 0x52, 0x16, 0x0e, 0xb9, 0x9f, 0x11, 0x82, 0x49, 0xc4, 0xf9, 0x78,
	0x72, 0x35, 0x1a, 0x0d, 0x07, 0xf2, 0xd1, 0x21, 0x1f, 0x33, 0x1a, 0xc7, 0x28, 0x54, 0x9e, 0x42,
	0x71, 0xd4, 0x9d, 0x8e, 0x87, 0x03, 0xf9, 0x58, 0x93, 0x6f, 0xef, 0xcc, 0x93, 0x12, 0x8f, 0x82,
	0x2c, 0x2d, 0xa7, 0x73, 0xca, 0xa7, 0xd7, 0x0f, 0x6f, 0x73, 0x8c, 0x49, 0xa4, 0x9d, 0x7d, 0xfc,
	0xac, 0xd7, 0xbe, 0x7e, 0xd1, 0xab, 0xdf, 0xf4, 0xd4, 0xcd, 0xbd, 0x5e, 0xfb, 0x71, 0xaf, 0xd7,
	0x3e, 0xe4, 0x3a, 0xd8, 0xe4, 0x3a, 0xf8, 0x96, 0xeb, 0xe0, 0x57, 0xae, 0x83, 0xb9, 0x58, 0xc4,
	0x70, 0xf9, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x19, 0xf7, 0x5b, 0x8f, 0x4e, 0x03, 0x00, 0x00,
}
