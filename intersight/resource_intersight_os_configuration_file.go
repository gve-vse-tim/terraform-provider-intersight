package intersight

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"strings"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceOsConfigurationFile() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceOsConfigurationFileCreate,
		ReadContext:   resourceOsConfigurationFileRead,
		UpdateContext: resourceOsConfigurationFileUpdate,
		DeleteContext: resourceOsConfigurationFileDelete,
		Importer:      &schema.ResourceImporter{StateContext: schema.ImportStatePassthroughContext},
		CustomizeDiff: CustomizeTagDiff,
		Schema: map[string]*schema.Schema{
			"account_moid": {
				Description: "The Account ID for this managed object.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					if val != nil {
						warns = append(warns, fmt.Sprintf("Cannot set read-only property: [%s]", key))
					}
					return
				}},
			"additional_properties": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: SuppressDiffAdditionProps,
			},
			"ancestors": {
				Description: "An array of relationships to moBaseMo resources.",
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				ConfigMode:  schema.SchemaConfigModeAttr,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
						},
						"class_id": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "mo.MoRef",
						},
						"moid": {
							Description: "The Moid of the referenced REST resource.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"object_type": {
							Description: "The fully-qualified name of the remote type referred by this relationship.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"selector": {
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
			},
			"catalog": {
				Description: "A reference to a osCatalog resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				ConfigMode:  schema.SchemaConfigModeAttr,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
						},
						"class_id": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "mo.MoRef",
						},
						"moid": {
							Description: "The Moid of the referenced REST resource.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"object_type": {
							Description: "The fully-qualified name of the remote type referred by this relationship.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"selector": {
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
			},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "os.ConfigurationFile",
			},
			"create_time": {
				Description: "The time when this managed object was created.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					if val != nil {
						warns = append(warns, fmt.Sprintf("Cannot set read-only property: [%s]", key))
					}
					return
				}},
			"description": {
				Description: "Description of the OS ConfigurationFile.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"distributions": {
				Description: "An array of relationships to hclOperatingSystem resources.",
				Type:        schema.TypeList,
				Optional:    true,
				ConfigMode:  schema.SchemaConfigModeAttr,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
						},
						"class_id": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "mo.MoRef",
						},
						"moid": {
							Description: "The Moid of the referenced REST resource.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"object_type": {
							Description: "The fully-qualified name of the remote type referred by this relationship.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"selector": {
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
			},
			"domain_group_moid": {
				Description: "The DomainGroup ID for this managed object.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					if val != nil {
						warns = append(warns, fmt.Sprintf("Cannot set read-only property: [%s]", key))
					}
					return
				}},
			"file_content": {
				Description: "The content of the entire configuration file is stored as value. The content\ncan either be a static file content or a template content.\nThe template is expected to conform to the golang template syntax. The values\nfrom os.Answers properties will be used to populate this template.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"internal": {
				Description: "The internal flag is set to true when configuration file is uploaded from OS Install wizard. Internal Configuration files will not be displayed in Answer Management Page.",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
			},
			"mod_time": {
				Description: "The time when this managed object was last modified.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					if val != nil {
						warns = append(warns, fmt.Sprintf("Cannot set read-only property: [%s]", key))
					}
					return
				}},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
			},
			"name": {
				Description: "The name of the OS ConfigurationFile that uniquely identifies the configuration file.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "os.ConfigurationFile",
			},
			"owners": {
				Type:       schema.TypeList,
				Optional:   true,
				Computed:   true,
				ConfigMode: schema.SchemaConfigModeAttr,
				Elem: &schema.Schema{
					Type: schema.TypeString}},
			"parent": {
				Description: "A reference to a moBaseMo resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Computed:    true,
				ConfigMode:  schema.SchemaConfigModeAttr,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
						},
						"class_id": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "mo.MoRef",
						},
						"moid": {
							Description: "The Moid of the referenced REST resource.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"object_type": {
							Description: "The fully-qualified name of the remote type referred by this relationship.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"selector": {
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
			},
			"permission_resources": {
				Description: "An array of relationships to moBaseMo resources.",
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				ConfigMode:  schema.SchemaConfigModeAttr,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
						},
						"class_id": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "mo.MoRef",
						},
						"moid": {
							Description: "The Moid of the referenced REST resource.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"object_type": {
							Description: "The fully-qualified name of the remote type referred by this relationship.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"selector": {
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
			},
			"placeholders": {
				Type:       schema.TypeList,
				Optional:   true,
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
						},
						"class_id": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "os.PlaceHolder",
						},
						"is_value_set": {
							Description: "Flag to indicate if value is set. Value will be used to check if any edit.",
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     true,
						},
						"object_type": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "os.PlaceHolder",
						},
						"type": {
							Description: "Definition of place holder.",
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							ConfigMode:  schema.SchemaConfigModeAttr,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"additional_properties": {
										Type:             schema.TypeString,
										Optional:         true,
										DiffSuppressFunc: SuppressDiffAdditionProps,
									},
									"class_id": {
										Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
										Type:        schema.TypeString,
										Optional:    true,
										Default:     "workflow.PrimitiveDataType",
									},
									"default": {
										Description: "Default value for the data type. If default value was provided and the input was required the default value will be used as the input.",
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
										ConfigMode:  schema.SchemaConfigModeAttr,
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"additional_properties": {
													Type:             schema.TypeString,
													Optional:         true,
													DiffSuppressFunc: SuppressDiffAdditionProps,
												},
												"class_id": {
													Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
													Type:        schema.TypeString,
													Optional:    true,
													Default:     "workflow.DefaultValue",
												},
												"is_value_set": {
													Description: "A flag that indicates whether a default value is given or not. This flag will be useful in case of the secure parameter where the value will be filtered out in API responses.",
													Type:        schema.TypeBool,
													Optional:    true,
													Computed:    true,
													ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
														if val != nil {
															warns = append(warns, fmt.Sprintf("Cannot set read-only property: [%s]", key))
														}
														return
													}},
												"object_type": {
													Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
													Type:        schema.TypeString,
													Optional:    true,
													Default:     "workflow.DefaultValue",
												},
												"override": {
													Description: "Override the default value provided for the data type. When true, allow the user to enter value for the data type.",
													Type:        schema.TypeBool,
													Optional:    true,
												},
												"value": {
													Description: "Default value for the data type. If default value was provided and the input was required the default value will be used as the input.",
													Type:        schema.TypeString,
													Optional:    true,
												},
											},
										},
									},
									"description": {
										Description: "Provide a detailed description of the data type.",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"display_meta": {
										Description: "Captures the meta data needed for displaying workflow data types in Intersight User Interface.",
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
										ConfigMode:  schema.SchemaConfigModeAttr,
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"additional_properties": {
													Type:             schema.TypeString,
													Optional:         true,
													DiffSuppressFunc: SuppressDiffAdditionProps,
												},
												"class_id": {
													Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
													Type:        schema.TypeString,
													Optional:    true,
													Default:     "workflow.DisplayMeta",
												},
												"inventory_selector": {
													Description: "Inventory selector specified for primitive data property should be used in Intersight User Interface.",
													Type:        schema.TypeBool,
													Optional:    true,
													Default:     true,
												},
												"object_type": {
													Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
													Type:        schema.TypeString,
													Optional:    true,
													Default:     "workflow.DisplayMeta",
												},
												"widget_type": {
													Description: "Specify the widget type for data display.\n* `None` - Display none of the widget types.\n* `Radio` - Display the widget as a radio button.\n* `Dropdown` - Display the widget as a dropdown.\n* `GridSelector` - Display the widget as a selector.\n* `DrawerSelector` - Display the widget as a selector.",
													Type:        schema.TypeString,
													Optional:    true,
													Default:     "None",
												},
											},
										},
									},
									"input_parameters": {
										Description: "JSON formatted mapping from other property of the definition to the current property. Input parameter mapping is supported only for custom data type property in workflow definition and custom data type definition. The format to specify mapping ina workflow definition when source property is of scalar types is '${workflow.input.property}'. The format to specify mapping when the source property is of object reference and mapping needs to be made to the property of the object is '${workflow.input.property.subproperty}'. The format to specify mapping in a custom data type definition is '${datatype.type.property}'. When the current property is of non-scalar type like composite custom data type, then mapping can be provided to the individual property of the custom data type like 'cdt_property:${workflow.input.property}'.",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"label": {
										Description: "Descriptive label for the data type. Label can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), space ( ) or an underscore (_). The first and last character in label must be an alphanumeric character.",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"name": {
										Description: "Descriptive name for the data type. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-) or an underscore (_). The first and last character in name must be an alphanumeric character.",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"object_type": {
										Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
										Type:        schema.TypeString,
										Optional:    true,
										Default:     "workflow.PrimitiveDataType",
									},
									"properties": {
										Description: "Primitive data type properties.",
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
										ConfigMode:  schema.SchemaConfigModeAttr,
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"additional_properties": {
													Type:             schema.TypeString,
													Optional:         true,
													DiffSuppressFunc: SuppressDiffAdditionProps,
												},
												"class_id": {
													Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
													Type:        schema.TypeString,
													Optional:    true,
													Default:     "workflow.PrimitiveDataProperty",
												},
												"constraints": {
													Description: "Constraints that must be applied to the parameter value supplied for this data type.",
													Type:        schema.TypeList,
													MaxItems:    1,
													Optional:    true,
													ConfigMode:  schema.SchemaConfigModeAttr,
													Computed:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"additional_properties": {
																Type:             schema.TypeString,
																Optional:         true,
																DiffSuppressFunc: SuppressDiffAdditionProps,
															},
															"class_id": {
																Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
																Type:        schema.TypeString,
																Optional:    true,
																Default:     "workflow.Constraints",
															},
															"enum_list": {
																Type:       schema.TypeList,
																Optional:   true,
																ConfigMode: schema.SchemaConfigModeAttr,
																Computed:   true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"additional_properties": {
																			Type:             schema.TypeString,
																			Optional:         true,
																			DiffSuppressFunc: SuppressDiffAdditionProps,
																		},
																		"class_id": {
																			Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
																			Type:        schema.TypeString,
																			Optional:    true,
																			Default:     "workflow.EnumEntry",
																		},
																		"label": {
																			Description: "Label for the enum value. A user friendly short string to identify the enum value. Label can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:), space ( ), single quote ('), forward slash (/), or an underscore (_) and must have an alphanumeric character.",
																			Type:        schema.TypeString,
																			Optional:    true,
																		},
																		"object_type": {
																			Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
																			Type:        schema.TypeString,
																			Optional:    true,
																			Default:     "workflow.EnumEntry",
																		},
																		"value": {
																			Description: "Enum value for this enum entry. Value will be passed to the workflow as string type for execution. Value can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:), space ( ), forward slash (/), or an underscore (_).",
																			Type:        schema.TypeString,
																			Optional:    true,
																		},
																	},
																},
															},
															"max": {
																Description: "Allowed maximum value of the parameter if parameter is integer/float or maximum length of the parameter if the parameter is string. When max and min are set to 0, then the limits are not checked. If parameter is integer/float, then maximum number supported is 1.797693134862315708145274237317043567981e+308 or (2**1023 * (2**53 - 1) / 2**52). When a number bigger than this is given as Maximum value, the constraints will not be enforced.",
																Type:        schema.TypeFloat,
																Optional:    true,
															},
															"min": {
																Description: "Allowed minimum value of the parameter if parameter is integer/float or minimum length of the parameter if the parameter is string. When max and min are set to 0, then the limits are not checked. If parameter is integer/float, then minimum number supported is 4.940656458412465441765687928682213723651e-324 or (1 / 2 ** (1023 - 1 + 52)). When a number smaller than this is given as minimum value, the constraints will not be enforced.",
																Type:        schema.TypeFloat,
																Optional:    true,
															},
															"object_type": {
																Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
																Type:        schema.TypeString,
																Optional:    true,
																Default:     "workflow.Constraints",
															},
															"regex": {
																Description: "When the parameter is a string this regular expression is used to ensure the value is valid.",
																Type:        schema.TypeString,
																Optional:    true,
															},
														},
													},
												},
												"inventory_selector": {
													Type:       schema.TypeList,
													Optional:   true,
													ConfigMode: schema.SchemaConfigModeAttr,
													Computed:   true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"additional_properties": {
																Type:             schema.TypeString,
																Optional:         true,
																DiffSuppressFunc: SuppressDiffAdditionProps,
															},
															"class_id": {
																Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
																Type:        schema.TypeString,
																Optional:    true,
																Default:     "workflow.MoReferenceProperty",
															},
															"display_attributes": {
																Type:       schema.TypeList,
																Optional:   true,
																ConfigMode: schema.SchemaConfigModeAttr,
																Computed:   true,
																Elem: &schema.Schema{
																	Type: schema.TypeString}},
															"object_type": {
																Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
																Type:        schema.TypeString,
																Optional:    true,
																Default:     "workflow.MoReferenceProperty",
															},
															"selector": {
																Description: "Field to hold an Intersight API along with an optional filter to narrow down the search options.",
																Type:        schema.TypeString,
																Optional:    true,
															},
															"selector_property": {
																Description: "Selector properties to define HTTP method and 'body' in case of upsert operation.",
																Type:        schema.TypeList,
																MaxItems:    1,
																Optional:    true,
																ConfigMode:  schema.SchemaConfigModeAttr,
																Computed:    true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"additional_properties": {
																			Type:             schema.TypeString,
																			Optional:         true,
																			DiffSuppressFunc: SuppressDiffAdditionProps,
																		},
																		"body": {
																			Description: "Content of the request body to send for POST request.",
																			Type:        schema.TypeString,
																			Optional:    true,
																		},
																		"class_id": {
																			Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
																			Type:        schema.TypeString,
																			Optional:    true,
																			Default:     "workflow.SelectorProperty",
																		},
																		"method": {
																			Description: "The HTTP method to be used.\n* `GET` - The HTTP GET method requests a representation of the specified resource.\n* `POST` - The HTTP POST method sends data to the server.",
																			Type:        schema.TypeString,
																			Optional:    true,
																			Default:     "GET",
																		},
																		"object_type": {
																			Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
																			Type:        schema.TypeString,
																			Optional:    true,
																			Default:     "workflow.SelectorProperty",
																		},
																	},
																},
															},
															"value_attribute": {
																Description: "A property from the Intersight object, value of which can be used as value for referenced input definition.",
																Type:        schema.TypeString,
																Optional:    true,
															},
														},
													},
												},
												"object_type": {
													Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
													Type:        schema.TypeString,
													Optional:    true,
													Default:     "workflow.PrimitiveDataProperty",
												},
												"secure": {
													Description: "Intersight supports secure properties as task input/output. The values of\nthese properties are encrypted and stored in Intersight.\nThis flag marks the property to be secure when it is set to true.",
													Type:        schema.TypeBool,
													Optional:    true,
												},
												"type": {
													Description: "Specify the enum type for primitive data type.\n* `string` - Enum to specify a string data type.\n* `integer` - Enum to specify an integer32 data type.\n* `float` - Enum to specify a float64 data type.\n* `boolean` - Enum to specify a boolean data type.\n* `json` - Enum to specify a json data type.\n* `enum` - Enum to specify a enum data type which is a list of pre-defined strings.",
													Type:        schema.TypeString,
													Optional:    true,
													Default:     "string",
												},
											},
										},
									},
									"required": {
										Description: "Specifies whether this parameter is required. The field is applicable for task and workflow.",
										Type:        schema.TypeBool,
										Optional:    true,
									},
								},
							},
						},
						"value": {
							Description: "Value for placeholder provided by user.",
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
			},
			"shared_scope": {
				Description: "Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.\nObjects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					if val != nil {
						warns = append(warns, fmt.Sprintf("Cannot set read-only property: [%s]", key))
					}
					return
				}},
			"supported": {
				Description: "An internal property that is used to distinguish between the pre-canned OS\nconfiguration file entries and user provided entries.",
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					if val != nil {
						warns = append(warns, fmt.Sprintf("Cannot set read-only property: [%s]", key))
					}
					return
				}},
			"tags": {
				Type:       schema.TypeList,
				Optional:   true,
				ConfigMode: schema.SchemaConfigModeAttr,
				Computed:   true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
						},
						"key": {
							Description: "The string representation of a tag key.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"value": {
							Description: "The string representation of a tag value.",
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
			},
			"version_context": {
				Description: "The versioning info for this managed object.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Computed:    true,
				ConfigMode:  schema.SchemaConfigModeAttr,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
						},
						"class_id": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "mo.VersionContext",
						},
						"interested_mos": {
							Type:       schema.TypeList,
							Optional:   true,
							ConfigMode: schema.SchemaConfigModeAttr,
							Computed:   true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"additional_properties": {
										Type:             schema.TypeString,
										Optional:         true,
										DiffSuppressFunc: SuppressDiffAdditionProps,
									},
									"class_id": {
										Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
										Type:        schema.TypeString,
										Optional:    true,
										Default:     "mo.MoRef",
									},
									"moid": {
										Description: "The Moid of the referenced REST resource.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"object_type": {
										Description: "The fully-qualified name of the remote type referred by this relationship.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"selector": {
										Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
										Type:        schema.TypeString,
										Optional:    true,
									},
								},
							},
						},
						"object_type": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "mo.VersionContext",
						},
						"ref_mo": {
							Description: "A reference to the original Managed Object.",
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Computed:    true,
							ConfigMode:  schema.SchemaConfigModeAttr,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"additional_properties": {
										Type:             schema.TypeString,
										Optional:         true,
										DiffSuppressFunc: SuppressDiffAdditionProps,
									},
									"class_id": {
										Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
										Type:        schema.TypeString,
										Optional:    true,
										Default:     "mo.MoRef",
									},
									"moid": {
										Description: "The Moid of the referenced REST resource.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"object_type": {
										Description: "The fully-qualified name of the remote type referred by this relationship.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"selector": {
										Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
										Type:        schema.TypeString,
										Optional:    true,
									},
								},
							},
						},
						"timestamp": {
							Description: "The time this versioned Managed Object was created.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
								if val != nil {
									warns = append(warns, fmt.Sprintf("Cannot set read-only property: [%s]", key))
								}
								return
							}},
						"nr_version": {
							Description: "The version of the Managed Object, e.g. an incrementing number or a hash id.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
								if val != nil {
									warns = append(warns, fmt.Sprintf("Cannot set read-only property: [%s]", key))
								}
								return
							}},
						"version_type": {
							Description: "Specifies type of version. Currently the only supported value is \"Configured\"\nthat is used to keep track of snapshots of policies and profiles that are intended\nto be configured to target endpoints.\n* `Modified` - Version created every time an object is modified.\n* `Configured` - Version created every time an object is configured to the service profile.\n* `Deployed` - Version created for objects related to a service profile when it is deployed.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
								if val != nil {
									warns = append(warns, fmt.Sprintf("Cannot set read-only property: [%s]", key))
								}
								return
							}},
					},
				},
			},
		},
	}
}

