# Python Bazel demo


## Usage:

1. Run main function: `bazel run //jupyter-run`

2. To run Python3: 

    1. Firstly, you need to locate Python3 using: `bazel cquery 'deps(//jupyter-run)' | fgrep bin/python3`

    1. In my desktop(Linux), it is:

    ```
    Analyzing: target //jupyter-run:jupyter-run (0 packages loaded, 0 targets configured)
    INFO: Analyzed target //jupyter-run:jupyter-run (0 packages loaded, 0 targets configured).
    INFO: Found 1 target...
    @python3_11_x86_64-unknown-linux-gnu//:bin/python3 (null)
    @python3_11_x86_64-unknown-linux-gnu//:bin/python3-config (null)
    @python3_11_x86_64-unknown-linux-gnu//:bin/python3.11 (null)
    @python3_11_x86_64-unknown-linux-gnu//:bin/python3.11-config (null)
    INFO: Elapsed time: 0.147s, Critical Path: 0.00s
    INFO: 0 processes.
    INFO: Build completed successfully, 0 total actions

    ```
    Then I can run Python using: `bazel run @python3_11_x86_64-unknown-linux-gnu//:bin/python3`
