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
