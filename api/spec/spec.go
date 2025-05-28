package spec

import (
	"encoding/json"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
)

// RoutePrefixKey is the prefix keyword for the routes.
const RoutePrefixKey = "prefix"

type (
	// Doc describes document
	Doc []string

	// Annotation defines key-value
	Annotation struct {
		Properties map[string]string
	}

	// ApiSyntax describes the syntax grammar
	ApiSyntax struct {
		Version string
		Doc     Doc
		Comment Doc
	}

	// ApiSpec describes an api file
	ApiSpec struct {
		Info    Info
		Syntax  ApiSyntax // Deprecated: useless expression
		Imports []Import  // Deprecated: useless expression
		Types   []Type
		Service Service
	}

	// Import describes api import
	Import struct {
		Value   string
		Doc     Doc
		Comment Doc
	}

	// Group defines a set of routing information
	Group struct {
		Annotation Annotation
		Routes     []Route
	}

	// Info describes info grammar block
	Info struct {
		// Deprecated: use Properties instead
		Title string
		// Deprecated: use Properties instead
		Desc string
		// Deprecated: use Properties instead
		Version string
		// Deprecated: use Properties instead
		Author string
		// Deprecated: use Properties instead
		Email      string
		Properties map[string]string
	}

	// Member describes the field of a structure
	Member struct {
		Name string
		// data type, for example, string、map[int]string、[]int64、[]*User
		Type    Type
		Tag     string
		Comment string
		// document for the field
		Docs     Doc
		IsInline bool
	}

	// Route describes api route
	Route struct {
		// Deprecated: Use Service AtServer instead.
		AtServerAnnotation Annotation
		Method             string
		Path               string
		RequestType        Type
		ResponseType       Type
		Docs               Doc
		Handler            string
		AtDoc              AtDoc
		HandlerDoc         Doc
		HandlerComment     Doc
		Doc                Doc
		Comment            Doc
	}

	// Service describes api service
	Service struct {
		Name   string
		Groups []Group
	}

	// Type defines api type
	Type interface {
		Name() string
		Comments() []string
		Documents() []string
	}

	// DefineStruct describes api structure
	DefineStruct struct {
		RawName string
		Members []Member
		Docs    Doc
	}

	// NestedStruct describes a structure nested in structure.
	NestedStruct struct {
		RawName string
		Members []Member
		Docs    Doc
	}

	// PrimitiveType describes the basic golang type, such as bool,int32,int64, ...
	PrimitiveType struct {
		RawName string
	}

	QualifiedType struct {
		PackageName string
		RawName     string
	}

	// MapType describes a map for api
	MapType struct {
		RawName string
		// only support the PrimitiveType
		Key string
		// it can be asserted as PrimitiveType: int、bool、
		// PointerType: *string、*User、
		// MapType: map[${PrimitiveType}]interface、
		// ArrayType:[]int、[]User、[]*User
		// InterfaceType: interface{}
		// Type
		Value Type
	}

	// ArrayType describes a slice for api
	ArrayType struct {
		RawName string
		Value   Type
	}

	// InterfaceType describes an interface for api
	InterfaceType struct {
		RawName string
	}

	// PointerType describes a pointer for api
	PointerType struct {
		RawName string
		Type    Type
	}

	// AtDoc describes a metadata for api grammar: @doc(...)
	AtDoc struct {
		Properties map[string]string
		Text       string
	}
)

/////////////////////////////////////////////////////////////////////////

func (p DefineStruct) MarshalJSON() ([]byte, error) {
	type Alias DefineStruct
	return json.Marshal(&struct {
		Kind string `json:"kind"` // 提取类型标识
		Alias
	}{
		Kind:  "DefineStruct", // 类型标识值
		Alias: (Alias)(p),
	})
}

func (p NestedStruct) MarshalJSON() ([]byte, error) {
	type Alias NestedStruct
	return json.Marshal(&struct {
		Kind string `json:"kind"` // 提取类型标识
		Alias
	}{
		Kind:  "NestedStruct", // 类型标识值
		Alias: (Alias)(p),
	})
}

func (p PrimitiveType) MarshalJSON() ([]byte, error) {
	type Alias PrimitiveType
	return json.Marshal(&struct {
		Kind string `json:"kind"` // 提取类型标识
		Alias
	}{
		Kind:  "PrimitiveType", // 类型标识值
		Alias: (Alias)(p),
	})
}

func (p QualifiedType) MarshalJSON() ([]byte, error) {
	type Alias QualifiedType
	return json.Marshal(&struct {
		Kind string `json:"kind"` // 提取类型标识
		Alias
	}{
		Kind:  "QualifiedType", // 类型标识值
		Alias: (Alias)(p),
	})
}

func (p MapType) MarshalJSON() ([]byte, error) {
	type Alias MapType
	return json.Marshal(&struct {
		Kind string `json:"kind"` // 提取类型标识
		Alias
	}{
		Kind:  "MapType", // 类型标识值
		Alias: (Alias)(p),
	})
}

