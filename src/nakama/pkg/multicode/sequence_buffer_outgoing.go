// Copyright 2017 The Nakama Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package multicode

type PacketSet struct {
	messageIDs []uint16
}

type SequenceBufferOutgoing struct {
	sequence      uint16
	numEntries    uint16
	entrySequence []uint32
	entryData     []*PacketSet
}

func NewSequenceBufferOutgoing(bufferSize int) *SequenceBufferOutgoing {
	s := &SequenceBufferOutgoing{
		sequence:      0,
		numEntries:    uint16(bufferSize),
		entrySequence: make([]uint32, bufferSize),
		entryData:     make([]*PacketSet, bufferSize),
	}
	for i := 0; i < bufferSize; i++ {
		s.entrySequence[i] = NULL_SEQUENCE
		s.entryData[i] = &PacketSet{}
	}
	return s
}

func (s *SequenceBufferOutgoing) Find(sequence uint16) *PacketSet {
	index := sequence % s.numEntries
	sequenceNum := s.entrySequence[index]
	if sequenceNum == uint32(sequence) {
		return s.entryData[index]
	}
	return nil
}

func (s *SequenceBufferOutgoing) Insert(sequence uint16) *PacketSet {
	if SequenceLessThan(sequence, s.sequence-s.numEntries) {
		return nil
	}

	if SequenceGreaterThan(sequence+1, s.sequence) {
		s.RemoveEntries(int32(s.sequence), int32(sequence))
		s.sequence = sequence + 1
	}

	index := sequence % s.numEntries
	s.entrySequence[index] = uint32(sequence)
	return s.entryData[index]
}

func (s *SequenceBufferOutgoing) RemoveEntries(startSequence, finishSequence int32) {
	if finishSequence < startSequence {
		finishSequence += 65536
	}
	if uint16(finishSequence-startSequence) < s.numEntries {
		for sequence := startSequence; sequence <= finishSequence; sequence++ {
			s.entrySequence[uint16(sequence)%s.numEntries] = NULL_SEQUENCE
		}
	} else {
		for i := uint16(0); i < s.numEntries; i++ {
			s.entrySequence[i] = NULL_SEQUENCE
		}
	}
}
