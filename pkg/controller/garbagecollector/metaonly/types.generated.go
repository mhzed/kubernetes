/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// ************************************************************
// DO NOT EDIT.
// THIS FILE IS AUTO-GENERATED BY codecgen.
// ************************************************************

package metaonly

import (
	"errors"
	"fmt"
	codec1978 "github.com/ugorji/go/codec"
	pkg2_v1 "k8s.io/kubernetes/pkg/api/v1"
	pkg1_v1 "k8s.io/kubernetes/pkg/apis/meta/v1"
	pkg3_types "k8s.io/kubernetes/pkg/types"
	"reflect"
	"runtime"
	time "time"
)

const (
	// ----- content types ----
	codecSelferC_UTF81234 = 1
	codecSelferC_RAW1234  = 0
	// ----- value types used ----
	codecSelferValueTypeArray1234 = 10
	codecSelferValueTypeMap1234   = 9
	// ----- containerStateValues ----
	codecSelfer_containerMapKey1234    = 2
	codecSelfer_containerMapValue1234  = 3
	codecSelfer_containerMapEnd1234    = 4
	codecSelfer_containerArrayElem1234 = 6
	codecSelfer_containerArrayEnd1234  = 7
)

var (
	codecSelferBitsize1234                         = uint8(reflect.TypeOf(uint(0)).Bits())
	codecSelferOnlyMapOrArrayEncodeToStructErr1234 = errors.New(`only encoded map or array can be decoded into a struct`)
)

type codecSelfer1234 struct{}

func init() {
	if codec1978.GenVersion != 5 {
		_, file, _, _ := runtime.Caller(0)
		err := fmt.Errorf("codecgen version mismatch: current: %v, need %v. Re-generate file: %v",
			5, codec1978.GenVersion, file)
		panic(err)
	}
	if false { // reference the types, but skip this branch at build/run time
		var v0 pkg2_v1.ObjectMeta
		var v1 pkg1_v1.TypeMeta
		var v2 pkg3_types.UID
		var v3 time.Time
		_, _, _, _ = v0, v1, v2, v3
	}
}

func (x *MetadataOnlyObject) CodecEncodeSelf(e *codec1978.Encoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		yym1 := z.EncBinary()
		_ = yym1
		if false {
		} else if z.HasExtensions() && z.EncExt(x) {
		} else {
			yysep2 := !z.EncBinary()
			yy2arr2 := z.EncBasicHandle().StructToArray
			var yyq2 [3]bool
			_, _, _ = yysep2, yyq2, yy2arr2
			const yyr2 bool = false
			yyq2[0] = x.Kind != ""
			yyq2[1] = x.APIVersion != ""
			yyq2[2] = true
			var yynn2 int
			if yyr2 || yy2arr2 {
				r.EncodeArrayStart(3)
			} else {
				yynn2 = 0
				for _, b := range yyq2 {
					if b {
						yynn2++
					}
				}
				r.EncodeMapStart(yynn2)
				yynn2 = 0
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq2[0] {
					yym4 := z.EncBinary()
					_ = yym4
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.Kind))
					}
				} else {
					r.EncodeString(codecSelferC_UTF81234, "")
				}
			} else {
				if yyq2[0] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("kind"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yym5 := z.EncBinary()
					_ = yym5
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.Kind))
					}
				}
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq2[1] {
					yym7 := z.EncBinary()
					_ = yym7
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.APIVersion))
					}
				} else {
					r.EncodeString(codecSelferC_UTF81234, "")
				}
			} else {
				if yyq2[1] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("apiVersion"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yym8 := z.EncBinary()
					_ = yym8
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.APIVersion))
					}
				}
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq2[2] {
					yy10 := &x.ObjectMeta
					yy10.CodecEncodeSelf(e)
				} else {
					r.EncodeNil()
				}
			} else {
				if yyq2[2] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("metadata"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yy11 := &x.ObjectMeta
					yy11.CodecEncodeSelf(e)
				}
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayEnd1234)
			} else {
				z.EncSendContainerState(codecSelfer_containerMapEnd1234)
			}
		}
	}
}

