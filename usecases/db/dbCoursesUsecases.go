package usecases

import (
	"akademia-api/entities"
	"errors"
)

func (d *DbUsecase) GetAllDrafts() ([]entities.CoursesClassesOutput, error) {
	courses, err := d.Repository.Content.GetDraftCourses()
	if err != nil {
		return nil, err
	}

	ids := make([]string, len(courses))
	for i, course := range courses {
		ids[i] = course.ID
	}

	classes, err := d.Repository.Content.GetClassesByCoursesID(ids)
	if err != nil {
		return nil, err
	}

	var output []entities.CoursesClassesOutput

	output = append(output, entities.CoursesClassesOutput{
		Courses: courses,
		Classes: classes,
	})

	return output, nil
}

func (d *DbUsecase) UpdateCourse(course *entities.CourseInput) error {
	if course.ID == "" {
		return errors.New("course ID cannot be empty")
	}

	var courseUpdate entities.Courses
	courseUpdate.ID = course.ID
	courseUpdate.Name = course.Name
	courseUpdate.Description = course.Description
	courseUpdate.Status = course.Status
	courseUpdate.ProductID = course.ProductID

	if err := d.Repository.Content.UpdateCourse(&courseUpdate); err != nil {
		return err
	}
	if len(course.Classes) > 0 {
		d.Repository.Content.UpdateClasses(course.Classes)
	}

	return nil
}

func (d *DbUsecase) CreateFullCourse(course entities.CourseInput) error {
	var courseInput entities.Courses
	courseInput.Name = course.Name
	courseInput.Description = course.Description
	courseInput.Status = course.Status

	if err := d.Repository.Content.CreateCourse(courseInput); err != nil {
		return err
	}
	if len(course.Classes) > 0 {
		for i := range course.Classes {
			course.Classes[i].CourseID = courseInput.ID
		}
		if err := d.Repository.Content.CreateClasses(course.Classes); err != nil {
			return err
		}
	}

	return nil
}

func (d *DbUsecase) GetFullCourseInfo(courseID string) (*entities.CourseClassesOutput, error) {
	course, err := d.Repository.Content.GetCourseByID(courseID)
	if err != nil {
		return nil, err
	}

	classes, err := d.Repository.Content.GetAllClassesByCourseID(courseID)
	if err != nil {
		return nil, err
	}

	return &entities.CourseClassesOutput{
		Course:  *course,
		Classes: classes,
	}, nil
}
