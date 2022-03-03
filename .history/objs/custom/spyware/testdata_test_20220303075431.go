package spyware

import (
	"github.com/willhu-commit/pango/version"
)

type tc struct {
	desc    string
	version version.Number
	conf    Entry
}

func getTests() []tc {
	return []tc{
		{"v1 standard less than no qualifiers", version.Number{7, 0, 0, ""}, Entry{
			Name:          "t1",
			ThreatName:    "threat1",
			Comment:       "foobar",
			Severity:      "low",
			Direction:     "client2server",
			DefaultAction: DefaultActionAllow,
			Cves:          []string{"cve1", "cve2"},
			Bugtraqs:      []string{"bug1", "bug2"},
			Vendors:       []string{"vendor1", "vendor2"},
			References:    []string{"ref1", "ref2"},
			StandardSignatureType: &StandardSignatureType{
				Signatures: []StandardSignature{
					StandardSignature{
						Name:      "sig1",
						Comment:   "sig comment",
						Scope:     "transaction",
						OrderFree: true,
						StandardAnds: []StandardAnd{
							StandardAnd{
								Name: "And 1",
								StandardOrs: []StandardOr{
									StandardOr{
										Name: "Or 1",
										LessThan: &Condition{
											Value:   4096,
											Context: "http-rsp-code",
										},
									},
								},
							},
						},
					},
				},
			},
		}},
		{"v1 standard less than with qualifiers", version.Number{7, 0, 0, ""}, Entry{
			Name:          "t2",
			ThreatName:    "threat1",
			Comment:       "foobar",
			Severity:      "low",
			Direction:     "client2server",
			DefaultAction: DefaultActionAlert,
			StandardSignatureType: &StandardSignatureType{
				Signatures: []StandardSignature{
					StandardSignature{
						Name:      "sig1",
						Comment:   "sig comment",
						Scope:     "transaction",
						OrderFree: true,
						StandardAnds: []StandardAnd{
							StandardAnd{
								Name: "And 1",
								StandardOrs: []StandardOr{
									StandardOr{
										Name: "Or 1",
										LessThan: &Condition{
											Value:   4096,
											Context: "http-rsp-code",
											Qualifiers: []Qualifier{
												Qualifier{
													Qualifier: "foo",
													Value:     "bar",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		}},
		{"v1 standard equal to no qualifiers", version.Number{7, 0, 0, ""}, Entry{
			Name:          "t3",
			ThreatName:    "threat1",
			Comment:       "foobar",
			Severity:      "low",
			Direction:     "client2server",
			DefaultAction: DefaultActionDrop,
			StandardSignatureType: &StandardSignatureType{
				Signatures: []StandardSignature{
					StandardSignature{
						Name:      "sig1",
						Comment:   "sig comment",
						Scope:     "transaction",
						OrderFree: true,
						StandardAnds: []StandardAnd{
							StandardAnd{
								Name: "And 1",
								StandardOrs: []StandardOr{
									StandardOr{
										Name: "Or 1",
										EqualTo: &EqualTo{
											Value:   4096,
											Context: "http-rsp-code",
										},
									},
								},
							},
						},
					},
				},
			},
		}},
		{"v1 standard equal to with qualifiers", version.Number{7, 0, 0, ""}, Entry{
			Name:          "t4",
			ThreatName:    "threat1",
			Comment:       "foobar",
			Severity:      "low",
			Direction:     "client2server",
			DefaultAction: DefaultActionResetClient,
			StandardSignatureType: &StandardSignatureType{
				Signatures: []StandardSignature{
					StandardSignature{
						Name:      "sig1",
						Comment:   "sig comment",
						Scope:     "transaction",
						OrderFree: true,
						StandardAnds: []StandardAnd{
							StandardAnd{
								Name: "And 1",
								StandardOrs: []StandardOr{
									StandardOr{
										Name: "Or 1",
										EqualTo: &EqualTo{
											Value:   4096,
											Context: "http-rsp-code",
											Qualifiers: []Qualifier{
												Qualifier{
													Qualifier: "foo",
													Value:     "bar",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		}},
		{"v1 standard greater than no qualifiers", version.Number{7, 0, 0, ""}, Entry{
			Name:          "t5",
			ThreatName:    "threat1",
			Comment:       "foobar",
			Severity:      "low",
			Direction:     "client2server",
			DefaultAction: DefaultActionResetServer,
			StandardSignatureType: &StandardSignatureType{
				Signatures: []StandardSignature{
					StandardSignature{
						Name:      "sig1",
						Comment:   "sig comment",
						Scope:     "transaction",
						OrderFree: true,
						StandardAnds: []StandardAnd{
							StandardAnd{
								Name: "And 1",
								StandardOrs: []StandardOr{
									StandardOr{
										Name: "Or 1",
										GreaterThan: &Condition{
											Value:   4096,
											Context: "http-rsp-code",
										},
									},
								},
							},
						},
					},
				},
			},
		}},
		{"v1 standard greater than with qualifiers", version.Number{7, 0, 0, ""}, Entry{
			Name:          "t6",
			ThreatName:    "threat1",
			Comment:       "foobar",
			Severity:      "low",
			Direction:     "client2server",
			DefaultAction: DefaultActionResetBoth,
			StandardSignatureType: &StandardSignatureType{
				Signatures: []StandardSignature{
					StandardSignature{
						Name:      "sig1",
						Comment:   "sig comment",
						Scope:     "transaction",
						OrderFree: true,
						StandardAnds: []StandardAnd{
							StandardAnd{
								Name: "And 1",
								StandardOrs: []StandardOr{
									StandardOr{
										Name: "Or 1",
										GreaterThan: &Condition{
											Value:   4096,
											Context: "http-rsp-code",
											Qualifiers: []Qualifier{
												Qualifier{
													Qualifier: "foo",
													Value:     "bar",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		}},
		{"v1 standard pattern match no qualifiers", version.Number{7, 0, 0, ""}, Entry{
			Name:            "t7",
			ThreatName:      "threat1",
			Comment:         "foobar",
			Severity:        "low",
			Direction:       "client2server",
			DefaultAction:   DefaultActionBlockIp,
			BlockIpTrackBy:  "source-and-destination",
			BlockIpDuration: 1024,
			StandardSignatureType: &StandardSignatureType{
				Signatures: []StandardSignature{
					StandardSignature{
						Name:      "sig1",
						Comment:   "sig comment",
						Scope:     "transaction",
						OrderFree: true,
						StandardAnds: []StandardAnd{
							StandardAnd{
								Name: "And 1",
								StandardOrs: []StandardOr{
									StandardOr{
										Name: "Or 1",
										PatternMatch: &Pattern{
											Pattern: "abcdefgh",
											Context: "http-rsp-code",
											Negate:  true,
										},
									},
								},
							},
						},
					},
				},
			},
		}},
		{"v1 standard pattern match with qualifiers", version.Number{7, 0, 0, ""}, Entry{
			Name:          "t8",
			ThreatName:    "threat1",
			Comment:       "foobar",
			Severity:      "low",
			Direction:     "client2server",
			DefaultAction: DefaultActionAllow,
			StandardSignatureType: &StandardSignatureType{
				Signatures: []StandardSignature{
					StandardSignature{
						Name:      "sig1",
						Comment:   "sig comment",
						Scope:     "transaction",
						OrderFree: true,
						StandardAnds: []StandardAnd{
							StandardAnd{
								Name: "And 1",
								StandardOrs: []StandardOr{
									StandardOr{
										Name: "Or 1",
										PatternMatch: &Pattern{
											Pattern: "abcdefgh",
											Context: "http-rsp-code",
											Negate:  false,
											Qualifiers: []Qualifier{
												Qualifier{
													Qualifier: "foo",
													Value:     "bar",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		}},
		{"v1 combination", version.Number{7, 0, 0, ""}, Entry{
			Name:          "t9",
			ThreatName:    "myThreat",
			Comment:       "foobar",
			Severity:      "medium",
			Direction:     "client2server",
			DefaultAction: DefaultActionAlert,
			CombinationSignatureType: &CombinationSignatureType{
				OrderFree:           true,
				ThresholdTime:       7,
				IntervalTime:        8,
				AggregationCriteria: "source-and-destination",
				Signatures: []CombinationSignature{
					CombinationSignature{
						Name: "mySig",
						CombinationOrs: []CombinationOr{
							CombinationOr{
								Name:     "Or Condition 1",
								ThreatId: "10052",
							},
						},
					},
				},
			},
		}},
		{"v1 standard less than no qualifiers", version.Number{7, 0, 0, ""}, Entry{
			Name:          "t10",
			ThreatName:    "threat1",
			Comment:       "foobar",
			Severity:      "low",
			Direction:     "client2server",
			DefaultAction: DefaultActionAllow,
			Cves:          []string{"cve1", "cve2"},
			Bugtraqs:      []string{"bug1", "bug2"},
			Vendors:       []string{"vendor1", "vendor2"},
			References:    []string{"ref1", "ref2"},
			StandardSignatureType: &StandardSignatureType{
				Signatures: []StandardSignature{
					StandardSignature{
						Name:      "sig1",
						Comment:   "sig comment",
						Scope:     "transaction",
						OrderFree: true,
						StandardAnds: []StandardAnd{
							StandardAnd{
								Name: "And 1",
								StandardOrs: []StandardOr{
									StandardOr{
										Name: "Or 1",
										LessThan: &Condition{
											Value:   4096,
											Context: "http-rsp-code",
										},
									},
								},
							},
						},
					},
				},
			},
		}},
		{"v1 standard less than with qualifiers", version.Number{7, 0, 0, ""}, Entry{
			Name:          "t11",
			ThreatName:    "threat1",
			Comment:       "foobar",
			Severity:      "low",
			Direction:     "client2server",
			DefaultAction: DefaultActionAlert,
			StandardSignatureType: &StandardSignatureType{
				Signatures: []StandardSignature{
					StandardSignature{
						Name:      "sig1",
						Comment:   "sig comment",
						Scope:     "transaction",
						OrderFree: true,
						StandardAnds: []StandardAnd{
							StandardAnd{
								Name: "And 1",
								StandardOrs: []StandardOr{
									StandardOr{
										Name: "Or 1",
										LessThan: &Condition{
											Value:   4096,
											Context: "http-rsp-code",
											Qualifiers: []Qualifier{
												Qualifier{
													Qualifier: "foo",
													Value:     "bar",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		}},
		{"v1 standard equal to no qualifiers", version.Number{7, 0, 0, ""}, Entry{
			Name:          "t12",
			ThreatName:    "threat1",
			Comment:       "foobar",
			Severity:      "low",
			Direction:     "client2server",
			DefaultAction: DefaultActionDrop,
			StandardSignatureType: &StandardSignatureType{
				Signatures: []StandardSignature{
					StandardSignature{
						Name:      "sig1",
						Comment:   "sig comment",
						Scope:     "transaction",
						OrderFree: true,
						StandardAnds: []StandardAnd{
							StandardAnd{
								Name: "And 1",
								StandardOrs: []StandardOr{
									StandardOr{
										Name: "Or 1",
										EqualTo: &EqualTo{
											Value:   4096,
											Context: "http-rsp-code",
										},
									},
								},
							},
						},
					},
				},
			},
		}},
		{"v1 standard equal to with qualifiers", version.Number{7, 0, 0, ""}, Entry{
			Name:          "t13",
			ThreatName:    "threat1",
			Comment:       "foobar",
			Severity:      "low",
			Direction:     "client2server",
			DefaultAction: DefaultActionResetClient,
			StandardSignatureType: &StandardSignatureType{
				Signatures: []StandardSignature{
					StandardSignature{
						Name:      "sig1",
						Comment:   "sig comment",
						Scope:     "transaction",
						OrderFree: true,
						StandardAnds: []StandardAnd{
							StandardAnd{
								Name: "And 1",
								StandardOrs: []StandardOr{
									StandardOr{
										Name: "Or 1",
										EqualTo: &EqualTo{
											Value:   4096,
											Context: "http-rsp-code",
											Qualifiers: []Qualifier{
												Qualifier{
													Qualifier: "foo",
													Value:     "bar",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		}},
		{"v1 standard greater than no qualifiers", version.Number{7, 0, 0, ""}, Entry{
			Name:          "t14",
			ThreatName:    "threat1",
			Comment:       "foobar",
			Severity:      "low",
			Direction:     "client2server",
			DefaultAction: DefaultActionResetServer,
			StandardSignatureType: &StandardSignatureType{
				Signatures: []StandardSignature{
					StandardSignature{
						Name:      "sig1",
						Comment:   "sig comment",
						Scope:     "transaction",
						OrderFree: true,
						StandardAnds: []StandardAnd{
							StandardAnd{
								Name: "And 1",
								StandardOrs: []StandardOr{
									StandardOr{
										Name: "Or 1",
										GreaterThan: &Condition{
											Value:   4096,
											Context: "http-rsp-code",
										},
									},
								},
							},
						},
					},
				},
			},
		}},
		{"v1 standard greater than with qualifiers", version.Number{7, 0, 0, ""}, Entry{
			Name:          "t15",
			ThreatName:    "threat1",
			Comment:       "foobar",
			Severity:      "low",
			Direction:     "client2server",
			DefaultAction: DefaultActionResetBoth,
			StandardSignatureType: &StandardSignatureType{
				Signatures: []StandardSignature{
					StandardSignature{
						Name:      "sig1",
						Comment:   "sig comment",
						Scope:     "transaction",
						OrderFree: true,
						StandardAnds: []StandardAnd{
							StandardAnd{
								Name: "And 1",
								StandardOrs: []StandardOr{
									StandardOr{
										Name: "Or 1",
										GreaterThan: &Condition{
											Value:   4096,
											Context: "http-rsp-code",
											Qualifiers: []Qualifier{
												Qualifier{
													Qualifier: "foo",
													Value:     "bar",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		}},
		{"v1 standard pattern match no qualifiers", version.Number{7, 0, 0, ""}, Entry{
			Name:            "t16",
			ThreatName:      "threat1",
			Comment:         "foobar",
			Severity:        "low",
			Direction:       "client2server",
			DefaultAction:   DefaultActionBlockIp,
			BlockIpTrackBy:  "source-and-destination",
			BlockIpDuration: 1024,
			StandardSignatureType: &StandardSignatureType{
				Signatures: []StandardSignature{
					StandardSignature{
						Name:      "sig1",
						Comment:   "sig comment",
						Scope:     "transaction",
						OrderFree: true,
						StandardAnds: []StandardAnd{
							StandardAnd{
								Name: "And 1",
								StandardOrs: []StandardOr{
									StandardOr{
										Name: "Or 1",
										PatternMatch: &Pattern{
											Pattern: "abcdefgh",
											Context: "http-rsp-code",
											Negate:  true,
										},
									},
								},
							},
						},
					},
				},
			},
		}},
		{"v1 standard pattern match with qualifiers", version.Number{7, 0, 0, ""}, Entry{
			Name:          "t17",
			ThreatName:    "threat1",
			Comment:       "foobar",
			Severity:      "low",
			Direction:     "client2server",
			DefaultAction: DefaultActionAllow,
			StandardSignatureType: &StandardSignatureType{
				Signatures: []StandardSignature{
					StandardSignature{
						Name:      "sig1",
						Comment:   "sig comment",
						Scope:     "transaction",
						OrderFree: true,
						StandardAnds: []StandardAnd{
							StandardAnd{
								Name: "And 1",
								StandardOrs: []StandardOr{
									StandardOr{
										Name: "Or 1",
										PatternMatch: &Pattern{
											Pattern: "abcdefgh",
											Context: "http-rsp-code",
											Negate:  false,
											Qualifiers: []Qualifier{
												Qualifier{
													Qualifier: "foo",
													Value:     "bar",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		}},
		{"v1 combination order free", version.Number{7, 0, 0, ""}, Entry{
			Name:          "t18",
			ThreatName:    "myThreat",
			Comment:       "foobar",
			Severity:      "medium",
			Direction:     "client2server",
			DefaultAction: DefaultActionAlert,
			CombinationSignatureType: &CombinationSignatureType{
				OrderFree:           true,
				ThresholdTime:       7,
				IntervalTime:        8,
				AggregationCriteria: "source-and-destination",
				Signatures: []CombinationSignature{
					CombinationSignature{
						Name: "And Condition 1",
						CombinationOrs: []CombinationOr{
							CombinationOr{
								Name:     "Or Condition 1",
								ThreatId: "10052",
							},
						},
					},
				},
			},
		}},
		{"v1 combination ordered", version.Number{7, 0, 0, ""}, Entry{
			Name:          "t19",
			ThreatName:    "myThreat",
			Comment:       "foobar",
			Severity:      "medium",
			Direction:     "client2server",
			DefaultAction: DefaultActionAlert,
			CombinationSignatureType: &CombinationSignatureType{
				ThresholdTime:       7,
				IntervalTime:        8,
				AggregationCriteria: "source-and-destination",
				Signatures: []CombinationSignature{
					CombinationSignature{
						Name: "And Condition 1",
						CombinationOrs: []CombinationOr{
							CombinationOr{
								Name:     "Or Condition 1",
								ThreatId: "10052",
							},
						},
					},
				},
			},
		}},
		{"v2 standard less than no qualifiers", version.Number{10, 0, 0, ""}, Entry{
			Name:          "t1",
			ThreatName:    "threat1",
			Comment:       "foobar",
			Severity:      "low",
			Direction:     "client2server",
			DefaultAction: DefaultActionAllow,
			Cves:          []string{"cve1", "cve2"},
			Bugtraqs:      []string{"bug1", "bug2"},
			Vendors:       []string{"vendor1", "vendor2"},
			References:    []string{"ref1", "ref2"},
			StandardSignatureType: &StandardSignatureType{
				Signatures: []StandardSignature{
					StandardSignature{
						Name:      "sig1",
						Comment:   "sig comment",
						Scope:     "transaction",
						OrderFree: true,
						StandardAnds: []StandardAnd{
							StandardAnd{
								Name: "And 1",
								StandardOrs: []StandardOr{
									StandardOr{
										Name: "Or 1",
										LessThan: &Condition{
											Value:   4096,
											Context: "http-rsp-code",
										},
									},
								},
							},
						},
					},
				},
			},
		}},
		{"v2 standard less than with qualifiers", version.Number{10, 0, 0, ""}, Entry{
			Name:          "t2",
			ThreatName:    "threat1",
			Comment:       "foobar",
			Severity:      "low",
			Direction:     "client2server",
			DefaultAction: DefaultActionAlert,
			StandardSignatureType: &StandardSignatureType{
				Signatures: []StandardSignature{
					StandardSignature{
						Name:      "sig1",
						Comment:   "sig comment",
						Scope:     "transaction",
						OrderFree: true,
						StandardAnds: []StandardAnd{
							StandardAnd{
								Name: "And 1",
								StandardOrs: []StandardOr{
									StandardOr{
										Name: "Or 1",
										LessThan: &Condition{
											Value:   4096,
											Context: "http-rsp-code",
											Qualifiers: []Qualifier{
												Qualifier{
													Qualifier: "foo",
													Value:     "bar",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		}},
		{"v2 standard equal to no qualifiers", version.Number{10, 0, 0, ""}, Entry{
			Name:          "t3",
			ThreatName:    "threat1",
			Comment:       "foobar",
			Severity:      "low",
			Direction:     "client2server",
			DefaultAction: DefaultActionDrop,
			StandardSignatureType: &StandardSignatureType{
				Signatures: []StandardSignature{
					StandardSignature{
						Name:      "sig1",
						Comment:   "sig comment",
						Scope:     "transaction",
						OrderFree: true,
						StandardAnds: []StandardAnd{
							StandardAnd{
								Name: "And 1",
								StandardOrs: []StandardOr{
									StandardOr{
										Name: "Or 1",
										EqualTo: &EqualTo{
											Value:   4096,
											Context: "http-rsp-code",
										},
									},
								},
							},
						},
					},
				},
			},
		}},
		{"v2 standard equal to with qualifiers", version.Number{10, 0, 0, ""}, Entry{
			Name:          "t4",
			ThreatName:    "threat1",
			Comment:       "foobar",
			Severity:      "low",
			Direction:     "client2server",
			DefaultAction: DefaultActionResetClient,
			StandardSignatureType: &StandardSignatureType{
				Signatures: []StandardSignature{
					StandardSignature{
						Name:      "sig1",
						Comment:   "sig comment",
						Scope:     "transaction",
						OrderFree: true,
						StandardAnds: []StandardAnd{
							StandardAnd{
								Name: "And 1",
								StandardOrs: []StandardOr{
									StandardOr{
										Name: "Or 1",
										EqualTo: &EqualTo{
											Value:   4096,
											Context: "http-rsp-code",
											Qualifiers: []Qualifier{
												Qualifier{
													Qualifier: "foo",
													Value:     "bar",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		}},
		{"v2 standard greater than no qualifiers", version.Number{10, 0, 0, ""}, Entry{
			Name:          "t5",
			ThreatName:    "threat1",
			Comment:       "foobar",
			Severity:      "low",
			Direction:     "client2server",
			DefaultAction: DefaultActionResetServer,
			StandardSignatureType: &StandardSignatureType{
				Signatures: []StandardSignature{
					StandardSignature{
						Name:      "sig1",
						Comment:   "sig comment",
						Scope:     "transaction",
						OrderFree: true,
						StandardAnds: []StandardAnd{
							StandardAnd{
								Name: "And 1",
								StandardOrs: []StandardOr{
									StandardOr{
										Name: "Or 1",
										GreaterThan: &Condition{
											Value:   4096,
											Context: "http-rsp-code",
										},
									},
								},
							},
						},
					},
				},
			},
		}},
		{"v2 standard greater than with qualifiers", version.Number{10, 0, 0, ""}, Entry{
			Name:          "t6",
			ThreatName:    "threat1",
			Comment:       "foobar",
			Severity:      "low",
			Direction:     "client2server",
			DefaultAction: DefaultActionResetBoth,
			StandardSignatureType: &StandardSignatureType{
				Signatures: []StandardSignature{
					StandardSignature{
						Name:      "sig1",
						Comment:   "sig comment",
						Scope:     "transaction",
						OrderFree: true,
						StandardAnds: []StandardAnd{
							StandardAnd{
								Name: "And 1",
								StandardOrs: []StandardOr{
									StandardOr{
										Name: "Or 1",
										GreaterThan: &Condition{
											Value:   4096,
											Context: "http-rsp-code",
											Qualifiers: []Qualifier{
												Qualifier{
													Qualifier: "foo",
													Value:     "bar",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		}},
		{"v2 standard pattern match no qualifiers", version.Number{10, 0, 0, ""}, Entry{
			Name:            "t7",
			ThreatName:      "threat1",
			Comment:         "foobar",
			Severity:        "low",
			Direction:       "client2server",
			DefaultAction:   DefaultActionBlockIp,
			BlockIpTrackBy:  "source-and-destination",
			BlockIpDuration: 1024,
			StandardSignatureType: &StandardSignatureType{
				Signatures: []StandardSignature{
					StandardSignature{
						Name:      "sig1",
						Comment:   "sig comment",
						Scope:     "transaction",
						OrderFree: true,
						StandardAnds: []StandardAnd{
							StandardAnd{
								Name: "And 1",
								StandardOrs: []StandardOr{
									StandardOr{
										Name: "Or 1",
										PatternMatch: &Pattern{
											Pattern: "abcdefgh",
											Context: "http-rsp-code",
											Negate:  true,
										},
									},
								},
							},
						},
					},
				},
			},
		}},
		{"v2 standard pattern match with qualifiers", version.Number{10, 0, 0, ""}, Entry{
			Name:          "t8",
			ThreatName:    "threat1",
			Comment:       "foobar",
			Severity:      "low",
			Direction:     "client2server",
			DefaultAction: DefaultActionAllow,
			StandardSignatureType: &StandardSignatureType{
				Signatures: []StandardSignature{
					StandardSignature{
						Name:      "sig1",
						Comment:   "sig comment",
						Scope:     "transaction",
						OrderFree: true,
						StandardAnds: []StandardAnd{
							StandardAnd{
								Name: "And 1",
								StandardOrs: []StandardOr{
									StandardOr{
										Name: "Or 1",
										PatternMatch: &Pattern{
											Pattern: "abcdefgh",
											Context: "http-rsp-code",
											Negate:  false,
											Qualifiers: []Qualifier{
												Qualifier{
													Qualifier: "foo",
													Value:     "bar",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		}},
		{"v2 combination", version.Number{10, 0, 0, ""}, Entry{
			Name:          "t9",
			ThreatName:    "myThreat",
			Comment:       "foobar",
			Severity:      "medium",
			Direction:     "client2server",
			DefaultAction: DefaultActionAlert,
			CombinationSignatureType: &CombinationSignatureType{
				OrderFree:           true,
				ThresholdTime:       7,
				IntervalTime:        8,
				AggregationCriteria: "source-and-destination",
				Signatures: []CombinationSignature{
					CombinationSignature{
						Name: "mySig",
						CombinationOrs: []CombinationOr{
							CombinationOr{
								Name:     "Or Condition 1",
								ThreatId: "10052",
							},
						},
					},
				},
			},
		}},
		{"v2 standard less than no qualifiers", version.Number{10, 0, 0, ""}, Entry{
			Name:          "t10",
			ThreatName:    "threat1",
			Comment:       "foobar",
			Severity:      "low",
			Direction:     "client2server",
			DefaultAction: DefaultActionAllow,
			Cves:          []string{"cve1", "cve2"},
			Bugtraqs:      []string{"bug1", "bug2"},
			Vendors:       []string{"vendor1", "vendor2"},
			References:    []string{"ref1", "ref2"},
			StandardSignatureType: &StandardSignatureType{
				Signatures: []StandardSignature{
					StandardSignature{
						Name:      "sig1",
						Comment:   "sig comment",
						Scope:     "transaction",
						OrderFree: true,
						StandardAnds: []StandardAnd{
							StandardAnd{
								Name: "And 1",
								StandardOrs: []StandardOr{
									StandardOr{
										Name: "Or 1",
										LessThan: &Condition{
											Value:   4096,
											Context: "http-rsp-code",
										},
									},
								},
							},
						},
					},
				},
			},
		}},
		{"v2 standard less than with qualifiers", version.Number{10, 0, 0, ""}, Entry{
			Name:          "t11",
			ThreatName:    "threat1",
			Comment:       "foobar",
			Severity:      "low",
			Direction:     "client2server",
			DefaultAction: DefaultActionAlert,
			StandardSignatureType: &StandardSignatureType{
				Signatures: []StandardSignature{
					StandardSignature{
						Name:      "sig1",
						Comment:   "sig comment",
						Scope:     "transaction",
						OrderFree: true,
						StandardAnds: []StandardAnd{
							StandardAnd{
								Name: "And 1",
								StandardOrs: []StandardOr{
									StandardOr{
										Name: "Or 1",
										LessThan: &Condition{
											Value:   4096,
											Context: "http-rsp-code",
											Qualifiers: []Qualifier{
												Qualifier{
													Qualifier: "foo",
													Value:     "bar",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		}},
		{"v2 standard equal to no qualifiers", version.Number{10, 0, 0, ""}, Entry{
			Name:          "t12",
			ThreatName:    "threat1",
			Comment:       "foobar",
			Severity:      "low",
			Direction:     "client2server",
			DefaultAction: DefaultActionDrop,
			StandardSignatureType: &StandardSignatureType{
				Signatures: []StandardSignature{
					StandardSignature{
						Name:      "sig1",
						Comment:   "sig comment",
						Scope:     "transaction",
						OrderFree: true,
						StandardAnds: []StandardAnd{
							StandardAnd{
								Name: "And 1",
								StandardOrs: []StandardOr{
									StandardOr{
										Name: "Or 1",
										EqualTo: &EqualTo{
											Value:   4096,
											Context: "http-rsp-code",
										},
									},
								},
							},
						},
					},
				},
			},
		}},
		{"v2 standard equal to with qualifiers", version.Number{10, 0, 0, ""}, Entry{
			Name:          "t13",
			ThreatName:    "threat1",
			Comment:       "foobar",
			Severity:      "low",
			Direction:     "client2server",
			DefaultAction: DefaultActionResetClient,
			StandardSignatureType: &StandardSignatureType{
				Signatures: []StandardSignature{
					StandardSignature{
						Name:      "sig1",
						Comment:   "sig comment",
						Scope:     "transaction",
						OrderFree: true,
						StandardAnds: []StandardAnd{
							StandardAnd{
								Name: "And 1",
								StandardOrs: []StandardOr{
									StandardOr{
										Name: "Or 1",
										EqualTo: &EqualTo{
											Value:   4096,
											Context: "http-rsp-code",
											Qualifiers: []Qualifier{
												Qualifier{
													Qualifier: "foo",
													Value:     "bar",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		}},
		{"v2 standard greater than no qualifiers", version.Number{10, 0, 0, ""}, Entry{
			Name:          "t14",
			ThreatName:    "threat1",
			Comment:       "foobar",
			Severity:      "low",
			Direction:     "client2server",
			DefaultAction: DefaultActionResetServer,
			StandardSignatureType: &StandardSignatureType{
				Signatures: []StandardSignature{
					StandardSignature{
						Name:      "sig1",
						Comment:   "sig comment",
						Scope:     "transaction",
						OrderFree: true,
						StandardAnds: []StandardAnd{
							StandardAnd{
								Name: "And 1",
								StandardOrs: []StandardOr{
									StandardOr{
										Name: "Or 1",
										GreaterThan: &Condition{
											Value:   4096,
											Context: "http-rsp-code",
										},
									},
								},
							},
						},
					},
				},
			},
		}},
		{"v2 standard greater than with qualifiers", version.Number{10, 0, 0, ""}, Entry{
			Name:          "t15",
			ThreatName:    "threat1",
			Comment:       "foobar",
			Severity:      "low",
			Direction:     "client2server",
			DefaultAction: DefaultActionResetBoth,
			StandardSignatureType: &StandardSignatureType{
				Signatures: []StandardSignature{
					StandardSignature{
						Name:      "sig1",
						Comment:   "sig comment",
						Scope:     "transaction",
						OrderFree: true,
						StandardAnds: []StandardAnd{
							StandardAnd{
								Name: "And 1",
								StandardOrs: []StandardOr{
									StandardOr{
										Name: "Or 1",
										GreaterThan: &Condition{
											Value:   4096,
											Context: "http-rsp-code",
											Qualifiers: []Qualifier{
												Qualifier{
													Qualifier: "foo",
													Value:     "bar",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		}},
		{"v2 standard pattern match no qualifiers", version.Number{10, 0, 0, ""}, Entry{
			Name:            "t16",
			ThreatName:      "threat1",
			Comment:         "foobar",
			Severity:        "low",
			Direction:       "client2server",
			DefaultAction:   DefaultActionBlockIp,
			BlockIpTrackBy:  "source-and-destination",
			BlockIpDuration: 1024,
			StandardSignatureType: &StandardSignatureType{
				Signatures: []StandardSignature{
					StandardSignature{
						Name:      "sig1",
						Comment:   "sig comment",
						Scope:     "transaction",
						OrderFree: true,
						StandardAnds: []StandardAnd{
							StandardAnd{
								Name: "And 1",
								StandardOrs: []StandardOr{
									StandardOr{
										Name: "Or 1",
										PatternMatch: &Pattern{
											Pattern: "abcdefgh",
											Context: "http-rsp-code",
											Negate:  true,
										},
									},
								},
							},
						},
					},
				},
			},
		}},
		{"v2 standard pattern match with qualifiers", version.Number{10, 0, 0, ""}, Entry{
			Name:          "t17",
			ThreatName:    "threat1",
			Comment:       "foobar",
			Severity:      "low",
			Direction:     "client2server",
			DefaultAction: DefaultActionAllow,
			StandardSignatureType: &StandardSignatureType{
				Signatures: []StandardSignature{
					StandardSignature{
						Name:      "sig1",
						Comment:   "sig comment",
						Scope:     "transaction",
						OrderFree: true,
						StandardAnds: []StandardAnd{
							StandardAnd{
								Name: "And 1",
								StandardOrs: []StandardOr{
									StandardOr{
										Name: "Or 1",
										PatternMatch: &Pattern{
											Pattern: "abcdefgh",
											Context: "http-rsp-code",
											Negate:  false,
											Qualifiers: []Qualifier{
												Qualifier{
													Qualifier: "foo",
													Value:     "bar",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		}},
		{"v2 combination order free", version.Number{10, 0, 0, ""}, Entry{
			Name:          "t18",
			ThreatName:    "myThreat",
			Comment:       "foobar",
			Severity:      "medium",
			Direction:     "client2server",
			DefaultAction: DefaultActionAlert,
			CombinationSignatureType: &CombinationSignatureType{
				OrderFree:           true,
				ThresholdTime:       7,
				IntervalTime:        8,
				AggregationCriteria: "source-and-destination",
				Signatures: []CombinationSignature{
					CombinationSignature{
						Name: "And Condition 1",
						CombinationOrs: []CombinationOr{
							CombinationOr{
								Name:     "Or Condition 1",
								ThreatId: "10052",
							},
						},
					},
				},
			},
		}},
		{"v2 combination ordered", version.Number{10, 0, 0, ""}, Entry{
			Name:          "t19",
			ThreatName:    "myThreat",
			Comment:       "foobar",
			Severity:      "medium",
			Direction:     "client2server",
			DefaultAction: DefaultActionAlert,
			CombinationSignatureType: &CombinationSignatureType{
				ThresholdTime:       7,
				IntervalTime:        8,
				AggregationCriteria: "source-and-destination",
				Signatures: []CombinationSignature{
					CombinationSignature{
						Name: "And Condition 1",
						CombinationOrs: []CombinationOr{
							CombinationOr{
								Name:     "Or Condition 1",
								ThreatId: "10052",
							},
						},
					},
				},
			},
		}},
		{"v2 standard equal to and negate", version.Number{10, 0, 0, ""}, Entry{
			Name:          "t20",
			ThreatName:    "threat1",
			Comment:       "foobar",
			Severity:      "high",
			Direction:     "client2server",
			DefaultAction: DefaultActionDrop,
			StandardSignatureType: &StandardSignatureType{
				Signatures: []StandardSignature{
					StandardSignature{
						Name:      "sig1",
						Comment:   "sig comment",
						Scope:     "transaction",
						OrderFree: true,
						StandardAnds: []StandardAnd{
							StandardAnd{
								Name: "And 1",
								StandardOrs: []StandardOr{
									StandardOr{
										Name: "Or 1",
										EqualTo: &EqualTo{
											Value:   4096,
											Context: "http-rsp-code",
											Negate:  true,
										},
									},
								},
							},
						},
					},
				},
			},
		}},
	}
}
