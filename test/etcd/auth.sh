#!/bin/sh

etcdctl put "/GRAPE/dev/AUTH/APP/app1" \
'{"id": 1, "app_id": "app1", "app_sec": "app_sec1", "username": "app1", "endpoints": ["GET::/ping", "GET::/echo"]}'

etcdctl put "/GRAPE/dev/AUTH/APP/app2" \
'{"id": 2, "app_id": "app2", "app_sec": "app_sec2", "username": "app2", "endpoints": ["GET::/fail", "GET::/timeout"]}'

etcdctl put "/GRAPE/dev/AUTH/APP/app3" \
'{"id": 3, "app_id": "app3", "app_sec": "app_sec3", "username": "app3", "endpoints": ["GET::/health", "GET::/echo"]}'


etcdctl put "/GRAPE/dev/AUTH/TOKEN/t1" \
'{"id": "tttttt-ttttt-1", "app_id": 1}'

etcdctl put "/GRAPE/dev/AUTH/TOKEN/t2" \
'{"id": "tttttt-ttttt-2", "app_id": 2}'

etcdctl put "/GRAPE/dev/AUTH/TOKEN/t3" \
'{"id": "tttttt-ttttt-3", "app_id": 3, "expired_at": 66666}'

etcdctl put "/GRAPE/dev/AUTH/TOKEN/t31" \
'{"id": "tttttt-ttttt-31", "app_id": 3, "expired_at": 1920816239}'