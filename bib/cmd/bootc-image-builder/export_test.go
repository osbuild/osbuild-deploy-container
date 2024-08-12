package main

var (
	CanChownInPath                = canChownInPath
	CheckFilesystemCustomizations = checkFilesystemCustomizations
	GetDistroAndRunner            = getDistroAndRunner
	CheckMountpoints              = checkMountpoints
	PartitionTables               = partitionTables
	UpdateFilesystemSizes         = updateFilesystemSizes
)

func MockOsGetuid(new func() int) (restore func()) {
	saved := osGetuid
	osGetuid = new
	return func() {
		osGetuid = saved
	}
}
