package vendor_dep

//go:generate mockgen -package vendor_dep -destination mock.go github.com/johnpena/mock/mockgen/internal/tests/vendor_dep VendorsDep
//go:generate mockgen -destination source_mock_package/mock.go -source=vendor_dep.go
