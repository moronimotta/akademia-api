package usecases

func (d *DbUsecase) MarkClassAsCompleted(userID, courseID, classID string) error {
	if err := d.Repository.UserProgress.UpdateClassStatus(userID, courseID, classID); err != nil {
		return err
	}

	if err := d.Repository.UserProgress.UpdateUserCourseProgress(userID, courseID); err != nil {
		return err
	}
	return nil
}
