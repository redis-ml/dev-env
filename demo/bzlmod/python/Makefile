BAZEL = bazelisk

python:
	echo "Check comments for different OS"
	# For Linux
	# bazel run '@@rules_python~0.25.0~python~python_3_11_x86_64-unknown-linux-gnu//:bin/python3
	# For MacOS
	# bazel run '@@rules_python~0.25.0~python~python_3_11_x86_64-apple-darwin//:bin/python3' --
	#
	# OR run 'make locate_python' to find out
locate_python:
	$(BAZEL) cquery 'deps(:bzlmod)' | fgrep bin/python
	
