# cudnn_log_parser

Example input = [alexnet_cudnn.log](https://github.com/rai-project/cudnn_log_parser/blob/master/_fixtures/alexnet_cudnn.log)

Example output = [alexnet_cudnn.csv](https://github.com/rai-project/cudnn_log_parser/blob/master/_fixtures/alexnet_cudnn.csv)



### Run with cublas logging

```
CUBLAS_LOGINFO_DBG=1 CUBLAS_LOGDEST_DBG=cublas.log prog
```



### Run with cudnn logging

```
CUDNN_LOGINFO_DBG=1 CUDNN_LOGDEST_DBG=cudnn.log prog
```


### Run with full logging

```
$ CUBLAS_LOGINFO_DBG=1 CUBLAS_LOGDEST_DBG=cublas.log CUDNN_LOGINFO_DBG=1 CUDNN_LOGDEST_DBG=cudnn.log prog
```


### MKL

https://github.com/intel/mkl-dnn/blob/master/doc/performance_considerations/verbose.md


### CUDA Driver

https://docs.nvidia.com/cuda/cuda-c-programming-guide/index.html#env-vars

Driver-Based Profiler (these variables have no impact on the Visual Profiler or the command line profiler nvprof)

- `CUDA_​DEVICE`	Integer (default is 0)	Specifies the index of the device to profile.
- `COMPUTE_​PROFILE	0` or 1 (default is 0)	Disables profiling (when set to 0) or enables profiling (when set to 1).
- `COMPUTE_​PROFILE_​CONFIG`	Path	Specifies the configuration file to set profiling options and select performance counters.
- `COMPUTE_​PROFILE_​LOG`	Path	Specifies the file used to save the profiling output. In case of multiple contexts, use '%d' in the COMPUTE_PROFILE_LOG to generate separate output files for each context - with '%d' substituted by the context number.
- `COMPUTE_​PROFILE_​CSV`	0 or 1 (default is 0)	When set to 1, the output will be in comma-separated format.
