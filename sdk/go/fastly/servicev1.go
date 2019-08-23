// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fastly

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a Fastly Service, representing the configuration for a website, app,
// API, or anything else to be served through Fastly. A Service encompasses Domains
// and Backends.
// 
// The Service resource requires a domain name that is correctly set up to direct
// traffic to the Fastly service. See Fastly's guide on [Adding CNAME Records][fastly-cname]
// on their documentation site for guidance.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-fastly/blob/master/website/docs/r/service_v1.html.markdown.
type Servicev1 struct {
	s *pulumi.ResourceState
}

// NewServicev1 registers a new resource with the given unique name, arguments, and options.
func NewServicev1(ctx *pulumi.Context,
	name string, args *Servicev1Args, opts ...pulumi.ResourceOpt) (*Servicev1, error) {
	if args == nil || args.Domains == nil {
		return nil, errors.New("missing required argument 'Domains'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["acls"] = nil
		inputs["activate"] = nil
		inputs["backends"] = nil
		inputs["bigqueryloggings"] = nil
		inputs["blobstorageloggings"] = nil
		inputs["cacheSettings"] = nil
		inputs["comment"] = nil
		inputs["conditions"] = nil
		inputs["defaultHost"] = nil
		inputs["defaultTtl"] = nil
		inputs["dictionaries"] = nil
		inputs["directors"] = nil
		inputs["domains"] = nil
		inputs["dynamicsnippets"] = nil
		inputs["forceDestroy"] = nil
		inputs["gcsloggings"] = nil
		inputs["gzips"] = nil
		inputs["headers"] = nil
		inputs["healthchecks"] = nil
		inputs["logentries"] = nil
		inputs["name"] = nil
		inputs["papertrails"] = nil
		inputs["requestSettings"] = nil
		inputs["responseObjects"] = nil
		inputs["s3loggings"] = nil
		inputs["snippets"] = nil
		inputs["splunks"] = nil
		inputs["sumologics"] = nil
		inputs["syslogs"] = nil
		inputs["vcls"] = nil
		inputs["versionComment"] = nil
	} else {
		inputs["acls"] = args.Acls
		inputs["activate"] = args.Activate
		inputs["backends"] = args.Backends
		inputs["bigqueryloggings"] = args.Bigqueryloggings
		inputs["blobstorageloggings"] = args.Blobstorageloggings
		inputs["cacheSettings"] = args.CacheSettings
		inputs["comment"] = args.Comment
		inputs["conditions"] = args.Conditions
		inputs["defaultHost"] = args.DefaultHost
		inputs["defaultTtl"] = args.DefaultTtl
		inputs["dictionaries"] = args.Dictionaries
		inputs["directors"] = args.Directors
		inputs["domains"] = args.Domains
		inputs["dynamicsnippets"] = args.Dynamicsnippets
		inputs["forceDestroy"] = args.ForceDestroy
		inputs["gcsloggings"] = args.Gcsloggings
		inputs["gzips"] = args.Gzips
		inputs["headers"] = args.Headers
		inputs["healthchecks"] = args.Healthchecks
		inputs["logentries"] = args.Logentries
		inputs["name"] = args.Name
		inputs["papertrails"] = args.Papertrails
		inputs["requestSettings"] = args.RequestSettings
		inputs["responseObjects"] = args.ResponseObjects
		inputs["s3loggings"] = args.S3loggings
		inputs["snippets"] = args.Snippets
		inputs["splunks"] = args.Splunks
		inputs["sumologics"] = args.Sumologics
		inputs["syslogs"] = args.Syslogs
		inputs["vcls"] = args.Vcls
		inputs["versionComment"] = args.VersionComment
	}
	inputs["activeVersion"] = nil
	s, err := ctx.RegisterResource("fastly:index/servicev1:Servicev1", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Servicev1{s: s}, nil
}

// GetServicev1 gets an existing Servicev1 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServicev1(ctx *pulumi.Context,
	name string, id pulumi.ID, state *Servicev1State, opts ...pulumi.ResourceOpt) (*Servicev1, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["acls"] = state.Acls
		inputs["activate"] = state.Activate
		inputs["activeVersion"] = state.ActiveVersion
		inputs["backends"] = state.Backends
		inputs["bigqueryloggings"] = state.Bigqueryloggings
		inputs["blobstorageloggings"] = state.Blobstorageloggings
		inputs["cacheSettings"] = state.CacheSettings
		inputs["comment"] = state.Comment
		inputs["conditions"] = state.Conditions
		inputs["defaultHost"] = state.DefaultHost
		inputs["defaultTtl"] = state.DefaultTtl
		inputs["dictionaries"] = state.Dictionaries
		inputs["directors"] = state.Directors
		inputs["domains"] = state.Domains
		inputs["dynamicsnippets"] = state.Dynamicsnippets
		inputs["forceDestroy"] = state.ForceDestroy
		inputs["gcsloggings"] = state.Gcsloggings
		inputs["gzips"] = state.Gzips
		inputs["headers"] = state.Headers
		inputs["healthchecks"] = state.Healthchecks
		inputs["logentries"] = state.Logentries
		inputs["name"] = state.Name
		inputs["papertrails"] = state.Papertrails
		inputs["requestSettings"] = state.RequestSettings
		inputs["responseObjects"] = state.ResponseObjects
		inputs["s3loggings"] = state.S3loggings
		inputs["snippets"] = state.Snippets
		inputs["splunks"] = state.Splunks
		inputs["sumologics"] = state.Sumologics
		inputs["syslogs"] = state.Syslogs
		inputs["vcls"] = state.Vcls
		inputs["versionComment"] = state.VersionComment
	}
	s, err := ctx.ReadResource("fastly:index/servicev1:Servicev1", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Servicev1{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Servicev1) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Servicev1) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// A set of ACL configuration blocks.  Defined below.
func (r *Servicev1) Acls() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["acls"])
}

