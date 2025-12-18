import json
from django.shortcuts import render
from django.utils.safestring import mark_safe


def render_vue(request, component_path, props=None):
    if props is None:
        props = {}

    context = {
        "component_path": component_path,
        "props": mark_safe(json.dumps(props)),
    }

    return render(request, "app.html", context)
