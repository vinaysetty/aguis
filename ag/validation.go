package ag

func (x Enrollment_Status) Valid() bool {
	return Enrollment_None <= x && x <= Enrollment_Teacher
}

func (m *EnrollmentRequest) Valid() bool {
	if m != nil {
		return m.GetEnrolled().Valid() && m.GetCourseID() > 0 && m.GetUserID() > 0
	}
	return false
}
