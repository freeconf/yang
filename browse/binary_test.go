package browse

import (
	"testing"
	"github.com/blitter/meta/yang"
	"strings"
	"bytes"
	"encoding/hex"
	"github.com/blitter/node"
)

func TestBinaryBrowser(t *testing.T) {
	mstr := `
module m {
	namespace "";
	prefix "";
	revision 0;
	leaf c {
		type string;
	}
	container a {
		container b {
			leaf s {
				type string;
			}
			leaf b {
				type boolean;
			}
			leaf i {
				type int32;
			}
			leaf l {
				type int64;
			}
			leaf d {
				type decimal64;
			}
			leaf e {
				type enumeration {
					enum one;
					enum two;
				}
			}
			leaf-list sl {
				type string;
			}
			leaf-list bl {
				type boolean;
			}
			leaf-list il {
				type int32;
			}
			leaf-list ll {
				type int64;
			}
			leaf-list dl {
				type decimal64;
			}
			leaf-list el {
				type enumeration {
					enum one;
					enum two;
				}
			}
		}
	}
	list p {
		key "k";
		leaf k {
			type string;
		}
		container q {
			leaf s {
				type string;
			}
		}
		list r {
			leaf z {
				type int32;
			}
		}
	}
}
`
	m, err := yang.LoadModuleCustomImport(mstr, nil)
	if err != nil {
		t.Fatal(err)
	}
	tests := []string {
		`{"c":"hello"}`,
		`{"a":{"b":{"s":"waldo","b":true,"i":99,"l":100,"d":1.5,"e":"one"}}}`,
		`{"a":{"b":{"sl":["waldo"],"bl":[true],"il":[99,100],"ll":[100,101],"dl":[1.5,2.5],"el":["one","two"]}}}`,
		`{"p":[{"k":"walter"}]}`,
		`{"p":[{"k":"walter"},{"k":"waldo"},{"k":"weirdo"}]}`,
	}
	c := node.NewContext()
	for _, test := range tests {
		var buff bytes.Buffer
		w := NewBinaryWriter(&buff)
		if err = c.Select(m, w.Node()).InsertFrom(node.NewJsonReader(strings.NewReader(test)).Node()).LastErr; err != nil {
			t.Error(err)
		}
		r := NewBinaryReader(&buff)
		var actualBuff bytes.Buffer
		if err = c.Select(m, r.Node()).InsertInto(node.NewJsonWriter(&actualBuff).Node()).LastErr; err != nil {
			t.Log("\n" + hex.Dump(buff.Bytes()))
			t.Error(err)
		}
		actual := actualBuff.String()
		if test != actual {
			t.Log("\n" + hex.Dump(buff.Bytes()))
			t.Errorf("\nExpected:%s\n  Actual:%s", test, actual)
		}
	}
}

func TestBinaryComplex(t *testing.T) {
	pipe := NewPipe()
	pull, push := pipe.PullPush()
	orig := node.NewJsonReader(strings.NewReader(binaryTestData)).Node()
	inline := NewInline()
	c := node.NewContext()
	onSchemaLoad := make(chan error)
	go func() {
		defer close(onSchemaLoad)
		err := inline.Load(c, orig, push, onSchemaLoad)
		pipe.Close(err)
	}()

	t.Log("wating for meta to load...")
	loadErr := <- onSchemaLoad
	t.Log("meta loaded")
	if loadErr != nil {
		t.Error(loadErr)
	} else {
		var buff bytes.Buffer
		sel := c.Select(inline.DataMeta, pull)
		if err := sel.InsertInto(node.NewJsonWriter(&buff).Node()).LastErr; err != nil {
			t.Error(err)
		}
		t.Log(buff.String())
	}
}

