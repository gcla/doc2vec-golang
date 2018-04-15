package doc2vec

// NOTE: THIS FILE WAS PRODUCED BY THE
// MSGP CODE GENERATION TOOL (github.com/tinylib/msgp)
// DO NOT EDIT

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *SortItem) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "Idx":
			z.Idx, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "Dis":
			z.Dis, err = dc.ReadFloat64()
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
func (z SortItem) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "Idx"
	err = en.Append(0x82, 0xa3, 0x49, 0x64, 0x78)
	if err != nil {
		return
	}
	err = en.WriteInt32(z.Idx)
	if err != nil {
		return
	}
	// write "Dis"
	err = en.Append(0xa3, 0x44, 0x69, 0x73)
	if err != nil {
		return
	}
	err = en.WriteFloat64(z.Dis)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z SortItem) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "Idx"
	o = append(o, 0x82, 0xa3, 0x49, 0x64, 0x78)
	o = msgp.AppendInt32(o, z.Idx)
	// string "Dis"
	o = append(o, 0xa3, 0x44, 0x69, 0x73)
	o = msgp.AppendFloat64(o, z.Dis)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *SortItem) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "Idx":
			z.Idx, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "Dis":
			z.Dis, bts, err = msgp.ReadFloat64Bytes(bts)
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
func (z SortItem) Msgsize() (s int) {
	s = 1 + 4 + msgp.Int32Size + 4 + msgp.Float64Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *TDoc2VecImpl) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "Dim":
			z.Dim, err = dc.ReadInt()
			if err != nil {
				return
			}
		case "UseCbow":
			z.UseCbow, err = dc.ReadBool()
			if err != nil {
				return
			}
		case "WindowSize":
			z.WindowSize, err = dc.ReadInt()
			if err != nil {
				return
			}
		case "UseHS":
			z.UseHS, err = dc.ReadBool()
			if err != nil {
				return
			}
		case "UseNEG":
			z.UseNEG, err = dc.ReadBool()
			if err != nil {
				return
			}
		case "Negative":
			z.Negative, err = dc.ReadInt()
			if err != nil {
				return
			}
		case "StartAlpha":
			z.StartAlpha, err = dc.ReadFloat64()
			if err != nil {
				return
			}
		case "Iters":
			z.Iters, err = dc.ReadInt()
			if err != nil {
				return
			}
		case "TrainedWords":
			z.TrainedWords, err = dc.ReadInt()
			if err != nil {
				return
			}
		case "Corpus":
			err = z.Corpus.DecodeMsg(dc)
			if err != nil {
				return
			}
		case "NN":
			err = z.NN.DecodeMsg(dc)
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
func (z *TDoc2VecImpl) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 11
	// write "Dim"
	err = en.Append(0x8b, 0xa3, 0x44, 0x69, 0x6d)
	if err != nil {
		return
	}
	err = en.WriteInt(z.Dim)
	if err != nil {
		return
	}
	// write "UseCbow"
	err = en.Append(0xa7, 0x55, 0x73, 0x65, 0x43, 0x62, 0x6f, 0x77)
	if err != nil {
		return
	}
	err = en.WriteBool(z.UseCbow)
	if err != nil {
		return
	}
	// write "WindowSize"
	err = en.Append(0xaa, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x53, 0x69, 0x7a, 0x65)
	if err != nil {
		return
	}
	err = en.WriteInt(z.WindowSize)
	if err != nil {
		return
	}
	// write "UseHS"
	err = en.Append(0xa5, 0x55, 0x73, 0x65, 0x48, 0x53)
	if err != nil {
		return
	}
	err = en.WriteBool(z.UseHS)
	if err != nil {
		return
	}
	// write "UseNEG"
	err = en.Append(0xa6, 0x55, 0x73, 0x65, 0x4e, 0x45, 0x47)
	if err != nil {
		return
	}
	err = en.WriteBool(z.UseNEG)
	if err != nil {
		return
	}
	// write "Negative"
	err = en.Append(0xa8, 0x4e, 0x65, 0x67, 0x61, 0x74, 0x69, 0x76, 0x65)
	if err != nil {
		return
	}
	err = en.WriteInt(z.Negative)
	if err != nil {
		return
	}
	// write "StartAlpha"
	err = en.Append(0xaa, 0x53, 0x74, 0x61, 0x72, 0x74, 0x41, 0x6c, 0x70, 0x68, 0x61)
	if err != nil {
		return
	}
	err = en.WriteFloat64(z.StartAlpha)
	if err != nil {
		return
	}
	// write "Iters"
	err = en.Append(0xa5, 0x49, 0x74, 0x65, 0x72, 0x73)
	if err != nil {
		return
	}
	err = en.WriteInt(z.Iters)
	if err != nil {
		return
	}
	// write "TrainedWords"
	err = en.Append(0xac, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x64, 0x57, 0x6f, 0x72, 0x64, 0x73)
	if err != nil {
		return
	}
	err = en.WriteInt(z.TrainedWords)
	if err != nil {
		return
	}
	// write "Corpus"
	err = en.Append(0xa6, 0x43, 0x6f, 0x72, 0x70, 0x75, 0x73)
	if err != nil {
		return
	}
	err = z.Corpus.EncodeMsg(en)
	if err != nil {
		return
	}
	// write "NN"
	err = en.Append(0xa2, 0x4e, 0x4e)
	if err != nil {
		return
	}
	err = z.NN.EncodeMsg(en)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *TDoc2VecImpl) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 11
	// string "Dim"
	o = append(o, 0x8b, 0xa3, 0x44, 0x69, 0x6d)
	o = msgp.AppendInt(o, z.Dim)
	// string "UseCbow"
	o = append(o, 0xa7, 0x55, 0x73, 0x65, 0x43, 0x62, 0x6f, 0x77)
	o = msgp.AppendBool(o, z.UseCbow)
	// string "WindowSize"
	o = append(o, 0xaa, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x53, 0x69, 0x7a, 0x65)
	o = msgp.AppendInt(o, z.WindowSize)
	// string "UseHS"
	o = append(o, 0xa5, 0x55, 0x73, 0x65, 0x48, 0x53)
	o = msgp.AppendBool(o, z.UseHS)
	// string "UseNEG"
	o = append(o, 0xa6, 0x55, 0x73, 0x65, 0x4e, 0x45, 0x47)
	o = msgp.AppendBool(o, z.UseNEG)
	// string "Negative"
	o = append(o, 0xa8, 0x4e, 0x65, 0x67, 0x61, 0x74, 0x69, 0x76, 0x65)
	o = msgp.AppendInt(o, z.Negative)
	// string "StartAlpha"
	o = append(o, 0xaa, 0x53, 0x74, 0x61, 0x72, 0x74, 0x41, 0x6c, 0x70, 0x68, 0x61)
	o = msgp.AppendFloat64(o, z.StartAlpha)
	// string "Iters"
	o = append(o, 0xa5, 0x49, 0x74, 0x65, 0x72, 0x73)
	o = msgp.AppendInt(o, z.Iters)
	// string "TrainedWords"
	o = append(o, 0xac, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x64, 0x57, 0x6f, 0x72, 0x64, 0x73)
	o = msgp.AppendInt(o, z.TrainedWords)
	// string "Corpus"
	o = append(o, 0xa6, 0x43, 0x6f, 0x72, 0x70, 0x75, 0x73)
	o, err = z.Corpus.MarshalMsg(o)
	if err != nil {
		return
	}
	// string "NN"
	o = append(o, 0xa2, 0x4e, 0x4e)
	o, err = z.NN.MarshalMsg(o)
	if err != nil {
		return
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *TDoc2VecImpl) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "Dim":
			z.Dim, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "UseCbow":
			z.UseCbow, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				return
			}
		case "WindowSize":
			z.WindowSize, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "UseHS":
			z.UseHS, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				return
			}
		case "UseNEG":
			z.UseNEG, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				return
			}
		case "Negative":
			z.Negative, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "StartAlpha":
			z.StartAlpha, bts, err = msgp.ReadFloat64Bytes(bts)
			if err != nil {
				return
			}
		case "Iters":
			z.Iters, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "TrainedWords":
			z.TrainedWords, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "Corpus":
			bts, err = z.Corpus.UnmarshalMsg(bts)
			if err != nil {
				return
			}
		case "NN":
			bts, err = z.NN.UnmarshalMsg(bts)
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
func (z *TDoc2VecImpl) Msgsize() (s int) {
	s = 1 + 4 + msgp.IntSize + 8 + msgp.BoolSize + 11 + msgp.IntSize + 6 + msgp.BoolSize + 7 + msgp.BoolSize + 9 + msgp.IntSize + 11 + msgp.Float64Size + 6 + msgp.IntSize + 13 + msgp.IntSize + 7 + z.Corpus.Msgsize() + 3 + z.NN.Msgsize()
	return
}