// Conditionally prevents the Service from being activated. The apply step will continue to create a new draft version but will not activate it if this is set to false. Default true.
func (r *Servicev1) Activate() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["activate"])
}

// The currently active version of your Fastly Service.
func (r *Servicev1) ActiveVersion() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["activeVersion"])
}

// A set of Backends to service requests from your Domains.
// Defined below. Backends must be defined in this argument, or defined in the
// `vcl` argument below
func (r *Servicev1) Backends() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["backends"])
}

// A BigQuery endpoint to send streaming logs too.
// Defined below.
func (r *Servicev1) Bigqueryloggings() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["bigqueryloggings"])
}

// An Azure Blob Storage endpoint to send streaming logs too.
// Defined below.
func (r *Servicev1) Blobstorageloggings() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["blobstorageloggings"])
}

// A set of Cache Settings, allowing you to override
func (r *Servicev1) CacheSettings() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["cacheSettings"])
}

// A personal freeform descriptive note
func (r *Servicev1) Comment() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["comment"])
}

// A set of conditions to add logic to any basic
// configuration object in this service. Defined below.
func (r *Servicev1) Conditions() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["conditions"])
}

// The default hostname.
func (r *Servicev1) DefaultHost() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["defaultHost"])
}

// The default Time-to-live (TTL) for
// requests.
func (r *Servicev1) DefaultTtl() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["defaultTtl"])
}

// A set of dictionaries that allow the storing of key values pair for use within VCL functions. Defined below.
func (r *Servicev1) Dictionaries() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["dictionaries"])
}

// A director to allow more control over balancing traffic over backends.
// when an item is not to be cached based on an above `condition`. Defined below
func (r *Servicev1) Directors() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["directors"])
}

// A set of Domain names to serve as entry points for your
// Service. Defined below.
func (r *Servicev1) Domains() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["domains"])
}

// A set of custom, "dynamic" VCL Snippet configuration blocks.  Defined below.
func (r *Servicev1) Dynamicsnippets() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["dynamicsnippets"])
}

// Services that are active cannot be destroyed. In
// order to destroy the Service, set `forceDestroy` to `true`. Default `false`.
func (r *Servicev1) ForceDestroy() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["forceDestroy"])
}

// A gcs endpoint to send streaming logs too.
// Defined below.
func (r *Servicev1) Gcsloggings() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["gcsloggings"])
}

// A set of gzip rules to control automatic gzipping of
// content. Defined below.
func (r *Servicev1) Gzips() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["gzips"])
}

// A set of Headers to manipulate for each request. Defined
// below.
func (r *Servicev1) Headers() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["headers"])
}

// Automated healthchecks on the cache that can change how Fastly interacts with the cache based on its health.
func (r *Servicev1) Healthchecks() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["healthchecks"])
}

// A logentries endpoint to send streaming logs too.
// Defined below.
func (r *Servicev1) Logentries() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["logentries"])
}

