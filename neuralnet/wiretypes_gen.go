package neuralnet

// NOTE: THIS FILE WAS PRODUCED BY THE
// MSGP CODE GENERATION TOOL (github.com/tinylib/msgp)
// DO NOT EDIT

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *TNeuralNetImpl) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "Syn0":
			var zb0002 uint32
			zb0002, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Syn0) >= int(zb0002) {
				z.Syn0 = (z.Syn0)[:zb0002]
			} else {
				z.Syn0 = make([]TVector, zb0002)
			}
			for za0001 := range z.Syn0 {
				var zb0003 uint32
				zb0003, err = dc.ReadArrayHeader()
				if err != nil {
					return
				}
				if cap(z.Syn0[za0001]) >= int(zb0003) {
					z.Syn0[za0001] = (z.Syn0[za0001])[:zb0003]
				} else {
					z.Syn0[za0001] = make(TVector, zb0003)
				}
				for za0002 := range z.Syn0[za0001] {
					z.Syn0[za0001][za0002], err = dc.ReadFloat32()
					if err != nil {
						return
					}
				}
			}
		case "Dsyn0":
			var zb0004 uint32
			zb0004, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Dsyn0) >= int(zb0004) {
				z.Dsyn0 = (z.Dsyn0)[:zb0004]
			} else {
				z.Dsyn0 = make([]TVector, zb0004)
			}
			for za0003 := range z.Dsyn0 {
				var zb0005 uint32
				zb0005, err = dc.ReadArrayHeader()
				if err != nil {
					return
				}
				if cap(z.Dsyn0[za0003]) >= int(zb0005) {
					z.Dsyn0[za0003] = (z.Dsyn0[za0003])[:zb0005]
				} else {
					z.Dsyn0[za0003] = make(TVector, zb0005)
				}
				for za0004 := range z.Dsyn0[za0003] {
					z.Dsyn0[za0003][za0004], err = dc.ReadFloat32()
					if err != nil {
						return
					}
				}
			}
		case "Syn1":
			var zb0006 uint32
			zb0006, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Syn1) >= int(zb0006) {
				z.Syn1 = (z.Syn1)[:zb0006]
			} else {
				z.Syn1 = make([]TVector, zb0006)
			}
			for za0005 := range z.Syn1 {
				var zb0007 uint32
				zb0007, err = dc.ReadArrayHeader()
				if err != nil {
					return
				}
				if cap(z.Syn1[za0005]) >= int(zb0007) {
					z.Syn1[za0005] = (z.Syn1[za0005])[:zb0007]
				} else {
					z.Syn1[za0005] = make(TVector, zb0007)
				}
				for za0006 := range z.Syn1[za0005] {
					z.Syn1[za0005][za0006], err = dc.ReadFloat32()
					if err != nil {
						return
					}
				}
			}
		case "Syn1neg":
			var zb0008 uint32
			zb0008, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Syn1neg) >= int(zb0008) {
				z.Syn1neg = (z.Syn1neg)[:zb0008]
			} else {
				z.Syn1neg = make([]TVector, zb0008)
			}
			for za0007 := range z.Syn1neg {
				var zb0009 uint32
				zb0009, err = dc.ReadArrayHeader()
				if err != nil {
					return
				}
				if cap(z.Syn1neg[za0007]) >= int(zb0009) {
					z.Syn1neg[za0007] = (z.Syn1neg[za0007])[:zb0009]
				} else {
					z.Syn1neg[za0007] = make(TVector, zb0009)
				}
				for za0008 := range z.Syn1neg[za0007] {
					z.Syn1neg[za0007][za0008], err = dc.ReadFloat32()
					if err != nil {
						return
					}
				}
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
func (z *TNeuralNetImpl) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 4
	// write "Syn0"
	err = en.Append(0x84, 0xa4, 0x53, 0x79, 0x6e, 0x30)
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.Syn0)))
	if err != nil {
		return
	}
	for za0001 := range z.Syn0 {
		err = en.WriteArrayHeader(uint32(len(z.Syn0[za0001])))
		if err != nil {
			return
		}
		for za0002 := range z.Syn0[za0001] {
			err = en.WriteFloat32(z.Syn0[za0001][za0002])
			if err != nil {
				return
			}
		}
	}
	// write "Dsyn0"
	err = en.Append(0xa5, 0x44, 0x73, 0x79, 0x6e, 0x30)
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.Dsyn0)))
	if err != nil {
		return
	}
	for za0003 := range z.Dsyn0 {
		err = en.WriteArrayHeader(uint32(len(z.Dsyn0[za0003])))
		if err != nil {
			return
		}
		for za0004 := range z.Dsyn0[za0003] {
			err = en.WriteFloat32(z.Dsyn0[za0003][za0004])
			if err != nil {
				return
			}
		}
	}
	// write "Syn1"
	err = en.Append(0xa4, 0x53, 0x79, 0x6e, 0x31)
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.Syn1)))
	if err != nil {
		return
	}
	for za0005 := range z.Syn1 {
		err = en.WriteArrayHeader(uint32(len(z.Syn1[za0005])))
		if err != nil {
			return
		}
		for za0006 := range z.Syn1[za0005] {
			err = en.WriteFloat32(z.Syn1[za0005][za0006])
			if err != nil {
				return
			}
		}
	}
	// write "Syn1neg"
	err = en.Append(0xa7, 0x53, 0x79, 0x6e, 0x31, 0x6e, 0x65, 0x67)
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.Syn1neg)))
	if err != nil {
		return
	}
	for za0007 := range z.Syn1neg {
		err = en.WriteArrayHeader(uint32(len(z.Syn1neg[za0007])))
		if err != nil {
			return
		}
		for za0008 := range z.Syn1neg[za0007] {
			err = en.WriteFloat32(z.Syn1neg[za0007][za0008])
			if err != nil {
				return
			}
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *TNeuralNetImpl) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 4
	// string "Syn0"
	o = append(o, 0x84, 0xa4, 0x53, 0x79, 0x6e, 0x30)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Syn0)))
	for za0001 := range z.Syn0 {
		o = msgp.AppendArrayHeader(o, uint32(len(z.Syn0[za0001])))
		for za0002 := range z.Syn0[za0001] {
			o = msgp.AppendFloat32(o, z.Syn0[za0001][za0002])
		}
	}
	// string "Dsyn0"
	o = append(o, 0xa5, 0x44, 0x73, 0x79, 0x6e, 0x30)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Dsyn0)))
	for za0003 := range z.Dsyn0 {
		o = msgp.AppendArrayHeader(o, uint32(len(z.Dsyn0[za0003])))
		for za0004 := range z.Dsyn0[za0003] {
			o = msgp.AppendFloat32(o, z.Dsyn0[za0003][za0004])
		}
	}
	// string "Syn1"
	o = append(o, 0xa4, 0x53, 0x79, 0x6e, 0x31)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Syn1)))
	for za0005 := range z.Syn1 {
		o = msgp.AppendArrayHeader(o, uint32(len(z.Syn1[za0005])))
		for za0006 := range z.Syn1[za0005] {
			o = msgp.AppendFloat32(o, z.Syn1[za0005][za0006])
		}
	}
	// string "Syn1neg"
	o = append(o, 0xa7, 0x53, 0x79, 0x6e, 0x31, 0x6e, 0x65, 0x67)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Syn1neg)))
	for za0007 := range z.Syn1neg {
		o = msgp.AppendArrayHeader(o, uint32(len(z.Syn1neg[za0007])))
		for za0008 := range z.Syn1neg[za0007] {
			o = msgp.AppendFloat32(o, z.Syn1neg[za0007][za0008])
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *TNeuralNetImpl) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "Syn0":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Syn0) >= int(zb0002) {
				z.Syn0 = (z.Syn0)[:zb0002]
			} else {
				z.Syn0 = make([]TVector, zb0002)
			}
			for za0001 := range z.Syn0 {
				var zb0003 uint32
				zb0003, bts, err = msgp.ReadArrayHeaderBytes(bts)
				if err != nil {
					return
				}
				if cap(z.Syn0[za0001]) >= int(zb0003) {
					z.Syn0[za0001] = (z.Syn0[za0001])[:zb0003]
				} else {
					z.Syn0[za0001] = make(TVector, zb0003)
				}
				for za0002 := range z.Syn0[za0001] {
					z.Syn0[za0001][za0002], bts, err = msgp.ReadFloat32Bytes(bts)
					if err != nil {
						return
					}
				}
			}
		case "Dsyn0":
			var zb0004 uint32
			zb0004, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Dsyn0) >= int(zb0004) {
				z.Dsyn0 = (z.Dsyn0)[:zb0004]
			} else {
				z.Dsyn0 = make([]TVector, zb0004)
			}
			for za0003 := range z.Dsyn0 {
				var zb0005 uint32
				zb0005, bts, err = msgp.ReadArrayHeaderBytes(bts)
				if err != nil {
					return
				}
				if cap(z.Dsyn0[za0003]) >= int(zb0005) {
					z.Dsyn0[za0003] = (z.Dsyn0[za0003])[:zb0005]
				} else {
					z.Dsyn0[za0003] = make(TVector, zb0005)
				}
				for za0004 := range z.Dsyn0[za0003] {
					z.Dsyn0[za0003][za0004], bts, err = msgp.ReadFloat32Bytes(bts)
					if err != nil {
						return
					}
				}
			}
		case "Syn1":
			var zb0006 uint32
			zb0006, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Syn1) >= int(zb0006) {
				z.Syn1 = (z.Syn1)[:zb0006]
			} else {
				z.Syn1 = make([]TVector, zb0006)
			}
			for za0005 := range z.Syn1 {
				var zb0007 uint32
				zb0007, bts, err = msgp.ReadArrayHeaderBytes(bts)
				if err != nil {
					return
				}
				if cap(z.Syn1[za0005]) >= int(zb0007) {
					z.Syn1[za0005] = (z.Syn1[za0005])[:zb0007]
				} else {
					z.Syn1[za0005] = make(TVector, zb0007)
				}
				for za0006 := range z.Syn1[za0005] {
					z.Syn1[za0005][za0006], bts, err = msgp.ReadFloat32Bytes(bts)
					if err != nil {
						return
					}
				}
			}
		case "Syn1neg":
			var zb0008 uint32
			zb0008, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Syn1neg) >= int(zb0008) {
				z.Syn1neg = (z.Syn1neg)[:zb0008]
			} else {
				z.Syn1neg = make([]TVector, zb0008)
			}
			for za0007 := range z.Syn1neg {
				var zb0009 uint32
				zb0009, bts, err = msgp.ReadArrayHeaderBytes(bts)
				if err != nil {
					return
				}
				if cap(z.Syn1neg[za0007]) >= int(zb0009) {
					z.Syn1neg[za0007] = (z.Syn1neg[za0007])[:zb0009]
				} else {
					z.Syn1neg[za0007] = make(TVector, zb0009)
				}
				for za0008 := range z.Syn1neg[za0007] {
					z.Syn1neg[za0007][za0008], bts, err = msgp.ReadFloat32Bytes(bts)
					if err != nil {
						return
					}
				}
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
func (z *TNeuralNetImpl) Msgsize() (s int) {
	s = 1 + 5 + msgp.ArrayHeaderSize
	for za0001 := range z.Syn0 {
		s += msgp.ArrayHeaderSize + (len(z.Syn0[za0001]) * (msgp.Float32Size))
	}
	s += 6 + msgp.ArrayHeaderSize
	for za0003 := range z.Dsyn0 {
		s += msgp.ArrayHeaderSize + (len(z.Dsyn0[za0003]) * (msgp.Float32Size))
	}
	s += 5 + msgp.ArrayHeaderSize
	for za0005 := range z.Syn1 {
		s += msgp.ArrayHeaderSize + (len(z.Syn1[za0005]) * (msgp.Float32Size))
	}
	s += 8 + msgp.ArrayHeaderSize
	for za0007 := range z.Syn1neg {
		s += msgp.ArrayHeaderSize + (len(z.Syn1neg[za0007]) * (msgp.Float32Size))
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *TVector) DecodeMsg(dc *msgp.Reader) (err error) {
	var zb0002 uint32
	zb0002, err = dc.ReadArrayHeader()
	if err != nil {
		return
	}
	if cap((*z)) >= int(zb0002) {
		(*z) = (*z)[:zb0002]
	} else {
		(*z) = make(TVector, zb0002)
	}
	for zb0001 := range *z {
		(*z)[zb0001], err = dc.ReadFloat32()
		if err != nil {
			return
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z TVector) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteArrayHeader(uint32(len(z)))
	if err != nil {
		return
	}
	for zb0003 := range z {
		err = en.WriteFloat32(z[zb0003])
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z TVector) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendArrayHeader(o, uint32(len(z)))
	for zb0003 := range z {
		o = msgp.AppendFloat32(o, z[zb0003])
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *TVector) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0002 uint32
	zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		return
	}
	if cap((*z)) >= int(zb0002) {
		(*z) = (*z)[:zb0002]
	} else {
		(*z) = make(TVector, zb0002)
	}
	for zb0001 := range *z {
		(*z)[zb0001], bts, err = msgp.ReadFloat32Bytes(bts)
		if err != nil {
			return
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z TVector) Msgsize() (s int) {
	s = msgp.ArrayHeaderSize + (len(z) * (msgp.Float32Size))
	return
}