func resourceOsConfigurationFileCreate(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = models.NewOsConfigurationFileWithDefaults()

	if v, ok := d.GetOk("additional_properties"); ok {
		x := []byte(v.(string))
		var x1 interface{}
		err := json.Unmarshal(x, &x1)
		if err == nil && x1 != nil {
			o.AdditionalProperties = x1.(map[string]interface{})
		}
	}

	if v, ok := d.GetOk("catalog"); ok {
		p := make([]models.OsCatalogRelationship, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := models.NewMoMoRefWithDefaults()
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("")
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			p = append(p, models.MoMoRefAsOsCatalogRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetCatalog(x)
		}
	}

	o.SetClassId("os.ConfigurationFile")

	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.SetDescription(x)
	}

	if v, ok := d.GetOk("distributions"); ok {
		x := make([]models.HclOperatingSystemRelationship, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := models.NewMoMoRefWithDefaults()
			l := s[i].(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("mo.MoRef")
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			x = append(x, models.MoMoRefAsHclOperatingSystemRelationship(o))
		}
		if len(x) > 0 {
			o.SetDistributions(x)
		}
	}

	if v, ok := d.GetOk("file_content"); ok {
		x := (v.(string))
		o.SetFileContent(x)
	}

	if v, ok := d.GetOkExists("internal"); ok {
		x := (v.(bool))
		o.SetInternal(x)
	}

	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}

	if v, ok := d.GetOk("name"); ok {
		x := (v.(string))
		o.SetName(x)
	}

	o.SetObjectType("os.ConfigurationFile")

	if v, ok := d.GetOk("placeholders"); ok {
		x := make([]models.OsPlaceHolder, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := models.NewOsPlaceHolderWithDefaults()
			l := s[i].(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("os.PlaceHolder")
			if v, ok := l["is_value_set"]; ok {
				{
					x := (v.(bool))
					o.SetIsValueSet(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["type"]; ok {
				{
					p := make([]models.WorkflowPrimitiveDataType, 0, 1)
					s := v.([]interface{})
					for i := 0; i < len(s); i++ {
						l := s[i].(map[string]interface{})
						o := models.NewWorkflowPrimitiveDataTypeWithDefaults()
						if v, ok := l["additional_properties"]; ok {
							{
								x := []byte(v.(string))
								var x1 interface{}
								err := json.Unmarshal(x, &x1)
								if err == nil && x1 != nil {
									o.AdditionalProperties = x1.(map[string]interface{})
								}
							}
						}
						o.SetClassId("")
						if v, ok := l["default"]; ok {
							{
								p := make([]models.WorkflowDefaultValue, 0, 1)
								s := v.([]interface{})
								for i := 0; i < len(s); i++ {
									l := s[i].(map[string]interface{})
									o := models.NewWorkflowDefaultValueWithDefaults()
									if v, ok := l["additional_properties"]; ok {
										{
											x := []byte(v.(string))
											var x1 interface{}
											err := json.Unmarshal(x, &x1)
											if err == nil && x1 != nil {
												o.AdditionalProperties = x1.(map[string]interface{})
											}
										}
									}
									o.SetClassId("")
									if v, ok := l["object_type"]; ok {
										{
											x := (v.(string))
											o.SetObjectType(x)
										}
									}
									if v, ok := l["override"]; ok {
										{
											x := (v.(bool))
											o.SetOverride(x)
										}
									}
									if v, ok := l["value"]; ok {
										{
											x := []byte(v.(string))
											var x1 interface{}
											err := json.Unmarshal(x, &x1)
											if err == nil && x1 != nil {
												x2 := x1.(map[string]interface{})
												o.SetValue(x2)
											}
										}
									}
									p = append(p, *o)
								}
								if len(p) > 0 {
									x := p[0]
									o.SetDefault(x)
								}
							}
						}
						if v, ok := l["description"]; ok {
							{
								x := (v.(string))
								o.SetDescription(x)
							}
						}
						if v, ok := l["display_meta"]; ok {
							{
								p := make([]models.WorkflowDisplayMeta, 0, 1)
								s := v.([]interface{})
								for i := 0; i < len(s); i++ {
									l := s[i].(map[string]interface{})
									o := models.NewWorkflowDisplayMetaWithDefaults()
									if v, ok := l["additional_properties"]; ok {
										{
											x := []byte(v.(string))
											var x1 interface{}
											err := json.Unmarshal(x, &x1)
											if err == nil && x1 != nil {
												o.AdditionalProperties = x1.(map[string]interface{})
											}
										}
									}
									o.SetClassId("")
									if v, ok := l["inventory_selector"]; ok {
										{
											x := (v.(bool))
											o.SetInventorySelector(x)
										}
									}
									if v, ok := l["object_type"]; ok {
										{
											x := (v.(string))
											o.SetObjectType(x)
										}
									}
									if v, ok := l["widget_type"]; ok {
										{
											x := (v.(string))
											o.SetWidgetType(x)
										}
									}
									p = append(p, *o)
								}
								if len(p) > 0 {
									x := p[0]
									o.SetDisplayMeta(x)
								}
							}
						}
						if v, ok := l["input_parameters"]; ok {
							{
								x := []byte(v.(string))
								var x1 interface{}
								err := json.Unmarshal(x, &x1)
								if err == nil && x1 != nil {
									x2 := x1.(map[string]interface{})
									o.SetInputParameters(x2)
								}
							}
						}
						if v, ok := l["label"]; ok {
							{
								x := (v.(string))
								o.SetLabel(x)
							}
						}
						if v, ok := l["name"]; ok {
							{
								x := (v.(string))
								o.SetName(x)
							}
						}
						if v, ok := l["object_type"]; ok {
							{
								x := (v.(string))
								o.SetObjectType(x)
							}
						}
						if v, ok := l["properties"]; ok {
							{
								p := make([]models.WorkflowPrimitiveDataProperty, 0, 1)
								s := v.([]interface{})
								for i := 0; i < len(s); i++ {
									l := s[i].(map[string]interface{})
									o := models.NewWorkflowPrimitiveDataPropertyWithDefaults()
									if v, ok := l["additional_properties"]; ok {
										{
											x := []byte(v.(string))
											var x1 interface{}
											err := json.Unmarshal(x, &x1)
											if err == nil && x1 != nil {
												o.AdditionalProperties = x1.(map[string]interface{})
											}
										}
									}
									o.SetClassId("")
									if v, ok := l["constraints"]; ok {
										{
											p := make([]models.WorkflowConstraints, 0, 1)
											s := v.([]interface{})
											for i := 0; i < len(s); i++ {
												l := s[i].(map[string]interface{})
												o := models.NewWorkflowConstraintsWithDefaults()
												if v, ok := l["additional_properties"]; ok {
													{
														x := []byte(v.(string))
														var x1 interface{}
														err := json.Unmarshal(x, &x1)
														if err == nil && x1 != nil {
															o.AdditionalProperties = x1.(map[string]interface{})
														}
													}
												}
												o.SetClassId("")
												if v, ok := l["enum_list"]; ok {
													{
														x := make([]models.WorkflowEnumEntry, 0)
														s := v.([]interface{})
														for i := 0; i < len(s); i++ {
															o := models.NewWorkflowEnumEntryWithDefaults()
															l := s[i].(map[string]interface{})
															if v, ok := l["additional_properties"]; ok {
																{
																	x := []byte(v.(string))
																	var x1 interface{}
																	err := json.Unmarshal(x, &x1)
																	if err == nil && x1 != nil {
																		o.AdditionalProperties = x1.(map[string]interface{})
																	}
																}
															}
															o.SetClassId("workflow.EnumEntry")
															if v, ok := l["label"]; ok {
																{
																	x := (v.(string))
																	o.SetLabel(x)
																}
															}
															if v, ok := l["object_type"]; ok {
																{
																	x := (v.(string))
																	o.SetObjectType(x)
																}
															}
															if v, ok := l["value"]; ok {
																{
																	x := (v.(string))
																	o.SetValue(x)
																}
															}
															x = append(x, *o)
														}
														if len(x) > 0 {
															o.SetEnumList(x)
														}
													}
												}
												if v, ok := l["max"]; ok {
													{
														x := v.(float64)
														o.SetMax(x)
													}
												}
												if v, ok := l["min"]; ok {
													{
														x := v.(float64)
														o.SetMin(x)
													}
												}
												if v, ok := l["object_type"]; ok {
													{
														x := (v.(string))
														o.SetObjectType(x)
													}
												}
												if v, ok := l["regex"]; ok {
													{
														x := (v.(string))
														o.SetRegex(x)
													}
												}
												p = append(p, *o)
											}
											if len(p) > 0 {
												x := p[0]
												o.SetConstraints(x)
											}
										}
									}
									if v, ok := l["inventory_selector"]; ok {
										{
											x := make([]models.WorkflowMoReferenceProperty, 0)
											s := v.([]interface{})
											for i := 0; i < len(s); i++ {
												o := models.NewWorkflowMoReferencePropertyWithDefaults()
												l := s[i].(map[string]interface{})
												if v, ok := l["additional_properties"]; ok {
													{
														x := []byte(v.(string))
														var x1 interface{}
														err := json.Unmarshal(x, &x1)
														if err == nil && x1 != nil {
															o.AdditionalProperties = x1.(map[string]interface{})
														}
													}
												}
												o.SetClassId("workflow.MoReferenceProperty")
												if v, ok := l["display_attributes"]; ok {
													{
														x := make([]string, 0)
														y := reflect.ValueOf(v)
														for i := 0; i < y.Len(); i++ {
															if y.Index(i).Interface() != nil {
																x = append(x, y.Index(i).Interface().(string))
															}
														}
														if len(x) > 0 {
															o.SetDisplayAttributes(x)
														}
													}
												}
												if v, ok := l["object_type"]; ok {
													{
														x := (v.(string))
														o.SetObjectType(x)
													}
												}
												if v, ok := l["selector"]; ok {
													{
														x := (v.(string))
														o.SetSelector(x)
													}
												}
												if v, ok := l["selector_property"]; ok {
													{
														p := make([]models.WorkflowSelectorProperty, 0, 1)
														s := v.([]interface{})
														for i := 0; i < len(s); i++ {
															l := s[i].(map[string]interface{})
															o := models.NewWorkflowSelectorPropertyWithDefaults()
															if v, ok := l["additional_properties"]; ok {
																{
																	x := []byte(v.(string))
																	var x1 interface{}
																	err := json.Unmarshal(x, &x1)
																	if err == nil && x1 != nil {
																		o.AdditionalProperties = x1.(map[string]interface{})
																	}
																}
															}
															if v, ok := l["body"]; ok {
																{
																	x := []byte(v.(string))
																	var x1 interface{}
																	err := json.Unmarshal(x, &x1)
																	if err == nil && x1 != nil {
																		x2 := x1.(map[string]interface{})
																		o.SetBody(x2)
																	}
																}
															}
															o.SetClassId("")
															if v, ok := l["method"]; ok {
																{
																	x := (v.(string))
																	o.SetMethod(x)
																}
															}
															if v, ok := l["object_type"]; ok {
																{
																	x := (v.(string))
																	o.SetObjectType(x)
																}
															}
															p = append(p, *o)
														}
														if len(p) > 0 {
															x := p[0]
															o.SetSelectorProperty(x)
														}
													}
												}
												if v, ok := l["value_attribute"]; ok {
													{
														x := (v.(string))
														o.SetValueAttribute(x)
													}
												}
												x = append(x, *o)
											}
											if len(x) > 0 {
												o.SetInventorySelector(x)
											}
										}
									}
									if v, ok := l["object_type"]; ok {
										{
											x := (v.(string))
											o.SetObjectType(x)
										}
									}
									if v, ok := l["secure"]; ok {
										{
											x := (v.(bool))
											o.SetSecure(x)
										}
									}
									if v, ok := l["type"]; ok {
										{
											x := (v.(string))
											o.SetType(x)
										}
									}
									p = append(p, *o)
								}
								if len(p) > 0 {
									x := p[0]
									o.SetProperties(x)
								}
							}
						}
						if v, ok := l["required"]; ok {
							{
								x := (v.(bool))
								o.SetRequired(x)
							}
						}
						p = append(p, *o)
					}
					if len(p) > 0 {
						x := p[0]
						o.SetType(x)
					}
				}
			}
			if v, ok := l["value"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						x2 := x1.(map[string]interface{})
						o.SetValue(x2)
					}
				}
			}
			x = append(x, *o)
		}
		if len(x) > 0 {
			o.SetPlaceholders(x)
		}
	}

	if v, ok := d.GetOk("tags"); ok {
		x := make([]models.MoTag, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := models.NewMoTagWithDefaults()
			l := s[i].(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			if v, ok := l["key"]; ok {
				{
					x := (v.(string))
					o.SetKey(x)
				}
			}
			if v, ok := l["value"]; ok {
				{
					x := (v.(string))
					o.SetValue(x)
				}
			}
			x = append(x, *o)
		}
		if len(x) > 0 {
			o.SetTags(x)
		}
	}

	r := conn.ApiClient.OsApi.CreateOsConfigurationFile(conn.ctx).OsConfigurationFile(*o)
	resultMo, _, responseErr := r.Execute()
	if responseErr != nil {
		errorType := fmt.Sprintf("%T", responseErr)
		if strings.Contains(errorType, "GenericOpenAPIError") {
			responseErr := responseErr.(*models.GenericOpenAPIError)
			return diag.Errorf("error occurred while creating OsConfigurationFile: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		return diag.Errorf("error occurred while creating OsConfigurationFile: %s", responseErr.Error())
	}
	log.Printf("Moid: %s", resultMo.GetMoid())
	d.SetId(resultMo.GetMoid())
	return append(de, resourceOsConfigurationFileRead(c, d, meta)...)
}

func resourceOsConfigurationFileRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	var de diag.Diagnostics
	conn := meta.(*Config)
	r := conn.ApiClient.OsApi.GetOsConfigurationFileByMoid(conn.ctx, d.Id())
	s, _, responseErr := r.Execute()
	if responseErr != nil {
		if strings.Contains(responseErr.Error(), "404") {
			de = append(de, diag.Diagnostic{Summary: "OsConfigurationFile object " + d.Id() + " not found. Removing from statefile", Severity: diag.Warning})
			d.SetId("")
			return de
		}
		errorType := fmt.Sprintf("%T", responseErr)
		if strings.Contains(errorType, "GenericOpenAPIError") {
			responseErr := responseErr.(*models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching OsConfigurationFile: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		return diag.Errorf("error occurred while fetching OsConfigurationFile: %s", responseErr.Error())
	}

	if err := d.Set("account_moid", (s.GetAccountMoid())); err != nil {
		return diag.Errorf("error occurred while setting property AccountMoid in OsConfigurationFile object: %s", err.Error())
	}

	if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
		return diag.Errorf("error occurred while setting property AdditionalProperties in OsConfigurationFile object: %s", err.Error())
	}

	if err := d.Set("ancestors", flattenListMoBaseMoRelationship(s.GetAncestors(), d)); err != nil {
		return diag.Errorf("error occurred while setting property Ancestors in OsConfigurationFile object: %s", err.Error())
	}

	if err := d.Set("catalog", flattenMapOsCatalogRelationship(s.GetCatalog(), d)); err != nil {
		return diag.Errorf("error occurred while setting property Catalog in OsConfigurationFile object: %s", err.Error())
	}

	if err := d.Set("class_id", (s.GetClassId())); err != nil {
		return diag.Errorf("error occurred while setting property ClassId in OsConfigurationFile object: %s", err.Error())
	}

	if err := d.Set("create_time", (s.GetCreateTime()).String()); err != nil {
		return diag.Errorf("error occurred while setting property CreateTime in OsConfigurationFile object: %s", err.Error())
	}

	if err := d.Set("description", (s.GetDescription())); err != nil {
		return diag.Errorf("error occurred while setting property Description in OsConfigurationFile object: %s", err.Error())
	}

	if err := d.Set("distributions", flattenListHclOperatingSystemRelationship(s.GetDistributions(), d)); err != nil {
		return diag.Errorf("error occurred while setting property Distributions in OsConfigurationFile object: %s", err.Error())
	}

	if err := d.Set("domain_group_moid", (s.GetDomainGroupMoid())); err != nil {
		return diag.Errorf("error occurred while setting property DomainGroupMoid in OsConfigurationFile object: %s", err.Error())
	}

	if err := d.Set("file_content", (s.GetFileContent())); err != nil {
		return diag.Errorf("error occurred while setting property FileContent in OsConfigurationFile object: %s", err.Error())
	}

	if err := d.Set("internal", (s.GetInternal())); err != nil {
		return diag.Errorf("error occurred while setting property Internal in OsConfigurationFile object: %s", err.Error())
	}

	if err := d.Set("mod_time", (s.GetModTime()).String()); err != nil {
		return diag.Errorf("error occurred while setting property ModTime in OsConfigurationFile object: %s", err.Error())
	}

	if err := d.Set("moid", (s.GetMoid())); err != nil {
		return diag.Errorf("error occurred while setting property Moid in OsConfigurationFile object: %s", err.Error())
	}

	if err := d.Set("name", (s.GetName())); err != nil {
		return diag.Errorf("error occurred while setting property Name in OsConfigurationFile object: %s", err.Error())
	}

	if err := d.Set("object_type", (s.GetObjectType())); err != nil {
		return diag.Errorf("error occurred while setting property ObjectType in OsConfigurationFile object: %s", err.Error())
	}

	if err := d.Set("owners", (s.GetOwners())); err != nil {
		return diag.Errorf("error occurred while setting property Owners in OsConfigurationFile object: %s", err.Error())
	}

	if err := d.Set("parent", flattenMapMoBaseMoRelationship(s.GetParent(), d)); err != nil {
		return diag.Errorf("error occurred while setting property Parent in OsConfigurationFile object: %s", err.Error())
	}

	if err := d.Set("permission_resources", flattenListMoBaseMoRelationship(s.GetPermissionResources(), d)); err != nil {
		return diag.Errorf("error occurred while setting property PermissionResources in OsConfigurationFile object: %s", err.Error())
	}

	if err := d.Set("placeholders", flattenListOsPlaceHolder(s.GetPlaceholders(), d)); err != nil {
		return diag.Errorf("error occurred while setting property Placeholders in OsConfigurationFile object: %s", err.Error())
	}

	if err := d.Set("shared_scope", (s.GetSharedScope())); err != nil {
		return diag.Errorf("error occurred while setting property SharedScope in OsConfigurationFile object: %s", err.Error())
	}

	if err := d.Set("supported", (s.GetSupported())); err != nil {
		return diag.Errorf("error occurred while setting property Supported in OsConfigurationFile object: %s", err.Error())
	}

	if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
		return diag.Errorf("error occurred while setting property Tags in OsConfigurationFile object: %s", err.Error())
	}

	if err := d.Set("version_context", flattenMapMoVersionContext(s.GetVersionContext(), d)); err != nil {
		return diag.Errorf("error occurred while setting property VersionContext in OsConfigurationFile object: %s", err.Error())
	}

	log.Printf("s: %v", s)
	log.Printf("Moid: %s", s.GetMoid())
	return de
}

func resourceOsConfigurationFileUpdate(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.OsConfigurationFile{}

	if d.HasChange("additional_properties") {
		v := d.Get("additional_properties")
		x := []byte(v.(string))
		var x1 interface{}
		err := json.Unmarshal(x, &x1)
		if err == nil && x1 != nil {
			o.AdditionalProperties = x1.(map[string]interface{})
		}
	}

	if d.HasChange("catalog") {
		v := d.Get("catalog")
		p := make([]models.OsCatalogRelationship, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := &models.MoMoRef{}
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("")
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			p = append(p, models.MoMoRefAsOsCatalogRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetCatalog(x)
		}
	}

	o.SetClassId("os.ConfigurationFile")

	if d.HasChange("description") {
		v := d.Get("description")
		x := (v.(string))
		o.SetDescription(x)
	}

	if d.HasChange("distributions") {
		v := d.Get("distributions")
		x := make([]models.HclOperatingSystemRelationship, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.MoMoRef{}
			l := s[i].(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("mo.MoRef")
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			x = append(x, models.MoMoRefAsHclOperatingSystemRelationship(o))
		}
		o.SetDistributions(x)
	}

	if d.HasChange("file_content") {
		v := d.Get("file_content")
		x := (v.(string))
		o.SetFileContent(x)
	}

	if d.HasChange("internal") {
		v := d.Get("internal")
		x := (v.(bool))
		o.SetInternal(x)
	}

	if d.HasChange("moid") {
		v := d.Get("moid")
		x := (v.(string))
		o.SetMoid(x)
	}

	if d.HasChange("name") {
		v := d.Get("name")
		x := (v.(string))
		o.SetName(x)
	}

	o.SetObjectType("os.ConfigurationFile")

	if d.HasChange("placeholders") {
		v := d.Get("placeholders")
		x := make([]models.OsPlaceHolder, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.OsPlaceHolder{}
			l := s[i].(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("os.PlaceHolder")
			if v, ok := l["is_value_set"]; ok {
				{
					x := (v.(bool))
					o.SetIsValueSet(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["type"]; ok {
				{
					p := make([]models.WorkflowPrimitiveDataType, 0, 1)
					s := v.([]interface{})
					for i := 0; i < len(s); i++ {
						l := s[i].(map[string]interface{})
						o := models.NewWorkflowPrimitiveDataTypeWithDefaults()
						if v, ok := l["additional_properties"]; ok {
							{
								x := []byte(v.(string))
								var x1 interface{}
								err := json.Unmarshal(x, &x1)
								if err == nil && x1 != nil {
									o.AdditionalProperties = x1.(map[string]interface{})
								}
							}
						}
						o.SetClassId("")
						if v, ok := l["default"]; ok {
							{
								p := make([]models.WorkflowDefaultValue, 0, 1)
								s := v.([]interface{})
								for i := 0; i < len(s); i++ {
									l := s[i].(map[string]interface{})
									o := models.NewWorkflowDefaultValueWithDefaults()
									if v, ok := l["additional_properties"]; ok {
										{
											x := []byte(v.(string))
											var x1 interface{}
											err := json.Unmarshal(x, &x1)
											if err == nil && x1 != nil {
												o.AdditionalProperties = x1.(map[string]interface{})
											}
										}
									}
									o.SetClassId("")
									if v, ok := l["object_type"]; ok {
										{
											x := (v.(string))
											o.SetObjectType(x)
										}
									}
									if v, ok := l["override"]; ok {
										{
											x := (v.(bool))
											o.SetOverride(x)
										}
									}
									if v, ok := l["value"]; ok {
										{
											x := []byte(v.(string))
											var x1 interface{}
											err := json.Unmarshal(x, &x1)
											if err == nil && x1 != nil {
												x2 := x1.(map[string]interface{})
												o.SetValue(x2)
											}
										}
									}
									p = append(p, *o)
								}
								if len(p) > 0 {
									x := p[0]
									o.SetDefault(x)
								}
							}
						}
						if v, ok := l["description"]; ok {
							{
								x := (v.(string))
								o.SetDescription(x)
							}
						}
						if v, ok := l["display_meta"]; ok {
							{
								p := make([]models.WorkflowDisplayMeta, 0, 1)
								s := v.([]interface{})
								for i := 0; i < len(s); i++ {
									l := s[i].(map[string]interface{})
									o := models.NewWorkflowDisplayMetaWithDefaults()
									if v, ok := l["additional_properties"]; ok {
										{
											x := []byte(v.(string))
											var x1 interface{}
											err := json.Unmarshal(x, &x1)
											if err == nil && x1 != nil {
												o.AdditionalProperties = x1.(map[string]interface{})
											}
										}
									}
									o.SetClassId("")
									if v, ok := l["inventory_selector"]; ok {
										{
											x := (v.(bool))
											o.SetInventorySelector(x)
										}
									}
									if v, ok := l["object_type"]; ok {
										{
											x := (v.(string))
											o.SetObjectType(x)
										}
									}
									if v, ok := l["widget_type"]; ok {
										{
											x := (v.(string))
											o.SetWidgetType(x)
										}
									}
									p = append(p, *o)
								}
								if len(p) > 0 {
									x := p[0]
									o.SetDisplayMeta(x)
								}
							}
						}
						if v, ok := l["input_parameters"]; ok {
							{
								x := []byte(v.(string))
								var x1 interface{}
								err := json.Unmarshal(x, &x1)
								if err == nil && x1 != nil {
									x2 := x1.(map[string]interface{})
									o.SetInputParameters(x2)
								}
							}
						}
						if v, ok := l["label"]; ok {
							{
								x := (v.(string))
								o.SetLabel(x)
							}
						}
						if v, ok := l["name"]; ok {
							{
								x := (v.(string))
								o.SetName(x)
							}
						}
						if v, ok := l["object_type"]; ok {
							{
								x := (v.(string))
								o.SetObjectType(x)
							}
						}
						if v, ok := l["properties"]; ok {
							{
								p := make([]models.WorkflowPrimitiveDataProperty, 0, 1)
								s := v.([]interface{})
								for i := 0; i < len(s); i++ {
									l := s[i].(map[string]interface{})
									o := models.NewWorkflowPrimitiveDataPropertyWithDefaults()
									if v, ok := l["additional_properties"]; ok {
										{
											x := []byte(v.(string))
											var x1 interface{}
											err := json.Unmarshal(x, &x1)
											if err == nil && x1 != nil {
												o.AdditionalProperties = x1.(map[string]interface{})
											}
										}
									}
									o.SetClassId("")
									if v, ok := l["constraints"]; ok {
										{
											p := make([]models.WorkflowConstraints, 0, 1)
											s := v.([]interface{})
											for i := 0; i < len(s); i++ {
												l := s[i].(map[string]interface{})
												o := models.NewWorkflowConstraintsWithDefaults()
												if v, ok := l["additional_properties"]; ok {
													{
														x := []byte(v.(string))
														var x1 interface{}
														err := json.Unmarshal(x, &x1)
														if err == nil && x1 != nil {
															o.AdditionalProperties = x1.(map[string]interface{})
														}
													}
												}
												o.SetClassId("")
												if v, ok := l["enum_list"]; ok {
													{
														x := make([]models.WorkflowEnumEntry, 0)
														s := v.([]interface{})
														for i := 0; i < len(s); i++ {
															o := models.NewWorkflowEnumEntryWithDefaults()
															l := s[i].(map[string]interface{})
															if v, ok := l["additional_properties"]; ok {
																{
																	x := []byte(v.(string))
																	var x1 interface{}
																	err := json.Unmarshal(x, &x1)
																	if err == nil && x1 != nil {
																		o.AdditionalProperties = x1.(map[string]interface{})
																	}
																}
															}
															o.SetClassId("workflow.EnumEntry")
															if v, ok := l["label"]; ok {
																{
																	x := (v.(string))
																	o.SetLabel(x)
																}
															}
															if v, ok := l["object_type"]; ok {
																{
																	x := (v.(string))
																	o.SetObjectType(x)
																}
															}
															if v, ok := l["value"]; ok {
																{
																	x := (v.(string))
																	o.SetValue(x)
																}
															}
															x = append(x, *o)
														}
														if len(x) > 0 {
															o.SetEnumList(x)
														}
													}
												}
												if v, ok := l["max"]; ok {
													{
														x := v.(float64)
														o.SetMax(x)
													}
												}
												if v, ok := l["min"]; ok {
													{
														x := v.(float64)
														o.SetMin(x)
													}
												}
												if v, ok := l["object_type"]; ok {
													{
														x := (v.(string))
														o.SetObjectType(x)
													}
												}
												if v, ok := l["regex"]; ok {
													{
														x := (v.(string))
														o.SetRegex(x)
													}
												}
												p = append(p, *o)
											}
											if len(p) > 0 {
												x := p[0]
												o.SetConstraints(x)
											}
										}
									}
									if v, ok := l["inventory_selector"]; ok {
										{
											x := make([]models.WorkflowMoReferenceProperty, 0)
											s := v.([]interface{})
											for i := 0; i < len(s); i++ {
												o := models.NewWorkflowMoReferencePropertyWithDefaults()
												l := s[i].(map[string]interface{})
												if v, ok := l["additional_properties"]; ok {
													{
														x := []byte(v.(string))
														var x1 interface{}
														err := json.Unmarshal(x, &x1)
														if err == nil && x1 != nil {
															o.AdditionalProperties = x1.(map[string]interface{})
														}
													}
												}
												o.SetClassId("workflow.MoReferenceProperty")
												if v, ok := l["display_attributes"]; ok {
													{
														x := make([]string, 0)
														y := reflect.ValueOf(v)
														for i := 0; i < y.Len(); i++ {
															if y.Index(i).Interface() != nil {
																x = append(x, y.Index(i).Interface().(string))
															}
														}
														if len(x) > 0 {
															o.SetDisplayAttributes(x)
														}
													}
												}
												if v, ok := l["object_type"]; ok {
													{
														x := (v.(string))
														o.SetObjectType(x)
													}
												}
												if v, ok := l["selector"]; ok {
													{
														x := (v.(string))
														o.SetSelector(x)
													}
												}
												if v, ok := l["selector_property"]; ok {
													{
														p := make([]models.WorkflowSelectorProperty, 0, 1)
														s := v.([]interface{})
														for i := 0; i < len(s); i++ {
															l := s[i].(map[string]interface{})
															o := models.NewWorkflowSelectorPropertyWithDefaults()
															if v, ok := l["additional_properties"]; ok {
																{
																	x := []byte(v.(string))
																	var x1 interface{}
																	err := json.Unmarshal(x, &x1)
																	if err == nil && x1 != nil {
																		o.AdditionalProperties = x1.(map[string]interface{})
																	}
																}
															}
															if v, ok := l["body"]; ok {
																{
																	x := []byte(v.(string))
																	var x1 interface{}
																	err := json.Unmarshal(x, &x1)
																	if err == nil && x1 != nil {
																		x2 := x1.(map[string]interface{})
																		o.SetBody(x2)
																	}
																}
															}
															o.SetClassId("")
															if v, ok := l["method"]; ok {
																{
																	x := (v.(string))
																	o.SetMethod(x)
																}
															}
															if v, ok := l["object_type"]; ok {
																{
																	x := (v.(string))
																	o.SetObjectType(x)
																}
															}
															p = append(p, *o)
														}
														if len(p) > 0 {
															x := p[0]
															o.SetSelectorProperty(x)
														}
													}
												}
												if v, ok := l["value_attribute"]; ok {
													{
														x := (v.(string))
														o.SetValueAttribute(x)
													}
												}
												x = append(x, *o)
											}
											if len(x) > 0 {
												o.SetInventorySelector(x)
											}
										}
									}
									if v, ok := l["object_type"]; ok {
										{
											x := (v.(string))
											o.SetObjectType(x)
										}
									}
									if v, ok := l["secure"]; ok {
										{
											x := (v.(bool))
											o.SetSecure(x)
										}
									}
									if v, ok := l["type"]; ok {
										{
											x := (v.(string))
											o.SetType(x)
										}
									}
									p = append(p, *o)
								}
								if len(p) > 0 {
									x := p[0]
									o.SetProperties(x)
								}
							}
						}
						if v, ok := l["required"]; ok {
							{
								x := (v.(bool))
								o.SetRequired(x)
							}
						}
						p = append(p, *o)
					}
					if len(p) > 0 {
						x := p[0]
						o.SetType(x)
					}
				}
			}
			if v, ok := l["value"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						x2 := x1.(map[string]interface{})
						o.SetValue(x2)
					}
				}
			}
			x = append(x, *o)
		}
		o.SetPlaceholders(x)
	}

	if d.HasChange("tags") {
		v := d.Get("tags")
		x := make([]models.MoTag, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.MoTag{}
			l := s[i].(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			if v, ok := l["key"]; ok {
				{
					x := (v.(string))
					o.SetKey(x)
				}
			}
			if v, ok := l["value"]; ok {
				{
					x := (v.(string))
					o.SetValue(x)
				}
			}
			x = append(x, *o)
		}
		o.SetTags(x)
	}

	r := conn.ApiClient.OsApi.UpdateOsConfigurationFile(conn.ctx, d.Id()).OsConfigurationFile(*o)
	result, _, responseErr := r.Execute()
	if responseErr != nil {
		errorType := fmt.Sprintf("%T", responseErr)
		if strings.Contains(errorType, "GenericOpenAPIError") {
			responseErr := responseErr.(*models.GenericOpenAPIError)
			return diag.Errorf("error occurred while updating OsConfigurationFile: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		return diag.Errorf("error occurred while updating OsConfigurationFile: %s", responseErr.Error())
	}
	log.Printf("Moid: %s", result.GetMoid())
	d.SetId(result.GetMoid())
	return append(de, resourceOsConfigurationFileRead(c, d, meta)...)
}

func resourceOsConfigurationFileDelete(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	var de diag.Diagnostics
	conn := meta.(*Config)
	p := conn.ApiClient.OsApi.DeleteOsConfigurationFile(conn.ctx, d.Id())
	_, deleteErr := p.Execute()
	if deleteErr != nil {
		errorType := fmt.Sprintf("%T", deleteErr)
		if strings.Contains(deleteErr.Error(), "404") {
			de = append(de, diag.Diagnostic{Summary: "OsConfigurationFileDelete: OsConfigurationFile object " + d.Id() + " not found. Removing from statefile", Severity: diag.Warning})
			return de
		}
		if strings.Contains(errorType, "GenericOpenAPIError") {
			deleteErr := deleteErr.(*models.GenericOpenAPIError)
			return diag.Errorf("error occurred while deleting OsConfigurationFile object: %s Response from endpoint: %s", deleteErr.Error(), string(deleteErr.Body()))
		}
		return diag.Errorf("error occurred while deleting OsConfigurationFile object: %s", deleteErr.Error())
	}
	return de
}