func (x *MetadataOnlyObject) CodecDecodeSelf(d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	yym12 := z.DecBinary()
	_ = yym12
	if false {
	} else if z.HasExtensions() && z.DecExt(x) {
	} else {
		yyct13 := r.ContainerType()
		if yyct13 == codecSelferValueTypeMap1234 {
			yyl13 := r.ReadMapStart()
			if yyl13 == 0 {
				z.DecSendContainerState(codecSelfer_containerMapEnd1234)
			} else {
				x.codecDecodeSelfFromMap(yyl13, d)
			}
		} else if yyct13 == codecSelferValueTypeArray1234 {
			yyl13 := r.ReadArrayStart()
			if yyl13 == 0 {
				z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
			} else {
				x.codecDecodeSelfFromArray(yyl13, d)
			}
		} else {
			panic(codecSelferOnlyMapOrArrayEncodeToStructErr1234)
		}
	}
}

func (x *MetadataOnlyObject) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yys14Slc = z.DecScratchBuffer() // default slice to decode into
	_ = yys14Slc
	var yyhl14 bool = l >= 0
	for yyj14 := 0; ; yyj14++ {
		if yyhl14 {
			if yyj14 >= l {
				break
			}
		} else {
			if r.CheckBreak() {
				break
			}
		}
		z.DecSendContainerState(codecSelfer_containerMapKey1234)
		yys14Slc = r.DecodeBytes(yys14Slc, true, true)
		yys14 := string(yys14Slc)
		z.DecSendContainerState(codecSelfer_containerMapValue1234)
		switch yys14 {
		case "kind":
			if r.TryDecodeAsNil() {
				x.Kind = ""
			} else {
				x.Kind = string(r.DecodeString())
			}
		case "apiVersion":
			if r.TryDecodeAsNil() {
				x.APIVersion = ""
			} else {
				x.APIVersion = string(r.DecodeString())
			}
		case "metadata":
			if r.TryDecodeAsNil() {
				x.ObjectMeta = pkg2_v1.ObjectMeta{}
			} else {
				yyv17 := &x.ObjectMeta
				yyv17.CodecDecodeSelf(d)
			}
		default:
			z.DecStructFieldNotFound(-1, yys14)
		} // end switch yys14
	} // end for yyj14
	z.DecSendContainerState(codecSelfer_containerMapEnd1234)
}

func (x *MetadataOnlyObject) codecDecodeSelfFromArray(l int, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yyj18 int
	var yyb18 bool
	var yyhl18 bool = l >= 0
	yyj18++
	if yyhl18 {
		yyb18 = yyj18 > l
	} else {
		yyb18 = r.CheckBreak()
	}
	if yyb18 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.Kind = ""
	} else {
		x.Kind = string(r.DecodeString())
	}
	yyj18++
	if yyhl18 {
		yyb18 = yyj18 > l
	} else {
		yyb18 = r.CheckBreak()
	}
	if yyb18 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.APIVersion = ""
	} else {
		x.APIVersion = string(r.DecodeString())
	}
	yyj18++
	if yyhl18 {
		yyb18 = yyj18 > l
	} else {
		yyb18 = r.CheckBreak()
	}
	if yyb18 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.ObjectMeta = pkg2_v1.ObjectMeta{}
	} else {
		yyv21 := &x.ObjectMeta
		yyv21.CodecDecodeSelf(d)
	}
	for {
		yyj18++
		if yyhl18 {
			yyb18 = yyj18 > l
		} else {
			yyb18 = r.CheckBreak()
		}
		if yyb18 {
			break
		}
		z.DecSendContainerState(codecSelfer_containerArrayElem1234)
		z.DecStructFieldNotFound(yyj18-1, "")
	}
	z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
}