func (p ArrayType) MarshalJSON() ([]byte, error) {
	type Alias ArrayType
	return json.Marshal(&struct {
		Kind string `json:"kind"` // 提取类型标识
		Alias
	}{
		Kind:  "ArrayType", // 类型标识值
		Alias: (Alias)(p),
	})
}

func (p InterfaceType) MarshalJSON() ([]byte, error) {
	type Alias InterfaceType
	return json.Marshal(&struct {
		Kind string `json:"kind"` // 提取类型标识
		Alias
	}{
		Kind:  "InterfaceType", // 类型标识值
		Alias: (Alias)(p),
	})
}

func (p PointerType) MarshalJSON() ([]byte, error) {
	type Alias PointerType
	return json.Marshal(&struct {
		Kind string `json:"kind"` // 提取类型标识
		Alias
	}{
		Kind:  "PointerType", // 类型标识值
		Alias: (Alias)(p),
	})
}

/////////////////////////////////////////////////////////////////////////

func (p *ApiSpec) UnmarshalJSON(data []byte) error {
	type Alias ApiSpec // 创建别名避免递归调用
	aux := &struct {
		*Alias
		Types []json.RawMessage //`json:"types"`  // 延迟解析原始数据[8](@ref)
	}{
		Alias: (*Alias)(p),
	}

	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}

	p.Types = make([]Type, 0)
	for _, raw := range aux.Types {
		tp, err := createType(raw)
		if err != nil {
			return err
		}
		p.Types = append(p.Types, tp)
	}
	return nil
}

func (p *Route) UnmarshalJSON(data []byte) error {
	type Alias Route // 创建别名避免递归调用
	aux := &struct {
		*Alias
		RequestType  json.RawMessage `json:"RequestType"`
		ResponseType json.RawMessage `json:"ResponseType"`
	}{
		Alias: (*Alias)(p),
	}
	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}

	tp, err := createType(aux.RequestType)
	if err != nil {
		return err
	}
	p.RequestType = tp

	tp, err = createType(aux.ResponseType)
	if err != nil {
		return err
	}
	p.ResponseType = tp

	return nil
}

func (p *Member) UnmarshalJSON(data []byte) error {
	type Alias Member // 创建别名避免递归调用
	aux := &struct {
		*Alias
		Type json.RawMessage
	}{
		Alias: (*Alias)(p),
	}
	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}

	tp, err := createType(aux.Type)
	if err != nil {
		return err
	}
	p.Type = tp

	return nil
}

func (p *MapType) UnmarshalJSON(data []byte) error {
	type Alias MapType
	aux := &struct {
		*Alias
		Value json.RawMessage
	}{
		Alias: (*Alias)(p),
	}
	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}

	tp, err := createType(aux.Value)
	if err != nil {
		return err
	}
	p.Value = tp

	return nil
}

func (p *ArrayType) UnmarshalJSON(data []byte) error {
	type Alias ArrayType
	aux := &struct {
		*Alias
		Value json.RawMessage
	}{
		Alias: (*Alias)(p),
	}
	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}

	tp, err := createType(aux.Value)
	if err != nil {
		return err
	}
	p.Value = tp

	return nil
}

func createType(raw json.RawMessage) (Type, error) {
	var base struct {
		Kind string `json:"kind"`
	}

	// process RequestType
	if err := json.Unmarshal(raw, &base); err != nil {
		return nil, err
	}

	switch base.Kind {
	case "DefineStruct":
		var tp DefineStruct
		if err := json.Unmarshal(raw, &tp); err != nil {
			return nil, err
		}
		return tp, nil
	case "NestedStruct":
		var tp NestedStruct
		if err := json.Unmarshal(raw, &tp); err != nil {
			return nil, err
		}
		return tp, nil
	case "PrimitiveType":
		var tp PrimitiveType
		if err := json.Unmarshal(raw, &tp); err != nil {
			return nil, err
		}
		return tp, nil
	case "QualifiedType":
		var tp QualifiedType
		if err := json.Unmarshal(raw, &tp); err != nil {
			return nil, err
		}
		return tp, nil
	case "MapType":
		var tp MapType
		if err := json.Unmarshal(raw, &tp); err != nil {
			return nil, err
		}
		return tp, nil
	case "ArrayType":
		var tp ArrayType
		if err := json.Unmarshal(raw, &tp); err != nil {
			return nil, err
		}
		return tp, nil
	case "InterfaceType":
		var tp InterfaceType
		if err := json.Unmarshal(raw, &tp); err != nil {
			return nil, err
		}
		return tp, nil
	case "PointerType":
		var tp QualifiedType
		if err := json.Unmarshal(raw, &tp); err != nil {
			return nil, err
		}
		return tp, nil
	default:
		logx.Error("unknown type: %s", base.Kind)
		return nil, fmt.Errorf("unknown type: %s", base.Kind)
	}
}
