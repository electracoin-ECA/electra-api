package models

// This function will create the indexes for each of the collections we have present.
// we cannot currently have a global only-once operation to create index (that should be part of
// a CI/CD process) so we will have a create if not present mechanism here so we dont
// create index or even try to create index every time someone starts an application.
func init() {

}
