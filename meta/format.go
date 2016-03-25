package meta

type DataFormat int

// matches list in browse.h
const (
	FMT_BINARY DataFormat = iota + 1
	FMT_BITS
	FMT_BOOLEAN
	FMT_DECIMAL64
	FMT_ENUMERATION
	FMT_IDENTITYREF
	FMT_INSTANCE_IDENTIFIER
	FMT_INT8
	FMT_INT16
	FMT_INT32
	FMT_INT64
	FMT_LEAFREF
	FMT_STRING
	FMT_UINT8
	FMT_UINT16
	FMT_UINT32
	FMT_UINT64
	FMT_UNION
	FMT_ANYDATA
)

const (
	FMT_BINARY_LIST              DataFormat = iota + 1025
	FMT_BITS_LIST                           //1026
	FMT_BOOLEAN_LIST                        //1027
	FMT_DECIMAL64_LIST                      //1028
	FMT_ENUMERATION_LIST                    //1029
	FMT_IDENTITYREF_LIST                    //1030
	FMT_INSTANCE_IDENTIFIER_LIST            //1031
	FMT_INT8_LIST                           //1032
	FMT_INT16_LIST                          //1033
	FMT_INT32_LIST                          //1034
	FMT_INT64_LIST                          //1035
	FMT_LEAFREF_LIST                        //1036
	FMT_STRING_LIST                         //1037
	FMT_UINT8_LIST                          //1038
	FMT_UINT16_LIST                         //1039
	FMT_UINT32_LIST                         //1040
	FMT_UINT64_LIST                         //1041
	FMT_UNION_LIST                          //1042
	FMT_ANYDATA_LIST                        //1043
)

func (f DataFormat) String() string {
	for name, candidate := range internalTypes {
		if f == candidate {
			return name
		}
		if f - 1024 == candidate {
			return name + "-list"
		}
	}
	return "?unknown?"
}

func IsListFormat(f DataFormat) bool {
	return f >= FMT_BINARY_LIST && f <= FMT_UNION_LIST
}

func DataTypeImplicitFormat(typeIdent string) DataFormat {
	return internalTypes[typeIdent]
}

var internalTypes = map[string]DataFormat{
	"binary":              FMT_BINARY,
	"bits":                FMT_BITS,
	"boolean":             FMT_BOOLEAN,
	"decimal64":           FMT_DECIMAL64,
	"enumeration":         FMT_ENUMERATION,
	"identitydef":         FMT_IDENTITYREF,
	"instance-identifier": FMT_INSTANCE_IDENTIFIER,
	"int8":                FMT_INT8,
	"int16":               FMT_INT16,
	"int32":               FMT_INT32,
	"int64":               FMT_INT64,
	"leafref":             FMT_LEAFREF,
	"string":              FMT_STRING,
	"uint8":               FMT_UINT8,
	"uint16":              FMT_UINT16,
	"uint32":              FMT_UINT32,
	"uint64":              FMT_UINT64,
	"union":               FMT_UNION,
	"any":                 FMT_ANYDATA,
}
