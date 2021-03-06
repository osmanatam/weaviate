//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2020 SeMI Technologies B.V. All rights reserved.
//
//  CONTACT: hello@semi.technology
//

package local

import (
	"fmt"
	"runtime/debug"
	"testing"

	"github.com/graphql-go/graphql"
	"github.com/semi-technologies/weaviate/entities/models"
	"github.com/semi-technologies/weaviate/entities/schema"
	"github.com/semi-technologies/weaviate/usecases/config"
	"github.com/semi-technologies/weaviate/usecases/network/common/peers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// These tests are component tests for the local package including all its
// subpackages, such as get, getmeta, etc.. However, they only assert that the
// graphql tree can be built under certain circumstances. This helps us to
// catch errors on edge cases like empty schemas, classes with empty
// properties, empty peer lists, peers with empty schemas, etc. However, we
// don't get any guaruantuee of whether the individual queries resolve
// correctly. For those cases we have unit tests in die individual subpackages
// (i.e. get, getmeta, aggreagate, etc.).  Additionally we have (a few) e2e
// tests.

func Test_GraphQLNetworkBuild(t *testing.T) {
	tests := testCases{

		// This tests asserts that an action-only schema doesn't lead to errors.
		testCase{
			name:  "with only actions locally",
			peers: peers.Peers{},
			localSchema: schema.Schema{
				Actions: &models.Schema{
					Classes: []*models.Class{
						&models.Class{
							Class: "BestLocalAction",
							Properties: []*models.Property{
								&models.Property{
									DataType: []string{"string"},
									Name:     "myStringProp",
								},
							},
						},
					},
				},
				Things: &models.Schema{},
			},
		},

		// This tests asserts that a things-only schema doesn't lead to errors.
		testCase{
			name:  "with only things locally",
			peers: peers.Peers{},
			localSchema: schema.Schema{
				Things: &models.Schema{
					Classes: []*models.Class{
						&models.Class{
							Class: "BestLocalThing",
							Properties: []*models.Property{
								&models.Property{
									DataType: []string{"string"},
									Name:     "myStringProp",
								},
							},
						},
					},
				},
				Actions: &models.Schema{},
			},
		},

		// This tests asserts that a class without any properties doesn't lead to
		// errors.
		testCase{
			name:  "with things without properties locally",
			peers: peers.Peers{},
			localSchema: schema.Schema{
				Things: &models.Schema{
					Classes: []*models.Class{
						&models.Class{
							Class:      "BestLocalThing",
							Properties: []*models.Property{},
						},
					},
				},
				Actions: &models.Schema{},
			},
		},

		testCase{
			name:        "without any peers",
			peers:       peers.Peers{},
			localSchema: validSchema(),
		},

		testCase{
			name:        "one peer with empty schema",
			localSchema: validSchema(),
			peers: peers.Peers{
				peers.Peer{
					Name:   "SomePeer",
					Schema: schema.Schema{},
				},
			},
		},

		testCase{
			name:        "one peer with a thing schema without classes",
			localSchema: validSchema(),
			peers: peers.Peers{
				peers.Peer{
					Name: "SomePeer",
					Schema: schema.Schema{
						Things: &models.Schema{
							Classes: []*models.Class{},
						},
					},
				},
			},
		},

		// this test asserts that we don't error with property-less classes, as we
		// could otherwise end up with empty Fields which would lead to a graphQL
		// build error
		testCase{
			name:        "one peer with a thing schema without properties, but no actions",
			localSchema: validSchema(),
			peers: peers.Peers{
				peers.Peer{
					Name: "SomePeer",
					Schema: schema.Schema{
						Things: &models.Schema{
							Classes: []*models.Class{
								&models.Class{
									Class:      "BestClass",
									Properties: []*models.Property{},
								},
							},
						},
						Actions: &models.Schema{},
					},
				},
			},
		},

		// this test asserts that we don't error with half-empty schemas, as we
		// could otherwise end up with empty Fields which would lead to a graphQL
		// build error
		testCase{
			name:        "one peer with a thing schema with properties, but no actions",
			localSchema: validSchema(),
			peers: peers.Peers{
				peers.Peer{
					Name: "SomePeer",
					Schema: schema.Schema{
						Things: &models.Schema{
							Classes: []*models.Class{
								&models.Class{
									Class: "BestClass",
									Properties: []*models.Property{
										&models.Property{
											Name:     "stringProp",
											DataType: []string{"string"},
										},
									},
								},
							},
						},
						Actions: &models.Schema{},
					},
				},
			},
		},

		testCase{
			name:        "one peer with a both a thing and an action class",
			localSchema: validSchema(),
			peers: peers.Peers{
				peers.Peer{
					Name: "SomePeer",
					Schema: schema.Schema{
						Things: &models.Schema{
							Classes: []*models.Class{
								&models.Class{
									Class:      "BestThing",
									Properties: []*models.Property{},
								},
							},
						},
						Actions: &models.Schema{
							Classes: []*models.Class{
								&models.Class{
									Class:      "BestAction",
									Properties: []*models.Property{},
								},
							},
						},
					},
				},
			},
		},

		// The properties, albeit on different classes have the same names. This
		// test asserts that there is no naming collision, i.e. that the properties
		// are namespaced correctly by their respective classes.
		testCase{
			name:        "one peer with a both a thing and an action class with properties",
			localSchema: validSchema(),
			peers: peers.Peers{
				peers.Peer{
					Name: "SomePeer",
					Schema: schema.Schema{
						Things: &models.Schema{
							Classes: []*models.Class{
								&models.Class{
									Class: "BestThing",
									Properties: []*models.Property{
										&models.Property{
											DataType: []string{"string"},
											Name:     "myStringProp",
										},
									},
								},
							},
						},
						Actions: &models.Schema{
							Classes: []*models.Class{
								&models.Class{
									Class: "BestAction",
									Properties: []*models.Property{
										&models.Property{
											DataType: []string{"string"},
											Name:     "myStringProp",
										},
									},
								},
							},
						},
					},
				},
			},
		},

		// This tests assert that there are no name collisions with multiple peers,
		// i.e. that every peer has their objects and fields namespaced correctly.
		testCase{
			name:        "two peers with identical schemas",
			localSchema: validSchema(),
			peers: peers.Peers{
				peers.Peer{
					Name: "SomePeer",
					Schema: schema.Schema{
						Things: &models.Schema{
							Classes: []*models.Class{
								&models.Class{
									Class: "BestThing",
									Properties: []*models.Property{
										&models.Property{
											DataType: []string{"string"},
											Name:     "myStringProp",
										},
									},
								},
							},
						},
						Actions: &models.Schema{
							Classes: []*models.Class{
								&models.Class{
									Class: "BestAction",
									Properties: []*models.Property{
										&models.Property{
											DataType: []string{"string"},
											Name:     "myStringProp",
										},
									},
								},
							},
						},
					},
				},
				peers.Peer{
					Name: "SomeOtherPeer",
					Schema: schema.Schema{
						Things: &models.Schema{
							Classes: []*models.Class{
								&models.Class{
									Class: "BestThing",
									Properties: []*models.Property{
										&models.Property{
											DataType: []string{"string"},
											Name:     "myStringProp",
										},
									},
								},
							},
						},
						Actions: &models.Schema{
							Classes: []*models.Class{
								&models.Class{
									Class: "BestAction",
									Properties: []*models.Property{
										&models.Property{
											DataType: []string{"string"},
											Name:     "myStringProp",
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	tests.AssertNoError(t)
}

type testCase struct {
	name        string
	peers       peers.Peers
	localSchema schema.Schema
}

type testCases []testCase

func (tests testCases) AssertNoError(t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			localSchema, err := Build(&test.localSchema, test.peers, nil, config.Config{})
			require.Nil(t, err, test.name)

			schemaObject := graphql.ObjectConfig{
				Name:        "WeaviateObj",
				Description: "Location of the root query",
				Fields:      localSchema,
			}

			func() {
				defer func() {
					if r := recover(); r != nil {
						err = fmt.Errorf("%v at %s", r, debug.Stack())
					}
				}()

				_, err = graphql.NewSchema(graphql.SchemaConfig{
					Query: graphql.NewObject(schemaObject),
				})
			}()

			assert.Nil(t, err, test.name)
		})
	}
}

func validSchema() schema.Schema {
	return schema.Schema{
		Things: &models.Schema{
			Classes: []*models.Class{
				&models.Class{
					Class: "BestLocalThing",
					Properties: []*models.Property{
						&models.Property{
							DataType: []string{"string"},
							Name:     "myStringProp",
						},
					},
				},
			},
		},
		Actions: &models.Schema{
			Classes: []*models.Class{
				&models.Class{
					Class: "BestLocalAction",
					Properties: []*models.Property{
						&models.Property{
							DataType: []string{"string"},
							Name:     "myStringProp",
						},
					},
				},
			},
		},
	}
}
