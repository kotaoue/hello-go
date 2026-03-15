# go-exportedtest
Test for effect on build by "exported names".

## Result
```ShellSession
$ ll | grep -v ".go"
total 15624
drwxr-xr-x  18 kota.oue  staff      576  6 26 17:23 .
drwxr-xr-x   8 kota.oue  staff      256  6 26 16:49 ..
drwxr-xr-x  14 kota.oue  staff      448  6 26 16:54 .git
-rw-r--r--   1 kota.oue  staff      269  6 26 16:49 .gitignore
-rw-r--r--   1 kota.oue  staff     1065  6 26 16:49 LICENSE
-rw-r--r--   1 kota.oue  staff       64  6 26 16:49 README.md
-rwxr-xr-x   1 kota.oue  staff  1200912  6 26 17:23 bigimport_only
-rwxr-xr-x   1 kota.oue  staff  2174008  6 26 16:56 import
-rwxr-xr-x   1 kota.oue  staff  1200912  6 26 16:55 noimport
drwxr-xr-x   4 kota.oue  staff      128  6 26 17:21 packages
-rwxr-xr-x   1 kota.oue  staff  2174008  6 26 17:10 smallimport
-rwxr-xr-x   1 kota.oue  staff  1200912  6 26 17:11 smallimport_only
```

## Conclusion
There was no big difference in file size.