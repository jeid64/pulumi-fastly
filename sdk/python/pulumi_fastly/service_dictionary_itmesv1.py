# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables

class ServiceDictionaryItmesv1(pulumi.CustomResource):
    dictionary_id: pulumi.Output[str]
    """
    The ID of the dictionary that the items belong to
    """
    items: pulumi.Output[dict]
    """
    A map representing an entry in the dictionary, (key/value)
    """
    service_id: pulumi.Output[str]
    """
    The ID of the service that the dictionary belongs to
    """
    def __init__(__self__, resource_name, opts=None, dictionary_id=None, items=None, service_id=None, __props__=None, __name__=None, __opts__=None):
        """
        Create a ServiceDictionaryItmesv1 resource with the given unique name, props, and options.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] dictionary_id: The ID of the dictionary that the items belong to
        :param pulumi.Input[dict] items: A map representing an entry in the dictionary, (key/value)
        :param pulumi.Input[str] service_id: The ID of the service that the dictionary belongs to

        > This content is derived from https://github.com/terraform-providers/terraform-provider-fastly/blob/master/website/docs/r/service_dictionary_items_v1.html.markdown.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if dictionary_id is None:
                raise TypeError("Missing required property 'dictionary_id'")
            __props__['dictionary_id'] = dictionary_id
            __props__['items'] = items
            if service_id is None:
                raise TypeError("Missing required property 'service_id'")
            __props__['service_id'] = service_id
        super(ServiceDictionaryItmesv1, __self__).__init__(
            'fastly:index/serviceDictionaryItmesv1:ServiceDictionaryItmesv1',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, dictionary_id=None, items=None, service_id=None):
        """
        Get an existing ServiceDictionaryItmesv1 resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] dictionary_id: The ID of the dictionary that the items belong to
        :param pulumi.Input[dict] items: A map representing an entry in the dictionary, (key/value)
        :param pulumi.Input[str] service_id: The ID of the service that the dictionary belongs to

        > This content is derived from https://github.com/terraform-providers/terraform-provider-fastly/blob/master/website/docs/r/service_dictionary_items_v1.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["dictionary_id"] = dictionary_id
        __props__["items"] = items
        __props__["service_id"] = service_id
        return ServiceDictionaryItmesv1(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

