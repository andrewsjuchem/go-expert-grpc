package service

import (
	"context"
	"io"

	"github.com/andrewsjuchem/go-expert-grpc/internal/database"
	"github.com/andrewsjuchem/go-expert-grpc/internal/pb"
)

type CourseService struct {
	pb.UnimplementedCourseServiceServer
	CourseDB database.Course
}

func NewCourseService(courseDB database.Course) *CourseService {
	return &CourseService{
		CourseDB: courseDB,
	}
}

func (c *CourseService) CreateCourse(ctx context.Context, req *pb.CreateCourseRequest) (*pb.Course, error) {
	course, err := c.CourseDB.Create(req.Name, req.Description, req.CategoryId)
	if err != nil {
		return nil, err
	}

	courseResponse := &pb.Course{
		Id:          course.ID,
		Name:        course.Name,
		Description: course.Description,
		CategoryId:  course.CategoryID,
	}

	return courseResponse, nil
}

func (c *CourseService) ListCourses(ctx context.Context, req *pb.Blank) (*pb.CourseList, error) {
	courses, err := c.CourseDB.FindAll()
	if err != nil {
		return nil, err
	}

	var coursesResponse []*pb.Course

	for _, course := range courses {
		courseResponse := &pb.Course{
			Id:          course.ID,
			Name:        course.Name,
			Description: course.Description,
			CategoryId:  course.CategoryID,
		}
		coursesResponse = append(coursesResponse, courseResponse)
	}

	return &pb.CourseList{Courses: coursesResponse}, nil
}

func (c *CourseService) GetCourse(ctx context.Context, req *pb.CourseGetRequest) (*pb.Course, error) {
	course, err := c.CourseDB.Find(req.Id)
	if err != nil {
		return nil, err
	}

	courseResponse := &pb.Course{
		Id:          course.ID,
		Name:        course.Name,
		Description: course.Description,
		CategoryId:  course.CategoryID,
	}

	return courseResponse, nil
}

func (c *CourseService) CreateCategoryStream(stream pb.CourseService_CreateCourseStreamServer) error {
	courses := &pb.CourseList{}

	for {
		course, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(courses)
		}
		if err != nil {
			return err
		}

		courseResult, err := c.CourseDB.Create(course.Name, course.Description, course.CategoryId)
		if err != nil {
			return err
		}

		courses.Courses = append(courses.Courses, &pb.Course{
			Id:          courseResult.ID,
			Name:        courseResult.Name,
			Description: courseResult.Description,
			CategoryId:  courseResult.CategoryID,
		})
	}
}

func (c *CourseService) CreateCourseStreamBidirectional(stream pb.CourseService_CreateCourseStreamBidirectionalServer) error {
	for {
		course, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		courseResult, err := c.CourseDB.Create(course.Name, course.Description, course.CategoryId)
		if err != nil {
			return err
		}

		err = stream.Send(&pb.Course{
			Id:          courseResult.ID,
			Name:        courseResult.Name,
			Description: courseResult.Description,
			CategoryId:  courseResult.CategoryID,
		})
		if err != nil {
			return err
		}
	}
}
