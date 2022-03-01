JinjaFiles = provider("transitive_sources")

def get_transitive_srcs(srcs, deps):
    """Obtain the source files for a target and its transitive dependencies.
    Args:
      srcs: a list of source files
      deps: a list of targets that are direct dependencies
    Returns:
      a collection of the transitive sources
    """
    return depset(
        srcs,
        transitive = [dep[JinjaFiles].transitive_sources for dep in deps],
    )

def _jinja_library_impl(ctx):
    trans_srcs = get_transitive_srcs(ctx.files.srcs, ctx.attr.deps)
    return [
        JinjaFiles(transitive_sources = trans_srcs),
        DefaultInfo(files = trans_srcs),
    ]

jinja_library = rule(
    implementation = _jinja_library_impl,
    attrs = {
        "srcs": attr.label_list(allow_files = True),
        "deps": attr.label_list(),
    },
)

def _jinja_binary_impl(ctx):
    jinjacc = ctx.executable._jinjacc
    out = ctx.outputs.out
    trans_srcs = get_transitive_srcs(ctx.files.srcs, ctx.attr.deps)
    srcs_list = trans_srcs.to_list()
    ctx.actions.run(
        executable = jinjacc,
        arguments = [out.path] + [src.path for src in srcs_list],
        inputs = srcs_list,
        tools = [jinjacc],
        outputs = [out],
    )

jinja_binary = rule(
    implementation = _jinja_binary_impl,
    attrs = {
        "srcs": attr.label_list(allow_files = True),
        "deps": attr.label_list(),
        "_jinjacc": attr.label(
            default = Label("//depsets:jinjacc"),
            allow_files = True,
            executable = True,
            cfg = "exec",
        ),
    },
    outputs = {"out": "%{name}.out"},
)
