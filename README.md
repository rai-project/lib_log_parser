# cudnn_log_parser



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
