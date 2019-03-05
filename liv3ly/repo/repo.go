package irepo

import "github.com/ngdlong91/cucumbers/errs"

type LoadMode int

type ActionMode int

const (
	Insert ActionMode = 1
	Update ActionMode = 2
	Delete ActionMode = 3
	Other  ActionMode = 4
)

const (
	LoadAll    LoadMode = 1
	LoadByPage LoadMode = 2
)

type ActionRepo interface {
	Build() ierr.CustomError
	Execute() ierr.CustomError

	SetMode(mode ActionMode)
	Mode() ActionMode
}

type Repo interface {
	IsCached() bool
	BuildForCache() ierr.CustomError
	LoadFromCache() ierr.CustomError
	BuildForStorage() ierr.CustomError
	LoadFromStorage() ierr.CustomError
}

type DetailRepo interface {
	Repo
	IsExist() (bool, ierr.CustomError)
}

type ListRepo interface {
	Repo
	SetMode(mode LoadMode)
	Mode() LoadMode
}

func Prepare(c Repo) ierr.CustomError {
	if c.IsCached() {
		return c.BuildForCache()
	} else {
		return c.BuildForStorage()
	}
}

func Execute(c Repo) ierr.CustomError {
	if c.IsCached() {
		return c.LoadFromCache()
	} else {
		return c.LoadFromStorage()
	}
}

func ExecuteSelect(c DetailRepo) ierr.CustomError {
	var err ierr.CustomError
	if c.IsCached() {
		err = c.BuildForCache()
		if !err.IsSuccess() {
			return err
		}
		err = c.LoadFromCache()
		if !err.IsSuccess() {
			return err
		}
	} else {
		err = c.BuildForStorage()
		if !err.IsSuccess() {
			return err
		}
		err = c.LoadFromStorage()
		if !err.IsSuccess() {
			return err
		}
	}
	return err
}
func ExecuteMultiSelect(c ListRepo) ierr.CustomError {
	var err ierr.CustomError
	if c.IsCached() {
		err = c.BuildForCache()
		if !err.IsSuccess() {
			return err
		}
		err = c.LoadFromCache()
		if !err.IsSuccess() {
			return err
		}
	} else {
		err = c.BuildForStorage()
		if !err.IsSuccess() {
			return err
		}
		err = c.LoadFromStorage()
		if !err.IsSuccess() {
			return err
		}
	}
	return err.Success()
}
func ExecuteAction(c ActionRepo) ierr.CustomError {
	err := c.Build()
	if !err.IsSuccess() {
		return err
	}
	err = c.Execute()
	if !err.IsSuccess() {
		return err
	}
	return err.Success()
}
