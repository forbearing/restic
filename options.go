package restic

/*
TODO:
	1.limit "restic dump flags --archive value range, only support "tar" and "zip"
	2."restic dump 4e63f524 / > file.tar " need specific output file path.
	3."restic find" 的某些选项改成 time.Duration 类型
	4.restic forget 的某些选项改成 time.Duration 类型
	4.restic mount 的某些选项改成 time.Duration 类型
	- restic backup 的某些选项改成 time.Duration 类型
	- 考虑子命令
	- 注释都加一个 Usage

NOTE:
	- Global flag 放最后面
	- 如果 flag 的值类行为 bool 可以不用写值

all restic sub-command man page
	restic-backup
	restic-cache
	restic-cat
	restic-check
	restic-copy
	restic-diff				no
	restic-dump
	restic-find
	restic-forget
	restic-generate			no
	restic-init
	restic-key
	restic-list
	restic-ls
	restic-migrate
	restic-mount
	restic-prune
	restic-rebuild-index
	restic-recover
	restic-restore
	restic-self-update
	restic-snapshots
	restic-stats
	restic-tag
	restic-unlock
	restic-version

*/
