package repo

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
	Build() errs.CustomError
	Execute() errs.CustomError

	SetMode(mode ActionMode)
	Mode() ActionMode
}

type Repo interface {
	IsCached() bool
	BuildForCache() errs.CustomError
	LoadFromCache() errs.CustomError

	BuildForStorage() errs.CustomError
	LoadFromStorage() errs.CustomError
}

type DetailRepo interface {
	Repo
	IsExist() bool
}

type ListRepo interface {
	Repo
	SetMode(mode LoadMode)
	Mode() LoadMode
}

func ExecuteSelect(c DetailRepo) errs.CustomError {
	var err errs.CustomError
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
		//err = c.BuildForStorage()
		//if !err.IsSuccess() {
		//	return err
		//}
		err = c.LoadFromStorage()
		if !err.IsSuccess() {
			return err
		}
	}
	return err
}
func ExecuteMultiSelect(c ListRepo) errs.CustomError {
	var err errs.CustomError
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

func ExecuteAction(c ActionRepo) errs.CustomError {
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