var binaryTestData = `
{
  "meta": {
    "definitions": [
      {
        "ident": "upload",
        "container": {
          "ident": "upload",
          "definitions": [
            {
              "ident": "rows",
              "list": {
                "key": [],
                "ident": "rows",
                "definitions": [
                  {
                    "ident": "student_local_id",
                    "leaf": {
                      "ident": "student_local_id",
                      "config": true,
                      "mandatory": false,
                      "type": {
                        "ident": "string",
                        "path": "",
                        "minLength": 0,
                        "maxLength": 2147483647
                      }
                    }
                  },
                  {
                    "ident": "student_sasid",
                    "leaf": {
                      "ident": "student_sasid",
                      "config": true,
                      "mandatory": false,
                      "type": {
                        "ident": "string",
                        "path": "",
                        "minLength": 0,
                        "maxLength": 2147483647
                      }
                    }
                  },
                  {
                    "ident": "enrollment_status",
                    "leaf": {
                      "ident": "enrollment_status",
                      "config": true,
                      "mandatory": false,
                      "type": {
                        "ident": "string",
                        "path": "",
                        "minLength": 0,
                        "maxLength": 2147483647
                      }
                    }
                  },
                  {
                    "ident": "student_first",
                    "leaf": {
                      "ident": "student_first",
                      "config": true,
                      "mandatory": false,
                      "type": {
                        "ident": "string",
                        "path": "",
                        "minLength": 0,
                        "maxLength": 2147483647
                      }
                    }
                  },
                  {
                    "ident": "student_middle",
                    "leaf": {
                      "ident": "student_middle",
                      "config": true,
                      "mandatory": false,
                      "type": {
                        "ident": "string",
                        "path": "",
                        "minLength": 0,
                        "maxLength": 2147483647
                      }
                    }
                  },
                  {
                    "ident": "student_last",
                    "leaf": {
                      "ident": "student_last",
                      "config": true,
                      "mandatory": false,
                      "type": {
                        "ident": "string",
                        "path": "",
                        "minLength": 0,
                        "maxLength": 2147483647
                      }
                    }
                  },
                  {
                    "ident": "address",
                    "leaf": {
                      "ident": "address",
                      "config": true,
                      "mandatory": false,
                      "type": {
                        "ident": "string",
                        "path": "",
                        "minLength": 0,
                        "maxLength": 2147483647
                      }
                    }
                  },
                  {
                    "ident": "city_state_zip",
                    "leaf": {
                      "ident": "city_state_zip",
                      "config": true,
                      "mandatory": false,
                      "type": {
                        "ident": "string",
                        "path": "",
                        "minLength": 0,
                        "maxLength": 2147483647
                      }
                    }
                  },
                  {
                    "ident": "school_name",
                    "leaf": {
                      "ident": "school_name",
                      "config": true,
                      "mandatory": false,
                      "type": {
                        "ident": "string",
                        "path": "",
                        "minLength": 0,
                        "maxLength": 2147483647
                      }
                    }
                  },
                  {
                    "ident": "school_code",
                    "leaf": {
                      "ident": "school_code",
                      "config": true,
                      "mandatory": false,
                      "type": {
                        "ident": "string",
                        "path": "",
                        "minLength": 0,
                        "maxLength": 2147483647
                      }
                    }
                  },
                  {
                    "ident": "school_id",
                    "leaf": {
                      "ident": "school_id",
                      "config": true,
                      "mandatory": false,
                      "type": {
                        "ident": "string",
                        "path": "",
                        "minLength": 0,
                        "maxLength": 2147483647
                      }
                    }
                  },
                  {
                    "ident": "grade_level",
                    "leaf": {
                      "ident": "grade_level",
                      "config": true,
                      "mandatory": false,
                      "type": {
                        "ident": "string",
                        "path": "",
                        "minLength": 0,
                        "maxLength": 2147483647
                      }
                    }
                  },
                  {
                    "ident": "next_year_school",
                    "leaf": {
                      "ident": "next_year_school",
                      "config": true,
                      "mandatory": false,
                      "type": {
                        "ident": "string",
                        "path": "",
                        "minLength": 0,
                        "maxLength": 2147483647
                      }
                    }
                  },
                  {
                    "ident": "next_grade",
                    "leaf": {
                      "ident": "next_grade",
                      "config": true,
                      "mandatory": false,
                      "type": {
                        "ident": "string",
                        "path": "",
                        "minLength": 0,
                        "maxLength": 2147483647
                      }
                    }
                  },
                  {
                    "ident": "transfer_school",
                    "leaf": {
                      "ident": "transfer_school",
                      "config": true,
                      "mandatory": false,
                      "type": {
                        "ident": "string",
                        "path": "",
                        "minLength": 0,
                        "maxLength": 2147483647
                      }
                    }
                  },
                  {
                    "ident": "trans_Eligible",
                    "leaf": {
                      "ident": "trans_Eligible",
                      "config": true,
                      "mandatory": false,
                      "type": {
                        "ident": "string",
                        "path": "",
                        "minLength": 0,
                        "maxLength": 2147483647
                      }
                    }
                  },
                  {
                    "ident": "trans_Eligible_NextYr",
                    "leaf": {
                      "ident": "trans_Eligible_NextYr",
                      "config": true,
                      "mandatory": false,
                      "type": {
                        "ident": "string",
                        "path": "",
                        "minLength": 0,
                        "maxLength": 2147483647
                      }
                    }
                  },
                  {
                    "ident": "student_dob",
                    "leaf": {
                      "ident": "student_dob",
                      "config": true,
                      "mandatory": false,
                      "type": {
                        "ident": "string",
                        "path": "",
                        "minLength": 0,
                        "maxLength": 2147483647
                      }
                    }
                  },
                  {
                    "ident": "student_home_language",
                    "leaf": {
                      "ident": "student_home_language",
                      "config": true,
                      "mandatory": false,
                      "type": {
                        "ident": "string",
                        "path": "",
                        "minLength": 0,
                        "maxLength": 2147483647
                      }
                    }
                  },
                  {
                    "ident": "home_phone",
                    "leaf": {
                      "ident": "home_phone",
                      "config": true,
                      "mandatory": false,
                      "type": {
                        "ident": "string",
                        "path": "",
                        "minLength": 0,
                        "maxLength": 2147483647
                      }
                    }
                  },
                  {
                    "ident": "student_email",
                    "leaf": {
                      "ident": "student_email",
                      "config": true,
                      "mandatory": false,
                      "type": {
                        "ident": "string",
                        "path": "",
                        "minLength": 0,
                        "maxLength": 2147483647
                      }
                    }
                  },
                  {
                    "ident": "student_mkv_status",
                    "leaf": {
                      "ident": "student_mkv_status",
                      "config": true,
                      "mandatory": false,
                      "type": {
                        "ident": "string",
                        "path": "",
                        "minLength": 0,
                        "maxLength": 2147483647
                      }
                    }
                  },
                  {
                    "ident": "student_mkv_exitDate",
                    "leaf": {
                      "ident": "student_mkv_exitDate",
                      "config": true,
                      "mandatory": false,
                      "type": {
                        "ident": "string",
                        "path": "",
                        "minLength": 0,
                        "maxLength": 2147483647
                      }
                    }
                  },
                  {
                    "ident": "student_frLunchAppStatus",
                    "leaf": {
                      "ident": "student_frLunchAppStatus",
                      "config": true,
                      "mandatory": false,
                      "type": {
                        "ident": "string",
                        "path": "",
                        "minLength": 0,
                        "maxLength": 2147483647
                      }
                    }
                  },
                  {
                    "ident": "student_freereduced",
                    "leaf": {
                      "ident": "student_freereduced",
                      "config": true,
                      "mandatory": false,
                      "type": {
                        "ident": "string",
                        "path": "",
                        "minLength": 0,
                        "maxLength": 2147483647
                      }
                    }
                  },
                  {
                    "ident": "student_last_year_fr",
                    "leaf": {
                      "ident": "student_last_year_fr",
                      "config": true,
                      "mandatory": false,
                      "type": {
                        "ident": "string",
                        "path": "",
                        "minLength": 0,
                        "maxLength": 2147483647
                      }
                    }
                  },
                  {
                    "ident": "Bus_am_route",
                    "leaf": {
                      "ident": "Bus_am_route",
                      "config": true,
                      "mandatory": false,
                      "type": {
                        "ident": "string",
                        "path": "",
                        "minLength": 0,
                        "maxLength": 2147483647
                      }
                    }
                  },
                  {
                    "ident": "Bus_am_description",
                    "leaf": {
                      "ident": "Bus_am_description",
                      "config": true,
                      "mandatory": false,
                      "type": {
                        "ident": "string",
                        "path": "",
                        "minLength": 0,
                        "maxLength": 2147483647
                      }
                    }
                  },
                  {
                    "ident": "Bus_am_time",
                    "leaf": {
                      "ident": "Bus_am_time",
                      "config": true,
                      "mandatory": false,
                      "type": {
                        "ident": "string",
                        "path": "",
                        "minLength": 0,
                        "maxLength": 2147483647
                      }
                    }
                  },
                  {
                    "ident": "Bus_pm_route",
                    "leaf": {
                      "ident": "Bus_pm_route",
                      "config": true,
                      "mandatory": false,
                      "type": {
                        "ident": "string",
                        "path": "",
                        "minLength": 0,
                        "maxLength": 2147483647
                      }
                    }
                  },
                  {
                    "ident": "Bus_pm_description",
                    "leaf": {
                      "ident": "Bus_pm_description",
                      "config": true,
                      "mandatory": false,
                      "type": {
                        "ident": "string",
                        "path": "",
                        "minLength": 0,
                        "maxLength": 2147483647
                      }
                    }
                  },
                  {
                    "ident": "Bus_pm_time",
                    "leaf": {
                      "ident": "Bus_pm_time",
                      "config": true,
                      "mandatory": false,
                      "type": {
                        "ident": "string",
                        "path": "",
                        "minLength": 0,
                        "maxLength": 2147483647
                      }
                    }
                  },
                  {
                    "ident": "bus_tran_am_arrival",
                    "leaf": {
                      "ident": "bus_tran_am_arrival",
                      "config": true,
                      "mandatory": false,
                      "type": {
                        "ident": "string",
                        "path": "",
                        "minLength": 0,
                        "maxLength": 2147483647
                      }
                    }
                  },
                  {
                    "ident": "bus_tran_am_location",
                    "leaf": {
                      "ident": "bus_tran_am_location",
                      "config": true,
                      "mandatory": false,
                      "type": {
                        "ident": "string",
                        "path": "",
                        "minLength": 0,
                        "maxLength": 2147483647
                      }
                    }
                  },
                  {
                    "ident": "bus_tran_am_alt",
                    "leaf": {
                      "ident": "bus_tran_am_alt",
                      "config": true,
                      "mandatory": false,
                      "type": {
                        "ident": "string",
                        "path": "",
                        "minLength": 0,
                        "maxLength": 2147483647
                      }
                    }
                  },
                  {
                    "ident": "bus_tran_pm_arrival",
                    "leaf": {
                      "ident": "bus_tran_pm_arrival",
                      "config": true,
                      "mandatory": false,
                      "type": {
                        "ident": "string",
                        "path": "",
                        "minLength": 0,
                        "maxLength": 2147483647
                      }
                    }
                  },
                  {
                    "ident": "bus_tran_pm_location",
                    "leaf": {
                      "ident": "bus_tran_pm_location",
                      "config": true,
                      "mandatory": false,
                      "type": {
                        "ident": "string",
                        "path": "",
                        "minLength": 0,
                        "maxLength": 2147483647
                      }
                    }
                  },
                  {
                    "ident": "bus_tran_pm_alt",
                    "leaf": {
                      "ident": "bus_tran_pm_alt",
                      "config": true,
                      "mandatory": false,
                      "type": {
                        "ident": "string",
                        "path": "",
                        "minLength": 0,
                        "maxLength": 2147483647
                      }
                    }
                  },
                  {
                    "ident": "parent_1_name",
                    "leaf": {
                      "ident": "parent_1_name",
                      "config": true,
                      "mandatory": false,
                      "type": {
                        "ident": "string",
                        "path": "",
                        "minLength": 0,
                        "maxLength": 2147483647
                      }
                    }
                  },
                  {
                    "ident": "parent_1_emergency_priority",
                    "leaf": {
                      "ident": "parent_1_emergency_priority",
                      "config": true,
                      "mandatory": false,
                      "type": {
                        "ident": "string",
                        "path": "",
                        "minLength": 0,
                        "maxLength": 2147483647
                      }
                    }
                  },
                  {
                    "ident": "parent_1_relationship",
                    "leaf": {
                      "ident": "parent_1_relationship",
                      "config": true,
                      "mandatory": false,
                      "type": {
                        "ident": "string",
                        "path": "",
                        "minLength": 0,
                        "maxLength": 2147483647
                      }
                    }
                  },
                  {
                    "ident": "parent_1_street_address",
                    "leaf": {
                      "ident": "parent_1_street_address",
                      "config": true,
                      "mandatory": false,
                      "type": {
                        "ident": "string",
                        "path": "",
                        "minLength": 0,
                        "maxLength": 2147483647
                      }
                    }
                  },
                  {
                    "ident": "parent_1_city_state_zip",
                    "leaf": {
                      "ident": "parent_1_city_state_zip",
                      "config": true,
                      "mandatory": false,
                      "type": {
                        "ident": "string",
                        "path": "",
                        "minLength": 0,
                        "maxLength": 2147483647
                      }
                    }
                  },
                  {
                    "ident": "parent_1_home_phone",
                    "leaf": {
                      "ident": "parent_1_home_phone",
                      "config": true,
                      "mandatory": false,
                      "type": {
                        "ident": "string",
                        "path": "",
                        "minLength": 0,
                        "maxLength": 2147483647
                      }
                    }
                  },
                  {
                    "ident": "parent_1_work_phone",
                    "leaf": {
                      "ident": "parent_1_work_phone",
                      "config": true,
                      "mandatory": false,
                      "type": {
                        "ident": "string",
                        "path": "",
                        "minLength": 0,
                        "maxLength": 2147483647
                      }
                    }
                  },
                  {
                    "ident": "parent_1_cell_phone",
                    "leaf": {
                      "ident": "parent_1_cell_phone",
                      "config": true,
                      "mandatory": false,
                      "type": {
                        "ident": "string",
                        "path": "",
                        "minLength": 0,
                        "maxLength": 2147483647
                      }
                    }
                  },
                  {
                    "ident": "parent_1_email",
                    "leaf": {
                      "ident": "parent_1_email",
                      "config": true,
                      "mandatory": false,
                      "type": {
                        "ident": "string",
                        "path": "",
                        "minLength": 0,
                        "maxLength": 2147483647
                      }
                    }
                  },
                  {
                    "ident": "parent_1_receive_email",
                    "leaf": {
                      "ident": "parent_1_receive_email",
                      "config": true,
                      "mandatory": false,
                      "type": {
                        "ident": "string",
                        "path": "",
                        "minLength": 0,
                        "maxLength": 2147483647
                      }
                    }
                  },
                  {
                    "ident": "parent_1_legal_guardian",
                    "leaf": {
                      "ident": "parent_1_legal_guardian",
                      "config": true,
                      "mandatory": false,
                      "type": {
                        "ident": "string",
                        "path": "",
                        "minLength": 0,
                        "maxLength": 2147483647
                      }
                    }
                  },
                  {
                    "ident": "parent_2_name",
                    "leaf": {
                      "ident": "parent_2_name",
                      "config": true,
                      "mandatory": false,
                      "type": {
                        "ident": "string",
                        "path": "",
                        "minLength": 0,
                        "maxLength": 2147483647
                      }
                    }
                  },
                  {
                    "ident": "parent_2_emergency_priority",
                    "leaf": {
                      "ident": "parent_2_emergency_priority",
                      "config": true,
                      "mandatory": false,
                      "type": {
                        "ident": "string",
                        "path": "",
                        "minLength": 0,
                        "maxLength": 2147483647
                      }
                    }
                  },
                  {
                    "ident": "parent_2_relationship",
                    "leaf": {
                      "ident": "parent_2_relationship",
                      "config": true,
                      "mandatory": false,
                      "type": {
                        "ident": "string",
                        "path": "",
                        "minLength": 0,
                        "maxLength": 2147483647
                      }
                    }
                  },
                  {
                    "ident": "parent_2_street_address",
                    "leaf": {
                      "ident": "parent_2_street_address",
                      "config": true,
                      "mandatory": false,
                      "type": {
                        "ident": "string",
                        "path": "",
                        "minLength": 0,
                        "maxLength": 2147483647
                      }
                    }
                  },
                  {
                    "ident": "parent_2_city_state_zip",
                    "leaf": {
                      "ident": "parent_2_city_state_zip",
                      "config": true,
                      "mandatory": false,
                      "type": {
                        "ident": "string",
                        "path": "",
                        "minLength": 0,
                        "maxLength": 2147483647
                      }
                    }
                  },
                  {
                    "ident": "parent_2_home_phone",
                    "leaf": {
                      "ident": "parent_2_home_phone",
                      "config": true,
                      "mandatory": false,
                      "type": {
                        "ident": "string",
                        "path": "",
                        "minLength": 0,
                        "maxLength": 2147483647
                      }
                    }
                  },
                  {
                    "ident": "parent_2_work_phone",
                    "leaf": {
                      "ident": "parent_2_work_phone",
                      "config": true,
                      "mandatory": false,
                      "type": {
                        "ident": "string",
                        "path": "",
                        "minLength": 0,
                        "maxLength": 2147483647
                      }
                    }
                  },
                  {
                    "ident": "parent_2_cell_phone",
                    "leaf": {
                      "ident": "parent_2_cell_phone",
                      "config": true,
                      "mandatory": false,
                      "type": {
                        "ident": "string",
                        "path": "",
                        "minLength": 0,
                        "maxLength": 2147483647
                      }
                    }
                  },
                  {
                    "ident": "parent_2_email",
                    "leaf": {
                      "ident": "parent_2_email",
                      "config": true,
                      "mandatory": false,
                      "type": {
                        "ident": "string",
                        "path": "",
                        "minLength": 0,
                        "maxLength": 2147483647
                      }
                    }
                  },
                  {
                    "ident": "parent_2_receive_email",
                    "leaf": {
                      "ident": "parent_2_receive_email",
                      "config": true,
                      "mandatory": false,
                      "type": {
                        "ident": "string",
                        "path": "",
                        "minLength": 0,
                        "maxLength": 2147483647
                      }
                    }
                  },
                  {
                    "ident": "parent_2_legal_guardian",
                    "leaf": {
                      "ident": "parent_2_legal_guardian",
                      "config": true,
                      "mandatory": false,
                      "type": {
                        "ident": "string",
                        "path": "",
                        "minLength": 0,
                        "maxLength": 2147483647
                      }
                    }
                  },
                  {
                    "ident": "parent_3_name",
                    "leaf": {
                      "ident": "parent_3_name",
                      "config": true,
                      "mandatory": false,
                      "type": {
                        "ident": "string",
                        "path": "",
                        "minLength": 0,
                        "maxLength": 2147483647
                      }
                    }
                  },
                  {
                    "ident": "parent_3_emergency_priority",
                    "leaf": {
                      "ident": "parent_3_emergency_priority",
                      "config": true,
                      "mandatory": false,
                      "type": {
                        "ident": "string",
                        "path": "",
                        "minLength": 0,
                        "maxLength": 2147483647
                      }
                    }
                  },
                  {
                    "ident": "parent_3_relationship",
                    "leaf": {
                      "ident": "parent_3_relationship",
                      "config": true,
                      "mandatory": false,
                      "type": {
                        "ident": "string",
                        "path": "",
                        "minLength": 0,
                        "maxLength": 2147483647
                      }
                    }
                  },
                  {
                    "ident": "parent_3_street_address",
                    "leaf": {
                      "ident": "parent_3_street_address",
                      "config": true,
                      "mandatory": false,
                      "type": {
                        "ident": "string",
                        "path": "",
                        "minLength": 0,
                        "maxLength": 2147483647
                      }
                    }
                  },
                  {
                    "ident": "parent_3_city_state_zip",
                    "leaf": {
                      "ident": "parent_3_city_state_zip",
                      "config": true,
                      "mandatory": false,
                      "type": {
                        "ident": "string",
                        "path": "",
                        "minLength": 0,
                        "maxLength": 2147483647
                      }
                    }
                  },
                  {
                    "ident": "parent_3_home_phone",
                    "leaf": {
                      "ident": "parent_3_home_phone",
                      "config": true,
                      "mandatory": false,
                      "type": {
                        "ident": "string",
                        "path": "",
                        "minLength": 0,
                        "maxLength": 2147483647
                      }
                    }
                  },
                  {
                    "ident": "parent_3_work_phone",
                    "leaf": {
                      "ident": "parent_3_work_phone",
                      "config": true,
                      "mandatory": false,
                      "type": {
                        "ident": "string",
                        "path": "",
                        "minLength": 0,
                        "maxLength": 2147483647
                      }
                    }
                  },
                  {
                    "ident": "parent_3_cell_phone",
                    "leaf": {
                      "ident": "parent_3_cell_phone",
                      "config": true,
                      "mandatory": false,
                      "type": {
                        "ident": "string",
                        "path": "",
                        "minLength": 0,
                        "maxLength": 2147483647
                      }
                    }
                  },
                  {
                    "ident": "parent_3_email",
                    "leaf": {
                      "ident": "parent_3_email",
                      "config": true,
                      "mandatory": false,
                      "type": {
                        "ident": "string",
                        "path": "",
                        "minLength": 0,
                        "maxLength": 2147483647
                      }
                    }
                  },
                  {
                    "ident": "parent_3_receive_email",
                    "leaf": {
                      "ident": "parent_3_receive_email",
                      "config": true,
                      "mandatory": false,
                      "type": {
                        "ident": "string",
                        "path": "",
                        "minLength": 0,
                        "maxLength": 2147483647
                      }
                    }
                  },
                  {
                    "ident": "parent_3_legal_guardian",
                    "leaf": {
                      "ident": "parent_3_legal_guardian",
                      "config": true,
                      "mandatory": false,
                      "type": {
                        "ident": "string",
                        "path": "",
                        "minLength": 0,
                        "maxLength": 2147483647
                      }
                    }
                  },
                  {
                    "ident": "parent_4_name",
                    "leaf": {
                      "ident": "parent_4_name",
                      "config": true,
                      "mandatory": false,
                      "type": {
                        "ident": "string",
                        "path": "",
                        "minLength": 0,
                        "maxLength": 2147483647
                      }
                    }
                  },
                  {
                    "ident": "parent_4_emergency_priority",
                    "leaf": {
                      "ident": "parent_4_emergency_priority",
                      "config": true,
                      "mandatory": false,
                      "type": {
                        "ident": "string",
                        "path": "",
                        "minLength": 0,
                        "maxLength": 2147483647
                      }
                    }
                  },
                  {
                    "ident": "parent_4_relationship",
                    "leaf": {
                      "ident": "parent_4_relationship",
                      "config": true,
                      "mandatory": false,
                      "type": {
                        "ident": "string",
                        "path": "",
                        "minLength": 0,
                        "maxLength": 2147483647
                      }
                    }
                  },
                  {
                    "ident": "parent_4_street_address",
                    "leaf": {
                      "ident": "parent_4_street_address",
                      "config": true,
                      "mandatory": false,
                      "type": {
                        "ident": "string",
                        "path": "",
                        "minLength": 0,
                        "maxLength": 2147483647
                      }
                    }
                  },
                  {
                    "ident": "parent_4_city_state_zip",
                    "leaf": {
                      "ident": "parent_4_city_state_zip",
                      "config": true,
                      "mandatory": false,
                      "type": {
                        "ident": "string",
                        "path": "",
                        "minLength": 0,
                        "maxLength": 2147483647
                      }
                    }
                  },
                  {
                    "ident": "parent_4_home_phone",
                    "leaf": {
                      "ident": "parent_4_home_phone",
                      "config": true,
                      "mandatory": false,
                      "type": {
                        "ident": "string",
                        "path": "",
                        "minLength": 0,
                        "maxLength": 2147483647
                      }
                    }
                  },
                  {
                    "ident": "parent_4_work_phone",
                    "leaf": {
                      "ident": "parent_4_work_phone",
                      "config": true,
                      "mandatory": false,
                      "type": {
                        "ident": "string",
                        "path": "",
                        "minLength": 0,
                        "maxLength": 2147483647
                      }
                    }
                  },
                  {
                    "ident": "parent_4_cell_phone",
                    "leaf": {
                      "ident": "parent_4_cell_phone",
                      "config": true,
                      "mandatory": false,
                      "type": {
                        "ident": "string",
                        "path": "",
                        "minLength": 0,
                        "maxLength": 2147483647
                      }
                    }
                  },
                  {
                    "ident": "parent_4_email",
                    "leaf": {
                      "ident": "parent_4_email",
                      "config": true,
                      "mandatory": false,
                      "type": {
                        "ident": "string",
                        "path": "",
                        "minLength": 0,
                        "maxLength": 2147483647
                      }
                    }
                  },
                  {
                    "ident": "parent_4_receive_email",
                    "leaf": {
                      "ident": "parent_4_receive_email",
                      "config": true,
                      "mandatory": false,
                      "type": {
                        "ident": "string",
                        "path": "",
                        "minLength": 0,
                        "maxLength": 2147483647
                      }
                    }
                  },
                  {
                    "ident": "parent_4_legal_guardian",
                    "leaf": {
                      "ident": "parent_4_legal_guardian",
                      "config": true,
                      "mandatory": false,
                      "type": {
                        "ident": "string",
                        "path": "",
                        "minLength": 0,
                        "maxLength": 2147483647
                      }
                    }
                  }
                ]
              }
            }
          ]
        }
      }
    ]
  },
  "data": {
    "upload": {
      "rows": [
        {
          "student_local_id": "000000",
          "student_sasid": "1024608925",
          "enrollment_status": "Active",
          "student_first": "Charles",
          "student_middle": "Robert",
          "student_last": "Abany",
          "address": "250 Baldwin AV",
          "city_state_zip": "Framingham MA 01701",
          "school_name": "Walsh Middle School",
          "school_code": "WAL",
          "school_id": "01000310",
          "grade_level": "06",
          "next_year_school": "N/A",
          "next_grade": "06",
          "transfer_school": "",
          "trans_Eligible": "Y",
          "trans_Eligible_NextYr": "Y",
          "student_dob": "02/20/2004",
          "student_home_language": "",
          "home_phone": "508-877-4827",
          "student_email": "CAbany4213@fpsed.org",
          "student_mkv_status": "Not applicable",
          "student_mkv_exitDate": "",
          "student_frLunchAppStatus": "Received",
          "student_freereduced": "F",
          "student_last_year_fr": "F",
          "Bus_am_route": "24",
          "Bus_am_description": "Baldwin Av \u0026 Nob Hill Dr",
          "Bus_am_time": "07:39 AM",
          "Bus_pm_route": "16",
          "Bus_pm_description": "Baldwin Av \u0026 Nob Hill Dr",
          "Bus_pm_time": "02:53 PM",
          "bus_tran_am_arrival": "",
          "bus_tran_am_location": "",
          "bus_tran_am_alt": "",
          "bus_tran_pm_arrival": "",
          "bus_tran_pm_location": "",
          "bus_tran_pm_alt": "",
          "parent_1_name": "Elizabeth Abany",
          "parent_1_emergency_priority": "1",
          "parent_1_relationship": "Mother",
          "parent_1_street_address": "250 Baldwin AV",
          "parent_1_city_state_zip": "Framingham MA 01701",
          "parent_1_home_phone": "508-309-3996",
          "parent_1_work_phone": "508-641-1977",
          "parent_1_cell_phone": "",
          "parent_1_email": "abany825@gmail.com",
          "parent_1_receive_email": "true",
          "parent_1_legal_guardian": "1",
          "parent_2_name": "Robert Abany",
          "parent_2_emergency_priority": "2",
          "parent_2_relationship": "Father",
          "parent_2_street_address": "118 Evelina ST",
          "parent_2_city_state_zip": "Marlboro MA 01752",
          "parent_2_home_phone": "",
          "parent_2_work_phone": "508-380-4752",
          "parent_2_cell_phone": "617-722-3589",
          "parent_2_email": "robabany@gmail.com",
          "parent_2_receive_email": "false",
          "parent_2_legal_guardian": "1",
          "parent_3_name": "",
          "parent_3_emergency_priority": "",
          "parent_3_relationship": "",
          "parent_3_street_address": "",
          "parent_3_city_state_zip": "",
          "parent_3_home_phone": "",
          "parent_3_work_phone": "",
          "parent_3_cell_phone": "",
          "parent_3_email": "",
          "parent_3_receive_email": "",
          "parent_3_legal_guardian": "",
          "parent_4_name": "",
          "parent_4_emergency_priority": "",
          "parent_4_relationship": "",
          "parent_4_street_address": "",
          "parent_4_city_state_zip": "",
          "parent_4_home_phone": "",
          "parent_4_work_phone": "",
          "parent_4_cell_phone": "",
          "parent_4_email": "",
          "parent_4_receive_email": "",
          "parent_4_legal_guardian": ""
        },
        {
          "student_local_id": "140756",
          "student_sasid": "1072730327",
          "enrollment_status": "Active",
          "student_first": "Isabelle",
          "student_middle": "Jeanne",
          "student_last": "Abany",
          "address": "250 Baldwin AV",
          "city_state_zip": "Framingham MA 01701",
          "school_name": "Dunning Elementary School",
          "school_code": "DUN",
          "school_id": "01000007",
          "grade_level": "04",
          "next_year_school": "N/A",
          "next_grade": "04",
          "transfer_school": "",
          "trans_Eligible": "N",
          "trans_Eligible_NextYr": "N",
          "student_dob": "12/27/2005",
          "student_home_language": "267",
          "home_phone": "508-877-4827",
          "student_email": "IAbany0756@fpsed.org",
          "student_mkv_status": "Not applicable",
          "student_mkv_exitDate": "",
          "student_frLunchAppStatus": "Received",
          "student_freereduced": "F",
          "student_last_year_fr": "F",
          "Bus_am_route": "33",
          "Bus_am_description": "Baldwin Av \u0026 Hiram Rd ***",
          "Bus_am_time": "08:33 AM",
          "Bus_pm_route": "41",
          "Bus_pm_description": "Baldwin Av \u0026 Hiram Rd ***",
          "Bus_pm_time": "03:32 PM",
          "bus_tran_am_arrival": "",
          "bus_tran_am_location": "",
          "bus_tran_am_alt": "",
          "bus_tran_pm_arrival": "",
          "bus_tran_pm_location": "",
          "bus_tran_pm_alt": "",
          "parent_1_name": "Elizabeth Abany",
          "parent_1_emergency_priority": "1",
          "parent_1_relationship": "Mother",
          "parent_1_street_address": "250 Baldwin AV",
          "parent_1_city_state_zip": "Framingham MA 01701",
          "parent_1_home_phone": "508-309-3996",
          "parent_1_work_phone": "508-641-1977",
          "parent_1_cell_phone": "",
          "parent_1_email": "abany825@gmail.com",
          "parent_1_receive_email": "false",
          "parent_1_legal_guardian": "1",
          "parent_2_name": "Robert Abany",
          "parent_2_emergency_priority": "2",
          "parent_2_relationship": "Father",
          "parent_2_street_address": "118 Evelina ST",
          "parent_2_city_state_zip": "Marlboro MA 01752",
          "parent_2_home_phone": "",
          "parent_2_work_phone": "508-380-4752",
          "parent_2_cell_phone": "617-722-3589",
          "parent_2_email": "robabany@gmail.com",
          "parent_2_receive_email": "false",
          "parent_2_legal_guardian": "1",
          "parent_3_name": "",
          "parent_3_emergency_priority": "",
          "parent_3_relationship": "",
          "parent_3_street_address": "",
          "parent_3_city_state_zip": "",
          "parent_3_home_phone": "",
          "parent_3_work_phone": "",
          "parent_3_cell_phone": "",
          "parent_3_email": "",
          "parent_3_receive_email": "",
          "parent_3_legal_guardian": "",
          "parent_4_name": "",
          "parent_4_emergency_priority": "",
          "parent_4_relationship": "",
          "parent_4_street_address": "",
          "parent_4_city_state_zip": "",
          "parent_4_home_phone": "",
          "parent_4_work_phone": "",
          "parent_4_cell_phone": "",
          "parent_4_email": "",
          "parent_4_receive_email": "",
          "parent_4_legal_guardian": ""
        }
      ]
    }
  }
}
`