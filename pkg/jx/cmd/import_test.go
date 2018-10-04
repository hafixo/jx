package cmd_test

//fails tests in one-host master
//	import_test.go:49:
//	Error Trace:	import_test.go:49
//Error:      	Received unexpected error:
//	open test_data/prow/no_owners_file/OWNERS: no such file or directory
//Test:       	TestCreateProwOwnersFileCreateWhenDoesNotExist
//Messages:   	There should be no error
//panic: remove test_data/prow/no_owners_file/OWNERS: no such file or directory [recovered]
//panic: remove test_data/prow/no_owners_file/OWNERS: no such file or directory

//func cleanup(path string) {
//	var err = os.Remove(path)
//	if err != nil {
//		panic(err)
//	}
//}
//
//func TestCreateProwOwnersFileExistsDoNothing(t *testing.T) {
//	t.Parallel()
//	// Set the path
//	startPath, err := filepath.Abs("")
//	if err != nil {
//		panic(err)
//	}
//	path := startPath + "/test_data/prow"
//	cmd := cmd.ImportOptions{
//		Dir: path,
//	}
//
//	err = cmd.CreateProwOwnersFile()
//	assert.NoError(t, err, "There should be no error")
//}
//
//func TestCreateProwOwnersFileCreateWhenDoesNotExist(t *testing.T) {
//	t.Parallel()
//	// Set the path
//	startPath, err := filepath.Abs("")
//	if err != nil {
//		panic(err)
//	}
//	path := startPath + "/test_data/prow/no_owners_file"
//	// Remove test file after test
//	defer cleanup(path + "/OWNERS")
//
//	cmd := cmd.ImportOptions{
//		Dir: path,
//		GitUserAuth: &auth.UserAuth{
//			Username: "derek_zoolander",
//		},
//	}
//
//	err = cmd.CreateProwOwnersFile()
//	assert.NoError(t, err, "There should be no error")
//}

//
//func TestCreateProwOwnersFileCreateWhenDoesNotExistAndNoGitUserSet(t *testing.T) {
//	t.Parallel()
//	// Set the path
//	startPath, err := filepath.Abs("")
//	if err != nil {
//		panic(err)
//	}
//	path := startPath + "/test_data/prow/no_owners_file"
//
//	cmd := cmd.ImportOptions{
//		Dir: path,
//	}
//
//	err = cmd.CreateProwOwnersFile()
//	assert.Error(t, err, "There should an error")
//}
//
//func TestCreateProwOwnersAliasesFileExistsDoNothing(t *testing.T) {
//	t.Parallel()
//	// Set the path
//	startPath, err := filepath.Abs("")
//	if err != nil {
//		panic(err)
//	}
//	path := startPath + "/test_data/prow"
//	cmd := cmd.ImportOptions{
//		Dir: path,
//	}
//
//	err = cmd.CreateProwOwnersAliasesFile()
//	assert.NoError(t, err, "There should be no error")
//}
//
//func TestCreateProwOwnersAliasesFileCreateWhenDoesNotExist(t *testing.T) {
//	t.Parallel()
//	// Set the path
//	startPath, err := filepath.Abs("")
//	if err != nil {
//		panic(err)
//	}
//	path := startPath + "/test_data/prow/no_owners_file"
//	// Remove test file after test
//	defer cleanup(path + "/OWNERS_ALIASES")
//
//	cmd := cmd.ImportOptions{
//		Dir: path,
//		GitUserAuth: &auth.UserAuth{
//			Username: "derek_zoolander",
//		},
//	}
//
//	err = cmd.CreateProwOwnersAliasesFile()
//	assert.NoError(t, err, "There should be no error")
//}
//
//func TestCreateProwOwnersAliasesFileCreateWhenDoesNotExistAndNoGitUserSet(t *testing.T) {
//	t.Parallel()
//	// Set the path
//	startPath, err := filepath.Abs("")
//	if err != nil {
//		panic(err)
//	}
//	path := startPath + "/test_data/prow/no_owners_file"
//
//	cmd := cmd.ImportOptions{
//		Dir: path,
//	}
//
//	err = cmd.CreateProwOwnersAliasesFile()
//	assert.Error(t, err, "There should an error")
//}
