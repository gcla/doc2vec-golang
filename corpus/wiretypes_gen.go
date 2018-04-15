package corpus

// NOTE: THIS FILE WAS PRODUCED BY THE
// MSGP CODE GENERATION TOOL (github.com/tinylib/msgp)
// DO NOT EDIT

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *TCorpusImpl) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Words":
			var zb0002 uint32
			zb0002, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Words) >= int(zb0002) {
				z.Words = (z.Words)[:zb0002]
			} else {
				z.Words = make(TWordItemSlice, zb0002)
			}
			for za0001 := range z.Words {
				err = z.Words[za0001].DecodeMsg(dc)
				if err != nil {
					return
				}
			}
		case "Word2Idx":
			var zb0003 uint32
			zb0003, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.Word2Idx == nil && zb0003 > 0 {
				z.Word2Idx = make(map[string]int32, zb0003)
			} else if len(z.Word2Idx) > 0 {
				for key, _ := range z.Word2Idx {
					delete(z.Word2Idx, key)
				}
			}
			for zb0003 > 0 {
				zb0003--
				var za0002 string
				var za0003 int32
				za0002, err = dc.ReadString()
				if err != nil {
					return
				}
				za0003, err = dc.ReadInt32()
				if err != nil {
					return
				}
				z.Word2Idx[za0002] = za0003
			}
		case "Doc2WordsIdx":
			var zb0004 uint32
			zb0004, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Doc2WordsIdx) >= int(zb0004) {
				z.Doc2WordsIdx = (z.Doc2WordsIdx)[:zb0004]
			} else {
				z.Doc2WordsIdx = make([][]int32, zb0004)
			}
			for za0004 := range z.Doc2WordsIdx {
				var zb0005 uint32
				zb0005, err = dc.ReadArrayHeader()
				if err != nil {
					return
				}
				if cap(z.Doc2WordsIdx[za0004]) >= int(zb0005) {
					z.Doc2WordsIdx[za0004] = (z.Doc2WordsIdx[za0004])[:zb0005]
				} else {
					z.Doc2WordsIdx[za0004] = make([]int32, zb0005)
				}
				for za0005 := range z.Doc2WordsIdx[za0004] {
					z.Doc2WordsIdx[za0004][za0005], err = dc.ReadInt32()
					if err != nil {
						return
					}
				}
			}
		case "Doc2Idx":
			var zb0006 uint32
			zb0006, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.Doc2Idx == nil && zb0006 > 0 {
				z.Doc2Idx = make(map[string]int32, zb0006)
			} else if len(z.Doc2Idx) > 0 {
				for key, _ := range z.Doc2Idx {
					delete(z.Doc2Idx, key)
				}
			}
			for zb0006 > 0 {
				zb0006--
				var za0006 string
				var za0007 int32
				za0006, err = dc.ReadString()
				if err != nil {
					return
				}
				za0007, err = dc.ReadInt32()
				if err != nil {
					return
				}
				z.Doc2Idx[za0006] = za0007
			}
		case "MinReduce":
			z.MinReduce, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "MinCnt":
			z.MinCnt, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "WordsCnt":
			z.WordsCnt, err = dc.ReadInt()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *TCorpusImpl) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 7
	// write "Words"
	err = en.Append(0x87, 0xa5, 0x57, 0x6f, 0x72, 0x64, 0x73)
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.Words)))
	if err != nil {
		return
	}
	for za0001 := range z.Words {
		err = z.Words[za0001].EncodeMsg(en)
		if err != nil {
			return
		}
	}
	// write "Word2Idx"
	err = en.Append(0xa8, 0x57, 0x6f, 0x72, 0x64, 0x32, 0x49, 0x64, 0x78)
	if err != nil {
		return
	}
	err = en.WriteMapHeader(uint32(len(z.Word2Idx)))
	if err != nil {
		return
	}
	for za0002, za0003 := range z.Word2Idx {
		err = en.WriteString(za0002)
		if err != nil {
			return
		}
		err = en.WriteInt32(za0003)
		if err != nil {
			return
		}
	}
	// write "Doc2WordsIdx"
	err = en.Append(0xac, 0x44, 0x6f, 0x63, 0x32, 0x57, 0x6f, 0x72, 0x64, 0x73, 0x49, 0x64, 0x78)
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.Doc2WordsIdx)))
	if err != nil {
		return
	}
	for za0004 := range z.Doc2WordsIdx {
		err = en.WriteArrayHeader(uint32(len(z.Doc2WordsIdx[za0004])))
		if err != nil {
			return
		}
		for za0005 := range z.Doc2WordsIdx[za0004] {
			err = en.WriteInt32(z.Doc2WordsIdx[za0004][za0005])
			if err != nil {
				return
			}
		}
	}
	// write "Doc2Idx"
	err = en.Append(0xa7, 0x44, 0x6f, 0x63, 0x32, 0x49, 0x64, 0x78)
	if err != nil {
		return
	}
	err = en.WriteMapHeader(uint32(len(z.Doc2Idx)))
	if err != nil {
		return
	}
	for za0006, za0007 := range z.Doc2Idx {
		err = en.WriteString(za0006)
		if err != nil {
			return
		}
		err = en.WriteInt32(za0007)
		if err != nil {
			return
		}
	}
	// write "MinReduce"
	err = en.Append(0xa9, 0x4d, 0x69, 0x6e, 0x52, 0x65, 0x64, 0x75, 0x63, 0x65)
	if err != nil {
		return
	}
	err = en.WriteInt32(z.MinReduce)
	if err != nil {
		return
	}
	// write "MinCnt"
	err = en.Append(0xa6, 0x4d, 0x69, 0x6e, 0x43, 0x6e, 0x74)
	if err != nil {
		return
	}
	err = en.WriteInt32(z.MinCnt)
	if err != nil {
		return
	}
	// write "WordsCnt"
	err = en.Append(0xa8, 0x57, 0x6f, 0x72, 0x64, 0x73, 0x43, 0x6e, 0x74)
	if err != nil {
		return
	}
	err = en.WriteInt(z.WordsCnt)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *TCorpusImpl) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 7
	// string "Words"
	o = append(o, 0x87, 0xa5, 0x57, 0x6f, 0x72, 0x64, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Words)))
	for za0001 := range z.Words {
		o, err = z.Words[za0001].MarshalMsg(o)
		if err != nil {
			return
		}
	}
	// string "Word2Idx"
	o = append(o, 0xa8, 0x57, 0x6f, 0x72, 0x64, 0x32, 0x49, 0x64, 0x78)
	o = msgp.AppendMapHeader(o, uint32(len(z.Word2Idx)))
	for za0002, za0003 := range z.Word2Idx {
		o = msgp.AppendString(o, za0002)
		o = msgp.AppendInt32(o, za0003)
	}
	// string "Doc2WordsIdx"
	o = append(o, 0xac, 0x44, 0x6f, 0x63, 0x32, 0x57, 0x6f, 0x72, 0x64, 0x73, 0x49, 0x64, 0x78)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Doc2WordsIdx)))
	for za0004 := range z.Doc2WordsIdx {
		o = msgp.AppendArrayHeader(o, uint32(len(z.Doc2WordsIdx[za0004])))
		for za0005 := range z.Doc2WordsIdx[za0004] {
			o = msgp.AppendInt32(o, z.Doc2WordsIdx[za0004][za0005])
		}
	}
	// string "Doc2Idx"
	o = append(o, 0xa7, 0x44, 0x6f, 0x63, 0x32, 0x49, 0x64, 0x78)
	o = msgp.AppendMapHeader(o, uint32(len(z.Doc2Idx)))
	for za0006, za0007 := range z.Doc2Idx {
		o = msgp.AppendString(o, za0006)
		o = msgp.AppendInt32(o, za0007)
	}
	// string "MinReduce"
	o = append(o, 0xa9, 0x4d, 0x69, 0x6e, 0x52, 0x65, 0x64, 0x75, 0x63, 0x65)
	o = msgp.AppendInt32(o, z.MinReduce)
	// string "MinCnt"
	o = append(o, 0xa6, 0x4d, 0x69, 0x6e, 0x43, 0x6e, 0x74)
	o = msgp.AppendInt32(o, z.MinCnt)
	// string "WordsCnt"
	o = append(o, 0xa8, 0x57, 0x6f, 0x72, 0x64, 0x73, 0x43, 0x6e, 0x74)
	o = msgp.AppendInt(o, z.WordsCnt)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *TCorpusImpl) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Words":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Words) >= int(zb0002) {
				z.Words = (z.Words)[:zb0002]
			} else {
				z.Words = make(TWordItemSlice, zb0002)
			}
			for za0001 := range z.Words {
				bts, err = z.Words[za0001].UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		case "Word2Idx":
			var zb0003 uint32
			zb0003, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				return
			}
			if z.Word2Idx == nil && zb0003 > 0 {
				z.Word2Idx = make(map[string]int32, zb0003)
			} else if len(z.Word2Idx) > 0 {
				for key, _ := range z.Word2Idx {
					delete(z.Word2Idx, key)
				}
			}
			for zb0003 > 0 {
				var za0002 string
				var za0003 int32
				zb0003--
				za0002, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				za0003, bts, err = msgp.ReadInt32Bytes(bts)
				if err != nil {
					return
				}
				z.Word2Idx[za0002] = za0003
			}
		case "Doc2WordsIdx":
			var zb0004 uint32
			zb0004, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Doc2WordsIdx) >= int(zb0004) {
				z.Doc2WordsIdx = (z.Doc2WordsIdx)[:zb0004]
			} else {
				z.Doc2WordsIdx = make([][]int32, zb0004)
			}
			for za0004 := range z.Doc2WordsIdx {
				var zb0005 uint32
				zb0005, bts, err = msgp.ReadArrayHeaderBytes(bts)
				if err != nil {
					return
				}
				if cap(z.Doc2WordsIdx[za0004]) >= int(zb0005) {
					z.Doc2WordsIdx[za0004] = (z.Doc2WordsIdx[za0004])[:zb0005]
				} else {
					z.Doc2WordsIdx[za0004] = make([]int32, zb0005)
				}
				for za0005 := range z.Doc2WordsIdx[za0004] {
					z.Doc2WordsIdx[za0004][za0005], bts, err = msgp.ReadInt32Bytes(bts)
					if err != nil {
						return
					}
				}
			}
		case "Doc2Idx":
			var zb0006 uint32
			zb0006, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				return
			}
			if z.Doc2Idx == nil && zb0006 > 0 {
				z.Doc2Idx = make(map[string]int32, zb0006)
			} else if len(z.Doc2Idx) > 0 {
				for key, _ := range z.Doc2Idx {
					delete(z.Doc2Idx, key)
				}
			}
			for zb0006 > 0 {
				var za0006 string
				var za0007 int32
				zb0006--
				za0006, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				za0007, bts, err = msgp.ReadInt32Bytes(bts)
				if err != nil {
					return
				}
				z.Doc2Idx[za0006] = za0007
			}
		case "MinReduce":
			z.MinReduce, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "MinCnt":
			z.MinCnt, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "WordsCnt":
			z.WordsCnt, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *TCorpusImpl) Msgsize() (s int) {
	s = 1 + 6 + msgp.ArrayHeaderSize
	for za0001 := range z.Words {
		s += z.Words[za0001].Msgsize()
	}
	s += 9 + msgp.MapHeaderSize
	if z.Word2Idx != nil {
		for za0002, za0003 := range z.Word2Idx {
			_ = za0003
			s += msgp.StringPrefixSize + len(za0002) + msgp.Int32Size
		}
	}
	s += 13 + msgp.ArrayHeaderSize
	for za0004 := range z.Doc2WordsIdx {
		s += msgp.ArrayHeaderSize + (len(z.Doc2WordsIdx[za0004]) * (msgp.Int32Size))
	}
	s += 8 + msgp.MapHeaderSize
	if z.Doc2Idx != nil {
		for za0006, za0007 := range z.Doc2Idx {
			_ = za0007
			s += msgp.StringPrefixSize + len(za0006) + msgp.Int32Size
		}
	}
	s += 10 + msgp.Int32Size + 7 + msgp.Int32Size + 9 + msgp.IntSize
	return
}