// The unique name for the Service to create.
func (r *Servicev1) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// A Papertrail endpoint to send streaming logs too.
// Defined below.
func (r *Servicev1) Papertrails() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["papertrails"])
}

// A set of Request modifiers. Defined below
func (r *Servicev1) RequestSettings() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["requestSettings"])
}

// Allows you to create synthetic responses that exist entirely on the varnish machine. Useful for creating error or maintenance pages that exists outside the scope of your datacenter. Best when used with Condition objects.
func (r *Servicev1) ResponseObjects() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["responseObjects"])
}

// A set of S3 Buckets to send streaming logs too.
// Defined below.
func (r *Servicev1) S3loggings() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["s3loggings"])
}

// A set of custom, "regular" (non-dynamic) VCL Snippet configuration blocks.  Defined below.
func (r *Servicev1) Snippets() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["snippets"])
}

// A Splunk endpoint to send streaming logs too.
// Defined below.
func (r *Servicev1) Splunks() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["splunks"])
}

// A Sumologic endpoint to send streaming logs too.
// Defined below.
func (r *Servicev1) Sumologics() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["sumologics"])
}

// A syslog endpoint to send streaming logs too.
// Defined below.
func (r *Servicev1) Syslogs() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["syslogs"])
}

// A set of custom VCL configuration blocks. The
// ability to upload custom VCL code is not enabled by default for new Fastly
// accounts (see the [Fastly documentation](https://docs.fastly.com/guides/vcl/uploading-custom-vcl) for details).
func (r *Servicev1) Vcls() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["vcls"])
}

// Description field for the version.
func (r *Servicev1) VersionComment() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["versionComment"])
}

// Input properties used for looking up and filtering Servicev1 resources.
type Servicev1State struct {
	// A set of ACL configuration blocks.  Defined below.
	Acls interface{}
	// Conditionally prevents the Service from being activated. The apply step will continue to create a new draft version but will not activate it if this is set to false. Default true.
	Activate interface{}
	// The currently active version of your Fastly Service.
	ActiveVersion interface{}
	// A set of Backends to service requests from your Domains.
	// Defined below. Backends must be defined in this argument, or defined in the
	// `vcl` argument below
	Backends interface{}
	// A BigQuery endpoint to send streaming logs too.
	// Defined below.
	Bigqueryloggings interface{}
	// An Azure Blob Storage endpoint to send streaming logs too.
	// Defined below.
	Blobstorageloggings interface{}
	// A set of Cache Settings, allowing you to override
	CacheSettings interface{}
	// A personal freeform descriptive note
	Comment interface{}
	// A set of conditions to add logic to any basic
	// configuration object in this service. Defined below.
	Conditions interface{}
	// The default hostname.
	DefaultHost interface{}
	// The default Time-to-live (TTL) for
	// requests.
	DefaultTtl interface{}
	// A set of dictionaries that allow the storing of key values pair for use within VCL functions. Defined below.
	Dictionaries interface{}
	// A director to allow more control over balancing traffic over backends.
	// when an item is not to be cached based on an above `condition`. Defined below
	Directors interface{}
	// A set of Domain names to serve as entry points for your
	// Service. Defined below.
	Domains interface{}
	// A set of custom, "dynamic" VCL Snippet configuration blocks.  Defined below.
	Dynamicsnippets interface{}
	// Services that are active cannot be destroyed. In
	// order to destroy the Service, set `forceDestroy` to `true`. Default `false`.
	ForceDestroy interface{}
	// A gcs endpoint to send streaming logs too.
	// Defined below.
	Gcsloggings interface{}
	// A set of gzip rules to control automatic gzipping of
	// content. Defined below.
	Gzips interface{}
	// A set of Headers to manipulate for each request. Defined
	// below.
	Headers interface{}
	// Automated healthchecks on the cache that can change how Fastly interacts with the cache based on its health.
	Healthchecks interface{}
	// A logentries endpoint to send streaming logs too.
	// Defined below.
	Logentries interface{}
	// The unique name for the Service to create.
	Name interface{}
	// A Papertrail endpoint to send streaming logs too.
	// Defined below.
	Papertrails interface{}
	// A set of Request modifiers. Defined below
	RequestSettings interface{}
	// Allows you to create synthetic responses that exist entirely on the varnish machine. Useful for creating error or maintenance pages that exists outside the scope of your datacenter. Best when used with Condition objects.
	ResponseObjects interface{}
	// A set of S3 Buckets to send streaming logs too.
	// Defined below.
	S3loggings interface{}
	// A set of custom, "regular" (non-dynamic) VCL Snippet configuration blocks.  Defined below.
	Snippets interface{}
	// A Splunk endpoint to send streaming logs too.
	// Defined below.
	Splunks interface{}
	// A Sumologic endpoint to send streaming logs too.
	// Defined below.
	Sumologics interface{}
	// A syslog endpoint to send streaming logs too.
	// Defined below.
	Syslogs interface{}
	// A set of custom VCL configuration blocks. The
	// ability to upload custom VCL code is not enabled by default for new Fastly
	// accounts (see the [Fastly documentation](https://docs.fastly.com/guides/vcl/uploading-custom-vcl) for details).
	Vcls interface{}
	// Description field for the version.
	VersionComment interface{}
}

