// Copyright 2015 Tamás Gulácsi
//
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.

package oerr

import (
	"encoding/binary"
	"errors"
	"fmt"
)

type MsgID struct {
	Prefix string
	Code   uint32
}

func (id MsgID) String() string {
	return fmt.Sprintf("%s-%05d", id.Prefix, id.Code)
}
func (id MsgID) MarshalBinary() (data []byte, err error) {
	data = make([]byte, 3+4)
	data[0], data[1], data[2] = id.Prefix[0], id.Prefix[1], id.Prefix[2]
	binary.BigEndian.PutUint32(data[3:], id.Code)
	return
}
func (id *MsgID) UnmarshalBinary(data []byte) error {
	if len(data) < 7 {
		return errors.New("data too short")
	}
	id.Prefix = string(data[:3])
	id.Code = binary.BigEndian.Uint32(data[3:7])
	return nil
}

type MsgData struct {
	Description, Cause, Action string
}

func (m MsgData) String() string {
	return fmt.Sprintf("%s\nCause: %s\nAction: %s", m.Description, m.Cause, m.Action)
}
func (d MsgData) MarshalBinary() (data []byte, err error) {
	vv := []string{d.Description, d.Cause, d.Action}
	n := 3 * 2
	for _, s := range vv {
		n += len(s)
	}
	data = make([]byte, n)
	off := 0
	for _, s := range vv {
		binary.BigEndian.PutUint16(data[off:], uint16(len(s)))
		off += 2
		off += copy(data[off:], []byte(s))
	}
	return data, nil
}
func (d *MsgData) UnmarshalBinary(data []byte) error {
	off := 0
	for _, p := range []*string{&d.Description, &d.Cause, &d.Action} {
		length := int(binary.BigEndian.Uint16(data[off:]))
		off += 2
		*p = string(data[off : off+length])
		off += length
	}
	return nil
}

type Message struct {
	MsgID
	MsgData
}
