// 招聘市场服务 
package elemeOpenApi

// 简历回传
// resume 简历信息
func (recruitment *Recruitment) UploadResume(resume_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["resume"] = resume_
	return APIInterface(recruitment.config, "eleme.recruitment.uploadResume", params)
}

// 职位状态变更
// positionId 职位id
// status 变更状态
func (recruitment *Recruitment) UpdateJobPositionStatus(positionId_ string, status_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["positionId"] = positionId_
	params["status"] = status_
	return APIInterface(recruitment.config, "eleme.recruitment.updateJobPositionStatus", params)
}

// 简历状态变更
// resumeId 职位id
// status 变更状态
func (recruitment *Recruitment) UpdateResumeStatus(resumeId_ string, status_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["resumeId"] = resumeId_
	params["status"] = status_
	return APIInterface(recruitment.config, "eleme.recruitment.updateResumeStatus", params)
}

