package yang

import (
	"github.com/c2g/meta"
	"strings"
	"fmt"
)

type YangPool map[string]string

var internal = make(YangPool)

func (self YangPool) OpenStream(streamId string) (meta.DataStream, error) {
	if s, found := self[streamId]; found {
		return strings.NewReader(s), nil
	}
	return nil, nil
}

func InternalModule(name string) *meta.Module {
	// TODO: performance - return deep copy of cached copy
	inlineYang, err := LoadModule(InternalYang(), name)
	if err != nil {
		msg := fmt.Sprintf("Error parsing %s yang, %s", name, err.Error())
		panic(msg)
	}
	return inlineYang
}

func InternalYang() YangPool {
	return internal
}

func init() {
	InternalYang()["yanglib"] = `
module yanglib {
    namespace "http://schema.org/yang";
    prefix "schema";
    description "Yang definition of yang";
    revision 0 {
        description "Yang 1.0 with some 1.1 features";
    }

    grouping def-header {
        leaf ident {
            type string;
        }
        leaf description {
            type string;
        }
    }

    grouping type {
        container type {
            leaf ident {
                type string;
            }
            leaf range {
                type string;
            }
            leaf-list enumeration {
                type string;
            }
            leaf path {
                type string;
            }
            leaf minLength {
                type int32;
            }
            leaf maxLength {
                type int32;
            }
        }
    }

    grouping groupings-typedefs {
        list groupings {
            key "ident";
            uses def-header;

            /*
              !! CIRCULAR
            */
            uses groupings-typedefs;
            uses containers-lists-leafs-uses-choice;
        }
        list typedefs {
            key "ident";
            uses def-header;
            uses type;
        }
    }

    grouping actions {
        list actions {
            key "ident";
            uses def-header;
            container input {
                uses groupings-typedefs;
                uses containers-lists-leafs-uses-choice;
            }
            container output {
                uses groupings-typedefs;
                uses containers-lists-leafs-uses-choice;
            }
        }
    }

    grouping notifications {
        list notifications {
            key "ident";
            uses def-header;
            uses groupings-typedefs;
            uses containers-lists-leafs-uses-choice;
        }
    }

    grouping containers-lists-leafs-uses-choice {
        list definitions {
            key "ident";
            leaf ident {
            	type string;
            }
            choice body-stmt {
                case container {
                    container container {
                        uses def-header;
                        uses groupings-typedefs;
                        uses containers-lists-leafs-uses-choice;
                        /*uses notifications; */
                    }
                }
                case list {
                    container list {
                        leaf-list key {
                            type string;
                        }
                        uses def-header;
                        uses groupings-typedefs;
                        uses containers-lists-leafs-uses-choice;
                        /* uses notifications; */
                    }
                }
                case leaf {
                    container leaf {
                        uses def-header;
                        leaf config {
                            type boolean;
                        }
                        leaf mandatory {
                            type boolean;
                        }
                        uses type;
                    }
                }
                case anyxml {
                    container anyxml {
                        uses def-header;
                        leaf config {
                            type boolean;
                        }
                        leaf mandatory {
                            type boolean;
                        }
                        uses type;
                    }
                }
                case leaf-list {
                    container leaf-list {
                        uses def-header;
                        leaf config {
                            type boolean;
                        }
                        leaf mandatory {
                            type boolean;
                        }
                        uses type;
                    }
                }
                case uses {
                    container uses {
                        uses def-header;
                        /* need to expand this to use refine */
                    }
                }
                case choice {
                    container choice {
                        uses def-header;
                        list cases {
                            key "ident";
                            leaf ident {
                                type string;
                            }
                            /*
                             !! CIRCULAR
                            */
                            uses containers-lists-leafs-uses-choice;
                        }
                    }
                }
            }
        }
    }

    grouping module {
    	container module {
			uses def-header;
			leaf namespace {
				type string;
			}
			leaf prefix {
				type string;
			}
			container revision {
				leaf rev-date {
					type string;
				}
				leaf description {
					type string;
				}
			}
			list rpcs {
				key "ident";
				uses def-header;
				container input {
					uses groupings-typedefs;
					uses containers-lists-leafs-uses-choice;
				}
				container output {
					uses groupings-typedefs;
					uses containers-lists-leafs-uses-choice;
				}
			}
			uses notifications;
			uses groupings-typedefs;
			uses containers-lists-leafs-uses-choice;
		}
	}
}
`

	InternalYang()["turing-machine"] = `
module turing-machine {

  namespace "http://example.net/turing-machine";

  prefix "tm";

  description
    "Data model for the Turing Machine.";

  revision 2013-12-27 {
    description
      "Initial revision.";
  }

  /* Typedefs */

  typedef tape-symbol {
    description
      "Type of symbols appearing in tape cells.

       A blank is represented as an empty string where necessary.";
    type string {
      length "0..1";
    }
  }

  typedef cell-index {
    description
      "Type for indexing tape cells.";
    type int64;
  }

  typedef state-index {
    description
      "Type for indexing states of the control unit.";
    type uint16;
  }

  typedef head-dir {
    type enumeration {
      enum left;
      enum right;
    }
    default "right";
    description
      "Possible directions for moving the read/write head, one cell
       to the left or right (default).";
  }

  /* Groupings */

  grouping tape-cells {
    description
      "The tape of the Turing Machine is represented as a sparse
       array.";
    list cell {
      description
        "List of non-blank cells.";
      key "coord";
      leaf coord {
        type cell-index;
        description
          "Coordinate (index) of the tape cell.";
      }
      leaf symbol {
        type tape-symbol {
          length "1";
        }
        description
          "Symbol appearing in the tape cell.

           Blank (empty string) is not allowed here because the
           'cell' list only contains non-blank cells.";
      }
    }
  }

  /* State data and Configuration */

  container turing-machine {
    description
      "State data and configuration of a Turing Machine.";
    leaf state {
      config "false";
      mandatory "true";
      type state-index;
      description
        "Current state of the control unit.

         The initial state is 0.";
    }
    leaf head-position {
      config "false";
      mandatory "true";
      type cell-index;
      description
        "Position of tape read/write head.";
    }
    container tape {
      description
        "The contents of the tape.";
      config "false";
      uses tape-cells;
      action rewind {
        description "be kind";
        input {
          leaf position {
            type int32;
          }
        }
        output {
          leaf estimatedTime {
            type int32;
          }
        }
      }
    }
    container transition-function {
      description
        "The Turing Machine is configured by specifying the
         transition function.";
      list delta {
        description
          "The list of transition rules.";
        key "label";
        unique "input/state input/symbol";
        leaf label {
          type string;
          description
            "An arbitrary label of the transition rule.";
        }
        container input {
          description
            "Output values of the transition rule.";
          leaf state {
            type state-index;
            description
              "New state of the control unit. If this leaf is not
               present, the state doesn't change.";
          }
          leaf symbol {
            type tape-symbol;
            description
              "Symbol to be written to the tape cell. If this leaf is
               not present, the symbol doesn't change.";
          }
          leaf head-move {
            type head-dir;
            description
              "Move the head one cell to the left or right";
          }
        }
      }
    }
  }

  /* RPCs */

  rpc initialize {
    description
      "Initialize the Turing Machine as follows:

       1. Put the control unit into the initial state (0).

       2. Move the read/write head to the tape cell with coordinate
          zero.

       3. Write the string from the 'tape-content' input parameter to
          the tape, character by character, starting at cell 0. The
          tape is othewise empty.";
    input {
      leaf tape-content {
        type string;
        default "";
        description
          "The string with which the tape shall be initialized. The
           leftmost symbol will be at tape coordinate 0.";
      }
    }
  }

  rpc run {
    description
      "Start the Turing Machine operation.";
  }

  /* Notifications */

  notification halted {
    description
      "The Turing Machine has halted. This means that there is no
       transition rule for the current state and tape symbol.";
    leaf state {
      mandatory "true";
      type state-index;
    }
  }
}
`
}