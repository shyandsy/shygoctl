package spec

import "encoding/json"

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

//////////////////////////////////////////////////////////////////////

func (p DefineStruct) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Action string `json:"action"` // 注入类型标识
		Type
	}{
		Action: "DefineStruct", // 类型标识值
		Type:   (Type)(p),
	})
}
func (p NestedStruct) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Action string `json:"action"` // 注入类型标识
		Type
	}{
		Action: "NestedStruct", // 类型标识值
		Type:   (Type)(p),
	})
}
func (p PrimitiveType) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Action string `json:"action"` // 注入类型标识
		Type
	}{
		Action: "PrimitiveType", // 类型标识值
		Type:   (Type)(p),
	})
}
func (p QualifiedType) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Action string `json:"action"` // 注入类型标识
		Type
	}{
		Action: "QualifiedType", // 类型标识值
		Type:   (Type)(p),
	})
}
func (p MapType) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Action string `json:"action"` // 注入类型标识
		Type
	}{
		Action: "MapType", // 类型标识值
		Type:   (Type)(p),
	})
}
func (p ArrayType) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Action string `json:"action"` // 注入类型标识
		Type
	}{
		Action: "ArrayType", // 类型标识值
		Type:   (Type)(p),
	})
}
func (p InterfaceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Action string `json:"action"` // 注入类型标识
		Type
	}{
		Action: "InterfaceType", // 类型标识值
		Type:   (Type)(p),
	})
}
func (p PointerType) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Action string `json:"action"` // 注入类型标识
		Type
	}{
		Action: "PointerType", // 类型标识值
		Type:   (Type)(p),
	})
}

/////////////////////////////////////////////////////////////////////////

func (p *Member) UnmarshalJSON(data []byte) error {
	//to: Type    Type
	// implements:
	// - DefineStruct
	// - NestedStruct
	// - PrimitiveType
	// - QualifiedType
	// - MapType
	// - ArrayType
	// - InterfaceType
	// - PointerType
	type tempMember struct {
		Name        string
		Tag         string
		Comment     string
		Docs        Doc
		IsInline    bool
		TypePayload *json.RawMessage `json:"Type"` //Type Type
	}
	temp := tempMember{}
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	switch temp["method"] {
	case "credit_card":
		cc := &CreditCard{}
		if err := json.Unmarshal(data, cc); err != nil {
			return err
		}
		*p = cc
	case "paypal":
		pp := &PayPal{}
		if err := json.Unmarshal(data, pp); err != nil {
			return err
		}
		*p = pp
	}
	return nil
}

func (p *Route) UnmarshalJSON(data []byte) error {
	// to: RequestType        Type
	// to: ResponseType       Type
	// implements:
	// - DefineStruct
	// - NestedStruct
	// - PrimitiveType
	// - QualifiedType
	// - MapType
	// - ArrayType
	// - InterfaceType
	// - PointerType
	var temp map[string]interface{}
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	switch temp["method"] {
	case "credit_card":
		cc := &CreditCard{}
		if err := json.Unmarshal(data, cc); err != nil {
			return err
		}
		*p = cc
	case "paypal":
		pp := &PayPal{}
		if err := json.Unmarshal(data, pp); err != nil {
			return err
		}
		*p = pp
	}
	return nil
}

func (p *MapType) UnmarshalJSON(data []byte) error {
	//to: Value Type
	// implements:
	// - DefineStruct
	// - NestedStruct
	// - PrimitiveType
	// - QualifiedType
	// - MapType
	// - ArrayType
	// - InterfaceType
	// - PointerType
	var temp map[string]interface{}
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	switch temp["method"] {
	case "credit_card":
		cc := &CreditCard{}
		if err := json.Unmarshal(data, cc); err != nil {
			return err
		}
		*p = cc
	case "paypal":
		pp := &PayPal{}
		if err := json.Unmarshal(data, pp); err != nil {
			return err
		}
		*p = pp
	}
	return nil
}

func (p *ArrayType) UnmarshalJSON(data []byte) error {
	// to: Value   Type
	// implements:
	// - DefineStruct
	// - NestedStruct
	// - PrimitiveType
	// - QualifiedType
	// - MapType
	// - ArrayType
	// - InterfaceType
	// - PointerType
	var temp map[string]interface{}
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	switch temp["method"] {
	case "credit_card":
		cc := &CreditCard{}
		if err := json.Unmarshal(data, cc); err != nil {
			return err
		}
		*p = cc
	case "paypal":
		pp := &PayPal{}
		if err := json.Unmarshal(data, pp); err != nil {
			return err
		}
		*p = pp
	}
	return nil
}

func (p *PointerType) UnmarshalJSON(data []byte) error {
	//to: Type    Type
	// implements:
	// - DefineStruct
	// - NestedStruct
	// - PrimitiveType
	// - QualifiedType
	// - MapType
	// - ArrayType
	// - InterfaceType
	// - PointerType
	var temp map[string]interface{}
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	switch temp["method"] {
	case "credit_card":
		cc := &CreditCard{}
		if err := json.Unmarshal(data, cc); err != nil {
			return err
		}
		*p = cc
	case "paypal":
		pp := &PayPal{}
		if err := json.Unmarshal(data, pp); err != nil {
			return err
		}
		*p = pp
	}
	return nil
}