// DecodeMsg implements msgp.Decodable
func (z *TWordItem) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Cnt":
			z.Cnt, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "Point":
			var zb0002 uint32
			zb0002, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Point) >= int(zb0002) {
				z.Point = (z.Point)[:zb0002]
			} else {
				z.Point = make([]int32, zb0002)
			}
			for za0001 := range z.Point {
				z.Point[za0001], err = dc.ReadInt32()
				if err != nil {
					return
				}
			}
		case "Code":
			var zb0003 uint32
			zb0003, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Code) >= int(zb0003) {
				z.Code = (z.Code)[:zb0003]
			} else {
				z.Code = make([]bool, zb0003)
			}
			for za0002 := range z.Code {
				z.Code[za0002], err = dc.ReadBool()
				if err != nil {
					return
				}
			}
		case "Word":
			z.Word, err = dc.ReadString()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *TWordItem) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 4
	// write "Cnt"
	err = en.Append(0x84, 0xa3, 0x43, 0x6e, 0x74)
	if err != nil {
		return
	}
	err = en.WriteInt32(z.Cnt)
	if err != nil {
		return
	}
	// write "Point"
	err = en.Append(0xa5, 0x50, 0x6f, 0x69, 0x6e, 0x74)
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.Point)))
	if err != nil {
		return
	}
	for za0001 := range z.Point {
		err = en.WriteInt32(z.Point[za0001])
		if err != nil {
			return
		}
	}
	// write "Code"
	err = en.Append(0xa4, 0x43, 0x6f, 0x64, 0x65)
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.Code)))
	if err != nil {
		return
	}
	for za0002 := range z.Code {
		err = en.WriteBool(z.Code[za0002])
		if err != nil {
			return
		}
	}
	// write "Word"
	err = en.Append(0xa4, 0x57, 0x6f, 0x72, 0x64)
	if err != nil {
		return
	}
	err = en.WriteString(z.Word)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *TWordItem) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 4
	// string "Cnt"
	o = append(o, 0x84, 0xa3, 0x43, 0x6e, 0x74)
	o = msgp.AppendInt32(o, z.Cnt)
	// string "Point"
	o = append(o, 0xa5, 0x50, 0x6f, 0x69, 0x6e, 0x74)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Point)))
	for za0001 := range z.Point {
		o = msgp.AppendInt32(o, z.Point[za0001])
	}
	// string "Code"
	o = append(o, 0xa4, 0x43, 0x6f, 0x64, 0x65)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Code)))
	for za0002 := range z.Code {
		o = msgp.AppendBool(o, z.Code[za0002])
	}
	// string "Word"
	o = append(o, 0xa4, 0x57, 0x6f, 0x72, 0x64)
	o = msgp.AppendString(o, z.Word)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *TWordItem) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Cnt":
			z.Cnt, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "Point":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Point) >= int(zb0002) {
				z.Point = (z.Point)[:zb0002]
			} else {
				z.Point = make([]int32, zb0002)
			}
			for za0001 := range z.Point {
				z.Point[za0001], bts, err = msgp.ReadInt32Bytes(bts)
				if err != nil {
					return
				}
			}
		case "Code":
			var zb0003 uint32
			zb0003, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Code) >= int(zb0003) {
				z.Code = (z.Code)[:zb0003]
			} else {
				z.Code = make([]bool, zb0003)
			}
			for za0002 := range z.Code {
				z.Code[za0002], bts, err = msgp.ReadBoolBytes(bts)
				if err != nil {
					return
				}
			}
		case "Word":
			z.Word, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *TWordItem) Msgsize() (s int) {
	s = 1 + 4 + msgp.Int32Size + 6 + msgp.ArrayHeaderSize + (len(z.Point) * (msgp.Int32Size)) + 5 + msgp.ArrayHeaderSize + (len(z.Code) * (msgp.BoolSize)) + 5 + msgp.StringPrefixSize + len(z.Word)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *TWordItemSlice) DecodeMsg(dc *msgp.Reader) (err error) {
	var zb0002 uint32
	zb0002, err = dc.ReadArrayHeader()
	if err != nil {
		return
	}
	if cap((*z)) >= int(zb0002) {
		(*z) = (*z)[:zb0002]
	} else {
		(*z) = make(TWordItemSlice, zb0002)
	}
	for zb0001 := range *z {
		err = (*z)[zb0001].DecodeMsg(dc)
		if err != nil {
			return
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z TWordItemSlice) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteArrayHeader(uint32(len(z)))
	if err != nil {
		return
	}
	for zb0003 := range z {
		err = z[zb0003].EncodeMsg(en)
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z TWordItemSlice) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendArrayHeader(o, uint32(len(z)))
	for zb0003 := range z {
		o, err = z[zb0003].MarshalMsg(o)
		if err != nil {
			return
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *TWordItemSlice) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0002 uint32
	zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		return
	}
	if cap((*z)) >= int(zb0002) {
		(*z) = (*z)[:zb0002]
	} else {
		(*z) = make(TWordItemSlice, zb0002)
	}
	for zb0001 := range *z {
		bts, err = (*z)[zb0001].UnmarshalMsg(bts)
		if err != nil {
			return
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z TWordItemSlice) Msgsize() (s int) {
	s = msgp.ArrayHeaderSize
	for zb0003 := range z {
		s += z[zb0003].Msgsize()
	}
	return
}
