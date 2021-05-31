_JINJA2_FILETYPE = [
    ".jinja2",
    ".jinja",
]

def _add_prefix_to_imports(label, imports):
    imports_prefix = ""
    if label.workspace_root:
        imports_prefix += label.workspace_root + "/"
    if label.package:
        imports_prefix += label.package + "/"
    return [imports_prefix + im for im in imports]

def _setup_deps(deps):
    """Collects source files and import flags of transitive dependencies.

    Args:
      deps: List of deps labels from ctx.attr.deps.

    Returns:
      Returns a struct containing the following fields:
        transitive_sources: List of Files containing sources of transitive
            dependencies
        imports: List of Strings containing import flags set by transitive
            dependency targets.
    """
    transitive_sources = []
    imports = []
    for dep in deps:
        transitive_sources.append(dep.transitive_jinja2_files)
        imports.append(dep.imports)

    return struct(
        imports = depset(transitive = imports),
        transitive_sources = depset(transitive = transitive_sources, order = "postorder"),
    )

def _toolchain(ctx):
    return struct(
        binary_path = ctx.executable.jinja2.path,
    )

def _quote(s):
    return '"' + s.replace('"', '\\"') + '"'

def _stamp_resolve(ctx, string, output):
    stamps = [ctx.info_file, ctx.version_file]
    stamp_args = [
        "--stamp-info-file=%s" % sf.path
        for sf in stamps
    ]
    ctx.actions.run(
        executable = ctx.executable._stamper,
        arguments = [
            "--format=%s" % string,
            "--output=%s" % output.path,
        ] + stamp_args,
        inputs = stamps,
        tools = [ctx.executable._stamper],
        outputs = [output],
        mnemonic = "Stamp",
    )

def _jinja2_generate_impl(ctx):
    """Implementation of the jinja2_generate rule."""

    depinfo = _setup_deps(ctx.attr.deps)
    toolchain = _toolchain(ctx)
    # jinja2_ext_str_files = ctx.files.ext_str_files

    command = (
        [
            "set -e;",
        ] + ["%s %s %s.yaml" % (toolchain.binary_path, im, im) for im in _add_prefix_to_imports(ctx.label, ctx.attr.imports)] +
    )

    outputs = []

		outputs += ctx.outputs.outs
		# command += ["-m", ctx.outputs.outs[0].dirname, ctx.file.src.path]

    transitive_data = depset(transitive = [dep.data_runfiles.files for dep in ctx.attr.deps]
    # NB(sparkprime): (1) transitive_data is never used, since runfiles is only
    # used when .files is pulled from it.  (2) This makes sense - jinja2 does
    # not need transitive dependencies to be passed on the commandline. It
    # needs the -J but that is handled separately.

    # files = jinja2_ext_str_files

    runfiles = ctx.runfiles(
        collect_data = True,
        files = files,
        transitive_files = depset(transitive_data),
    )

    compile_inputs = (
        [ctx.file.src] +
        runfiles.files.to_list() +
        depinfo.transitive_sources.to_list()
    )

    tools = [ctx.executable.jinja2]

    ctx.actions.run_shell(
        inputs = compile_inputs,
        tools = tools,
        outputs = outputs,
        mnemonic = "Jinja2",
        command = " ".join(command),
        use_default_shell_env = True,
        progress_message = "Generating Jinja2 template for " + ctx.label.name,
    )

_jinja2_common_attrs = {
    "data": attr.label_list(
        allow_files = True,
    ),
    "imports": attr.string_list(),
    "jinja2": attr.label(
        default = Label("//messaging/alerting_rules:jinja2"),
        cfg = "host",
        executable = True,
        allow_single_file = True,
    ),
    "deps": attr.label_list(
        providers = ["transitive_jinja2_files"],
        allow_files = False,
    ),
}

"""Creates a logical set of Jinja2 files.

Args:
    name: A unique name for this rule.
    srcs: List of `.jinja2` files that comprises this Jinja2 library.
    deps: List of targets that are required by the `srcs` Jinja2 files.
    imports: List of import `-J` flags to be passed to the `jinja2` compiler.

Example:
  Suppose you have the following directory structure:

  ```
  [workspace]/
      WORKSPACE
      configs/
          BUILD
          backend.jinja2
          frontend.jinja2
  ```

  You can use the `jinja2_library` rule to build a collection of `.jinja2`
  files that can be imported by other `.jinja2` files as dependencies:

  `configs/BUILD`:

  ```python
  load("@io_bazel_rules_jinja2//jinja2:jinja2.bzl", "jinja2_library")

  jinja2_library(
      name = "configs",
      srcs = [
          "backend.jinja2",
          "frontend.jinja2",
      ],
  )
  ```
"""

_jinja2_compile_attrs = {
    "srcs": attr.label(allow_files = True),
    # "ext_str_files": attr.label_list(
    #     allow_files = True,
    # ),
    "extra_args": attr.string_list(),
    # "_stamper": attr.label(
    #     default = Label("//jinja2:stamper"),
    #     cfg = "host",
    #     executable = True,
    #     allow_files = True,
    # ),
}

_jinja2_generate_attrs = {
    "outs": attr.output_list(mandatory = True),
}

jinja2_generate = rule(
    _jinja2_impl,
    attrs = dict(_jinja2_compile_attrs.items() +
                 _jinja2_generate_attrs.items() +
                 _jinja2_common_attrs.items()),
)