// The set of arguments for constructing a Servicev1 resource.
type Servicev1Args struct {
	// A set of ACL configuration blocks.  Defined below.
	Acls interface{}
	// Conditionally prevents the Service from being activated. The apply step will continue to create a new draft version but will not activate it if this is set to false. Default true.
	Activate interface{}
	// A set of Backends to service requests from your Domains.
	// Defined below. Backends must be defined in this argument, or defined in the
	// `vcl` argument below
	Backends interface{}
	// A BigQuery endpoint to send streaming logs too.
	// Defined below.
	Bigqueryloggings interface{}
	// An Azure Blob Storage endpoint to send streaming logs too.
	// Defined below.
	Blobstorageloggings interface{}
	// A set of Cache Settings, allowing you to override
	CacheSettings interface{}
	// A personal freeform descriptive note
	Comment interface{}
	// A set of conditions to add logic to any basic
	// configuration object in this service. Defined below.
	Conditions interface{}
	// The default hostname.
	DefaultHost interface{}
	// The default Time-to-live (TTL) for
	// requests.
	DefaultTtl interface{}
	// A set of dictionaries that allow the storing of key values pair for use within VCL functions. Defined below.
	Dictionaries interface{}
	// A director to allow more control over balancing traffic over backends.
	// when an item is not to be cached based on an above `condition`. Defined below
	Directors interface{}
	// A set of Domain names to serve as entry points for your
	// Service. Defined below.
	Domains interface{}
	// A set of custom, "dynamic" VCL Snippet configuration blocks.  Defined below.
	Dynamicsnippets interface{}
	// Services that are active cannot be destroyed. In
	// order to destroy the Service, set `forceDestroy` to `true`. Default `false`.
	ForceDestroy interface{}
	// A gcs endpoint to send streaming logs too.
	// Defined below.
	Gcsloggings interface{}
	// A set of gzip rules to control automatic gzipping of
	// content. Defined below.
	Gzips interface{}
	// A set of Headers to manipulate for each request. Defined
	// below.
	Headers interface{}
	// Automated healthchecks on the cache that can change how Fastly interacts with the cache based on its health.
	Healthchecks interface{}
	// A logentries endpoint to send streaming logs too.
	// Defined below.
	Logentries interface{}
	// The unique name for the Service to create.
	Name interface{}
	// A Papertrail endpoint to send streaming logs too.
	// Defined below.
	Papertrails interface{}
	// A set of Request modifiers. Defined below
	RequestSettings interface{}
	// Allows you to create synthetic responses that exist entirely on the varnish machine. Useful for creating error or maintenance pages that exists outside the scope of your datacenter. Best when used with Condition objects.
	ResponseObjects interface{}
	// A set of S3 Buckets to send streaming logs too.
	// Defined below.
	S3loggings interface{}
	// A set of custom, "regular" (non-dynamic) VCL Snippet configuration blocks.  Defined below.
	Snippets interface{}
	// A Splunk endpoint to send streaming logs too.
	// Defined below.
	Splunks interface{}
	// A Sumologic endpoint to send streaming logs too.
	// Defined below.
	Sumologics interface{}
	// A syslog endpoint to send streaming logs too.
	// Defined below.
	Syslogs interface{}
	// A set of custom VCL configuration blocks. The
	// ability to upload custom VCL code is not enabled by default for new Fastly
	// accounts (see the [Fastly documentation](https://docs.fastly.com/guides/vcl/uploading-custom-vcl) for details).
	Vcls interface{}
	// Description field for the version.
	VersionComment interface{}
}