// DecodeMsg implements msgp.Decodable
func (z *TSortItemSlice) DecodeMsg(dc *msgp.Reader) (err error) {
	var zb0002 uint32
	zb0002, err = dc.ReadArrayHeader()
	if err != nil {
		return
	}
	if cap((*z)) >= int(zb0002) {
		(*z) = (*z)[:zb0002]
	} else {
		(*z) = make(TSortItemSlice, zb0002)
	}
	for zb0001 := range *z {
		if dc.IsNil() {
			err = dc.ReadNil()
			if err != nil {
				return
			}
			(*z)[zb0001] = nil
		} else {
			if (*z)[zb0001] == nil {
				(*z)[zb0001] = new(SortItem)
			}
			var field []byte
			_ = field
			var zb0003 uint32
			zb0003, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			for zb0003 > 0 {
				zb0003--
				field, err = dc.ReadMapKeyPtr()
				if err != nil {
					return
				}
				switch msgp.UnsafeString(field) {
				case "Idx":
					(*z)[zb0001].Idx, err = dc.ReadInt32()
					if err != nil {
						return
					}
				case "Dis":
					(*z)[zb0001].Dis, err = dc.ReadFloat64()
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
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z TSortItemSlice) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteArrayHeader(uint32(len(z)))
	if err != nil {
		return
	}
	for zb0004 := range z {
		if z[zb0004] == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			// map header, size 2
			// write "Idx"
			err = en.Append(0x82, 0xa3, 0x49, 0x64, 0x78)
			if err != nil {
				return
			}
			err = en.WriteInt32(z[zb0004].Idx)
			if err != nil {
				return
			}
			// write "Dis"
			err = en.Append(0xa3, 0x44, 0x69, 0x73)
			if err != nil {
				return
			}
			err = en.WriteFloat64(z[zb0004].Dis)
			if err != nil {
				return
			}
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z TSortItemSlice) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendArrayHeader(o, uint32(len(z)))
	for zb0004 := range z {
		if z[zb0004] == nil {
			o = msgp.AppendNil(o)
		} else {
			// map header, size 2
			// string "Idx"
			o = append(o, 0x82, 0xa3, 0x49, 0x64, 0x78)
			o = msgp.AppendInt32(o, z[zb0004].Idx)
			// string "Dis"
			o = append(o, 0xa3, 0x44, 0x69, 0x73)
			o = msgp.AppendFloat64(o, z[zb0004].Dis)
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *TSortItemSlice) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0002 uint32
	zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		return
	}
	if cap((*z)) >= int(zb0002) {
		(*z) = (*z)[:zb0002]
	} else {
		(*z) = make(TSortItemSlice, zb0002)
	}
	for zb0001 := range *z {
		if msgp.IsNil(bts) {
			bts, err = msgp.ReadNilBytes(bts)
			if err != nil {
				return
			}
			(*z)[zb0001] = nil
		} else {
			if (*z)[zb0001] == nil {
				(*z)[zb0001] = new(SortItem)
			}
			var field []byte
			_ = field
			var zb0003 uint32
			zb0003, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				return
			}
			for zb0003 > 0 {
				zb0003--
				field, bts, err = msgp.ReadMapKeyZC(bts)
				if err != nil {
					return
				}
				switch msgp.UnsafeString(field) {
				case "Idx":
					(*z)[zb0001].Idx, bts, err = msgp.ReadInt32Bytes(bts)
					if err != nil {
						return
					}
				case "Dis":
					(*z)[zb0001].Dis, bts, err = msgp.ReadFloat64Bytes(bts)
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
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z TSortItemSlice) Msgsize() (s int) {
	s = msgp.ArrayHeaderSize
	for zb0004 := range z {
		if z[zb0004] == nil {
			s += msgp.NilSize
		} else {
			s += 1 + 4 + msgp.Int32Size + 4 + msgp.Float64Size
		}
	}
	return
}