func (x *MetadataOnlyObjectList) CodecEncodeSelf(e *codec1978.Encoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		yym22 := z.EncBinary()
		_ = yym22
		if false {
		} else if z.HasExtensions() && z.EncExt(x) {
		} else {
			yysep23 := !z.EncBinary()
			yy2arr23 := z.EncBasicHandle().StructToArray
			var yyq23 [4]bool
			_, _, _ = yysep23, yyq23, yy2arr23
			const yyr23 bool = false
			yyq23[0] = x.Kind != ""
			yyq23[1] = x.APIVersion != ""
			yyq23[2] = true
			var yynn23 int
			if yyr23 || yy2arr23 {
				r.EncodeArrayStart(4)
			} else {
				yynn23 = 1
				for _, b := range yyq23 {
					if b {
						yynn23++
					}
				}
				r.EncodeMapStart(yynn23)
				yynn23 = 0
			}
			if yyr23 || yy2arr23 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq23[0] {
					yym25 := z.EncBinary()
					_ = yym25
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.Kind))
					}
				} else {
					r.EncodeString(codecSelferC_UTF81234, "")
				}
			} else {
				if yyq23[0] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("kind"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yym26 := z.EncBinary()
					_ = yym26
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.Kind))
					}
				}
			}
			if yyr23 || yy2arr23 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq23[1] {
					yym28 := z.EncBinary()
					_ = yym28
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.APIVersion))
					}
				} else {
					r.EncodeString(codecSelferC_UTF81234, "")
				}
			} else {
				if yyq23[1] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("apiVersion"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yym29 := z.EncBinary()
					_ = yym29
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.APIVersion))
					}
				}
			}
			if yyr23 || yy2arr23 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq23[2] {
					yy31 := &x.ListMeta
					yym32 := z.EncBinary()
					_ = yym32
					if false {
					} else if z.HasExtensions() && z.EncExt(yy31) {
					} else {
						z.EncFallback(yy31)
					}
				} else {
					r.EncodeNil()
				}
			} else {
				if yyq23[2] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("metadata"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yy33 := &x.ListMeta
					yym34 := z.EncBinary()
					_ = yym34
					if false {
					} else if z.HasExtensions() && z.EncExt(yy33) {
					} else {
						z.EncFallback(yy33)
					}
				}
			}
			if yyr23 || yy2arr23 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if x.Items == nil {
					r.EncodeNil()
				} else {
					yym36 := z.EncBinary()
					_ = yym36
					if false {
					} else {
						h.encSliceMetadataOnlyObject(([]MetadataOnlyObject)(x.Items), e)
					}
				}
			} else {
				z.EncSendContainerState(codecSelfer_containerMapKey1234)
				r.EncodeString(codecSelferC_UTF81234, string("items"))
				z.EncSendContainerState(codecSelfer_containerMapValue1234)
				if x.Items == nil {
					r.EncodeNil()
				} else {
					yym37 := z.EncBinary()
					_ = yym37
					if false {
					} else {
						h.encSliceMetadataOnlyObject(([]MetadataOnlyObject)(x.Items), e)
					}
				}
			}
			if yyr23 || yy2arr23 {
				z.EncSendContainerState(codecSelfer_containerArrayEnd1234)
			} else {
				z.EncSendContainerState(codecSelfer_containerMapEnd1234)
			}
		}
	}
}

func (x *MetadataOnlyObjectList) CodecDecodeSelf(d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	yym38 := z.DecBinary()
	_ = yym38
	if false {
	} else if z.HasExtensions() && z.DecExt(x) {
	} else {
		yyct39 := r.ContainerType()
		if yyct39 == codecSelferValueTypeMap1234 {
			yyl39 := r.ReadMapStart()
			if yyl39 == 0 {
				z.DecSendContainerState(codecSelfer_containerMapEnd1234)
			} else {
				x.codecDecodeSelfFromMap(yyl39, d)
			}
		} else if yyct39 == codecSelferValueTypeArray1234 {
			yyl39 := r.ReadArrayStart()
			if yyl39 == 0 {
				z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
			} else {
				x.codecDecodeSelfFromArray(yyl39, d)
			}
		} else {
			panic(codecSelferOnlyMapOrArrayEncodeToStructErr1234)
		}
	}
}

