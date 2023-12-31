// go run linux/mksysnum.go -Wall -Werror -static -I/tmp/include /tmp/include/asm/unistd.h
// Code generated by the command above; see README.md. DO NOT EDIT.

//go:build sw64 && linux
// +build sw64,linux

package unix

const (
	SYS_EXIT                   = 1
	SYS_FORK                   = 2
	SYS_READ                   = 3
	SYS_WRITE                  = 4
	SYS_CLOSE                  = 6
	SYS_LINK                   = 9
	SYS_UNLINK                 = 10
	SYS_CHDIR                  = 12
	SYS_FCHDIR                 = 13
	SYS_MKNOD                  = 14
	SYS_CHMOD                  = 15
	SYS_CHOWN                  = 16
	SYS_BRK                    = 17
	SYS_LSEEK                  = 19
	SYS_GETXPID                = 20
	SYS_UMOUNT2                = 22
	SYS_SETUID                 = 23
	SYS_GETXUID                = 24
	SYS_PTRACE                 = 26
	SYS_ACCESS                 = 33
	SYS_SYNC                   = 36
	SYS_KILL                   = 37
	SYS_SETPGID                = 39
	SYS_DUP                    = 41
	SYS_PIPE                   = 42
	SYS_OPEN                   = 45
	SYS_GETXGID                = 47
	SYS_ODD_SIGPROCMASK        = 48
	SYS_ACCT                   = 51
	SYS_SIGPENDING             = 52
	SYS_IOCTL                  = 54
	SYS_SYMLINK                = 57
	SYS_READLINK               = 58
	SYS_EXECVE                 = 59
	SYS_UMASK                  = 60
	SYS_CHROOT                 = 61
	SYS_GETPGRP                = 63
	SYS_VFORK                  = 66
	SYS_STAT                   = 67
	SYS_LSTAT                  = 68
	SYS_MMAP                   = 71
	SYS_MUNMAP                 = 73
	SYS_MPROTECT               = 74
	SYS_MADVISE                = 75
	SYS_VHANGUP                = 76
	SYS_GETGROUPS              = 79
	SYS_SETGROUPS              = 80
	SYS_SETPGRP                = 82
	SYS_GETHOSTNAME            = 87
	SYS_SETHOSTNAME            = 88
	SYS_DUP2                   = 90
	SYS_FSTAT                  = 91
	SYS_FCNTL                  = 92
	SYS_POLL                   = 94
	SYS_FSYNC                  = 95
	SYS_SETPRIORITY            = 96
	SYS_SOCKET                 = 97
	SYS_CONNECT                = 98
	SYS_ACCEPT                 = 99
	SYS_ODD_GETPRIORITY        = 100
	SYS_SEND                   = 101
	SYS_RECV                   = 102
	SYS_SIGRETURN              = 103
	SYS_BIND                   = 104
	SYS_SETSOCKOPT             = 105
	SYS_LISTEN                 = 106
	SYS_SIGSUSPEND             = 111
	SYS_RECVMSG                = 113
	SYS_SENDMSG                = 114
	SYS_GETSOCKOPT             = 118
	SYS_SOCKETCALL             = 119
	SYS_READV                  = 120
	SYS_WRITEV                 = 121
	SYS_FCHOWN                 = 123
	SYS_FCHMOD                 = 124
	SYS_RECVFROM               = 125
	SYS_SETREUID               = 126
	SYS_SETREGID               = 127
	SYS_RENAME                 = 128
	SYS_TRUNCATE               = 129
	SYS_FTRUNCATE              = 130
	SYS_FLOCK                  = 131
	SYS_SETGID                 = 132
	SYS_SENDTO                 = 133
	SYS_SHUTDOWN               = 134
	SYS_SOCKETPAIR             = 135
	SYS_MKDIR                  = 136
	SYS_RMDIR                  = 137
	SYS_GETPEERNAME            = 141
	SYS_GETRLIMIT              = 144
	SYS_SETRLIMIT              = 145
	SYS_SETSID                 = 147
	SYS_QUOTACTL               = 148
	SYS_GETSOCKNAME            = 150
	SYS_SIGACTION              = 156
	SYS_SETDOMAINNAME          = 166
	SYS_BPF                    = 170
	SYS_USERFAULTFD            = 171
	SYS_MEMBARRIER             = 172
	SYS_MLOCK2                 = 173
	SYS_GETPID                 = 174
	SYS_GETPPID                = 175
	SYS_GETUID                 = 176
	SYS_GETEUID                = 177
	SYS_GETGID                 = 178
	SYS_GETEGID                = 179
	SYS_MSGCTL                 = 200
	SYS_MSGGET                 = 201
	SYS_MSGRCV                 = 202
	SYS_MSGSND                 = 203
	SYS_SEMCTL                 = 204
	SYS_SEMGET                 = 205
	SYS_SEMOP                  = 206
	SYS_LCHOWN                 = 208
	SYS_SHMAT                  = 209
	SYS_SHMCTL                 = 210
	SYS_SHMDT                  = 211
	SYS_SHMGET                 = 212
	SYS_MSYNC                  = 217
	SYS_STATFS64               = 229
	SYS_FSTATFS64              = 230
	SYS_GETPGID                = 233
	SYS_GETSID                 = 234
	SYS_SIGALTSTACK            = 235
	SYS_SYSFS                  = 254
	SYS_GETSYSINFO             = 256
	SYS_SETSYSINFO             = 257
	SYS_PIDFD_SEND_SIGNAL      = 271
	SYS_IO_URING_SETUP         = 272
	SYS_IO_URING_ENTER         = 273
	SYS_IO_URING_REGISTER      = 274
	SYS_OPEN_TREE              = 275
	SYS_MOVE_MOUNT             = 276
	SYS_FSOPEN                 = 277
	SYS_FSCONFIG               = 278
	SYS_FSMOUNT                = 279
	SYS_FSPICK                 = 280
	SYS_PIDFD_OPEN             = 281
	SYS_CLONE3                 = 282
	SYS_CLOSE_RANGE            = 283
	SYS_OPENAT2                = 284
	SYS_PIDFD_GETFD            = 285
	SYS_FACCESSAT2             = 286
	SYS_PROCESS_MADVISE        = 287
	SYS_PKEY_MPROTECT          = 288
	SYS_PKEY_ALLOC             = 289
	SYS_PKEY_FREE              = 290
	SYS_GETPRIORITY            = 298
	SYS_SIGPROCMASK            = 299
	SYS_BDFLUSH                = 300
	SYS_MOUNT                  = 302
	SYS_SWAPOFF                = 304
	SYS_GETDENTS               = 305
	SYS_CREATE_MODULE          = 306
	SYS_INIT_MODULE            = 307
	SYS_DELETE_MODULE          = 308
	SYS_GET_KERNEL_SYMS        = 309
	SYS_SYSLOG                 = 310
	SYS_REBOOT                 = 311
	SYS_CLONE                  = 312
	SYS_USELIB                 = 313
	SYS_MLOCK                  = 314
	SYS_MUNLOCK                = 315
	SYS_MLOCKALL               = 316
	SYS_MUNLOCKALL             = 317
	SYS_SYSINFO                = 318
	SYS_OLDUMOUNT              = 321
	SYS_SWAPON                 = 322
	SYS_TIMES                  = 323
	SYS_PERSONALITY            = 324
	SYS_SETFSUID               = 325
	SYS_SETFSGID               = 326
	SYS_USTAT                  = 327
	SYS_STATFS                 = 328
	SYS_FSTATFS                = 329
	SYS_SCHED_SETPARAM         = 330
	SYS_SCHED_GETPARAM         = 331
	SYS_SCHED_SETSCHEDULER     = 332
	SYS_SCHED_GETSCHEDULER     = 333
	SYS_SCHED_YIELD            = 334
	SYS_SCHED_GET_PRIORITY_MAX = 335
	SYS_SCHED_GET_PRIORITY_MIN = 336
	SYS_SCHED_RR_GET_INTERVAL  = 337
	SYS_AFS_SYSCALL            = 338
	SYS_UNAME                  = 339
	SYS_NANOSLEEP              = 340
	SYS_MREMAP                 = 341
	SYS_NFSSERVCTL             = 342
	SYS_SETRESUID              = 343
	SYS_GETRESUID              = 344
	SYS_PCICONFIG_READ         = 345
	SYS_PCICONFIG_WRITE        = 346
	SYS_QUERY_MODULE           = 347
	SYS_PRCTL                  = 348
	SYS_PREAD64                = 349
	SYS_PWRITE64               = 350
	SYS_RT_SIGRETURN           = 351
	SYS_RT_SIGACTION           = 352
	SYS_RT_SIGPROCMASK         = 353
	SYS_RT_SIGPENDING          = 354
	SYS_RT_SIGTIMEDWAIT        = 355
	SYS_RT_SIGQUEUEINFO        = 356
	SYS_RT_SIGSUSPEND          = 357
	SYS_SELECT                 = 358
	SYS_GETTIMEOFDAY           = 359
	SYS_SETTIMEOFDAY           = 360
	SYS_GETITIMER              = 361
	SYS_SETITIMER              = 362
	SYS_UTIMES                 = 363
	SYS_GETRUSAGE              = 364
	SYS_WAIT4                  = 365
	SYS_ADJTIMEX               = 366
	SYS_GETCWD                 = 367
	SYS_CAPGET                 = 368
	SYS_CAPSET                 = 369
	SYS_SENDFILE               = 370
	SYS_SETRESGID              = 371
	SYS_GETRESGID              = 372
	SYS_DIPC                   = 373
	SYS_PIVOT_ROOT             = 374
	SYS_MINCORE                = 375
	SYS_PCICONFIG_IOBASE       = 376
	SYS_GETDENTS64             = 377
	SYS_GETTID                 = 378
	SYS_READAHEAD              = 379
	SYS_TKILL                  = 381
	SYS_SETXATTR               = 382
	SYS_LSETXATTR              = 383
	SYS_FSETXATTR              = 384
	SYS_GETXATTR               = 385
	SYS_LGETXATTR              = 386
	SYS_FGETXATTR              = 387
	SYS_LISTXATTR              = 388
	SYS_LLISTXATTR             = 389
	SYS_FLISTXATTR             = 390
	SYS_REMOVEXATTR            = 391
	SYS_LREMOVEXATTR           = 392
	SYS_FREMOVEXATTR           = 393
	SYS_FUTEX                  = 394
	SYS_SCHED_SETAFFINITY      = 395
	SYS_SCHED_GETAFFINITY      = 396
	SYS_TUXCALL                = 397
	SYS_IO_SETUP               = 398
	SYS_IO_DESTROY             = 399
	SYS_IO_GETEVENTS           = 400
	SYS_IO_SUBMIT              = 401
	SYS_IO_CANCEL              = 402
	SYS_IO_PGETEVENTS          = 403
	SYS_RSEQ                   = 404
	SYS_EXIT_GROUP             = 405
	SYS_LOOKUP_DCOOKIE         = 406
	SYS_EPOLL_CREATE           = 407
	SYS_EPOLL_CTL              = 408
	SYS_EPOLL_WAIT             = 409
	SYS_REMAP_FILE_PAGES       = 410
	SYS_SET_TID_ADDRESS        = 411
	SYS_RESTART_SYSCALL        = 412
	SYS_FADVISE64              = 413
	SYS_TIMER_CREATE           = 414
	SYS_TIMER_SETTIME          = 415
	SYS_TIMER_GETTIME          = 416
	SYS_TIMER_GETOVERRUN       = 417
	SYS_TIMER_DELETE           = 418
	SYS_CLOCK_SETTIME          = 419
	SYS_CLOCK_GETTIME          = 420
	SYS_CLOCK_GETRES           = 421
	SYS_CLOCK_NANOSLEEP        = 422
	SYS_SEMTIMEDOP             = 423
	SYS_TGKILL                 = 424
	SYS_STAT64                 = 425
	SYS_LSTAT64                = 426
	SYS_FSTAT64                = 427
	SYS_VSERVER                = 428
	SYS_MBIND                  = 429
	SYS_GET_MEMPOLICY          = 430
	SYS_SET_MEMPOLICY          = 431
	SYS_MQ_OPEN                = 432
	SYS_MQ_UNLINK              = 433
	SYS_MQ_TIMEDSEND           = 434
	SYS_MQ_TIMEDRECEIVE        = 435
	SYS_MQ_NOTIFY              = 436
	SYS_MQ_GETSETATTR          = 437
	SYS_WAITID                 = 438
	SYS_ADD_KEY                = 439
	SYS_REQUEST_KEY            = 440
	SYS_KEYCTL                 = 441
	SYS_IOPRIO_SET             = 442
	SYS_IOPRIO_GET             = 443
	SYS_INOTIFY_INIT           = 444
	SYS_INOTIFY_ADD_WATCH      = 445
	SYS_INOTIFY_RM_WATCH       = 446
	SYS_FDATASYNC              = 447
	SYS_KEXEC_LOAD             = 448
	SYS_MIGRATE_PAGES          = 449
	SYS_OPENAT                 = 450
	SYS_MKDIRAT                = 451
	SYS_MKNODAT                = 452
	SYS_FCHOWNAT               = 453
	SYS_FUTIMESAT              = 454
	SYS_FSTATAT64              = 455
	SYS_UNLINKAT               = 456
	SYS_RENAMEAT               = 457
	SYS_LINKAT                 = 458
	SYS_SYMLINKAT              = 459
	SYS_READLINKAT             = 460
	SYS_FCHMODAT               = 461
	SYS_FACCESSAT              = 462
	SYS_PSELECT6               = 463
	SYS_PPOLL                  = 464
	SYS_UNSHARE                = 465
	SYS_SET_ROBUST_LIST        = 466
	SYS_GET_ROBUST_LIST        = 467
	SYS_SPLICE                 = 468
	SYS_SYNC_FILE_RANGE        = 469
	SYS_TEE                    = 470
	SYS_VMSPLICE               = 471
	SYS_MOVE_PAGES             = 472
	SYS_GETCPU                 = 473
	SYS_EPOLL_PWAIT            = 474
	SYS_UTIMENSAT              = 475
	SYS_SIGNALFD               = 476
	SYS_TIMERFD                = 477
	SYS_EVENTFD                = 478
	SYS_RECVMMSG               = 479
	SYS_FALLOCATE              = 480
	SYS_TIMERFD_CREATE         = 481
	SYS_TIMERFD_SETTIME        = 482
	SYS_TIMERFD_GETTIME        = 483
	SYS_SIGNALFD4              = 484
	SYS_EVENTFD2               = 485
	SYS_EPOLL_CREATE1          = 486
	SYS_DUP3                   = 487
	SYS_PIPE2                  = 488
	SYS_INOTIFY_INIT1          = 489
	SYS_PREADV                 = 490
	SYS_PWRITEV                = 491
	SYS_RT_TGSIGQUEUEINFO      = 492
	SYS_PERF_EVENT_OPEN        = 493
	SYS_FANOTIFY_INIT          = 494
	SYS_FANOTIFY_MARK          = 495
	SYS_PRLIMIT64              = 496
	SYS_NAME_TO_HANDLE_AT      = 497
	SYS_OPEN_BY_HANDLE_AT      = 498
	SYS_CLOCK_ADJTIME          = 499
	SYS_SYNCFS                 = 500
	SYS_SETNS                  = 501
	SYS_ACCEPT4                = 502
	SYS_SENDMMSG               = 503
	SYS_PROCESS_VM_READV       = 504
	SYS_PROCESS_VM_WRITEV      = 505
	SYS_KCMP                   = 506
	SYS_FINIT_MODULE           = 507
	SYS_SCHED_SETATTR          = 508
	SYS_SCHED_GETATTR          = 509
	SYS_RENAMEAT2              = 510
	SYS_GETRANDOM              = 511
	SYS_MEMFD_CREATE           = 512
	SYS_EXECVEAT               = 513
	SYS_SECCOMP                = 514
	SYS_COPY_FILE_RANGE        = 515
	SYS_PREADV2                = 516
	SYS_PWRITEV2               = 517
	SYS_STATX                  = 518
)
