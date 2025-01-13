BAZEL = bazelisk

run:
	$(BAZEL) run //jupyter-run:jupyter -- --notebook-dir=$(PWD)/notebooks --ip=0.0.0.0

gazelle:
	$(BAZEL) run //:gazelle