func (x *MetadataOnlyObjectList) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yys40Slc = z.DecScratchBuffer() // default slice to decode into
	_ = yys40Slc
	var yyhl40 bool = l >= 0
	for yyj40 := 0; ; yyj40++ {
		if yyhl40 {
			if yyj40 >= l {
				break
			}
		} else {
			if r.CheckBreak() {
				break
			}
		}
		z.DecSendContainerState(codecSelfer_containerMapKey1234)
		yys40Slc = r.DecodeBytes(yys40Slc, true, true)
		yys40 := string(yys40Slc)
		z.DecSendContainerState(codecSelfer_containerMapValue1234)
		switch yys40 {
		case "kind":
			if r.TryDecodeAsNil() {
				x.Kind = ""
			} else {
				x.Kind = string(r.DecodeString())
			}
		case "apiVersion":
			if r.TryDecodeAsNil() {
				x.APIVersion = ""
			} else {
				x.APIVersion = string(r.DecodeString())
			}
		case "metadata":
			if r.TryDecodeAsNil() {
				x.ListMeta = pkg1_v1.ListMeta{}
			} else {
				yyv43 := &x.ListMeta
				yym44 := z.DecBinary()
				_ = yym44
				if false {
				} else if z.HasExtensions() && z.DecExt(yyv43) {
				} else {
					z.DecFallback(yyv43, false)
				}
			}
		case "items":
			if r.TryDecodeAsNil() {
				x.Items = nil
			} else {
				yyv45 := &x.Items
				yym46 := z.DecBinary()
				_ = yym46
				if false {
				} else {
					h.decSliceMetadataOnlyObject((*[]MetadataOnlyObject)(yyv45), d)
				}
			}
		default:
			z.DecStructFieldNotFound(-1, yys40)
		} // end switch yys40
	} // end for yyj40
	z.DecSendContainerState(codecSelfer_containerMapEnd1234)
}

func (x *MetadataOnlyObjectList) codecDecodeSelfFromArray(l int, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yyj47 int
	var yyb47 bool
	var yyhl47 bool = l >= 0
	yyj47++
	if yyhl47 {
		yyb47 = yyj47 > l
	} else {
		yyb47 = r.CheckBreak()
	}
	if yyb47 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.Kind = ""
	} else {
		x.Kind = string(r.DecodeString())
	}
	yyj47++
	if yyhl47 {
		yyb47 = yyj47 > l
	} else {
		yyb47 = r.CheckBreak()
	}
	if yyb47 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.APIVersion = ""
	} else {
		x.APIVersion = string(r.DecodeString())
	}
	yyj47++
	if yyhl47 {
		yyb47 = yyj47 > l
	} else {
		yyb47 = r.CheckBreak()
	}
	if yyb47 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.ListMeta = pkg1_v1.ListMeta{}
	} else {
		yyv50 := &x.ListMeta
		yym51 := z.DecBinary()
		_ = yym51
		if false {
		} else if z.HasExtensions() && z.DecExt(yyv50) {
		} else {
			z.DecFallback(yyv50, false)
		}
	}
	yyj47++
	if yyhl47 {
		yyb47 = yyj47 > l
	} else {
		yyb47 = r.CheckBreak()
	}
	if yyb47 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.Items = nil
	} else {
		yyv52 := &x.Items
		yym53 := z.DecBinary()
		_ = yym53
		if false {
		} else {
			h.decSliceMetadataOnlyObject((*[]MetadataOnlyObject)(yyv52), d)
		}
	}
	for {
		yyj47++
		if yyhl47 {
			yyb47 = yyj47 > l
		} else {
			yyb47 = r.CheckBreak()
		}
		if yyb47 {
			break
		}
		z.DecSendContainerState(codecSelfer_containerArrayElem1234)
		z.DecStructFieldNotFound(yyj47-1, "")
	}
	z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
}

