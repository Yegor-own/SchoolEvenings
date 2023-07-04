package sqlite

import "backend/src/domain/entity"

func (d SqlxDriver) CreateTag(tag *entity.Tag) error {
	return nil
}

// DeleteTag deletes relations between course and tag NOT TAG
func (d SqlxDriver) DeleteTag(id uint) error {
	return nil
}
