# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import importlib
# Make subpackages available:
__all__ = ['config']
for pkg in __all__:
    if pkg != 'config':
        importlib.import_module(f'{__name__}.{pkg}')

# Export this package's modules as members:
from .service_acl_entriesv1 import *
from .service_dictionary_itemsv1 import *
from .service_dynamic_snippet_contentv1 import *
from .servicev1 import *
from .get_fastly_ip_ranges import *
from .provider import *