func (x codecSelfer1234) encSliceMetadataOnlyObject(v []MetadataOnlyObject, e *codec1978.Encoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	r.EncodeArrayStart(len(v))
	for _, yyv54 := range v {
		z.EncSendContainerState(codecSelfer_containerArrayElem1234)
		yy55 := &yyv54
		yy55.CodecEncodeSelf(e)
	}
	z.EncSendContainerState(codecSelfer_containerArrayEnd1234)
}

func (x codecSelfer1234) decSliceMetadataOnlyObject(v *[]MetadataOnlyObject, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r

	yyv56 := *v
	yyh56, yyl56 := z.DecSliceHelperStart()
	var yyc56 bool
	if yyl56 == 0 {
		if yyv56 == nil {
			yyv56 = []MetadataOnlyObject{}
			yyc56 = true
		} else if len(yyv56) != 0 {
			yyv56 = yyv56[:0]
			yyc56 = true
		}
	} else if yyl56 > 0 {
		var yyrr56, yyrl56 int
		var yyrt56 bool
		if yyl56 > cap(yyv56) {

			yyrg56 := len(yyv56) > 0
			yyv256 := yyv56
			yyrl56, yyrt56 = z.DecInferLen(yyl56, z.DecBasicHandle().MaxInitLen, 256)
			if yyrt56 {
				if yyrl56 <= cap(yyv56) {
					yyv56 = yyv56[:yyrl56]
				} else {
					yyv56 = make([]MetadataOnlyObject, yyrl56)
				}
			} else {
				yyv56 = make([]MetadataOnlyObject, yyrl56)
			}
			yyc56 = true
			yyrr56 = len(yyv56)
			if yyrg56 {
				copy(yyv56, yyv256)
			}
		} else if yyl56 != len(yyv56) {
			yyv56 = yyv56[:yyl56]
			yyc56 = true
		}
		yyj56 := 0
		for ; yyj56 < yyrr56; yyj56++ {
			yyh56.ElemContainerState(yyj56)
			if r.TryDecodeAsNil() {
				yyv56[yyj56] = MetadataOnlyObject{}
			} else {
				yyv57 := &yyv56[yyj56]
				yyv57.CodecDecodeSelf(d)
			}

		}
		if yyrt56 {
			for ; yyj56 < yyl56; yyj56++ {
				yyv56 = append(yyv56, MetadataOnlyObject{})
				yyh56.ElemContainerState(yyj56)
				if r.TryDecodeAsNil() {
					yyv56[yyj56] = MetadataOnlyObject{}
				} else {
					yyv58 := &yyv56[yyj56]
					yyv58.CodecDecodeSelf(d)
				}

			}
		}

	} else {
		yyj56 := 0
		for ; !r.CheckBreak(); yyj56++ {

			if yyj56 >= len(yyv56) {
				yyv56 = append(yyv56, MetadataOnlyObject{}) // var yyz56 MetadataOnlyObject
				yyc56 = true
			}
			yyh56.ElemContainerState(yyj56)
			if yyj56 < len(yyv56) {
				if r.TryDecodeAsNil() {
					yyv56[yyj56] = MetadataOnlyObject{}
				} else {
					yyv59 := &yyv56[yyj56]
					yyv59.CodecDecodeSelf(d)
				}

			} else {
				z.DecSwallow()
			}

		}
		if yyj56 < len(yyv56) {
			yyv56 = yyv56[:yyj56]
			yyc56 = true
		} else if yyj56 == 0 && yyv56 == nil {
			yyv56 = []MetadataOnlyObject{}
			yyc56 = true
		}
	}
	yyh56.End()
	if yyc56 {
		*v = yyv56
	}
}
